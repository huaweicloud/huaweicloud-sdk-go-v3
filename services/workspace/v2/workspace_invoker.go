package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

type ListAccessAddressBackupConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessAddressBackupConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessAddressBackupConfigInvoker) Invoke() (*model.ListAccessAddressBackupConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessAddressBackupConfigResponse), nil
	}
}

type UpdateAccessAddressBackupConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessAddressBackupConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessAddressBackupConfigInvoker) Invoke() (*model.UpdateAccessAddressBackupConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessAddressBackupConfigResponse), nil
	}
}

type BatchDeleteAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAccessPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAccessPoliciesInvoker) Invoke() (*model.BatchDeleteAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAccessPoliciesResponse), nil
	}
}

type CreateAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccessPolicyInvoker) Invoke() (*model.CreateAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessPolicyResponse), nil
	}
}

type ExportIpTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportIpTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportIpTemplateInvoker) Invoke() (*model.ExportIpTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportIpTemplateResponse), nil
	}
}

type ImportIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportIpInvoker) Invoke() (*model.ImportIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportIpResponse), nil
	}
}

type ListAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessPoliciesInvoker) Invoke() (*model.ListAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPoliciesResponse), nil
	}
}

type ListAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPolicyObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessPolicyObjectsInvoker) Invoke() (*model.ListAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPolicyObjectsResponse), nil
	}
}

type UpdateAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessPolicyInvoker) Invoke() (*model.UpdateAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessPolicyResponse), nil
	}
}

type UpdateAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessPolicyObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessPolicyObjectsInvoker) Invoke() (*model.UpdateAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessPolicyObjectsResponse), nil
	}
}

type CreateAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgenciesInvoker) Invoke() (*model.CreateAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgenciesResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type ListAlarmStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmStatisticsInvoker) Invoke() (*model.ListAlarmStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmStatisticsResponse), nil
	}
}

type ListAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmsInvoker) Invoke() (*model.ListAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmsResponse), nil
	}
}

type BatchDeleteAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppsInvoker) Invoke() (*model.BatchDeleteAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppsResponse), nil
	}
}

type BatchDeleteJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) Invoke() (*model.BatchDeleteJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsResponse), nil
	}
}

type BatchDisableAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisableAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisableAppsInvoker) Invoke() (*model.BatchDisableAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisableAppsResponse), nil
	}
}

type BatchEnableAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableAppsInvoker) Invoke() (*model.BatchEnableAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableAppsResponse), nil
	}
}

type BatchInstallAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchInstallAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchInstallAppsInvoker) Invoke() (*model.BatchInstallAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchInstallAppsResponse), nil
	}
}

type BatchUpdateAppAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateAppAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateAppAuthorizationsInvoker) Invoke() (*model.BatchUpdateAppAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateAppAuthorizationsResponse), nil
	}
}

type CreateAndAuthorizeBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAndAuthorizeBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAndAuthorizeBucketInvoker) Invoke() (*model.CreateAndAuthorizeBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAndAuthorizeBucketResponse), nil
	}
}

type CreateBucketCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBucketCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBucketCredentialInvoker) Invoke() (*model.CreateBucketCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBucketCredentialResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type InstallAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InstallAppInvoker) Invoke() (*model.InstallAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallAppResponse), nil
	}
}

type ListAppAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppAuthorizationsInvoker) Invoke() (*model.ListAppAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppAuthorizationsResponse), nil
	}
}

type ListAppCatalogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppCatalogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppCatalogsInvoker) Invoke() (*model.ListAppCatalogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppCatalogsResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type RetryJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryJobsInvoker) Invoke() (*model.RetryJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryJobsResponse), nil
	}
}

type UpdateAppAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppAuthorizationsInvoker) Invoke() (*model.UpdateAppAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppAuthorizationsResponse), nil
	}
}

type UpdateUploadedAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUploadedAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUploadedAppInvoker) Invoke() (*model.UpdateUploadedAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUploadedAppResponse), nil
	}
}

type UploadAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAppInvoker) Invoke() (*model.UploadAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAppResponse), nil
	}
}

type AddRestrictedRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRestrictedRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRestrictedRuleInvoker) Invoke() (*model.AddRestrictedRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRestrictedRuleResponse), nil
	}
}

type BatchDeleteAppRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppRulesInvoker) Invoke() (*model.BatchDeleteAppRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppRulesResponse), nil
	}
}

type CreateAppRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppRuleInvoker) Invoke() (*model.CreateAppRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppRuleResponse), nil
	}
}

type DeleteAppRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppRuleInvoker) Invoke() (*model.DeleteAppRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppRuleResponse), nil
	}
}

type DeleteRestrictedRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRestrictedRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRestrictedRuleInvoker) Invoke() (*model.DeleteRestrictedRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRestrictedRuleResponse), nil
	}
}

type DisableRuleRestrictionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableRuleRestrictionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableRuleRestrictionInvoker) Invoke() (*model.DisableRuleRestrictionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableRuleRestrictionResponse), nil
	}
}

type EnableRuleRestrictionInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableRuleRestrictionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableRuleRestrictionInvoker) Invoke() (*model.EnableRuleRestrictionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableRuleRestrictionResponse), nil
	}
}

type ListAppRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppRuleInvoker) Invoke() (*model.ListAppRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppRuleResponse), nil
	}
}

type ListRestrictedRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestrictedRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRestrictedRuleInvoker) Invoke() (*model.ListRestrictedRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestrictedRuleResponse), nil
	}
}

type SetRuleRestrictionInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRuleRestrictionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRuleRestrictionInvoker) Invoke() (*model.SetRuleRestrictionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRuleRestrictionResponse), nil
	}
}

type ShowRuleRestrictionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRuleRestrictionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRuleRestrictionInvoker) Invoke() (*model.ShowRuleRestrictionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRuleRestrictionResponse), nil
	}
}

type UpdateAppRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppRuleInvoker) Invoke() (*model.UpdateAppRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppRuleResponse), nil
	}
}

type DownloadMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadMetadataInvoker) Invoke() (*model.DownloadMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadMetadataResponse), nil
	}
}

type ShowAssistAuthConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssistAuthConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssistAuthConfigInvoker) Invoke() (*model.ShowAssistAuthConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssistAuthConfigResponse), nil
	}
}

type ShowAssistAuthConfigApplyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssistAuthConfigApplyObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssistAuthConfigApplyObjectsInvoker) Invoke() (*model.ShowAssistAuthConfigApplyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssistAuthConfigApplyObjectsResponse), nil
	}
}

type ShowAuthConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthConfigInvoker) Invoke() (*model.ShowAuthConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthConfigResponse), nil
	}
}

type UpdateAssistAuthConfigApplyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssistAuthConfigApplyObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssistAuthConfigApplyObjectsInvoker) Invoke() (*model.UpdateAssistAuthConfigApplyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssistAuthConfigApplyObjectsResponse), nil
	}
}

type UpdateAssistAuthMethodConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssistAuthMethodConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssistAuthMethodConfigInvoker) Invoke() (*model.UpdateAssistAuthMethodConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssistAuthMethodConfigResponse), nil
	}
}

type UpdateAuthMethodConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuthMethodConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuthMethodConfigInvoker) Invoke() (*model.UpdateAuthMethodConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuthMethodConfigResponse), nil
	}
}

type ValidateConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateConfigInvoker) Invoke() (*model.ValidateConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateConfigResponse), nil
	}
}

type ListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailabilityZonesInvoker) Invoke() (*model.ListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZonesResponse), nil
	}
}

type ListAzsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAzsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAzsInvoker) Invoke() (*model.ListAzsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAzsResponse), nil
	}
}

type ShowAzDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAzDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAzDetailsInvoker) Invoke() (*model.ShowAzDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAzDetailsResponse), nil
	}
}

type CreateCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertInvoker) Invoke() (*model.CreateCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertResponse), nil
	}
}

type DeleteCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCertInvoker) Invoke() (*model.DeleteCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertResponse), nil
	}
}

type ExportCertCrlFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertCrlFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportCertCrlFileInvoker) Invoke() (*model.ExportCertCrlFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertCrlFileResponse), nil
	}
}

type ExportCertCsrFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertCsrFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportCertCsrFileInvoker) Invoke() (*model.ExportCertCsrFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertCsrFileResponse), nil
	}
}

type ExportCertPemFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertPemFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportCertPemFileInvoker) Invoke() (*model.ExportCertPemFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertPemFileResponse), nil
	}
}

type ImportCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportCertInvoker) Invoke() (*model.ImportCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCertResponse), nil
	}
}

type ListCertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertsInvoker) Invoke() (*model.ListCertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertsResponse), nil
	}
}

type SetCertStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetCertStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetCertStatusInvoker) Invoke() (*model.SetCertStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetCertStatusResponse), nil
	}
}

type ShowCertDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertDetailInvoker) Invoke() (*model.ShowCertDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertDetailResponse), nil
	}
}

type CheckDesktopImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDesktopImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDesktopImagesInvoker) Invoke() (*model.CheckDesktopImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDesktopImagesResponse), nil
	}
}

type CheckSysprepInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckSysprepInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckSysprepInfoInvoker) Invoke() (*model.CheckSysprepInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckSysprepInfoResponse), nil
	}
}

type ExportUserLoginInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserLoginInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserLoginInfoNewInvoker) Invoke() (*model.ExportUserLoginInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserLoginInfoNewResponse), nil
	}
}

type ListHistoryOnlineInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryOnlineInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryOnlineInfoNewInvoker) Invoke() (*model.ListHistoryOnlineInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryOnlineInfoNewResponse), nil
	}
}

type ListInstancesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesStatusInvoker) Invoke() (*model.ListInstancesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesStatusResponse), nil
	}
}

type ListLoginRecordsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginRecordsNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginRecordsNewInvoker) Invoke() (*model.ListLoginRecordsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginRecordsNewResponse), nil
	}
}

type ExportUserConnectionNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserConnectionNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserConnectionNewInvoker) Invoke() (*model.ExportUserConnectionNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserConnectionNewResponse), nil
	}
}

type AttachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachInstancesInvoker) Invoke() (*model.AttachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachInstancesResponse), nil
	}
}

type BatchAssociateInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAssociateInstancesInvoker) Invoke() (*model.BatchAssociateInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateInstancesResponse), nil
	}
}

type BatchAttachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAttachInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAttachInstancesInvoker) Invoke() (*model.BatchAttachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAttachInstancesResponse), nil
	}
}

type BatchChangeDesktopNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeDesktopNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchChangeDesktopNetworkInvoker) Invoke() (*model.BatchChangeDesktopNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeDesktopNetworkResponse), nil
	}
}

type BatchDeleteDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDesktopsInvoker) Invoke() (*model.BatchDeleteDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopsResponse), nil
	}
}

type BatchDetachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDetachInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDetachInstancesInvoker) Invoke() (*model.BatchDetachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDetachInstancesResponse), nil
	}
}

type BatchInstallAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchInstallAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchInstallAgentInvoker) Invoke() (*model.BatchInstallAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchInstallAgentResponse), nil
	}
}

type BatchLogoffDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchLogoffDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchLogoffDesktopsInvoker) Invoke() (*model.BatchLogoffDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchLogoffDesktopsResponse), nil
	}
}

type BatchRebuildDesktopsSystemDiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebuildDesktopsSystemDiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRebuildDesktopsSystemDiskInvoker) Invoke() (*model.BatchRebuildDesktopsSystemDiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebuildDesktopsSystemDiskResponse), nil
	}
}

type BatchRunDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRunDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRunDesktopsInvoker) Invoke() (*model.BatchRunDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRunDesktopsResponse), nil
	}
}

type CancelRemoteAssistanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRemoteAssistanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelRemoteAssistanceInvoker) Invoke() (*model.CancelRemoteAssistanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRemoteAssistanceResponse), nil
	}
}

type ChangeDesktopNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDesktopNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDesktopNetworkInvoker) Invoke() (*model.ChangeDesktopNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDesktopNetworkResponse), nil
	}
}

type ChangeDesktopToImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDesktopToImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDesktopToImageInvoker) Invoke() (*model.ChangeDesktopToImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDesktopToImageResponse), nil
	}
}

type ChangeUserPrivilegeGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeUserPrivilegeGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeUserPrivilegeGroupInvoker) Invoke() (*model.ChangeUserPrivilegeGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeUserPrivilegeGroupResponse), nil
	}
}

type CreateDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopInvoker) Invoke() (*model.CreateDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopResponse), nil
	}
}

type CreateRemoteAssistanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRemoteAssistanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRemoteAssistanceInvoker) Invoke() (*model.CreateRemoteAssistanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRemoteAssistanceResponse), nil
	}
}

type DeleteDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesktopInvoker) Invoke() (*model.DeleteDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopResponse), nil
	}
}

type DetachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachInstancesInvoker) Invoke() (*model.DetachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachInstancesResponse), nil
	}
}

type ListAgentsInstallConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentsInstallConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentsInstallConditionInvoker) Invoke() (*model.ListAgentsInstallConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentsInstallConditionResponse), nil
	}
}

type ListDesktopActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopActionsInvoker) Invoke() (*model.ListDesktopActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopActionsResponse), nil
	}
}

type ListDesktopDetachInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopDetachInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopDetachInfoInvoker) Invoke() (*model.ListDesktopDetachInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopDetachInfoResponse), nil
	}
}

type ListDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsInvoker) Invoke() (*model.ListDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsResponse), nil
	}
}

type ListDesktopsConnectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsConnectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsConnectStatusInvoker) Invoke() (*model.ListDesktopsConnectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsConnectStatusResponse), nil
	}
}

type ListDesktopsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsDetailInvoker) Invoke() (*model.ListDesktopsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsDetailResponse), nil
	}
}

type RegisterDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterDomainInvoker) Invoke() (*model.RegisterDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterDomainResponse), nil
	}
}

type ResizeDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeDesktopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeDesktopInvoker) Invoke() (*model.ResizeDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeDesktopResponse), nil
	}
}

type SendNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendNotificationsInvoker) Invoke() (*model.SendNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendNotificationsResponse), nil
	}
}

type SetMaintenanceModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetMaintenanceModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetMaintenanceModeInvoker) Invoke() (*model.SetMaintenanceModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetMaintenanceModeResponse), nil
	}
}

type ShowDesktopDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopDetailInvoker) Invoke() (*model.ShowDesktopDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopDetailResponse), nil
	}
}

type ShowDesktopMonitorDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopMonitorDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopMonitorDataInvoker) Invoke() (*model.ShowDesktopMonitorDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopMonitorDataResponse), nil
	}
}

type ShowDesktopNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopNetworkInvoker) Invoke() (*model.ShowDesktopNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopNetworkResponse), nil
	}
}

type ShowDesktopNetworksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopNetworksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopNetworksInvoker) Invoke() (*model.ShowDesktopNetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopNetworksResponse), nil
	}
}

type ShowDesktopRemoteAssistanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopRemoteAssistanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopRemoteAssistanceInfoInvoker) Invoke() (*model.ShowDesktopRemoteAssistanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopRemoteAssistanceInfoResponse), nil
	}
}

type ShowRemoteConsoleAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRemoteConsoleAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRemoteConsoleAddressInvoker) Invoke() (*model.ShowRemoteConsoleAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRemoteConsoleAddressResponse), nil
	}
}

type ShowSysprepInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSysprepInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSysprepInfoInvoker) Invoke() (*model.ShowSysprepInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSysprepInfoResponse), nil
	}
}

type UpdateDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesktopInvoker) Invoke() (*model.UpdateDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopResponse), nil
	}
}

type UpdateDesktopSidsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopSidsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesktopSidsInvoker) Invoke() (*model.UpdateDesktopSidsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopSidsResponse), nil
	}
}

type UpdateDesktopUsernameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopUsernameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesktopUsernameInvoker) Invoke() (*model.UpdateDesktopUsernameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopUsernameResponse), nil
	}
}

type BatchDeleteDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopNamePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDesktopNamePolicyInvoker) Invoke() (*model.BatchDeleteDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopNamePolicyResponse), nil
	}
}

type CreateDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopNamePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopNamePolicyInvoker) Invoke() (*model.CreateDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopNamePolicyResponse), nil
	}
}

type ListDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopNamePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopNamePolicyInvoker) Invoke() (*model.ListDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopNamePolicyResponse), nil
	}
}

type UpdateDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopNamePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesktopNamePolicyInvoker) Invoke() (*model.UpdateDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopNamePolicyResponse), nil
	}
}

type AddDesktopPoolVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDesktopPoolVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDesktopPoolVolumesInvoker) Invoke() (*model.AddDesktopPoolVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDesktopPoolVolumesResponse), nil
	}
}

type CreateDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopPoolInvoker) Invoke() (*model.CreateDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopPoolResponse), nil
	}
}

type CreateDesktopPoolAuthorizedObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopPoolAuthorizedObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopPoolAuthorizedObjectsInvoker) Invoke() (*model.CreateDesktopPoolAuthorizedObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopPoolAuthorizedObjectsResponse), nil
	}
}

type DeleteDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesktopPoolInvoker) Invoke() (*model.DeleteDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopPoolResponse), nil
	}
}

type DeleteDesktopPoolVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopPoolVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesktopPoolVolumesInvoker) Invoke() (*model.DeleteDesktopPoolVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopPoolVolumesResponse), nil
	}
}

type ExecuteDesktopPoolActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDesktopPoolActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteDesktopPoolActionInvoker) Invoke() (*model.ExecuteDesktopPoolActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDesktopPoolActionResponse), nil
	}
}

type ExecuteDesktopPoolScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDesktopPoolScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteDesktopPoolScriptInvoker) Invoke() (*model.ExecuteDesktopPoolScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDesktopPoolScriptResponse), nil
	}
}

type ExpandDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandDesktopPoolInvoker) Invoke() (*model.ExpandDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandDesktopPoolResponse), nil
	}
}

type ExpandDesktopPoolVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandDesktopPoolVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandDesktopPoolVolumesInvoker) Invoke() (*model.ExpandDesktopPoolVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandDesktopPoolVolumesResponse), nil
	}
}

type ListDesktopPoolAuthorizedObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopPoolAuthorizedObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopPoolAuthorizedObjectsInvoker) Invoke() (*model.ListDesktopPoolAuthorizedObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopPoolAuthorizedObjectsResponse), nil
	}
}

type ListDesktopPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopPoolsInvoker) Invoke() (*model.ListDesktopPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopPoolsResponse), nil
	}
}

type ListDesktopPoolsByUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopPoolsByUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopPoolsByUsersInvoker) Invoke() (*model.ListDesktopPoolsByUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopPoolsByUsersResponse), nil
	}
}

type ListInconsistentStaticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInconsistentStaticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInconsistentStaticsInvoker) Invoke() (*model.ListInconsistentStaticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInconsistentStaticsResponse), nil
	}
}

type ListPoolDesktopsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolDesktopsDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolDesktopsDetailInvoker) Invoke() (*model.ListPoolDesktopsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolDesktopsDetailResponse), nil
	}
}

type RebuildDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebuildDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebuildDesktopPoolInvoker) Invoke() (*model.RebuildDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebuildDesktopPoolResponse), nil
	}
}

type ResizeDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeDesktopPoolInvoker) Invoke() (*model.ResizeDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeDesktopPoolResponse), nil
	}
}

type SendDesktopPoolNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendDesktopPoolNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendDesktopPoolNotificationsInvoker) Invoke() (*model.SendDesktopPoolNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendDesktopPoolNotificationsResponse), nil
	}
}

type ShowDesktopPoolDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopPoolDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopPoolDetailInvoker) Invoke() (*model.ShowDesktopPoolDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopPoolDetailResponse), nil
	}
}

type ShowDesktopPoolsScriptExecTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopPoolsScriptExecTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDesktopPoolsScriptExecTasksInvoker) Invoke() (*model.ShowDesktopPoolsScriptExecTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopPoolsScriptExecTasksResponse), nil
	}
}

type UpdateDesktopPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesktopPoolInvoker) Invoke() (*model.UpdateDesktopPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopPoolResponse), nil
	}
}

type BatchAddDesktopsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDesktopsTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddDesktopsTagsInvoker) Invoke() (*model.BatchAddDesktopsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDesktopsTagsResponse), nil
	}
}

type BatchChangeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchChangeTagsInvoker) Invoke() (*model.BatchChangeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeTagsResponse), nil
	}
}

type BatchDeleteDesktopsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopsTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDesktopsTagsInvoker) Invoke() (*model.BatchDeleteDesktopsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopsTagsResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListDesktopByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopByTagsInvoker) Invoke() (*model.ListDesktopByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopByTagsResponse), nil
	}
}

type ListDesktopsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsByTagsInvoker) Invoke() (*model.ListDesktopsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsByTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ShowTagByDesktopIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagByDesktopIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagByDesktopIdInvoker) Invoke() (*model.ShowTagByDesktopIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagByDesktopIdResponse), nil
	}
}

type ExportDesktopListNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDesktopListNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDesktopListNewInvoker) Invoke() (*model.ExportDesktopListNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDesktopListNewResponse), nil
	}
}

type ShowHibernateTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHibernateTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHibernateTypeInvoker) Invoke() (*model.ShowHibernateTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHibernateTypeResponse), nil
	}
}

type DeleteExportTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExportTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteExportTasksInvoker) Invoke() (*model.DeleteExportTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExportTasksResponse), nil
	}
}

type DownloadExportFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadExportFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadExportFileInvoker) Invoke() (*model.DownloadExportFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadExportFileResponse), nil
	}
}

type ListExportTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExportTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExportTasksInvoker) Invoke() (*model.ListExportTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExportTasksResponse), nil
	}
}

type BatchDeleteUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteUserGroupsInvoker) Invoke() (*model.BatchDeleteUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteUserGroupsResponse), nil
	}
}

type CreateUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserGroupInvoker) Invoke() (*model.CreateUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserGroupResponse), nil
	}
}

type DeleteUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserGroupInvoker) Invoke() (*model.DeleteUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserGroupResponse), nil
	}
}

type ExportUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserGroupsInvoker) Invoke() (*model.ExportUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserGroupsResponse), nil
	}
}

type ListUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserGroupsInvoker) Invoke() (*model.ListUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserGroupsResponse), nil
	}
}

type ListUsersOfGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersOfGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersOfGroupInvoker) Invoke() (*model.ListUsersOfGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersOfGroupResponse), nil
	}
}

type RunActionsOnGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunActionsOnGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunActionsOnGroupInvoker) Invoke() (*model.RunActionsOnGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunActionsOnGroupResponse), nil
	}
}

type UpdateUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserGroupInvoker) Invoke() (*model.UpdateUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserGroupResponse), nil
	}
}

type ExportUserGroupUsersNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserGroupUsersNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserGroupUsersNewInvoker) Invoke() (*model.ExportUserGroupUsersNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserGroupUsersNewResponse), nil
	}
}

type ListHostsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostsDetailInvoker) Invoke() (*model.ListHostsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsDetailResponse), nil
	}
}

type ListServersByHostIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersByHostIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersByHostIdInvoker) Invoke() (*model.ListServersByHostIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersByHostIdResponse), nil
	}
}

type UpdateHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostsInvoker) Invoke() (*model.UpdateHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostsResponse), nil
	}
}

type ListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImagesInvoker) Invoke() (*model.ListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesResponse), nil
	}
}

type ListMarketImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMarketImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMarketImagesInvoker) Invoke() (*model.ListMarketImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMarketImagesResponse), nil
	}
}

type EstimateAddResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateAddResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateAddResourcesInvoker) Invoke() (*model.EstimateAddResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateAddResourcesResponse), nil
	}
}

type EstimateChangeImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateChangeImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateChangeImagesInvoker) Invoke() (*model.EstimateChangeImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateChangeImagesResponse), nil
	}
}

type EstimateDesktopPoolAddVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateDesktopPoolAddVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateDesktopPoolAddVolumeInvoker) Invoke() (*model.EstimateDesktopPoolAddVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateDesktopPoolAddVolumeResponse), nil
	}
}

type EstimateDesktopPoolChangeImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateDesktopPoolChangeImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateDesktopPoolChangeImageInvoker) Invoke() (*model.EstimateDesktopPoolChangeImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateDesktopPoolChangeImageResponse), nil
	}
}

type EstimateDesktopPoolExtendVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateDesktopPoolExtendVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateDesktopPoolExtendVolumeInvoker) Invoke() (*model.EstimateDesktopPoolExtendVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateDesktopPoolExtendVolumeResponse), nil
	}
}

type EstimateDesktopPoolResizeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateDesktopPoolResizeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EstimateDesktopPoolResizeInvoker) Invoke() (*model.EstimateDesktopPoolResizeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateDesktopPoolResizeResponse), nil
	}
}

type RunActionsOnWorkspaceJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunActionsOnWorkspaceJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunActionsOnWorkspaceJobInvoker) Invoke() (*model.RunActionsOnWorkspaceJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunActionsOnWorkspaceJobResponse), nil
	}
}

type BatchDeleteSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSubJobsInvoker) Invoke() (*model.BatchDeleteSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSubJobsResponse), nil
	}
}

type ListItaSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListItaSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListItaSubJobsInvoker) Invoke() (*model.ListItaSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListItaSubJobsResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ListNatMappingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatMappingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatMappingConfigsInvoker) Invoke() (*model.ListNatMappingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatMappingConfigsResponse), nil
	}
}

type UpdateNatMappingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatMappingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNatMappingConfigsInvoker) Invoke() (*model.UpdateNatMappingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatMappingConfigsResponse), nil
	}
}

type ApplyDesktopsInternetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyDesktopsInternetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyDesktopsInternetInvoker) Invoke() (*model.ApplyDesktopsInternetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyDesktopsInternetResponse), nil
	}
}

type ApplyInternetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyInternetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyInternetInvoker) Invoke() (*model.ApplyInternetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyInternetResponse), nil
	}
}

type ApplySubnetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplySubnetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplySubnetBandwidthInvoker) Invoke() (*model.ApplySubnetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplySubnetBandwidthResponse), nil
	}
}

type AssociateDesktopsEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateDesktopsEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateDesktopsEipInvoker) Invoke() (*model.AssociateDesktopsEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateDesktopsEipResponse), nil
	}
}

type BatchDisassociateDesktopsEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociateDesktopsEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisassociateDesktopsEipInvoker) Invoke() (*model.BatchDisassociateDesktopsEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociateDesktopsEipResponse), nil
	}
}

type CheckCidrInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckCidrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckCidrInvoker) Invoke() (*model.CheckCidrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckCidrResponse), nil
	}
}

type DeleteSubnetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubnetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubnetBandwidthInvoker) Invoke() (*model.DeleteSubnetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubnetBandwidthResponse), nil
	}
}

type ListDesktopsEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsEipsInvoker) Invoke() (*model.ListDesktopsEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsEipsResponse), nil
	}
}

type ListInternetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetInvoker) Invoke() (*model.ListInternetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetResponse), nil
	}
}

type ListNatGatewaysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewaysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewaysInvoker) Invoke() (*model.ListNatGatewaysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewaysResponse), nil
	}
}

type ListPortsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPortsInvoker) Invoke() (*model.ListPortsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortsResponse), nil
	}
}

type ListSecurityGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityGroupsInvoker) Invoke() (*model.ListSecurityGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupsResponse), nil
	}
}

type ListSubnetBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubnetBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubnetBandwidthsInvoker) Invoke() (*model.ListSubnetBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubnetBandwidthsResponse), nil
	}
}

type ListSubnetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubnetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubnetsInvoker) Invoke() (*model.ListSubnetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubnetsResponse), nil
	}
}

type ListVpcInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVpcInfoInvoker) Invoke() (*model.ListVpcInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcInfoResponse), nil
	}
}

type ShowSubnetBandwidthControlListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubnetBandwidthControlListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSubnetBandwidthControlListInvoker) Invoke() (*model.ShowSubnetBandwidthControlListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubnetBandwidthControlListResponse), nil
	}
}

type ShowUsingSubnetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUsingSubnetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUsingSubnetsInvoker) Invoke() (*model.ShowUsingSubnetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUsingSubnetsResponse), nil
	}
}

type UpdateSubnetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubnetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubnetBandwidthInvoker) Invoke() (*model.UpdateSubnetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubnetBandwidthResponse), nil
	}
}

type UpdateSubnetBandwidthControlListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubnetBandwidthControlListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubnetBandwidthControlListInvoker) Invoke() (*model.UpdateSubnetBandwidthControlListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubnetBandwidthControlListResponse), nil
	}
}

type CreateChangeOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateChangeOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateChangeOrderInvoker) Invoke() (*model.CreateChangeOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateChangeOrderResponse), nil
	}
}

type CreateDesktopBatchOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopBatchOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopBatchOrderInvoker) Invoke() (*model.CreateDesktopBatchOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopBatchOrderResponse), nil
	}
}

type CreateDesktopOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopOrderInvoker) Invoke() (*model.CreateDesktopOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopOrderResponse), nil
	}
}

type CreateDesktopPoolChangeOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopPoolChangeOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopPoolChangeOrderInvoker) Invoke() (*model.CreateDesktopPoolChangeOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopPoolChangeOrderResponse), nil
	}
}

type CreateOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrderInvoker) Invoke() (*model.CreateOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrderResponse), nil
	}
}

type CreateResourcePackagesOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourcePackagesOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourcePackagesOrderInvoker) Invoke() (*model.CreateResourcePackagesOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourcePackagesOrderResponse), nil
	}
}

type CreateSubnetBandwidthChangeOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubnetBandwidthChangeOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubnetBandwidthChangeOrderInvoker) Invoke() (*model.CreateSubnetBandwidthChangeOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubnetBandwidthChangeOrderResponse), nil
	}
}

type AddOuInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddOuInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddOuInvoker) Invoke() (*model.AddOuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddOuResponse), nil
	}
}

type DeleteOuInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOuInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOuInvoker) Invoke() (*model.DeleteOuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOuResponse), nil
	}
}

type ListAdOuUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAdOuUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAdOuUsersInvoker) Invoke() (*model.ListAdOuUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAdOuUsersResponse), nil
	}
}

type ListAdOusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAdOusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAdOusInvoker) Invoke() (*model.ListAdOusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAdOusResponse), nil
	}
}

type ListOuDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOuDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOuDetailsInvoker) Invoke() (*model.ListOuDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOuDetailsResponse), nil
	}
}

type UpdateOuInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOuInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOuInfoInvoker) Invoke() (*model.UpdateOuInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOuInfoResponse), nil
	}
}

type BatchUpdateTargetOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateTargetOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateTargetOfPolicyGroupInvoker) Invoke() (*model.BatchUpdateTargetOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateTargetOfPolicyGroupResponse), nil
	}
}

type CreatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyGroupInvoker) Invoke() (*model.CreatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyGroupResponse), nil
	}
}

type CreatePolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyTemplateInvoker) Invoke() (*model.CreatePolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyTemplateResponse), nil
	}
}

type DeletePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) Invoke() (*model.DeletePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyGroupResponse), nil
	}
}

type ExportPolicyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportPolicyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportPolicyGroupsInvoker) Invoke() (*model.ExportPolicyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportPolicyGroupsResponse), nil
	}
}

type ImportPolicyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportPolicyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportPolicyGroupsInvoker) Invoke() (*model.ImportPolicyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportPolicyGroupsResponse), nil
	}
}

type ListOriginalPolicyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOriginalPolicyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOriginalPolicyInfoInvoker) Invoke() (*model.ListOriginalPolicyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOriginalPolicyInfoResponse), nil
	}
}

type ListPoliciesOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesOfPolicyGroupInvoker) Invoke() (*model.ListPoliciesOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesOfPolicyGroupResponse), nil
	}
}

type ListPolicyDetailInfoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyDetailInfoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyDetailInfoByIdInvoker) Invoke() (*model.ListPolicyDetailInfoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyDetailInfoByIdResponse), nil
	}
}

type ListPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyGroupInvoker) Invoke() (*model.ListPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupResponse), nil
	}
}

type ListPolicyGroupInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyGroupInfoInvoker) Invoke() (*model.ListPolicyGroupInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupInfoResponse), nil
	}
}

type ListTargetOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTargetOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTargetOfPolicyGroupInvoker) Invoke() (*model.ListTargetOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTargetOfPolicyGroupResponse), nil
	}
}

type UpdatePoliciesOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePoliciesOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePoliciesOfPolicyGroupInvoker) Invoke() (*model.UpdatePoliciesOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePoliciesOfPolicyGroupResponse), nil
	}
}

type UpdatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyGroupInvoker) Invoke() (*model.UpdatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyGroupResponse), nil
	}
}

type ListHourPackagesTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHourPackagesTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHourPackagesTypeInvoker) Invoke() (*model.ListHourPackagesTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHourPackagesTypeResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ListResourcePackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcePackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcePackagesInvoker) Invoke() (*model.ListResourcePackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcePackagesResponse), nil
	}
}

type ListSharerProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharerProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharerProductsInvoker) Invoke() (*model.ListSharerProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharerProductsResponse), nil
	}
}

type ListTenantProfilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantProfilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantProfilesInvoker) Invoke() (*model.ListTenantProfilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantProfilesResponse), nil
	}
}

type UpdateTenantProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTenantProfileInvoker) Invoke() (*model.UpdateTenantProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantProfileResponse), nil
	}
}

type ShowQuotaDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaDetailsInvoker) Invoke() (*model.ShowQuotaDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaDetailsResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type BatchDeleteScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteScheduledTasksInvoker) Invoke() (*model.BatchDeleteScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScheduledTasksResponse), nil
	}
}

type CreateScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScheduledTasksInvoker) Invoke() (*model.CreateScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduledTasksResponse), nil
	}
}

type DeleteScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduledTasksInvoker) Invoke() (*model.DeleteScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduledTasksResponse), nil
	}
}

type ExportScheduledTasksRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportScheduledTasksRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportScheduledTasksRecordsInvoker) Invoke() (*model.ExportScheduledTasksRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportScheduledTasksRecordsResponse), nil
	}
}

type ListFutureExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFutureExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFutureExecutionsInvoker) Invoke() (*model.ListFutureExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFutureExecutionsResponse), nil
	}
}

type ListScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTasksInvoker) Invoke() (*model.ListScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksResponse), nil
	}
}

type ListScheduledTasksRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTasksRecordsInvoker) Invoke() (*model.ListScheduledTasksRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksRecordsResponse), nil
	}
}

type ListScheduledTasksRecordsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksRecordsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTasksRecordsDetailsInvoker) Invoke() (*model.ListScheduledTasksRecordsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksRecordsDetailsResponse), nil
	}
}

type ListTimeZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTimeZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTimeZonesInvoker) Invoke() (*model.ListTimeZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTimeZonesResponse), nil
	}
}

type ShowScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScheduledTasksInvoker) Invoke() (*model.ShowScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduledTasksResponse), nil
	}
}

type UpdateScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScheduledTasksInvoker) Invoke() (*model.UpdateScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduledTasksResponse), nil
	}
}

type BatchDeleteScreenRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScreenRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteScreenRecordsInvoker) Invoke() (*model.BatchDeleteScreenRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScreenRecordsResponse), nil
	}
}

type ListDesktopOperationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopOperationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopOperationsInvoker) Invoke() (*model.ListDesktopOperationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopOperationsResponse), nil
	}
}

type ListDownloadAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDownloadAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDownloadAddressInvoker) Invoke() (*model.ListDownloadAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDownloadAddressResponse), nil
	}
}

type ListScreenRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScreenRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScreenRecordsInvoker) Invoke() (*model.ListScreenRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScreenRecordsResponse), nil
	}
}

type ShowScreenRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScreenRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScreenRecordInvoker) Invoke() (*model.ShowScreenRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScreenRecordResponse), nil
	}
}

type CreateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScriptInvoker) Invoke() (*model.CreateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScriptResponse), nil
	}
}

type DeleteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScriptInvoker) Invoke() (*model.DeleteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScriptResponse), nil
	}
}

type ExecuteScriptByDesktopTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteScriptByDesktopTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteScriptByDesktopTagInvoker) Invoke() (*model.ExecuteScriptByDesktopTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteScriptByDesktopTagResponse), nil
	}
}

type ExecuteScriptOrCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteScriptOrCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteScriptOrCommandInvoker) Invoke() (*model.ExecuteScriptOrCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteScriptOrCommandResponse), nil
	}
}

type ListScriptRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScriptRecordsInvoker) Invoke() (*model.ListScriptRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptRecordsResponse), nil
	}
}

type ListScriptTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScriptTasksInvoker) Invoke() (*model.ListScriptTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptTasksResponse), nil
	}
}

type ListScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScriptsInvoker) Invoke() (*model.ListScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptsResponse), nil
	}
}

type RetryScriptExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryScriptExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryScriptExecutionInvoker) Invoke() (*model.RetryScriptExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryScriptExecutionResponse), nil
	}
}

type ShowScriptDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScriptDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScriptDetailInvoker) Invoke() (*model.ShowScriptDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScriptDetailResponse), nil
	}
}

type ShowScriptRecordDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScriptRecordDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScriptRecordDetailInvoker) Invoke() (*model.ShowScriptRecordDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScriptRecordDetailResponse), nil
	}
}

type StopScriptExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopScriptExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopScriptExecutionInvoker) Invoke() (*model.StopScriptExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopScriptExecutionResponse), nil
	}
}

type UpdateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScriptInvoker) Invoke() (*model.UpdateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScriptResponse), nil
	}
}

type AddDesktopSubResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDesktopSubResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDesktopSubResourcesInvoker) Invoke() (*model.AddDesktopSubResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDesktopSubResourcesResponse), nil
	}
}

type DeleteDesktopSubResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopSubResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesktopSubResourcesInvoker) Invoke() (*model.DeleteDesktopSubResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopSubResourcesResponse), nil
	}
}

type ShowShareSpaceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShareSpaceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShareSpaceConfigInvoker) Invoke() (*model.ShowShareSpaceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShareSpaceConfigResponse), nil
	}
}

type UpdateShareSpaceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateShareSpaceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateShareSpaceConfigInvoker) Invoke() (*model.UpdateShareSpaceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateShareSpaceConfigResponse), nil
	}
}

type AddSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSiteInvoker) Invoke() (*model.AddSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSiteResponse), nil
	}
}

type CheckEdgeSiteResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckEdgeSiteResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckEdgeSiteResourcesInvoker) Invoke() (*model.CheckEdgeSiteResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckEdgeSiteResourcesResponse), nil
	}
}

type DeleteSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSiteInvoker) Invoke() (*model.DeleteSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSiteResponse), nil
	}
}

type ListSiteConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSiteConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSiteConfigsInvoker) Invoke() (*model.ListSiteConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSiteConfigsResponse), nil
	}
}

type ListWksEdgeSitesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWksEdgeSitesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWksEdgeSitesInvoker) Invoke() (*model.ListWksEdgeSitesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWksEdgeSitesResponse), nil
	}
}

type UpdateAccessModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessModeInvoker) Invoke() (*model.UpdateAccessModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessModeResponse), nil
	}
}

type UpdateSubnetIdsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubnetIdsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubnetIdsInvoker) Invoke() (*model.UpdateSubnetIdsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubnetIdsResponse), nil
	}
}

type BatchCreateDesktopSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDesktopSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDesktopSnapshotInvoker) Invoke() (*model.BatchCreateDesktopSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDesktopSnapshotResponse), nil
	}
}

type BatchDeleteDesktopSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDesktopSnapshotInvoker) Invoke() (*model.BatchDeleteDesktopSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopSnapshotResponse), nil
	}
}

type BatchRestoreDesktopSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRestoreDesktopSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRestoreDesktopSnapshotInvoker) Invoke() (*model.BatchRestoreDesktopSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRestoreDesktopSnapshotResponse), nil
	}
}

type ListDesktopSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopSnapshotInvoker) Invoke() (*model.ListDesktopSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopSnapshotResponse), nil
	}
}

type AddMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddMetricNotifyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddMetricNotifyRuleInvoker) Invoke() (*model.AddMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddMetricNotifyRuleResponse), nil
	}
}

type DeleteMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMetricNotifyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMetricNotifyRuleInvoker) Invoke() (*model.DeleteMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMetricNotifyRuleResponse), nil
	}
}

type ExportAppUserAccessDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportAppUserAccessDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportAppUserAccessDataInvoker) Invoke() (*model.ExportAppUserAccessDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportAppUserAccessDataResponse), nil
	}
}

type ListAppUserAccessDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppUserAccessDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppUserAccessDataInvoker) Invoke() (*model.ListAppUserAccessDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppUserAccessDataResponse), nil
	}
}

type ListDesktopUsageMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopUsageMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopUsageMetricInvoker) Invoke() (*model.ListDesktopUsageMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopUsageMetricResponse), nil
	}
}

type ListDesktopsStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDesktopsStatisticsInvoker) Invoke() (*model.ListDesktopsStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsStatisticsResponse), nil
	}
}

type ListLoginStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginStateInvoker) Invoke() (*model.ListLoginStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginStateResponse), nil
	}
}

type ListMetricNotifyRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricNotifyRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricNotifyRecordInvoker) Invoke() (*model.ListMetricNotifyRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricNotifyRecordResponse), nil
	}
}

type ListMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricNotifyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricNotifyRuleInvoker) Invoke() (*model.ListMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricNotifyRuleResponse), nil
	}
}

type ListMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricsInvoker) Invoke() (*model.ListMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsResponse), nil
	}
}

type ListMetricsTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricsTrendInvoker) Invoke() (*model.ListMetricsTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsTrendResponse), nil
	}
}

type ListRunStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRunStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRunStateInvoker) Invoke() (*model.ListRunStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRunStateResponse), nil
	}
}

type ListUnusedDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnusedDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUnusedDesktopsInvoker) Invoke() (*model.ListUnusedDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnusedDesktopsResponse), nil
	}
}

type ListUsedDesktopInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsedDesktopInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsedDesktopInfoInvoker) Invoke() (*model.ListUsedDesktopInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsedDesktopInfoResponse), nil
	}
}

type ListUserUsageMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserUsageMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserUsageMetricInvoker) Invoke() (*model.ListUserUsageMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserUsageMetricResponse), nil
	}
}

type ShowGrowthRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGrowthRateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGrowthRateInvoker) Invoke() (*model.ShowGrowthRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGrowthRateResponse), nil
	}
}

type ShowUserAccessStagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserAccessStagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserAccessStagesInvoker) Invoke() (*model.ShowUserAccessStagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserAccessStagesResponse), nil
	}
}

type UpdateMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMetricNotifyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMetricNotifyRuleInvoker) Invoke() (*model.UpdateMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMetricNotifyRuleResponse), nil
	}
}

type ExportDesktopUsageMetricNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDesktopUsageMetricNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDesktopUsageMetricNewInvoker) Invoke() (*model.ExportDesktopUsageMetricNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDesktopUsageMetricNewResponse), nil
	}
}

type ExportUserUsageMetricNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserUsageMetricNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserUsageMetricNewInvoker) Invoke() (*model.ExportUserUsageMetricNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserUsageMetricNewResponse), nil
	}
}

type ShowAvailableIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAvailableIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAvailableIpInvoker) Invoke() (*model.ShowAvailableIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAvailableIpResponse), nil
	}
}

type ListTenantConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantConfigsInvoker) Invoke() (*model.ListTenantConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantConfigsResponse), nil
	}
}

type UpdateTenantConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTenantConfigInvoker) Invoke() (*model.UpdateTenantConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantConfigResponse), nil
	}
}

type CreateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTerminalsBindingDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTerminalsBindingDesktopsInvoker) Invoke() (*model.CreateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTerminalsBindingDesktopsResponse), nil
	}
}

type DeleteTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTerminalsBindingDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTerminalsBindingDesktopsInvoker) Invoke() (*model.DeleteTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTerminalsBindingDesktopsResponse), nil
	}
}

type ExportTerminalsBindingDesktopsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTerminalsBindingDesktopsTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTerminalsBindingDesktopsTemplateInvoker) Invoke() (*model.ExportTerminalsBindingDesktopsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTerminalsBindingDesktopsTemplateResponse), nil
	}
}

type ListTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsInvoker) Invoke() (*model.ListTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsResponse), nil
	}
}

type ListTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.ListTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsConfigResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsConfigResponse), nil
	}
}

type ExportTerminalsBindingDesktopsInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTerminalsBindingDesktopsInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTerminalsBindingDesktopsInfoNewInvoker) Invoke() (*model.ExportTerminalsBindingDesktopsInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTerminalsBindingDesktopsInfoNewResponse), nil
	}
}

type BatchCreateUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateUsersInvoker) Invoke() (*model.BatchCreateUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateUsersResponse), nil
	}
}

type BatchDeleteOtpDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteOtpDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteOtpDevicesInvoker) Invoke() (*model.BatchDeleteOtpDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteOtpDevicesResponse), nil
	}
}

type BatchDeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteUserInvoker) Invoke() (*model.BatchDeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteUserResponse), nil
	}
}

type ChangeUserStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeUserStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeUserStatusInvoker) Invoke() (*model.ChangeUserStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeUserStatusResponse), nil
	}
}

type CreateDesktopUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesktopUserInvoker) Invoke() (*model.CreateDesktopUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ExportUserListTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserListTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUserListTemplateInvoker) Invoke() (*model.ExportUserListTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserListTemplateResponse), nil
	}
}

type ExportUsersTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUsersTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUsersTemplateInvoker) Invoke() (*model.ExportUsersTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUsersTemplateResponse), nil
	}
}

type ImportUserListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportUserListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportUserListInvoker) Invoke() (*model.ImportUserListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportUserListResponse), nil
	}
}

type ListNotificationRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationRecordsInvoker) Invoke() (*model.ListNotificationRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationRecordsResponse), nil
	}
}

type ListOtpDevicesByUserIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOtpDevicesByUserIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOtpDevicesByUserIdInvoker) Invoke() (*model.ListOtpDevicesByUserIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOtpDevicesByUserIdResponse), nil
	}
}

type ListUserDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserDetailInvoker) Invoke() (*model.ListUserDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserDetailResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ResetRandomPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetRandomPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetRandomPasswordInvoker) Invoke() (*model.ResetRandomPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetRandomPasswordResponse), nil
	}
}

type SendEmailInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendEmailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendEmailInvoker) Invoke() (*model.SendEmailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendEmailResponse), nil
	}
}

type UpdateUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserInfoInvoker) Invoke() (*model.UpdateUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserInfoResponse), nil
	}
}

type ListUserEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserEventsInvoker) Invoke() (*model.ListUserEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserEventsResponse), nil
	}
}

type ListUserEventsLtsConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserEventsLtsConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserEventsLtsConfigurationsInvoker) Invoke() (*model.ListUserEventsLtsConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserEventsLtsConfigurationsResponse), nil
	}
}

type SetUserEventsLtsConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetUserEventsLtsConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetUserEventsLtsConfigurationsInvoker) Invoke() (*model.SetUserEventsLtsConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetUserEventsLtsConfigurationsResponse), nil
	}
}

type ExportUsersNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUsersNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportUsersNewInvoker) Invoke() (*model.ExportUsersNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUsersNewResponse), nil
	}
}

type AddDesktopVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDesktopVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDesktopVolumesInvoker) Invoke() (*model.AddDesktopVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDesktopVolumesResponse), nil
	}
}

type AddVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddVolumesInvoker) Invoke() (*model.AddVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVolumesResponse), nil
	}
}

type BatchModifyQosVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchModifyQosVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchModifyQosVolumesInvoker) Invoke() (*model.BatchModifyQosVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchModifyQosVolumesResponse), nil
	}
}

type DeleteDesktopVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesktopVolumesInvoker) Invoke() (*model.DeleteDesktopVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopVolumesResponse), nil
	}
}

type ExpandDesktopVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandDesktopVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandDesktopVolumeInvoker) Invoke() (*model.ExpandDesktopVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandDesktopVolumeResponse), nil
	}
}

type ExpandVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandVolumesInvoker) Invoke() (*model.ExpandVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandVolumesResponse), nil
	}
}

type ListVolumeProductInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumeProductInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumeProductInfoInvoker) Invoke() (*model.ListVolumeProductInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumeProductInfoResponse), nil
	}
}

type ApplyWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyWorkspaceInvoker) Invoke() (*model.ApplyWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyWorkspaceResponse), nil
	}
}

type CancelWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelWorkspaceInvoker) Invoke() (*model.CancelWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelWorkspaceResponse), nil
	}
}

type CheckEnterpriseIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckEnterpriseIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckEnterpriseIdInvoker) Invoke() (*model.CheckEnterpriseIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckEnterpriseIdResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type ShowWorkspaceLockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceLockInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkspaceLockInvoker) Invoke() (*model.ShowWorkspaceLockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceLockResponse), nil
	}
}

type UnlockWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnlockWorkspaceInvoker) Invoke() (*model.UnlockWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockWorkspaceResponse), nil
	}
}

type UpdateEnterpriseIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnterpriseIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnterpriseIdInvoker) Invoke() (*model.UpdateEnterpriseIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnterpriseIdResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}

type ValidateDomainControllerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateDomainControllerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateDomainControllerInvoker) Invoke() (*model.ValidateDomainControllerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateDomainControllerResponse), nil
	}
}
