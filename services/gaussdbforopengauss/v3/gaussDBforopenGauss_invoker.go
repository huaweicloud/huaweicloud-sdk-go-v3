package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbforopengauss/v3/model"
)

type AddInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AllowDbPrivilegesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AllowDbPrivilegesInvoker) Invoke() (*model.AllowDbPrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowDbPrivilegesResponse), nil
	}
}

type AllowDbRolePrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AllowDbRolePrivilegesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AllowDbRolePrivilegesInvoker) Invoke() (*model.AllowDbRolePrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowDbRolePrivilegesResponse), nil
	}
}

type AttachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachEipInvoker) Invoke() (*model.AttachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEipResponse), nil
	}
}

type BatchShowUpgradeCandidateVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowUpgradeCandidateVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowUpgradeCandidateVersionsInvoker) Invoke() (*model.BatchShowUpgradeCandidateVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowUpgradeCandidateVersionsResponse), nil
	}
}

type CancelScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelScheduleTaskInvoker) Invoke() (*model.CancelScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelScheduleTaskResponse), nil
	}
}

type ConfirmRestoredDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmRestoredDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CopyConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateConfigurationTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigurationTemplateInvoker) Invoke() (*model.CreateConfigurationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationTemplateResponse), nil
	}
}

type CreateCrossCloudConstructDisasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCrossCloudConstructDisasterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCrossCloudConstructDisasterInvoker) Invoke() (*model.CreateCrossCloudConstructDisasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCrossCloudConstructDisasterResponse), nil
	}
}

type CreateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatabaseInvoker) Invoke() (*model.CreateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResponse), nil
	}
}

type CreateDatabaseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatabaseInstanceInvoker) Invoke() (*model.CreateDatabaseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseInstanceResponse), nil
	}
}

type CreateDatabaseSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateDbInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDbInstanceInvoker) Invoke() (*model.CreateDbInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbInstanceResponse), nil
	}
}

type CreateDbRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDbRoleInvoker) Invoke() (*model.CreateDbRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbRoleResponse), nil
	}
}

type CreateDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussDbInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateManualBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateRestoreInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRestoreInstanceInvoker) Invoke() (*model.CreateRestoreInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRestoreInstanceResponse), nil
	}
}

type CreateScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScheduleTaskInvoker) Invoke() (*model.CreateScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduleTaskResponse), nil
	}
}

type CreateSlowLogDownloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSlowLogDownloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatabaseInvoker) Invoke() (*model.DeleteDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResponse), nil
	}
}

type DeleteDatabaseSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseSchemaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatabaseSchemaInvoker) Invoke() (*model.DeleteDatabaseSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseSchemaResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteManualBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteManualBackupInvoker) Invoke() (*model.DeleteManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteManualBackupResponse), nil
	}
}

type DeleteScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduleTaskInvoker) Invoke() (*model.DeleteScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleTaskResponse), nil
	}
}

type DownloadBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadBackupInvoker) Invoke() (*model.DownloadBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBackupResponse), nil
	}
}

type ExecuteCrossCloudDisasterDataCacheEndInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterDataCacheEndInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterDataCacheEndInvoker) Invoke() (*model.ExecuteCrossCloudDisasterDataCacheEndResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterDataCacheEndResponse), nil
	}
}

type ExecuteCrossCloudDisasterDataCacheStartInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterDataCacheStartInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterDataCacheStartInvoker) Invoke() (*model.ExecuteCrossCloudDisasterDataCacheStartResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterDataCacheStartResponse), nil
	}
}

type ExecuteCrossCloudDisasterEndSimulationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterEndSimulationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterEndSimulationInvoker) Invoke() (*model.ExecuteCrossCloudDisasterEndSimulationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterEndSimulationResponse), nil
	}
}

type ExecuteCrossCloudDisasterRecoveryFailoverInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterRecoveryFailoverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterRecoveryFailoverInvoker) Invoke() (*model.ExecuteCrossCloudDisasterRecoveryFailoverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterRecoveryFailoverResponse), nil
	}
}

type ExecuteCrossCloudDisasterRestoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterRestoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterRestoreInvoker) Invoke() (*model.ExecuteCrossCloudDisasterRestoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterRestoreResponse), nil
	}
}

type ExecuteCrossCloudDisasterStartSimulationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterStartSimulationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterStartSimulationInvoker) Invoke() (*model.ExecuteCrossCloudDisasterStartSimulationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterStartSimulationResponse), nil
	}
}

type ExecuteCrossCloudDisasterSwitchoverInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterSwitchoverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudDisasterSwitchoverInvoker) Invoke() (*model.ExecuteCrossCloudDisasterSwitchoverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudDisasterSwitchoverResponse), nil
	}
}

type ExecuteCrossCloudReleaseDisasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCrossCloudReleaseDisasterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCrossCloudReleaseDisasterInvoker) Invoke() (*model.ExecuteCrossCloudReleaseDisasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCrossCloudReleaseDisasterResponse), nil
	}
}

type InstallKernelPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallKernelPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListApplicableInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAppliedHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAvailableFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListBackupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListBackupsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupsDetailsInvoker) Invoke() (*model.ListBackupsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsDetailsResponse), nil
	}
}

type ListBindedEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBindedEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCnInfosBeforeReduceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListComponentInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListConfigurationsDiffInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigurationsDiffInvoker) Invoke() (*model.ListConfigurationsDiffResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsDiffResponse), nil
	}
}

type ListDatabaseInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseInstancesInvoker) Invoke() (*model.ListDatabaseInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseInstancesResponse), nil
	}
}

type ListDatabaseRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseRolesInvoker) Invoke() (*model.ListDatabaseRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseRolesResponse), nil
	}
}

type ListDatabaseSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDatastoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatastoresInvoker) Invoke() (*model.ListDatastoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatastoresResponse), nil
	}
}

type ListDatastoresDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatastoresDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatastoresDetailsInvoker) Invoke() (*model.ListDatastoresDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatastoresDetailsResponse), nil
	}
}

type ListDbBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbBackupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDbFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDbUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDbUsersInvoker) Invoke() (*model.ListDbUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbUsersResponse), nil
	}
}

type ListDisasterRecoveryRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDisasterRecoveryRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDisasterRecoveryRecordInvoker) Invoke() (*model.ListDisasterRecoveryRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDisasterRecoveryRecordResponse), nil
	}
}

type ListEpsQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEpsQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEpsQuotasInvoker) Invoke() (*model.ListEpsQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEpsQuotasResponse), nil
	}
}

type ListFeaturesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeaturesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeaturesInvoker) Invoke() (*model.ListFeaturesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeaturesResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListFlavorsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlavorsDetailsInvoker) Invoke() (*model.ListFlavorsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsDetailsResponse), nil
	}
}

type ListGaussDbDatastoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussDbDatastoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListHistoryOperationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceDetailsInvoker) Invoke() (*model.ListInstanceDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceDetailsResponse), nil
	}
}

type ListInstanceEngineDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceEngineDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceEngineDetailInvoker) Invoke() (*model.ListInstanceEngineDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceEngineDetailResponse), nil
	}
}

type ListInstanceErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceErrorLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstancesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListKernelPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListParamGroupTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListParamGroupTemplatesInvoker) Invoke() (*model.ListParamGroupTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListParamGroupTemplatesResponse), nil
	}
}

type ListParameterGroupTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListParameterGroupTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListParameterGroupTemplatesInvoker) Invoke() (*model.ListParameterGroupTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListParameterGroupTemplatesResponse), nil
	}
}

type ListPluginExtensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginExtensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPredefinedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ListRecycleInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecycleInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecycleInstancesInvoker) Invoke() (*model.ListRecycleInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleInstancesResponse), nil
	}
}

type ListRecycleInstancesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecycleInstancesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecycleInstancesDetailsInvoker) Invoke() (*model.ListRecycleInstancesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleInstancesDetailsResponse), nil
	}
}

type ListRestorableInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestorableInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRestorableInstancesInvoker) Invoke() (*model.ListRestorableInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestorableInstancesResponse), nil
	}
}

type ListRestorableInstancesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestorableInstancesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRestorableInstancesDetailsInvoker) Invoke() (*model.ListRestorableInstancesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestorableInstancesDetailsResponse), nil
	}
}

type ListRestoreTimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTimesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRestoreTimesInvoker) Invoke() (*model.ListRestoreTimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTimesResponse), nil
	}
}

type ListScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduleTaskInvoker) Invoke() (*model.ListScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduleTaskResponse), nil
	}
}

type ListStorageTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSupportKernelPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListTopIoTrafficsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ModifyEpsQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ResetConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetConfigurationInvoker) Invoke() (*model.ResetConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConfigurationResponse), nil
	}
}

type ResetDrConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetDrConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetDrConfigInvoker) Invoke() (*model.ResetDrConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetDrConfigResponse), nil
	}
}

type ResetPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ResizeInstanceFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestartInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestoreInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ResumePluginExtensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunInstanceActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SearchAutoEnlargePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetDbUserPwdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetKernelPluginLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetNewBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBalanceStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBatchUpgradeCandidateVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowConfigurationDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigurationDetailInvoker) Invoke() (*model.ShowConfigurationDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationDetailResponse), nil
	}
}

type ShowCrossCloudDisasterInstanceMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCrossCloudDisasterInstanceMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCrossCloudDisasterInstanceMonitorInvoker) Invoke() (*model.ShowCrossCloudDisasterInstanceMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCrossCloudDisasterInstanceMonitorResponse), nil
	}
}

type ShowCrossCloudDisasterRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCrossCloudDisasterRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCrossCloudDisasterRelationsInvoker) Invoke() (*model.ShowCrossCloudDisasterRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCrossCloudDisasterRelationsResponse), nil
	}
}

type ShowDeploymentFormInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentFormInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowErrorLogSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowInstanceConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowInstanceDiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowInstanceParamGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceParamGroupInvoker) Invoke() (*model.ShowInstanceParamGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceParamGroupResponse), nil
	}
}

type ShowInstanceParamGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceParamGroupDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceParamGroupDetailInvoker) Invoke() (*model.ShowInstanceParamGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceParamGroupDetailResponse), nil
	}
}

type ShowInstanceSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowParameterGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowParameterGroupDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowParameterGroupDetailInvoker) Invoke() (*model.ShowParameterGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowParameterGroupDetailResponse), nil
	}
}

type ShowProjectQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowSlowLogDownloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlowLogDownloadInvoker) Invoke() (*model.ShowSlowLogDownloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowLogDownloadResponse), nil
	}
}

type ShowSourceInstanceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSourceInstanceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSourceInstanceDetailInvoker) Invoke() (*model.ShowSourceInstanceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSourceInstanceDetailResponse), nil
	}
}

type ShowSslCertDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSslCertDownloadLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowUpgradeCandidateVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUpgradeCandidateVersionsInvoker) Invoke() (*model.ShowUpgradeCandidateVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpgradeCandidateVersionsResponse), nil
	}
}

type ShowUpgradeCandidateVersionsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpgradeCandidateVersionsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUpgradeCandidateVersionsDetailsInvoker) Invoke() (*model.ShowUpgradeCandidateVersionsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpgradeCandidateVersionsDetailsResponse), nil
	}
}

type StartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartInstanceInvoker) Invoke() (*model.StartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResponse), nil
	}
}

type StartMysqlCompatibilityInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartMysqlCompatibilityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartMysqlCompatibilityInvoker) Invoke() (*model.StartMysqlCompatibilityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartMysqlCompatibilityResponse), nil
	}
}

type StopBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopBackupInvoker) Invoke() (*model.StopBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBackupResponse), nil
	}
}

type StopInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInstanceInvoker) Invoke() (*model.StopInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInstanceResponse), nil
	}
}

type SwitchConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SwitchShardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchShardInvoker) Invoke() (*model.SwitchShardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchShardResponse), nil
	}
}

type UpdateFeaturesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFeaturesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFeaturesInvoker) Invoke() (*model.UpdateFeaturesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFeaturesResponse), nil
	}
}

type UpdateInstanceConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateInstanceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceNameInvoker) Invoke() (*model.UpdateInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNameResponse), nil
	}
}

type UpdateMysqlCompatibilityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMysqlCompatibilityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMysqlCompatibilityInvoker) Invoke() (*model.UpdateMysqlCompatibilityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMysqlCompatibilityResponse), nil
	}
}

type UpgradeInstanceVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeInstanceVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpgradeInstancesVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ValidateParaGroupNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ValidateWeakPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateWeakPasswordInvoker) Invoke() (*model.ValidateWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateWeakPasswordResponse), nil
	}
}

type CreateLimitTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLimitTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLimitTaskInvoker) Invoke() (*model.CreateLimitTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLimitTaskResponse), nil
	}
}

type DeleteLimitTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLimitTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLimitTaskInvoker) Invoke() (*model.DeleteLimitTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLimitTaskResponse), nil
	}
}

type ListLimitTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLimitTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLimitTaskInvoker) Invoke() (*model.ListLimitTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLimitTaskResponse), nil
	}
}

type ListNodeLimitSqlModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeLimitSqlModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodeLimitSqlModelInvoker) Invoke() (*model.ListNodeLimitSqlModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeLimitSqlModelResponse), nil
	}
}

type ShowLimitTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLimitTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLimitTaskInvoker) Invoke() (*model.ShowLimitTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLimitTaskResponse), nil
	}
}

type SyncLimitDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncLimitDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncLimitDataInvoker) Invoke() (*model.SyncLimitDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncLimitDataResponse), nil
	}
}

type UpdateLimitTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLimitTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLimitTaskInvoker) Invoke() (*model.UpdateLimitTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLimitTaskResponse), nil
	}
}
