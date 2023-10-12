package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1/model"
)

type AsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAsClient(hcClient *http_client.HcHttpClient) *AsClient {
	return &AsClient{HcClient: hcClient}
}

func AsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AttachCallbackInstanceLifeCycleHook 伸缩实例生命周期回调
//
// 通过生命周期操作令牌或者通过实例ID和生命周期挂钩名称对伸缩实例指定的挂钩进行回调操作。如果在超时时间结束前已完成自定义操作，选择终止或继续完成生命周期操作。如果需要更多时间完成自定义操作，选择延长超时时间，实例保持等待状态的时间将增加1小时。只有实例的生命周期挂钩状态为 HANGING 时才可以进行回调操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) AttachCallbackInstanceLifeCycleHook(request *model.AttachCallbackInstanceLifeCycleHookRequest) (*model.AttachCallbackInstanceLifeCycleHookResponse, error) {
	requestDef := GenReqDefForAttachCallbackInstanceLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachCallbackInstanceLifeCycleHookResponse), nil
	}
}

// AttachCallbackInstanceLifeCycleHookInvoker 伸缩实例生命周期回调
func (c *AsClient) AttachCallbackInstanceLifeCycleHookInvoker(request *model.AttachCallbackInstanceLifeCycleHookRequest) *AttachCallbackInstanceLifeCycleHookInvoker {
	requestDef := GenReqDefForAttachCallbackInstanceLifeCycleHook()
	return &AttachCallbackInstanceLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddScalingInstances 批量添加实例
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
// 说明：
// - 单次最多批量操作实例个数为10。批量添加后实例数不能大于伸缩组的最大实例数，批量移出后实例数不能小于伸缩组的最小实例数。
// - 当伸缩组处于INSERVICE状态且没有伸缩活动时，才能添加实例。
// - 当伸缩组没有伸缩活动时，才能移出实例。
// - 向伸缩组中添加实例时，必须保证实例所在的可用区包含于伸缩组的可用区内。
// - 实例处于INSERVICE状态时才可以进行移出、设置或取消实例保护属性等操作。
// - 当伸缩组发生自动缩容活动时，设置了实例保护的实例不会被移出伸缩组。
// - 批量移出弹性伸缩组中的实例时，若该实例加入伸缩组时绑定的监听器和伸缩组本身的监听器相同，会解绑定实例和监听器。若该实例加入伸缩组时绑定的监听器和伸缩组本身的监听器不同，会保留实例和监听器的绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchAddScalingInstances(request *model.BatchAddScalingInstancesRequest) (*model.BatchAddScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchAddScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddScalingInstancesResponse), nil
	}
}

// BatchAddScalingInstancesInvoker 批量添加实例
func (c *AsClient) BatchAddScalingInstancesInvoker(request *model.BatchAddScalingInstancesRequest) *BatchAddScalingInstancesInvoker {
	requestDef := GenReqDefForBatchAddScalingInstances()
	return &BatchAddScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteScalingConfigs 批量删除弹性伸缩配置
//
// 批量删除指定弹性伸缩配置。被伸缩组使用的伸缩配置不能被删除。单次最多删除伸缩配置个数为50。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchDeleteScalingConfigs(request *model.BatchDeleteScalingConfigsRequest) (*model.BatchDeleteScalingConfigsResponse, error) {
	requestDef := GenReqDefForBatchDeleteScalingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScalingConfigsResponse), nil
	}
}

// BatchDeleteScalingConfigsInvoker 批量删除弹性伸缩配置
func (c *AsClient) BatchDeleteScalingConfigsInvoker(request *model.BatchDeleteScalingConfigsRequest) *BatchDeleteScalingConfigsInvoker {
	requestDef := GenReqDefForBatchDeleteScalingConfigs()
	return &BatchDeleteScalingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteScalingPolicies 批量删除弹性伸缩策略。
//
// 批量启用、停用或者删除弹性伸缩策略。单次最多批量操作伸缩策略个数为20。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchDeleteScalingPolicies(request *model.BatchDeleteScalingPoliciesRequest) (*model.BatchDeleteScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchDeleteScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScalingPoliciesResponse), nil
	}
}

// BatchDeleteScalingPoliciesInvoker 批量删除弹性伸缩策略。
func (c *AsClient) BatchDeleteScalingPoliciesInvoker(request *model.BatchDeleteScalingPoliciesRequest) *BatchDeleteScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchDeleteScalingPolicies()
	return &BatchDeleteScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchPauseScalingPolicies 批量停用弹性伸缩策略。
//
// 批量启用、停用或者删除弹性伸缩策略。单次最多批量操作伸缩策略个数为20。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchPauseScalingPolicies(request *model.BatchPauseScalingPoliciesRequest) (*model.BatchPauseScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchPauseScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPauseScalingPoliciesResponse), nil
	}
}

// BatchPauseScalingPoliciesInvoker 批量停用弹性伸缩策略。
func (c *AsClient) BatchPauseScalingPoliciesInvoker(request *model.BatchPauseScalingPoliciesRequest) *BatchPauseScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchPauseScalingPolicies()
	return &BatchPauseScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchProtectScalingInstances 批量设置实例保护
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchProtectScalingInstances(request *model.BatchProtectScalingInstancesRequest) (*model.BatchProtectScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchProtectScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchProtectScalingInstancesResponse), nil
	}
}

// BatchProtectScalingInstancesInvoker 批量设置实例保护
func (c *AsClient) BatchProtectScalingInstancesInvoker(request *model.BatchProtectScalingInstancesRequest) *BatchProtectScalingInstancesInvoker {
	requestDef := GenReqDefForBatchProtectScalingInstances()
	return &BatchProtectScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveScalingInstances 批量移除实例
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchRemoveScalingInstances(request *model.BatchRemoveScalingInstancesRequest) (*model.BatchRemoveScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchRemoveScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveScalingInstancesResponse), nil
	}
}

