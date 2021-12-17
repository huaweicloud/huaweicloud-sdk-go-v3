package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotanalytics/v1/model"
)

type IoTAnalyticsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIoTAnalyticsClient(hcClient *http_client.HcHttpClient) *IoTAnalyticsClient {
	return &IoTAnalyticsClient{HcClient: hcClient}
}

func IoTAnalyticsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建资产模型
func (c *IoTAnalyticsClient) CreateAssetModel(request *model.CreateAssetModelRequest) (*model.CreateAssetModelResponse, error) {
	requestDef := GenReqDefForCreateAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetModelResponse), nil
	}
}

//删除资产模型
func (c *IoTAnalyticsClient) DeleteAssetModel(request *model.DeleteAssetModelRequest) (*model.DeleteAssetModelResponse, error) {
	requestDef := GenReqDefForDeleteAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetModelResponse), nil
	}
}

//获取资产模型列表
func (c *IoTAnalyticsClient) ListAssetModels(request *model.ListAssetModelsRequest) (*model.ListAssetModelsResponse, error) {
	requestDef := GenReqDefForListAssetModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetModelsResponse), nil
	}
}

//获取资产模型详情
func (c *IoTAnalyticsClient) ShowAssetModel(request *model.ShowAssetModelRequest) (*model.ShowAssetModelResponse, error) {
	requestDef := GenReqDefForShowAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetModelResponse), nil
	}
}

//修改资产模型
func (c *IoTAnalyticsClient) UpdateAssetModel(request *model.UpdateAssetModelRequest) (*model.UpdateAssetModelResponse, error) {
	requestDef := GenReqDefForUpdateAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetModelResponse), nil
	}
}

//创建资产
func (c *IoTAnalyticsClient) CreateAssetNew(request *model.CreateAssetNewRequest) (*model.CreateAssetNewResponse, error) {
	requestDef := GenReqDefForCreateAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetNewResponse), nil
	}
}

//删除资产
func (c *IoTAnalyticsClient) DeleteAssetNew(request *model.DeleteAssetNewRequest) (*model.DeleteAssetNewResponse, error) {
	requestDef := GenReqDefForDeleteAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetNewResponse), nil
	}
}

//获取资产列表
func (c *IoTAnalyticsClient) ListAssetsNew(request *model.ListAssetsNewRequest) (*model.ListAssetsNewResponse, error) {
	requestDef := GenReqDefForListAssetsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetsNewResponse), nil
	}
}

//发布资产
func (c *IoTAnalyticsClient) PublishRootAsset(request *model.PublishRootAssetRequest) (*model.PublishRootAssetResponse, error) {
	requestDef := GenReqDefForPublishRootAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishRootAssetResponse), nil
	}
}

//获取资产详情
func (c *IoTAnalyticsClient) ShowAssetNew(request *model.ShowAssetNewRequest) (*model.ShowAssetNewResponse, error) {
	requestDef := GenReqDefForShowAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetNewResponse), nil
	}
}

//修改资产
func (c *IoTAnalyticsClient) UpdateAssetNew(request *model.UpdateAssetNewRequest) (*model.UpdateAssetNewResponse, error) {
	requestDef := GenReqDefForUpdateAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetNewResponse), nil
	}
}

//获取资产属性最新值
func (c *IoTAnalyticsClient) ShowLastPropertyValue(request *model.ShowLastPropertyValueRequest) (*model.ShowLastPropertyValueResponse, error) {
	requestDef := GenReqDefForShowLastPropertyValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastPropertyValueResponse), nil
	}
}

//获取资产属性聚合值
func (c *IoTAnalyticsClient) ShowMetricValue(request *model.ShowMetricValueRequest) (*model.ShowMetricValueResponse, error) {
	requestDef := GenReqDefForShowMetricValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricValueResponse), nil
	}
}

//获取资产属性历史值
func (c *IoTAnalyticsClient) ShowPropertyRawValue(request *model.ShowPropertyRawValueRequest) (*model.ShowPropertyRawValueResponse, error) {
	requestDef := GenReqDefForShowPropertyRawValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyRawValueResponse), nil
	}
}

//创建批计算资源。
func (c *IoTAnalyticsClient) CreateComputingResource(request *model.CreateComputingResourceRequest) (*model.CreateComputingResourceResponse, error) {
	requestDef := GenReqDefForCreateComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComputingResourceResponse), nil
	}
}

