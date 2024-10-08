package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilopezhe/terraform-provider-awx/internal/client"
	"github.com/ilopezhe/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &credentialResource{}
	_ resource.ResourceWithConfigure   = &credentialResource{}
	_ resource.ResourceWithImportState = &credentialResource{}
)

// NewCredentialResource is a helper function to simplify the provider implementation.
func NewCredentialResource() resource.Resource {
	return &credentialResource{}
}

type credentialResource struct {
	client   c.Client
	endpoint string
}

func (o *credentialResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/credentials/"
}

func (o *credentialResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_credential"
}

func (o *credentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"credential_type": schema.Int64Attribute{
				Description:   "Specify the type of credential you want to create. Refer to the documentation for details on each type.",
				Sensitive:     false,
				Required:      true,
				Optional:      false,
				Computed:      false,
				PlanModifiers: []planmodifier.Int64{},
				Validators:    []validator.Int64{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this credential.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"inputs": schema.StringAttribute{
				Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(`{}`),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"name": schema.StringAttribute{
				Description:   "Name of this credential.",
				Sensitive:     false,
				Required:      true,
				Optional:      false,
				Computed:      false,
				PlanModifiers: []planmodifier.String{},
				Validators: []validator.String{
					stringvalidator.LengthAtMost(512),
				},
			},
			"organization": schema.Int64Attribute{
				Description: "Inherit permissions from organization roles. If provided on creation, do not give either user or team.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					// exactly_one_of_org_user_team
					int64validator.ExactlyOneOf(
						path.MatchRoot("organization"),
						path.MatchRoot("team"),
						path.MatchRoot("user"),
					),
				},
			},
			// Write only elements
			"team": schema.Int64Attribute{
				Description: "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					// exactly_one_of_org_user_team
					int64validator.ExactlyOneOf(
						path.MatchRoot("organization"),
						path.MatchRoot("team"),
						path.MatchRoot("user"),
					),
				},
			},
			"user": schema.Int64Attribute{
				Description: "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.",
				Sensitive:   false,
				Required:    false,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					// exactly_one_of_org_user_team
					int64validator.ExactlyOneOf(
						path.MatchRoot("organization"),
						path.MatchRoot("team"),
						path.MatchRoot("user"),
					),
				},
			},
			// Data only elements
			"cloud": schema.BoolAttribute{
				Description: "",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this credential.",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"kind": schema.StringAttribute{
				Description: "",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"kubernetes": schema.BoolAttribute{
				Description: "",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"managed": schema.BoolAttribute{
				Description: "",
				Required:    false,
				Optional:    false,
				Computed:    true,
				Sensitive:   false,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o *credentialResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Credential.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *credentialResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state credentialTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Credential
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	bodyRequest.Team = plan.Team.ValueInt64()
	bodyRequest.User = plan.User.ValueInt64()
	tflog.Debug(ctx, "[Credential/Create] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Credential on %s for create", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Credential resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Credential on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	state.Team = types.Int64Value(plan.Team.ValueInt64())
	state.User = types.Int64Value(plan.User.ValueInt64())

	if err = hookCredential(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Credential",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *credentialResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state credentialTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for Credential
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Credential on %s for read", endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Credential from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		if o.client.IsResourceNotFound(err) {
			// Remove the resource from the state to signal that it needs to be recreated
			response.State.RemoveResource(ctx)
			return
		}
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Credential on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookCredential(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Credential",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *credentialResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state credentialTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Credential
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[Credential/Update] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	bodyRequest.Team = plan.Team.ValueInt64()
	bodyRequest.User = plan.User.ValueInt64()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Credential on %s for update", endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Credential resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Credential on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	state.Team = types.Int64Value(plan.Team.ValueInt64())
	state.User = types.Int64Value(plan.User.ValueInt64())

	if err = hookCredential(ctx, ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on Credential",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *credentialResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state credentialTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Credential
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Credential on %s for delete", endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Credential
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Credential on %s", endpoint),
			err.Error(),
		)
		return
	}
}