// BatchRemoveScalingInstancesInvoker 批量移除实例
func (c *AsClient) BatchRemoveScalingInstancesInvoker(request *model.BatchRemoveScalingInstancesRequest) *BatchRemoveScalingInstancesInvoker {
	requestDef := GenReqDefForBatchRemoveScalingInstances()
	return &BatchRemoveScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResumeScalingPolicies 批量启用弹性伸缩策略。
//
// 批量启用、停用或者删除弹性伸缩策略。单次最多批量操作伸缩策略个数为20。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchResumeScalingPolicies(request *model.BatchResumeScalingPoliciesRequest) (*model.BatchResumeScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchResumeScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResumeScalingPoliciesResponse), nil
	}
}

// BatchResumeScalingPoliciesInvoker 批量启用弹性伸缩策略。
func (c *AsClient) BatchResumeScalingPoliciesInvoker(request *model.BatchResumeScalingPoliciesRequest) *BatchResumeScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchResumeScalingPolicies()
	return &BatchResumeScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetScalingInstancesStandby 批量将实例转为备用状态
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchSetScalingInstancesStandby(request *model.BatchSetScalingInstancesStandbyRequest) (*model.BatchSetScalingInstancesStandbyResponse, error) {
	requestDef := GenReqDefForBatchSetScalingInstancesStandby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetScalingInstancesStandbyResponse), nil
	}
}

// BatchSetScalingInstancesStandbyInvoker 批量将实例转为备用状态
func (c *AsClient) BatchSetScalingInstancesStandbyInvoker(request *model.BatchSetScalingInstancesStandbyRequest) *BatchSetScalingInstancesStandbyInvoker {
	requestDef := GenReqDefForBatchSetScalingInstancesStandby()
	return &BatchSetScalingInstancesStandbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUnprotectScalingInstances 批量取消实例保护
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchUnprotectScalingInstances(request *model.BatchUnprotectScalingInstancesRequest) (*model.BatchUnprotectScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchUnprotectScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnprotectScalingInstancesResponse), nil
	}
}

// BatchUnprotectScalingInstancesInvoker 批量取消实例保护
func (c *AsClient) BatchUnprotectScalingInstancesInvoker(request *model.BatchUnprotectScalingInstancesRequest) *BatchUnprotectScalingInstancesInvoker {
	requestDef := GenReqDefForBatchUnprotectScalingInstances()
	return &BatchUnprotectScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUnsetScalingInstancesStantby 批量将实例移出备用状态
//
// 批量移出伸缩组中的实例或批量添加伸缩组外的实例。批量对伸缩组中的实例设置或取消其实例保护属性。批量将伸缩组中的实例转入或移出备用状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) BatchUnsetScalingInstancesStantby(request *model.BatchUnsetScalingInstancesStantbyRequest) (*model.BatchUnsetScalingInstancesStantbyResponse, error) {
	requestDef := GenReqDefForBatchUnsetScalingInstancesStantby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnsetScalingInstancesStantbyResponse), nil
	}
}

// BatchUnsetScalingInstancesStantbyInvoker 批量将实例移出备用状态
func (c *AsClient) BatchUnsetScalingInstancesStantbyInvoker(request *model.BatchUnsetScalingInstancesStantbyRequest) *BatchUnsetScalingInstancesStantbyInvoker {
	requestDef := GenReqDefForBatchUnsetScalingInstancesStantby()
	return &BatchUnsetScalingInstancesStantbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupScheduledTask 创建计划任务
//
// 创建计划任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateGroupScheduledTask(request *model.CreateGroupScheduledTaskRequest) (*model.CreateGroupScheduledTaskResponse, error) {
	requestDef := GenReqDefForCreateGroupScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupScheduledTaskResponse), nil
	}
}

// CreateGroupScheduledTaskInvoker 创建计划任务
func (c *AsClient) CreateGroupScheduledTaskInvoker(request *model.CreateGroupScheduledTaskRequest) *CreateGroupScheduledTaskInvoker {
	requestDef := GenReqDefForCreateGroupScheduledTask()
	return &CreateGroupScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLifyCycleHook 创建生命周期挂钩
//
// 创建生命周期挂钩，可为伸缩组添加一个或多个生命周期挂钩，最多添加5个。添加生命周期挂钩后，当伸缩组进行伸缩活动时，实例将被生命周期挂钩挂起并置于等待状态（正在加入伸缩组或正在移出伸缩组），实例将保持此状态直至超时时间结束或者用户手动回调。用户能够在实例保持等待状态的时间段内执行自定义操作，例如，用户可以在新启动的实例上安装或配置软件，也可以在实例终止前从实例中下载日志文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateLifyCycleHook(request *model.CreateLifyCycleHookRequest) (*model.CreateLifyCycleHookResponse, error) {
	requestDef := GenReqDefForCreateLifyCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLifyCycleHookResponse), nil
	}
}

// CreateLifyCycleHookInvoker 创建生命周期挂钩
func (c *AsClient) CreateLifyCycleHookInvoker(request *model.CreateLifyCycleHookRequest) *CreateLifyCycleHookInvoker {
	requestDef := GenReqDefForCreateLifyCycleHook()
	return &CreateLifyCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingConfig 创建弹性伸缩配置
//
// 创建弹性伸缩配置。伸缩配置是伸缩组内实例（弹性云服务器云主机）的模板，定义了伸缩组内待添加的实例的规格数据。伸缩配置与伸缩组是解耦的，同一伸缩配置可以被多个伸缩组使用。默认最多可以创建100个伸缩配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingConfig(request *model.CreateScalingConfigRequest) (*model.CreateScalingConfigResponse, error) {
	requestDef := GenReqDefForCreateScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingConfigResponse), nil
	}
}

