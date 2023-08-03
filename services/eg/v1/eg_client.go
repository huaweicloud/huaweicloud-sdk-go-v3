package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v1/model"
)

type EgClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEgClient(hcClient *http_client.HcHttpClient) *EgClient {
	return &EgClient{HcClient: hcClient}
}

func EgClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CheckPutEvents 预校验指定事件源发布事件成功
//
// 发布事件到事件源成功需要有订阅等条件，预先校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CheckPutEvents(request *model.CheckPutEventsRequest) (*model.CheckPutEventsResponse, error) {
	requestDef := GenReqDefForCheckPutEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPutEventsResponse), nil
	}
}

// CheckPutEventsInvoker 预校验指定事件源发布事件成功
func (c *EgClient) CheckPutEventsInvoker(request *model.CheckPutEventsRequest) *CheckPutEventsInvoker {
	requestDef := GenReqDefForCheckPutEvents()
	return &CheckPutEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgencies 创建服务委托
//
// 按照业务场景，一键创建服务委托授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateAgencies(request *model.CreateAgenciesRequest) (*model.CreateAgenciesResponse, error) {
	requestDef := GenReqDefForCreateAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgenciesResponse), nil
	}
}

// CreateAgenciesInvoker 创建服务委托
func (c *EgClient) CreateAgenciesInvoker(request *model.CreateAgenciesRequest) *CreateAgenciesInvoker {
	requestDef := GenReqDefForCreateAgencies()
	return &CreateAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateChannel 创建自定义事件通道
//
// 创建自定义事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateChannel(request *model.CreateChannelRequest) (*model.CreateChannelResponse, error) {
	requestDef := GenReqDefForCreateChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateChannelResponse), nil
	}
}

// CreateChannelInvoker 创建自定义事件通道
func (c *EgClient) CreateChannelInvoker(request *model.CreateChannelRequest) *CreateChannelInvoker {
	requestDef := GenReqDefForCreateChannel()
	return &CreateChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnection 创建目标连接
//
// 创建目标连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

// CreateConnectionInvoker 创建目标连接
func (c *EgClient) CreateConnectionInvoker(request *model.CreateConnectionRequest) *CreateConnectionInvoker {
	requestDef := GenReqDefForCreateConnection()
	return &CreateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建访问端点
//
// 创建访问端点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建访问端点
func (c *EgClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEventSource 创建自定义事件源
//
// 创建用户自定义类型的事件源，只能指定自定义通道，不能指定系统通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateEventSource(request *model.CreateEventSourceRequest) (*model.CreateEventSourceResponse, error) {
	requestDef := GenReqDefForCreateEventSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventSourceResponse), nil
	}
}

// CreateEventSourceInvoker 创建自定义事件源
func (c *EgClient) CreateEventSourceInvoker(request *model.CreateEventSourceRequest) *CreateEventSourceInvoker {
	requestDef := GenReqDefForCreateEventSource()
	return &CreateEventSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEventStreaming 创建事件流
//
// 创建事件流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateEventStreaming(request *model.CreateEventStreamingRequest) (*model.CreateEventStreamingResponse, error) {
	requestDef := GenReqDefForCreateEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventStreamingResponse), nil
	}
}

// CreateEventStreamingInvoker 创建事件流
func (c *EgClient) CreateEventStreamingInvoker(request *model.CreateEventStreamingRequest) *CreateEventStreamingInvoker {
	requestDef := GenReqDefForCreateEventStreaming()
	return &CreateEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubscription 创建事件订阅
//
// 创建事件订阅。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateSubscription(request *model.CreateSubscriptionRequest) (*model.CreateSubscriptionResponse, error) {
	requestDef := GenReqDefForCreateSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubscriptionResponse), nil
	}
}

