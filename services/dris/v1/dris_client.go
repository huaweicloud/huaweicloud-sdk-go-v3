package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dris/v1/model"
)

type DrisClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDrisClient(hcClient *http_client.HcHttpClient) *DrisClient {
	return &DrisClient{HcClient: hcClient}
}

func DrisClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddForwardingConfigs 创建数据转发配置
//
// 创建数据转发配置。当前仅支持数据转发至kafka，数据转发配置成功添加后配置中的Topic消息将会转发至指定的brokers。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) AddForwardingConfigs(request *model.AddForwardingConfigsRequest) (*model.AddForwardingConfigsResponse, error) {
	requestDef := GenReqDefForAddForwardingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddForwardingConfigsResponse), nil
	}
}

// AddForwardingConfigsInvoker 创建数据转发配置
func (c *DrisClient) AddForwardingConfigsInvoker(request *model.AddForwardingConfigsRequest) *AddForwardingConfigsInvoker {
	requestDef := GenReqDefForAddForwardingConfigs()
	return &AddForwardingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteForwardingConfig 删除数据转发配置
//
// 根据转发配置的唯一ID（forwarding_config_id）删除数据转发配置，删除后配置中订阅的topic消息将不会被转发至brokers。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteForwardingConfig(request *model.DeleteForwardingConfigRequest) (*model.DeleteForwardingConfigResponse, error) {
	requestDef := GenReqDefForDeleteForwardingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteForwardingConfigResponse), nil
	}
}

// DeleteForwardingConfigInvoker 删除数据转发配置
func (c *DrisClient) DeleteForwardingConfigInvoker(request *model.DeleteForwardingConfigRequest) *DeleteForwardingConfigInvoker {
	requestDef := GenReqDefForDeleteForwardingConfig()
	return &DeleteForwardingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowForwardingConfig 查询数据转发配置
//
// 根据转发配置的唯一ID（forwarding_config_id）查询数据转发配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowForwardingConfig(request *model.ShowForwardingConfigRequest) (*model.ShowForwardingConfigResponse, error) {
	requestDef := GenReqDefForShowForwardingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowForwardingConfigResponse), nil
	}
}

// ShowForwardingConfigInvoker 查询数据转发配置
func (c *DrisClient) ShowForwardingConfigInvoker(request *model.ShowForwardingConfigRequest) *ShowForwardingConfigInvoker {
	requestDef := GenReqDefForShowForwardingConfig()
	return &ShowForwardingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowForwardingConfigs 查询数据转发配置列表
//
// 查询数据转发配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowForwardingConfigs(request *model.ShowForwardingConfigsRequest) (*model.ShowForwardingConfigsResponse, error) {
	requestDef := GenReqDefForShowForwardingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowForwardingConfigsResponse), nil
	}
}

// ShowForwardingConfigsInvoker 查询数据转发配置列表
func (c *DrisClient) ShowForwardingConfigsInvoker(request *model.ShowForwardingConfigsRequest) *ShowForwardingConfigsInvoker {
	requestDef := GenReqDefForShowForwardingConfigs()
	return &ShowForwardingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateForwardingConfig 修改数据转发配置
//
// 根据转发配置的唯一ID（forwarding_config_id）修改数据转发配置，当前支持更新的字段有topicPrefix、userTopics、brokers，需要把该字段新的对应值全量写入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateForwardingConfig(request *model.UpdateForwardingConfigRequest) (*model.UpdateForwardingConfigResponse, error) {
	requestDef := GenReqDefForUpdateForwardingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateForwardingConfigResponse), nil
	}
}

// UpdateForwardingConfigInvoker 修改数据转发配置
func (c *DrisClient) UpdateForwardingConfigInvoker(request *model.UpdateForwardingConfigRequest) *UpdateForwardingConfigInvoker {
	requestDef := GenReqDefForUpdateForwardingConfig()
	return &UpdateForwardingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeFlows 查询历史交通统计信息列表
//
// 查询历史交通统计信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ListEdgeFlows(request *model.ListEdgeFlowsRequest) (*model.ListEdgeFlowsResponse, error) {
	requestDef := GenReqDefForListEdgeFlows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeFlowsResponse), nil
	}
}

// ListEdgeFlowsInvoker 查询历史交通统计信息列表
func (c *DrisClient) ListEdgeFlowsInvoker(request *model.ListEdgeFlowsRequest) *ListEdgeFlowsInvoker {
	requestDef := GenReqDefForListEdgeFlows()
	return &ListEdgeFlowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHistoryTrafficEvents 查询历史交通事件列表
//
// 查询历史交通事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowHistoryTrafficEvents(request *model.ShowHistoryTrafficEventsRequest) (*model.ShowHistoryTrafficEventsResponse, error) {
	requestDef := GenReqDefForShowHistoryTrafficEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTrafficEventsResponse), nil
	}
}

