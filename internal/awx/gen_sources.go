package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	ApiVersion string = "24.6.1"
)

// DataSources is a helper function to return all defined data sources
func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAdHocCommandDataSource,
		NewApplicationDataSource,
		NewConstructedInventoriesDataSource,
		NewConstructedInventoriesObjectRolesDataSource,
		NewCredentialDataSource,
		NewCredentialInputSourceDataSource,
		NewCredentialObjectRolesDataSource,
		NewCredentialTypeDataSource,
		NewExecutionEnvironmentDataSource,
		NewGroupDataSource,
		NewHostDataSource,
		NewHostObjectRolesDataSource,
		NewInstanceGroupDataSource,
		NewInstanceGroupObjectRolesDataSource,
		NewInventoryDataSource,
		NewInventoryObjectRolesDataSource,
		NewInventorySourceDataSource,
		NewJobTemplateDataSource,
		NewJobTemplateObjectRolesDataSource,
		NewLabelDataSource,
		NewMeDataSource,
		NewNotificationTemplateDataSource,
		NewOrganizationDataSource,
		NewOrganizationObjectRolesDataSource,
		NewProjectDataSource,
		NewProjectObjectRolesDataSource,
		NewScheduleDataSource,
		NewSettingsAuthAzureADOauth2DataSource,
		NewSettingsAuthGithubDataSource,
		NewSettingsAuthGithubEnterpriseDataSource,
		NewSettingsAuthGithubEnterpriseOrgDataSource,
		NewSettingsAuthGithubEnterpriseTeamDataSource,
		NewSettingsAuthGithubOrgDataSource,
		NewSettingsAuthGithubTeamDataSource,
		NewSettingsAuthGoogleOauth2DataSource,
		NewSettingsAuthLDAPDataSource,
		NewSettingsAuthSAMLDataSource,
		NewSettingsJobsDataSource,
		NewSettingsMiscAuthenticationDataSource,
		NewSettingsMiscLoggingDataSource,
		NewSettingsMiscSystemDataSource,
		NewSettingsOpenIDConnectDataSource,
		NewSettingsUIDataSource,
		NewTeamDataSource,
		NewTeamObjectRolesDataSource,
		NewTokensDataSource,
		NewUserDataSource,
		NewWorkflowJobTemplateDataSource,
		NewWorkflowJobTemplateObjectRolesDataSource,
	}
}

// Resources is a helper function to return all defined resources
func Resources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAdHocCommandResource,
		NewApplicationResource,
		NewConstructedInventoriesResource,
		NewCredentialResource,
		NewCredentialInputSourceResource,
		NewCredentialTypeResource,
		NewExecutionEnvironmentResource,
		NewGroupResource,
		NewHostResource,
		NewHostAssociateDisassociateGroupResource,
		NewInstanceGroupResource,
		NewInventoryResource,
		NewInventorySourceResource,
		NewJobTemplateResource,
		NewJobTemplateAssociateDisassociateCredentialResource,
		NewJobTemplateAssociateDisassociateInstanceGroupResource,
		NewJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewJobTemplateSurveyResource,
		NewLabelResource,
		NewNotificationTemplateResource,
		NewOrganizationResource,
		NewOrganizationAssociateDisassociateGalaxyCredentialResource,
		NewOrganizationAssociateDisassociateInstanceGroupResource,
		NewProjectResource,
		NewScheduleResource,
		NewSettingsAuthAzureADOauth2Resource,
		NewSettingsAuthGithubResource,
		NewSettingsAuthGithubEnterpriseResource,
		NewSettingsAuthGithubEnterpriseOrgResource,
		NewSettingsAuthGithubEnterpriseTeamResource,
		NewSettingsAuthGithubOrgResource,
		NewSettingsAuthGithubTeamResource,
		NewSettingsAuthGoogleOauth2Resource,
		NewSettingsAuthLDAPResource,
		NewSettingsAuthSAMLResource,
		NewSettingsJobsResource,
		NewSettingsMiscAuthenticationResource,
		NewSettingsMiscLoggingResource,
		NewSettingsMiscSystemResource,
		NewSettingsOpenIDConnectResource,
		NewSettingsUIResource,
		NewTeamResource,
		NewTeamAssociateDisassociateRoleResource,
		NewTokensResource,
		NewUserResource,
		NewUserAssociateDisassociateRoleResource,
		NewWorkflowJobTemplateResource,
		NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewWorkflowJobTemplateSurveyResource,
	}
}