// CreateSubscriptionInvoker 创建事件订阅
func (c *EgClient) CreateSubscriptionInvoker(request *model.CreateSubscriptionRequest) *CreateSubscriptionInvoker {
	requestDef := GenReqDefForCreateSubscription()
	return &CreateSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubscriptionTarget 创建事件订阅目标
//
// 创建单个事件订阅目标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) CreateSubscriptionTarget(request *model.CreateSubscriptionTargetRequest) (*model.CreateSubscriptionTargetResponse, error) {
	requestDef := GenReqDefForCreateSubscriptionTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubscriptionTargetResponse), nil
	}
}

// CreateSubscriptionTargetInvoker 创建事件订阅目标
func (c *EgClient) CreateSubscriptionTargetInvoker(request *model.CreateSubscriptionTargetRequest) *CreateSubscriptionTargetInvoker {
	requestDef := GenReqDefForCreateSubscriptionTarget()
	return &CreateSubscriptionTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteChannel 删除自定义事件通道
//
// 删除指定自定义事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteChannel(request *model.DeleteChannelRequest) (*model.DeleteChannelResponse, error) {
	requestDef := GenReqDefForDeleteChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteChannelResponse), nil
	}
}

// DeleteChannelInvoker 删除自定义事件通道
func (c *EgClient) DeleteChannelInvoker(request *model.DeleteChannelRequest) *DeleteChannelInvoker {
	requestDef := GenReqDefForDeleteChannel()
	return &DeleteChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnection 删除目标连接
//
// 删除目标连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteConnection(request *model.DeleteConnectionRequest) (*model.DeleteConnectionResponse, error) {
	requestDef := GenReqDefForDeleteConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionResponse), nil
	}
}

// DeleteConnectionInvoker 删除目标连接
func (c *EgClient) DeleteConnectionInvoker(request *model.DeleteConnectionRequest) *DeleteConnectionInvoker {
	requestDef := GenReqDefForDeleteConnection()
	return &DeleteConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpoint 删除访问端点
//
// 删除访问端点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

// DeleteEndpointInvoker 删除访问端点
func (c *EgClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEventSource 删除自定义事件源
//
// 删除指定的自定义事件源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteEventSource(request *model.DeleteEventSourceRequest) (*model.DeleteEventSourceResponse, error) {
	requestDef := GenReqDefForDeleteEventSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventSourceResponse), nil
	}
}

// DeleteEventSourceInvoker 删除自定义事件源
func (c *EgClient) DeleteEventSourceInvoker(request *model.DeleteEventSourceRequest) *DeleteEventSourceInvoker {
	requestDef := GenReqDefForDeleteEventSource()
	return &DeleteEventSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEventStreaming 删除事件流
//
// 删除事件流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteEventStreaming(request *model.DeleteEventStreamingRequest) (*model.DeleteEventStreamingResponse, error) {
	requestDef := GenReqDefForDeleteEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventStreamingResponse), nil
	}
}

// DeleteEventStreamingInvoker 删除事件流
func (c *EgClient) DeleteEventStreamingInvoker(request *model.DeleteEventStreamingRequest) *DeleteEventStreamingInvoker {
	requestDef := GenReqDefForDeleteEventStreaming()
	return &DeleteEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubscription 删除事件订阅
//
// 删除事件订阅。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteSubscription(request *model.DeleteSubscriptionRequest) (*model.DeleteSubscriptionResponse, error) {
	requestDef := GenReqDefForDeleteSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubscriptionResponse), nil
	}
}

// DeleteSubscriptionInvoker 删除事件订阅
func (c *EgClient) DeleteSubscriptionInvoker(request *model.DeleteSubscriptionRequest) *DeleteSubscriptionInvoker {
	requestDef := GenReqDefForDeleteSubscription()
	return &DeleteSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubscriptionTarget 删除事件订阅目标
//
// 删除事件订阅目标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) DeleteSubscriptionTarget(request *model.DeleteSubscriptionTargetRequest) (*model.DeleteSubscriptionTargetResponse, error) {
	requestDef := GenReqDefForDeleteSubscriptionTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubscriptionTargetResponse), nil
	}
}