// ShowHistoryTrafficEventsInvoker 查询历史交通事件列表
func (c *DrisClient) ShowHistoryTrafficEventsInvoker(request *model.ShowHistoryTrafficEventsRequest) *ShowHistoryTrafficEventsInvoker {
	requestDef := GenReqDefForShowHistoryTrafficEvents()
	return &ShowHistoryTrafficEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendImmediateEvent 创建即时交通事件
//
// 创建即时交通事件，平台分发即时交通事件给目标设备的接口。事件一旦创建便会立即下发且只会下发一次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) SendImmediateEvent(request *model.SendImmediateEventRequest) (*model.SendImmediateEventResponse, error) {
	requestDef := GenReqDefForSendImmediateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendImmediateEventResponse), nil
	}
}

// SendImmediateEventInvoker 创建即时交通事件
func (c *DrisClient) SendImmediateEventInvoker(request *model.SendImmediateEventRequest) *SendImmediateEventInvoker {
	requestDef := GenReqDefForSendImmediateEvent()
	return &SendImmediateEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataChannel 创建业务通道
//
// 创建业务通道，用于创建Edge消息上报的数据通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateDataChannel(request *model.CreateDataChannelRequest) (*model.CreateDataChannelResponse, error) {
	requestDef := GenReqDefForCreateDataChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataChannelResponse), nil
	}
}

// CreateDataChannelInvoker 创建业务通道
func (c *DrisClient) CreateDataChannelInvoker(request *model.CreateDataChannelRequest) *CreateDataChannelInvoker {
	requestDef := GenReqDefForCreateDataChannel()
	return &CreateDataChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataChannel 删除业务通道
//
// 删除业务通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteDataChannel(request *model.DeleteDataChannelRequest) (*model.DeleteDataChannelResponse, error) {
	requestDef := GenReqDefForDeleteDataChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataChannelResponse), nil
	}
}

// DeleteDataChannelInvoker 删除业务通道
func (c *DrisClient) DeleteDataChannelInvoker(request *model.DeleteDataChannelRequest) *DeleteDataChannelInvoker {
	requestDef := GenReqDefForDeleteDataChannel()
	return &DeleteDataChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataChannel 查询业务通道
//
// 查询业务通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowDataChannel(request *model.ShowDataChannelRequest) (*model.ShowDataChannelResponse, error) {
	requestDef := GenReqDefForShowDataChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataChannelResponse), nil
	}
}

// ShowDataChannelInvoker 查询业务通道
func (c *DrisClient) ShowDataChannelInvoker(request *model.ShowDataChannelRequest) *ShowDataChannelInvoker {
	requestDef := GenReqDefForShowDataChannel()
	return &ShowDataChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataChannel 修改业务通道
//
// 修改业务通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateDataChannel(request *model.UpdateDataChannelRequest) (*model.UpdateDataChannelResponse, error) {
	requestDef := GenReqDefForUpdateDataChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataChannelResponse), nil
	}
}

// UpdateDataChannelInvoker 修改业务通道
func (c *DrisClient) UpdateDataChannelInvoker(request *model.UpdateDataChannelRequest) *UpdateDataChannelInvoker {
	requestDef := GenReqDefForUpdateDataChannel()
	return &UpdateDataChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2xEdge 创建Edge
//
// 创建Edge
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateV2xEdge(request *model.CreateV2xEdgeRequest) (*model.CreateV2xEdgeResponse, error) {
	requestDef := GenReqDefForCreateV2xEdge()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2xEdgeResponse), nil
	}
}

// CreateV2xEdgeInvoker 创建Edge
func (c *DrisClient) CreateV2xEdgeInvoker(request *model.CreateV2xEdgeRequest) *CreateV2xEdgeInvoker {
	requestDef := GenReqDefForCreateV2xEdge()
	return &CreateV2xEdgeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2XEdgeByV2xEdgeId 删除Edge
//
// 删除Edge之前需要删除Edge下的业务通道和关联设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteV2XEdgeByV2xEdgeId(request *model.DeleteV2XEdgeByV2xEdgeIdRequest) (*model.DeleteV2XEdgeByV2xEdgeIdResponse, error) {
	requestDef := GenReqDefForDeleteV2XEdgeByV2xEdgeId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2XEdgeByV2xEdgeIdResponse), nil
	}
}

