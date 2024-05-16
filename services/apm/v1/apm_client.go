package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type ApmClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewApmClient(hcClient *httpclient.HcHttpClient) *ApmClient {
	return &ApmClient{HcClient: hcClient}
}

func ApmClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAkSk 创建aksk
//
// 创建aksk。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 删除aksk。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 查询租户的aksk。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListAlarmData 查询告警列表
//
// 查询系统中存在的告警。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ListAlarmData(request *model.ListAlarmDataRequest) (*model.ListAlarmDataResponse, error) {
	requestDef := GenReqDefForListAlarmData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmDataResponse), nil
	}
}

// ListAlarmDataInvoker 查询告警列表
func (c *ApmClient) ListAlarmDataInvoker(request *model.ListAlarmDataRequest) *ListAlarmDataInvoker {
	requestDef := GenReqDefForListAlarmData()
	return &ListAlarmDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmNotify 查询告警消息列表
//
// 查询单个告警的触发详情与历史。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ListAlarmNotify(request *model.ListAlarmNotifyRequest) (*model.ListAlarmNotifyResponse, error) {
	requestDef := GenReqDefForListAlarmNotify()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmNotifyResponse), nil
	}
}

// ListAlarmNotifyInvoker 查询告警消息列表
func (c *ApmClient) ListAlarmNotifyInvoker(request *model.ListAlarmNotifyRequest) *ListAlarmNotifyInvoker {
	requestDef := GenReqDefForListAlarmNotify()
	return &ListAlarmNotifyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAgentStatus 更改实例的采集状态
//
// 改变指定实例的采集状态：开启和关闭。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ChangeAgentStatus(request *model.ChangeAgentStatusRequest) (*model.ChangeAgentStatusResponse, error) {
	requestDef := GenReqDefForChangeAgentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAgentStatusResponse), nil
	}
}

// ChangeAgentStatusInvoker 更改实例的采集状态
func (c *ApmClient) ChangeAgentStatusInvoker(request *model.ChangeAgentStatusRequest) *ChangeAgentStatusInvoker {
	requestDef := GenReqDefForChangeAgentStatus()
	return &ChangeAgentStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgent 删除agent
//
// 删除agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) DeleteAgent(request *model.DeleteAgentRequest) (*model.DeleteAgentResponse, error) {
	requestDef := GenReqDefForDeleteAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgentResponse), nil
	}
}

// DeleteAgentInvoker 删除agent
func (c *ApmClient) DeleteAgentInvoker(request *model.DeleteAgentRequest) *DeleteAgentInvoker {
	requestDef := GenReqDefForDeleteAgent()
	return &DeleteAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAkSk 获取ak/sk
//
// 获取该用户创建的ak/sk列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListBusiness 查询应用列表
//
// 该接口用于查询对应用户下的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}

// ListBusinessInvoker 查询应用列表
func (c *ApmClient) ListBusinessInvoker(request *model.ListBusinessRequest) *ListBusinessInvoker {
	requestDef := GenReqDefForListBusiness()
	return &ListBusinessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvMonitorItem 查询监控项列表
//
// 查询监控项列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 保存监控项。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// SearchAgent 查询应用下所有探针
//
// 该接口用于搜索应用下所有探针情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchAgent(request *model.SearchAgentRequest) (*model.SearchAgentResponse, error) {
	requestDef := GenReqDefForSearchAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAgentResponse), nil
	}
}

// SearchAgentInvoker 查询应用下所有探针
func (c *ApmClient) SearchAgentInvoker(request *model.SearchAgentRequest) *SearchAgentInvoker {
	requestDef := GenReqDefForSearchAgent()
	return &SearchAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchApplication 对指定区域下的组件和环境及其探针情况进行搜索
//
// 对指定区域下的组件和环境及其探针情况进行搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchApplication(request *model.SearchApplicationRequest) (*model.SearchApplicationResponse, error) {
	requestDef := GenReqDefForSearchApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchApplicationResponse), nil
	}
}

// SearchApplicationInvoker 对指定区域下的组件和环境及其探针情况进行搜索
func (c *ApmClient) SearchApplicationInvoker(request *model.SearchApplicationRequest) *SearchApplicationInvoker {
	requestDef := GenReqDefForSearchApplication()
	return &SearchApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMasterAddress 查询master地址
//
// 根据region名称获取该region下的master服务podlb地址信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 该接口用于删除指定的组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListAppEnvs 获取组件下的环境列表
//
// 获取组件下的环境列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取组件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 查询环境标签接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowBusinessDetail 查询单个应用的详情
//
// 查询单个应用的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowBusinessDetail(request *model.ShowBusinessDetailRequest) (*model.ShowBusinessDetailResponse, error) {
	requestDef := GenReqDefForShowBusinessDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBusinessDetailResponse), nil
	}
}

