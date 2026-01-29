package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rocketmq/v2/model"
)

type RocketMQClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRocketMQClient(hcClient *httpclient.HcHttpClient) *RocketMQClient {
	return &RocketMQClient{HcClient: hcClient}
}

func RocketMQClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateOrDeleteRocketmqTag 批量添加或删除实例标签
//
// 批量添加或删除实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchCreateOrDeleteRocketmqTag(request *model.BatchCreateOrDeleteRocketmqTagRequest) (*model.BatchCreateOrDeleteRocketmqTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteRocketmqTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteRocketmqTagResponse), nil
	}
}

// BatchCreateOrDeleteRocketmqTagInvoker 批量添加或删除实例标签
func (c *RocketMQClient) BatchCreateOrDeleteRocketmqTagInvoker(request *model.BatchCreateOrDeleteRocketmqTagRequest) *BatchCreateOrDeleteRocketmqTagInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteRocketmqTag()
	return &BatchCreateOrDeleteRocketmqTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDiagnosisRecords 批量删除实例诊断报告
//
// 批量删除实例诊断报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchDeleteDiagnosisRecords(request *model.BatchDeleteDiagnosisRecordsRequest) (*model.BatchDeleteDiagnosisRecordsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDiagnosisRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDiagnosisRecordsResponse), nil
	}
}

// BatchDeleteDiagnosisRecordsInvoker 批量删除实例诊断报告
func (c *RocketMQClient) BatchDeleteDiagnosisRecordsInvoker(request *model.BatchDeleteDiagnosisRecordsRequest) *BatchDeleteDiagnosisRecordsInvoker {
	requestDef := GenReqDefForBatchDeleteDiagnosisRecords()
	return &BatchDeleteDiagnosisRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDiagnosisRecordsForRocketMq 批量删除实例诊断报告
//
// 批量删除实例诊断报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchDeleteDiagnosisRecordsForRocketMq(request *model.BatchDeleteDiagnosisRecordsForRocketMqRequest) (*model.BatchDeleteDiagnosisRecordsForRocketMqResponse, error) {
	requestDef := GenReqDefForBatchDeleteDiagnosisRecordsForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDiagnosisRecordsForRocketMqResponse), nil
	}
}

// BatchDeleteDiagnosisRecordsForRocketMqInvoker 批量删除实例诊断报告
func (c *RocketMQClient) BatchDeleteDiagnosisRecordsForRocketMqInvoker(request *model.BatchDeleteDiagnosisRecordsForRocketMqRequest) *BatchDeleteDiagnosisRecordsForRocketMqInvoker {
	requestDef := GenReqDefForBatchDeleteDiagnosisRecordsForRocketMq()
	return &BatchDeleteDiagnosisRecordsForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInstances 批量删除实例
//
// 批量删除实例。**实例删除后，实例中原有的数据将被删除，且没有备份，请谨慎操作。**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

// BatchDeleteInstancesInvoker 批量删除实例
func (c *RocketMQClient) BatchDeleteInstancesInvoker(request *model.BatchDeleteInstancesRequest) *BatchDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchDeleteInstances()
	return &BatchDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRocketMqMigrationTask 批量删除元数据迁移任务
//
// 批量删除元数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchDeleteRocketMqMigrationTask(request *model.BatchDeleteRocketMqMigrationTaskRequest) (*model.BatchDeleteRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForBatchDeleteRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRocketMqMigrationTaskResponse), nil
	}
}

// BatchDeleteRocketMqMigrationTaskInvoker 批量删除元数据迁移任务
func (c *RocketMQClient) BatchDeleteRocketMqMigrationTaskInvoker(request *model.BatchDeleteRocketMqMigrationTaskRequest) *BatchDeleteRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForBatchDeleteRocketMqMigrationTask()
	return &BatchDeleteRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateConsumerGroup 批量修改消费组
//
// 批量修改消费组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) BatchUpdateConsumerGroup(request *model.BatchUpdateConsumerGroupRequest) (*model.BatchUpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForBatchUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateConsumerGroupResponse), nil
	}
}

// BatchUpdateConsumerGroupInvoker 批量修改消费组
func (c *RocketMQClient) BatchUpdateConsumerGroupInvoker(request *model.BatchUpdateConsumerGroupRequest) *BatchUpdateConsumerGroupInvoker {
	requestDef := GenReqDefForBatchUpdateConsumerGroup()
	return &BatchUpdateConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConsumerGroupOrBatchDeleteConsumerGroup 创建消费组或批量删除消费组
//
// 创建消费组或批量删除消费组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateConsumerGroupOrBatchDeleteConsumerGroup(request *model.CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) (*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForCreateConsumerGroupOrBatchDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse), nil
	}
}

// CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker 创建消费组或批量删除消费组
func (c *RocketMQClient) CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker(request *model.CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) *CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker {
	requestDef := GenReqDefForCreateConsumerGroupOrBatchDeleteConsumerGroup()
	return &CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnosisTask 创建实例诊断任务
//
// 创建实例诊断任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateDiagnosisTask(request *model.CreateDiagnosisTaskRequest) (*model.CreateDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskResponse), nil
	}
}

// CreateDiagnosisTaskInvoker 创建实例诊断任务
func (c *RocketMQClient) CreateDiagnosisTaskInvoker(request *model.CreateDiagnosisTaskRequest) *CreateDiagnosisTaskInvoker {
	requestDef := GenReqDefForCreateDiagnosisTask()
	return &CreateDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnosisTaskForRocketMq 创建实例诊断任务
//
// 创建实例诊断任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateDiagnosisTaskForRocketMq(request *model.CreateDiagnosisTaskForRocketMqRequest) (*model.CreateDiagnosisTaskForRocketMqResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTaskForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskForRocketMqResponse), nil
	}
}

// CreateDiagnosisTaskForRocketMqInvoker 创建实例诊断任务
func (c *RocketMQClient) CreateDiagnosisTaskForRocketMqInvoker(request *model.CreateDiagnosisTaskForRocketMqRequest) *CreateDiagnosisTaskForRocketMqInvoker {
	requestDef := GenReqDefForCreateDiagnosisTaskForRocketMq()
	return &CreateDiagnosisTaskForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceByEngine 创建实例
//
// 创建实例[，该接口支持创建按需和包周期两种计费方式的实例](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,cmcc,ax)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateInstanceByEngine(request *model.CreateInstanceByEngineRequest) (*model.CreateInstanceByEngineResponse, error) {
	requestDef := GenReqDefForCreateInstanceByEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceByEngineResponse), nil
	}
}