// DeleteV2XEdgeByV2xEdgeIdInvoker 删除Edge
func (c *DrisClient) DeleteV2XEdgeByV2xEdgeIdInvoker(request *model.DeleteV2XEdgeByV2xEdgeIdRequest) *DeleteV2XEdgeByV2xEdgeIdInvoker {
	requestDef := GenReqDefForDeleteV2XEdgeByV2xEdgeId()
	return &DeleteV2XEdgeByV2xEdgeIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2xEdges 查询Edge列表
//
// 查询Edge列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ListV2xEdges(request *model.ListV2xEdgesRequest) (*model.ListV2xEdgesResponse, error) {
	requestDef := GenReqDefForListV2xEdges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2xEdgesResponse), nil
	}
}

// ListV2xEdgesInvoker 查询Edge列表
func (c *DrisClient) ListV2xEdgesInvoker(request *model.ListV2xEdgesRequest) *ListV2xEdgesInvoker {
	requestDef := GenReqDefForListV2xEdges()
	return &ListV2xEdgesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentCode 生成部署应用安装命令
//
// 生成部署应用安装命令,然后在ITS800或者ATLAS500上通过Shell执行
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowDeploymentCode(request *model.ShowDeploymentCodeRequest) (*model.ShowDeploymentCodeResponse, error) {
	requestDef := GenReqDefForShowDeploymentCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentCodeResponse), nil
	}
}

// ShowDeploymentCodeInvoker 生成部署应用安装命令
func (c *DrisClient) ShowDeploymentCodeInvoker(request *model.ShowDeploymentCodeRequest) *ShowDeploymentCodeInvoker {
	requestDef := GenReqDefForShowDeploymentCode()
	return &ShowDeploymentCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2xEdgeDetail 查询Edge
//
// 查询Edge
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowV2xEdgeDetail(request *model.ShowV2xEdgeDetailRequest) (*model.ShowV2xEdgeDetailResponse, error) {
	requestDef := GenReqDefForShowV2xEdgeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2xEdgeDetailResponse), nil
	}
}

// ShowV2xEdgeDetailInvoker 查询Edge
func (c *DrisClient) ShowV2xEdgeDetailInvoker(request *model.ShowV2xEdgeDetailRequest) *ShowV2xEdgeDetailInvoker {
	requestDef := GenReqDefForShowV2xEdgeDetail()
	return &ShowV2xEdgeDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2xEdge 修改Edge
//
// 修改Edge
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateV2xEdge(request *model.UpdateV2xEdgeRequest) (*model.UpdateV2xEdgeResponse, error) {
	requestDef := GenReqDefForUpdateV2xEdge()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2xEdgeResponse), nil
	}
}

// UpdateV2xEdgeInvoker 修改Edge
func (c *DrisClient) UpdateV2xEdgeInvoker(request *model.UpdateV2xEdgeRequest) *UpdateV2xEdgeInvoker {
	requestDef := GenReqDefForUpdateV2xEdge()
	return &UpdateV2xEdgeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowIpcs 查询IPC列表
//
// 获取多个IPC资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowIpcs(request *model.BatchShowIpcsRequest) (*model.BatchShowIpcsResponse, error) {
	requestDef := GenReqDefForBatchShowIpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowIpcsResponse), nil
	}
}

// BatchShowIpcsInvoker 查询IPC列表
func (c *DrisClient) BatchShowIpcsInvoker(request *model.BatchShowIpcsRequest) *BatchShowIpcsInvoker {
	requestDef := GenReqDefForBatchShowIpcs()
	return &BatchShowIpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpc 查询IPC
//
// 查询IPC
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowIpc(request *model.ShowIpcRequest) (*model.ShowIpcResponse, error) {
	requestDef := GenReqDefForShowIpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpcResponse), nil
	}
}

// ShowIpcInvoker 查询IPC
func (c *DrisClient) ShowIpcInvoker(request *model.ShowIpcRequest) *ShowIpcInvoker {
	requestDef := GenReqDefForShowIpc()
	return &ShowIpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowRadars 查询雷达列表
//
// 查询雷达列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowRadars(request *model.BatchShowRadarsRequest) (*model.BatchShowRadarsResponse, error) {
	requestDef := GenReqDefForBatchShowRadars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowRadarsResponse), nil
	}
}

// BatchShowRadarsInvoker 查询雷达列表
func (c *DrisClient) BatchShowRadarsInvoker(request *model.BatchShowRadarsRequest) *BatchShowRadarsInvoker {
	requestDef := GenReqDefForBatchShowRadars()
	return &BatchShowRadarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowRsus 查询RSU列表
//
// 查询RSU列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowRsus(request *model.BatchShowRsusRequest) (*model.BatchShowRsusResponse, error) {
	requestDef := GenReqDefForBatchShowRsus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowRsusResponse), nil
	}
}

// BatchShowRsusInvoker 查询RSU列表
func (c *DrisClient) BatchShowRsusInvoker(request *model.BatchShowRsusRequest) *BatchShowRsusInvoker {
	requestDef := GenReqDefForBatchShowRsus()
	return &BatchShowRsusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRsu 创建RSU
//
// 创建RSU
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateRsu(request *model.CreateRsuRequest) (*model.CreateRsuResponse, error) {
	requestDef := GenReqDefForCreateRsu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRsuResponse), nil
	}
}