// DeleteSubscriptionTargetInvoker 删除事件订阅目标
func (c *EgClient) DeleteSubscriptionTargetInvoker(request *model.DeleteSubscriptionTargetRequest) *DeleteSubscriptionTargetInvoker {
	requestDef := GenReqDefForDeleteSubscriptionTarget()
	return &DeleteSubscriptionTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgencies 查询服务委托
//
// 查询服务委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesResponse), nil
	}
}

// ListAgenciesInvoker 查询服务委托
func (c *EgClient) ListAgenciesInvoker(request *model.ListAgenciesRequest) *ListAgenciesInvoker {
	requestDef := GenReqDefForListAgencies()
	return &ListAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListChannels 查询事件通道列表
//
// 查询事件通道列表，包括系统通道和自定义通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListChannels(request *model.ListChannelsRequest) (*model.ListChannelsResponse, error) {
	requestDef := GenReqDefForListChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChannelsResponse), nil
	}
}

// ListChannelsInvoker 查询事件通道列表
func (c *EgClient) ListChannelsInvoker(request *model.ListChannelsRequest) *ListChannelsInvoker {
	requestDef := GenReqDefForListChannels()
	return &ListChannelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 查询目标连接列表
//
// 查询目标连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 查询目标连接列表
func (c *EgClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询访问端点
//
// 查询访问端点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询访问端点
func (c *EgClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventSources 查询事件源列表
//
// 支持条件查询，如可以指定事件通道ID来查询某个事件通道下的所有事件源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListEventSources(request *model.ListEventSourcesRequest) (*model.ListEventSourcesResponse, error) {
	requestDef := GenReqDefForListEventSources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventSourcesResponse), nil
	}
}

// ListEventSourcesInvoker 查询事件源列表
func (c *EgClient) ListEventSourcesInvoker(request *model.ListEventSourcesRequest) *ListEventSourcesInvoker {
	requestDef := GenReqDefForListEventSources()
	return &ListEventSourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventStreaming 查询事件流列表
//
// 查询事件流列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListEventStreaming(request *model.ListEventStreamingRequest) (*model.ListEventStreamingResponse, error) {
	requestDef := GenReqDefForListEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventStreamingResponse), nil
	}
}

// ListEventStreamingInvoker 查询事件流列表
func (c *EgClient) ListEventStreamingInvoker(request *model.ListEventStreamingRequest) *ListEventStreamingInvoker {
	requestDef := GenReqDefForListEventStreaming()
	return &ListEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventTarget 查询事件目标分类
//
// 查询预置的事件目标分类，获取每个事件目标分类的字段定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListEventTarget(request *model.ListEventTargetRequest) (*model.ListEventTargetResponse, error) {
	requestDef := GenReqDefForListEventTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventTargetResponse), nil
	}
}

// ListEventTargetInvoker 查询事件目标分类
func (c *EgClient) ListEventTargetInvoker(request *model.ListEventTargetRequest) *ListEventTargetInvoker {
	requestDef := GenReqDefForListEventTarget()
	return &ListEventTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPubMetrics 查询事件通道监控指标数据
//
// 查询事件通道监控指标数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListPubMetrics(request *model.ListPubMetricsRequest) (*model.ListPubMetricsResponse, error) {
	requestDef := GenReqDefForListPubMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPubMetricsResponse), nil
	}
}

// ListPubMetricsInvoker 查询事件通道监控指标数据
func (c *EgClient) ListPubMetricsInvoker(request *model.ListPubMetricsRequest) *ListPubMetricsInvoker {
	requestDef := GenReqDefForListPubMetrics()
	return &ListPubMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询当前租户的配额，未特殊配置过的会返回默认配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *EgClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubMetrics 查询事件订阅监控指标数据
//
// 查询事件订阅监控指标数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListSubMetrics(request *model.ListSubMetricsRequest) (*model.ListSubMetricsResponse, error) {
	requestDef := GenReqDefForListSubMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubMetricsResponse), nil
	}
}

