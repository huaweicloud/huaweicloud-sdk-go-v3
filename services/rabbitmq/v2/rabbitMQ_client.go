package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rabbitmq/v2/model"
)

type RabbitMQClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRabbitMQClient(hcClient *httpclient.HcHttpClient) *RabbitMQClient {
	return &RabbitMQClient{HcClient: hcClient}
}

func RabbitMQClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateOrDeleteRabbitMqTag 批量添加或删除实例标签
//
// 批量添加或删除实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) BatchCreateOrDeleteRabbitMqTag(request *model.BatchCreateOrDeleteRabbitMqTagRequest) (*model.BatchCreateOrDeleteRabbitMqTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteRabbitMqTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteRabbitMqTagResponse), nil
	}
}

// BatchCreateOrDeleteRabbitMqTagInvoker 批量添加或删除实例标签
func (c *RabbitMQClient) BatchCreateOrDeleteRabbitMqTagInvoker(request *model.BatchCreateOrDeleteRabbitMqTagRequest) *BatchCreateOrDeleteRabbitMqTagInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteRabbitMqTag()
	return &BatchCreateOrDeleteRabbitMqTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRestartOrDeleteInstances 批量删除实例
//
// 批量删除实例。
//
// 实例删除后，实例中原有的数据将被删除，且没有备份，请谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) BatchRestartOrDeleteInstances(request *model.BatchRestartOrDeleteInstancesRequest) (*model.BatchRestartOrDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestartOrDeleteInstancesResponse), nil
	}
}

// BatchRestartOrDeleteInstancesInvoker 批量删除实例
func (c *RabbitMQClient) BatchRestartOrDeleteInstancesInvoker(request *model.BatchRestartOrDeleteInstancesRequest) *BatchRestartOrDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()
	return &BatchRestartOrDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidInstance 创建实例(按需)
//
// 创建实例，该接口创建的实例为按需计费的方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

// CreatePostPaidInstanceInvoker 创建实例(按需)
func (c *RabbitMQClient) CreatePostPaidInstanceInvoker(request *model.CreatePostPaidInstanceRequest) *CreatePostPaidInstanceInvoker {
	requestDef := GenReqDefForCreatePostPaidInstance()
	return &CreatePostPaidInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidInstanceByEngine 创建实例
//
// 创建实例[，该接口支持创建按需[和包周期](tag:hws,hws_eu,hws_hk,ctc,cmcc)计费方式的实例](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,cmcc)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreatePostPaidInstanceByEngine(request *model.CreatePostPaidInstanceByEngineRequest) (*model.CreatePostPaidInstanceByEngineResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstanceByEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceByEngineResponse), nil
	}
}

// CreatePostPaidInstanceByEngineInvoker 创建实例
func (c *RabbitMQClient) CreatePostPaidInstanceByEngineInvoker(request *model.CreatePostPaidInstanceByEngineRequest) *CreatePostPaidInstanceByEngineInvoker {
	requestDef := GenReqDefForCreatePostPaidInstanceByEngine()
	return &CreatePostPaidInstanceByEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 创建用户（仅AMQP版本支持）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *RabbitMQClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackgroundTask 删除后台任务管理中的指定记录
//
// 删除后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// DeleteBackgroundTaskInvoker 删除后台任务管理中的指定记录
func (c *RabbitMQClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除指定的实例
//
// 删除指定的实例，释放该实例的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除指定的实例
func (c *RabbitMQClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 删除用户（仅AMQP版本支持）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *RabbitMQClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 查询可用区信息
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 查询可用区信息
func (c *RabbitMQClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackgroundTasks 查询实例的后台任务列表
//
// 查询实例的后台任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListBackgroundTasks(request *model.ListBackgroundTasksRequest) (*model.ListBackgroundTasksResponse, error) {
	requestDef := GenReqDefForListBackgroundTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTasksResponse), nil
	}
}

// ListBackgroundTasksInvoker 查询实例的后台任务列表
func (c *RabbitMQClient) ListBackgroundTasksInvoker(request *model.ListBackgroundTasksRequest) *ListBackgroundTasksInvoker {
	requestDef := GenReqDefForListBackgroundTasks()
	return &ListBackgroundTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngineProducts 查询产品规格列表
//
// 查询产品规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListEngineProducts(request *model.ListEngineProductsRequest) (*model.ListEngineProductsResponse, error) {
	requestDef := GenReqDefForListEngineProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEngineProductsResponse), nil
	}
}

