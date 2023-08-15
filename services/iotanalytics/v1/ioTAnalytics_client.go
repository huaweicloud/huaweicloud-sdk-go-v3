package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/iotanalytics/v1/model"
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

// CreateAssetModel 创建资产模型
//
// 创建资产模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateAssetModel(request *model.CreateAssetModelRequest) (*model.CreateAssetModelResponse, error) {
	requestDef := GenReqDefForCreateAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetModelResponse), nil
	}
}

// CreateAssetModelInvoker 创建资产模型
func (c *IoTAnalyticsClient) CreateAssetModelInvoker(request *model.CreateAssetModelRequest) *CreateAssetModelInvoker {
	requestDef := GenReqDefForCreateAssetModel()
	return &CreateAssetModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAssetModel 删除资产模型
//
// 删除资产模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteAssetModel(request *model.DeleteAssetModelRequest) (*model.DeleteAssetModelResponse, error) {
	requestDef := GenReqDefForDeleteAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetModelResponse), nil
	}
}

// DeleteAssetModelInvoker 删除资产模型
func (c *IoTAnalyticsClient) DeleteAssetModelInvoker(request *model.DeleteAssetModelRequest) *DeleteAssetModelInvoker {
	requestDef := GenReqDefForDeleteAssetModel()
	return &DeleteAssetModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssetModels 获取资产模型列表
//
// 获取资产模型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListAssetModels(request *model.ListAssetModelsRequest) (*model.ListAssetModelsResponse, error) {
	requestDef := GenReqDefForListAssetModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetModelsResponse), nil
	}
}

// ListAssetModelsInvoker 获取资产模型列表
func (c *IoTAnalyticsClient) ListAssetModelsInvoker(request *model.ListAssetModelsRequest) *ListAssetModelsInvoker {
	requestDef := GenReqDefForListAssetModels()
	return &ListAssetModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetModel 获取资产模型详情
//
// 获取资产模型详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowAssetModel(request *model.ShowAssetModelRequest) (*model.ShowAssetModelResponse, error) {
	requestDef := GenReqDefForShowAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetModelResponse), nil
	}
}

// ShowAssetModelInvoker 获取资产模型详情
func (c *IoTAnalyticsClient) ShowAssetModelInvoker(request *model.ShowAssetModelRequest) *ShowAssetModelInvoker {
	requestDef := GenReqDefForShowAssetModel()
	return &ShowAssetModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssetModel 修改资产模型
//
// 修改资产模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateAssetModel(request *model.UpdateAssetModelRequest) (*model.UpdateAssetModelResponse, error) {
	requestDef := GenReqDefForUpdateAssetModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetModelResponse), nil
	}
}

// UpdateAssetModelInvoker 修改资产模型
func (c *IoTAnalyticsClient) UpdateAssetModelInvoker(request *model.UpdateAssetModelRequest) *UpdateAssetModelInvoker {
	requestDef := GenReqDefForUpdateAssetModel()
	return &UpdateAssetModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAssetNew 创建资产
//
// 创建资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateAssetNew(request *model.CreateAssetNewRequest) (*model.CreateAssetNewResponse, error) {
	requestDef := GenReqDefForCreateAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetNewResponse), nil
	}
}

// CreateAssetNewInvoker 创建资产
func (c *IoTAnalyticsClient) CreateAssetNewInvoker(request *model.CreateAssetNewRequest) *CreateAssetNewInvoker {
	requestDef := GenReqDefForCreateAssetNew()
	return &CreateAssetNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAssetNew 删除资产
//
// 删除资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteAssetNew(request *model.DeleteAssetNewRequest) (*model.DeleteAssetNewResponse, error) {
	requestDef := GenReqDefForDeleteAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetNewResponse), nil
	}
}

// DeleteAssetNewInvoker 删除资产
func (c *IoTAnalyticsClient) DeleteAssetNewInvoker(request *model.DeleteAssetNewRequest) *DeleteAssetNewInvoker {
	requestDef := GenReqDefForDeleteAssetNew()
	return &DeleteAssetNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssetsNew 获取资产列表
//
// 获取资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListAssetsNew(request *model.ListAssetsNewRequest) (*model.ListAssetsNewResponse, error) {
	requestDef := GenReqDefForListAssetsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetsNewResponse), nil
	}
}

