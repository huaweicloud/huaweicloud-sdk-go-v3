package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kafka/v2/model"
)

type KafkaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKafkaClient(hcClient *http_client.HcHttpClient) *KafkaClient {
	return &KafkaClient{HcClient: hcClient}
}

func KafkaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 批量添加或删除实例标签
//
// 批量添加或删除实例标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) BatchCreateOrDeleteKafkaTag(request *model.BatchCreateOrDeleteKafkaTagRequest) (*model.BatchCreateOrDeleteKafkaTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteKafkaTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteKafkaTagResponse), nil
	}
}

// Kafka实例批量删除Topic
//
// 该接口用于向Kafka实例批量删除Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) BatchDeleteInstanceTopic(request *model.BatchDeleteInstanceTopicRequest) (*model.BatchDeleteInstanceTopicResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceTopicResponse), nil
	}
}

// 批量删除用户
//
// 批量删除Kafka实例的用户
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) BatchDeleteInstanceUsers(request *model.BatchDeleteInstanceUsersRequest) (*model.BatchDeleteInstanceUsersResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceUsersResponse), nil
	}
}

// 批量重启或删除实例
//
// 批量重启或删除实例。
//
// 在实例重启过程中，客户端的生产与消费消息等请求会被拒绝。
//
// 实例删除后，实例中原有的数据将被删除，且没有备份，请谨慎操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) BatchRestartOrDeleteInstances(request *model.BatchRestartOrDeleteInstancesRequest) (*model.BatchRestartOrDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestartOrDeleteInstancesResponse), nil
	}
}

// 创建实例的转储节点
//
// 创建实例的转储节点。
//
// **当前通过调用API，只支持按需实例创建转储节点。**
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreateConnector(request *model.CreateConnectorRequest) (*model.CreateConnectorResponse, error) {
	requestDef := GenReqDefForCreateConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectorResponse), nil
	}
}

// Kafka实例创建Topic
//
// 该接口用于向Kafka实例创建Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreateInstanceTopic(request *model.CreateInstanceTopicRequest) (*model.CreateInstanceTopicResponse, error) {
	requestDef := GenReqDefForCreateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceTopicResponse), nil
	}
}

// 创建用户
//
// 创建Kafka实例的用户，用户可连接开启SASL的Kafka实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreateInstanceUser(request *model.CreateInstanceUserRequest) (*model.CreateInstanceUserResponse, error) {
	requestDef := GenReqDefForCreateInstanceUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceUserResponse), nil
	}
}

// 新增Kafka实例指定Topic分区
//
// 新增Kafka实例指定Topic分区。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreatePartition(request *model.CreatePartitionRequest) (*model.CreatePartitionResponse, error) {
	requestDef := GenReqDefForCreatePartition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartitionResponse), nil
	}
}

// 创建实例（按需）
//
// 创建实例，该接口创建的实例为按需计费的方式。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

// 创建转储任务
//
// 创建转储任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) CreateSinkTask(request *model.CreateSinkTaskRequest) (*model.CreateSinkTaskResponse, error) {
	requestDef := GenReqDefForCreateSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSinkTaskResponse), nil
	}
}

// 删除后台任务管理中的指定记录
//
// 删除后台任务管理中的指定记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// 删除指定的实例
//
// 删除指定的实例，释放该实例的所有资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// 删除单个转储任务
//
// 删除单个转储任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) DeleteSinkTask(request *model.DeleteSinkTaskRequest) (*model.DeleteSinkTaskResponse, error) {
	requestDef := GenReqDefForDeleteSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSinkTaskResponse), nil
	}
}

// 查询可用区信息
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// 查询实例的后台任务列表
//
// 查询实例的后台任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListBackgroundTasks(request *model.ListBackgroundTasksRequest) (*model.ListBackgroundTasksResponse, error) {
	requestDef := GenReqDefForListBackgroundTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTasksResponse), nil
	}
}

// 查询产品规格列表
//
// 查询产品规格列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListEngineProducts(request *model.ListEngineProductsRequest) (*model.ListEngineProductsResponse, error) {
	requestDef := GenReqDefForListEngineProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEngineProductsResponse), nil
	}
}

// Kafka实例查询Topic
//
// 该接口用于查询指定Kafka实例的Topic详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListInstanceTopics(request *model.ListInstanceTopicsRequest) (*model.ListInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTopicsResponse), nil
	}
}