// CreateInstanceByEngineInvoker 创建实例
func (c *RocketMQClient) CreateInstanceByEngineInvoker(request *model.CreateInstanceByEngineRequest) *CreateInstanceByEngineInvoker {
	requestDef := GenReqDefForCreateInstanceByEngine()
	return &CreateInstanceByEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidInstance 创建实例（按需）
//
// 创建实例，该接口创建的实例为按需计费的方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

// CreatePostPaidInstanceInvoker 创建实例（按需）
func (c *RocketMQClient) CreatePostPaidInstanceInvoker(request *model.CreatePostPaidInstanceRequest) *CreatePostPaidInstanceInvoker {
	requestDef := GenReqDefForCreatePostPaidInstance()
	return &CreatePostPaidInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidInstanceForRocketMq 创建实例
//
// 创建实例[，该接口支持创建按需和包周期两种计费方式的实例](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,cmcc,ax)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreatePostPaidInstanceForRocketMq(request *model.CreatePostPaidInstanceForRocketMqRequest) (*model.CreatePostPaidInstanceForRocketMqResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstanceForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceForRocketMqResponse), nil
	}
}

// CreatePostPaidInstanceForRocketMqInvoker 创建实例
func (c *RocketMQClient) CreatePostPaidInstanceForRocketMqInvoker(request *model.CreatePostPaidInstanceForRocketMqRequest) *CreatePostPaidInstanceForRocketMqInvoker {
	requestDef := GenReqDefForCreatePostPaidInstanceForRocketMq()
	return &CreatePostPaidInstanceForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRocketMqMigrationTask 新建元数据迁移任务
//
// 新建元数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateRocketMqMigrationTask(request *model.CreateRocketMqMigrationTaskRequest) (*model.CreateRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRocketMqMigrationTaskResponse), nil
	}
}

// CreateRocketMqMigrationTaskInvoker 新建元数据迁移任务
func (c *RocketMQClient) CreateRocketMqMigrationTaskInvoker(request *model.CreateRocketMqMigrationTaskRequest) *CreateRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForCreateRocketMqMigrationTask()
	return &CreateRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 创建用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *RocketMQClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackgroundTask 删除后台任务管理中的指定记录
//
// 删除后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// DeleteBackgroundTaskInvoker 删除后台任务管理中的指定记录
func (c *RocketMQClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConsumerGroup 删除指定消费组
//
// 删除指定消费组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteConsumerGroup(request *model.DeleteConsumerGroupRequest) (*model.DeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConsumerGroupResponse), nil
	}
}

// DeleteConsumerGroupInvoker 删除指定消费组
func (c *RocketMQClient) DeleteConsumerGroupInvoker(request *model.DeleteConsumerGroupRequest) *DeleteConsumerGroupInvoker {
	requestDef := GenReqDefForDeleteConsumerGroup()
	return &DeleteConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除指定的实例
//
// 删除指定的实例，释放该实例的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除指定的实例
func (c *RocketMQClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduledTask 删除定时任务管理中的指定记录
//
// 删除定时任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteScheduledTask(request *model.DeleteScheduledTaskRequest) (*model.DeleteScheduledTaskResponse, error) {
	requestDef := GenReqDefForDeleteScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduledTaskResponse), nil
	}
}

// DeleteScheduledTaskInvoker 删除定时任务管理中的指定记录
func (c *RocketMQClient) DeleteScheduledTaskInvoker(request *model.DeleteScheduledTaskRequest) *DeleteScheduledTaskInvoker {
	requestDef := GenReqDefForDeleteScheduledTask()
	return &DeleteScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 删除用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *RocketMQClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDns 开启RocketMQ实例域名访问能力
//
// 开启RocketMQ实例域名访问能力。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) EnableDns(request *model.EnableDnsRequest) (*model.EnableDnsResponse, error) {
	requestDef := GenReqDefForEnableDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDnsResponse), nil
	}
}

// EnableDnsInvoker 开启RocketMQ实例域名访问能力
func (c *RocketMQClient) EnableDnsInvoker(request *model.EnableDnsRequest) *EnableDnsInvoker {
	requestDef := GenReqDefForEnableDns()
	return &EnableDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 查询可用区信息
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 查询可用区信息
func (c *RocketMQClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackgroundTasks 查询实例的后台任务列表
//
// 查询实例的后台任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListBackgroundTasks(request *model.ListBackgroundTasksRequest) (*model.ListBackgroundTasksResponse, error) {
	requestDef := GenReqDefForListBackgroundTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTasksResponse), nil
	}
}

// ListBackgroundTasksInvoker 查询实例的后台任务列表
func (c *RocketMQClient) ListBackgroundTasksInvoker(request *model.ListBackgroundTasksRequest) *ListBackgroundTasksInvoker {
	requestDef := GenReqDefForListBackgroundTasks()
	return &ListBackgroundTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBrokers 查询代理列表
//
// 查询代理列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListBrokers(request *model.ListBrokersRequest) (*model.ListBrokersResponse, error) {
	requestDef := GenReqDefForListBrokers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBrokersResponse), nil
	}
}

// ListBrokersInvoker 查询代理列表
func (c *RocketMQClient) ListBrokersInvoker(request *model.ListBrokersRequest) *ListBrokersInvoker {
	requestDef := GenReqDefForListBrokers()
	return &ListBrokersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigFeatures 获取特性开关列表
//
// 获取特性开关列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListConfigFeatures(request *model.ListConfigFeaturesRequest) (*model.ListConfigFeaturesResponse, error) {
	requestDef := GenReqDefForListConfigFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigFeaturesResponse), nil
	}
}

// ListConfigFeaturesInvoker 获取特性开关列表
func (c *RocketMQClient) ListConfigFeaturesInvoker(request *model.ListConfigFeaturesRequest) *ListConfigFeaturesInvoker {
	requestDef := GenReqDefForListConfigFeatures()
	return &ListConfigFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsumeGroupAccessPolicy 查询消费组的授权用户列表
//
// 查询消费组的授权用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListConsumeGroupAccessPolicy(request *model.ListConsumeGroupAccessPolicyRequest) (*model.ListConsumeGroupAccessPolicyResponse, error) {
	requestDef := GenReqDefForListConsumeGroupAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumeGroupAccessPolicyResponse), nil
	}
}