// ListAssetsNewInvoker 获取资产列表
func (c *IoTAnalyticsClient) ListAssetsNewInvoker(request *model.ListAssetsNewRequest) *ListAssetsNewInvoker {
	requestDef := GenReqDefForListAssetsNew()
	return &ListAssetsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishRootAsset 发布资产
//
// 发布资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) PublishRootAsset(request *model.PublishRootAssetRequest) (*model.PublishRootAssetResponse, error) {
	requestDef := GenReqDefForPublishRootAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishRootAssetResponse), nil
	}
}

// PublishRootAssetInvoker 发布资产
func (c *IoTAnalyticsClient) PublishRootAssetInvoker(request *model.PublishRootAssetRequest) *PublishRootAssetInvoker {
	requestDef := GenReqDefForPublishRootAsset()
	return &PublishRootAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetNew 获取资产详情
//
// 获取资产详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowAssetNew(request *model.ShowAssetNewRequest) (*model.ShowAssetNewResponse, error) {
	requestDef := GenReqDefForShowAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetNewResponse), nil
	}
}

// ShowAssetNewInvoker 获取资产详情
func (c *IoTAnalyticsClient) ShowAssetNewInvoker(request *model.ShowAssetNewRequest) *ShowAssetNewInvoker {
	requestDef := GenReqDefForShowAssetNew()
	return &ShowAssetNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssetNew 修改资产
//
// 修改资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateAssetNew(request *model.UpdateAssetNewRequest) (*model.UpdateAssetNewResponse, error) {
	requestDef := GenReqDefForUpdateAssetNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetNewResponse), nil
	}
}

// UpdateAssetNewInvoker 修改资产
func (c *IoTAnalyticsClient) UpdateAssetNewInvoker(request *model.UpdateAssetNewRequest) *UpdateAssetNewInvoker {
	requestDef := GenReqDefForUpdateAssetNew()
	return &UpdateAssetNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLastPropertyValue 获取资产属性最新值
//
// 获取资产属性最新值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowLastPropertyValue(request *model.ShowLastPropertyValueRequest) (*model.ShowLastPropertyValueResponse, error) {
	requestDef := GenReqDefForShowLastPropertyValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastPropertyValueResponse), nil
	}
}

// ShowLastPropertyValueInvoker 获取资产属性最新值
func (c *IoTAnalyticsClient) ShowLastPropertyValueInvoker(request *model.ShowLastPropertyValueRequest) *ShowLastPropertyValueInvoker {
	requestDef := GenReqDefForShowLastPropertyValue()
	return &ShowLastPropertyValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricValue 获取资产属性聚合值
//
// 获取资产属性聚合值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowMetricValue(request *model.ShowMetricValueRequest) (*model.ShowMetricValueResponse, error) {
	requestDef := GenReqDefForShowMetricValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricValueResponse), nil
	}
}

// ShowMetricValueInvoker 获取资产属性聚合值
func (c *IoTAnalyticsClient) ShowMetricValueInvoker(request *model.ShowMetricValueRequest) *ShowMetricValueInvoker {
	requestDef := GenReqDefForShowMetricValue()
	return &ShowMetricValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPropertyRawValue 获取资产属性历史值
//
// 获取资产属性历史值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowPropertyRawValue(request *model.ShowPropertyRawValueRequest) (*model.ShowPropertyRawValueResponse, error) {
	requestDef := GenReqDefForShowPropertyRawValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyRawValueResponse), nil
	}
}

// ShowPropertyRawValueInvoker 获取资产属性历史值
func (c *IoTAnalyticsClient) ShowPropertyRawValueInvoker(request *model.ShowPropertyRawValueRequest) *ShowPropertyRawValueInvoker {
	requestDef := GenReqDefForShowPropertyRawValue()
	return &ShowPropertyRawValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComputingResource 创建批计算资源
//
// 创建批计算资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateComputingResource(request *model.CreateComputingResourceRequest) (*model.CreateComputingResourceResponse, error) {
	requestDef := GenReqDefForCreateComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComputingResourceResponse), nil
	}
}

// CreateComputingResourceInvoker 创建批计算资源
func (c *IoTAnalyticsClient) CreateComputingResourceInvoker(request *model.CreateComputingResourceRequest) *CreateComputingResourceInvoker {
	requestDef := GenReqDefForCreateComputingResource()
	return &CreateComputingResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComputingResource 删除批计算资源
//
// 删除批计算资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteComputingResource(request *model.DeleteComputingResourceRequest) (*model.DeleteComputingResourceResponse, error) {
	requestDef := GenReqDefForDeleteComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComputingResourceResponse), nil
	}
}