// CreateRsuInvoker 创建RSU
func (c *DrisClient) CreateRsuInvoker(request *model.CreateRsuRequest) *CreateRsuInvoker {
	requestDef := GenReqDefForCreateRsu()
	return &CreateRsuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRsu 删除RSU
//
// 删除RSU
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteRsu(request *model.DeleteRsuRequest) (*model.DeleteRsuResponse, error) {
	requestDef := GenReqDefForDeleteRsu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRsuResponse), nil
	}
}

// DeleteRsuInvoker 删除RSU
func (c *DrisClient) DeleteRsuInvoker(request *model.DeleteRsuRequest) *DeleteRsuInvoker {
	requestDef := GenReqDefForDeleteRsu()
	return &DeleteRsuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRsu 修改RSU
//
// 修改RSU
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateRsu(request *model.UpdateRsuRequest) (*model.UpdateRsuResponse, error) {
	requestDef := GenReqDefForUpdateRsu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRsuResponse), nil
	}
}

// UpdateRsuInvoker 修改RSU
func (c *DrisClient) UpdateRsuInvoker(request *model.UpdateRsuRequest) *UpdateRsuInvoker {
	requestDef := GenReqDefForUpdateRsu()
	return &UpdateRsuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowTrafficControllers 查询信号机列表
//
// 查询信号机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowTrafficControllers(request *model.BatchShowTrafficControllersRequest) (*model.BatchShowTrafficControllersResponse, error) {
	requestDef := GenReqDefForBatchShowTrafficControllers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowTrafficControllersResponse), nil
	}
}

// BatchShowTrafficControllersInvoker 查询信号机列表
func (c *DrisClient) BatchShowTrafficControllersInvoker(request *model.BatchShowTrafficControllersRequest) *BatchShowTrafficControllersInvoker {
	requestDef := GenReqDefForBatchShowTrafficControllers()
	return &BatchShowTrafficControllersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrafficController 创建信号机
//
// 创建信号机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateTrafficController(request *model.CreateTrafficControllerRequest) (*model.CreateTrafficControllerResponse, error) {
	requestDef := GenReqDefForCreateTrafficController()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrafficControllerResponse), nil
	}
}

// CreateTrafficControllerInvoker 创建信号机
func (c *DrisClient) CreateTrafficControllerInvoker(request *model.CreateTrafficControllerRequest) *CreateTrafficControllerInvoker {
	requestDef := GenReqDefForCreateTrafficController()
	return &CreateTrafficControllerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrafficController 删除信号机
//
// 删除信号机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteTrafficController(request *model.DeleteTrafficControllerRequest) (*model.DeleteTrafficControllerResponse, error) {
	requestDef := GenReqDefForDeleteTrafficController()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrafficControllerResponse), nil
	}
}

// DeleteTrafficControllerInvoker 删除信号机
func (c *DrisClient) DeleteTrafficControllerInvoker(request *model.DeleteTrafficControllerRequest) *DeleteTrafficControllerInvoker {
	requestDef := GenReqDefForDeleteTrafficController()
	return &DeleteTrafficControllerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrafficController 修改信号机
//
// 修改信号机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateTrafficController(request *model.UpdateTrafficControllerRequest) (*model.UpdateTrafficControllerResponse, error) {
	requestDef := GenReqDefForUpdateTrafficController()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrafficControllerResponse), nil
	}
}

// UpdateTrafficControllerInvoker 修改信号机
func (c *DrisClient) UpdateTrafficControllerInvoker(request *model.UpdateTrafficControllerRequest) *UpdateTrafficControllerInvoker {
	requestDef := GenReqDefForUpdateTrafficController()
	return &UpdateTrafficControllerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowTrafficEvents 查询长期交通事件列表
//
// 条件查询交通事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowTrafficEvents(request *model.BatchShowTrafficEventsRequest) (*model.BatchShowTrafficEventsResponse, error) {
	requestDef := GenReqDefForBatchShowTrafficEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowTrafficEventsResponse), nil
	}
}

// BatchShowTrafficEventsInvoker 查询长期交通事件列表
func (c *DrisClient) BatchShowTrafficEventsInvoker(request *model.BatchShowTrafficEventsRequest) *BatchShowTrafficEventsInvoker {
	requestDef := GenReqDefForBatchShowTrafficEvents()
	return &BatchShowTrafficEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrafficEvent 创建长期交通事件
//
// 创建长期交通事件时，平台根据事件的起始时间和结束时间确定当前长期交通事件的状态。对于活跃状态的交通事件会立即下发给在事件影响范围内的RSU，对于未来事件则是在事件开始时间点下发到在事件影响范围内的RSU，过期事件不会下发。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateTrafficEvent(request *model.CreateTrafficEventRequest) (*model.CreateTrafficEventResponse, error) {
	requestDef := GenReqDefForCreateTrafficEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrafficEventResponse), nil
	}
}