// ListConsumeGroupAccessPolicyInvoker 查询消费组的授权用户列表
func (c *RocketMQClient) ListConsumeGroupAccessPolicyInvoker(request *model.ListConsumeGroupAccessPolicyRequest) *ListConsumeGroupAccessPolicyInvoker {
	requestDef := GenReqDefForListConsumeGroupAccessPolicy()
	return &ListConsumeGroupAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsumeGroupAccessPolicyForRocketMq 查询消费组的授权用户列表
//
// 查询消费组的授权用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListConsumeGroupAccessPolicyForRocketMq(request *model.ListConsumeGroupAccessPolicyForRocketMqRequest) (*model.ListConsumeGroupAccessPolicyForRocketMqResponse, error) {
	requestDef := GenReqDefForListConsumeGroupAccessPolicyForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumeGroupAccessPolicyForRocketMqResponse), nil
	}
}

// ListConsumeGroupAccessPolicyForRocketMqInvoker 查询消费组的授权用户列表
func (c *RocketMQClient) ListConsumeGroupAccessPolicyForRocketMqInvoker(request *model.ListConsumeGroupAccessPolicyForRocketMqRequest) *ListConsumeGroupAccessPolicyForRocketMqInvoker {
	requestDef := GenReqDefForListConsumeGroupAccessPolicyForRocketMq()
	return &ListConsumeGroupAccessPolicyForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnosisReports 查询实例诊断报告列表
//
// 查询实例诊断报告列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListDiagnosisReports(request *model.ListDiagnosisReportsRequest) (*model.ListDiagnosisReportsResponse, error) {
	requestDef := GenReqDefForListDiagnosisReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisReportsResponse), nil
	}
}

// ListDiagnosisReportsInvoker 查询实例诊断报告列表
func (c *RocketMQClient) ListDiagnosisReportsInvoker(request *model.ListDiagnosisReportsRequest) *ListDiagnosisReportsInvoker {
	requestDef := GenReqDefForListDiagnosisReports()
	return &ListDiagnosisReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnosisReportsForRocketMq 查询实例诊断报告列表
//
// 查询实例诊断报告列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListDiagnosisReportsForRocketMq(request *model.ListDiagnosisReportsForRocketMqRequest) (*model.ListDiagnosisReportsForRocketMqResponse, error) {
	requestDef := GenReqDefForListDiagnosisReportsForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisReportsForRocketMqResponse), nil
	}
}

// ListDiagnosisReportsForRocketMqInvoker 查询实例诊断报告列表
func (c *RocketMQClient) ListDiagnosisReportsForRocketMqInvoker(request *model.ListDiagnosisReportsForRocketMqRequest) *ListDiagnosisReportsForRocketMqInvoker {
	requestDef := GenReqDefForListDiagnosisReportsForRocketMq()
	return &ListDiagnosisReportsForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngineProducts 查询产品规格列表
//
// 查询相应引擎的产品规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListEngineProducts(request *model.ListEngineProductsRequest) (*model.ListEngineProductsResponse, error) {
	requestDef := GenReqDefForListEngineProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEngineProductsResponse), nil
	}
}

// ListEngineProductsInvoker 查询产品规格列表
func (c *RocketMQClient) ListEngineProductsInvoker(request *model.ListEngineProductsRequest) *ListEngineProductsInvoker {
	requestDef := GenReqDefForListEngineProducts()
	return &ListEngineProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceConsumerGroups 查询消费组列表
//
// 查询消费组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListInstanceConsumerGroups(request *model.ListInstanceConsumerGroupsRequest) (*model.ListInstanceConsumerGroupsResponse, error) {
	requestDef := GenReqDefForListInstanceConsumerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

// ListInstanceConsumerGroupsInvoker 查询消费组列表
func (c *RocketMQClient) ListInstanceConsumerGroupsInvoker(request *model.ListInstanceConsumerGroupsRequest) *ListInstanceConsumerGroupsInvoker {
	requestDef := GenReqDefForListInstanceConsumerGroups()
	return &ListInstanceConsumerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询所有实例列表
//
// 查询租户的实例列表，支持按照条件查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询所有实例列表
func (c *RocketMQClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessageTrace 查询消息轨迹
//
// 查询消息轨迹。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListMessageTrace(request *model.ListMessageTraceRequest) (*model.ListMessageTraceResponse, error) {
	requestDef := GenReqDefForListMessageTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTraceResponse), nil
	}
}

// ListMessageTraceInvoker 查询消息轨迹
func (c *RocketMQClient) ListMessageTraceInvoker(request *model.ListMessageTraceRequest) *ListMessageTraceInvoker {
	requestDef := GenReqDefForListMessageTrace()
	return &ListMessageTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessageTraceForRocketMq 查询消息轨迹
//
// 查询消息轨迹。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListMessageTraceForRocketMq(request *model.ListMessageTraceForRocketMqRequest) (*model.ListMessageTraceForRocketMqResponse, error) {
	requestDef := GenReqDefForListMessageTraceForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTraceForRocketMqResponse), nil
	}
}

// ListMessageTraceForRocketMqInvoker 查询消息轨迹
func (c *RocketMQClient) ListMessageTraceForRocketMqInvoker(request *model.ListMessageTraceForRocketMqRequest) *ListMessageTraceForRocketMqInvoker {
	requestDef := GenReqDefForListMessageTraceForRocketMq()
	return &ListMessageTraceForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessages 查询消息
//
// 查询消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListMessages(request *model.ListMessagesRequest) (*model.ListMessagesResponse, error) {
	requestDef := GenReqDefForListMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesResponse), nil
	}
}

// ListMessagesInvoker 查询消息
func (c *RocketMQClient) ListMessagesInvoker(request *model.ListMessagesRequest) *ListMessagesInvoker {
	requestDef := GenReqDefForListMessages()
	return &ListMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessagesForRocketMq 查询消息
//
// 查询消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListMessagesForRocketMq(request *model.ListMessagesForRocketMqRequest) (*model.ListMessagesForRocketMqResponse, error) {
	requestDef := GenReqDefForListMessagesForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesForRocketMqResponse), nil
	}
}

// ListMessagesForRocketMqInvoker 查询消息
func (c *RocketMQClient) ListMessagesForRocketMqInvoker(request *model.ListMessagesForRocketMqRequest) *ListMessagesForRocketMqInvoker {
	requestDef := GenReqDefForListMessagesForRocketMq()
	return &ListMessagesForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRocketMqMigrationTask 查询实例下所有迁移任务或查询指定迁移任务信息
//
// 1. 查询实例下所有迁移任务
// 2. 查询指定迁移任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListRocketMqMigrationTask(request *model.ListRocketMqMigrationTaskRequest) (*model.ListRocketMqMigrationTaskResponse, error) {
	requestDef := GenReqDefForListRocketMqMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRocketMqMigrationTaskResponse), nil
	}
}