// ShowBusinessDetailInvoker 查询单个应用的详情
func (c *ApmClient) ShowBusinessDetailInvoker(request *model.ShowBusinessDetailRequest) *ShowBusinessDetailInvoker {
	requestDef := GenReqDefForShowBusinessDetail()
	return &ShowBusinessDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubBusinessDetail 查询子应用详情
//
// 查询单个子应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowSubBusinessDetail(request *model.ShowSubBusinessDetailRequest) (*model.ShowSubBusinessDetailResponse, error) {
	requestDef := GenReqDefForShowSubBusinessDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubBusinessDetailResponse), nil
	}
}

// ShowSubBusinessDetailInvoker 查询子应用详情
func (c *ApmClient) ShowSubBusinessDetailInvoker(request *model.ShowSubBusinessDetailRequest) *ShowSubBusinessDetailInvoker {
	requestDef := GenReqDefForShowSubBusinessDetail()
	return &ShowSubBusinessDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopologyTree 获取应用树
//
// 获取应用树。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowTopologyTree(request *model.ShowTopologyTreeRequest) (*model.ShowTopologyTreeResponse, error) {
	requestDef := GenReqDefForShowTopologyTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopologyTreeResponse), nil
	}
}

// ShowTopologyTreeInvoker 获取应用树
func (c *ApmClient) ShowTopologyTreeInvoker(request *model.ShowTopologyTreeRequest) *ShowTopologyTreeInvoker {
	requestDef := GenReqDefForShowTopologyTree()
	return &ShowTopologyTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlameLineTree
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowFlameLineTree(request *model.ShowFlameLineTreeRequest) (*model.ShowFlameLineTreeResponse, error) {
	requestDef := GenReqDefForShowFlameLineTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlameLineTreeResponse), nil
	}
}

// ShowFlameLineTreeInvoker
func (c *ApmClient) ShowFlameLineTreeInvoker(request *model.ShowFlameLineTreeRequest) *ShowFlameLineTreeInvoker {
	requestDef := GenReqDefForShowFlameLineTree()
	return &ShowFlameLineTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOpenRegion 查询开通的region
//
// 该接口用于查询用户开通的region信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// SearchBusinessTopology 查询应用全局拓扑图
//
// 查询应用级别全局拓扑图信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchBusinessTopology(request *model.SearchBusinessTopologyRequest) (*model.SearchBusinessTopologyResponse, error) {
	requestDef := GenReqDefForSearchBusinessTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchBusinessTopologyResponse), nil
	}
}

// SearchBusinessTopologyInvoker 查询应用全局拓扑图
func (c *ApmClient) SearchBusinessTopologyInvoker(request *model.SearchBusinessTopologyRequest) *SearchBusinessTopologyInvoker {
	requestDef := GenReqDefForSearchBusinessTopology()
	return &SearchBusinessTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchEnvTopology 查询组件环境拓扑图
//
// 查询组件环境级别全局拓扑图信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchEnvTopology(request *model.SearchEnvTopologyRequest) (*model.SearchEnvTopologyResponse, error) {
	requestDef := GenReqDefForSearchEnvTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchEnvTopologyResponse), nil
	}
}

// SearchEnvTopologyInvoker 查询组件环境拓扑图
func (c *ApmClient) SearchEnvTopologyInvoker(request *model.SearchEnvTopologyRequest) *SearchEnvTopologyInvoker {
	requestDef := GenReqDefForSearchEnvTopology()
	return &SearchEnvTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBusiness 创建链路追踪应用
//
// 创建链路追踪应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) CreateBusiness(request *model.CreateBusinessRequest) (*model.CreateBusinessResponse, error) {
	requestDef := GenReqDefForCreateBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBusinessResponse), nil
	}
}

// CreateBusinessInvoker 创建链路追踪应用
func (c *ApmClient) CreateBusinessInvoker(request *model.CreateBusinessRequest) *CreateBusinessInvoker {
	requestDef := GenReqDefForCreateBusiness()
	return &CreateBusinessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessPoint 获取链路追踪应用接入地址
//
// 获取链路追踪应用接入地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowAccessPoint(request *model.ShowAccessPointRequest) (*model.ShowAccessPointResponse, error) {
	requestDef := GenReqDefForShowAccessPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessPointResponse), nil
	}
}

// ShowAccessPointInvoker 获取链路追踪应用接入地址
func (c *ApmClient) ShowAccessPointInvoker(request *model.ShowAccessPointRequest) *ShowAccessPointInvoker {
	requestDef := GenReqDefForShowAccessPoint()
	return &ShowAccessPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowToken 获取链路追踪应用的token
//
// 获取链路追踪应用的token
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowToken(request *model.ShowTokenRequest) (*model.ShowTokenResponse, error) {
	requestDef := GenReqDefForShowToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTokenResponse), nil
	}
}