// ListSubMetricsInvoker 查询事件订阅监控指标数据
func (c *EgClient) ListSubMetricsInvoker(request *model.ListSubMetricsRequest) *ListSubMetricsInvoker {
	requestDef := GenReqDefForListSubMetrics()
	return &ListSubMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubscriptions 查询事件订阅列表
//
// 查询事件订阅列表，支持指定事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListSubscriptions(request *model.ListSubscriptionsRequest) (*model.ListSubscriptionsResponse, error) {
	requestDef := GenReqDefForListSubscriptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionsResponse), nil
	}
}

// ListSubscriptionsInvoker 查询事件订阅列表
func (c *EgClient) ListSubscriptionsInvoker(request *model.ListSubscriptionsRequest) *ListSubscriptionsInvoker {
	requestDef := GenReqDefForListSubscriptions()
	return &ListSubscriptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTracedEvents 查询事件追踪列表
//
// 查询事件追踪列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListTracedEvents(request *model.ListTracedEventsRequest) (*model.ListTracedEventsResponse, error) {
	requestDef := GenReqDefForListTracedEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTracedEventsResponse), nil
	}
}

// ListTracedEventsInvoker 查询事件追踪列表
func (c *EgClient) ListTracedEventsInvoker(request *model.ListTracedEventsRequest) *ListTracedEventsInvoker {
	requestDef := GenReqDefForListTracedEvents()
	return &ListTracedEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTriggers 查询单个函数的EG触发器
//
// 查询触发器，支持指定函数urn，一个以函数urn为目标的订阅为一个触发器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListTriggers(request *model.ListTriggersRequest) (*model.ListTriggersResponse, error) {
	requestDef := GenReqDefForListTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTriggersResponse), nil
	}
}

// ListTriggersInvoker 查询单个函数的EG触发器
func (c *EgClient) ListTriggersInvoker(request *model.ListTriggersRequest) *ListTriggersInvoker {
	requestDef := GenReqDefForListTriggers()
	return &ListTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowTriggers 查询单个函数流的EG触发器
//
// 查询触发器，支持指定函数流id，一个以函数流id为目标的订阅为一个触发器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListWorkflowTriggers(request *model.ListWorkflowTriggersRequest) (*model.ListWorkflowTriggersResponse, error) {
	requestDef := GenReqDefForListWorkflowTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowTriggersResponse), nil
	}
}

// ListWorkflowTriggersInvoker 查询单个函数流的EG触发器
func (c *EgClient) ListWorkflowTriggersInvoker(request *model.ListWorkflowTriggersRequest) *ListWorkflowTriggersInvoker {
	requestDef := GenReqDefForListWorkflowTriggers()
	return &ListWorkflowTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OperateSubscription 操作事件订阅
//
// 操作事件订阅，支持启用、禁用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) OperateSubscription(request *model.OperateSubscriptionRequest) (*model.OperateSubscriptionResponse, error) {
	requestDef := GenReqDefForOperateSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.OperateSubscriptionResponse), nil
	}
}

// OperateSubscriptionInvoker 操作事件订阅
func (c *EgClient) OperateSubscriptionInvoker(request *model.OperateSubscriptionRequest) *OperateSubscriptionInvoker {
	requestDef := GenReqDefForOperateSubscription()
	return &OperateSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutEvents 发布事件到事件通道
//
// 发布事件到事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) PutEvents(request *model.PutEventsRequest) (*model.PutEventsResponse, error) {
	requestDef := GenReqDefForPutEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutEventsResponse), nil
	}
}