// ListRocketMqMigrationTaskInvoker 查询实例下所有迁移任务或查询指定迁移任务信息
func (c *RocketMQClient) ListRocketMqMigrationTaskInvoker(request *model.ListRocketMqMigrationTaskRequest) *ListRocketMqMigrationTaskInvoker {
	requestDef := GenReqDefForListRocketMqMigrationTask()
	return &ListRocketMqMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTasks 查询实例的定时任务列表
//
// 查询实例的定时任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListScheduledTasks(request *model.ListScheduledTasksRequest) (*model.ListScheduledTasksResponse, error) {
	requestDef := GenReqDefForListScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTasksResponse), nil
	}
}

// ListScheduledTasksInvoker 查询实例的定时任务列表
func (c *RocketMQClient) ListScheduledTasksInvoker(request *model.ListScheduledTasksRequest) *ListScheduledTasksInvoker {
	requestDef := GenReqDefForListScheduledTasks()
	return &ListScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopicAccessPolicy 查询主题的授权用户列表
//
// 查询主题的授权用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListTopicAccessPolicy(request *model.ListTopicAccessPolicyRequest) (*model.ListTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForListTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicAccessPolicyResponse), nil
	}
}

// ListTopicAccessPolicyInvoker 查询主题的授权用户列表
func (c *RocketMQClient) ListTopicAccessPolicyInvoker(request *model.ListTopicAccessPolicyRequest) *ListTopicAccessPolicyInvoker {
	requestDef := GenReqDefForListTopicAccessPolicy()
	return &ListTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUser 查询用户列表
//
// 查询用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

// ListUserInvoker 查询用户列表
func (c *RocketMQClient) ListUserInvoker(request *model.ListUserRequest) *ListUserInvoker {
	requestDef := GenReqDefForListUser()
	return &ListUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyInstanceSslConfig 修改实例SSL相关配置
//
// 修改实例SSL相关配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ModifyInstanceSslConfig(request *model.ModifyInstanceSslConfigRequest) (*model.ModifyInstanceSslConfigResponse, error) {
	requestDef := GenReqDefForModifyInstanceSslConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyInstanceSslConfigResponse), nil
	}
}

// ModifyInstanceSslConfigInvoker 修改实例SSL相关配置
func (c *RocketMQClient) ModifyInstanceSslConfigInvoker(request *model.ModifyInstanceSslConfigRequest) *ModifyInstanceSslConfigInvoker {
	requestDef := GenReqDefForModifyInstanceSslConfig()
	return &ModifyInstanceSslConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyRecyclePolicy 更新回收站策略
//
// 更新回收站策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ModifyRecyclePolicy(request *model.ModifyRecyclePolicyRequest) (*model.ModifyRecyclePolicyResponse, error) {
	requestDef := GenReqDefForModifyRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyRecyclePolicyResponse), nil
	}
}

// ModifyRecyclePolicyInvoker 更新回收站策略
func (c *RocketMQClient) ModifyRecyclePolicyInvoker(request *model.ModifyRecyclePolicyRequest) *ModifyRecyclePolicyInvoker {
	requestDef := GenReqDefForModifyRecyclePolicy()
	return &ModifyRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetConsumeOffset 重置消费进度
//
// 重置消费进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ResetConsumeOffset(request *model.ResetConsumeOffsetRequest) (*model.ResetConsumeOffsetResponse, error) {
	requestDef := GenReqDefForResetConsumeOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConsumeOffsetResponse), nil
	}
}

// ResetConsumeOffsetInvoker 重置消费进度
func (c *RocketMQClient) ResetConsumeOffsetInvoker(request *model.ResetConsumeOffsetRequest) *ResetConsumeOffsetInvoker {
	requestDef := GenReqDefForResetConsumeOffset()
	return &ResetConsumeOffsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetConsumeOffsetForRocketMq 重置消费进度
//
// 重置消费进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ResetConsumeOffsetForRocketMq(request *model.ResetConsumeOffsetForRocketMqRequest) (*model.ResetConsumeOffsetForRocketMqResponse, error) {
	requestDef := GenReqDefForResetConsumeOffsetForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConsumeOffsetForRocketMqResponse), nil
	}
}

// ResetConsumeOffsetForRocketMqInvoker 重置消费进度
func (c *RocketMQClient) ResetConsumeOffsetForRocketMqInvoker(request *model.ResetConsumeOffsetForRocketMqRequest) *ResetConsumeOffsetForRocketMqInvoker {
	requestDef := GenReqDefForResetConsumeOffsetForRocketMq()
	return &ResetConsumeOffsetForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 实例规格变更
//
// 实例规格变更。
//
// [**当前通过调用API，只支持按需实例进行实例规格变更。**](tag:hws,hws_hk,ctc,cmcc,ax)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 实例规格变更
func (c *RocketMQClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstanceForRocketMq 实例规格变更
//
// 实例规格变更。
//
// [**当前通过调用API，只支持按需实例进行实例规格变更。**](tag:hws,hws_hk,ctc,cmcc)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ResizeInstanceForRocketMq(request *model.ResizeInstanceForRocketMqRequest) (*model.ResizeInstanceForRocketMqResponse, error) {
	requestDef := GenReqDefForResizeInstanceForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceForRocketMqResponse), nil
	}
}

// ResizeInstanceForRocketMqInvoker 实例规格变更
func (c *RocketMQClient) ResizeInstanceForRocketMqInvoker(request *model.ResizeInstanceForRocketMqRequest) *ResizeInstanceForRocketMqInvoker {
	requestDef := GenReqDefForResizeInstanceForRocketMq()
	return &ResizeInstanceForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreRecycleInstance 恢复回收站实例
//
// 恢复回收站实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) RestoreRecycleInstance(request *model.RestoreRecycleInstanceRequest) (*model.RestoreRecycleInstanceResponse, error) {
	requestDef := GenReqDefForRestoreRecycleInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreRecycleInstanceResponse), nil
	}
}