// CreateScalingConfigInvoker 创建弹性伸缩配置
func (c *AsClient) CreateScalingConfigInvoker(request *model.CreateScalingConfigRequest) *CreateScalingConfigInvoker {
	requestDef := GenReqDefForCreateScalingConfig()
	return &CreateScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingGroup 创建弹性伸缩组
//
// 伸缩组是具有相同应用场景的实例的集合，是启停伸缩策略和进行伸缩活动的基本单位。伸缩组内定义了最大实例数、期望实例数、最小实例数、虚拟私有云、子网、负载均衡等信息。默认最多可以创建10个伸缩组。如果伸缩组配置了负载均衡，在添加或移除实例时，会自动为实例绑定或解绑负载均衡监听器。如果伸缩组使用负载均衡健康检查方式，伸缩组中的实例需要启用负载均衡器的监听端口才能通过健康检查。端口启用可在安全组中进行配置，可参考添加安全组规则进行操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingGroup(request *model.CreateScalingGroupRequest) (*model.CreateScalingGroupResponse, error) {
	requestDef := GenReqDefForCreateScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingGroupResponse), nil
	}
}

// CreateScalingGroupInvoker 创建弹性伸缩组
func (c *AsClient) CreateScalingGroupInvoker(request *model.CreateScalingGroupRequest) *CreateScalingGroupInvoker {
	requestDef := GenReqDefForCreateScalingGroup()
	return &CreateScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingNotification 配置伸缩组通知
//
// 给弹性伸缩组配置通知功能。每调用一次该接口，伸缩组即配置一个通知主题及其通知场景，每个伸缩组最多可以增加5个主题。通知主题由用户事先在SMN创建并进行订阅，当通知主题对应的通知场景出现时，伸缩组会向用户的订阅终端发送通知。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingNotification(request *model.CreateScalingNotificationRequest) (*model.CreateScalingNotificationResponse, error) {
	requestDef := GenReqDefForCreateScalingNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingNotificationResponse), nil
	}
}

// CreateScalingNotificationInvoker 配置伸缩组通知
func (c *AsClient) CreateScalingNotificationInvoker(request *model.CreateScalingNotificationRequest) *CreateScalingNotificationInvoker {
	requestDef := GenReqDefForCreateScalingNotification()
	return &CreateScalingNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingPolicy 创建弹性伸缩策略
//
// 创建弹性伸缩策略。伸缩策略定义了伸缩组内实例的扩张和收缩操作。如果执行伸缩策略造成伸缩组期望实例数与伸缩组内实例数不符，弹性伸缩会自动调整实例资源，以匹配期望实例数。当前伸缩策略支持告警触发策略，周期触发策略，定时触发策略。在策略执行具体动作中，可设置实例变化的个数，或根据当前实例的百分比数进行伸缩。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingPolicy(request *model.CreateScalingPolicyRequest) (*model.CreateScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingPolicyResponse), nil
	}
}

// CreateScalingPolicyInvoker 创建弹性伸缩策略
func (c *AsClient) CreateScalingPolicyInvoker(request *model.CreateScalingPolicyRequest) *CreateScalingPolicyInvoker {
	requestDef := GenReqDefForCreateScalingPolicy()
	return &CreateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingTagInfo 创建标签
//
// 创建或删除指定资源的标签。每个伸缩组最多添加10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingTagInfo(request *model.CreateScalingTagInfoRequest) (*model.CreateScalingTagInfoResponse, error) {
	requestDef := GenReqDefForCreateScalingTagInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingTagInfoResponse), nil
	}
}

// CreateScalingTagInfoInvoker 创建标签
func (c *AsClient) CreateScalingTagInfoInvoker(request *model.CreateScalingTagInfoRequest) *CreateScalingTagInfoInvoker {
	requestDef := GenReqDefForCreateScalingTagInfo()
	return &CreateScalingTagInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupScheduledTask 删除计划任务
//
// 删除计划任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteGroupScheduledTask(request *model.DeleteGroupScheduledTaskRequest) (*model.DeleteGroupScheduledTaskResponse, error) {
	requestDef := GenReqDefForDeleteGroupScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupScheduledTaskResponse), nil
	}
}

// DeleteGroupScheduledTaskInvoker 删除计划任务
func (c *AsClient) DeleteGroupScheduledTaskInvoker(request *model.DeleteGroupScheduledTaskRequest) *DeleteGroupScheduledTaskInvoker {
	requestDef := GenReqDefForDeleteGroupScheduledTask()
	return &DeleteGroupScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLifecycleHook 删除生命周期挂钩
//
// 删除一个指定生命周期挂钩。伸缩组进行伸缩活动时，不允许删除该伸缩组内的生命周期挂钩。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteLifecycleHook(request *model.DeleteLifecycleHookRequest) (*model.DeleteLifecycleHookResponse, error) {
	requestDef := GenReqDefForDeleteLifecycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLifecycleHookResponse), nil
	}
}

// DeleteLifecycleHookInvoker 删除生命周期挂钩
func (c *AsClient) DeleteLifecycleHookInvoker(request *model.DeleteLifecycleHookRequest) *DeleteLifecycleHookInvoker {
	requestDef := GenReqDefForDeleteLifecycleHook()
	return &DeleteLifecycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingConfig 删除弹性伸缩配置
//
// 删除一个指定弹性伸缩配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingConfig(request *model.DeleteScalingConfigRequest) (*model.DeleteScalingConfigResponse, error) {
	requestDef := GenReqDefForDeleteScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingConfigResponse), nil
	}
}

// DeleteScalingConfigInvoker 删除弹性伸缩配置
func (c *AsClient) DeleteScalingConfigInvoker(request *model.DeleteScalingConfigRequest) *DeleteScalingConfigInvoker {
	requestDef := GenReqDefForDeleteScalingConfig()
	return &DeleteScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingGroup 删除弹性伸缩组
//
// 删除一个指定弹性伸缩组。force_delete属性表示如果伸缩组存在ECS实例或正在进行伸缩活动，是否强制删除伸缩组并移出和释放ECS实例。默认值为no，表示不强制删除伸缩组。如果force_delete的值为no，必须满足以下两个条件，才能删除伸缩组：条件一：伸缩组没有正在进行的伸缩活动。条件二：伸缩组当前的ECS实例数量（current_instance_number）为0。如果force_delete的值为yes，伸缩组会被置于DELETING状态，拒绝接收新的伸缩活动请求，然后等待已有的伸缩活动完成，最后将伸缩组内所有ECS实例移出伸缩组（用户手动添加的ECS实例会被移出伸缩组，弹性伸缩自动创建的ECS实例会被自动删除）并删除伸缩组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingGroup(request *model.DeleteScalingGroupRequest) (*model.DeleteScalingGroupResponse, error) {
	requestDef := GenReqDefForDeleteScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingGroupResponse), nil
	}
}

