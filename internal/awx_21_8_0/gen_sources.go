package awx_21_8_0

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	ApiVersion string = "21.8.0"
)

// DataSources is a helper function to return all defined data sources
func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewInventoryDataSource,
		NewInventoryObjectRolesDataSource,
		NewProjectDataSource,
		NewProjectObjectRolesDataSource,
		NewOrganizationDataSource,
		NewOrganizationObjectRolesDataSource,
		NewTeamDataSource,
		NewTeamObjectRolesDataSource,
		NewExecutionEnvironmentDataSource,
		NewCredentialDataSource,
		NewCredentialObjectRolesDataSource,
		NewCredentialTypeDataSource,
		NewUserDataSource,
		NewCredentialInputSourceDataSource,
		NewLabelDataSource,
		NewApplicationDataSource,
		NewScheduleDataSource,
		NewInstanceGroupDataSource,
		NewGroupDataSource,
		NewTokensDataSource,
		NewAdHocCommandDataSource,
		NewMeDataSource,
		NewSettingsUIDataSource,
		NewSettingsOpenIDConnectDataSource,
		NewSettingsMiscLoggingDataSource,
		NewSettingsMiscSystemDataSource,
		NewSettingsMiscAuthenticationDataSource,
		NewSettingsJobsDataSource,
		NewSettingsAuthAzureADOauth2DataSource,
		NewSettingsAuthGoogleOauth2DataSource,
		NewSettingsAuthGithubDataSource,
		NewSettingsAuthGithubEnterpriseDataSource,
		NewSettingsAuthGithubEnterpriseOrgDataSource,
		NewSettingsAuthGithubEnterpriseTeamDataSource,
		NewSettingsAuthGithubOrgDataSource,
		NewSettingsAuthGithubTeamDataSource,
		NewSettingsAuthSAMLDataSource,
		NewSettingsAuthLDAPDataSource,
		NewNotificationTemplateDataSource,
		NewWorkflowJobTemplateDataSource,
		NewWorkflowJobTemplateObjectRolesDataSource,
		NewInventorySourceDataSource,
		NewJobTemplateDataSource,
		NewJobTemplateObjectRolesDataSource,
		NewHostDataSource,
		NewHostObjectRolesDataSource,
	}
}

// Resources is a helper function to return all defined resources
func Resources() []func() resource.Resource {
	return []func() resource.Resource{
		NewInventoryResource,
		NewProjectResource,
		NewOrganizationResource,
		NewOrganizationAssociateDisassociateInstanceGroupResource,
		NewOrganizationAssociateDisassociateGalaxyCredentialResource,
		NewTeamResource,
		NewTeamAssociateDisassociateRoleResource,
		NewExecutionEnvironmentResource,
		NewCredentialResource,
		NewCredentialTypeResource,
		NewUserResource,
		NewUserAssociateDisassociateRoleResource,
		NewCredentialInputSourceResource,
		NewLabelResource,
		NewApplicationResource,
		NewScheduleResource,
		NewInstanceGroupResource,
		NewGroupResource,
		NewTokensResource,
		NewAdHocCommandResource,
		NewSettingsUIResource,
		NewSettingsOpenIDConnectResource,
		NewSettingsMiscLoggingResource,
		NewSettingsMiscSystemResource,
		NewSettingsMiscAuthenticationResource,
		NewSettingsJobsResource,
		NewSettingsAuthAzureADOauth2Resource,
		NewSettingsAuthGoogleOauth2Resource,
		NewSettingsAuthGithubResource,
		NewSettingsAuthGithubEnterpriseResource,
		NewSettingsAuthGithubEnterpriseOrgResource,
		NewSettingsAuthGithubEnterpriseTeamResource,
		NewSettingsAuthGithubOrgResource,
		NewSettingsAuthGithubTeamResource,
		NewSettingsAuthSAMLResource,
		NewSettingsAuthLDAPResource,
		NewNotificationTemplateResource,
		NewWorkflowJobTemplateResource,
		NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewWorkflowJobTemplateSurveyResource,
		NewInventorySourceResource,
		NewJobTemplateResource,
		NewJobTemplateAssociateDisassociateCredentialResource,
		NewJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewJobTemplateSurveyResource,
		NewHostResource,
		NewHostAssociateDisassociateGroupResource,
	}
}
