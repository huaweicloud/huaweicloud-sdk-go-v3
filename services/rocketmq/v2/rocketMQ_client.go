package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rocketmq/v2/model"
)

type RocketMQClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRocketMQClient(hcClient *http_client.HcHttpClient) *RocketMQClient {
	return &RocketMQClient{HcClient: hcClient}
}

func RocketMQClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 批量删除实例
//
// 批量删除实例。**实例删除后，实例中原有的数据将被删除，且没有备份，请谨慎操。**
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

// 批量修改消费组
//
// 批量修改消费组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) BatchUpdateConsumerGroup(request *model.BatchUpdateConsumerGroupRequest) (*model.BatchUpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForBatchUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateConsumerGroupResponse), nil
	}
}

// 创建消费组或批量删除消费组
//
// 创建消费组或批量删除消费组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) CreateConsumerGroupOrBatchDeleteConsumerGroup(request *model.CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) (*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForCreateConsumerGroupOrBatchDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse), nil
	}
}

// 创建实例（按需）
//
// 创建实例，该接口创建的实例为按需计费的方式。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

// 创建用户
//
// 创建用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// 删除指定消费组
//
// 删除指定消费组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) DeleteConsumerGroup(request *model.DeleteConsumerGroupRequest) (*model.DeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConsumerGroupResponse), nil
	}
}

// 删除指定的实例
//
// 删除指定的实例，释放该实例的所有资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// 删除用户
//
// 删除用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// 导出死信消息
//
// 导出死信消息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ExportDlqMessage(request *model.ExportDlqMessageRequest) (*model.ExportDlqMessageResponse, error) {
	requestDef := GenReqDefForExportDlqMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDlqMessageResponse), nil
	}
}

// 查询可用区信息
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// 查询代理列表
//
// 查询代理列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListBrokers(request *model.ListBrokersRequest) (*model.ListBrokersResponse, error) {
	requestDef := GenReqDefForListBrokers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBrokersResponse), nil
	}
}

// 查询消费组的授权用户列表
//
// 查询消费组的授权用户列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListConsumeGroupAccessPolicy(request *model.ListConsumeGroupAccessPolicyRequest) (*model.ListConsumeGroupAccessPolicyResponse, error) {
	requestDef := GenReqDefForListConsumeGroupAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumeGroupAccessPolicyResponse), nil
	}
}

// 查询消费组列表
//
// 查询消费组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListInstanceConsumerGroups(request *model.ListInstanceConsumerGroupsRequest) (*model.ListInstanceConsumerGroupsResponse, error) {
	requestDef := GenReqDefForListInstanceConsumerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

// 查询所有实例列表
//
// 查询租户的实例列表，支持按照条件查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// 查询消息轨迹
//
// 查询消息轨迹。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListMessageTrace(request *model.ListMessageTraceRequest) (*model.ListMessageTraceResponse, error) {
	requestDef := GenReqDefForListMessageTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTraceResponse), nil
	}
}

// 查询消息
//
// 查询消息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListMessages(request *model.ListMessagesRequest) (*model.ListMessagesResponse, error) {
	requestDef := GenReqDefForListMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesResponse), nil
	}
}

// 查询主题的授权用户列表
//
// 查询主题的授权用户列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListTopicAccessPolicy(request *model.ListTopicAccessPolicyRequest) (*model.ListTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForListTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicAccessPolicyResponse), nil
	}
}

// 查询用户列表
//
// 查询用户列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

// 重置消费进度
//
// 重置消费进度。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ResetConsumeOffset(request *model.ResetConsumeOffsetRequest) (*model.ResetConsumeOffsetResponse, error) {
	requestDef := GenReqDefForResetConsumeOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConsumeOffsetResponse), nil
	}
}

// 查询消费列表或详情
//
// 查询消费列表或详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowConsumerListOrDetails(request *model.ShowConsumerListOrDetailsRequest) (*model.ShowConsumerListOrDetailsResponse, error) {
	requestDef := GenReqDefForShowConsumerListOrDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerListOrDetailsResponse), nil
	}
}

// 查询指定消费组
//
// 查询指定消费组详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

// 查询指定实例
//
// 查询指定实例的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// 查询用户详情
//
// 查询用户详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserResponse), nil
	}
}

// 修改消费组
//
// 修改指定消费组参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) UpdateConsumerGroup(request *model.UpdateConsumerGroupRequest) (*model.UpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConsumerGroupResponse), nil
	}
}

// 修改实例信息
//
// 修改实例的名称和描述信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// 修改用户参数
//
// 修改用户参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// 创建主题或批量删除主题
//
// 创建主题或批量删除主题。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) CreateTopicOrBatchDeleteTopic(request *model.CreateTopicOrBatchDeleteTopicRequest) (*model.CreateTopicOrBatchDeleteTopicResponse, error) {
	requestDef := GenReqDefForCreateTopicOrBatchDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTopicOrBatchDeleteTopicResponse), nil
	}
}

// 删除指定主题
//
// 删除指定主题。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) DeleteTopic(request *model.DeleteTopicRequest) (*model.DeleteTopicResponse, error) {
	requestDef := GenReqDefForDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicResponse), nil
	}
}

// 查询主题消费组列表
//
// 查询主题消费组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ListConsumerGroupOfTopic(request *model.ListConsumerGroupOfTopicRequest) (*model.ListConsumerGroupOfTopicResponse, error) {
	requestDef := GenReqDefForListConsumerGroupOfTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumerGroupOfTopicResponse), nil
	}
}

// 查询单个主题
//
// 查询单个主题。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowOneTopic(request *model.ShowOneTopicRequest) (*model.ShowOneTopicResponse, error) {
	requestDef := GenReqDefForShowOneTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOneTopicResponse), nil
	}
}

// 查询主题的消息数
//
// 查询主题的消息数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) ShowTopicStatus(request *model.ShowTopicStatusRequest) (*model.ShowTopicStatusResponse, error) {
	requestDef := GenReqDefForShowTopicStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicStatusResponse), nil
	}
}

// 修改主题
//
// 修改主题。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RocketMQClient) UpdateTopic(request *model.UpdateTopicRequest) (*model.UpdateTopicResponse, error) {
	requestDef := GenReqDefForUpdateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicResponse), nil
	}
}