// CreateTrafficEventInvoker 创建长期交通事件
func (c *DrisClient) CreateTrafficEventInvoker(request *model.CreateTrafficEventRequest) *CreateTrafficEventInvoker {
	requestDef := GenReqDefForCreateTrafficEvent()
	return &CreateTrafficEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrafficEvent 删除长期交通事件
//
// 刪除长期交通事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteTrafficEvent(request *model.DeleteTrafficEventRequest) (*model.DeleteTrafficEventResponse, error) {
	requestDef := GenReqDefForDeleteTrafficEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrafficEventResponse), nil
	}
}

// DeleteTrafficEventInvoker 删除长期交通事件
func (c *DrisClient) DeleteTrafficEventInvoker(request *model.DeleteTrafficEventRequest) *DeleteTrafficEventInvoker {
	requestDef := GenReqDefForDeleteTrafficEvent()
	return &DeleteTrafficEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrafficEvent 查询长期交通事件
//
// 查询长期交通事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowTrafficEvent(request *model.ShowTrafficEventRequest) (*model.ShowTrafficEventResponse, error) {
	requestDef := GenReqDefForShowTrafficEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrafficEventResponse), nil
	}
}

// ShowTrafficEventInvoker 查询长期交通事件
func (c *DrisClient) ShowTrafficEventInvoker(request *model.ShowTrafficEventRequest) *ShowTrafficEventInvoker {
	requestDef := GenReqDefForShowTrafficEvent()
	return &ShowTrafficEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrafficEvent 修改长期交通事件
//
// 修改长期交通事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateTrafficEvent(request *model.UpdateTrafficEventRequest) (*model.UpdateTrafficEventResponse, error) {
	requestDef := GenReqDefForUpdateTrafficEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrafficEventResponse), nil
	}
}

// UpdateTrafficEventInvoker 修改长期交通事件
func (c *DrisClient) UpdateTrafficEventInvoker(request *model.UpdateTrafficEventRequest) *UpdateTrafficEventInvoker {
	requestDef := GenReqDefForUpdateTrafficEvent()
	return &UpdateTrafficEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2xEdgeApp 部署边缘应用
//
// **部署边缘应用前需确保**：
//
// - Edge已创建且处于在线状态。相关方法请参见：[创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)，[查询Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0003.html)。
//
// - 待部署的应用已创建且应用版本状态已更新至发布。相关方法请参见：[创建应用](https://support.huaweicloud.com/api-v2x/v2x_04_0026.html)，[创建应用版本](https://support.huaweicloud.com/api-v2x/v2x_04_0038.html)，[更新应用版本状态](https://support.huaweicloud.com/api-v2x/v2x_04_0105.html)
//
// 如部署边缘应用接口调用成功，稍后将会自动安装至边缘设备无需手动操作。自动安装完成后应用将处于运行中的状态。
//
// **关于应用在设备侧部署的耗时问题**：
//
// &amp;emsp;&amp;emsp;从边缘应用部署成功到处于运行中状态的耗时取决于边缘设备所处的网络状况以及应用镜像包的大小，可通过查询边缘应用获取边缘应用部署状态。相关方法请参见：[查询边缘应用](https://support.huaweicloud.com/api-v2x/v2x_04_0115.html)\&quot;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateV2xEdgeApp(request *model.CreateV2xEdgeAppRequest) (*model.CreateV2xEdgeAppResponse, error) {
	requestDef := GenReqDefForCreateV2xEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2xEdgeAppResponse), nil
	}
}

// CreateV2xEdgeAppInvoker 部署边缘应用
func (c *DrisClient) CreateV2xEdgeAppInvoker(request *model.CreateV2xEdgeAppRequest) *CreateV2xEdgeAppInvoker {
	requestDef := GenReqDefForCreateV2xEdgeApp()
	return &CreateV2xEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2XEdgeAppByEdgeAppId 删除边缘应用
//
// 删除系统应用（$edgetepa）前应先删除业务通道。如删除边缘应用接口调用成功，稍后边缘设备将会自动删除应用无需手动操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteV2XEdgeAppByEdgeAppId(request *model.DeleteV2XEdgeAppByEdgeAppIdRequest) (*model.DeleteV2XEdgeAppByEdgeAppIdResponse, error) {
	requestDef := GenReqDefForDeleteV2XEdgeAppByEdgeAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2XEdgeAppByEdgeAppIdResponse), nil
	}
}

