package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbfornosql/v3/model"
)

type ApplyConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyConfigurationInvoker) Invoke() (*model.ApplyConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyConfigurationResponse), nil
	}
}

type BatchTagActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagActionInvoker) Invoke() (*model.BatchTagActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagActionResponse), nil
	}
}

type CheckDisasterRecoveryOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDisasterRecoveryOperationInvoker) Invoke() (*model.CheckDisasterRecoveryOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDisasterRecoveryOperationResponse), nil
	}
}

type CheckWeekPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckWeekPasswordInvoker) Invoke() (*model.CheckWeekPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckWeekPasswordResponse), nil
	}
}

type CompareConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompareConfigurationInvoker) Invoke() (*model.CompareConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareConfigurationResponse), nil
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

type CreateBackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackInvoker) Invoke() (*model.CreateBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackResponse), nil
	}
}

type CreateColdVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateColdVolumeInvoker) Invoke() (*model.CreateColdVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateColdVolumeResponse), nil
	}
}

type CreateConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationInvoker) Invoke() (*model.CreateConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationResponse), nil
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

type CreateDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDisasterRecoveryInvoker) Invoke() (*model.CreateDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDisasterRecoveryResponse), nil
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

type DeleteBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupInvoker) Invoke() (*model.DeleteBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupResponse), nil
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

type DeleteDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDbUserInvoker) Invoke() (*model.DeleteDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDbUserResponse), nil
	}
}

type DeleteDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDisasterRecoveryInvoker) Invoke() (*model.DeleteDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDisasterRecoveryResponse), nil
	}
}

type DeleteEnlargeFailNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnlargeFailNodeInvoker) Invoke() (*model.DeleteEnlargeFailNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnlargeFailNodeResponse), nil
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

type DeleteInstancesSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstancesSessionInvoker) Invoke() (*model.DeleteInstancesSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesSessionResponse), nil
	}
}

type DeleteLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLtsConfigsInvoker) Invoke() (*model.DeleteLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLtsConfigsResponse), nil
	}
}

type ExpandInstanceNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandInstanceNodeInvoker) Invoke() (*model.ExpandInstanceNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandInstanceNodeResponse), nil
	}
}

type ListAvailableFlavorInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableFlavorInfosInvoker) Invoke() (*model.ListAvailableFlavorInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableFlavorInfosResponse), nil
	}
}

type ListCassandraSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCassandraSlowLogsInvoker) Invoke() (*model.ListCassandraSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCassandraSlowLogsResponse), nil
	}
}

type ListConfigurationDatastoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationDatastoresInvoker) Invoke() (*model.ListConfigurationDatastoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationDatastoresResponse), nil
	}
}

type ListConfigurationTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationTemplatesInvoker) Invoke() (*model.ListConfigurationTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationTemplatesResponse), nil
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

type ListDedicatedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDedicatedResourcesInvoker) Invoke() (*model.ListDedicatedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDedicatedResourcesResponse), nil
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

type ListFlavorInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorInfosInvoker) Invoke() (*model.ListFlavorInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorInfosResponse), nil
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

type ListInfluxdbSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInfluxdbSlowLogsInvoker) Invoke() (*model.ListInfluxdbSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInfluxdbSlowLogsResponse), nil
	}
}

type ListInstanceDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceDatabasesInvoker) Invoke() (*model.ListInstanceDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceDatabasesResponse), nil
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

type ListInstancesByResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesByResourceTagsInvoker) Invoke() (*model.ListInstancesByResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesByResourceTagsResponse), nil
	}
}

type ListInstancesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesByTagsInvoker) Invoke() (*model.ListInstancesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesByTagsResponse), nil
	}
}

type ListInstancesSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesSessionInvoker) Invoke() (*model.ListInstancesSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesSessionResponse), nil
	}
}

type ListInstancesSessionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesSessionStatisticsInvoker) Invoke() (*model.ListInstancesSessionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesSessionStatisticsResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsConfigsInvoker) Invoke() (*model.ListLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsConfigsResponse), nil
	}
}

type ListMongodbErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMongodbErrorLogsInvoker) Invoke() (*model.ListMongodbErrorLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMongodbErrorLogsResponse), nil
	}
}

type ListMongodbSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMongodbSlowLogsInvoker) Invoke() (*model.ListMongodbSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMongodbSlowLogsResponse), nil
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

type ListRedisSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRedisSlowLogsInvoker) Invoke() (*model.ListRedisSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRedisSlowLogsResponse), nil
	}
}

type ListRestoreDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreDatabasesInvoker) Invoke() (*model.ListRestoreDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreDatabasesResponse), nil
	}
}

type ListRestoreTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTablesInvoker) Invoke() (*model.ListRestoreTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTablesResponse), nil
	}
}

type ListRestoreTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTimeInvoker) Invoke() (*model.ListRestoreTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTimeResponse), nil
	}
}

type ListSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSlowLogsInvoker) Invoke() (*model.ListSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSlowLogsResponse), nil
	}
}

type ModifyDbUserPrivilegeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyDbUserPrivilegeInvoker) Invoke() (*model.ModifyDbUserPrivilegeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyDbUserPrivilegeResponse), nil
	}
}

type ModifyEpsQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyEpsQuotasInvoker) Invoke() (*model.ModifyEpsQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyEpsQuotasResponse), nil
	}
}

type ModifyPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyPortInvoker) Invoke() (*model.ModifyPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyPortResponse), nil
	}
}

type ModifyPublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyPublicIpInvoker) Invoke() (*model.ModifyPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyPublicIpResponse), nil
	}
}

type ModifyVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyVolumeInvoker) Invoke() (*model.ModifyVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyVolumeResponse), nil
	}
}

type OfflineNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *OfflineNodesInvoker) Invoke() (*model.OfflineNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.OfflineNodesResponse), nil
	}
}

type PauseResumeDataSynchronizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseResumeDataSynchronizationInvoker) Invoke() (*model.PauseResumeDataSynchronizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseResumeDataSynchronizationResponse), nil
	}
}

type ResetDbUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetDbUserPasswordInvoker) Invoke() (*model.ResetDbUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetDbUserPasswordResponse), nil
	}
}

type ResetParamGroupTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetParamGroupTemplateInvoker) Invoke() (*model.ResetParamGroupTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetParamGroupTemplateResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeColdVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeColdVolumeInvoker) Invoke() (*model.ResizeColdVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeColdVolumeResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type ResizeInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceVolumeInvoker) Invoke() (*model.ResizeInstanceVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceVolumeResponse), nil
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

type RestoreExistingInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreExistingInstanceInvoker) Invoke() (*model.RestoreExistingInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreExistingInstanceResponse), nil
	}
}

type SaveLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveLtsConfigsInvoker) Invoke() (*model.SaveLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveLtsConfigsResponse), nil
	}
}

type SetAutoEnlargePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAutoEnlargePolicyInvoker) Invoke() (*model.SetAutoEnlargePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAutoEnlargePolicyResponse), nil
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

type ShowAllInstancesBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllInstancesBackupsInvoker) Invoke() (*model.ShowAllInstancesBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllInstancesBackupsResponse), nil
	}
}

type ShowAllInstancesBackupsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllInstancesBackupsNewInvoker) Invoke() (*model.ShowAllInstancesBackupsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllInstancesBackupsNewResponse), nil
	}
}

type ShowApplicableInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicableInstancesInvoker) Invoke() (*model.ShowApplicableInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicableInstancesResponse), nil
	}
}

type ShowApplyHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplyHistoryInvoker) Invoke() (*model.ShowApplyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplyHistoryResponse), nil
	}
}

type ShowAutoEnlargePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoEnlargePolicyInvoker) Invoke() (*model.ShowAutoEnlargePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoEnlargePolicyResponse), nil
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

type ShowElbIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowElbIpGroupInvoker) Invoke() (*model.ShowElbIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowElbIpGroupResponse), nil
	}
}

type ShowErrorLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowErrorLogInvoker) Invoke() (*model.ShowErrorLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowErrorLogResponse), nil
	}
}

type ShowHighRiskCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHighRiskCommandsInvoker) Invoke() (*model.ShowHighRiskCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHighRiskCommandsResponse), nil
	}
}

type ShowInstanceBiactiveRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceBiactiveRegionsInvoker) Invoke() (*model.ShowInstanceBiactiveRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceBiactiveRegionsResponse), nil
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

type ShowInstanceRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceRoleInvoker) Invoke() (*model.ShowInstanceRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceRoleResponse), nil
	}
}

type ShowIpNumRequirementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpNumRequirementInvoker) Invoke() (*model.ShowIpNumRequirementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpNumRequirementResponse), nil
	}
}

type ShowModifyHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModifyHistoryInvoker) Invoke() (*model.ShowModifyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModifyHistoryResponse), nil
	}
}

type ShowPasswordlessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPasswordlessConfigInvoker) Invoke() (*model.ShowPasswordlessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPasswordlessConfigResponse), nil
	}
}

type ShowPauseResumeStutusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPauseResumeStutusInvoker) Invoke() (*model.ShowPauseResumeStutusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPauseResumeStutusResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
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

type ShowRedisBigKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRedisBigKeysInvoker) Invoke() (*model.ShowRedisBigKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRedisBigKeysResponse), nil
	}
}

type ShowRestorableListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRestorableListInvoker) Invoke() (*model.ShowRestorableListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRestorableListResponse), nil
	}
}

type ShowSlowLogDesensitizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowLogDesensitizationInvoker) Invoke() (*model.ShowSlowLogDesensitizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowLogDesensitizationResponse), nil
	}
}

type ShrinkInstanceNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkInstanceNodeInvoker) Invoke() (*model.ShrinkInstanceNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkInstanceNodeResponse), nil
	}
}

type SwitchIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchIpGroupInvoker) Invoke() (*model.SwitchIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchIpGroupResponse), nil
	}
}

type SwitchSlowlogDesensitizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSlowlogDesensitizationInvoker) Invoke() (*model.SwitchSlowlogDesensitizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

type SwitchSslInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSslInvoker) Invoke() (*model.SwitchSslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSslResponse), nil
	}
}

type SwitchToMasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchToMasterInvoker) Invoke() (*model.SwitchToMasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchToMasterResponse), nil
	}
}

type SwitchToSlaveInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchToSlaveInvoker) Invoke() (*model.SwitchToSlaveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchToSlaveResponse), nil
	}
}

type UpdateClientNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClientNetworkInvoker) Invoke() (*model.UpdateClientNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClientNetworkResponse), nil
	}
}

type UpdateConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationInvoker) Invoke() (*model.UpdateConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationResponse), nil
	}
}

type UpdateDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabasesInvoker) Invoke() (*model.UpdateDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabasesResponse), nil
	}
}

type UpdateHighRiskCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHighRiskCommandsInvoker) Invoke() (*model.UpdateHighRiskCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHighRiskCommandsResponse), nil
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

type UpdatePasswordlessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePasswordlessConfigInvoker) Invoke() (*model.UpdatePasswordlessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePasswordlessConfigResponse), nil
	}
}

type UpdateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityGroupInvoker) Invoke() (*model.UpdateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityGroupResponse), nil
	}
}

type UpgradeDbVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeDbVersionInvoker) Invoke() (*model.UpgradeDbVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeDbVersionResponse), nil
	}
}

type ListApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionInvoker) Invoke() (*model.ListApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}