// PutEventsInvoker 发布事件到事件通道
func (c *EgClient) PutEventsInvoker(request *model.PutEventsRequest) *PutEventsInvoker {
	requestDef := GenReqDefForPutEvents()
	return &PutEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutOfficialEvents 发布官方事件到事件通道
//
// 发布官方事件到事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) PutOfficialEvents(request *model.PutOfficialEventsRequest) (*model.PutOfficialEventsResponse, error) {
	requestDef := GenReqDefForPutOfficialEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutOfficialEventsResponse), nil
	}
}

// PutOfficialEventsInvoker 发布官方事件到事件通道
func (c *EgClient) PutOfficialEventsInvoker(request *model.PutOfficialEventsRequest) *PutOfficialEventsInvoker {
	requestDef := GenReqDefForPutOfficialEvents()
	return &PutOfficialEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeEventStreaming 操作事件流
//
// 操作事件流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ResumeEventStreaming(request *model.ResumeEventStreamingRequest) (*model.ResumeEventStreamingResponse, error) {
	requestDef := GenReqDefForResumeEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeEventStreamingResponse), nil
	}
}

// ResumeEventStreamingInvoker 操作事件流
func (c *EgClient) ResumeEventStreamingInvoker(request *model.ResumeEventStreamingRequest) *ResumeEventStreamingInvoker {
	requestDef := GenReqDefForResumeEventStreaming()
	return &ResumeEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfChannel 查询事件通道详情
//
// 查询指定事件通道详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfChannel(request *model.ShowDetailOfChannelRequest) (*model.ShowDetailOfChannelResponse, error) {
	requestDef := GenReqDefForShowDetailOfChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfChannelResponse), nil
	}
}

// ShowDetailOfChannelInvoker 查询事件通道详情
func (c *EgClient) ShowDetailOfChannelInvoker(request *model.ShowDetailOfChannelRequest) *ShowDetailOfChannelInvoker {
	requestDef := GenReqDefForShowDetailOfChannel()
	return &ShowDetailOfChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfConnection 查询目标连接详情
//
// 查询目标连接详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfConnection(request *model.ShowDetailOfConnectionRequest) (*model.ShowDetailOfConnectionResponse, error) {
	requestDef := GenReqDefForShowDetailOfConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfConnectionResponse), nil
	}
}

// ShowDetailOfConnectionInvoker 查询目标连接详情
func (c *EgClient) ShowDetailOfConnectionInvoker(request *model.ShowDetailOfConnectionRequest) *ShowDetailOfConnectionInvoker {
	requestDef := GenReqDefForShowDetailOfConnection()
	return &ShowDetailOfConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfEvent 查询发送事件的内容
//
// 根据事件ID查询事件详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfEvent(request *model.ShowDetailOfEventRequest) (*model.ShowDetailOfEventResponse, error) {
	requestDef := GenReqDefForShowDetailOfEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfEventResponse), nil
	}
}

// ShowDetailOfEventInvoker 查询发送事件的内容
func (c *EgClient) ShowDetailOfEventInvoker(request *model.ShowDetailOfEventRequest) *ShowDetailOfEventInvoker {
	requestDef := GenReqDefForShowDetailOfEvent()
	return &ShowDetailOfEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfEventSource 查询事件源详情
//
// 查询事件源详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfEventSource(request *model.ShowDetailOfEventSourceRequest) (*model.ShowDetailOfEventSourceResponse, error) {
	requestDef := GenReqDefForShowDetailOfEventSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfEventSourceResponse), nil
	}
}

// ShowDetailOfEventSourceInvoker 查询事件源详情
func (c *EgClient) ShowDetailOfEventSourceInvoker(request *model.ShowDetailOfEventSourceRequest) *ShowDetailOfEventSourceInvoker {
	requestDef := GenReqDefForShowDetailOfEventSource()
	return &ShowDetailOfEventSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfEventTrace 事件轨迹详情
//
// 事件轨迹详情，展示事件源到投递目标的投递情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfEventTrace(request *model.ShowDetailOfEventTraceRequest) (*model.ShowDetailOfEventTraceResponse, error) {
	requestDef := GenReqDefForShowDetailOfEventTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfEventTraceResponse), nil
	}
}

