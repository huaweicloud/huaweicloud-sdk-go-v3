package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type ApmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewApmClient(hcClient *http_client.HcHttpClient) *ApmClient {
	return &ApmClient{HcClient: hcClient}
}

func ApmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateAkSk 创建aksk
//
// 创建aksk
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) CreateAkSk(request *model.CreateAkSkRequest) (*model.CreateAkSkResponse, error) {
	requestDef := GenReqDefForCreateAkSk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAkSkResponse), nil
	}
}

// CreateAkSkInvoker 创建aksk
func (c *ApmClient) CreateAkSkInvoker(request *model.CreateAkSkRequest) *CreateAkSkInvoker {
	requestDef := GenReqDefForCreateAkSk()
	return &CreateAkSkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAkSk 删除aksk
//
// 删除aksk
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) DeleteAkSk(request *model.DeleteAkSkRequest) (*model.DeleteAkSkResponse, error) {
	requestDef := GenReqDefForDeleteAkSk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAkSkResponse), nil
	}
}

// DeleteAkSkInvoker 删除aksk
func (c *ApmClient) DeleteAkSkInvoker(request *model.DeleteAkSkRequest) *DeleteAkSkInvoker {
	requestDef := GenReqDefForDeleteAkSk()
	return &DeleteAkSkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAkSks 查询租户的aksk
//
// 查询租户的aksk
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowAkSks(request *model.ShowAkSksRequest) (*model.ShowAkSksResponse, error) {
	requestDef := GenReqDefForShowAkSks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAkSksResponse), nil
	}
}

// ShowAkSksInvoker 查询租户的aksk
func (c *ApmClient) ShowAkSksInvoker(request *model.ShowAkSksRequest) *ShowAkSksInvoker {
	requestDef := GenReqDefForShowAkSks()
	return &ShowAkSksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAkSk 获取ak/sk
//
// 获取该用户创建的ak/sk列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListAkSk(request *model.ListAkSkRequest) (*model.ListAkSkResponse, error) {
	requestDef := GenReqDefForListAkSk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAkSkResponse), nil
	}
}

// ListAkSkInvoker 获取ak/sk
func (c *ApmClient) ListAkSkInvoker(request *model.ListAkSkRequest) *ListAkSkInvoker {
	requestDef := GenReqDefForListAkSk()
	return &ListAkSkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusiness 查询业务列表
//
// 该接口用于查询对应用户下的业务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}

// ListBusinessInvoker 查询业务列表
func (c *ApmClient) ListBusinessInvoker(request *model.ListBusinessRequest) *ListBusinessInvoker {
	requestDef := GenReqDefForListBusiness()
	return &ListBusinessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvMonitorItem 查询监控项列表
//
// 查询监控项列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListEnvMonitorItem(request *model.ListEnvMonitorItemRequest) (*model.ListEnvMonitorItemResponse, error) {
	requestDef := GenReqDefForListEnvMonitorItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvMonitorItemResponse), nil
	}
}

// ListEnvMonitorItemInvoker 查询监控项列表
func (c *ApmClient) ListEnvMonitorItemInvoker(request *model.ListEnvMonitorItemRequest) *ListEnvMonitorItemInvoker {
	requestDef := GenReqDefForListEnvMonitorItem()
	return &ListEnvMonitorItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveMonitorItemConfig 保存监控项
//
// 保存监控项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) SaveMonitorItemConfig(request *model.SaveMonitorItemConfigRequest) (*model.SaveMonitorItemConfigResponse, error) {
	requestDef := GenReqDefForSaveMonitorItemConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveMonitorItemConfigResponse), nil
	}
}

// SaveMonitorItemConfigInvoker 保存监控项
func (c *ApmClient) SaveMonitorItemConfigInvoker(request *model.SaveMonitorItemConfigRequest) *SaveMonitorItemConfigInvoker {
	requestDef := GenReqDefForSaveMonitorItemConfig()
	return &SaveMonitorItemConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchApplication 对指定区域下的应用和环境及其探针情况进行搜索
//
// 对指定区域下的应用和环境及其探针情况进行搜索
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) SearchApplication(request *model.SearchApplicationRequest) (*model.SearchApplicationResponse, error) {
	requestDef := GenReqDefForSearchApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchApplicationResponse), nil
	}
}

// SearchApplicationInvoker 对指定区域下的应用和环境及其探针情况进行搜索
func (c *ApmClient) SearchApplicationInvoker(request *model.SearchApplicationRequest) *SearchApplicationInvoker {
	requestDef := GenReqDefForSearchApplication()
	return &SearchApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMasterAddress 查询master地址
//
// 根据region名称获取该名称下的master服务podlb地址信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowMasterAddress(request *model.ShowMasterAddressRequest) (*model.ShowMasterAddressResponse, error) {
	requestDef := GenReqDefForShowMasterAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterAddressResponse), nil
	}
}