// 查询所有实例列表
//
// 查询租户的实例列表，支持按照条件查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// 查询产品规格列表
//
// 在创建kafka实例时，需要配置订购的产品ID（即product_id），可通过该接口查询产品规格。
//
// 例如，要订购按需计费、基准带宽为100MB的kafka实例，可从接口响应消息中，查找Hourly的消息体，然后找到bandwidth为100MB的记录对应的product_id，该product_id的值即是创建上述kafka实例时需要配置的产品ID。
//
// 同时，unavailable_zones字段表示资源不足的可用区列表，如果为空，则表示所有可用区都有资源，如果不为空，则表示字段值的可用区没有资源。所以必须确保您购买的资源所在的可用区有资源，不在该字段列表内。
//
// 例如，响应消息中bandwidth字段为1200MB的记录，unavailable_zones字段包含cn-east-2b、cn-east-2a和cn-east-2d，表示在华东-上海2的可用区1、可用区2、可用区3都没有该资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// 查询转储任务列表
//
// 查询转储任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ListSinkTasks(request *model.ListSinkTasksRequest) (*model.ListSinkTasksResponse, error) {
	requestDef := GenReqDefForListSinkTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSinkTasksResponse), nil
	}
}

// 重置Manager密码
//
// 重置Manager密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ResetManagerPassword(request *model.ResetManagerPasswordRequest) (*model.ResetManagerPasswordResponse, error) {
	requestDef := GenReqDefForResetManagerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetManagerPasswordResponse), nil
	}
}

// 重置消费组消费进度到指定位置
//
// Kafka实例不支持在线重置消费进度。在执行重置消费进度之前，必须停止被重置消费组客户端。
//
// &gt; 在停止被重置消费组客户端后，需要经过ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG配置的时间（默认10000毫秒），服务端才认为消费组客户端真正下线。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ResetMessageOffset(request *model.ResetMessageOffsetRequest) (*model.ResetMessageOffsetResponse, error) {
	requestDef := GenReqDefForResetMessageOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMessageOffsetResponse), nil
	}
}

// 重置密码
//
// 重置密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// 重置用户密码
//
// 重置用户密码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ResetUserPasswrod(request *model.ResetUserPasswrodRequest) (*model.ResetUserPasswrodResponse, error) {
	requestDef := GenReqDefForResetUserPasswrod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetUserPasswrodResponse), nil
	}
}

// 实例规格变更
//
// 实例规格变更。
//
// **当前通过调用API，只支持按需实例进行实例规格变更。**
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// 重启Manager
//
// 重启Manager。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) RestartManager(request *model.RestartManagerRequest) (*model.RestartManagerResponse, error) {
	requestDef := GenReqDefForRestartManager()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartManagerResponse), nil
	}
}

// 查询后台任务管理中的指定记录
//
// 查询后台任务管理中的指定记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowBackgroundTask(request *model.ShowBackgroundTaskRequest) (*model.ShowBackgroundTaskResponse, error) {
	requestDef := GenReqDefForShowBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskResponse), nil
	}
}

// 查询实例在CES的监控层级关系
//
// 查询实例在CES的监控层级关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowCesHierarchy(request *model.ShowCesHierarchyRequest) (*model.ShowCesHierarchyResponse, error) {
	requestDef := GenReqDefForShowCesHierarchy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCesHierarchyResponse), nil
	}
}

// 查询Kafka集群元数据信息
//
// 查询Kafka集群元数据信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// 查询Kafka实例的协调器信息
//
// 查询Kafka实例的协调器信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowCoordinators(request *model.ShowCoordinatorsRequest) (*model.ShowCoordinatorsResponse, error) {
	requestDef := GenReqDefForShowCoordinators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCoordinatorsResponse), nil
	}
}

// 查询消费组信息
//
// 查询消费组信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowGroups(request *model.ShowGroupsRequest) (*model.ShowGroupsResponse, error) {
	requestDef := GenReqDefForShowGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsResponse), nil
	}
}

// 查询指定实例
//
// 查询指定实例的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// 查询实例的扩容规格列表
//
// 查询实例的扩容规格列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceExtendProductInfo(request *model.ShowInstanceExtendProductInfoRequest) (*model.ShowInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceExtendProductInfoResponse), nil
	}
}

// 查询消息
//
// 查询消息的偏移量和消息内容。
// 先根据时间戳查询消息的偏移量，再根据偏移量查询消息内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceMessages(request *model.ShowInstanceMessagesRequest) (*model.ShowInstanceMessagesResponse, error) {
	requestDef := GenReqDefForShowInstanceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMessagesResponse), nil
	}
}

// 查询Kafka实例Topic详细信息
//
// 查询Kafka实例Topic详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceTopicDetail(request *model.ShowInstanceTopicDetailRequest) (*model.ShowInstanceTopicDetailResponse, error) {
	requestDef := GenReqDefForShowInstanceTopicDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceTopicDetailResponse), nil
	}
}

