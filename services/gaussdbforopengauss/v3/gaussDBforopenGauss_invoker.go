package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbforopengauss/v3/model"
)

type AddInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddInstanceTagsInvoker) Invoke() (*model.AddInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddInstanceTagsResponse), nil
	}
}

type AllowDbPrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AllowDbPrivilegesInvoker) Invoke() (*model.AllowDbPrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowDbPrivilegesResponse), nil
	}
}

type AttachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEipInvoker) Invoke() (*model.AttachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEipResponse), nil
	}
}

type ConfirmRestoredDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmRestoredDataInvoker) Invoke() (*model.ConfirmRestoredDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmRestoredDataResponse), nil
	}
}

type CopyConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyConfigurationInvoker) Invoke() (*model.CopyConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyConfigurationResponse), nil
	}
}

type CreateConfigurationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationTemplateInvoker) Invoke() (*model.CreateConfigurationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationTemplateResponse), nil
	}
}

type CreateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseInvoker) Invoke() (*model.CreateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResponse), nil
	}
}

type CreateDatabaseSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseSchemasInvoker) Invoke() (*model.CreateDatabaseSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseSchemasResponse), nil
	}
}

type CreateDbInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbInstanceInvoker) Invoke() (*model.CreateDbInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbInstanceResponse), nil
	}
}

type CreateDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbUserInvoker) Invoke() (*model.CreateDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbUserResponse), nil
	}
}

type CreateGaussDbInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussDbInstanceInvoker) Invoke() (*model.CreateGaussDbInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussDbInstanceResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManualBackupInvoker) Invoke() (*model.CreateManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManualBackupResponse), nil
	}
}

type CreateRestoreInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRestoreInstanceInvoker) Invoke() (*model.CreateRestoreInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRestoreInstanceResponse), nil
	}
}

type CreateSlowLogDownloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSlowLogDownloadInvoker) Invoke() (*model.CreateSlowLogDownloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSlowLogDownloadResponse), nil
	}
}

type DeleteConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigurationInvoker) Invoke() (*model.DeleteConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigurationResponse), nil
	}
}

type DeleteDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseInvoker) Invoke() (*model.DeleteDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceTagInvoker) Invoke() (*model.DeleteInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceTagResponse), nil
	}
}

type DeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobInvoker) Invoke() (*model.DeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobResponse), nil
	}
}

type DeleteManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteManualBackupInvoker) Invoke() (*model.DeleteManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteManualBackupResponse), nil
	}
}

type DownloadBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBackupInvoker) Invoke() (*model.DownloadBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBackupResponse), nil
	}
}

type InstallKernelPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallKernelPluginInvoker) Invoke() (*model.InstallKernelPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallKernelPluginResponse), nil
	}
}

type ListApplicableInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicableInstancesInvoker) Invoke() (*model.ListApplicableInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicableInstancesResponse), nil
	}
}

type ListAppliedHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppliedHistoriesInvoker) Invoke() (*model.ListAppliedHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppliedHistoriesResponse), nil
	}
}

type ListAvailableFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableFlavorsInvoker) Invoke() (*model.ListAvailableFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableFlavorsResponse), nil
	}
}

type ListBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListBindedEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBindedEipsInvoker) Invoke() (*model.ListBindedEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBindedEipsResponse), nil
	}
}

type ListCnInfosBeforeReduceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCnInfosBeforeReduceInvoker) Invoke() (*model.ListCnInfosBeforeReduceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCnInfosBeforeReduceResponse), nil
	}
}

type ListComponentInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentInfosInvoker) Invoke() (*model.ListComponentInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentInfosResponse), nil
	}
}

type ListConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsInvoker) Invoke() (*model.ListConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsResponse), nil
	}
}

type ListConfigurationsDiffInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsDiffInvoker) Invoke() (*model.ListConfigurationsDiffResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsDiffResponse), nil
	}
}

type ListDatabaseSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseSchemasInvoker) Invoke() (*model.ListDatabaseSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseSchemasResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ListDatastoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatastoresInvoker) Invoke() (*model.ListDatastoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatastoresResponse), nil
	}
}

type ListDbBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbBackupsInvoker) Invoke() (*model.ListDbBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbBackupsResponse), nil
	}
}

type ListDbFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbFlavorsInvoker) Invoke() (*model.ListDbFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbFlavorsResponse), nil
	}
}

type ListDbUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbUsersInvoker) Invoke() (*model.ListDbUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbUsersResponse), nil
	}
}

type ListEpsQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEpsQuotasInvoker) Invoke() (*model.ListEpsQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEpsQuotasResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListGaussDbDatastoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussDbDatastoresInvoker) Invoke() (*model.ListGaussDbDatastoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussDbDatastoresResponse), nil
	}
}

type ListHistoryOperationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryOperationsInvoker) Invoke() (*model.ListHistoryOperationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryOperationsResponse), nil
	}
}

type ListInstanceDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceDetailsInvoker) Invoke() (*model.ListInstanceDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceDetailsResponse), nil
	}
}

type ListInstanceErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceErrorLogsInvoker) Invoke() (*model.ListInstanceErrorLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceErrorLogsResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListInstancesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesDetailsInvoker) Invoke() (*model.ListInstancesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesDetailsResponse), nil
	}
}

type ListKernelPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKernelPluginsInvoker) Invoke() (*model.ListKernelPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKernelPluginsResponse), nil
	}
}

type ListParamGroupTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListParamGroupTemplatesInvoker) Invoke() (*model.ListParamGroupTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListParamGroupTemplatesResponse), nil
	}
}

type ListPluginExtensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginExtensionsInvoker) Invoke() (*model.ListPluginExtensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginExtensionsResponse), nil
	}
}

type ListPredefinedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPredefinedTagsInvoker) Invoke() (*model.ListPredefinedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPredefinedTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListRecycleInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecycleInstancesInvoker) Invoke() (*model.ListRecycleInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleInstancesResponse), nil
	}
}

type ListRestorableInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestorableInstancesInvoker) Invoke() (*model.ListRestorableInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestorableInstancesResponse), nil
	}
}

type ListRestoreTimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTimesInvoker) Invoke() (*model.ListRestoreTimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTimesResponse), nil
	}
}

type ListStorageTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageTypesInvoker) Invoke() (*model.ListStorageTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageTypesResponse), nil
	}
}

type ListSupportKernelPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportKernelPluginsInvoker) Invoke() (*model.ListSupportKernelPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportKernelPluginsResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListTopIoTrafficsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopIoTrafficsInvoker) Invoke() (*model.ListTopIoTrafficsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopIoTrafficsResponse), nil
	}
}

type ModifyEpsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyEpsQuotaInvoker) Invoke() (*model.ModifyEpsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyEpsQuotaResponse), nil
	}
}

type ResetConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetConfigurationInvoker) Invoke() (*model.ResetConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConfigurationResponse), nil
	}
}

type ResetPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdInvoker) Invoke() (*model.ResetPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPwdResponse), nil
	}
}

type ResizeInstanceFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceFlavorInvoker) Invoke() (*model.ResizeInstanceFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceFlavorResponse), nil
	}
}

type RestartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartInstanceInvoker) Invoke() (*model.RestartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartInstanceResponse), nil
	}
}

type RestoreInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreInstanceInvoker) Invoke() (*model.RestoreInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreInstanceResponse), nil
	}
}

type ResumePluginExtensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumePluginExtensionsInvoker) Invoke() (*model.ResumePluginExtensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumePluginExtensionsResponse), nil
	}
}

type RunInstanceActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunInstanceActionInvoker) Invoke() (*model.RunInstanceActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunInstanceActionResponse), nil
	}
}

type SearchAutoEnlargePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchAutoEnlargePolicyInvoker) Invoke() (*model.SearchAutoEnlargePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchAutoEnlargePolicyResponse), nil
	}
}

type SetBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBackupPolicyInvoker) Invoke() (*model.SetBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBackupPolicyResponse), nil
	}
}

type SetDbUserPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetDbUserPwdInvoker) Invoke() (*model.SetDbUserPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetDbUserPwdResponse), nil
	}
}

type SetKernelPluginLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetKernelPluginLicenseInvoker) Invoke() (*model.SetKernelPluginLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetKernelPluginLicenseResponse), nil
	}
}

type SetNewBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetNewBackupPolicyInvoker) Invoke() (*model.SetNewBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetNewBackupPolicyResponse), nil
	}
}

type SetRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRecyclePolicyInvoker) Invoke() (*model.SetRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRecyclePolicyResponse), nil
	}
}

type ShowBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPolicyInvoker) Invoke() (*model.ShowBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPolicyResponse), nil
	}
}

type ShowBalanceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBalanceStatusInvoker) Invoke() (*model.ShowBalanceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBalanceStatusResponse), nil
	}
}

type ShowBatchUpgradeCandidateVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchUpgradeCandidateVersionsInvoker) Invoke() (*model.ShowBatchUpgradeCandidateVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchUpgradeCandidateVersionsResponse), nil
	}
}

type ShowConfigurationDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationDetailInvoker) Invoke() (*model.ShowConfigurationDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationDetailResponse), nil
	}
}

type ShowDeploymentFormInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentFormInvoker) Invoke() (*model.ShowDeploymentFormResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentFormResponse), nil
	}
}

type ShowErrorLogSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowErrorLogSwitchStatusInvoker) Invoke() (*model.ShowErrorLogSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowErrorLogSwitchStatusResponse), nil
	}
}

type ShowInstanceConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceConfigurationInvoker) Invoke() (*model.ShowInstanceConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceConfigurationResponse), nil
	}
}

type ShowInstanceDiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceDiskInvoker) Invoke() (*model.ShowInstanceDiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceDiskResponse), nil
	}
}

type ShowInstanceParamGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceParamGroupInvoker) Invoke() (*model.ShowInstanceParamGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceParamGroupResponse), nil
	}
}

type ShowInstanceSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceSnapshotInvoker) Invoke() (*model.ShowInstanceSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceSnapshotResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowProjectQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectQuotasInvoker) Invoke() (*model.ShowProjectQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectQuotasResponse), nil
	}
}

type ShowRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) Invoke() (*model.ShowRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecyclePolicyResponse), nil
	}
}

type ShowSlowLogDownloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowLogDownloadInvoker) Invoke() (*model.ShowSlowLogDownloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowLogDownloadResponse), nil
	}
}

type ShowSslCertDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSslCertDownloadLinkInvoker) Invoke() (*model.ShowSslCertDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSslCertDownloadLinkResponse), nil
	}
}

type ShowUpgradeCandidateVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpgradeCandidateVersionsInvoker) Invoke() (*model.ShowUpgradeCandidateVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpgradeCandidateVersionsResponse), nil
	}
}

type StartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceInvoker) Invoke() (*model.StartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResponse), nil
	}
}

type StopBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBackupInvoker) Invoke() (*model.StopBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBackupResponse), nil
	}
}

type SwitchConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchConfigurationInvoker) Invoke() (*model.SwitchConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchConfigurationResponse), nil
	}
}

type SwitchShardInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchShardInvoker) Invoke() (*model.SwitchShardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchShardResponse), nil
	}
}

type UpdateInstanceConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigurationInvoker) Invoke() (*model.UpdateInstanceConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

type UpdateInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceNameInvoker) Invoke() (*model.UpdateInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNameResponse), nil
	}
}

type UpgradeInstanceVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeInstanceVersionInvoker) Invoke() (*model.UpgradeInstanceVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeInstanceVersionResponse), nil
	}
}

type UpgradeInstancesVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeInstancesVersionInvoker) Invoke() (*model.UpgradeInstancesVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeInstancesVersionResponse), nil
	}
}

type ValidateParaGroupNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateParaGroupNameInvoker) Invoke() (*model.ValidateParaGroupNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateParaGroupNameResponse), nil
	}
}

type ValidateWeakPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateWeakPasswordInvoker) Invoke() (*model.ValidateWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateWeakPasswordResponse), nil
	}
}