//删除批计算资源。
func (c *IoTAnalyticsClient) DeleteComputingResource(request *model.DeleteComputingResourceRequest) (*model.DeleteComputingResourceResponse, error) {
	requestDef := GenReqDefForDeleteComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComputingResourceResponse), nil
	}
}

//查询批计算资源列表。
func (c *IoTAnalyticsClient) ListComputingResources(request *model.ListComputingResourcesRequest) (*model.ListComputingResourcesResponse, error) {
	requestDef := GenReqDefForListComputingResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComputingResourcesResponse), nil
	}
}

//创建数据源
func (c *IoTAnalyticsClient) CreateDatasource(request *model.CreateDatasourceRequest) (*model.CreateDatasourceResponse, error) {
	requestDef := GenReqDefForCreateDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatasourceResponse), nil
	}
}

//删除数据源
func (c *IoTAnalyticsClient) DeleteDatasource(request *model.DeleteDatasourceRequest) (*model.DeleteDatasourceResponse, error) {
	requestDef := GenReqDefForDeleteDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatasourceResponse), nil
	}
}

//查询数据源列表
func (c *IoTAnalyticsClient) ShowAllDataSource(request *model.ShowAllDataSourceRequest) (*model.ShowAllDataSourceResponse, error) {
	requestDef := GenReqDefForShowAllDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllDataSourceResponse), nil
	}
}

//查询数据源详情
func (c *IoTAnalyticsClient) ShowDataSource(request *model.ShowDataSourceRequest) (*model.ShowDataSourceResponse, error) {
	requestDef := GenReqDefForShowDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataSourceResponse), nil
	}
}

//修改数据源
func (c *IoTAnalyticsClient) UpdateDataSource(request *model.UpdateDataSourceRequest) (*model.UpdateDataSourceResponse, error) {
	requestDef := GenReqDefForUpdateDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataSourceResponse), nil
	}
}

//创建存储组
func (c *IoTAnalyticsClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

//删除存储组
func (c *IoTAnalyticsClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

//查询存储组列表
func (c *IoTAnalyticsClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

//更新存储组
func (c *IoTAnalyticsClient) UpdateGroup(request *model.UpdateGroupRequest) (*model.UpdateGroupResponse, error) {
	requestDef := GenReqDefForUpdateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupResponse), nil
	}
}

//删除存储
func (c *IoTAnalyticsClient) DeleteDataStore(request *model.DeleteDataStoreRequest) (*model.DeleteDataStoreResponse, error) {
	requestDef := GenReqDefForDeleteDataStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataStoreResponse), nil
	}
}

//查询存储列表
func (c *IoTAnalyticsClient) ListDataStores(request *model.ListDataStoresRequest) (*model.ListDataStoresResponse, error) {
	requestDef := GenReqDefForListDataStores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataStoresResponse), nil
	}
}

//更新存储
func (c *IoTAnalyticsClient) UpdateDataStore(request *model.UpdateDataStoreRequest) (*model.UpdateDataStoreResponse, error) {
	requestDef := GenReqDefForUpdateDataStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataStoreResponse), nil
	}
}

//根据标签查询设备历史值
func (c *IoTAnalyticsClient) ListHistory(request *model.ListHistoryRequest) (*model.ListHistoryResponse, error) {
	requestDef := GenReqDefForListHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryResponse), nil
	}
}

//根据标签聚合、查询数据
func (c *IoTAnalyticsClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

//查询设备属性最新状态值
func (c *IoTAnalyticsClient) ShowPropertyValues(request *model.ShowPropertyValuesRequest) (*model.ShowPropertyValuesResponse, error) {
	requestDef := GenReqDefForShowPropertyValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyValuesResponse), nil
	}
}

//查询标签的值列表
func (c *IoTAnalyticsClient) ListTagValues(request *model.ListTagValuesRequest) (*model.ListTagValuesResponse, error) {
	requestDef := GenReqDefForListTagValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagValuesResponse), nil
	}
}

//将SQL语句的查询结果下载到本地，只支持下载“QUERY”类型作业的查询结果。
func (c *IoTAnalyticsClient) ExportDataset(request *model.ExportDatasetRequest) (*model.ExportDatasetResponse, error) {
	requestDef := GenReqDefForExportDataset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDatasetResponse), nil
	}
}

//将数据从文件导入OBS表。
func (c *IoTAnalyticsClient) ImportData(request *model.ImportDataRequest) (*model.ImportDataResponse, error) {
	requestDef := GenReqDefForImportData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataResponse), nil
	}
}

