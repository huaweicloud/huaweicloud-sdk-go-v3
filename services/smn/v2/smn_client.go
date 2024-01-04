package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smn/v2/model"
)

type SmnClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSmnClient(hcClient *http_client.HcHttpClient) *SmnClient {
	return &SmnClient{HcClient: hcClient}
}

func SmnClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddSubscription 订阅
//
// 为指定Topic添加一个订阅者，如果订阅者的状态为未确认，则向订阅者发送一个确认的消息。待订阅者进行ConfirmSubscription确认后，该订阅者才能收到Topic发布的消息。单Topic默认可添加10000个订阅者，高并发场景下，可能会出现订阅者数量超过10000仍添加成功的情况，此为正常现象。接口是幂等的，如果添加已存在的订阅者，则返回成功，且status code为200，否则status code为201。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) AddSubscription(request *model.AddSubscriptionRequest) (*model.AddSubscriptionResponse, error) {
	requestDef := GenReqDefForAddSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubscriptionResponse), nil
	}
}

// AddSubscriptionInvoker 订阅
func (c *SmnClient) AddSubscriptionInvoker(request *model.AddSubscriptionRequest) *AddSubscriptionInvoker {
	requestDef := GenReqDefForAddSubscription()
	return &AddSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSubscriptionFromSubscriptionUser 导入订阅
//
// 为指定的Topic添加订阅者，订阅者信息来源为订阅用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) AddSubscriptionFromSubscriptionUser(request *model.AddSubscriptionFromSubscriptionUserRequest) (*model.AddSubscriptionFromSubscriptionUserResponse, error) {
	requestDef := GenReqDefForAddSubscriptionFromSubscriptionUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubscriptionFromSubscriptionUserResponse), nil
	}
}

// AddSubscriptionFromSubscriptionUserInvoker 导入订阅
func (c *SmnClient) AddSubscriptionFromSubscriptionUserInvoker(request *model.AddSubscriptionFromSubscriptionUserRequest) *AddSubscriptionFromSubscriptionUserInvoker {
	requestDef := GenReqDefForAddSubscriptionFromSubscriptionUser()
	return &AddSubscriptionFromSubscriptionUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateOrDeleteResourceTags 批量添加删除资源标签
//
// 为指定实例批量添加或删除标签。一个资源上最多有10个标签。
// 此接口为幂等接口：创建时如果请求体中存在重复key则报错。
// 创建时，不允许重复key，如果数据库存在就覆盖。
// 删除时，如果删除的标签不存在，默认处理成功，删除时不对标签字符集范围做校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) BatchCreateOrDeleteResourceTags(request *model.BatchCreateOrDeleteResourceTagsRequest) (*model.BatchCreateOrDeleteResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteResourceTagsResponse), nil
	}
}

// BatchCreateOrDeleteResourceTagsInvoker 批量添加删除资源标签
func (c *SmnClient) BatchCreateOrDeleteResourceTagsInvoker(request *model.BatchCreateOrDeleteResourceTagsRequest) *BatchCreateOrDeleteResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteResourceTags()
	return &BatchCreateOrDeleteResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSubscriptionsFilterPolices 批量创建订阅过滤策略
//
// 创建订阅者的消息过滤策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) BatchCreateSubscriptionsFilterPolices(request *model.BatchCreateSubscriptionsFilterPolicesRequest) (*model.BatchCreateSubscriptionsFilterPolicesResponse, error) {
	requestDef := GenReqDefForBatchCreateSubscriptionsFilterPolices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSubscriptionsFilterPolicesResponse), nil
	}
}