// DeleteScalingGroupInvoker 删除弹性伸缩组
func (c *AsClient) DeleteScalingGroupInvoker(request *model.DeleteScalingGroupRequest) *DeleteScalingGroupInvoker {
	requestDef := GenReqDefForDeleteScalingGroup()
	return &DeleteScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingInstance 移出弹性伸缩组实例
//
// 从弹性伸缩组中移出一个指定实例。实例处于INSERVICE且移出后实例数不能小于伸缩组的最小实例数时才可以移出。当伸缩组没有伸缩活动时，才能移出实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingInstance(request *model.DeleteScalingInstanceRequest) (*model.DeleteScalingInstanceResponse, error) {
	requestDef := GenReqDefForDeleteScalingInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingInstanceResponse), nil
	}
}

// DeleteScalingInstanceInvoker 移出弹性伸缩组实例
func (c *AsClient) DeleteScalingInstanceInvoker(request *model.DeleteScalingInstanceRequest) *DeleteScalingInstanceInvoker {
	requestDef := GenReqDefForDeleteScalingInstance()
	return &DeleteScalingInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingNotification 删除伸缩组通知
//
// 删除指定的弹性伸缩组中指定的通知。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingNotification(request *model.DeleteScalingNotificationRequest) (*model.DeleteScalingNotificationResponse, error) {
	requestDef := GenReqDefForDeleteScalingNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingNotificationResponse), nil
	}
}

// DeleteScalingNotificationInvoker 删除伸缩组通知
func (c *AsClient) DeleteScalingNotificationInvoker(request *model.DeleteScalingNotificationRequest) *DeleteScalingNotificationInvoker {
	requestDef := GenReqDefForDeleteScalingNotification()
	return &DeleteScalingNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingPolicy 删除弹性伸缩策略
//
// 删除一个指定弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingPolicy(request *model.DeleteScalingPolicyRequest) (*model.DeleteScalingPolicyResponse, error) {
	requestDef := GenReqDefForDeleteScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingPolicyResponse), nil
	}
}

// DeleteScalingPolicyInvoker 删除弹性伸缩策略
func (c *AsClient) DeleteScalingPolicyInvoker(request *model.DeleteScalingPolicyRequest) *DeleteScalingPolicyInvoker {
	requestDef := GenReqDefForDeleteScalingPolicy()
	return &DeleteScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingTagInfo 删除标签
//
// 创建或删除指定资源的标签。每个伸缩组最多添加10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) DeleteScalingTagInfo(request *model.DeleteScalingTagInfoRequest) (*model.DeleteScalingTagInfoResponse, error) {
	requestDef := GenReqDefForDeleteScalingTagInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingTagInfoResponse), nil
	}
}

// DeleteScalingTagInfoInvoker 删除标签
func (c *AsClient) DeleteScalingTagInfoInvoker(request *model.DeleteScalingTagInfoRequest) *DeleteScalingTagInfoInvoker {
	requestDef := GenReqDefForDeleteScalingTagInfo()
	return &DeleteScalingTagInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteScalingPolicy 执行弹性伸缩策略。
//
// 立即执行或启用或停止一个指定弹性伸缩策略。当伸缩组、伸缩策略状态处于INSERVICE时，伸缩策略才能被正确执行，否则会执行失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ExecuteScalingPolicy(request *model.ExecuteScalingPolicyRequest) (*model.ExecuteScalingPolicyResponse, error) {
	requestDef := GenReqDefForExecuteScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScalingPolicyResponse), nil
	}
}

// ExecuteScalingPolicyInvoker 执行弹性伸缩策略。
func (c *AsClient) ExecuteScalingPolicyInvoker(request *model.ExecuteScalingPolicyRequest) *ExecuteScalingPolicyInvoker {
	requestDef := GenReqDefForExecuteScalingPolicy()
	return &ExecuteScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupScheduledTasks 查询计划任务列表
//
// 查询计划任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListGroupScheduledTasks(request *model.ListGroupScheduledTasksRequest) (*model.ListGroupScheduledTasksResponse, error) {
	requestDef := GenReqDefForListGroupScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupScheduledTasksResponse), nil
	}
}

// ListGroupScheduledTasksInvoker 查询计划任务列表
func (c *AsClient) ListGroupScheduledTasksInvoker(request *model.ListGroupScheduledTasksRequest) *ListGroupScheduledTasksInvoker {
	requestDef := GenReqDefForListGroupScheduledTasks()
	return &ListGroupScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHookInstances 查询伸缩实例挂起信息
//
// 添加生命周期挂钩后，当伸缩组进行伸缩活动时，实例将被挂钩挂起并置于等待状态，根据输入条件过滤查询弹性伸缩组中伸缩实例的挂起信息。可根据实例ID进行条件过滤查询。若不加过滤条件默认查询指定伸缩组内所有实例挂起信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListHookInstances(request *model.ListHookInstancesRequest) (*model.ListHookInstancesResponse, error) {
	requestDef := GenReqDefForListHookInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHookInstancesResponse), nil
	}
}

// ListHookInstancesInvoker 查询伸缩实例挂起信息
func (c *AsClient) ListHookInstancesInvoker(request *model.ListHookInstancesRequest) *ListHookInstancesInvoker {
	requestDef := GenReqDefForListHookInstances()
	return &ListHookInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLifeCycleHooks 查询生命周期挂钩列表
//
// 根据伸缩组ID查询生命周期挂钩列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListLifeCycleHooks(request *model.ListLifeCycleHooksRequest) (*model.ListLifeCycleHooksResponse, error) {
	requestDef := GenReqDefForListLifeCycleHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLifeCycleHooksResponse), nil
	}
}