// RestoreRecycleInstanceInvoker 恢复回收站实例
func (c *RocketMQClient) RestoreRecycleInstanceInvoker(request *model.RestoreRecycleInstanceRequest) *RestoreRecycleInstanceInvoker {
	requestDef := GenReqDefForRestoreRecycleInstance()
	return &RestoreRecycleInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendDlqMessage 重发死信消息
//
// 重发死信消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) SendDlqMessage(request *model.SendDlqMessageRequest) (*model.SendDlqMessageResponse, error) {
	requestDef := GenReqDefForSendDlqMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendDlqMessageResponse), nil
	}
}

// SendDlqMessageInvoker 重发死信消息
func (c *RocketMQClient) SendDlqMessageInvoker(request *model.SendDlqMessageRequest) *SendDlqMessageInvoker {
	requestDef := GenReqDefForSendDlqMessage()
	return &SendDlqMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendDlqMessageForRocketMq 重发死信消息
//
// 重发死信消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) SendDlqMessageForRocketMq(request *model.SendDlqMessageForRocketMqRequest) (*model.SendDlqMessageForRocketMqResponse, error) {
	requestDef := GenReqDefForSendDlqMessageForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendDlqMessageForRocketMqResponse), nil
	}
}

// SendDlqMessageForRocketMqInvoker 重发死信消息
func (c *RocketMQClient) SendDlqMessageForRocketMqInvoker(request *model.SendDlqMessageForRocketMqRequest) *SendDlqMessageForRocketMqInvoker {
	requestDef := GenReqDefForSendDlqMessageForRocketMq()
	return &SendDlqMessageForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendMessage 发送消息
//
// 发送消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) SendMessage(request *model.SendMessageRequest) (*model.SendMessageResponse, error) {
	requestDef := GenReqDefForSendMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendMessageResponse), nil
	}
}

// SendMessageInvoker 发送消息
func (c *RocketMQClient) SendMessageInvoker(request *model.SendMessageRequest) *SendMessageInvoker {
	requestDef := GenReqDefForSendMessage()
	return &SendMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendMessageForRocketMq 发送消息
//
// 发送消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) SendMessageForRocketMq(request *model.SendMessageForRocketMqRequest) (*model.SendMessageForRocketMqResponse, error) {
	requestDef := GenReqDefForSendMessageForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendMessageForRocketMqResponse), nil
	}
}

// SendMessageForRocketMqInvoker 发送消息
func (c *RocketMQClient) SendMessageForRocketMqInvoker(request *model.SendMessageForRocketMqRequest) *SendMessageForRocketMqInvoker {
	requestDef := GenReqDefForSendMessageForRocketMq()
	return &SendMessageForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackgroundTask 查询后台任务管理中的指定记录
//
// 查询后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowBackgroundTask(request *model.ShowBackgroundTaskRequest) (*model.ShowBackgroundTaskResponse, error) {
	requestDef := GenReqDefForShowBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskResponse), nil
	}
}

// ShowBackgroundTaskInvoker 查询后台任务管理中的指定记录
func (c *RocketMQClient) ShowBackgroundTaskInvoker(request *model.ShowBackgroundTaskRequest) *ShowBackgroundTaskInvoker {
	requestDef := GenReqDefForShowBackgroundTask()
	return &ShowBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackgroundTaskProgress 查询实例变更记录的进度
//
// 查询后台任务管理中的指定实例变更记录的进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowBackgroundTaskProgress(request *model.ShowBackgroundTaskProgressRequest) (*model.ShowBackgroundTaskProgressResponse, error) {
	requestDef := GenReqDefForShowBackgroundTaskProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskProgressResponse), nil
	}
}

// ShowBackgroundTaskProgressInvoker 查询实例变更记录的进度
func (c *RocketMQClient) ShowBackgroundTaskProgressInvoker(request *model.ShowBackgroundTaskProgressRequest) *ShowBackgroundTaskProgressInvoker {
	requestDef := GenReqDefForShowBackgroundTaskProgress()
	return &ShowBackgroundTaskProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCesHierarchy 查询实例在CES的监控层级关系
//
// 查询实例在CES的监控层级关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowCesHierarchy(request *model.ShowCesHierarchyRequest) (*model.ShowCesHierarchyResponse, error) {
	requestDef := GenReqDefForShowCesHierarchy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCesHierarchyResponse), nil
	}
}

// ShowCesHierarchyInvoker 查询实例在CES的监控层级关系
func (c *RocketMQClient) ShowCesHierarchyInvoker(request *model.ShowCesHierarchyRequest) *ShowCesHierarchyInvoker {
	requestDef := GenReqDefForShowCesHierarchy()
	return &ShowCesHierarchyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsumerConnections 查询消费者列表
//
// 查询消费组内消费者列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowConsumerConnections(request *model.ShowConsumerConnectionsRequest) (*model.ShowConsumerConnectionsResponse, error) {
	requestDef := GenReqDefForShowConsumerConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerConnectionsResponse), nil
	}
}

// ShowConsumerConnectionsInvoker 查询消费者列表
func (c *RocketMQClient) ShowConsumerConnectionsInvoker(request *model.ShowConsumerConnectionsRequest) *ShowConsumerConnectionsInvoker {
	requestDef := GenReqDefForShowConsumerConnections()
	return &ShowConsumerConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsumerConnectionsForRocketMq 查询消费者列表
//
// 查询消费组内消费者列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowConsumerConnectionsForRocketMq(request *model.ShowConsumerConnectionsForRocketMqRequest) (*model.ShowConsumerConnectionsForRocketMqResponse, error) {
	requestDef := GenReqDefForShowConsumerConnectionsForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerConnectionsForRocketMqResponse), nil
	}
}

// ShowConsumerConnectionsForRocketMqInvoker 查询消费者列表
func (c *RocketMQClient) ShowConsumerConnectionsForRocketMqInvoker(request *model.ShowConsumerConnectionsForRocketMqRequest) *ShowConsumerConnectionsForRocketMqInvoker {
	requestDef := GenReqDefForShowConsumerConnectionsForRocketMq()
	return &ShowConsumerConnectionsForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsumerListOrDetails 查询消费列表或详情
//
// 查询消费列表或详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowConsumerListOrDetails(request *model.ShowConsumerListOrDetailsRequest) (*model.ShowConsumerListOrDetailsResponse, error) {
	requestDef := GenReqDefForShowConsumerListOrDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerListOrDetailsResponse), nil
	}
}

// ShowConsumerListOrDetailsInvoker 查询消费列表或详情
func (c *RocketMQClient) ShowConsumerListOrDetailsInvoker(request *model.ShowConsumerListOrDetailsRequest) *ShowConsumerListOrDetailsInvoker {
	requestDef := GenReqDefForShowConsumerListOrDetails()
	return &ShowConsumerListOrDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisReport 查询实例诊断报告
//
// 查询实例诊断报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowDiagnosisReport(request *model.ShowDiagnosisReportRequest) (*model.ShowDiagnosisReportResponse, error) {
	requestDef := GenReqDefForShowDiagnosisReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisReportResponse), nil
	}
}