// ShowTokenInvoker 获取链路追踪应用的token
func (c *ApmClient) ShowTokenInvoker(request *model.ShowTokenRequest) *ShowTokenInvoker {
	requestDef := GenReqDefForShowToken()
	return &ShowTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusinessEnv 查询URL跟踪Region环境列表
//
// 查询所选Region下设置了URL跟踪的环境列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ListBusinessEnv(request *model.ListBusinessEnvRequest) (*model.ListBusinessEnvResponse, error) {
	requestDef := GenReqDefForListBusinessEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessEnvResponse), nil
	}
}

// ListBusinessEnvInvoker 查询URL跟踪Region环境列表
func (c *ApmClient) ListBusinessEnvInvoker(request *model.ListBusinessEnvRequest) *ListBusinessEnvInvoker {
	requestDef := GenReqDefForListBusinessEnv()
	return &ListBusinessEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchTransaction 查询URL跟踪视图列表
//
// 查询当前被调用的URL跟踪视图列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchTransaction(request *model.SearchTransactionRequest) (*model.SearchTransactionResponse, error) {
	requestDef := GenReqDefForSearchTransaction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchTransactionResponse), nil
	}
}

// SearchTransactionInvoker 查询URL跟踪视图列表
func (c *ApmClient) SearchTransactionInvoker(request *model.SearchTransactionRequest) *SearchTransactionInvoker {
	requestDef := GenReqDefForSearchTransaction()
	return &SearchTransactionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchTransactionConfig 查询URL跟踪配置列表
//
// 查询已配置好的URL跟踪配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) SearchTransactionConfig(request *model.SearchTransactionConfigRequest) (*model.SearchTransactionConfigResponse, error) {
	requestDef := GenReqDefForSearchTransactionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchTransactionConfigResponse), nil
	}
}

// SearchTransactionConfigInvoker 查询URL跟踪配置列表
func (c *ApmClient) SearchTransactionConfigInvoker(request *model.SearchTransactionConfigRequest) *SearchTransactionConfigInvoker {
	requestDef := GenReqDefForSearchTransactionConfig()
	return &SearchTransactionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransactionDetail 查询URL跟踪视图详情
//
// 查询某条URL跟踪视图详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowTransactionDetail(request *model.ShowTransactionDetailRequest) (*model.ShowTransactionDetailResponse, error) {
	requestDef := GenReqDefForShowTransactionDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransactionDetailResponse), nil
	}
}

// ShowTransactionDetailInvoker 查询URL跟踪视图详情
func (c *ApmClient) ShowTransactionDetailInvoker(request *model.ShowTransactionDetailRequest) *ShowTransactionDetailInvoker {
	requestDef := GenReqDefForShowTransactionDetail()
	return &ShowTransactionDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvInstances 获取实例信息列表
//
// 获取实例信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取原始数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowEnvMonitorItems 获取监控项信息
//
// 获取监控项信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowEnvMonitorItems(request *model.ShowEnvMonitorItemsRequest) (*model.ShowEnvMonitorItemsResponse, error) {
	requestDef := GenReqDefForShowEnvMonitorItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvMonitorItemsResponse), nil
	}
}

// ShowEnvMonitorItemsInvoker 获取监控项信息
func (c *ApmClient) ShowEnvMonitorItemsInvoker(request *model.ShowEnvMonitorItemsRequest) *ShowEnvMonitorItemsInvoker {
	requestDef := GenReqDefForShowEnvMonitorItems()
	return &ShowEnvMonitorItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEventDetail 获取event的详情
//
// 获取event的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowMonitorItemDetail 获取一个监控项的详情
//
// 获取一个监控项的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApmClient) ShowMonitorItemDetail(request *model.ShowMonitorItemDetailRequest) (*model.ShowMonitorItemDetailResponse, error) {
	requestDef := GenReqDefForShowMonitorItemDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitorItemDetailResponse), nil
	}
}

// ShowMonitorItemDetailInvoker 获取一个监控项的详情
func (c *ApmClient) ShowMonitorItemDetailInvoker(request *model.ShowMonitorItemDetailRequest) *ShowMonitorItemDetailInvoker {
	requestDef := GenReqDefForShowMonitorItemDetail()
	return &ShowMonitorItemDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitorItemViewConfig 查询监控项配置信息
//
// 查询监控项配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取原始数据表格。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// span数据查询接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取汇总表格数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 调用链拓扑图。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取一个trace的所有调用链数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 获取趋势图。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