// DeleteComputingResourceInvoker 删除批计算资源
func (c *IoTAnalyticsClient) DeleteComputingResourceInvoker(request *model.DeleteComputingResourceRequest) *DeleteComputingResourceInvoker {
	requestDef := GenReqDefForDeleteComputingResource()
	return &DeleteComputingResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComputingResources 查询批计算资源列表
//
// 查询批计算资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListComputingResources(request *model.ListComputingResourcesRequest) (*model.ListComputingResourcesResponse, error) {
	requestDef := GenReqDefForListComputingResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComputingResourcesResponse), nil
	}
}

// ListComputingResourcesInvoker 查询批计算资源列表
func (c *IoTAnalyticsClient) ListComputingResourcesInvoker(request *model.ListComputingResourcesRequest) *ListComputingResourcesInvoker {
	requestDef := GenReqDefForListComputingResources()
	return &ListComputingResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatasource 创建数据源
//
// 创建数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateDatasource(request *model.CreateDatasourceRequest) (*model.CreateDatasourceResponse, error) {
	requestDef := GenReqDefForCreateDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatasourceResponse), nil
	}
}

// CreateDatasourceInvoker 创建数据源
func (c *IoTAnalyticsClient) CreateDatasourceInvoker(request *model.CreateDatasourceRequest) *CreateDatasourceInvoker {
	requestDef := GenReqDefForCreateDatasource()
	return &CreateDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatasource 删除数据源
//
// 删除数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteDatasource(request *model.DeleteDatasourceRequest) (*model.DeleteDatasourceResponse, error) {
	requestDef := GenReqDefForDeleteDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatasourceResponse), nil
	}
}

// DeleteDatasourceInvoker 删除数据源
func (c *IoTAnalyticsClient) DeleteDatasourceInvoker(request *model.DeleteDatasourceRequest) *DeleteDatasourceInvoker {
	requestDef := GenReqDefForDeleteDatasource()
	return &DeleteDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllDataSource 查询数据源列表
//
// 查询数据源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowAllDataSource(request *model.ShowAllDataSourceRequest) (*model.ShowAllDataSourceResponse, error) {
	requestDef := GenReqDefForShowAllDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllDataSourceResponse), nil
	}
}

// ShowAllDataSourceInvoker 查询数据源列表
func (c *IoTAnalyticsClient) ShowAllDataSourceInvoker(request *model.ShowAllDataSourceRequest) *ShowAllDataSourceInvoker {
	requestDef := GenReqDefForShowAllDataSource()
	return &ShowAllDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataSource 查询数据源详情
//
// 查询数据源详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowDataSource(request *model.ShowDataSourceRequest) (*model.ShowDataSourceResponse, error) {
	requestDef := GenReqDefForShowDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataSourceResponse), nil
	}
}

// ShowDataSourceInvoker 查询数据源详情
func (c *IoTAnalyticsClient) ShowDataSourceInvoker(request *model.ShowDataSourceRequest) *ShowDataSourceInvoker {
	requestDef := GenReqDefForShowDataSource()
	return &ShowDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataSource 修改数据源
//
// 修改数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateDataSource(request *model.UpdateDataSourceRequest) (*model.UpdateDataSourceResponse, error) {
	requestDef := GenReqDefForUpdateDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataSourceResponse), nil
	}
}

// UpdateDataSourceInvoker 修改数据源
func (c *IoTAnalyticsClient) UpdateDataSourceInvoker(request *model.UpdateDataSourceRequest) *UpdateDataSourceInvoker {
	requestDef := GenReqDefForUpdateDataSource()
	return &UpdateDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroup 创建存储组
//
// 创建存储组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建存储组
func (c *IoTAnalyticsClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除存储组
//
// 删除存储组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除存储组
func (c *IoTAnalyticsClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 查询存储组列表
//
// 查询存储组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 查询存储组列表
func (c *IoTAnalyticsClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroup 更新存储组
//
// 更新存储组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateGroup(request *model.UpdateGroupRequest) (*model.UpdateGroupResponse, error) {
	requestDef := GenReqDefForUpdateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupResponse), nil
	}
}

// UpdateGroupInvoker 更新存储组
func (c *IoTAnalyticsClient) UpdateGroupInvoker(request *model.UpdateGroupRequest) *UpdateGroupInvoker {
	requestDef := GenReqDefForUpdateGroup()
	return &UpdateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataStore 删除存储
//
// 删除存储
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteDataStore(request *model.DeleteDataStoreRequest) (*model.DeleteDataStoreResponse, error) {
	requestDef := GenReqDefForDeleteDataStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataStoreResponse), nil
	}
}