// ShowDiagnosisReportInvoker 查询实例诊断报告
func (c *RocketMQClient) ShowDiagnosisReportInvoker(request *model.ShowDiagnosisReportRequest) *ShowDiagnosisReportInvoker {
	requestDef := GenReqDefForShowDiagnosisReport()
	return &ShowDiagnosisReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisReportForRocketMq 查询实例诊断报告
//
// 查询实例诊断报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowDiagnosisReportForRocketMq(request *model.ShowDiagnosisReportForRocketMqRequest) (*model.ShowDiagnosisReportForRocketMqResponse, error) {
	requestDef := GenReqDefForShowDiagnosisReportForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisReportForRocketMqResponse), nil
	}
}

// ShowDiagnosisReportForRocketMqInvoker 查询实例诊断报告
func (c *RocketMQClient) ShowDiagnosisReportForRocketMqInvoker(request *model.ShowDiagnosisReportForRocketMqRequest) *ShowDiagnosisReportForRocketMqInvoker {
	requestDef := GenReqDefForShowDiagnosisReportForRocketMq()
	return &ShowDiagnosisReportForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisStack 查询stack信息
//
// 查询stack信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowDiagnosisStack(request *model.ShowDiagnosisStackRequest) (*model.ShowDiagnosisStackResponse, error) {
	requestDef := GenReqDefForShowDiagnosisStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisStackResponse), nil
	}
}

// ShowDiagnosisStackInvoker 查询stack信息
func (c *RocketMQClient) ShowDiagnosisStackInvoker(request *model.ShowDiagnosisStackRequest) *ShowDiagnosisStackInvoker {
	requestDef := GenReqDefForShowDiagnosisStack()
	return &ShowDiagnosisStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisStackForRocketMq 查询stack信息
//
// 查询stack信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowDiagnosisStackForRocketMq(request *model.ShowDiagnosisStackForRocketMqRequest) (*model.ShowDiagnosisStackForRocketMqResponse, error) {
	requestDef := GenReqDefForShowDiagnosisStackForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisStackForRocketMqResponse), nil
	}
}

// ShowDiagnosisStackForRocketMqInvoker 查询stack信息
func (c *RocketMQClient) ShowDiagnosisStackForRocketMqInvoker(request *model.ShowDiagnosisStackForRocketMqRequest) *ShowDiagnosisStackForRocketMqInvoker {
	requestDef := GenReqDefForShowDiagnosisStackForRocketMq()
	return &ShowDiagnosisStackForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineInstanceExtendProductInfo 查询实例的扩容规格列表
//
// 查询实例的扩容规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowEngineInstanceExtendProductInfo(request *model.ShowEngineInstanceExtendProductInfoRequest) (*model.ShowEngineInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineInstanceExtendProductInfoResponse), nil
	}
}

// ShowEngineInstanceExtendProductInfoInvoker 查询实例的扩容规格列表
func (c *RocketMQClient) ShowEngineInstanceExtendProductInfoInvoker(request *model.ShowEngineInstanceExtendProductInfoRequest) *ShowEngineInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()
	return &ShowEngineInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineInstanceExtendProductInfoForRocketMq 查询实例的扩容规格列表
//
// 查询实例的扩容规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowEngineInstanceExtendProductInfoForRocketMq(request *model.ShowEngineInstanceExtendProductInfoForRocketMqRequest) (*model.ShowEngineInstanceExtendProductInfoForRocketMqResponse, error) {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfoForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineInstanceExtendProductInfoForRocketMqResponse), nil
	}
}

// ShowEngineInstanceExtendProductInfoForRocketMqInvoker 查询实例的扩容规格列表
func (c *RocketMQClient) ShowEngineInstanceExtendProductInfoForRocketMqInvoker(request *model.ShowEngineInstanceExtendProductInfoForRocketMqRequest) *ShowEngineInstanceExtendProductInfoForRocketMqInvoker {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfoForRocketMq()
	return &ShowEngineInstanceExtendProductInfoForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroup 查询指定消费组
//
// 查询指定消费组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

// ShowGroupInvoker 查询指定消费组
func (c *RocketMQClient) ShowGroupInvoker(request *model.ShowGroupRequest) *ShowGroupInvoker {
	requestDef := GenReqDefForShowGroup()
	return &ShowGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询指定实例
//
// 查询指定实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询指定实例
func (c *RocketMQClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceNodes 查询实例节点
//
// 查询实例节点信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowInstanceNodes(request *model.ShowInstanceNodesRequest) (*model.ShowInstanceNodesResponse, error) {
	requestDef := GenReqDefForShowInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceNodesResponse), nil
	}
}

// ShowInstanceNodesInvoker 查询实例节点
func (c *RocketMQClient) ShowInstanceNodesInvoker(request *model.ShowInstanceNodesRequest) *ShowInstanceNodesInvoker {
	requestDef := GenReqDefForShowInstanceNodes()
	return &ShowInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查看租户配额
//
// 查询租户最大可以创建的实例个数和已创建的实例个数，以及每个实例最大可以创建标签的个数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查看租户配额
func (c *RocketMQClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecycleInstances 查询回收站实例列表
//
// 查询回收站实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRecycleInstances(request *model.ShowRecycleInstancesRequest) (*model.ShowRecycleInstancesResponse, error) {
	requestDef := GenReqDefForShowRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecycleInstancesResponse), nil
	}
}

// ShowRecycleInstancesInvoker 查询回收站实例列表
func (c *RocketMQClient) ShowRecycleInstancesInvoker(request *model.ShowRecycleInstancesRequest) *ShowRecycleInstancesInvoker {
	requestDef := GenReqDefForShowRecycleInstances()
	return &ShowRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRocketMqConfigs 查询RocketMQ配置
//
// 该接口用于查询RocketMQ配置，若成功则返回配置的相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRocketMqConfigs(request *model.ShowRocketMqConfigsRequest) (*model.ShowRocketMqConfigsResponse, error) {
	requestDef := GenReqDefForShowRocketMqConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketMqConfigsResponse), nil
	}
}

