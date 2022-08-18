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

// CreateChannel 创建自定义事件通道
//
// 创建自定义事件通道。
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

// CreateEventSource 创建自定义事件源
//
// 创建用户自定义类型的事件源，只能指定自定义通道，不能指定官方通道。
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
// 创建事件订阅。
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
// 在事件订阅中增加一个事件目标。
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
// 删除指定自定义事件通道。
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

// DeleteEventSource 删除自定义事件源
//
// 删除指定的自定义事件源。
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
// 删除事件订阅。
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
// 删除事件订阅中的事件目标。
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

// ListChannels 查询事件通道列表
//
// 查询事件通道列表，包括官方通道和自定义通道。
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

// ListEventSources 查询事件源列表
//
// 查询事件源列表，支持条件查询，如可以指定事件通道ID来查询某个事件通道下的所有事件源。
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
// 查询预置的事件目标分类，获取每个事件目标分类的字段定义。
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
// 查询当前租户的配额，未特殊配置过的会返回默认配额。
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
// 查询事件订阅列表，支持指定事件通道。
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

// OperateSubscription 启用/禁用事件订阅
//
// 启用/禁用事件订阅。
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

// OperateSubscriptionInvoker 启用/禁用事件订阅
func (c *EgClient) OperateSubscriptionInvoker(request *model.OperateSubscriptionRequest) *OperateSubscriptionInvoker {
	requestDef := GenReqDefForOperateSubscription()
	return &OperateSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutEvents 发布事件到事件通道
//
// 发布事件到事件通道。
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
// 查询事件通道详情。
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

// ShowDetailOfEventSource 查询事件源详情
//
// 查询事件源详情信息。
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
// 查询事件订阅详情。
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
// 查询事件订阅中事件目标的详情。
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
// 修改自定义事件通道的描述信息。
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

// UpdateEventSource 更新自定义事件源
//
// 修改自定义事件源的描述信息。
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
// 更新事件订阅描述信息、事件源参数或者事件目标参数。
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
// 更新事件订阅中事件源的参数。
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
// 更新事件订阅中事件目标的参数。
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