// DeleteDataStoreInvoker 删除存储
func (c *IoTAnalyticsClient) DeleteDataStoreInvoker(request *model.DeleteDataStoreRequest) *DeleteDataStoreInvoker {
	requestDef := GenReqDefForDeleteDataStore()
	return &DeleteDataStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataStores 查询存储列表
//
// 查询存储列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListDataStores(request *model.ListDataStoresRequest) (*model.ListDataStoresResponse, error) {
	requestDef := GenReqDefForListDataStores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataStoresResponse), nil
	}
}

// ListDataStoresInvoker 查询存储列表
func (c *IoTAnalyticsClient) ListDataStoresInvoker(request *model.ListDataStoresRequest) *ListDataStoresInvoker {
	requestDef := GenReqDefForListDataStores()
	return &ListDataStoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataStore 更新存储
//
// 更新存储
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateDataStore(request *model.UpdateDataStoreRequest) (*model.UpdateDataStoreResponse, error) {
	requestDef := GenReqDefForUpdateDataStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataStoreResponse), nil
	}
}

// UpdateDataStoreInvoker 更新存储
func (c *IoTAnalyticsClient) UpdateDataStoreInvoker(request *model.UpdateDataStoreRequest) *UpdateDataStoreInvoker {
	requestDef := GenReqDefForUpdateDataStore()
	return &UpdateDataStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistory 根据标签查询设备历史值
//
// 根据标签查询设备历史值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListHistory(request *model.ListHistoryRequest) (*model.ListHistoryResponse, error) {
	requestDef := GenReqDefForListHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryResponse), nil
	}
}

// ListHistoryInvoker 根据标签查询设备历史值
func (c *IoTAnalyticsClient) ListHistoryInvoker(request *model.ListHistoryRequest) *ListHistoryInvoker {
	requestDef := GenReqDefForListHistory()
	return &ListHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 根据标签聚合、查询指标数据
//
// 根据标签聚合、查询数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 根据标签聚合、查询指标数据
func (c *IoTAnalyticsClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPropertyValues 查询设备属性最新状态值
//
// 查询设备属性最新状态值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowPropertyValues(request *model.ShowPropertyValuesRequest) (*model.ShowPropertyValuesResponse, error) {
	requestDef := GenReqDefForShowPropertyValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyValuesResponse), nil
	}
}

// ShowPropertyValuesInvoker 查询设备属性最新状态值
func (c *IoTAnalyticsClient) ShowPropertyValuesInvoker(request *model.ShowPropertyValuesRequest) *ShowPropertyValuesInvoker {
	requestDef := GenReqDefForShowPropertyValues()
	return &ShowPropertyValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagValues 查询标签的值列表
//
// 查询标签的值列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListTagValues(request *model.ListTagValuesRequest) (*model.ListTagValuesResponse, error) {
	requestDef := GenReqDefForListTagValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagValuesResponse), nil
	}
}

// ListTagValuesInvoker 查询标签的值列表
func (c *IoTAnalyticsClient) ListTagValuesInvoker(request *model.ListTagValuesRequest) *ListTagValuesInvoker {
	requestDef := GenReqDefForListTagValues()
	return &ListTagValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDataset 下载离线作业结果
//
// 将SQL语句的查询结果下载到本地，只支持下载“QUERY”类型作业的查询结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ExportDataset(request *model.ExportDatasetRequest) (*model.ExportDatasetResponse, error) {
	requestDef := GenReqDefForExportDataset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDatasetResponse), nil
	}
}

// ExportDatasetInvoker 下载离线作业结果
func (c *IoTAnalyticsClient) ExportDatasetInvoker(request *model.ExportDatasetRequest) *ExportDatasetInvoker {
	requestDef := GenReqDefForExportDataset()
	return &ExportDatasetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportData 执行导入数据离线作业
//
// 将数据从文件导入OBS表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ImportData(request *model.ImportDataRequest) (*model.ImportDataResponse, error) {
	requestDef := GenReqDefForImportData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataResponse), nil
	}
}