// ListEngineProductsInvoker 查询产品规格列表
func (c *RabbitMQClient) ListEngineProductsInvoker(request *model.ListEngineProductsRequest) *ListEngineProductsInvoker {
	requestDef := GenReqDefForListEngineProducts()
	return &ListEngineProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesDetails 查询所有实例列表
//
// 查询租户的实例列表，支持按照条件查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListInstancesDetails(request *model.ListInstancesDetailsRequest) (*model.ListInstancesDetailsResponse, error) {
	requestDef := GenReqDefForListInstancesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesDetailsResponse), nil
	}
}

// ListInstancesDetailsInvoker 查询所有实例列表
func (c *RabbitMQClient) ListInstancesDetailsInvoker(request *model.ListInstancesDetailsRequest) *ListInstancesDetailsInvoker {
	requestDef := GenReqDefForListInstancesDetails()
	return &ListInstancesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlugins 查询插件列表
//
// 查询插件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListPlugins(request *model.ListPluginsRequest) (*model.ListPluginsResponse, error) {
	requestDef := GenReqDefForListPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginsResponse), nil
	}
}

// ListPluginsInvoker 查询插件列表
func (c *RabbitMQClient) ListPluginsInvoker(request *model.ListPluginsRequest) *ListPluginsInvoker {
	requestDef := GenReqDefForListPlugins()
	return &ListPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品规格列表
//
// 在创建实例时，需要配置订购的产品ID（即product_id），可通过该接口查询产品规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询产品规格列表
func (c *RabbitMQClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUser 查询用户列表
//
// 查询用户列表（仅AMQP版本支持）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

// ListUserInvoker 查询用户列表
func (c *RabbitMQClient) ListUserInvoker(request *model.ListUserRequest) *ListUserInvoker {
	requestDef := GenReqDefForListUser()
	return &ListUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 重置密码
//
// 重置密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 重置密码
func (c *RabbitMQClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeEngineInstance 新规格实例的规格变更
//
// 实例规格变更。
//
// [**当前通过调用API，只支持按需实例进行实例规格变更。**](tag:hws,hws_hk,ctc,cmcc,hws_eu)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ResizeEngineInstance(request *model.ResizeEngineInstanceRequest) (*model.ResizeEngineInstanceResponse, error) {
	requestDef := GenReqDefForResizeEngineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeEngineInstanceResponse), nil
	}
}

// ResizeEngineInstanceInvoker 新规格实例的规格变更
func (c *RabbitMQClient) ResizeEngineInstanceInvoker(request *model.ResizeEngineInstanceRequest) *ResizeEngineInstanceInvoker {
	requestDef := GenReqDefForResizeEngineInstance()
	return &ResizeEngineInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 实例规格变更
//
// 实例规格变更。
//
// [**当前通过调用API，只支持按需实例进行实例规格变更。**](tag:hws,hws_hk,ctc,cmcc,hws_eu)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 实例规格变更
func (c *RabbitMQClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackgroundTask 查询后台任务管理中的指定记录
//
// 查询后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowBackgroundTask(request *model.ShowBackgroundTaskRequest) (*model.ShowBackgroundTaskResponse, error) {
	requestDef := GenReqDefForShowBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskResponse), nil
	}
}

// ShowBackgroundTaskInvoker 查询后台任务管理中的指定记录
func (c *RabbitMQClient) ShowBackgroundTaskInvoker(request *model.ShowBackgroundTaskRequest) *ShowBackgroundTaskInvoker {
	requestDef := GenReqDefForShowBackgroundTask()
	return &ShowBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCesHierarchy 查询实例在CES的监控层级关系
//
// 查询实例在CES的监控层级关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowCesHierarchy(request *model.ShowCesHierarchyRequest) (*model.ShowCesHierarchyResponse, error) {
	requestDef := GenReqDefForShowCesHierarchy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCesHierarchyResponse), nil
	}
}