// DeleteV2XEdgeAppByEdgeAppIdInvoker 删除边缘应用
func (c *DrisClient) DeleteV2XEdgeAppByEdgeAppIdInvoker(request *model.DeleteV2XEdgeAppByEdgeAppIdRequest) *DeleteV2XEdgeAppByEdgeAppIdInvoker {
	requestDef := GenReqDefForDeleteV2XEdgeAppByEdgeAppId()
	return &DeleteV2XEdgeAppByEdgeAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2xEdgeApp 查询边缘应用列表
//
// 查询边缘应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ListV2xEdgeApp(request *model.ListV2xEdgeAppRequest) (*model.ListV2xEdgeAppResponse, error) {
	requestDef := GenReqDefForListV2xEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2xEdgeAppResponse), nil
	}
}

// ListV2xEdgeAppInvoker 查询边缘应用列表
func (c *DrisClient) ListV2xEdgeAppInvoker(request *model.ListV2xEdgeAppRequest) *ListV2xEdgeAppInvoker {
	requestDef := GenReqDefForListV2xEdgeApp()
	return &ListV2xEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2XEdgeAppDetailByEdgeAppId 查询边缘应用
//
// 查询边缘应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowV2XEdgeAppDetailByEdgeAppId(request *model.ShowV2XEdgeAppDetailByEdgeAppIdRequest) (*model.ShowV2XEdgeAppDetailByEdgeAppIdResponse, error) {
	requestDef := GenReqDefForShowV2XEdgeAppDetailByEdgeAppId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2XEdgeAppDetailByEdgeAppIdResponse), nil
	}
}

// ShowV2XEdgeAppDetailByEdgeAppIdInvoker 查询边缘应用
func (c *DrisClient) ShowV2XEdgeAppDetailByEdgeAppIdInvoker(request *model.ShowV2XEdgeAppDetailByEdgeAppIdRequest) *ShowV2XEdgeAppDetailByEdgeAppIdInvoker {
	requestDef := GenReqDefForShowV2XEdgeAppDetailByEdgeAppId()
	return &ShowV2XEdgeAppDetailByEdgeAppIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2xEdgeApp 升级边缘应用
//
// **升级边缘应用前需确保**：
//
// - Edge处于在线状态。相关方法请参见：[查询Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0003.html)。
//
// - 待升级的应用版本状态已更新至发布。相关方法请参见：[更新应用版本状态](https://support.huaweicloud.com/api-v2x/v2x_04_0105.html)
//
// 如升级边缘应用接口调用成功，稍后边缘设备将会自动升级至新版本无需手动操作。自动安装完成后应用将处于运行中的状态。
//
// **关于应用在设备侧升级的耗时问题**：
//
// &amp;emsp;&amp;emsp;从边缘应用升级成功到处于运行中状态的耗时取决于边缘设备所处的网络状况以及应用镜像包的大小，可通过查询边缘应用获取边缘应用部署状态。相关方法请参见：[查询边缘应用](https://support.huaweicloud.com/api-v2x/v2x_04_0115.html)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateV2xEdgeApp(request *model.UpdateV2xEdgeAppRequest) (*model.UpdateV2xEdgeAppResponse, error) {
	requestDef := GenReqDefForUpdateV2xEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2xEdgeAppResponse), nil
	}
}

// UpdateV2xEdgeAppInvoker 升级边缘应用
func (c *DrisClient) UpdateV2xEdgeAppInvoker(request *model.UpdateV2xEdgeAppRequest) *UpdateV2xEdgeAppInvoker {
	requestDef := GenReqDefForUpdateV2xEdgeApp()
	return &UpdateV2xEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowVehicles 查询车辆列表
//
// 查询车辆列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowVehicles(request *model.BatchShowVehiclesRequest) (*model.BatchShowVehiclesResponse, error) {
	requestDef := GenReqDefForBatchShowVehicles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowVehiclesResponse), nil
	}
}

// BatchShowVehiclesInvoker 查询车辆列表
func (c *DrisClient) BatchShowVehiclesInvoker(request *model.BatchShowVehiclesRequest) *BatchShowVehiclesInvoker {
	requestDef := GenReqDefForBatchShowVehicles()
	return &BatchShowVehiclesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVehicle 创建车辆
//
// 创建车辆
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateVehicle(request *model.CreateVehicleRequest) (*model.CreateVehicleResponse, error) {
	requestDef := GenReqDefForCreateVehicle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVehicleResponse), nil
	}
}

// CreateVehicleInvoker 创建车辆
func (c *DrisClient) CreateVehicleInvoker(request *model.CreateVehicleRequest) *CreateVehicleInvoker {
	requestDef := GenReqDefForCreateVehicle()
	return &CreateVehicleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVehicle 删除车辆
//
// 删除车辆
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteVehicle(request *model.DeleteVehicleRequest) (*model.DeleteVehicleResponse, error) {
	requestDef := GenReqDefForDeleteVehicle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVehicleResponse), nil
	}
}