// ListLifeCycleHooksInvoker 查询生命周期挂钩列表
func (c *AsClient) ListLifeCycleHooksInvoker(request *model.ListLifeCycleHooksRequest) *ListLifeCycleHooksInvoker {
	requestDef := GenReqDefForListLifeCycleHooks()
	return &ListLifeCycleHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 查询资源实例
//
// 根据项目ID查询指定资源类型的资源实例。资源、资源tag默认按照创建时间倒序。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 查询资源实例
func (c *AsClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingActivityLogs 查询伸缩活动日志
//
// 根据输入条件过滤查询伸缩活动日志。查询结果分页显示。可根据起始时间，截止时间，起始行号，记录数进行条件过滤查询。若不加过滤条件默认查询最多20条伸缩活动日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingActivityLogs(request *model.ListScalingActivityLogsRequest) (*model.ListScalingActivityLogsResponse, error) {
	requestDef := GenReqDefForListScalingActivityLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingActivityLogsResponse), nil
	}
}

// ListScalingActivityLogsInvoker 查询伸缩活动日志
func (c *AsClient) ListScalingActivityLogsInvoker(request *model.ListScalingActivityLogsRequest) *ListScalingActivityLogsInvoker {
	requestDef := GenReqDefForListScalingActivityLogs()
	return &ListScalingActivityLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingActivityV2Logs 查询伸缩活动日志v2版本
//
// 根据输入条件过滤查询伸缩活动日志，支持查询实例伸缩、ELB迁移、实例备用等类型活动。查询结果分页显示。查询伸缩活动日志V2版本与V1版本区别在于，V2版本展示了更详细的实例伸缩日志，如ELB迁移日志，实例备用日志信息。可根据起始时间，截止时间，起始行号，记录数，伸缩活动类型等作为条件过滤查询。若不加过滤条件默认查询最多20条伸缩活动日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingActivityV2Logs(request *model.ListScalingActivityV2LogsRequest) (*model.ListScalingActivityV2LogsResponse, error) {
	requestDef := GenReqDefForListScalingActivityV2Logs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingActivityV2LogsResponse), nil
	}
}

// ListScalingActivityV2LogsInvoker 查询伸缩活动日志v2版本
func (c *AsClient) ListScalingActivityV2LogsInvoker(request *model.ListScalingActivityV2LogsRequest) *ListScalingActivityV2LogsInvoker {
	requestDef := GenReqDefForListScalingActivityV2Logs()
	return &ListScalingActivityV2LogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingConfigs 查询弹性伸缩配置列表
//
// 根据输入条件过滤查询弹性伸缩配置。查询结果分页显示。可以根据伸缩配置名称，镜像ID，起始行号，记录条数进行条件过滤查询。若不加过滤条件默认最多查询租户下20条伸缩配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingConfigs(request *model.ListScalingConfigsRequest) (*model.ListScalingConfigsResponse, error) {
	requestDef := GenReqDefForListScalingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingConfigsResponse), nil
	}
}

// ListScalingConfigsInvoker 查询弹性伸缩配置列表
func (c *AsClient) ListScalingConfigsInvoker(request *model.ListScalingConfigsRequest) *ListScalingConfigsInvoker {
	requestDef := GenReqDefForListScalingConfigs()
	return &ListScalingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingGroups 查询弹性伸缩组列表
//
// 根据输入条件过滤查询弹性伸缩组列表。查询结果分页显示。可根据伸缩组名称，伸缩配置ID，伸缩组状态，企业项目ID，起始行号，记录条数进行条件过滤查询。若不加过滤条件默认最多查询租户下20条伸缩组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingGroups(request *model.ListScalingGroupsRequest) (*model.ListScalingGroupsResponse, error) {
	requestDef := GenReqDefForListScalingGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingGroupsResponse), nil
	}
}

// ListScalingGroupsInvoker 查询弹性伸缩组列表
func (c *AsClient) ListScalingGroupsInvoker(request *model.ListScalingGroupsRequest) *ListScalingGroupsInvoker {
	requestDef := GenReqDefForListScalingGroups()
	return &ListScalingGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingInstances 查询弹性伸缩组中的实例列表
//
// 根据输入条件过滤查询弹性伸缩组中实例信息。查询结果分页显示。可根据实例在伸缩组中的生命周期状态，实例健康状态，实例保护状态，起始行号，记录条数进行条件过滤查询。若不加过滤条件默认查询组内最多20条实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingInstances(request *model.ListScalingInstancesRequest) (*model.ListScalingInstancesResponse, error) {
	requestDef := GenReqDefForListScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingInstancesResponse), nil
	}
}

// ListScalingInstancesInvoker 查询弹性伸缩组中的实例列表
func (c *AsClient) ListScalingInstancesInvoker(request *model.ListScalingInstancesRequest) *ListScalingInstancesInvoker {
	requestDef := GenReqDefForListScalingInstances()
	return &ListScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingNotifications 查询伸缩组通知列表
//
// 根据伸缩组ID查询指定弹性伸缩组的通知列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingNotifications(request *model.ListScalingNotificationsRequest) (*model.ListScalingNotificationsResponse, error) {
	requestDef := GenReqDefForListScalingNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingNotificationsResponse), nil
	}
}

// ListScalingNotificationsInvoker 查询伸缩组通知列表
func (c *AsClient) ListScalingNotificationsInvoker(request *model.ListScalingNotificationsRequest) *ListScalingNotificationsInvoker {
	requestDef := GenReqDefForListScalingNotifications()
	return &ListScalingNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingPolicies 查询弹性伸缩策略列表
//
// 根据输入条件过滤查询弹性伸缩策略。查询结果分页显示。可根据伸缩策略名称，策略类型，伸缩策略ID，起始行号，记录数进行条件过滤查询。若不加过滤条件默认查询租户下指定伸缩组内最多20条伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingPolicies(request *model.ListScalingPoliciesRequest) (*model.ListScalingPoliciesResponse, error) {
	requestDef := GenReqDefForListScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingPoliciesResponse), nil
	}
}

