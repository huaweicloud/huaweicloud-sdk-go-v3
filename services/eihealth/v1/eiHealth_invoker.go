package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v1/model"
)

type CreateCpiTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCpiTaskInvoker) Invoke() (*model.CreateCpiTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCpiTaskResponse), nil
	}
}

type ShowCpiTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCpiTaskResultInvoker) Invoke() (*model.ShowCpiTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCpiTaskResultResponse), nil
	}
}

type CreateGenerationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGenerationTaskInvoker) Invoke() (*model.CreateGenerationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGenerationTaskResponse), nil
	}
}

type ShowGenerationTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGenerationTaskResultInvoker) Invoke() (*model.ShowGenerationTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGenerationTaskResultResponse), nil
	}
}

type BatchImportAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportAppInvoker) Invoke() (*model.BatchImportAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportAppResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type ListAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppInvoker) Invoke() (*model.ListAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppResponse), nil
	}
}

type PublishAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishAppInvoker) Invoke() (*model.PublishAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishAppResponse), nil
	}
}

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type SubscribeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeAppInvoker) Invoke() (*model.SubscribeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeAppResponse), nil
	}
}

type UpdateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type DeleteAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetVersionInvoker) Invoke() (*model.DeleteAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetVersionResponse), nil
	}
}

type ExecuteAssetActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteAssetActionInvoker) Invoke() (*model.ExecuteAssetActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteAssetActionResponse), nil
	}
}

type ListAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetInvoker) Invoke() (*model.ListAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetResponse), nil
	}
}

type ListPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPropertyInvoker) Invoke() (*model.ListPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPropertyResponse), nil
	}
}

type ShowAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetInvoker) Invoke() (*model.ShowAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetResponse), nil
	}
}

type ShowAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetVersionInvoker) Invoke() (*model.ShowAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetVersionResponse), nil
	}
}

type UpdateAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssetVersionInvoker) Invoke() (*model.UpdateAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssetVersionResponse), nil
	}
}

type CreateAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoJobInvoker) Invoke() (*model.CreateAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoJobResponse), nil
	}
}

type DeleteAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAutoJobInvoker) Invoke() (*model.DeleteAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAutoJobResponse), nil
	}
}

type ListAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoJobInvoker) Invoke() (*model.ListAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoJobResponse), nil
	}
}

type ShowAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoJobInvoker) Invoke() (*model.ShowAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoJobResponse), nil
	}
}

type StartAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAutoJobInvoker) Invoke() (*model.StartAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAutoJobResponse), nil
	}
}

type StopAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAutoJobInvoker) Invoke() (*model.StopAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAutoJobResponse), nil
	}
}

type UpdateAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoJobInvoker) Invoke() (*model.UpdateAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoJobResponse), nil
	}
}

type CreateComputingResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComputingResourceInvoker) Invoke() (*model.CreateComputingResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComputingResourceResponse), nil
	}
}

type DeleteComputingResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComputingResourceInvoker) Invoke() (*model.DeleteComputingResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComputingResourceResponse), nil
	}
}

type ListComputingResourceFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingResourceFlavorsInvoker) Invoke() (*model.ListComputingResourceFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingResourceFlavorsResponse), nil
	}
}

type ListComputingResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingResourcesInvoker) Invoke() (*model.ListComputingResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingResourcesResponse), nil
	}
}

type RebootNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootNodeInvoker) Invoke() (*model.RebootNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootNodeResponse), nil
	}
}

type ShowBmsDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBmsDevicesInvoker) Invoke() (*model.ShowBmsDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBmsDevicesResponse), nil
	}
}

type ShowEvsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEvsQuotaInvoker) Invoke() (*model.ShowEvsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEvsQuotaResponse), nil
	}
}

type ShowLeftQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLeftQuotaInvoker) Invoke() (*model.ShowLeftQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLeftQuotaResponse), nil
	}
}

type ShowScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduleInvoker) Invoke() (*model.ShowScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduleResponse), nil
	}
}

type StartNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartNodeInvoker) Invoke() (*model.StartNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartNodeResponse), nil
	}
}

type StopNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopNodeInvoker) Invoke() (*model.StopNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopNodeResponse), nil
	}
}

type UpdateScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduleInvoker) Invoke() (*model.UpdateScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduleResponse), nil
	}
}

type CreateBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackupInvoker) Invoke() (*model.CreateBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackupResponse), nil
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

type ListBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupInvoker) Invoke() (*model.ListBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupResponse), nil
	}
}

type RestoreBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreBackupInvoker) Invoke() (*model.RestoreBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreBackupResponse), nil
	}
}

type ShowBackupPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPathInvoker) Invoke() (*model.ShowBackupPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPathResponse), nil
	}
}

type BatchDeleteDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDataInvoker) Invoke() (*model.BatchDeleteDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDataResponse), nil
	}
}

type CopyDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyDataInvoker) Invoke() (*model.CopyDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyDataResponse), nil
	}
}

type CreateDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataInvoker) Invoke() (*model.CreateDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataResponse), nil
	}
}

type ImportDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDataInvoker) Invoke() (*model.ImportDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDataResponse), nil
	}
}

type ImportNetworkDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportNetworkDataInvoker) Invoke() (*model.ImportNetworkDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportNetworkDataResponse), nil
	}
}

type ListBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBucketInvoker) Invoke() (*model.ListBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBucketResponse), nil
	}
}

type ListDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataInvoker) Invoke() (*model.ListDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataResponse), nil
	}
}

type PublishDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishDataInvoker) Invoke() (*model.PublishDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishDataResponse), nil
	}
}

type QuoteDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *QuoteDataInvoker) Invoke() (*model.QuoteDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuoteDataResponse), nil
	}
}

type ShowBucketStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBucketStorageInvoker) Invoke() (*model.ShowBucketStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBucketStorageResponse), nil
	}
}

type ShowDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataInvoker) Invoke() (*model.ShowDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataResponse), nil
	}
}

type ShowDataPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataPolicyInvoker) Invoke() (*model.ShowDataPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataPolicyResponse), nil
	}
}

type SubscribeDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeDataInvoker) Invoke() (*model.SubscribeDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeDataResponse), nil
	}
}

type UpdateDataPathPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataPathPolicyInvoker) Invoke() (*model.UpdateDataPathPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataPathPolicyResponse), nil
	}
}

type UpdateDataPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataPolicyInvoker) Invoke() (*model.UpdateDataPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataPolicyResponse), nil
	}
}

type UploadDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadDataInvoker) Invoke() (*model.UploadDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadDataResponse), nil
	}
}

type CancelDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelDataJobInvoker) Invoke() (*model.CancelDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelDataJobResponse), nil
	}
}

type DeleteDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataJobInvoker) Invoke() (*model.DeleteDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataJobResponse), nil
	}
}

type DownloadDataJobLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDataJobLogInvoker) Invoke() (*model.DownloadDataJobLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDataJobLogResponse), nil
	}
}

type ListCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckpointInvoker) Invoke() (*model.ListCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckpointResponse), nil
	}
}

type ListDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataJobInvoker) Invoke() (*model.ListDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataJobResponse), nil
	}
}

type RetryDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryDataJobInvoker) Invoke() (*model.RetryDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryDataJobResponse), nil
	}
}

type ShowDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataJobInvoker) Invoke() (*model.ShowDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataJobResponse), nil
	}
}

type CreateDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseDataInvoker) Invoke() (*model.CreateDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseDataResponse), nil
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

type DeleteDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseDataInvoker) Invoke() (*model.DeleteDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseDataResponse), nil
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

type ImportDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDatabaseDataInvoker) Invoke() (*model.ImportDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDatabaseDataResponse), nil
	}
}

type ListDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseDataInvoker) Invoke() (*model.ListDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseDataResponse), nil
	}
}

type ListInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInvoker) Invoke() (*model.ListInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResponse), nil
	}
}

type QuoteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *QuoteInstanceInvoker) Invoke() (*model.QuoteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuoteInstanceResponse), nil
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

type UpdateDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseDataInvoker) Invoke() (*model.UpdateDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseDataResponse), nil
	}
}

type CreateDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseResourceInvoker) Invoke() (*model.CreateDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResourceResponse), nil
	}
}

type DeleteDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseResourceInvoker) Invoke() (*model.DeleteDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResourceResponse), nil
	}
}

type ListDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseResourceInvoker) Invoke() (*model.ListDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseResourceResponse), nil
	}
}

type ListDatabaseResourceFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseResourceFlavorInvoker) Invoke() (*model.ListDatabaseResourceFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseResourceFlavorResponse), nil
	}
}

type BatchDeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagInvoker) Invoke() (*model.BatchDeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagResponse), nil
	}
}

type CreateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageInvoker) Invoke() (*model.CreateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageResponse), nil
	}
}

type DeleteImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageInvoker) Invoke() (*model.DeleteImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ImportImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportImageInvoker) Invoke() (*model.ImportImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportImageResponse), nil
	}
}

type ListImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageInvoker) Invoke() (*model.ListImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageResponse), nil
	}
}

type ListImageTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageTagInvoker) Invoke() (*model.ListImageTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageTagResponse), nil
	}
}

type PublishImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishImageInvoker) Invoke() (*model.PublishImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishImageResponse), nil
	}
}

type ShowDockerLoginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDockerLoginInvoker) Invoke() (*model.ShowDockerLoginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDockerLoginResponse), nil
	}
}

type SubscribeImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeImageInvoker) Invoke() (*model.SubscribeImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeImageResponse), nil
	}
}

type UpdateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImageInvoker) Invoke() (*model.UpdateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImageResponse), nil
	}
}

type ShowJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobConfigInvoker) Invoke() (*model.ShowJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobConfigResponse), nil
	}
}

type UpdateJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobConfigInvoker) Invoke() (*model.UpdateJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobConfigResponse), nil
	}
}

type BatchCancelJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCancelJobInvoker) Invoke() (*model.BatchCancelJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCancelJobResponse), nil
	}
}

type BatchDeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobInvoker) Invoke() (*model.BatchDeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobResponse), nil
	}
}

type BatchRetryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRetryJobInvoker) Invoke() (*model.BatchRetryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRetryJobResponse), nil
	}
}

type CancelJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelJobInvoker) Invoke() (*model.CancelJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelJobResponse), nil
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

type ExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteJobInvoker) Invoke() (*model.ExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteJobResponse), nil
	}
}

type ListJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobInvoker) Invoke() (*model.ListJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobResponse), nil
	}
}

type RetryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryJobInvoker) Invoke() (*model.RetryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryJobResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowJobEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobEventInvoker) Invoke() (*model.ShowJobEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobEventResponse), nil
	}
}

type ShowJobLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobLogInvoker) Invoke() (*model.ShowJobLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobLogResponse), nil
	}
}

type ShowTaskEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskEventsInvoker) Invoke() (*model.ShowTaskEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskEventsResponse), nil
	}
}

type ShowTaskInstanceEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstanceEventsInvoker) Invoke() (*model.ShowTaskInstanceEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstanceEventsResponse), nil
	}
}

type ShowTaskInstanceMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstanceMetricDataInvoker) Invoke() (*model.ShowTaskInstanceMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstanceMetricDataResponse), nil
	}
}

type ShowTaskInstancePodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstancePodInvoker) Invoke() (*model.ShowTaskInstancePodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstancePodResponse), nil
	}
}

type ShowTaskInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstancesInvoker) Invoke() (*model.ShowTaskInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstancesResponse), nil
	}
}

type UpdateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobInvoker) Invoke() (*model.UpdateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobResponse), nil
	}
}

type BatchDeleteLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLabelInvoker) Invoke() (*model.BatchDeleteLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLabelResponse), nil
	}
}

type CreateLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLabelInvoker) Invoke() (*model.CreateLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLabelResponse), nil
	}
}

type DeleteLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLabelInvoker) Invoke() (*model.DeleteLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLabelResponse), nil
	}
}

type ListLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLabelInvoker) Invoke() (*model.ListLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLabelResponse), nil
	}
}

type CreateLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLabelPageInvoker) Invoke() (*model.CreateLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLabelPageResponse), nil
	}
}

type DeleteLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLabelPageInvoker) Invoke() (*model.DeleteLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLabelPageResponse), nil
	}
}

type ListLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLabelPageInvoker) Invoke() (*model.ListLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLabelPageResponse), nil
	}
}

type BatchDeleteNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteNoticeInvoker) Invoke() (*model.BatchDeleteNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteNoticeResponse), nil
	}
}

type BatchUpdateNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNoticeInvoker) Invoke() (*model.BatchUpdateNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNoticeResponse), nil
	}
}

type CheckEmailConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckEmailConnectionInvoker) Invoke() (*model.CheckEmailConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckEmailConnectionResponse), nil
	}
}

type DeleteMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMessageEmailConfigInvoker) Invoke() (*model.DeleteMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMessageEmailConfigResponse), nil
	}
}

type ListMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageInvoker) Invoke() (*model.ListMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageResponse), nil
	}
}

type ListMessageStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageStatisticsInvoker) Invoke() (*model.ListMessageStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageStatisticsResponse), nil
	}
}

type ListNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNoticeInvoker) Invoke() (*model.ListNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNoticeResponse), nil
	}
}

type ShowMessageClearRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageClearRuleInvoker) Invoke() (*model.ShowMessageClearRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageClearRuleResponse), nil
	}
}

type ShowMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageEmailConfigInvoker) Invoke() (*model.ShowMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageEmailConfigResponse), nil
	}
}

type ShowMessageReceiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageReceiveConfigInvoker) Invoke() (*model.ShowMessageReceiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageReceiveConfigResponse), nil
	}
}

type UpdateMessageClearRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageClearRuleInvoker) Invoke() (*model.UpdateMessageClearRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageClearRuleResponse), nil
	}
}

type UpdateMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageEmailConfigInvoker) Invoke() (*model.UpdateMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageEmailConfigResponse), nil
	}
}

type UpdateMessageReceiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageReceiveConfigInvoker) Invoke() (*model.UpdateMessageReceiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageReceiveConfigResponse), nil
	}
}

type BatchUpdateNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNodeLabelInvoker) Invoke() (*model.BatchUpdateNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNodeLabelResponse), nil
	}
}

type ListClusterAllNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterAllNodeLabelInvoker) Invoke() (*model.ListClusterAllNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterAllNodeLabelResponse), nil
	}
}

type ListNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeLabelInvoker) Invoke() (*model.ListNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeLabelResponse), nil
	}
}

type ListPresetLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPresetLabelInvoker) Invoke() (*model.ListPresetLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPresetLabelResponse), nil
	}
}

type CreateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotebookInvoker) Invoke() (*model.CreateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotebookResponse), nil
	}
}

type DeleteNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotebookInvoker) Invoke() (*model.DeleteNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotebookResponse), nil
	}
}

type ListNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotebookInvoker) Invoke() (*model.ListNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotebookResponse), nil
	}
}

type ListNotebookToolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotebookToolInvoker) Invoke() (*model.ListNotebookToolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotebookToolResponse), nil
	}
}

type ShowNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookInvoker) Invoke() (*model.ShowNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookResponse), nil
	}
}

type ShowNotebookTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookTokenInvoker) Invoke() (*model.ShowNotebookTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookTokenResponse), nil
	}
}

type StopOrStartNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopOrStartNotebookInvoker) Invoke() (*model.StopOrStartNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopOrStartNotebookResponse), nil
	}
}

type UpdateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotebookInvoker) Invoke() (*model.UpdateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotebookResponse), nil
	}
}

type ListObsBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketInvoker) Invoke() (*model.ListObsBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketResponse), nil
	}
}

type ListObsBucketObjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketObjectInvoker) Invoke() (*model.ListObsBucketObjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketObjectResponse), nil
	}
}

type ShowOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOverviewInvoker) Invoke() (*model.ShowOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOverviewResponse), nil
	}
}

type CreatePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePerformanceResourceInvoker) Invoke() (*model.CreatePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePerformanceResourceResponse), nil
	}
}

type DeletePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePerformanceResourceInvoker) Invoke() (*model.DeletePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePerformanceResourceResponse), nil
	}
}

type ListPerformanceResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPerformanceResourcesInvoker) Invoke() (*model.ListPerformanceResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPerformanceResourcesResponse), nil
	}
}

type UpdatePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePerformanceResourceInvoker) Invoke() (*model.UpdatePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePerformanceResourceResponse), nil
	}
}

type BatchDeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMemberInvoker) Invoke() (*model.BatchDeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMemberResponse), nil
	}
}

type CreateProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectInvoker) Invoke() (*model.CreateProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectResponse), nil
	}
}

type DeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInvoker) Invoke() (*model.DeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberResponse), nil
	}
}

type DeleteProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectInvoker) Invoke() (*model.DeleteProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectResponse), nil
	}
}

type ListProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectInvoker) Invoke() (*model.ListProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectResponse), nil
	}
}

type ShowProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectInvoker) Invoke() (*model.ShowProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectResponse), nil
	}
}

type TransferProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *TransferProjectInvoker) Invoke() (*model.TransferProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TransferProjectResponse), nil
	}
}

type UpdateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberInvoker) Invoke() (*model.UpdateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberResponse), nil
	}
}

type UpdateProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectInvoker) Invoke() (*model.UpdateProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectResponse), nil
	}
}

type DownloadDataTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDataTraceInvoker) Invoke() (*model.DownloadDataTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDataTraceResponse), nil
	}
}

type ShowProjectTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTraceInvoker) Invoke() (*model.ShowProjectTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTraceResponse), nil
	}
}

type ShowProjectTraceDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTraceDataInvoker) Invoke() (*model.ShowProjectTraceDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTraceDataResponse), nil
	}
}

type ShowProjectTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTrackerInvoker) Invoke() (*model.ShowProjectTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTrackerResponse), nil
	}
}

type UpdateProjectTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectTrackerInvoker) Invoke() (*model.UpdateProjectTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectTrackerResponse), nil
	}
}

type BatchDownloadResourceStatDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDownloadResourceStatDataInvoker) Invoke() (*model.BatchDownloadResourceStatDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDownloadResourceStatDataResponse), nil
	}
}

type ShowResourceMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceMetricDataInvoker) Invoke() (*model.ShowResourceMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceMetricDataResponse), nil
	}
}

type DeleteStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStarInvoker) Invoke() (*model.DeleteStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStarResponse), nil
	}
}

type ListStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarInvoker) Invoke() (*model.ListStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarResponse), nil
	}
}

type UpdateStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStarInvoker) Invoke() (*model.UpdateStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStarResponse), nil
	}
}

type ListGlobalWorkflowStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalWorkflowStatisticInvoker) Invoke() (*model.ListGlobalWorkflowStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalWorkflowStatisticResponse), nil
	}
}

type ListPerformanceResourceStatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPerformanceResourceStatInvoker) Invoke() (*model.ListPerformanceResourceStatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPerformanceResourceStatResponse), nil
	}
}

type ListWorkflowStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowStatisticInvoker) Invoke() (*model.ListWorkflowStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowStatisticResponse), nil
	}
}

type ListStorageResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageResourcesInvoker) Invoke() (*model.ListStorageResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageResourcesResponse), nil
	}
}

type CreateStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStudyInvoker) Invoke() (*model.CreateStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStudyResponse), nil
	}
}

type CreateStudyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStudyJobInvoker) Invoke() (*model.CreateStudyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStudyJobResponse), nil
	}
}

type DeleteStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStudyInvoker) Invoke() (*model.DeleteStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStudyResponse), nil
	}
}

type ListStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStudyInvoker) Invoke() (*model.ListStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStudyResponse), nil
	}
}

type ListStudyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStudyJobInvoker) Invoke() (*model.ListStudyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStudyJobResponse), nil
	}
}

type Show3dStructureContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *Show3dStructureContentInvoker) Invoke() (*model.Show3dStructureContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Show3dStructureContentResponse), nil
	}
}

type ShowExtremumInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtremumInfoInvoker) Invoke() (*model.ShowExtremumInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtremumInfoResponse), nil
	}
}

type ListArchiveConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArchiveConfigsInvoker) Invoke() (*model.ListArchiveConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArchiveConfigsResponse), nil
	}
}

type ShowEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvInvoker) Invoke() (*model.ShowEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvResponse), nil
	}
}

type ShowVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVendorInvoker) Invoke() (*model.ShowVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVendorResponse), nil
	}
}

type UpdateArchiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateArchiveConfigInvoker) Invoke() (*model.UpdateArchiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateArchiveConfigResponse), nil
	}
}

type UpdateVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVendorInvoker) Invoke() (*model.UpdateVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVendorResponse), nil
	}
}

type ListQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaInvoker) Invoke() (*model.ListQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type ImportTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportTemplateInvoker) Invoke() (*model.ImportTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportTemplateResponse), nil
	}
}

type ListTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateInvoker) Invoke() (*model.ListTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateResponse), nil
	}
}

type ShowTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateInvoker) Invoke() (*model.ShowTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateResponse), nil
	}
}

type UploadTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadTemplateInvoker) Invoke() (*model.UploadTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadTemplateResponse), nil
	}
}

type ChangePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePasswordInvoker) Invoke() (*model.ChangePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePasswordResponse), nil
	}
}

type CheckTokenVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTokenVerificationInvoker) Invoke() (*model.CheckTokenVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTokenVerificationResponse), nil
	}
}

type CreateCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCodeInvoker) Invoke() (*model.CreateCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCodeResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ImportUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportUserInvoker) Invoke() (*model.ImportUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportUserResponse), nil
	}
}

type ListMfaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMfaInvoker) Invoke() (*model.ListMfaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMfaResponse), nil
	}
}

type ListUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserInvoker) Invoke() (*model.ListUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserResponse), nil
	}
}

type ShowUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserInvoker) Invoke() (*model.ShowUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserResponse), nil
	}
}

type ShowUserSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserSettingInvoker) Invoke() (*model.ShowUserSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserSettingResponse), nil
	}
}

type UpdateInitPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInitPasswordInvoker) Invoke() (*model.UpdateInitPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInitPasswordResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type UpdateUserByDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserByDomainInvoker) Invoke() (*model.UpdateUserByDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserByDomainResponse), nil
	}
}

type UpdateUserRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserRoleInvoker) Invoke() (*model.UpdateUserRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserRoleResponse), nil
	}
}

type UpdateUserSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserSettingInvoker) Invoke() (*model.UpdateUserSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserSettingResponse), nil
	}
}

type ValidateCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateCodeInvoker) Invoke() (*model.ValidateCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateCodeResponse), nil
	}
}

type ListVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVendorInvoker) Invoke() (*model.ListVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVendorResponse), nil
	}
}

type CreateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInvoker) Invoke() (*model.CreateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowResponse), nil
	}
}

type DeleteWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowInvoker) Invoke() (*model.DeleteWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowResponse), nil
	}
}

type ImportWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportWorkflowInvoker) Invoke() (*model.ImportWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportWorkflowResponse), nil
	}
}

type ListWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowInvoker) Invoke() (*model.ListWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowResponse), nil
	}
}

type PublishWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishWorkflowInvoker) Invoke() (*model.PublishWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishWorkflowResponse), nil
	}
}

type ShowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowInvoker) Invoke() (*model.ShowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowResponse), nil
	}
}

type SubscribeWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeWorkflowInvoker) Invoke() (*model.SubscribeWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeWorkflowResponse), nil
	}
}

type UpdateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowInvoker) Invoke() (*model.UpdateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowResponse), nil
	}
}

type CreateOptimizationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOptimizationTaskInvoker) Invoke() (*model.CreateOptimizationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOptimizationTaskResponse), nil
	}
}

type ShowOptimizationTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOptimizationTaskResultInvoker) Invoke() (*model.ShowOptimizationTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOptimizationTaskResultResponse), nil
	}
}

type CreateSearchTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchTaskInvoker) Invoke() (*model.CreateSearchTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchTaskResponse), nil
	}
}

type ShowSearchTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSearchTaskResultInvoker) Invoke() (*model.ShowSearchTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSearchTaskResultResponse), nil
	}
}