// BatchCreateSubscriptionsFilterPolicesInvoker 批量创建订阅过滤策略
func (c *SmnClient) BatchCreateSubscriptionsFilterPolicesInvoker(request *model.BatchCreateSubscriptionsFilterPolicesRequest) *BatchCreateSubscriptionsFilterPolicesInvoker {
	requestDef := GenReqDefForBatchCreateSubscriptionsFilterPolices()
	return &BatchCreateSubscriptionsFilterPolicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSubscriptionsFilterPolices 批量删除订阅过滤策略
//
// 删除订阅者的消息过滤策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) BatchDeleteSubscriptionsFilterPolices(request *model.BatchDeleteSubscriptionsFilterPolicesRequest) (*model.BatchDeleteSubscriptionsFilterPolicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteSubscriptionsFilterPolices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSubscriptionsFilterPolicesResponse), nil
	}
}

// BatchDeleteSubscriptionsFilterPolicesInvoker 批量删除订阅过滤策略
func (c *SmnClient) BatchDeleteSubscriptionsFilterPolicesInvoker(request *model.BatchDeleteSubscriptionsFilterPolicesRequest) *BatchDeleteSubscriptionsFilterPolicesInvoker {
	requestDef := GenReqDefForBatchDeleteSubscriptionsFilterPolices()
	return &BatchDeleteSubscriptionsFilterPolicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateSubscriptionsFilterPolices 批量更新订阅过滤策略
//
// 更新订阅者的消息过滤策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) BatchUpdateSubscriptionsFilterPolices(request *model.BatchUpdateSubscriptionsFilterPolicesRequest) (*model.BatchUpdateSubscriptionsFilterPolicesResponse, error) {
	requestDef := GenReqDefForBatchUpdateSubscriptionsFilterPolices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateSubscriptionsFilterPolicesResponse), nil
	}
}

// BatchUpdateSubscriptionsFilterPolicesInvoker 批量更新订阅过滤策略
func (c *SmnClient) BatchUpdateSubscriptionsFilterPolicesInvoker(request *model.BatchUpdateSubscriptionsFilterPolicesRequest) *BatchUpdateSubscriptionsFilterPolicesInvoker {
	requestDef := GenReqDefForBatchUpdateSubscriptionsFilterPolices()
	return &BatchUpdateSubscriptionsFilterPolicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSubscription 取消订阅
//
// 删除指定的订阅者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CancelSubscription(request *model.CancelSubscriptionRequest) (*model.CancelSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSubscriptionResponse), nil
	}
}

// CancelSubscriptionInvoker 取消订阅
func (c *SmnClient) CancelSubscriptionInvoker(request *model.CancelSubscriptionRequest) *CancelSubscriptionInvoker {
	requestDef := GenReqDefForCancelSubscription()
	return &CancelSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogtank 绑定云日志
//
// 为指定Topic绑定一个云日志，用于记录主题消息发送状态等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateLogtank(request *model.CreateLogtankRequest) (*model.CreateLogtankResponse, error) {
	requestDef := GenReqDefForCreateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogtankResponse), nil
	}
}

// CreateLogtankInvoker 绑定云日志
func (c *SmnClient) CreateLogtankInvoker(request *model.CreateLogtankRequest) *CreateLogtankInvoker {
	requestDef := GenReqDefForCreateLogtank()
	return &CreateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMessageTemplate 创建消息模板
//
// 创建一个模板，用户可以按照模板去发送消息，这样可以减少请求的数据量。
// 单用户默认可创建100个消息模板，高并发场景下，可能会出现消息模板数量超过100仍创建成功的情况，此为正常现象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateMessageTemplate(request *model.CreateMessageTemplateRequest) (*model.CreateMessageTemplateResponse, error) {
	requestDef := GenReqDefForCreateMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessageTemplateResponse), nil
	}
}

// CreateMessageTemplateInvoker 创建消息模板
func (c *SmnClient) CreateMessageTemplateInvoker(request *model.CreateMessageTemplateRequest) *CreateMessageTemplateInvoker {
	requestDef := GenReqDefForCreateMessageTemplate()
	return &CreateMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceTag 添加资源标签
//
// 一个资源上最多有10个标签。此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

// CreateResourceTagInvoker 添加资源标签
func (c *SmnClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTopic 创建主题
//
// 创建Topic，单用户默认配额为3000。高并发场景下，可能会出现Topic数量超过3000仍创建成功的情况，此为正常现象。
// 接口是幂等的，接口调用返回成功时，若已存在同名的Topic，返回的status code为200，否则返回的status code为201
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateTopic(request *model.CreateTopicRequest) (*model.CreateTopicResponse, error) {
	requestDef := GenReqDefForCreateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTopicResponse), nil
	}
}