//在执行SQL查询语句的作业完成后，查看该作业执行的结果。目前仅支持查看“QUERY”类型作业的执行结果。该API只能查看前1000条的结果记录，若要查看全部的结果记录，需要先导出查询结果再进行查看。
func (c *IoTAnalyticsClient) ShowDataset(request *model.ShowDatasetRequest) (*model.ShowDatasetResponse, error) {
	requestDef := GenReqDefForShowDataset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatasetResponse), nil
	}
}

//检查离线作业SQL语法。
func (c *IoTAnalyticsClient) ValidateSql(request *model.ValidateSqlRequest) (*model.ValidateSqlResponse, error) {
	requestDef := GenReqDefForValidateSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateSqlResponse), nil
	}
}

//通过API数据源上报设备数据
func (c *IoTAnalyticsClient) AddDevData(request *model.AddDevDataRequest) (*model.AddDevDataResponse, error) {
	requestDef := GenReqDefForAddDevData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDevDataResponse), nil
	}
}

//创建离线作业。
func (c *IoTAnalyticsClient) CreateBatchJob(request *model.CreateBatchJobRequest) (*model.CreateBatchJobResponse, error) {
	requestDef := GenReqDefForCreateBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchJobResponse), nil
	}
}

//删除离线作业。
func (c *IoTAnalyticsClient) DeleteBatchJob(request *model.DeleteBatchJobRequest) (*model.DeleteBatchJobResponse, error) {
	requestDef := GenReqDefForDeleteBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBatchJobResponse), nil
	}
}

//查询离线作业列表。
func (c *IoTAnalyticsClient) ListBatchJobs(request *model.ListBatchJobsRequest) (*model.ListBatchJobsResponse, error) {
	requestDef := GenReqDefForListBatchJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchJobsResponse), nil
	}
}

//查询离线作业详情。
func (c *IoTAnalyticsClient) ShowBatchJob(request *model.ShowBatchJobRequest) (*model.ShowBatchJobResponse, error) {
	requestDef := GenReqDefForShowBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchJobResponse), nil
	}
}

//修改离线作业。
func (c *IoTAnalyticsClient) UpdateBatchJob(request *model.UpdateBatchJobRequest) (*model.UpdateBatchJobResponse, error) {
	requestDef := GenReqDefForUpdateBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBatchJobResponse), nil
	}
}

//新建管道作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。（作业中各算子的详细配置请参考算子配置章节。） check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查。当检查不通过时，将作业状态修改为草稿；检查通过时，将作业状态修改为就绪，并返回成功。
func (c *IoTAnalyticsClient) AddPipelineJob(request *model.AddPipelineJobRequest) (*model.AddPipelineJobResponse, error) {
	requestDef := GenReqDefForAddPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPipelineJobResponse), nil
	}
}

//删除用户指定的管道作业
func (c *IoTAnalyticsClient) DeletePipelineJob(request *model.DeletePipelineJobRequest) (*model.DeletePipelineJobResponse, error) {
	requestDef := GenReqDefForDeletePipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipelineJobResponse), nil
	}
}

//获取用户下的所有管道作业，支持分页。
func (c *IoTAnalyticsClient) ListPipelineJobs(request *model.ListPipelineJobsRequest) (*model.ListPipelineJobsResponse, error) {
	requestDef := GenReqDefForListPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineJobsResponse), nil
	}
}

//获取指定管道作业的详情
func (c *IoTAnalyticsClient) ShowPipelineJob(request *model.ShowPipelineJobRequest) (*model.ShowPipelineJobResponse, error) {
	requestDef := GenReqDefForShowPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineJobResponse), nil
	}
}

//提交管道作业到运行环境，实时接收数据源的数据并按用户定义的数据清洗逻辑对数据进行处理。
func (c *IoTAnalyticsClient) StartPipelineJob(request *model.StartPipelineJobRequest) (*model.StartPipelineJobResponse, error) {
	requestDef := GenReqDefForStartPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineJobResponse), nil
	}
}

//停止一个正在运行中的管道作业
func (c *IoTAnalyticsClient) StopPipelineJob(request *model.StopPipelineJobRequest) (*model.StopPipelineJobResponse, error) {
	requestDef := GenReqDefForStopPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineJobResponse), nil
	}
}

//更新管道作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。（管道作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：数据管道算子配置指南。） check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查。当检查不通过时，将作业状态修改为草稿；检查通过时，将作业状态修改为就绪，并返回成功。
func (c *IoTAnalyticsClient) UpdatePipelineJob(request *model.UpdatePipelineJobRequest) (*model.UpdatePipelineJobResponse, error) {
	requestDef := GenReqDefForUpdatePipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipelineJobResponse), nil
	}
}

