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

// CreateAgencies 创建服务委托
//
// 按照业务场景，一键创建服务委托授权
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 创建自定义事件通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 创建目标连接
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// create endpoint
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateEventSchema 创建自定义事件模型
//
// 创建自定义事件模型
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) CreateEventSchema(request *model.CreateEventSchemaRequest) (*model.CreateEventSchemaResponse, error) {
	requestDef := GenReqDefForCreateEventSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventSchemaResponse), nil
	}
}

// CreateEventSchemaInvoker 创建自定义事件模型
func (c *EgClient) CreateEventSchemaInvoker(request *model.CreateEventSchemaRequest) *CreateEventSchemaInvoker {
	requestDef := GenReqDefForCreateEventSchema()
	return &CreateEventSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEventSchemaVersion 创建自定义事件模型版本
//
// 创建自定义事件模型版本，版本号后台自动生成
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) CreateEventSchemaVersion(request *model.CreateEventSchemaVersionRequest) (*model.CreateEventSchemaVersionResponse, error) {
	requestDef := GenReqDefForCreateEventSchemaVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventSchemaVersionResponse), nil
	}
}

// CreateEventSchemaVersionInvoker 创建自定义事件模型版本
func (c *EgClient) CreateEventSchemaVersionInvoker(request *model.CreateEventSchemaVersionRequest) *CreateEventSchemaVersionInvoker {
	requestDef := GenReqDefForCreateEventSchemaVersion()
	return &CreateEventSchemaVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEventSource 创建自定义事件源
//
// 创建用户自定义类型的事件源，只能指定自定义通道，不能指定系统通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateSubscription 创建事件订阅
//
// 创建事件订阅
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 创建单个事件订阅目标
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 删除指定自定义事件通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 删除目标连接
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteEventSchema 删除事件模型
//
// 删除事件模型
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) DeleteEventSchema(request *model.DeleteEventSchemaRequest) (*model.DeleteEventSchemaResponse, error) {
	requestDef := GenReqDefForDeleteEventSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventSchemaResponse), nil
	}
}

// DeleteEventSchemaInvoker 删除事件模型
func (c *EgClient) DeleteEventSchemaInvoker(request *model.DeleteEventSchemaRequest) *DeleteEventSchemaInvoker {
	requestDef := GenReqDefForDeleteEventSchema()
	return &DeleteEventSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEventSchemaVersion 删除事件模型版本
//
// 删除事件模型指定版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) DeleteEventSchemaVersion(request *model.DeleteEventSchemaVersionRequest) (*model.DeleteEventSchemaVersionResponse, error) {
	requestDef := GenReqDefForDeleteEventSchemaVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventSchemaVersionResponse), nil
	}
}

// DeleteEventSchemaVersionInvoker 删除事件模型版本
func (c *EgClient) DeleteEventSchemaVersionInvoker(request *model.DeleteEventSchemaVersionRequest) *DeleteEventSchemaVersionInvoker {
	requestDef := GenReqDefForDeleteEventSchemaVersion()
	return &DeleteEventSchemaVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEventSource 删除自定义事件源
//
// 删除指定的自定义事件源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteSubscription 删除事件订阅
//
// 删除事件订阅
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 删除事件订阅目标
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DiscoverEventSchemaFromData 事件模型自动发现
//
// 事件模型自动发现
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) DiscoverEventSchemaFromData(request *model.DiscoverEventSchemaFromDataRequest) (*model.DiscoverEventSchemaFromDataResponse, error) {
	requestDef := GenReqDefForDiscoverEventSchemaFromData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DiscoverEventSchemaFromDataResponse), nil
	}
}

// DiscoverEventSchemaFromDataInvoker 事件模型自动发现
func (c *EgClient) DiscoverEventSchemaFromDataInvoker(request *model.DiscoverEventSchemaFromDataRequest) *DiscoverEventSchemaFromDataInvoker {
	requestDef := GenReqDefForDiscoverEventSchemaFromData()
	return &DiscoverEventSchemaFromDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgencies 查询服务委托
//
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 查询事件通道列表，包括系统通道和自定义通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 查询目标连接列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// list all endpoints
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListEventSchema 查询事件模型列表
//
// 查询事件模型列表，包括系统事件模型和自定义事件模型
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) ListEventSchema(request *model.ListEventSchemaRequest) (*model.ListEventSchemaResponse, error) {
	requestDef := GenReqDefForListEventSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventSchemaResponse), nil
	}
}