// CreateTopicInvoker 创建主题
func (c *SmnClient) CreateTopicInvoker(request *model.CreateTopicRequest) *CreateTopicInvoker {
	requestDef := GenReqDefForCreateTopic()
	return &CreateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogtank 解绑云日志
//
// 解绑指定Topic绑定的云日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteLogtank(request *model.DeleteLogtankRequest) (*model.DeleteLogtankResponse, error) {
	requestDef := GenReqDefForDeleteLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogtankResponse), nil
	}
}

// DeleteLogtankInvoker 解绑云日志
func (c *SmnClient) DeleteLogtankInvoker(request *model.DeleteLogtankRequest) *DeleteLogtankInvoker {
	requestDef := GenReqDefForDeleteLogtank()
	return &DeleteLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMessageTemplate 删除消息模板
//
// 删除消息模板。删除模板之前的消息请求都可以使用该模板发送，删除之后无法再使用该模板发送消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteMessageTemplate(request *model.DeleteMessageTemplateRequest) (*model.DeleteMessageTemplateResponse, error) {
	requestDef := GenReqDefForDeleteMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMessageTemplateResponse), nil
	}
}

// DeleteMessageTemplateInvoker 删除消息模板
func (c *SmnClient) DeleteMessageTemplateInvoker(request *model.DeleteMessageTemplateRequest) *DeleteMessageTemplateInvoker {
	requestDef := GenReqDefForDeleteMessageTemplate()
	return &DeleteMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 删除资源标签
//
// 幂等接口：删除时，不对标签做校验。删除的key不存在报404，key不能为空或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 删除资源标签
func (c *SmnClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTopic 删除主题
//
// 删除主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteTopic(request *model.DeleteTopicRequest) (*model.DeleteTopicResponse, error) {
	requestDef := GenReqDefForDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicResponse), nil
	}
}

// DeleteTopicInvoker 删除主题
func (c *SmnClient) DeleteTopicInvoker(request *model.DeleteTopicRequest) *DeleteTopicInvoker {
	requestDef := GenReqDefForDeleteTopic()
	return &DeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTopicAttributeByName 删除指定名称的主题策略
//
// 删除指定名称的主题策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteTopicAttributeByName(request *model.DeleteTopicAttributeByNameRequest) (*model.DeleteTopicAttributeByNameResponse, error) {
	requestDef := GenReqDefForDeleteTopicAttributeByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicAttributeByNameResponse), nil
	}
}

// DeleteTopicAttributeByNameInvoker 删除指定名称的主题策略
func (c *SmnClient) DeleteTopicAttributeByNameInvoker(request *model.DeleteTopicAttributeByNameRequest) *DeleteTopicAttributeByNameInvoker {
	requestDef := GenReqDefForDeleteTopicAttributeByName()
	return &DeleteTopicAttributeByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTopicAttributes 删除所有主题策略
//
// 删除所有主题策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteTopicAttributes(request *model.DeleteTopicAttributesRequest) (*model.DeleteTopicAttributesResponse, error) {
	requestDef := GenReqDefForDeleteTopicAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicAttributesResponse), nil
	}
}

// DeleteTopicAttributesInvoker 删除所有主题策略
func (c *SmnClient) DeleteTopicAttributesInvoker(request *model.DeleteTopicAttributesRequest) *DeleteTopicAttributesInvoker {
	requestDef := GenReqDefForDeleteTopicAttributes()
	return &DeleteTopicAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogtank 查询云日志
//
// 查询指定Topic绑定的云日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListLogtank(request *model.ListLogtankRequest) (*model.ListLogtankResponse, error) {
	requestDef := GenReqDefForListLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogtankResponse), nil
	}
}