// DeleteVehicleInvoker 删除车辆
func (c *DrisClient) DeleteVehicleInvoker(request *model.DeleteVehicleRequest) *DeleteVehicleInvoker {
	requestDef := GenReqDefForDeleteVehicle()
	return &DeleteVehicleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVehicle 修改车辆
//
// 修改车辆
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateVehicle(request *model.UpdateVehicleRequest) (*model.UpdateVehicleResponse, error) {
	requestDef := GenReqDefForUpdateVehicle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVehicleResponse), nil
	}
}

// UpdateVehicleInvoker 修改车辆
func (c *DrisClient) UpdateVehicleInvoker(request *model.UpdateVehicleRequest) *UpdateVehicleInvoker {
	requestDef := GenReqDefForUpdateVehicle()
	return &UpdateVehicleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowEdgeApps 查询应用列表
//
// 查询应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowEdgeApps(request *model.BatchShowEdgeAppsRequest) (*model.BatchShowEdgeAppsResponse, error) {
	requestDef := GenReqDefForBatchShowEdgeApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowEdgeAppsResponse), nil
	}
}

// BatchShowEdgeAppsInvoker 查询应用列表
func (c *DrisClient) BatchShowEdgeAppsInvoker(request *model.BatchShowEdgeAppsRequest) *BatchShowEdgeAppsInvoker {
	requestDef := GenReqDefForBatchShowEdgeApps()
	return &BatchShowEdgeAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateEdgeApp(request *model.CreateEdgeAppRequest) (*model.CreateEdgeAppResponse, error) {
	requestDef := GenReqDefForCreateEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeAppResponse), nil
	}
}

// CreateEdgeAppInvoker 创建应用
func (c *DrisClient) CreateEdgeAppInvoker(request *model.CreateEdgeAppRequest) *CreateEdgeAppInvoker {
	requestDef := GenReqDefForCreateEdgeApp()
	return &CreateEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeApp 删除应用
//
// 删除应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteEdgeApp(request *model.DeleteEdgeAppRequest) (*model.DeleteEdgeAppResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeAppResponse), nil
	}
}

// DeleteEdgeAppInvoker 删除应用
func (c *DrisClient) DeleteEdgeAppInvoker(request *model.DeleteEdgeAppRequest) *DeleteEdgeAppInvoker {
	requestDef := GenReqDefForDeleteEdgeApp()
	return &DeleteEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeApp 修改应用
//
// 修改应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateEdgeApp(request *model.UpdateEdgeAppRequest) (*model.UpdateEdgeAppResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeAppResponse), nil
	}
}

// UpdateEdgeAppInvoker 修改应用
func (c *DrisClient) UpdateEdgeAppInvoker(request *model.UpdateEdgeAppRequest) *UpdateEdgeAppInvoker {
	requestDef := GenReqDefForUpdateEdgeApp()
	return &UpdateEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowEdgeAppVersions 查询应用版本列表
//
// 查询应用版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) BatchShowEdgeAppVersions(request *model.BatchShowEdgeAppVersionsRequest) (*model.BatchShowEdgeAppVersionsResponse, error) {
	requestDef := GenReqDefForBatchShowEdgeAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowEdgeAppVersionsResponse), nil
	}
}

// BatchShowEdgeAppVersionsInvoker 查询应用版本列表
func (c *DrisClient) BatchShowEdgeAppVersionsInvoker(request *model.BatchShowEdgeAppVersionsRequest) *BatchShowEdgeAppVersionsInvoker {
	requestDef := GenReqDefForBatchShowEdgeAppVersions()
	return &BatchShowEdgeAppVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeApplicationVersion 创建应用版本
//
// 创建应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateEdgeApplicationVersion(request *model.CreateEdgeApplicationVersionRequest) (*model.CreateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForCreateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeApplicationVersionResponse), nil
	}
}

// CreateEdgeApplicationVersionInvoker 创建应用版本
func (c *DrisClient) CreateEdgeApplicationVersionInvoker(request *model.CreateEdgeApplicationVersionRequest) *CreateEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForCreateEdgeApplicationVersion()
	return &CreateEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeApplicationVersion 删除应用版本
//
// 删除应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteEdgeApplicationVersion(request *model.DeleteEdgeApplicationVersionRequest) (*model.DeleteEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeApplicationVersionResponse), nil
	}
}

// DeleteEdgeApplicationVersionInvoker 删除应用版本
func (c *DrisClient) DeleteEdgeApplicationVersionInvoker(request *model.DeleteEdgeApplicationVersionRequest) *DeleteEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForDeleteEdgeApplicationVersion()
	return &DeleteEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeApplicationVersion 查询应用版本
