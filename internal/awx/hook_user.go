package awx

import (
	"context"
	"fmt"

	"github.com/ilopezhe/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookUser(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *userTerraformModel) (err error) { // nolint:unused
	if source == hooks.SourceResource && (state == nil || orig == nil) && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == hooks.SourceResource && callee == hooks.CalleeCreate ||
		(state.Password.Equal(types.StringValue("$encrypted$")) &&
			(source == hooks.SourceResource && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead)) &&
			!orig.Password.IsNull()) {
		state.Password = orig.Password
	}

	return nil
}