// ImportDataInvoker 执行导入数据离线作业
func (c *IoTAnalyticsClient) ImportDataInvoker(request *model.ImportDataRequest) *ImportDataInvoker {
	requestDef := GenReqDefForImportData()
	return &ImportDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataset 查询离线作业结果
//
// 在执行SQL查询语句的作业完成后，查看该作业执行的结果。目前仅支持查看“QUERY”类型作业的执行结果。该API只能查看前1000条的结果记录，若要查看全部的结果记录，需要先导出查询结果再进行查看。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowDataset(request *model.ShowDatasetRequest) (*model.ShowDatasetResponse, error) {
	requestDef := GenReqDefForShowDataset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatasetResponse), nil
	}
}

// ShowDatasetInvoker 查询离线作业结果
func (c *IoTAnalyticsClient) ShowDatasetInvoker(request *model.ShowDatasetRequest) *ShowDatasetInvoker {
	requestDef := GenReqDefForShowDataset()
	return &ShowDatasetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateSql 检查离线作业SQL语法
//
// 检查离线作业SQL语法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ValidateSql(request *model.ValidateSqlRequest) (*model.ValidateSqlResponse, error) {
	requestDef := GenReqDefForValidateSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateSqlResponse), nil
	}
}

// ValidateSqlInvoker 检查离线作业SQL语法
func (c *IoTAnalyticsClient) ValidateSqlInvoker(request *model.ValidateSqlRequest) *ValidateSqlInvoker {
	requestDef := GenReqDefForValidateSql()
	return &ValidateSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDevData 通过API数据源上报设备数据
//
// 通过API数据源上报设备数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) AddDevData(request *model.AddDevDataRequest) (*model.AddDevDataResponse, error) {
	requestDef := GenReqDefForAddDevData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDevDataResponse), nil
	}
}

// AddDevDataInvoker 通过API数据源上报设备数据
func (c *IoTAnalyticsClient) AddDevDataInvoker(request *model.AddDevDataRequest) *AddDevDataInvoker {
	requestDef := GenReqDefForAddDevData()
	return &AddDevDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchJob 创建离线作业
//
// 创建离线作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateBatchJob(request *model.CreateBatchJobRequest) (*model.CreateBatchJobResponse, error) {
	requestDef := GenReqDefForCreateBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchJobResponse), nil
	}
}

// CreateBatchJobInvoker 创建离线作业
func (c *IoTAnalyticsClient) CreateBatchJobInvoker(request *model.CreateBatchJobRequest) *CreateBatchJobInvoker {
	requestDef := GenReqDefForCreateBatchJob()
	return &CreateBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBatchJob 删除离线作业
//
// 删除离线作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteBatchJob(request *model.DeleteBatchJobRequest) (*model.DeleteBatchJobResponse, error) {
	requestDef := GenReqDefForDeleteBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBatchJobResponse), nil
	}
}

// DeleteBatchJobInvoker 删除离线作业
func (c *IoTAnalyticsClient) DeleteBatchJobInvoker(request *model.DeleteBatchJobRequest) *DeleteBatchJobInvoker {
	requestDef := GenReqDefForDeleteBatchJob()
	return &DeleteBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBatchJobs 查询离线作业列表
//
// 查询离线作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListBatchJobs(request *model.ListBatchJobsRequest) (*model.ListBatchJobsResponse, error) {
	requestDef := GenReqDefForListBatchJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchJobsResponse), nil
	}
}

// ListBatchJobsInvoker 查询离线作业列表
func (c *IoTAnalyticsClient) ListBatchJobsInvoker(request *model.ListBatchJobsRequest) *ListBatchJobsInvoker {
	requestDef := GenReqDefForListBatchJobs()
	return &ListBatchJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchJob 查询离线作业详情
//
// 查询离线作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowBatchJob(request *model.ShowBatchJobRequest) (*model.ShowBatchJobResponse, error) {
	requestDef := GenReqDefForShowBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchJobResponse), nil
	}
}

// ShowBatchJobInvoker 查询离线作业详情
func (c *IoTAnalyticsClient) ShowBatchJobInvoker(request *model.ShowBatchJobRequest) *ShowBatchJobInvoker {
	requestDef := GenReqDefForShowBatchJob()
	return &ShowBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBatchJob 修改离线作业
//
// 修改离线作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateBatchJob(request *model.UpdateBatchJobRequest) (*model.UpdateBatchJobResponse, error) {
	requestDef := GenReqDefForUpdateBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBatchJobResponse), nil
	}
}