// ListLogtankInvoker 查询云日志
func (c *SmnClient) ListLogtankInvoker(request *model.ListLogtankRequest) *ListLogtankInvoker {
	requestDef := GenReqDefForListLogtank()
	return &ListLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessageTemplateDetails 查询消息模板详情
//
// 查询模板详情，包括模板内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListMessageTemplateDetails(request *model.ListMessageTemplateDetailsRequest) (*model.ListMessageTemplateDetailsResponse, error) {
	requestDef := GenReqDefForListMessageTemplateDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTemplateDetailsResponse), nil
	}
}

// ListMessageTemplateDetailsInvoker 查询消息模板详情
func (c *SmnClient) ListMessageTemplateDetailsInvoker(request *model.ListMessageTemplateDetailsRequest) *ListMessageTemplateDetailsInvoker {
	requestDef := GenReqDefForListMessageTemplateDetails()
	return &ListMessageTemplateDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessageTemplates 查询消息模板列表
//
// 分页查询模板列表，模板列表按照创建时间进行升序排列。分页查询可以指定offset以及limit。如果不存在模板，则返回空列表。额外的查询参数分别有message_template_name和protocol。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListMessageTemplates(request *model.ListMessageTemplatesRequest) (*model.ListMessageTemplatesResponse, error) {
	requestDef := GenReqDefForListMessageTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTemplatesResponse), nil
	}
}

// ListMessageTemplatesInvoker 查询消息模板列表
func (c *SmnClient) ListMessageTemplatesInvoker(request *model.ListMessageTemplatesRequest) *ListMessageTemplatesInvoker {
	requestDef := GenReqDefForListMessageTemplates()
	return &ListMessageTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询租户在指定Region和实例类型的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *SmnClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 查询资源实例
//
// 使用标签过滤实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 查询资源实例
func (c *SmnClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询资源标签
func (c *SmnClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubscriptions 查询订阅者列表
//
// 分页返回请求者的所有的订阅列表，订阅列表按照订阅创建时间进行升序排列。分页查询可以指定offset以及limit。如果订阅者不存在，返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListSubscriptions(request *model.ListSubscriptionsRequest) (*model.ListSubscriptionsResponse, error) {
	requestDef := GenReqDefForListSubscriptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionsResponse), nil
	}
}

// ListSubscriptionsInvoker 查询订阅者列表
func (c *SmnClient) ListSubscriptionsInvoker(request *model.ListSubscriptionsRequest) *ListSubscriptionsInvoker {
	requestDef := GenReqDefForListSubscriptions()
	return &ListSubscriptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubscriptionsByTopic 查询指定Topic的订阅者列表
//
// 分页获取特定Topic的订阅列表，订阅列表按照订阅创建时间进行升序排列。分页查询可以指定offset以及limit。如果指定Topic不存在订阅者，返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListSubscriptionsByTopic(request *model.ListSubscriptionsByTopicRequest) (*model.ListSubscriptionsByTopicResponse, error) {
	requestDef := GenReqDefForListSubscriptionsByTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionsByTopicResponse), nil
	}
}

// ListSubscriptionsByTopicInvoker 查询指定Topic的订阅者列表
func (c *SmnClient) ListSubscriptionsByTopicInvoker(request *model.ListSubscriptionsByTopicRequest) *ListSubscriptionsByTopicInvoker {
	requestDef := GenReqDefForListSubscriptionsByTopic()
	return &ListSubscriptionsByTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopicAttributes 查询主题策略
//
// 查询主题的策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListTopicAttributes(request *model.ListTopicAttributesRequest) (*model.ListTopicAttributesResponse, error) {
	requestDef := GenReqDefForListTopicAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicAttributesResponse), nil
	}
}

