package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotanalytics/v1/model"
)

type CreateAssetModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssetModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssetModelInvoker) Invoke() (*model.CreateAssetModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssetModelResponse), nil
	}
}

type DeleteAssetModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssetModelInvoker) Invoke() (*model.DeleteAssetModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetModelResponse), nil
	}
}

type ListAssetModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssetModelsInvoker) Invoke() (*model.ListAssetModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetModelsResponse), nil
	}
}

type ShowAssetModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetModelInvoker) Invoke() (*model.ShowAssetModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetModelResponse), nil
	}
}

type UpdateAssetModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssetModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssetModelInvoker) Invoke() (*model.UpdateAssetModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssetModelResponse), nil
	}
}

type CreateAssetNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssetNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssetNewInvoker) Invoke() (*model.CreateAssetNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssetNewResponse), nil
	}
}

type DeleteAssetNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssetNewInvoker) Invoke() (*model.DeleteAssetNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetNewResponse), nil
	}
}

type ListAssetsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetsNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssetsNewInvoker) Invoke() (*model.ListAssetsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetsNewResponse), nil
	}
}

type PublishRootAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishRootAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishRootAssetInvoker) Invoke() (*model.PublishRootAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRootAssetResponse), nil
	}
}

type ShowAssetNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetNewInvoker) Invoke() (*model.ShowAssetNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetNewResponse), nil
	}
}

type UpdateAssetNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssetNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssetNewInvoker) Invoke() (*model.UpdateAssetNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssetNewResponse), nil
	}
}

type ShowLastPropertyValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLastPropertyValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLastPropertyValueInvoker) Invoke() (*model.ShowLastPropertyValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLastPropertyValueResponse), nil
	}
}

type ShowMetricValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetricValueInvoker) Invoke() (*model.ShowMetricValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricValueResponse), nil
	}
}

type ShowPropertyRawValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPropertyRawValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPropertyRawValueInvoker) Invoke() (*model.ShowPropertyRawValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPropertyRawValueResponse), nil
	}
}

type CreateComputingResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComputingResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteComputingResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteComputingResourceInvoker) Invoke() (*model.DeleteComputingResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComputingResourceResponse), nil
	}
}

type ListComputingResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComputingResourcesInvoker) Invoke() (*model.ListComputingResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingResourcesResponse), nil
	}
}

type CreateDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatasourceInvoker) Invoke() (*model.CreateDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatasourceResponse), nil
	}
}

type DeleteDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatasourceInvoker) Invoke() (*model.DeleteDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatasourceResponse), nil
	}
}

type ShowAllDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllDataSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAllDataSourceInvoker) Invoke() (*model.ShowAllDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllDataSourceResponse), nil
	}
}

type ShowDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataSourceInvoker) Invoke() (*model.ShowDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataSourceResponse), nil
	}
}

type UpdateDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataSourceInvoker) Invoke() (*model.UpdateDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataSourceResponse), nil
	}
}

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupInvoker) Invoke() (*model.CreateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupResponse), nil
	}
}

type DeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type ListGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsInvoker) Invoke() (*model.ListGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsResponse), nil
	}
}

type UpdateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupInvoker) Invoke() (*model.UpdateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupResponse), nil
	}
}

type DeleteDataStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataStoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataStoreInvoker) Invoke() (*model.DeleteDataStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataStoreResponse), nil
	}
}

type ListDataStoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataStoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataStoresInvoker) Invoke() (*model.ListDataStoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataStoresResponse), nil
	}
}

type UpdateDataStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataStoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataStoreInvoker) Invoke() (*model.UpdateDataStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataStoreResponse), nil
	}
}

type ListHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryInvoker) Invoke() (*model.ListHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryResponse), nil
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

type ShowPropertyValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPropertyValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPropertyValuesInvoker) Invoke() (*model.ShowPropertyValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPropertyValuesResponse), nil
	}
}

type ListTagValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagValuesInvoker) Invoke() (*model.ListTagValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagValuesResponse), nil
	}
}

type ExportDatasetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDatasetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDatasetInvoker) Invoke() (*model.ExportDatasetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDatasetResponse), nil
	}
}

type ImportDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportDataInvoker) Invoke() (*model.ImportDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDataResponse), nil
	}
}

type ShowDatasetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatasetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDatasetInvoker) Invoke() (*model.ShowDatasetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatasetResponse), nil
	}
}

type ValidateSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateSqlInvoker) Invoke() (*model.ValidateSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateSqlResponse), nil
	}
}

type AddDevDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDevDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDevDataInvoker) Invoke() (*model.AddDevDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDevDataResponse), nil
	}
}

type CreateBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBatchJobInvoker) Invoke() (*model.CreateBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchJobResponse), nil
	}
}

type DeleteBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBatchJobInvoker) Invoke() (*model.DeleteBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBatchJobResponse), nil
	}
}

type ListBatchJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBatchJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBatchJobsInvoker) Invoke() (*model.ListBatchJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBatchJobsResponse), nil
	}
}

type ShowBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBatchJobInvoker) Invoke() (*model.ShowBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchJobResponse), nil
	}
}

type UpdateBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBatchJobInvoker) Invoke() (*model.UpdateBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBatchJobResponse), nil
	}
}

type AddPipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddPipelineJobInvoker) Invoke() (*model.AddPipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPipelineJobResponse), nil
	}
}

type DeletePipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePipelineJobInvoker) Invoke() (*model.DeletePipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipelineJobResponse), nil
	}
}

type ListPipelineJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPipelineJobsInvoker) Invoke() (*model.ListPipelineJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineJobsResponse), nil
	}
}

type ShowPipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPipelineJobInvoker) Invoke() (*model.ShowPipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineJobResponse), nil
	}
}

type StartPipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartPipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartPipelineJobInvoker) Invoke() (*model.StartPipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartPipelineJobResponse), nil
	}
}

type StopPipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopPipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopPipelineJobInvoker) Invoke() (*model.StopPipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopPipelineJobResponse), nil
	}
}

type UpdatePipelineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipelineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePipelineJobInvoker) Invoke() (*model.UpdatePipelineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipelineJobResponse), nil
	}
}

type CreateStreamingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStreamingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStreamingJobInvoker) Invoke() (*model.CreateStreamingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamingJobResponse), nil
	}
}

type DeleteStreamingJobByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStreamingJobByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStreamingJobByIdInvoker) Invoke() (*model.DeleteStreamingJobByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStreamingJobByIdResponse), nil
	}
}

type ShowJobByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobByIdInvoker) Invoke() (*model.ShowJobByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobByIdResponse), nil
	}
}

type ShowJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobsInvoker) Invoke() (*model.ShowJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobsResponse), nil
	}
}

type UpdateStreamingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStreamingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStreamingJobInvoker) Invoke() (*model.UpdateStreamingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStreamingJobResponse), nil
	}
}

type StartJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartJobInvoker) Invoke() (*model.StartJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartJobResponse), nil
	}
}

type StopJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopJobInvoker) Invoke() (*model.StopJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobResponse), nil
	}
}

type CreateRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRunInvoker) Invoke() (*model.CreateRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRunResponse), nil
	}
}

type DeleteRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRunInvoker) Invoke() (*model.DeleteRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRunResponse), nil
	}
}

type ListRunsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRunsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRunsInvoker) Invoke() (*model.ListRunsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRunsResponse), nil
	}
}

type ShowRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRunInvoker) Invoke() (*model.ShowRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRunResponse), nil
	}
}

type CreateTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTableInvoker) Invoke() (*model.CreateTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTableResponse), nil
	}
}

type DeleteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTableInvoker) Invoke() (*model.DeleteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTableResponse), nil
	}
}

type ListTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTablesInvoker) Invoke() (*model.ListTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTablesResponse), nil
	}
}

type ShowTablePreviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTablePreviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTablePreviewInvoker) Invoke() (*model.ShowTablePreviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTablePreviewResponse), nil
	}
}

type ShowTableSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTableSchemaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTableSchemaInvoker) Invoke() (*model.ShowTableSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTableSchemaResponse), nil
	}
}