// UpdateBatchJobInvoker 修改离线作业
func (c *IoTAnalyticsClient) UpdateBatchJobInvoker(request *model.UpdateBatchJobRequest) *UpdateBatchJobInvoker {
	requestDef := GenReqDefForUpdateBatchJob()
	return &UpdateBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPipelineJob 新建管道作业
//
// 新建管道作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。（作业中各算子的详细配置请参考算子配置章节。） check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查。当检查不通过时，将作业状态修改为草稿；检查通过时，将作业状态修改为就绪，并返回成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) AddPipelineJob(request *model.AddPipelineJobRequest) (*model.AddPipelineJobResponse, error) {
	requestDef := GenReqDefForAddPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPipelineJobResponse), nil
	}
}

// AddPipelineJobInvoker 新建管道作业
func (c *IoTAnalyticsClient) AddPipelineJobInvoker(request *model.AddPipelineJobRequest) *AddPipelineJobInvoker {
	requestDef := GenReqDefForAddPipelineJob()
	return &AddPipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipelineJob 删除管道作业
//
// 删除用户指定的管道作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeletePipelineJob(request *model.DeletePipelineJobRequest) (*model.DeletePipelineJobResponse, error) {
	requestDef := GenReqDefForDeletePipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipelineJobResponse), nil
	}
}

// DeletePipelineJobInvoker 删除管道作业
func (c *IoTAnalyticsClient) DeletePipelineJobInvoker(request *model.DeletePipelineJobRequest) *DeletePipelineJobInvoker {
	requestDef := GenReqDefForDeletePipelineJob()
	return &DeletePipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineJobs 获取管道作业列表
//
// 获取用户下的所有管道作业，支持分页。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListPipelineJobs(request *model.ListPipelineJobsRequest) (*model.ListPipelineJobsResponse, error) {
	requestDef := GenReqDefForListPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineJobsResponse), nil
	}
}

// ListPipelineJobsInvoker 获取管道作业列表
func (c *IoTAnalyticsClient) ListPipelineJobsInvoker(request *model.ListPipelineJobsRequest) *ListPipelineJobsInvoker {
	requestDef := GenReqDefForListPipelineJobs()
	return &ListPipelineJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineJob 获取管道作业详情
//
// 获取指定管道作业的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowPipelineJob(request *model.ShowPipelineJobRequest) (*model.ShowPipelineJobResponse, error) {
	requestDef := GenReqDefForShowPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineJobResponse), nil
	}
}

// ShowPipelineJobInvoker 获取管道作业详情
func (c *IoTAnalyticsClient) ShowPipelineJobInvoker(request *model.ShowPipelineJobRequest) *ShowPipelineJobInvoker {
	requestDef := GenReqDefForShowPipelineJob()
	return &ShowPipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartPipelineJob 启动管道作业
//
// 提交管道作业到运行环境，实时接收数据源的数据并按用户定义的数据清洗逻辑对数据进行处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) StartPipelineJob(request *model.StartPipelineJobRequest) (*model.StartPipelineJobResponse, error) {
	requestDef := GenReqDefForStartPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineJobResponse), nil
	}
}

// StartPipelineJobInvoker 启动管道作业
func (c *IoTAnalyticsClient) StartPipelineJobInvoker(request *model.StartPipelineJobRequest) *StartPipelineJobInvoker {
	requestDef := GenReqDefForStartPipelineJob()
	return &StartPipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopPipelineJob 停止管道作业
//
// 停止一个正在运行中的管道作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) StopPipelineJob(request *model.StopPipelineJobRequest) (*model.StopPipelineJobResponse, error) {
	requestDef := GenReqDefForStopPipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineJobResponse), nil
	}
}

// StopPipelineJobInvoker 停止管道作业
func (c *IoTAnalyticsClient) StopPipelineJobInvoker(request *model.StopPipelineJobRequest) *StopPipelineJobInvoker {
	requestDef := GenReqDefForStopPipelineJob()
	return &StopPipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipelineJob 更新管道作业
//
// 更新管道作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。（管道作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：数据管道算子配置指南。） check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查。当检查不通过时，将作业状态修改为草稿；检查通过时，将作业状态修改为就绪，并返回成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdatePipelineJob(request *model.UpdatePipelineJobRequest) (*model.UpdatePipelineJobResponse, error) {
	requestDef := GenReqDefForUpdatePipelineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipelineJobResponse), nil
	}
}