//
// 查询应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowEdgeApplicationVersion(request *model.ShowEdgeApplicationVersionRequest) (*model.ShowEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForShowEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeApplicationVersionResponse), nil
	}
}

// ShowEdgeApplicationVersionInvoker 查询应用版本
func (c *DrisClient) ShowEdgeApplicationVersionInvoker(request *model.ShowEdgeApplicationVersionRequest) *ShowEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForShowEdgeApplicationVersion()
	return &ShowEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeApplicationVersion 修改应用版本
//
// 修改应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateEdgeApplicationVersion(request *model.UpdateEdgeApplicationVersionRequest) (*model.UpdateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionResponse), nil
	}
}

// UpdateEdgeApplicationVersionInvoker 修改应用版本
func (c *DrisClient) UpdateEdgeApplicationVersionInvoker(request *model.UpdateEdgeApplicationVersionRequest) *UpdateEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForUpdateEdgeApplicationVersion()
	return &UpdateEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeApplicationVersionState 更新应用版本状态
//
// 更新应用版本状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateEdgeApplicationVersionState(request *model.UpdateEdgeApplicationVersionStateRequest) (*model.UpdateEdgeApplicationVersionStateResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersionState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionStateResponse), nil
	}
}

// UpdateEdgeApplicationVersionStateInvoker 更新应用版本状态
func (c *DrisClient) UpdateEdgeApplicationVersionStateInvoker(request *model.UpdateEdgeApplicationVersionStateRequest) *UpdateEdgeApplicationVersionStateInvoker {
	requestDef := GenReqDefForUpdateEdgeApplicationVersionState()
	return &UpdateEdgeApplicationVersionStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRsuModel 创建RSU型号
//
// 调用此接口可创建RSU型号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) CreateRsuModel(request *model.CreateRsuModelRequest) (*model.CreateRsuModelResponse, error) {
	requestDef := GenReqDefForCreateRsuModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRsuModelResponse), nil
	}
}

// CreateRsuModelInvoker 创建RSU型号
func (c *DrisClient) CreateRsuModelInvoker(request *model.CreateRsuModelRequest) *CreateRsuModelInvoker {
	requestDef := GenReqDefForCreateRsuModel()
	return &CreateRsuModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRsuModel 删除RSU型号
//
// 可调用此接口删除已创建的RSU型号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) DeleteRsuModel(request *model.DeleteRsuModelRequest) (*model.DeleteRsuModelResponse, error) {
	requestDef := GenReqDefForDeleteRsuModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRsuModelResponse), nil
	}
}

// DeleteRsuModelInvoker 删除RSU型号
func (c *DrisClient) DeleteRsuModelInvoker(request *model.DeleteRsuModelRequest) *DeleteRsuModelInvoker {
	requestDef := GenReqDefForDeleteRsuModel()
	return &DeleteRsuModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRsuModels 查询RSU型号列表
//
// 可调用此接口查询已创建RSU型号列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ListRsuModels(request *model.ListRsuModelsRequest) (*model.ListRsuModelsResponse, error) {
	requestDef := GenReqDefForListRsuModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRsuModelsResponse), nil
	}
}

// ListRsuModelsInvoker 查询RSU型号列表
func (c *DrisClient) ListRsuModelsInvoker(request *model.ListRsuModelsRequest) *ListRsuModelsInvoker {
	requestDef := GenReqDefForListRsuModels()
	return &ListRsuModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRsuModel 查询RSU型号
//
// 可调用此接口查询已创建的RSU型号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) ShowRsuModel(request *model.ShowRsuModelRequest) (*model.ShowRsuModelResponse, error) {
	requestDef := GenReqDefForShowRsuModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRsuModelResponse), nil
	}
}

// ShowRsuModelInvoker 查询RSU型号
func (c *DrisClient) ShowRsuModelInvoker(request *model.ShowRsuModelRequest) *ShowRsuModelInvoker {
	requestDef := GenReqDefForShowRsuModel()
	return &ShowRsuModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRsuModel 修改RSU型号
//
// 可调用此接口修改已创建的RSU型号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrisClient) UpdateRsuModel(request *model.UpdateRsuModelRequest) (*model.UpdateRsuModelResponse, error) {
	requestDef := GenReqDefForUpdateRsuModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRsuModelResponse), nil
	}
}

// UpdateRsuModelInvoker 修改RSU型号
func (c *DrisClient) UpdateRsuModelInvoker(request *model.UpdateRsuModelRequest) *UpdateRsuModelInvoker {
	requestDef := GenReqDefForUpdateRsuModel()
	return &UpdateRsuModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