// ShowMasterAddressInvoker 查询master地址
func (c *ApmClient) ShowMasterAddressInvoker(request *model.ShowMasterAddressRequest) *ShowMasterAddressInvoker {
	requestDef := GenReqDefForShowMasterAddress()
	return &ShowMasterAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 根据组件id删除指定的组件
//
// 该接口用于删除指定的组件
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 根据组件id删除指定的组件
func (c *ApmClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnv 根据环境id删除指定的环境
//
// 该接口用于删除指定的环境
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) DeleteEnv(request *model.DeleteEnvRequest) (*model.DeleteEnvResponse, error) {
	requestDef := GenReqDefForDeleteEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvResponse), nil
	}
}

// DeleteEnvInvoker 根据环境id删除指定的环境
func (c *ApmClient) DeleteEnvInvoker(request *model.DeleteEnvRequest) *DeleteEnvInvoker {
	requestDef := GenReqDefForDeleteEnv()
	return &DeleteEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppEnvs 获取组件下的环境列表
//
// 获取组件下的环境列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListAppEnvs(request *model.ListAppEnvsRequest) (*model.ListAppEnvsResponse, error) {
	requestDef := GenReqDefForListAppEnvs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppEnvsResponse), nil
	}
}

// ListAppEnvsInvoker 获取组件下的环境列表
func (c *ApmClient) ListAppEnvsInvoker(request *model.ListAppEnvsRequest) *ListAppEnvsInvoker {
	requestDef := GenReqDefForListAppEnvs()
	return &ListAppEnvsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 获取组件列表
//
// 获取组件列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 获取组件列表
func (c *ApmClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvTags 查询环境标签
//
// 查询环境标签接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListEnvTags(request *model.ListEnvTagsRequest) (*model.ListEnvTagsResponse, error) {
	requestDef := GenReqDefForListEnvTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvTagsResponse), nil
	}
}

// ListEnvTagsInvoker 查询环境标签
func (c *ApmClient) ListEnvTagsInvoker(request *model.ListEnvTagsRequest) *ListEnvTagsInvoker {
	requestDef := GenReqDefForListEnvTags()
	return &ListEnvTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopologyTree 获取业务树
//
// 获取业务树
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowTopologyTree(request *model.ShowTopologyTreeRequest) (*model.ShowTopologyTreeResponse, error) {
	requestDef := GenReqDefForShowTopologyTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopologyTreeResponse), nil
	}
}

// ShowTopologyTreeInvoker 获取业务树
func (c *ApmClient) ShowTopologyTreeInvoker(request *model.ShowTopologyTreeRequest) *ShowTopologyTreeInvoker {
	requestDef := GenReqDefForShowTopologyTree()
	return &ShowTopologyTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOpenRegion 查询开通的region
//
// 该接口用于查询用户开通的region信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListOpenRegion(request *model.ListOpenRegionRequest) (*model.ListOpenRegionResponse, error) {
	requestDef := GenReqDefForListOpenRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOpenRegionResponse), nil
	}
}

// ListOpenRegionInvoker 查询开通的region
func (c *ApmClient) ListOpenRegionInvoker(request *model.ListOpenRegionRequest) *ListOpenRegionInvoker {
	requestDef := GenReqDefForListOpenRegion()
	return &ListOpenRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportedRegion 查询所有的支持的region
//
// 查询所有的支持的region信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListSupportedRegion(request *model.ListSupportedRegionRequest) (*model.ListSupportedRegionResponse, error) {
	requestDef := GenReqDefForListSupportedRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedRegionResponse), nil
	}
}

// ListSupportedRegionInvoker 查询所有的支持的region
func (c *ApmClient) ListSupportedRegionInvoker(request *model.ListSupportedRegionRequest) *ListSupportedRegionInvoker {
	requestDef := GenReqDefForListSupportedRegion()
	return &ListSupportedRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvInstances 获取实例信息列表
//
// 获取实例信息列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListEnvInstances(request *model.ListEnvInstancesRequest) (*model.ListEnvInstancesResponse, error) {
	requestDef := GenReqDefForListEnvInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvInstancesResponse), nil
	}
}

// ListEnvInstancesInvoker 获取实例信息列表
func (c *ApmClient) ListEnvInstancesInvoker(request *model.ListEnvInstancesRequest) *ListEnvInstancesInvoker {
	requestDef := GenReqDefForListEnvInstances()
	return &ListEnvInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClobDetail 获取原始数据详情
//
// 获取原始数据详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowClobDetail(request *model.ShowClobDetailRequest) (*model.ShowClobDetailResponse, error) {
	requestDef := GenReqDefForShowClobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClobDetailResponse), nil
	}
}

// ShowClobDetailInvoker 获取原始数据详情
func (c *ApmClient) ShowClobDetailInvoker(request *model.ShowClobDetailRequest) *ShowClobDetailInvoker {
	requestDef := GenReqDefForShowClobDetail()
	return &ShowClobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnvMonitorItems
//
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowEnvMonitorItems(request *model.ShowEnvMonitorItemsRequest) (*model.ShowEnvMonitorItemsResponse, error) {
	requestDef := GenReqDefForShowEnvMonitorItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvMonitorItemsResponse), nil
	}
}