// ListScalingPoliciesInvoker 查询弹性伸缩策略列表
func (c *AsClient) ListScalingPoliciesInvoker(request *model.ListScalingPoliciesRequest) *ListScalingPoliciesInvoker {
	requestDef := GenReqDefForListScalingPolicies()
	return &ListScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingPolicyExecuteLogs 查询策略执行日志
//
// 根据输入条件过滤查询策略执行的历史记录。查询结果分页显示。可根据日志ID，伸缩资源类型，伸缩资源ID，策略执行类型，查询额起始，查询截止时间，查询起始行号，查询记录数进行条件过滤查询。若不加过滤条件默认查询最多20条策略执行日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingPolicyExecuteLogs(request *model.ListScalingPolicyExecuteLogsRequest) (*model.ListScalingPolicyExecuteLogsResponse, error) {
	requestDef := GenReqDefForListScalingPolicyExecuteLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingPolicyExecuteLogsResponse), nil
	}
}

// ListScalingPolicyExecuteLogsInvoker 查询策略执行日志
func (c *AsClient) ListScalingPolicyExecuteLogsInvoker(request *model.ListScalingPolicyExecuteLogsRequest) *ListScalingPolicyExecuteLogsInvoker {
	requestDef := GenReqDefForListScalingPolicyExecuteLogs()
	return &ListScalingPolicyExecuteLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingTagInfosByResourceId 查询资源标签
//
// 根据项目ID和资源ID查询指定资源类型的资源标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingTagInfosByResourceId(request *model.ListScalingTagInfosByResourceIdRequest) (*model.ListScalingTagInfosByResourceIdResponse, error) {
	requestDef := GenReqDefForListScalingTagInfosByResourceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingTagInfosByResourceIdResponse), nil
	}
}

// ListScalingTagInfosByResourceIdInvoker 查询资源标签
func (c *AsClient) ListScalingTagInfosByResourceIdInvoker(request *model.ListScalingTagInfosByResourceIdRequest) *ListScalingTagInfosByResourceIdInvoker {
	requestDef := GenReqDefForListScalingTagInfosByResourceId()
	return &ListScalingTagInfosByResourceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingTagInfosByTenantId 查询标签
//
// 根据项目ID查询指定资源类型的标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingTagInfosByTenantId(request *model.ListScalingTagInfosByTenantIdRequest) (*model.ListScalingTagInfosByTenantIdResponse, error) {
	requestDef := GenReqDefForListScalingTagInfosByTenantId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingTagInfosByTenantIdResponse), nil
	}
}

// ListScalingTagInfosByTenantIdInvoker 查询标签
func (c *AsClient) ListScalingTagInfosByTenantIdInvoker(request *model.ListScalingTagInfosByTenantIdRequest) *ListScalingTagInfosByTenantIdInvoker {
	requestDef := GenReqDefForListScalingTagInfosByTenantId()
	return &ListScalingTagInfosByTenantIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseScalingGroup 停止弹性伸缩组
//
// 启用或停止一个指定弹性伸缩组。已停用状态的伸缩组，不会自动触发任何伸缩活动。当伸缩组正在进行伸缩活动，即使停用，正在进行的伸缩活动也不会立即停止。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) PauseScalingGroup(request *model.PauseScalingGroupRequest) (*model.PauseScalingGroupResponse, error) {
	requestDef := GenReqDefForPauseScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseScalingGroupResponse), nil
	}
}

// PauseScalingGroupInvoker 停止弹性伸缩组
func (c *AsClient) PauseScalingGroupInvoker(request *model.PauseScalingGroupRequest) *PauseScalingGroupInvoker {
	requestDef := GenReqDefForPauseScalingGroup()
	return &PauseScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseScalingPolicy 停止弹性伸缩策略。
//
// 立即执行或启用或停止一个指定弹性伸缩策略。当伸缩组、伸缩策略状态处于INSERVICE时，伸缩策略才能被正确执行，否则会执行失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) PauseScalingPolicy(request *model.PauseScalingPolicyRequest) (*model.PauseScalingPolicyResponse, error) {
	requestDef := GenReqDefForPauseScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseScalingPolicyResponse), nil
	}
}

// PauseScalingPolicyInvoker 停止弹性伸缩策略。
func (c *AsClient) PauseScalingPolicyInvoker(request *model.PauseScalingPolicyRequest) *PauseScalingPolicyInvoker {
	requestDef := GenReqDefForPauseScalingPolicy()
	return &PauseScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeScalingGroup 启用弹性伸缩组
//
// 启用或停止一个指定弹性伸缩组。已停用状态的伸缩组，不会自动触发任何伸缩活动。当伸缩组正在进行伸缩活动，即使停用，正在进行的伸缩活动也不会立即停止。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ResumeScalingGroup(request *model.ResumeScalingGroupRequest) (*model.ResumeScalingGroupResponse, error) {
	requestDef := GenReqDefForResumeScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeScalingGroupResponse), nil
	}
}

// ResumeScalingGroupInvoker 启用弹性伸缩组
func (c *AsClient) ResumeScalingGroupInvoker(request *model.ResumeScalingGroupRequest) *ResumeScalingGroupInvoker {
	requestDef := GenReqDefForResumeScalingGroup()
	return &ResumeScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeScalingPolicy 启用弹性伸缩策略。
//
// 立即执行或启用或停止一个指定弹性伸缩策略。当伸缩组、伸缩策略状态处于INSERVICE时，伸缩策略才能被正确执行，否则会执行失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ResumeScalingPolicy(request *model.ResumeScalingPolicyRequest) (*model.ResumeScalingPolicyResponse, error) {
	requestDef := GenReqDefForResumeScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeScalingPolicyResponse), nil
	}
}