// 查询用户列表
//
// 查询用户列表。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceUsers(request *model.ShowInstanceUsersRequest) (*model.ShowInstanceUsersResponse, error) {
	requestDef := GenReqDefForShowInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceUsersResponse), nil
	}
}

// 查询项目标签
//
// 查询项目标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaProjectTags(request *model.ShowKafkaProjectTagsRequest) (*model.ShowKafkaProjectTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaProjectTagsResponse), nil
	}
}

// 查询实例标签
//
// 查询实例标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaTags(request *model.ShowKafkaTagsRequest) (*model.ShowKafkaTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTagsResponse), nil
	}
}

// 查询topic的磁盘存储情况
//
// 查询topic在Broker上磁盘占用情况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaTopicPartitionDiskusage(request *model.ShowKafkaTopicPartitionDiskusageRequest) (*model.ShowKafkaTopicPartitionDiskusageResponse, error) {
	requestDef := GenReqDefForShowKafkaTopicPartitionDiskusage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTopicPartitionDiskusageResponse), nil
	}
}

// 查询维护时间窗时间段
//
// 查询维护时间窗开始时间和结束时间。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowMaintainWindows(request *model.ShowMaintainWindowsRequest) (*model.ShowMaintainWindowsResponse, error) {
	requestDef := GenReqDefForShowMaintainWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMaintainWindowsResponse), nil
	}
}

// 查询分区指定时间段的消息
//
// 查询分区指定时间段的消息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowMessages(request *model.ShowMessagesRequest) (*model.ShowMessagesResponse, error) {
	requestDef := GenReqDefForShowMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessagesResponse), nil
	}
}

// 查询分区最早消息的位置
//
// 查询分区最早消息的位置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionBeginningMessage(request *model.ShowPartitionBeginningMessageRequest) (*model.ShowPartitionBeginningMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionBeginningMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionBeginningMessageResponse), nil
	}
}

// 查询分区最新消息的位置
//
// 查询分区最新消息的位置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionEndMessage(request *model.ShowPartitionEndMessageRequest) (*model.ShowPartitionEndMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionEndMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionEndMessageResponse), nil
	}
}

// 查询分区指定偏移量的消息
//
// 查询分区指定偏移量的消息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionMessage(request *model.ShowPartitionMessageRequest) (*model.ShowPartitionMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionMessageResponse), nil
	}
}

// 查询单个转储任务
//
// 查询单个转储任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowSinkTaskDetail(request *model.ShowSinkTaskDetailRequest) (*model.ShowSinkTaskDetailResponse, error) {
	requestDef := GenReqDefForShowSinkTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSinkTaskDetailResponse), nil
	}
}

// 查询用户权限
//
// 查询用户权限。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) ShowTopicAccessPolicy(request *model.ShowTopicAccessPolicyRequest) (*model.ShowTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicAccessPolicyResponse), nil
	}
}

// 修改实例信息
//
// 修改实例的名称和描述信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// 开启或关闭实例自动创建topic功能
//
// 开启或关闭实例自动创建topic功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceAutoCreateTopic(request *model.UpdateInstanceAutoCreateTopicRequest) (*model.UpdateInstanceAutoCreateTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceAutoCreateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceAutoCreateTopicResponse), nil
	}
}

// 修改实例跨VPC访问的内网IP
//
// 修改实例跨VPC访问的内网IP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceCrossVpcIp(request *model.UpdateInstanceCrossVpcIpRequest) (*model.UpdateInstanceCrossVpcIpResponse, error) {
	requestDef := GenReqDefForUpdateInstanceCrossVpcIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceCrossVpcIpResponse), nil
	}
}

// 修改Kafka实例Topic
//
// 修改Kafka实例Topic
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceTopic(request *model.UpdateInstanceTopicRequest) (*model.UpdateInstanceTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceTopicResponse), nil
	}
}

// 修改转储任务的配额
//
// 修改转储任务的配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateSinkTaskQuota(request *model.UpdateSinkTaskQuotaRequest) (*model.UpdateSinkTaskQuotaResponse, error) {
	requestDef := GenReqDefForUpdateSinkTaskQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSinkTaskQuotaResponse), nil
	}
}

// 设置用户权限
//
// 设置用户权限。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateTopicAccessPolicy(request *model.UpdateTopicAccessPolicyRequest) (*model.UpdateTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

// 修改Kafka实例Topic分区的副本
//
// 修改Kafka实例Topic分区的副本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KafkaClient) UpdateTopicReplica(request *model.UpdateTopicReplicaRequest) (*model.UpdateTopicReplicaResponse, error) {
	requestDef := GenReqDefForUpdateTopicReplica()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicReplicaResponse), nil
	}
}