// ShowCesHierarchyInvoker 查询实例在CES的监控层级关系
func (c *RabbitMQClient) ShowCesHierarchyInvoker(request *model.ShowCesHierarchyRequest) *ShowCesHierarchyInvoker {
	requestDef := GenReqDefForShowCesHierarchy()
	return &ShowCesHierarchyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineInstanceExtendProductInfo 查询新规格可扩容规格列表
//
// 查询新规格实例可扩容列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowEngineInstanceExtendProductInfo(request *model.ShowEngineInstanceExtendProductInfoRequest) (*model.ShowEngineInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineInstanceExtendProductInfoResponse), nil
	}
}

// ShowEngineInstanceExtendProductInfoInvoker 查询新规格可扩容规格列表
func (c *RabbitMQClient) ShowEngineInstanceExtendProductInfoInvoker(request *model.ShowEngineInstanceExtendProductInfoRequest) *ShowEngineInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowEngineInstanceExtendProductInfo()
	return &ShowEngineInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询指定实例
//
// 查询指定实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询指定实例
func (c *RabbitMQClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceExtendProductInfo 查询可扩容规格列表
//
// 查询可扩容规格列表。
//
// RabbtiMQ只支持只增加节点数的扩容方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowInstanceExtendProductInfo(request *model.ShowInstanceExtendProductInfoRequest) (*model.ShowInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceExtendProductInfoResponse), nil
	}
}

// ShowInstanceExtendProductInfoInvoker 查询可扩容规格列表
func (c *RabbitMQClient) ShowInstanceExtendProductInfoInvoker(request *model.ShowInstanceExtendProductInfoRequest) *ShowInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()
	return &ShowInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMaintainWindows 查询维护时间窗时间段
//
// 查询维护时间窗开始时间和结束时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowMaintainWindows(request *model.ShowMaintainWindowsRequest) (*model.ShowMaintainWindowsResponse, error) {
	requestDef := GenReqDefForShowMaintainWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMaintainWindowsResponse), nil
	}
}

// ShowMaintainWindowsInvoker 查询维护时间窗时间段
func (c *RabbitMQClient) ShowMaintainWindowsInvoker(request *model.ShowMaintainWindowsRequest) *ShowMaintainWindowsInvoker {
	requestDef := GenReqDefForShowMaintainWindows()
	return &ShowMaintainWindowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRabbitMqProjectTags 查询项目标签
//
// 查询项目标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowRabbitMqProjectTags(request *model.ShowRabbitMqProjectTagsRequest) (*model.ShowRabbitMqProjectTagsResponse, error) {
	requestDef := GenReqDefForShowRabbitMqProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRabbitMqProjectTagsResponse), nil
	}
}

// ShowRabbitMqProjectTagsInvoker 查询项目标签
func (c *RabbitMQClient) ShowRabbitMqProjectTagsInvoker(request *model.ShowRabbitMqProjectTagsRequest) *ShowRabbitMqProjectTagsInvoker {
	requestDef := GenReqDefForShowRabbitMqProjectTags()
	return &ShowRabbitMqProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRabbitMqTags 查询实例标签
//
// 查询实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowRabbitMqTags(request *model.ShowRabbitMqTagsRequest) (*model.ShowRabbitMqTagsResponse, error) {
	requestDef := GenReqDefForShowRabbitMqTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRabbitMqTagsResponse), nil
	}
}

// ShowRabbitMqTagsInvoker 查询实例标签
func (c *RabbitMQClient) ShowRabbitMqTagsInvoker(request *model.ShowRabbitMqTagsRequest) *ShowRabbitMqTagsInvoker {
	requestDef := GenReqDefForShowRabbitMqTags()
	return &ShowRabbitMqTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改实例信息
//
// 修改实例的名称和描述信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改实例信息
func (c *RabbitMQClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlugins 开启或关闭插件
//
// 开启或关闭插件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) UpdatePlugins(request *model.UpdatePluginsRequest) (*model.UpdatePluginsResponse, error) {
	requestDef := GenReqDefForUpdatePlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePluginsResponse), nil
	}
}