// ResumeScalingPolicyInvoker 启用弹性伸缩策略。
func (c *AsClient) ResumeScalingPolicyInvoker(request *model.ResumeScalingPolicyRequest) *ResumeScalingPolicyInvoker {
	requestDef := GenReqDefForResumeScalingPolicy()
	return &ResumeScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLifeCycleHook 查询生命周期挂钩详情
//
// 根据伸缩组ID及生命周期挂钩名称查询指定的生命周期挂钩详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowLifeCycleHook(request *model.ShowLifeCycleHookRequest) (*model.ShowLifeCycleHookResponse, error) {
	requestDef := GenReqDefForShowLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLifeCycleHookResponse), nil
	}
}

// ShowLifeCycleHookInvoker 查询生命周期挂钩详情
func (c *AsClient) ShowLifeCycleHookInvoker(request *model.ShowLifeCycleHookRequest) *ShowLifeCycleHookInvoker {
	requestDef := GenReqDefForShowLifeCycleHook()
	return &ShowLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyAndInstanceQuota 查询弹性伸缩策略和伸缩实例配额
//
// 根据伸缩组ID查询指定弹性伸缩组下的伸缩策略和伸缩实例的配额总数及已使用配额数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowPolicyAndInstanceQuota(request *model.ShowPolicyAndInstanceQuotaRequest) (*model.ShowPolicyAndInstanceQuotaResponse, error) {
	requestDef := GenReqDefForShowPolicyAndInstanceQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyAndInstanceQuotaResponse), nil
	}
}

// ShowPolicyAndInstanceQuotaInvoker 查询弹性伸缩策略和伸缩实例配额
func (c *AsClient) ShowPolicyAndInstanceQuotaInvoker(request *model.ShowPolicyAndInstanceQuotaRequest) *ShowPolicyAndInstanceQuotaInvoker {
	requestDef := GenReqDefForShowPolicyAndInstanceQuota()
	return &ShowPolicyAndInstanceQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceQuota 查询配额
//
// 查询指定租户下的弹性伸缩组、伸缩配置、伸缩带宽策略、伸缩策略和伸缩实例的配额总数及已使用配额数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowResourceQuota(request *model.ShowResourceQuotaRequest) (*model.ShowResourceQuotaResponse, error) {
	requestDef := GenReqDefForShowResourceQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceQuotaResponse), nil
	}
}

// ShowResourceQuotaInvoker 查询配额
func (c *AsClient) ShowResourceQuotaInvoker(request *model.ShowResourceQuotaRequest) *ShowResourceQuotaInvoker {
	requestDef := GenReqDefForShowResourceQuota()
	return &ShowResourceQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScalingConfig 查询弹性伸缩配置详情
//
// 根据伸缩配置ID查询一个弹性伸缩配置的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowScalingConfig(request *model.ShowScalingConfigRequest) (*model.ShowScalingConfigResponse, error) {
	requestDef := GenReqDefForShowScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingConfigResponse), nil
	}
}

// ShowScalingConfigInvoker 查询弹性伸缩配置详情
func (c *AsClient) ShowScalingConfigInvoker(request *model.ShowScalingConfigRequest) *ShowScalingConfigInvoker {
	requestDef := GenReqDefForShowScalingConfig()
	return &ShowScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScalingGroup 查询弹性伸缩组详情
//
// 查询一个指定弹性伸缩组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowScalingGroup(request *model.ShowScalingGroupRequest) (*model.ShowScalingGroupResponse, error) {
	requestDef := GenReqDefForShowScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingGroupResponse), nil
	}
}

// ShowScalingGroupInvoker 查询弹性伸缩组详情
func (c *AsClient) ShowScalingGroupInvoker(request *model.ShowScalingGroupRequest) *ShowScalingGroupInvoker {
	requestDef := GenReqDefForShowScalingGroup()
	return &ShowScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScalingPolicy 查询弹性伸缩策略详情
//
// 查询指定弹性伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowScalingPolicy(request *model.ShowScalingPolicyRequest) (*model.ShowScalingPolicyResponse, error) {
	requestDef := GenReqDefForShowScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingPolicyResponse), nil
	}
}

// ShowScalingPolicyInvoker 查询弹性伸缩策略详情
func (c *AsClient) ShowScalingPolicyInvoker(request *model.ShowScalingPolicyRequest) *ShowScalingPolicyInvoker {
	requestDef := GenReqDefForShowScalingPolicy()
	return &ShowScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupScheduledTask 更新计划任务
//
// 更新计划任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) UpdateGroupScheduledTask(request *model.UpdateGroupScheduledTaskRequest) (*model.UpdateGroupScheduledTaskResponse, error) {
	requestDef := GenReqDefForUpdateGroupScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupScheduledTaskResponse), nil
	}
}

// UpdateGroupScheduledTaskInvoker 更新计划任务
func (c *AsClient) UpdateGroupScheduledTaskInvoker(request *model.UpdateGroupScheduledTaskRequest) *UpdateGroupScheduledTaskInvoker {
	requestDef := GenReqDefForUpdateGroupScheduledTask()
	return &UpdateGroupScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLifeCycleHook 修改生命周期挂钩
//
// 修改一个指定生命周期挂钩中的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) UpdateLifeCycleHook(request *model.UpdateLifeCycleHookRequest) (*model.UpdateLifeCycleHookResponse, error) {
	requestDef := GenReqDefForUpdateLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLifeCycleHookResponse), nil
	}
}