// ShowRocketMqConfigsInvoker 查询RocketMQ配置
func (c *RocketMQClient) ShowRocketMqConfigsInvoker(request *model.ShowRocketMqConfigsRequest) *ShowRocketMqConfigsInvoker {
	requestDef := GenReqDefForShowRocketMqConfigs()
	return &ShowRocketMqConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRocketMqProductCores 查询RocketMQ产品规格核数
//
// 查询RocketMQ产品规格核数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRocketMqProductCores(request *model.ShowRocketMqProductCoresRequest) (*model.ShowRocketMqProductCoresResponse, error) {
	requestDef := GenReqDefForShowRocketMqProductCores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketMqProductCoresResponse), nil
	}
}

// ShowRocketMqProductCoresInvoker 查询RocketMQ产品规格核数
func (c *RocketMQClient) ShowRocketMqProductCoresInvoker(request *model.ShowRocketMqProductCoresRequest) *ShowRocketMqProductCoresInvoker {
	requestDef := GenReqDefForShowRocketMqProductCores()
	return &ShowRocketMqProductCoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRocketMqScalePreCheckInfo 查询RocketMQ实例扩容前置检查信息
//
// 获取RocketMQ实例扩容前置检查信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRocketMqScalePreCheckInfo(request *model.ShowRocketMqScalePreCheckInfoRequest) (*model.ShowRocketMqScalePreCheckInfoResponse, error) {
	requestDef := GenReqDefForShowRocketMqScalePreCheckInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketMqScalePreCheckInfoResponse), nil
	}
}

// ShowRocketMqScalePreCheckInfoInvoker 查询RocketMQ实例扩容前置检查信息
func (c *RocketMQClient) ShowRocketMqScalePreCheckInfoInvoker(request *model.ShowRocketMqScalePreCheckInfoRequest) *ShowRocketMqScalePreCheckInfoInvoker {
	requestDef := GenReqDefForShowRocketMqScalePreCheckInfo()
	return &ShowRocketMqScalePreCheckInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRocketmqProjectTags 查询项目标签
//
// 查询项目标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRocketmqProjectTags(request *model.ShowRocketmqProjectTagsRequest) (*model.ShowRocketmqProjectTagsResponse, error) {
	requestDef := GenReqDefForShowRocketmqProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketmqProjectTagsResponse), nil
	}
}

// ShowRocketmqProjectTagsInvoker 查询项目标签
func (c *RocketMQClient) ShowRocketmqProjectTagsInvoker(request *model.ShowRocketmqProjectTagsRequest) *ShowRocketmqProjectTagsInvoker {
	requestDef := GenReqDefForShowRocketmqProjectTags()
	return &ShowRocketmqProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRocketmqTags 查询实例标签
//
// 查询实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowRocketmqTags(request *model.ShowRocketmqTagsRequest) (*model.ShowRocketmqTagsResponse, error) {
	requestDef := GenReqDefForShowRocketmqTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRocketmqTagsResponse), nil
	}
}

// ShowRocketmqTagsInvoker 查询实例标签
func (c *RocketMQClient) ShowRocketmqTagsInvoker(request *model.ShowRocketmqTagsRequest) *ShowRocketmqTagsInvoker {
	requestDef := GenReqDefForShowRocketmqTags()
	return &ShowRocketmqTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUser 查询用户详情
//
// 查询用户详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserResponse), nil
	}
}

// ShowUserInvoker 查询用户详情
func (c *RocketMQClient) ShowUserInvoker(request *model.ShowUserRequest) *ShowUserInvoker {
	requestDef := GenReqDefForShowUser()
	return &ShowUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolumeExpandConfig 查询磁盘自动扩容配置
//
// 查询磁盘自动扩容配置，包括磁盘自动扩容是否开启，以及开启后的扩容阈值、扩容步长、扩容上限信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowVolumeExpandConfig(request *model.ShowVolumeExpandConfigRequest) (*model.ShowVolumeExpandConfigResponse, error) {
	requestDef := GenReqDefForShowVolumeExpandConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeExpandConfigResponse), nil
	}
}

// ShowVolumeExpandConfigInvoker 查询磁盘自动扩容配置
func (c *RocketMQClient) ShowVolumeExpandConfigInvoker(request *model.ShowVolumeExpandConfigRequest) *ShowVolumeExpandConfigInvoker {
	requestDef := GenReqDefForShowVolumeExpandConfig()
	return &ShowVolumeExpandConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConsumerGroup 修改消费组
//
// 修改指定消费组参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateConsumerGroup(request *model.UpdateConsumerGroupRequest) (*model.UpdateConsumerGroupResponse, error) {
	requestDef := GenReqDefForUpdateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConsumerGroupResponse), nil
	}
}

// UpdateConsumerGroupInvoker 修改消费组
func (c *RocketMQClient) UpdateConsumerGroupInvoker(request *model.UpdateConsumerGroupRequest) *UpdateConsumerGroupInvoker {
	requestDef := GenReqDefForUpdateConsumerGroup()
	return &UpdateConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改实例信息
//
// 修改实例相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改实例信息
func (c *RocketMQClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRocketMqConfigs 修改RocketMQ配置
//
// 该接口用于修改RocketMQ配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateRocketMqConfigs(request *model.UpdateRocketMqConfigsRequest) (*model.UpdateRocketMqConfigsResponse, error) {
	requestDef := GenReqDefForUpdateRocketMqConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRocketMqConfigsResponse), nil
	}
}

// UpdateRocketMqConfigsInvoker 修改RocketMQ配置
func (c *RocketMQClient) UpdateRocketMqConfigsInvoker(request *model.UpdateRocketMqConfigsRequest) *UpdateRocketMqConfigsInvoker {
	requestDef := GenReqDefForUpdateRocketMqConfigs()
	return &UpdateRocketMqConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScheduledTask 修改定时任务管理中的指定记录
//
// 修改定时任务管理中的指定记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateScheduledTask(request *model.UpdateScheduledTaskRequest) (*model.UpdateScheduledTaskResponse, error) {
	requestDef := GenReqDefForUpdateScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduledTaskResponse), nil
	}
}

// UpdateScheduledTaskInvoker 修改定时任务管理中的指定记录
func (c *RocketMQClient) UpdateScheduledTaskInvoker(request *model.UpdateScheduledTaskRequest) *UpdateScheduledTaskInvoker {
	requestDef := GenReqDefForUpdateScheduledTask()
	return &UpdateScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户参数
//
// 修改用户参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改用户参数
func (c *RocketMQClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVolumeExpansionConfig 修改磁盘自动扩容配置
//
// 该接口用于修改磁盘自动扩容配置，包含磁盘自动扩容是否开启、扩容阈值、扩容步长，以及扩容上限的配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateVolumeExpansionConfig(request *model.UpdateVolumeExpansionConfigRequest) (*model.UpdateVolumeExpansionConfigResponse, error) {
	requestDef := GenReqDefForUpdateVolumeExpansionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVolumeExpansionConfigResponse), nil
	}
}

