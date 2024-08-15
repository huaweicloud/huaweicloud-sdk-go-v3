package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"
)

type BatchCreateOrDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteTagsInvoker) Invoke() (*model.BatchCreateOrDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

type BatchDeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstancesInvoker) Invoke() (*model.BatchDeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstancesResponse), nil
	}
}

type BatchRestartOnlineMigrationTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRestartOnlineMigrationTasksInvoker) Invoke() (*model.BatchRestartOnlineMigrationTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRestartOnlineMigrationTasksResponse), nil
	}
}

type BatchShowNodesInformationInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowNodesInformationInvoker) Invoke() (*model.BatchShowNodesInformationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowNodesInformationResponse), nil
	}
}

type BatchStopMigrationTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopMigrationTasksInvoker) Invoke() (*model.BatchStopMigrationTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopMigrationTasksResponse), nil
	}
}

type ChangeMasterStandbyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeMasterStandbyInvoker) Invoke() (*model.ChangeMasterStandbyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeMasterStandbyResponse), nil
	}
}

type ChangeMasterStandbyAsyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeMasterStandbyAsyncInvoker) Invoke() (*model.ChangeMasterStandbyAsyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeMasterStandbyAsyncResponse), nil
	}
}

type CopyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyInstanceInvoker) Invoke() (*model.CopyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyInstanceResponse), nil
	}
}

type CreateAclAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAclAccountInvoker) Invoke() (*model.CreateAclAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAclAccountResponse), nil
	}
}

type CreateAutoExpireScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoExpireScanTaskInvoker) Invoke() (*model.CreateAutoExpireScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoExpireScanTaskResponse), nil
	}
}

type CreateBigkeyScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBigkeyScanTaskInvoker) Invoke() (*model.CreateBigkeyScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBigkeyScanTaskResponse), nil
	}
}

type CreateCustomTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomTemplateInvoker) Invoke() (*model.CreateCustomTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomTemplateResponse), nil
	}
}

type CreateDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnosisTaskInvoker) Invoke() (*model.CreateDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnosisTaskResponse), nil
	}
}

type CreateHotkeyScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHotkeyScanTaskInvoker) Invoke() (*model.CreateHotkeyScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHotkeyScanTaskResponse), nil
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

type CreateMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMigrationTaskInvoker) Invoke() (*model.CreateMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMigrationTaskResponse), nil
	}
}

type CreateOnlineMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOnlineMigrationTaskInvoker) Invoke() (*model.CreateOnlineMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOnlineMigrationTaskResponse), nil
	}
}

type CreateRedislogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRedislogInvoker) Invoke() (*model.CreateRedislogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRedislogResponse), nil
	}
}

type CreateRedislogDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRedislogDownloadLinkInvoker) Invoke() (*model.CreateRedislogDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRedislogDownloadLinkResponse), nil
	}
}

type CreateResizeOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResizeOrderInvoker) Invoke() (*model.CreateResizeOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResizeOrderResponse), nil
	}
}

type DeleteAclAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAclAccountInvoker) Invoke() (*model.DeleteAclAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAclAccountResponse), nil
	}
}

type DeleteBackgroundTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackgroundTaskInvoker) Invoke() (*model.DeleteBackgroundTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackgroundTaskResponse), nil
	}
}

type DeleteBackupFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupFileInvoker) Invoke() (*model.DeleteBackupFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupFileResponse), nil
	}
}

type DeleteBigkeyScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBigkeyScanTaskInvoker) Invoke() (*model.DeleteBigkeyScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBigkeyScanTaskResponse), nil
	}
}

type DeleteCenterTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCenterTaskInvoker) Invoke() (*model.DeleteCenterTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCenterTaskResponse), nil
	}
}

type DeleteConfigTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigTemplateInvoker) Invoke() (*model.DeleteConfigTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigTemplateResponse), nil
	}
}

type DeleteDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDiagnosisTaskInvoker) Invoke() (*model.DeleteDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDiagnosisTaskResponse), nil
	}
}

type DeleteHotkeyScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHotkeyScanTaskInvoker) Invoke() (*model.DeleteHotkeyScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHotkeyScanTaskResponse), nil
	}
}

type DeleteIpFromDomainNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIpFromDomainNameInvoker) Invoke() (*model.DeleteIpFromDomainNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIpFromDomainNameResponse), nil
	}
}

type DeleteMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMigrationTaskInvoker) Invoke() (*model.DeleteMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMigrationTaskResponse), nil
	}
}

type DeleteSingleInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSingleInstanceInvoker) Invoke() (*model.DeleteSingleInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSingleInstanceResponse), nil
	}
}

type DownloadSslCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadSslCertInvoker) Invoke() (*model.DownloadSslCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSslCertResponse), nil
	}
}

type ExchangeInstanceIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExchangeInstanceIpInvoker) Invoke() (*model.ExchangeInstanceIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExchangeInstanceIpResponse), nil
	}
}

type ExecuteClusterSwitchoverInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteClusterSwitchoverInvoker) Invoke() (*model.ExecuteClusterSwitchoverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteClusterSwitchoverResponse), nil
	}
}

type ExecuteCommandMobilizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCommandMobilizationInvoker) Invoke() (*model.ExecuteCommandMobilizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCommandMobilizationResponse), nil
	}
}

type ExportExcelJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportExcelJobInvoker) Invoke() (*model.ExportExcelJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportExcelJobResponse), nil
	}
}

type ExportInstancesTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportInstancesTaskInvoker) Invoke() (*model.ExportInstancesTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInstancesTaskResponse), nil
	}
}

type HangUpClientsInvoker struct {
	*invoker.BaseInvoker
}

func (i *HangUpClientsInvoker) Invoke() (*model.HangUpClientsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HangUpClientsResponse), nil
	}
}

type HangUpKillAllClientsInvoker struct {
	*invoker.BaseInvoker
}

func (i *HangUpKillAllClientsInvoker) Invoke() (*model.HangUpKillAllClientsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HangUpKillAllClientsResponse), nil
	}
}

type ListAclAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAclAccountsInvoker) Invoke() (*model.ListAclAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclAccountsResponse), nil
	}
}

type ListAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableZonesInvoker) Invoke() (*model.ListAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableZonesResponse), nil
	}
}

type ListBackgroundTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackgroundTaskInvoker) Invoke() (*model.ListBackgroundTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackgroundTaskResponse), nil
	}
}

type ListBackupFileLinksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupFileLinksInvoker) Invoke() (*model.ListBackupFileLinksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupFileLinksResponse), nil
	}
}

type ListBackupRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupRecordsInvoker) Invoke() (*model.ListBackupRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupRecordsResponse), nil
	}
}

type ListBigkeyScanTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBigkeyScanTasksInvoker) Invoke() (*model.ListBigkeyScanTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBigkeyScanTasksResponse), nil
	}
}

type ListCenterTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCenterTaskInvoker) Invoke() (*model.ListCenterTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCenterTaskResponse), nil
	}
}

type ListClientsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClientsInvoker) Invoke() (*model.ListClientsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClientsResponse), nil
	}
}

type ListConfigHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigHistoriesInvoker) Invoke() (*model.ListConfigHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigHistoriesResponse), nil
	}
}

type ListConfigTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigTemplatesInvoker) Invoke() (*model.ListConfigTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigTemplatesResponse), nil
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

type ListDiagnosisTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnosisTasksInvoker) Invoke() (*model.ListDiagnosisTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnosisTasksResponse), nil
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

type ListGroupReplicationInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupReplicationInfoInvoker) Invoke() (*model.ListGroupReplicationInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupReplicationInfoResponse), nil
	}
}

type ListHotKeyScanTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHotKeyScanTasksInvoker) Invoke() (*model.ListHotKeyScanTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHotKeyScanTasksResponse), nil
	}
}

type ListInstanceOperationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceOperationsInvoker) Invoke() (*model.ListInstanceOperationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceOperationsResponse), nil
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

type ListMaintenanceWindowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMaintenanceWindowsInvoker) Invoke() (*model.ListMaintenanceWindowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMaintenanceWindowsResponse), nil
	}
}

type ListMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigrationTaskInvoker) Invoke() (*model.ListMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigrationTaskResponse), nil
	}
}

type ListMigrationTaskLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigrationTaskLogsInvoker) Invoke() (*model.ListMigrationTaskLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigrationTaskLogsResponse), nil
	}
}

type ListMonitoredObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitoredObjectsInvoker) Invoke() (*model.ListMonitoredObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitoredObjectsResponse), nil
	}
}

type ListMonitoredObjectsOfInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitoredObjectsOfInstanceInvoker) Invoke() (*model.ListMonitoredObjectsOfInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitoredObjectsOfInstanceResponse), nil
	}
}

type ListNumberOfInstancesInDifferentStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNumberOfInstancesInDifferentStatusInvoker) Invoke() (*model.ListNumberOfInstancesInDifferentStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNumberOfInstancesInDifferentStatusResponse), nil
	}
}

type ListRedislogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRedislogInvoker) Invoke() (*model.ListRedislogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRedislogResponse), nil
	}
}

type ListRestoreRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreRecordsInvoker) Invoke() (*model.ListRestoreRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreRecordsResponse), nil
	}
}

type ListSlowlogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSlowlogInvoker) Invoke() (*model.ListSlowlogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSlowlogResponse), nil
	}
}

type ListStatisticsOfRunningInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatisticsOfRunningInstancesInvoker) Invoke() (*model.ListStatisticsOfRunningInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatisticsOfRunningInstancesResponse), nil
	}
}

type ListTagsOfTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsOfTenantInvoker) Invoke() (*model.ListTagsOfTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsOfTenantResponse), nil
	}
}

type LoginWebCliInvoker struct {
	*invoker.BaseInvoker
}

func (i *LoginWebCliInvoker) Invoke() (*model.LoginWebCliResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginWebCliResponse), nil
	}
}

type LogoffWebCliInvoker struct {
	*invoker.BaseInvoker
}

func (i *LogoffWebCliInvoker) Invoke() (*model.LogoffWebCliResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogoffWebCliResponse), nil
	}
}

type ResetAclAccountPassWordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetAclAccountPassWordInvoker) Invoke() (*model.ResetAclAccountPassWordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetAclAccountPassWordResponse), nil
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

type RestartOrFlushInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartOrFlushInstancesInvoker) Invoke() (*model.RestartOrFlushInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartOrFlushInstancesResponse), nil
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

type ScanClientsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScanClientsInvoker) Invoke() (*model.ScanClientsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScanClientsResponse), nil
	}
}

type ScanExpireKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScanExpireKeyInvoker) Invoke() (*model.ScanExpireKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScanExpireKeyResponse), nil
	}
}

type SetOnlineMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetOnlineMigrationTaskInvoker) Invoke() (*model.SetOnlineMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetOnlineMigrationTaskResponse), nil
	}
}

type ShowBackgroundTaskProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackgroundTaskProgressInvoker) Invoke() (*model.ShowBackgroundTaskProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackgroundTaskProgressResponse), nil
	}
}

type ShowBigkeyAutoscanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBigkeyAutoscanConfigInvoker) Invoke() (*model.ShowBigkeyAutoscanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBigkeyAutoscanConfigResponse), nil
	}
}

type ShowBigkeyScanTaskDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBigkeyScanTaskDetailsInvoker) Invoke() (*model.ShowBigkeyScanTaskDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBigkeyScanTaskDetailsResponse), nil
	}
}

type ShowConfigHistoryDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigHistoryDetailInvoker) Invoke() (*model.ShowConfigHistoryDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigHistoryDetailResponse), nil
	}
}

type ShowConfigTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigTemplateInvoker) Invoke() (*model.ShowConfigTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigTemplateResponse), nil
	}
}

type ShowDiagnosisTaskDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisTaskDetailsInvoker) Invoke() (*model.ShowDiagnosisTaskDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisTaskDetailsResponse), nil
	}
}

type ShowExpireAutoScanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExpireAutoScanConfigInvoker) Invoke() (*model.ShowExpireAutoScanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExpireAutoScanConfigResponse), nil
	}
}

type ShowExpireKeyScanInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExpireKeyScanInfoInvoker) Invoke() (*model.ShowExpireKeyScanInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExpireKeyScanInfoResponse), nil
	}
}

type ShowHotkeyAutoscanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotkeyAutoscanConfigInvoker) Invoke() (*model.ShowHotkeyAutoscanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotkeyAutoscanConfigResponse), nil
	}
}

type ShowHotkeyTaskDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotkeyTaskDetailsInvoker) Invoke() (*model.ShowHotkeyTaskDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotkeyTaskDetailsResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowInstanceSslDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceSslDetailInvoker) Invoke() (*model.ShowInstanceSslDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceSslDetailResponse), nil
	}
}

type ShowJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInfoInvoker) Invoke() (*model.ShowJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobInfoResponse), nil
	}
}

type ShowMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigrationTaskInvoker) Invoke() (*model.ShowMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigrationTaskResponse), nil
	}
}

type ShowMigrationTaskStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigrationTaskStatsInvoker) Invoke() (*model.ShowMigrationTaskStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigrationTaskStatsResponse), nil
	}
}

type ShowNodesInformationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodesInformationInvoker) Invoke() (*model.ShowNodesInformationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodesInformationResponse), nil
	}
}

type ShowQuotaOfTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaOfTenantInvoker) Invoke() (*model.ShowQuotaOfTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaOfTenantResponse), nil
	}
}

type ShowReplicationStatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReplicationStatesInvoker) Invoke() (*model.ShowReplicationStatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReplicationStatesResponse), nil
	}
}

type ShowTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagsInvoker) Invoke() (*model.ShowTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagsResponse), nil
	}
}

type StartInstanceResizeCheckJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceResizeCheckJobInvoker) Invoke() (*model.StartInstanceResizeCheckJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResizeCheckJobResponse), nil
	}
}

type StopMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopMigrationTaskInvoker) Invoke() (*model.StopMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopMigrationTaskResponse), nil
	}
}

type StopMigrationTaskSyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopMigrationTaskSyncInvoker) Invoke() (*model.StopMigrationTaskSyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopMigrationTaskSyncResponse), nil
	}
}

type UpdateAclAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclAccountInvoker) Invoke() (*model.UpdateAclAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclAccountResponse), nil
	}
}

type UpdateAclAccountPassWordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclAccountPassWordInvoker) Invoke() (*model.UpdateAclAccountPassWordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclAccountPassWordResponse), nil
	}
}

type UpdateAclAccountRemarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclAccountRemarkInvoker) Invoke() (*model.UpdateAclAccountRemarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclAccountRemarkResponse), nil
	}
}

type UpdateBigkeyAutoscanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBigkeyAutoscanConfigInvoker) Invoke() (*model.UpdateBigkeyAutoscanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBigkeyAutoscanConfigResponse), nil
	}
}

type UpdateClientIpTransparentTransmissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClientIpTransparentTransmissionInvoker) Invoke() (*model.UpdateClientIpTransparentTransmissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClientIpTransparentTransmissionResponse), nil
	}
}

type UpdateConfigTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigTemplateInvoker) Invoke() (*model.UpdateConfigTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigTemplateResponse), nil
	}
}

type UpdateConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationsInvoker) Invoke() (*model.UpdateConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationsResponse), nil
	}
}

type UpdateExpireAutoScanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateExpireAutoScanConfigInvoker) Invoke() (*model.UpdateExpireAutoScanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateExpireAutoScanConfigResponse), nil
	}
}

type UpdateHotkeyAutoScanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHotkeyAutoScanConfigInvoker) Invoke() (*model.UpdateHotkeyAutoScanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHotkeyAutoScanConfigResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}

type UpdateInstanceBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceBandwidthInvoker) Invoke() (*model.UpdateInstanceBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceBandwidthResponse), nil
	}
}

type UpdateInstanceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigInvoker) Invoke() (*model.UpdateInstanceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConfigResponse), nil
	}
}

type UpdateMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMigrationTaskInvoker) Invoke() (*model.UpdateMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMigrationTaskResponse), nil
	}
}

type UpdatePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePasswordInvoker) Invoke() (*model.UpdatePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePasswordResponse), nil
	}
}

type UpdateSlavePriorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSlavePriorityInvoker) Invoke() (*model.UpdateSlavePriorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSlavePriorityResponse), nil
	}
}

type UpdateSslSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSslSwitchInvoker) Invoke() (*model.UpdateSslSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSslSwitchResponse), nil
	}
}

type ValidateDeletableReplicaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateDeletableReplicaInvoker) Invoke() (*model.ValidateDeletableReplicaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateDeletableReplicaResponse), nil
	}
}

type ShowIpWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpWhitelistInvoker) Invoke() (*model.ShowIpWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpWhitelistResponse), nil
	}
}

type UpdateIpWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpWhitelistInvoker) Invoke() (*model.UpdateIpWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpWhitelistResponse), nil
	}
}