// ShowDetailOfEventTraceInvoker 事件轨迹详情
func (c *EgClient) ShowDetailOfEventTraceInvoker(request *model.ShowDetailOfEventTraceRequest) *ShowDetailOfEventTraceInvoker {
	requestDef := GenReqDefForShowDetailOfEventTrace()
	return &ShowDetailOfEventTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfSubscription 查询事件订阅详情
//
// 查询事件订阅详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfSubscription(request *model.ShowDetailOfSubscriptionRequest) (*model.ShowDetailOfSubscriptionResponse, error) {
	requestDef := GenReqDefForShowDetailOfSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfSubscriptionResponse), nil
	}
}

// ShowDetailOfSubscriptionInvoker 查询事件订阅详情
func (c *EgClient) ShowDetailOfSubscriptionInvoker(request *model.ShowDetailOfSubscriptionRequest) *ShowDetailOfSubscriptionInvoker {
	requestDef := GenReqDefForShowDetailOfSubscription()
	return &ShowDetailOfSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfSubscriptionTarget 查询事件订阅目标详情
//
// 查询事件订阅目标详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowDetailOfSubscriptionTarget(request *model.ShowDetailOfSubscriptionTargetRequest) (*model.ShowDetailOfSubscriptionTargetResponse, error) {
	requestDef := GenReqDefForShowDetailOfSubscriptionTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfSubscriptionTargetResponse), nil
	}
}

// ShowDetailOfSubscriptionTargetInvoker 查询事件订阅目标详情
func (c *EgClient) ShowDetailOfSubscriptionTargetInvoker(request *model.ShowDetailOfSubscriptionTargetRequest) *ShowDetailOfSubscriptionTargetInvoker {
	requestDef := GenReqDefForShowDetailOfSubscriptionTarget()
	return &ShowDetailOfSubscriptionTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEventStreaming 查询事件流详情
//
// 查询事件流详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ShowEventStreaming(request *model.ShowEventStreamingRequest) (*model.ShowEventStreamingResponse, error) {
	requestDef := GenReqDefForShowEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventStreamingResponse), nil
	}
}

// ShowEventStreamingInvoker 查询事件流详情
func (c *EgClient) ShowEventStreamingInvoker(request *model.ShowEventStreamingRequest) *ShowEventStreamingInvoker {
	requestDef := GenReqDefForShowEventStreaming()
	return &ShowEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateChannel 更新自定义事件通道
//
// 更新自定义事件通道定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateChannel(request *model.UpdateChannelRequest) (*model.UpdateChannelResponse, error) {
	requestDef := GenReqDefForUpdateChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateChannelResponse), nil
	}
}

// UpdateChannelInvoker 更新自定义事件通道
func (c *EgClient) UpdateChannelInvoker(request *model.UpdateChannelRequest) *UpdateChannelInvoker {
	requestDef := GenReqDefForUpdateChannel()
	return &UpdateChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnection 更新目标连接
//
// 更新目标连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateConnection(request *model.UpdateConnectionRequest) (*model.UpdateConnectionResponse, error) {
	requestDef := GenReqDefForUpdateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionResponse), nil
	}
}

// UpdateConnectionInvoker 更新目标连接
func (c *EgClient) UpdateConnectionInvoker(request *model.UpdateConnectionRequest) *UpdateConnectionInvoker {
	requestDef := GenReqDefForUpdateConnection()
	return &UpdateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpoint 更新访问端点
//
// 更新访问端点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateEndpoint(request *model.UpdateEndpointRequest) (*model.UpdateEndpointResponse, error) {
	requestDef := GenReqDefForUpdateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointResponse), nil
	}
}