// UpdatePluginsInvoker 开启或关闭插件
func (c *RabbitMQClient) UpdatePluginsInvoker(request *model.UpdatePluginsRequest) *UpdatePluginsInvoker {
	requestDef := GenReqDefForUpdatePlugins()
	return &UpdatePluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户参数
//
// 修改用户参数（仅AMQP版本支持）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改用户参数
func (c *RabbitMQClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBinding 添加绑定
//
// 添加绑定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreateBinding(request *model.CreateBindingRequest) (*model.CreateBindingResponse, error) {
	requestDef := GenReqDefForCreateBinding()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBindingResponse), nil
	}
}

// CreateBindingInvoker 添加绑定
func (c *RabbitMQClient) CreateBindingInvoker(request *model.CreateBindingRequest) *CreateBindingInvoker {
	requestDef := GenReqDefForCreateBinding()
	return &CreateBindingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBinding 删除绑定
//
// 删除绑定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) DeleteBinding(request *model.DeleteBindingRequest) (*model.DeleteBindingResponse, error) {
	requestDef := GenReqDefForDeleteBinding()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBindingResponse), nil
	}
}

// DeleteBindingInvoker 删除绑定
func (c *RabbitMQClient) DeleteBindingInvoker(request *model.DeleteBindingRequest) *DeleteBindingInvoker {
	requestDef := GenReqDefForDeleteBinding()
	return &DeleteBindingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBindings 查询Exchange绑定信息列表
//
// 查询Exchange绑定信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListBindings(request *model.ListBindingsRequest) (*model.ListBindingsResponse, error) {
	requestDef := GenReqDefForListBindings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBindingsResponse), nil
	}
}

// ListBindingsInvoker 查询Exchange绑定信息列表
func (c *RabbitMQClient) ListBindingsInvoker(request *model.ListBindingsRequest) *ListBindingsInvoker {
	requestDef := GenReqDefForListBindings()
	return &ListBindingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteExchanges 批量删除指定Exchange
//
// 批量删除指定Exchange。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) BatchDeleteExchanges(request *model.BatchDeleteExchangesRequest) (*model.BatchDeleteExchangesResponse, error) {
	requestDef := GenReqDefForBatchDeleteExchanges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteExchangesResponse), nil
	}
}

// BatchDeleteExchangesInvoker 批量删除指定Exchange
func (c *RabbitMQClient) BatchDeleteExchangesInvoker(request *model.BatchDeleteExchangesRequest) *BatchDeleteExchangesInvoker {
	requestDef := GenReqDefForBatchDeleteExchanges()
	return &BatchDeleteExchangesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExchange 创建Exchange
//
// 创建Exchange。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreateExchange(request *model.CreateExchangeRequest) (*model.CreateExchangeResponse, error) {
	requestDef := GenReqDefForCreateExchange()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExchangeResponse), nil
	}
}

// CreateExchangeInvoker 创建Exchange
func (c *RabbitMQClient) CreateExchangeInvoker(request *model.CreateExchangeRequest) *CreateExchangeInvoker {
	requestDef := GenReqDefForCreateExchange()
	return &CreateExchangeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExchanges 查询Exchange列表
//
// 查询Exchange列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListExchanges(request *model.ListExchangesRequest) (*model.ListExchangesResponse, error) {
	requestDef := GenReqDefForListExchanges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExchangesResponse), nil
	}
}

// ListExchangesInvoker 查询Exchange列表
func (c *RabbitMQClient) ListExchangesInvoker(request *model.ListExchangesRequest) *ListExchangesInvoker {
	requestDef := GenReqDefForListExchanges()
	return &ListExchangesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteQueues 批量删除指定Queue
//
// 批量删除指定Queue。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) BatchDeleteQueues(request *model.BatchDeleteQueuesRequest) (*model.BatchDeleteQueuesResponse, error) {
	requestDef := GenReqDefForBatchDeleteQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteQueuesResponse), nil
	}
}