// UpdatePipelineJobInvoker 更新管道作业
func (c *IoTAnalyticsClient) UpdatePipelineJobInvoker(request *model.UpdatePipelineJobRequest) *UpdatePipelineJobInvoker {
	requestDef := GenReqDefForUpdatePipelineJob()
	return &UpdatePipelineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStreamingJob 新建实时作业
//
// 除名称和描述外，可先不提供作业的详细配置信息。 check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查，无论检查是否通过，都将作业及配置信息保存为草稿，当检查不通过时，返回失败及错误信息，检查通过时，将作业状态修改为就绪，并返回成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateStreamingJob(request *model.CreateStreamingJobRequest) (*model.CreateStreamingJobResponse, error) {
	requestDef := GenReqDefForCreateStreamingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamingJobResponse), nil
	}
}

// CreateStreamingJobInvoker 新建实时作业
func (c *IoTAnalyticsClient) CreateStreamingJobInvoker(request *model.CreateStreamingJobRequest) *CreateStreamingJobInvoker {
	requestDef := GenReqDefForCreateStreamingJob()
	return &CreateStreamingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStreamingJobById 删除实时作业
//
// 删除用户指定的作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteStreamingJobById(request *model.DeleteStreamingJobByIdRequest) (*model.DeleteStreamingJobByIdResponse, error) {
	requestDef := GenReqDefForDeleteStreamingJobById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStreamingJobByIdResponse), nil
	}
}

// DeleteStreamingJobByIdInvoker 删除实时作业
func (c *IoTAnalyticsClient) DeleteStreamingJobByIdInvoker(request *model.DeleteStreamingJobByIdRequest) *DeleteStreamingJobByIdInvoker {
	requestDef := GenReqDefForDeleteStreamingJobById()
	return &DeleteStreamingJobByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobById 获取实时作业详情
//
// 获取指定作业的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowJobById(request *model.ShowJobByIdRequest) (*model.ShowJobByIdResponse, error) {
	requestDef := GenReqDefForShowJobById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobByIdResponse), nil
	}
}

// ShowJobByIdInvoker 获取实时作业详情
func (c *IoTAnalyticsClient) ShowJobByIdInvoker(request *model.ShowJobByIdRequest) *ShowJobByIdInvoker {
	requestDef := GenReqDefForShowJobById()
	return &ShowJobByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobs 获取实时作业列表
//
// 获取用户下的所有实时分析作业，支持分页。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowJobs(request *model.ShowJobsRequest) (*model.ShowJobsResponse, error) {
	requestDef := GenReqDefForShowJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobsResponse), nil
	}
}

// ShowJobsInvoker 获取实时作业列表
func (c *IoTAnalyticsClient) ShowJobsInvoker(request *model.ShowJobsRequest) *ShowJobsInvoker {
	requestDef := GenReqDefForShowJobs()
	return &ShowJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStreamingJob 更新实时作业
//
// 更新作业时，需要在URL中指定是更新哪一个作业，将在body中附带完整的作业信息。 check参数表示是否需要对作业配置进行检查，若为false，则不检查，将作业保存为草稿；若为true，则对作业配置进行检查，无论检查是否通过，都将作业及配置信息保存为草稿，当检查不通过时，返回失败及错误信息，检查通过时，将作业状态修改为就绪，并返回成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) UpdateStreamingJob(request *model.UpdateStreamingJobRequest) (*model.UpdateStreamingJobResponse, error) {
	requestDef := GenReqDefForUpdateStreamingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStreamingJobResponse), nil
	}
}

// UpdateStreamingJobInvoker 更新实时作业
func (c *IoTAnalyticsClient) UpdateStreamingJobInvoker(request *model.UpdateStreamingJobRequest) *UpdateStreamingJobInvoker {
	requestDef := GenReqDefForUpdateStreamingJob()
	return &UpdateStreamingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartJob 启动实时作业
//
// 提交作业到运行环境，实时接收数据并按用户定义的业务逻辑对数据进行处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) StartJob(request *model.StartJobRequest) (*model.StartJobResponse, error) {
	requestDef := GenReqDefForStartJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartJobResponse), nil
	}
}