// ListTopicAttributesInvoker 查询主题策略
func (c *SmnClient) ListTopicAttributesInvoker(request *model.ListTopicAttributesRequest) *ListTopicAttributesInvoker {
	requestDef := GenReqDefForListTopicAttributes()
	return &ListTopicAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopicDetails 查询主题详情
//
// 查询Topic的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListTopicDetails(request *model.ListTopicDetailsRequest) (*model.ListTopicDetailsResponse, error) {
	requestDef := GenReqDefForListTopicDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicDetailsResponse), nil
	}
}

// ListTopicDetailsInvoker 查询主题详情
func (c *SmnClient) ListTopicDetailsInvoker(request *model.ListTopicDetailsRequest) *ListTopicDetailsInvoker {
	requestDef := GenReqDefForListTopicDetails()
	return &ListTopicDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopics 查询主题列表
//
// 分页查询Topic列表，Topic列表按照Topic创建时间进行降序排列。分页查询可以指定offset以及limit。如果不存在Topic，则返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListTopics(request *model.ListTopicsRequest) (*model.ListTopicsResponse, error) {
	requestDef := GenReqDefForListTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicsResponse), nil
	}
}

// ListTopicsInvoker 查询主题列表
func (c *SmnClient) ListTopicsInvoker(request *model.ListTopicsRequest) *ListTopicsInvoker {
	requestDef := GenReqDefForListTopics()
	return &ListTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersion 查询SMN API V2版本信息
//
// 查询SMN API V2版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListVersion(request *model.ListVersionRequest) (*model.ListVersionResponse, error) {
	requestDef := GenReqDefForListVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionResponse), nil
	}
}

// ListVersionInvoker 查询SMN API V2版本信息
func (c *SmnClient) ListVersionInvoker(request *model.ListVersionRequest) *ListVersionInvoker {
	requestDef := GenReqDefForListVersion()
	return &ListVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersions 查询SMN支持的API版本号信息
//
// 查询SMN开放API支持的版本号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

// ListVersionsInvoker 查询SMN支持的API版本号信息
func (c *SmnClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishHttpDetect 发布探测消息
//
// 基于主题发送http/https探测消息，探测当前http/https 终端是否可用，SMN出口是否能够正常访问该终端。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) PublishHttpDetect(request *model.PublishHttpDetectRequest) (*model.PublishHttpDetectResponse, error) {
	requestDef := GenReqDefForPublishHttpDetect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishHttpDetectResponse), nil
	}
}

// PublishHttpDetectInvoker 发布探测消息
func (c *SmnClient) PublishHttpDetectInvoker(request *model.PublishHttpDetectRequest) *PublishHttpDetectInvoker {
	requestDef := GenReqDefForPublishHttpDetect()
	return &PublishHttpDetectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishMessage 消息发布
//
// 将消息发送给Topic的所有订阅端点。当返回消息ID时，该消息已被保存并开始尝试将其推送给Topic的订阅者。三种消息发送方式
//
// message
//
// message_structure
//
// message_template_name
//
// 只需要设置其中一个，如果同时设置，生效的优先级为
// message_structure &gt; message_template_name &gt; message。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) PublishMessage(request *model.PublishMessageRequest) (*model.PublishMessageResponse, error) {
	requestDef := GenReqDefForPublishMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishMessageResponse), nil
	}
}

// PublishMessageInvoker 消息发布
func (c *SmnClient) PublishMessageInvoker(request *model.PublishMessageRequest) *PublishMessageInvoker {
	requestDef := GenReqDefForPublishMessage()
	return &PublishMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpDetectResult 获取http探测结果
//
// 根据http探测发送返回的task_id查询探测结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ShowHttpDetectResult(request *model.ShowHttpDetectResultRequest) (*model.ShowHttpDetectResultResponse, error) {
	requestDef := GenReqDefForShowHttpDetectResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpDetectResultResponse), nil
	}
}