// BatchDeleteQueuesInvoker 批量删除指定Queue
func (c *RabbitMQClient) BatchDeleteQueuesInvoker(request *model.BatchDeleteQueuesRequest) *BatchDeleteQueuesInvoker {
	requestDef := GenReqDefForBatchDeleteQueues()
	return &BatchDeleteQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQueue 创建Queue
//
// 创建Queue。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreateQueue(request *model.CreateQueueRequest) (*model.CreateQueueResponse, error) {
	requestDef := GenReqDefForCreateQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQueueResponse), nil
	}
}

// CreateQueueInvoker 创建Queue
func (c *RabbitMQClient) CreateQueueInvoker(request *model.CreateQueueRequest) *CreateQueueInvoker {
	requestDef := GenReqDefForCreateQueue()
	return &CreateQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQueueInfo 清空Queue消息
//
// 清空Queue消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) DeleteQueueInfo(request *model.DeleteQueueInfoRequest) (*model.DeleteQueueInfoResponse, error) {
	requestDef := GenReqDefForDeleteQueueInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueueInfoResponse), nil
	}
}

// DeleteQueueInfoInvoker 清空Queue消息
func (c *RabbitMQClient) DeleteQueueInfoInvoker(request *model.DeleteQueueInfoRequest) *DeleteQueueInfoInvoker {
	requestDef := GenReqDefForDeleteQueueInfo()
	return &DeleteQueueInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueues 查询所属Vhost下Queue的列表
//
// 查询所属Vhost下Queue的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListQueues(request *model.ListQueuesRequest) (*model.ListQueuesResponse, error) {
	requestDef := GenReqDefForListQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueuesResponse), nil
	}
}

// ListQueuesInvoker 查询所属Vhost下Queue的列表
func (c *RabbitMQClient) ListQueuesInvoker(request *model.ListQueuesRequest) *ListQueuesInvoker {
	requestDef := GenReqDefForListQueues()
	return &ListQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQueueDetails 查询指定Queue详情
//
// 查询指定Queue详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ShowQueueDetails(request *model.ShowQueueDetailsRequest) (*model.ShowQueueDetailsResponse, error) {
	requestDef := GenReqDefForShowQueueDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueueDetailsResponse), nil
	}
}

// ShowQueueDetailsInvoker 查询指定Queue详情
func (c *RabbitMQClient) ShowQueueDetailsInvoker(request *model.ShowQueueDetailsRequest) *ShowQueueDetailsInvoker {
	requestDef := GenReqDefForShowQueueDetails()
	return &ShowQueueDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteVhosts 批量删除指定Vhost
//
// 批量删除指定Vhost。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) BatchDeleteVhosts(request *model.BatchDeleteVhostsRequest) (*model.BatchDeleteVhostsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVhosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVhostsResponse), nil
	}
}

// BatchDeleteVhostsInvoker 批量删除指定Vhost
func (c *RabbitMQClient) BatchDeleteVhostsInvoker(request *model.BatchDeleteVhostsRequest) *BatchDeleteVhostsInvoker {
	requestDef := GenReqDefForBatchDeleteVhosts()
	return &BatchDeleteVhostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVhost 创建Vhost
//
// 创建Vhost。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) CreateVhost(request *model.CreateVhostRequest) (*model.CreateVhostResponse, error) {
	requestDef := GenReqDefForCreateVhost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVhostResponse), nil
	}
}

// CreateVhostInvoker 创建Vhost
func (c *RabbitMQClient) CreateVhostInvoker(request *model.CreateVhostRequest) *CreateVhostInvoker {
	requestDef := GenReqDefForCreateVhost()
	return &CreateVhostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVhosts 查询Vhost列表
//
// 查询Vhost列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RabbitMQClient) ListVhosts(request *model.ListVhostsRequest) (*model.ListVhostsResponse, error) {
	requestDef := GenReqDefForListVhosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVhostsResponse), nil
	}
}

// ListVhostsInvoker 查询Vhost列表
func (c *RabbitMQClient) ListVhostsInvoker(request *model.ListVhostsRequest) *ListVhostsInvoker {
	requestDef := GenReqDefForListVhosts()
	return &ListVhostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