// UpdateEndpointInvoker 更新访问端点
func (c *EgClient) UpdateEndpointInvoker(request *model.UpdateEndpointRequest) *UpdateEndpointInvoker {
	requestDef := GenReqDefForUpdateEndpoint()
	return &UpdateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEventSource 更新自定义事件源
//
// 更新自定义事件源定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateEventSource(request *model.UpdateEventSourceRequest) (*model.UpdateEventSourceResponse, error) {
	requestDef := GenReqDefForUpdateEventSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventSourceResponse), nil
	}
}

// UpdateEventSourceInvoker 更新自定义事件源
func (c *EgClient) UpdateEventSourceInvoker(request *model.UpdateEventSourceRequest) *UpdateEventSourceInvoker {
	requestDef := GenReqDefForUpdateEventSource()
	return &UpdateEventSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEventStreaming 更新事件流
//
// 更新事件流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateEventStreaming(request *model.UpdateEventStreamingRequest) (*model.UpdateEventStreamingResponse, error) {
	requestDef := GenReqDefForUpdateEventStreaming()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventStreamingResponse), nil
	}
}

// UpdateEventStreamingInvoker 更新事件流
func (c *EgClient) UpdateEventStreamingInvoker(request *model.UpdateEventStreamingRequest) *UpdateEventStreamingInvoker {
	requestDef := GenReqDefForUpdateEventStreaming()
	return &UpdateEventStreamingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubscription 更新事件订阅
//
// 更新事件订阅定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateSubscription(request *model.UpdateSubscriptionRequest) (*model.UpdateSubscriptionResponse, error) {
	requestDef := GenReqDefForUpdateSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubscriptionResponse), nil
	}
}

// UpdateSubscriptionInvoker 更新事件订阅
func (c *EgClient) UpdateSubscriptionInvoker(request *model.UpdateSubscriptionRequest) *UpdateSubscriptionInvoker {
	requestDef := GenReqDefForUpdateSubscription()
	return &UpdateSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubscriptionSource 更新事件订阅源
//
// 更新事件订阅源定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateSubscriptionSource(request *model.UpdateSubscriptionSourceRequest) (*model.UpdateSubscriptionSourceResponse, error) {
	requestDef := GenReqDefForUpdateSubscriptionSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubscriptionSourceResponse), nil
	}
}

// UpdateSubscriptionSourceInvoker 更新事件订阅源
func (c *EgClient) UpdateSubscriptionSourceInvoker(request *model.UpdateSubscriptionSourceRequest) *UpdateSubscriptionSourceInvoker {
	requestDef := GenReqDefForUpdateSubscriptionSource()
	return &UpdateSubscriptionSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubscriptionTarget 更新事件订阅目标
//
// 更新事件订阅目标定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) UpdateSubscriptionTarget(request *model.UpdateSubscriptionTargetRequest) (*model.UpdateSubscriptionTargetResponse, error) {
	requestDef := GenReqDefForUpdateSubscriptionTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubscriptionTargetResponse), nil
	}
}

// UpdateSubscriptionTargetInvoker 更新事件订阅目标
func (c *EgClient) UpdateSubscriptionTargetInvoker(request *model.UpdateSubscriptionTargetRequest) *UpdateSubscriptionTargetInvoker {
	requestDef := GenReqDefForUpdateSubscriptionTarget()
	return &UpdateSubscriptionTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 获取API版本列表
//
// 获取服务支持的API版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 获取API版本列表
func (c *EgClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBuckets 获取obs桶列表
//
// 获取obs桶列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) ListObsBuckets(request *model.ListObsBucketsRequest) (*model.ListObsBucketsResponse, error) {
	requestDef := GenReqDefForListObsBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketsResponse), nil
	}
}

// ListObsBucketsInvoker 获取obs桶列表
func (c *EgClient) ListObsBucketsInvoker(request *model.ListObsBucketsRequest) *ListObsBucketsInvoker {
	requestDef := GenReqDefForListObsBuckets()
	return &ListObsBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