// UpdateLifeCycleHookInvoker 修改生命周期挂钩
func (c *AsClient) UpdateLifeCycleHookInvoker(request *model.UpdateLifeCycleHookRequest) *UpdateLifeCycleHookInvoker {
	requestDef := GenReqDefForUpdateLifeCycleHook()
	return &UpdateLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScalingGroup 修改弹性伸缩组
//
// 修改一个指定弹性伸缩组中的信息。更换伸缩组的伸缩配置，伸缩组中已经存在的使用之前伸缩配置创建的云服务器云主机不受影响。伸缩组为没有正在进行的伸缩活动时，可以修改伸缩组的子网、可用区和负载均衡配置。当伸缩组的期望实例数改变时，会触发伸缩活动加入或移出实例。期望实例数必须大于或等于最小实例数，必须小于或等于最大实例数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) UpdateScalingGroup(request *model.UpdateScalingGroupRequest) (*model.UpdateScalingGroupResponse, error) {
	requestDef := GenReqDefForUpdateScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingGroupResponse), nil
	}
}

// UpdateScalingGroupInvoker 修改弹性伸缩组
func (c *AsClient) UpdateScalingGroupInvoker(request *model.UpdateScalingGroupRequest) *UpdateScalingGroupInvoker {
	requestDef := GenReqDefForUpdateScalingGroup()
	return &UpdateScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScalingPolicy 修改弹性伸缩策略
//
// 修改指定弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) UpdateScalingPolicy(request *model.UpdateScalingPolicyRequest) (*model.UpdateScalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingPolicyResponse), nil
	}
}

// UpdateScalingPolicyInvoker 修改弹性伸缩策略
func (c *AsClient) UpdateScalingPolicyInvoker(request *model.UpdateScalingPolicyRequest) *UpdateScalingPolicyInvoker {
	requestDef := GenReqDefForUpdateScalingPolicy()
	return &UpdateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询弹性伸缩API所有版本信息
//
// 查询弹性伸缩API所有版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询弹性伸缩API所有版本信息
func (c *AsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询弹性伸缩API指定版本信息
//
// 根据租户id和资源id查询指定资源类型的标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询弹性伸缩API指定版本信息
func (c *AsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingV2Policy 创建弹性伸缩策略（V2版本）
//
// 可针对不同类型资源如伸缩组或带宽，创建弹性伸缩策略。创建弹性伸缩策略V2版本与V1版本的区别在于，V2版本支持创建对带宽资源进行调整的策略，通过伸缩资源类型区分伸缩资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) CreateScalingV2Policy(request *model.CreateScalingV2PolicyRequest) (*model.CreateScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingV2PolicyResponse), nil
	}
}

// CreateScalingV2PolicyInvoker 创建弹性伸缩策略（V2版本）
func (c *AsClient) CreateScalingV2PolicyInvoker(request *model.CreateScalingV2PolicyRequest) *CreateScalingV2PolicyInvoker {
	requestDef := GenReqDefForCreateScalingV2Policy()
	return &CreateScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllScalingV2Policies 查询弹性伸缩策略全量列表（V2版本）
//
// 根据输入条件过滤查询弹性伸缩策略，支持查询当前租户下全量伸缩策略。查询结果分页显示。可根据伸缩资源ID，伸缩资源类型，伸缩策略名称，伸缩策略ID，告警ID，企业项目ID，起始行号，记录数，排序方式等条件进行过滤查询。若不加过滤添加默认查询该租户下最多20条伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListAllScalingV2Policies(request *model.ListAllScalingV2PoliciesRequest) (*model.ListAllScalingV2PoliciesResponse, error) {
	requestDef := GenReqDefForListAllScalingV2Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllScalingV2PoliciesResponse), nil
	}
}

// ListAllScalingV2PoliciesInvoker 查询弹性伸缩策略全量列表（V2版本）
func (c *AsClient) ListAllScalingV2PoliciesInvoker(request *model.ListAllScalingV2PoliciesRequest) *ListAllScalingV2PoliciesInvoker {
	requestDef := GenReqDefForListAllScalingV2Policies()
	return &ListAllScalingV2PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScalingV2Policies 查询弹性伸缩策略列表（V2版本）
//
// 根据输入条件过滤查询弹性伸缩策略。查询结果分页显示。查询弹性伸缩策略V2版本与V1版本的区别在于，V2版本响应含伸缩资源类型。可根据伸缩策略名称，策略类型，伸缩策略ID，起始行号，记录数进行条件过滤查询。若不加过滤条件默认查询该租户下指定资源下最多20条伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ListScalingV2Policies(request *model.ListScalingV2PoliciesRequest) (*model.ListScalingV2PoliciesResponse, error) {
	requestDef := GenReqDefForListScalingV2Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingV2PoliciesResponse), nil
	}
}

// ListScalingV2PoliciesInvoker 查询弹性伸缩策略列表（V2版本）
func (c *AsClient) ListScalingV2PoliciesInvoker(request *model.ListScalingV2PoliciesRequest) *ListScalingV2PoliciesInvoker {
	requestDef := GenReqDefForListScalingV2Policies()
	return &ListScalingV2PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScalingV2Policy 查询指定弹性伸缩策略详情（V2版本）
//
// 查询指定弹性伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) ShowScalingV2Policy(request *model.ShowScalingV2PolicyRequest) (*model.ShowScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForShowScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingV2PolicyResponse), nil
	}
}

// ShowScalingV2PolicyInvoker 查询指定弹性伸缩策略详情（V2版本）
func (c *AsClient) ShowScalingV2PolicyInvoker(request *model.ShowScalingV2PolicyRequest) *ShowScalingV2PolicyInvoker {
	requestDef := GenReqDefForShowScalingV2Policy()
	return &ShowScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScalingV2Policy 修改弹性伸缩策略（V2版本）
//
// 修改指定弹性伸缩策略。修改弹性伸缩策略V2版本与V1版本的区别在于，V2版本支持修改伸缩资源类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsClient) UpdateScalingV2Policy(request *model.UpdateScalingV2PolicyRequest) (*model.UpdateScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForUpdateScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingV2PolicyResponse), nil
	}
}

// UpdateScalingV2PolicyInvoker 修改弹性伸缩策略（V2版本）
func (c *AsClient) UpdateScalingV2PolicyInvoker(request *model.UpdateScalingV2PolicyRequest) *UpdateScalingV2PolicyInvoker {
	requestDef := GenReqDefForUpdateScalingV2Policy()
	return &UpdateScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