// ShowEnvMonitorItemsInvoker
func (c *ApmClient) ShowEnvMonitorItemsInvoker(request *model.ShowEnvMonitorItemsRequest) *ShowEnvMonitorItemsInvoker {
	requestDef := GenReqDefForShowEnvMonitorItems()
	return &ShowEnvMonitorItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEventDetail 获取event的详情
//
// 获取event的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowEventDetail(request *model.ShowEventDetailRequest) (*model.ShowEventDetailResponse, error) {
	requestDef := GenReqDefForShowEventDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventDetailResponse), nil
	}
}

// ShowEventDetailInvoker 获取event的详情
func (c *ApmClient) ShowEventDetailInvoker(request *model.ShowEventDetailRequest) *ShowEventDetailInvoker {
	requestDef := GenReqDefForShowEventDetail()
	return &ShowEventDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitorItemViewConfig 查询监控项配置信息
//
// 查询监控项配置信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowMonitorItemViewConfig(request *model.ShowMonitorItemViewConfigRequest) (*model.ShowMonitorItemViewConfigResponse, error) {
	requestDef := GenReqDefForShowMonitorItemViewConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitorItemViewConfigResponse), nil
	}
}

// ShowMonitorItemViewConfigInvoker 查询监控项配置信息
func (c *ApmClient) ShowMonitorItemViewConfigInvoker(request *model.ShowMonitorItemViewConfigRequest) *ShowMonitorItemViewConfigInvoker {
	requestDef := GenReqDefForShowMonitorItemViewConfig()
	return &ShowMonitorItemViewConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRawTable 获取原始数据表格
//
// 获取原始数据表格
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowRawTable(request *model.ShowRawTableRequest) (*model.ShowRawTableResponse, error) {
	requestDef := GenReqDefForShowRawTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRawTableResponse), nil
	}
}

// ShowRawTableInvoker 获取原始数据表格
func (c *ApmClient) ShowRawTableInvoker(request *model.ShowRawTableRequest) *ShowRawTableInvoker {
	requestDef := GenReqDefForShowRawTable()
	return &ShowRawTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpanSearch 查询span数据
//
// span数据查询接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowSpanSearch(request *model.ShowSpanSearchRequest) (*model.ShowSpanSearchResponse, error) {
	requestDef := GenReqDefForShowSpanSearch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpanSearchResponse), nil
	}
}

// ShowSpanSearchInvoker 查询span数据
func (c *ApmClient) ShowSpanSearchInvoker(request *model.ShowSpanSearchRequest) *ShowSpanSearchInvoker {
	requestDef := GenReqDefForShowSpanSearch()
	return &ShowSpanSearchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSumTable 获取汇总表格数据
//
// 获取汇总表格数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowSumTable(request *model.ShowSumTableRequest) (*model.ShowSumTableResponse, error) {
	requestDef := GenReqDefForShowSumTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSumTableResponse), nil
	}
}

// ShowSumTableInvoker 获取汇总表格数据
func (c *ApmClient) ShowSumTableInvoker(request *model.ShowSumTableRequest) *ShowSumTableInvoker {
	requestDef := GenReqDefForShowSumTable()
	return &ShowSumTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopology 调用链拓扑图
//
// 调用链拓扑图
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowTopology(request *model.ShowTopologyRequest) (*model.ShowTopologyResponse, error) {
	requestDef := GenReqDefForShowTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopologyResponse), nil
	}
}

// ShowTopologyInvoker 调用链拓扑图
func (c *ApmClient) ShowTopologyInvoker(request *model.ShowTopologyRequest) *ShowTopologyInvoker {
	requestDef := GenReqDefForShowTopology()
	return &ShowTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTraceEvents 获取一个trace的所有调用链数据
//
// 获取一个trace的所有调用链数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowTraceEvents(request *model.ShowTraceEventsRequest) (*model.ShowTraceEventsResponse, error) {
	requestDef := GenReqDefForShowTraceEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTraceEventsResponse), nil
	}
}

// ShowTraceEventsInvoker 获取一个trace的所有调用链数据
func (c *ApmClient) ShowTraceEventsInvoker(request *model.ShowTraceEventsRequest) *ShowTraceEventsInvoker {
	requestDef := GenReqDefForShowTraceEvents()
	return &ShowTraceEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrend 获取趋势图
//
// 获取趋势图
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowTrend(request *model.ShowTrendRequest) (*model.ShowTrendResponse, error) {
	requestDef := GenReqDefForShowTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrendResponse), nil
	}
}

// ShowTrendInvoker 获取趋势图
func (c *ApmClient) ShowTrendInvoker(request *model.ShowTrendRequest) *ShowTrendInvoker {
	requestDef := GenReqDefForShowTrend()
	return &ShowTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