//除名称和描述外，可先不提供作业的详细配置信息。 check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查，无论检查是否通过，都将作业及配置信息保存为草稿，当检查不通过时，返回失败及错误信息，检查通过时，将作业状态修改为就绪，并返回成功。
func (c *IoTAnalyticsClient) CreateStreamingJob(request *model.CreateStreamingJobRequest) (*model.CreateStreamingJobResponse, error) {
	requestDef := GenReqDefForCreateStreamingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamingJobResponse), nil
	}
}

//删除用户指定的作业
func (c *IoTAnalyticsClient) DeleteStreamingJobById(request *model.DeleteStreamingJobByIdRequest) (*model.DeleteStreamingJobByIdResponse, error) {
	requestDef := GenReqDefForDeleteStreamingJobById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStreamingJobByIdResponse), nil
	}
}

//获取指定作业的详情
func (c *IoTAnalyticsClient) ShowJobById(request *model.ShowJobByIdRequest) (*model.ShowJobByIdResponse, error) {
	requestDef := GenReqDefForShowJobById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobByIdResponse), nil
	}
}

//获取用户下的所有实时分析作业，支持分页。
func (c *IoTAnalyticsClient) ShowJobs(request *model.ShowJobsRequest) (*model.ShowJobsResponse, error) {
	requestDef := GenReqDefForShowJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobsResponse), nil
	}
}

//更新作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。 check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查，无论检查是否通过，都将作业及配置信息保存为草稿，当检查不通过时，返回失败及错误信息，检查通过时，将作业状态修改为就绪，并返回成功。
func (c *IoTAnalyticsClient) UpdateStreamingJob(request *model.UpdateStreamingJobRequest) (*model.UpdateStreamingJobResponse, error) {
	requestDef := GenReqDefForUpdateStreamingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStreamingJobResponse), nil
	}
}

//提交作业到运行环境，实时接收数据并按用户定义的业务逻辑对数据进行处理。
func (c *IoTAnalyticsClient) StartJob(request *model.StartJobRequest) (*model.StartJobResponse, error) {
	requestDef := GenReqDefForStartJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartJobResponse), nil
	}
}

//停止一个正在运行中的作业
func (c *IoTAnalyticsClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

//执行离线作业。
func (c *IoTAnalyticsClient) CreateRun(request *model.CreateRunRequest) (*model.CreateRunResponse, error) {
	requestDef := GenReqDefForCreateRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRunResponse), nil
	}
}

//停止提交中或运行中的离线作业，若作业已经执行结束或失败则无法停止。
func (c *IoTAnalyticsClient) DeleteRun(request *model.DeleteRunRequest) (*model.DeleteRunResponse, error) {
	requestDef := GenReqDefForDeleteRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRunResponse), nil
	}
}

//查询离线作业运行列表。
func (c *IoTAnalyticsClient) ListRuns(request *model.ListRunsRequest) (*model.ListRunsResponse, error) {
	requestDef := GenReqDefForListRuns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRunsResponse), nil
	}
}

//查询离线作业运行详情。
func (c *IoTAnalyticsClient) ShowRun(request *model.ShowRunRequest) (*model.ShowRunResponse, error) {
	requestDef := GenReqDefForShowRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRunResponse), nil
	}
}

//创建离线数据表。
func (c *IoTAnalyticsClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

//删除离线数据表。
func (c *IoTAnalyticsClient) DeleteTable(request *model.DeleteTableRequest) (*model.DeleteTableResponse, error) {
	requestDef := GenReqDefForDeleteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableResponse), nil
	}
}

//查询离线数据表列表。
func (c *IoTAnalyticsClient) ListTables(request *model.ListTablesRequest) (*model.ListTablesResponse, error) {
	requestDef := GenReqDefForListTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesResponse), nil
	}
}

//预览离线数据表内容。
func (c *IoTAnalyticsClient) ShowTablePreview(request *model.ShowTablePreviewRequest) (*model.ShowTablePreviewResponse, error) {
	requestDef := GenReqDefForShowTablePreview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTablePreviewResponse), nil
	}
}

//查询离线数据表结构。
func (c *IoTAnalyticsClient) ShowTableSchema(request *model.ShowTableSchemaRequest) (*model.ShowTableSchemaResponse, error) {
	requestDef := GenReqDefForShowTableSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableSchemaResponse), nil
	}
}