// ListEventSchemaInvoker 查询事件模型列表
func (c *EgClient) ListEventSchemaInvoker(request *model.ListEventSchemaRequest) *ListEventSchemaInvoker {
	requestDef := GenReqDefForListEventSchema()
	return &ListEventSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventSchemaVersions 查询事件模型版本列表
//
// 查询事件模型版本列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) ListEventSchemaVersions(request *model.ListEventSchemaVersionsRequest) (*model.ListEventSchemaVersionsResponse, error) {
	requestDef := GenReqDefForListEventSchemaVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventSchemaVersionsResponse), nil
	}
}

// ListEventSchemaVersionsInvoker 查询事件模型版本列表
func (c *EgClient) ListEventSchemaVersionsInvoker(request *model.ListEventSchemaVersionsRequest) *ListEventSchemaVersionsInvoker {
	requestDef := GenReqDefForListEventSchemaVersions()
	return &ListEventSchemaVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventSources 查询事件源列表
//
// 支持条件查询，如可以指定事件通道ID来查询某个事件通道下的所有事件源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListEventTarget 查询事件目标分类
//
// 查询预置的事件目标分类，获取每个事件目标分类的字段定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListQuotas 查询配额
//
// 查询当前租户的配额，未特殊配置过的会返回默认配额
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListSubscriptions 查询事件订阅列表
//
// 查询事件订阅列表，支持指定事件通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListTriggers 查询事件订阅列表
//
// 查询触发器，支持指定函数urn。一个以函数urn为目标的订阅为一个触发器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) ListTriggers(request *model.ListTriggersRequest) (*model.ListTriggersResponse, error) {
	requestDef := GenReqDefForListTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTriggersResponse), nil
	}
}

// ListTriggersInvoker 查询事件订阅列表
func (c *EgClient) ListTriggersInvoker(request *model.ListTriggersRequest) *ListTriggersInvoker {
	requestDef := GenReqDefForListTriggers()
	return &ListTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OperateSubscription 操作事件订阅
//
// 操作事件订阅，支持启用、禁用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 发布事件到事件通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowDetailOfChannel 查询事件通道详情
//
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowDetailOfEventSchema 查询事件模型详情
//
// 查询事件模型详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) ShowDetailOfEventSchema(request *model.ShowDetailOfEventSchemaRequest) (*model.ShowDetailOfEventSchemaResponse, error) {
	requestDef := GenReqDefForShowDetailOfEventSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfEventSchemaResponse), nil
	}
}

// ShowDetailOfEventSchemaInvoker 查询事件模型详情
func (c *EgClient) ShowDetailOfEventSchemaInvoker(request *model.ShowDetailOfEventSchemaRequest) *ShowDetailOfEventSchemaInvoker {
	requestDef := GenReqDefForShowDetailOfEventSchema()
	return &ShowDetailOfEventSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfEventSchemaVersion 查询事件模型版本详情
//
// 查询事件模型指定版本详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) ShowDetailOfEventSchemaVersion(request *model.ShowDetailOfEventSchemaVersionRequest) (*model.ShowDetailOfEventSchemaVersionResponse, error) {
	requestDef := GenReqDefForShowDetailOfEventSchemaVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailOfEventSchemaVersionResponse), nil
	}
}

// ShowDetailOfEventSchemaVersionInvoker 查询事件模型版本详情
func (c *EgClient) ShowDetailOfEventSchemaVersionInvoker(request *model.ShowDetailOfEventSchemaVersionRequest) *ShowDetailOfEventSchemaVersionInvoker {
	requestDef := GenReqDefForShowDetailOfEventSchemaVersion()
	return &ShowDetailOfEventSchemaVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailOfEventSource 查询事件源详情
//
// 查询事件源详情信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowDetailOfSubscription 查询事件订阅详情
//
// 查询事件订阅详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 查询事件订阅目标详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateChannel 更新自定义事件通道
//
// 更新自定义事件通道定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 更新目标连接
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// update endpoint
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateEventSchema 更新自定义事件模型
//
// 更新自定义事件模型定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EgClient) UpdateEventSchema(request *model.UpdateEventSchemaRequest) (*model.UpdateEventSchemaResponse, error) {
	requestDef := GenReqDefForUpdateEventSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventSchemaResponse), nil
	}
}

// UpdateEventSchemaInvoker 更新自定义事件模型
func (c *EgClient) UpdateEventSchemaInvoker(request *model.UpdateEventSchemaRequest) *UpdateEventSchemaInvoker {
	requestDef := GenReqDefForUpdateEventSchema()
	return &UpdateEventSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEventSource 更新自定义事件源
//
// 更新自定义事件源定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateSubscription 更新事件订阅
//
// 更新事件订阅定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 更新事件订阅目标定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 获取服务支持的API版本列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