// StartJobInvoker 启动实时作业
func (c *IoTAnalyticsClient) StartJobInvoker(request *model.StartJobRequest) *StartJobInvoker {
	requestDef := GenReqDefForStartJob()
	return &StartJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJob 停止实时作业
//
// 停止一个正在运行中的作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

// StopJobInvoker 停止实时作业
func (c *IoTAnalyticsClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRun 执行离线作业
//
// 执行离线作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateRun(request *model.CreateRunRequest) (*model.CreateRunResponse, error) {
	requestDef := GenReqDefForCreateRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRunResponse), nil
	}
}

// CreateRunInvoker 执行离线作业
func (c *IoTAnalyticsClient) CreateRunInvoker(request *model.CreateRunRequest) *CreateRunInvoker {
	requestDef := GenReqDefForCreateRun()
	return &CreateRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRun 停止离线作业
//
// 停止提交中或运行中的离线作业，若作业已经执行结束或失败则无法停止。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteRun(request *model.DeleteRunRequest) (*model.DeleteRunResponse, error) {
	requestDef := GenReqDefForDeleteRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRunResponse), nil
	}
}

// DeleteRunInvoker 停止离线作业
func (c *IoTAnalyticsClient) DeleteRunInvoker(request *model.DeleteRunRequest) *DeleteRunInvoker {
	requestDef := GenReqDefForDeleteRun()
	return &DeleteRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuns 查询离线作业运行列表
//
// 查询离线作业运行列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListRuns(request *model.ListRunsRequest) (*model.ListRunsResponse, error) {
	requestDef := GenReqDefForListRuns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRunsResponse), nil
	}
}

// ListRunsInvoker 查询离线作业运行列表
func (c *IoTAnalyticsClient) ListRunsInvoker(request *model.ListRunsRequest) *ListRunsInvoker {
	requestDef := GenReqDefForListRuns()
	return &ListRunsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRun 查询离线作业运行详情
//
// 查询离线作业运行详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowRun(request *model.ShowRunRequest) (*model.ShowRunResponse, error) {
	requestDef := GenReqDefForShowRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRunResponse), nil
	}
}

// ShowRunInvoker 查询离线作业运行详情
func (c *IoTAnalyticsClient) ShowRunInvoker(request *model.ShowRunRequest) *ShowRunInvoker {
	requestDef := GenReqDefForShowRun()
	return &ShowRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTable 创建离线数据表
//
// 创建离线数据表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

// CreateTableInvoker 创建离线数据表
func (c *IoTAnalyticsClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTable 删除离线数据表
//
// 删除离线数据表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) DeleteTable(request *model.DeleteTableRequest) (*model.DeleteTableResponse, error) {
	requestDef := GenReqDefForDeleteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableResponse), nil
	}
}

// DeleteTableInvoker 删除离线数据表
func (c *IoTAnalyticsClient) DeleteTableInvoker(request *model.DeleteTableRequest) *DeleteTableInvoker {
	requestDef := GenReqDefForDeleteTable()
	return &DeleteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTables 查询离线数据表列表
//
// 查询离线数据表列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ListTables(request *model.ListTablesRequest) (*model.ListTablesResponse, error) {
	requestDef := GenReqDefForListTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesResponse), nil
	}
}

// ListTablesInvoker 查询离线数据表列表
func (c *IoTAnalyticsClient) ListTablesInvoker(request *model.ListTablesRequest) *ListTablesInvoker {
	requestDef := GenReqDefForListTables()
	return &ListTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTablePreview 预览离线数据表内容
//
// 预览离线数据表内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowTablePreview(request *model.ShowTablePreviewRequest) (*model.ShowTablePreviewResponse, error) {
	requestDef := GenReqDefForShowTablePreview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTablePreviewResponse), nil
	}
}

// ShowTablePreviewInvoker 预览离线数据表内容
func (c *IoTAnalyticsClient) ShowTablePreviewInvoker(request *model.ShowTablePreviewRequest) *ShowTablePreviewInvoker {
	requestDef := GenReqDefForShowTablePreview()
	return &ShowTablePreviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTableSchema 查询离线数据表结构
//
// 查询离线数据表结构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTAnalyticsClient) ShowTableSchema(request *model.ShowTableSchemaRequest) (*model.ShowTableSchemaResponse, error) {
	requestDef := GenReqDefForShowTableSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableSchemaResponse), nil
	}
}

// ShowTableSchemaInvoker 查询离线数据表结构
func (c *IoTAnalyticsClient) ShowTableSchemaInvoker(request *model.ShowTableSchemaRequest) *ShowTableSchemaInvoker {
	requestDef := GenReqDefForShowTableSchema()
	return &ShowTableSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