// ShowHttpDetectResultInvoker 获取http探测结果
func (c *SmnClient) ShowHttpDetectResultInvoker(request *model.ShowHttpDetectResultRequest) *ShowHttpDetectResultInvoker {
	requestDef := GenReqDefForShowHttpDetectResult()
	return &ShowHttpDetectResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogtank 更新云日志
//
// 更新指定Topic绑定的云日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateLogtank(request *model.UpdateLogtankRequest) (*model.UpdateLogtankResponse, error) {
	requestDef := GenReqDefForUpdateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogtankResponse), nil
	}
}

// UpdateLogtankInvoker 更新云日志
func (c *SmnClient) UpdateLogtankInvoker(request *model.UpdateLogtankRequest) *UpdateLogtankInvoker {
	requestDef := GenReqDefForUpdateLogtank()
	return &UpdateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMessageTemplate 更新消息模板
//
// 修改消息模板的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateMessageTemplate(request *model.UpdateMessageTemplateRequest) (*model.UpdateMessageTemplateResponse, error) {
	requestDef := GenReqDefForUpdateMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMessageTemplateResponse), nil
	}
}

// UpdateMessageTemplateInvoker 更新消息模板
func (c *SmnClient) UpdateMessageTemplateInvoker(request *model.UpdateMessageTemplateRequest) *UpdateMessageTemplateInvoker {
	requestDef := GenReqDefForUpdateMessageTemplate()
	return &UpdateMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubscription 更新订阅者
//
// 更新订阅者备注。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateSubscription(request *model.UpdateSubscriptionRequest) (*model.UpdateSubscriptionResponse, error) {
	requestDef := GenReqDefForUpdateSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubscriptionResponse), nil
	}
}

// UpdateSubscriptionInvoker 更新订阅者
func (c *SmnClient) UpdateSubscriptionInvoker(request *model.UpdateSubscriptionRequest) *UpdateSubscriptionInvoker {
	requestDef := GenReqDefForUpdateSubscription()
	return &UpdateSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopic 更新主题
//
// 更新显示名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateTopic(request *model.UpdateTopicRequest) (*model.UpdateTopicResponse, error) {
	requestDef := GenReqDefForUpdateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicResponse), nil
	}
}

// UpdateTopicInvoker 更新主题
func (c *SmnClient) UpdateTopicInvoker(request *model.UpdateTopicRequest) *UpdateTopicInvoker {
	requestDef := GenReqDefForUpdateTopic()
	return &UpdateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopicAttribute 更新主题策略
//
// 更新主题的策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateTopicAttribute(request *model.UpdateTopicAttributeRequest) (*model.UpdateTopicAttributeResponse, error) {
	requestDef := GenReqDefForUpdateTopicAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAttributeResponse), nil
	}
}