// UpdateVolumeExpansionConfigInvoker 修改磁盘自动扩容配置
func (c *RocketMQClient) UpdateVolumeExpansionConfigInvoker(request *model.UpdateVolumeExpansionConfigRequest) *UpdateVolumeExpansionConfigInvoker {
	requestDef := GenReqDefForUpdateVolumeExpansionConfig()
	return &UpdateVolumeExpansionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateConsumedMessage 消费验证
//
// 消费验证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ValidateConsumedMessage(request *model.ValidateConsumedMessageRequest) (*model.ValidateConsumedMessageResponse, error) {
	requestDef := GenReqDefForValidateConsumedMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateConsumedMessageResponse), nil
	}
}

// ValidateConsumedMessageInvoker 消费验证
func (c *RocketMQClient) ValidateConsumedMessageInvoker(request *model.ValidateConsumedMessageRequest) *ValidateConsumedMessageInvoker {
	requestDef := GenReqDefForValidateConsumedMessage()
	return &ValidateConsumedMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateConsumedMessageForRocketMq 消费验证
//
// 消费验证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ValidateConsumedMessageForRocketMq(request *model.ValidateConsumedMessageForRocketMqRequest) (*model.ValidateConsumedMessageForRocketMqResponse, error) {
	requestDef := GenReqDefForValidateConsumedMessageForRocketMq()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateConsumedMessageForRocketMqResponse), nil
	}
}

// ValidateConsumedMessageForRocketMqInvoker 消费验证
func (c *RocketMQClient) ValidateConsumedMessageForRocketMqInvoker(request *model.ValidateConsumedMessageForRocketMqRequest) *ValidateConsumedMessageForRocketMqInvoker {
	requestDef := GenReqDefForValidateConsumedMessageForRocketMq()
	return &ValidateConsumedMessageForRocketMqInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTopicOrBatchDeleteTopic 创建主题或批量删除主题
//
// 创建主题或批量删除主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) CreateTopicOrBatchDeleteTopic(request *model.CreateTopicOrBatchDeleteTopicRequest) (*model.CreateTopicOrBatchDeleteTopicResponse, error) {
	requestDef := GenReqDefForCreateTopicOrBatchDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTopicOrBatchDeleteTopicResponse), nil
	}
}

// CreateTopicOrBatchDeleteTopicInvoker 创建主题或批量删除主题
func (c *RocketMQClient) CreateTopicOrBatchDeleteTopicInvoker(request *model.CreateTopicOrBatchDeleteTopicRequest) *CreateTopicOrBatchDeleteTopicInvoker {
	requestDef := GenReqDefForCreateTopicOrBatchDeleteTopic()
	return &CreateTopicOrBatchDeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTopic 删除指定主题
//
// 删除指定主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) DeleteTopic(request *model.DeleteTopicRequest) (*model.DeleteTopicResponse, error) {
	requestDef := GenReqDefForDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicResponse), nil
	}
}

// DeleteTopicInvoker 删除指定主题
func (c *RocketMQClient) DeleteTopicInvoker(request *model.DeleteTopicRequest) *DeleteTopicInvoker {
	requestDef := GenReqDefForDeleteTopic()
	return &DeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsumerGroupOfTopic 查询主题消费组列表
//
// 查询主题消费组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListConsumerGroupOfTopic(request *model.ListConsumerGroupOfTopicRequest) (*model.ListConsumerGroupOfTopicResponse, error) {
	requestDef := GenReqDefForListConsumerGroupOfTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumerGroupOfTopicResponse), nil
	}
}

// ListConsumerGroupOfTopicInvoker 查询主题消费组列表
func (c *RocketMQClient) ListConsumerGroupOfTopicInvoker(request *model.ListConsumerGroupOfTopicRequest) *ListConsumerGroupOfTopicInvoker {
	requestDef := GenReqDefForListConsumerGroupOfTopic()
	return &ListConsumerGroupOfTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRocketInstanceTopics 查询主题列表
//
// 该接口用于查询指定RocketMQ实例的Topic列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ListRocketInstanceTopics(request *model.ListRocketInstanceTopicsRequest) (*model.ListRocketInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListRocketInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRocketInstanceTopicsResponse), nil
	}
}

// ListRocketInstanceTopicsInvoker 查询主题列表
func (c *RocketMQClient) ListRocketInstanceTopicsInvoker(request *model.ListRocketInstanceTopicsRequest) *ListRocketInstanceTopicsInvoker {
	requestDef := GenReqDefForListRocketInstanceTopics()
	return &ListRocketInstanceTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOneTopic 查询单个主题
//
// 查询单个主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowOneTopic(request *model.ShowOneTopicRequest) (*model.ShowOneTopicResponse, error) {
	requestDef := GenReqDefForShowOneTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOneTopicResponse), nil
	}
}

// ShowOneTopicInvoker 查询单个主题
func (c *RocketMQClient) ShowOneTopicInvoker(request *model.ShowOneTopicRequest) *ShowOneTopicInvoker {
	requestDef := GenReqDefForShowOneTopic()
	return &ShowOneTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopicStatus 查询主题的消息数
//
// 查询主题的消息数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) ShowTopicStatus(request *model.ShowTopicStatusRequest) (*model.ShowTopicStatusResponse, error) {
	requestDef := GenReqDefForShowTopicStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicStatusResponse), nil
	}
}

// ShowTopicStatusInvoker 查询主题的消息数
func (c *RocketMQClient) ShowTopicStatusInvoker(request *model.ShowTopicStatusRequest) *ShowTopicStatusInvoker {
	requestDef := GenReqDefForShowTopicStatus()
	return &ShowTopicStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopic 修改主题
//
// 修改主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RocketMQClient) UpdateTopic(request *model.UpdateTopicRequest) (*model.UpdateTopicResponse, error) {
	requestDef := GenReqDefForUpdateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicResponse), nil
	}
}

// UpdateTopicInvoker 修改主题
func (c *RocketMQClient) UpdateTopicInvoker(request *model.UpdateTopicRequest) *UpdateTopicInvoker {
	requestDef := GenReqDefForUpdateTopic()
	return &UpdateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