// UpdateTopicAttributeInvoker 更新主题策略
func (c *SmnClient) UpdateTopicAttributeInvoker(request *model.UpdateTopicAttributeRequest) *UpdateTopicAttributeInvoker {
	requestDef := GenReqDefForUpdateTopicAttribute()
	return &UpdateTopicAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateApplication 创建Application
//
// 创建平台应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateApplicationInvoker 创建Application
func (c *SmnClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteApplication 删除Application
//
// 删除平台应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteApplicationInvoker 删除Application
func (c *SmnClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationAttributes 查询Application属性
//
// 获取应用平台属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListApplicationAttributes(request *model.ListApplicationAttributesRequest) (*model.ListApplicationAttributesResponse, error) {
	requestDef := GenReqDefForListApplicationAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationAttributesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationAttributesInvoker 查询Application属性
func (c *SmnClient) ListApplicationAttributesInvoker(request *model.ListApplicationAttributesRequest) *ListApplicationAttributesInvoker {
	requestDef := GenReqDefForListApplicationAttributes()
	return &ListApplicationAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplications 查询Application
//
// 查询应用平台列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationsInvoker 查询Application
func (c *SmnClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PublishAppMessage App消息发布
//
// 将消息直发给endpoint设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) PublishAppMessage(request *model.PublishAppMessageRequest) (*model.PublishAppMessageResponse, error) {
	requestDef := GenReqDefForPublishAppMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishAppMessageResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PublishAppMessageInvoker App消息发布
func (c *SmnClient) PublishAppMessageInvoker(request *model.PublishAppMessageRequest) *PublishAppMessageInvoker {
	requestDef := GenReqDefForPublishAppMessage()
	return &PublishAppMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateApplication 更新Application
//
// 更新应用平台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateApplicationInvoker 更新Application
func (c *SmnClient) UpdateApplicationInvoker(request *model.UpdateApplicationRequest) *UpdateApplicationInvoker {
	requestDef := GenReqDefForUpdateApplication()
	return &UpdateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateApplicationEndpoint 创建Application endpoint
//
// 创建应用平台的endpoint终端。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) CreateApplicationEndpoint(request *model.CreateApplicationEndpointRequest) (*model.CreateApplicationEndpointResponse, error) {
	requestDef := GenReqDefForCreateApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationEndpointResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateApplicationEndpointInvoker 创建Application endpoint
func (c *SmnClient) CreateApplicationEndpointInvoker(request *model.CreateApplicationEndpointRequest) *CreateApplicationEndpointInvoker {
	requestDef := GenReqDefForCreateApplicationEndpoint()
	return &CreateApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteApplicationEndpoint 删除Application endpoint
//
// 删除设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) DeleteApplicationEndpoint(request *model.DeleteApplicationEndpointRequest) (*model.DeleteApplicationEndpointResponse, error) {
	requestDef := GenReqDefForDeleteApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationEndpointResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteApplicationEndpointInvoker 删除Application endpoint
func (c *SmnClient) DeleteApplicationEndpointInvoker(request *model.DeleteApplicationEndpointRequest) *DeleteApplicationEndpointInvoker {
	requestDef := GenReqDefForDeleteApplicationEndpoint()
	return &DeleteApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationEndpointAttributes 查询Application的Endpoint属性
//
// 获取endpoint的属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListApplicationEndpointAttributes(request *model.ListApplicationEndpointAttributesRequest) (*model.ListApplicationEndpointAttributesResponse, error) {
	requestDef := GenReqDefForListApplicationEndpointAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationEndpointAttributesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationEndpointAttributesInvoker 查询Application的Endpoint属性
func (c *SmnClient) ListApplicationEndpointAttributesInvoker(request *model.ListApplicationEndpointAttributesRequest) *ListApplicationEndpointAttributesInvoker {
	requestDef := GenReqDefForListApplicationEndpointAttributes()
	return &ListApplicationEndpointAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationEndpoints 查询Application的Endpoint列表
//
// 查询平台的endpoint列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) ListApplicationEndpoints(request *model.ListApplicationEndpointsRequest) (*model.ListApplicationEndpointsResponse, error) {
	requestDef := GenReqDefForListApplicationEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationEndpointsResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListApplicationEndpointsInvoker 查询Application的Endpoint列表
func (c *SmnClient) ListApplicationEndpointsInvoker(request *model.ListApplicationEndpointsRequest) *ListApplicationEndpointsInvoker {
	requestDef := GenReqDefForListApplicationEndpoints()
	return &ListApplicationEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateApplicationEndpoint 更新Application endpoint
//
// 更新设备属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnClient) UpdateApplicationEndpoint(request *model.UpdateApplicationEndpointRequest) (*model.UpdateApplicationEndpointResponse, error) {
	requestDef := GenReqDefForUpdateApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationEndpointResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateApplicationEndpointInvoker 更新Application endpoint
func (c *SmnClient) UpdateApplicationEndpointInvoker(request *model.UpdateApplicationEndpointRequest) *UpdateApplicationEndpointInvoker {
	requestDef := GenReqDefForUpdateApplicationEndpoint()
	return &UpdateApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
