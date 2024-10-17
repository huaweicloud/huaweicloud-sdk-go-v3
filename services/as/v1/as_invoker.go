package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1/model"
)

type AttachCallbackInstanceLifeCycleHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachCallbackInstanceLifeCycleHookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachCallbackInstanceLifeCycleHookInvoker) Invoke() (*model.AttachCallbackInstanceLifeCycleHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachCallbackInstanceLifeCycleHookResponse), nil
	}
}

type BatchAddScalingInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddScalingInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddScalingInstancesInvoker) Invoke() (*model.BatchAddScalingInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddScalingInstancesResponse), nil
	}
}

type BatchDeleteScalingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScalingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteScalingConfigsInvoker) Invoke() (*model.BatchDeleteScalingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScalingConfigsResponse), nil
	}
}

type BatchDeleteScalingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScalingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteScalingPoliciesInvoker) Invoke() (*model.BatchDeleteScalingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScalingPoliciesResponse), nil
	}
}

type BatchPauseScalingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchPauseScalingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchPauseScalingPoliciesInvoker) Invoke() (*model.BatchPauseScalingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchPauseScalingPoliciesResponse), nil
	}
}

type BatchProtectScalingInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchProtectScalingInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchProtectScalingInstancesInvoker) Invoke() (*model.BatchProtectScalingInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchProtectScalingInstancesResponse), nil
	}
}

type BatchRemoveScalingInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemoveScalingInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRemoveScalingInstancesInvoker) Invoke() (*model.BatchRemoveScalingInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemoveScalingInstancesResponse), nil
	}
}

type BatchResumeScalingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResumeScalingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchResumeScalingPoliciesInvoker) Invoke() (*model.BatchResumeScalingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResumeScalingPoliciesResponse), nil
	}
}

type BatchSetScalingInstancesStandbyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetScalingInstancesStandbyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetScalingInstancesStandbyInvoker) Invoke() (*model.BatchSetScalingInstancesStandbyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetScalingInstancesStandbyResponse), nil
	}
}

type BatchUnprotectScalingInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUnprotectScalingInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUnprotectScalingInstancesInvoker) Invoke() (*model.BatchUnprotectScalingInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUnprotectScalingInstancesResponse), nil
	}
}

type BatchUnsetScalingInstancesStantbyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUnsetScalingInstancesStantbyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUnsetScalingInstancesStantbyInvoker) Invoke() (*model.BatchUnsetScalingInstancesStantbyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUnsetScalingInstancesStantbyResponse), nil
	}
}

type CreateGroupScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupScheduledTaskInvoker) Invoke() (*model.CreateGroupScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupScheduledTaskResponse), nil
	}
}

type CreateLifyCycleHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLifyCycleHookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLifyCycleHookInvoker) Invoke() (*model.CreateLifyCycleHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLifyCycleHookResponse), nil
	}
}

type CreateScalingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingConfigInvoker) Invoke() (*model.CreateScalingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingConfigResponse), nil
	}
}

type CreateScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingGroupInvoker) Invoke() (*model.CreateScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingGroupResponse), nil
	}
}

type CreateScalingNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingNotificationInvoker) Invoke() (*model.CreateScalingNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingNotificationResponse), nil
	}
}

type CreateScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingPolicyInvoker) Invoke() (*model.CreateScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingPolicyResponse), nil
	}
}

type CreateScalingTagInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingTagInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingTagInfoInvoker) Invoke() (*model.CreateScalingTagInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingTagInfoResponse), nil
	}
}

type DeleteGroupScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupScheduledTaskInvoker) Invoke() (*model.DeleteGroupScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupScheduledTaskResponse), nil
	}
}

type DeleteLifecycleHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLifecycleHookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLifecycleHookInvoker) Invoke() (*model.DeleteLifecycleHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLifecycleHookResponse), nil
	}
}

type DeleteScalingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingConfigInvoker) Invoke() (*model.DeleteScalingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingConfigResponse), nil
	}
}

type DeleteScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingGroupInvoker) Invoke() (*model.DeleteScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingGroupResponse), nil
	}
}

type DeleteScalingInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingInstanceInvoker) Invoke() (*model.DeleteScalingInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingInstanceResponse), nil
	}
}

type DeleteScalingNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingNotificationInvoker) Invoke() (*model.DeleteScalingNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingNotificationResponse), nil
	}
}

type DeleteScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingPolicyInvoker) Invoke() (*model.DeleteScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingPolicyResponse), nil
	}
}

type DeleteScalingTagInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingTagInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingTagInfoInvoker) Invoke() (*model.DeleteScalingTagInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingTagInfoResponse), nil
	}
}

type ExecuteScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteScalingPolicyInvoker) Invoke() (*model.ExecuteScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteScalingPolicyResponse), nil
	}
}

type ListGroupScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupScheduledTasksInvoker) Invoke() (*model.ListGroupScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupScheduledTasksResponse), nil
	}
}

type ListHookInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHookInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHookInstancesInvoker) Invoke() (*model.ListHookInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHookInstancesResponse), nil
	}
}

type ListLifeCycleHooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLifeCycleHooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLifeCycleHooksInvoker) Invoke() (*model.ListLifeCycleHooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLifeCycleHooksResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstancesInvoker) Invoke() (*model.ListResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstancesResponse), nil
	}
}

type ListScalingActivityLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingActivityLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingActivityLogsInvoker) Invoke() (*model.ListScalingActivityLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingActivityLogsResponse), nil
	}
}

type ListScalingActivityV2LogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingActivityV2LogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingActivityV2LogsInvoker) Invoke() (*model.ListScalingActivityV2LogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingActivityV2LogsResponse), nil
	}
}

type ListScalingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingConfigsInvoker) Invoke() (*model.ListScalingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingConfigsResponse), nil
	}
}

type ListScalingGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingGroupsInvoker) Invoke() (*model.ListScalingGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingGroupsResponse), nil
	}
}

type ListScalingInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingInstancesInvoker) Invoke() (*model.ListScalingInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingInstancesResponse), nil
	}
}

type ListScalingNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingNotificationsInvoker) Invoke() (*model.ListScalingNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingNotificationsResponse), nil
	}
}

type ListScalingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingPoliciesInvoker) Invoke() (*model.ListScalingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingPoliciesResponse), nil
	}
}

type ListScalingPolicyExecuteLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingPolicyExecuteLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingPolicyExecuteLogsInvoker) Invoke() (*model.ListScalingPolicyExecuteLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingPolicyExecuteLogsResponse), nil
	}
}

type ListScalingTagInfosByResourceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingTagInfosByResourceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingTagInfosByResourceIdInvoker) Invoke() (*model.ListScalingTagInfosByResourceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingTagInfosByResourceIdResponse), nil
	}
}

type ListScalingTagInfosByTenantIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingTagInfosByTenantIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingTagInfosByTenantIdInvoker) Invoke() (*model.ListScalingTagInfosByTenantIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingTagInfosByTenantIdResponse), nil
	}
}

type PauseScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PauseScalingGroupInvoker) Invoke() (*model.PauseScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseScalingGroupResponse), nil
	}
}

type PauseScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PauseScalingPolicyInvoker) Invoke() (*model.PauseScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseScalingPolicyResponse), nil
	}
}

type ResumeScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeScalingGroupInvoker) Invoke() (*model.ResumeScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeScalingGroupResponse), nil
	}
}

type ResumeScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeScalingPolicyInvoker) Invoke() (*model.ResumeScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeScalingPolicyResponse), nil
	}
}

type ShowLifeCycleHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLifeCycleHookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLifeCycleHookInvoker) Invoke() (*model.ShowLifeCycleHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLifeCycleHookResponse), nil
	}
}

type ShowPolicyAndInstanceQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyAndInstanceQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyAndInstanceQuotaInvoker) Invoke() (*model.ShowPolicyAndInstanceQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyAndInstanceQuotaResponse), nil
	}
}

type ShowResourceQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceQuotaInvoker) Invoke() (*model.ShowResourceQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceQuotaResponse), nil
	}
}

type ShowScalingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScalingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScalingConfigInvoker) Invoke() (*model.ShowScalingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScalingConfigResponse), nil
	}
}

type ShowScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScalingGroupInvoker) Invoke() (*model.ShowScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScalingGroupResponse), nil
	}
}

type ShowScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScalingPolicyInvoker) Invoke() (*model.ShowScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScalingPolicyResponse), nil
	}
}

type UpdateGroupScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupScheduledTaskInvoker) Invoke() (*model.UpdateGroupScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupScheduledTaskResponse), nil
	}
}

type UpdateLifeCycleHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLifeCycleHookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLifeCycleHookInvoker) Invoke() (*model.UpdateLifeCycleHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLifeCycleHookResponse), nil
	}
}

type UpdateScalingGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScalingGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScalingGroupInvoker) Invoke() (*model.UpdateScalingGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScalingGroupResponse), nil
	}
}

type UpdateScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScalingPolicyInvoker) Invoke() (*model.UpdateScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScalingPolicyResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}

type CreateScalingV2PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingV2PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScalingV2PolicyInvoker) Invoke() (*model.CreateScalingV2PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingV2PolicyResponse), nil
	}
}

type ListAllScalingV2PoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllScalingV2PoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllScalingV2PoliciesInvoker) Invoke() (*model.ListAllScalingV2PoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllScalingV2PoliciesResponse), nil
	}
}

type ListScalingV2PoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingV2PoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingV2PoliciesInvoker) Invoke() (*model.ListScalingV2PoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingV2PoliciesResponse), nil
	}
}

type ShowScalingV2PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScalingV2PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScalingV2PolicyInvoker) Invoke() (*model.ShowScalingV2PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScalingV2PolicyResponse), nil
	}
}

type UpdateScalingV2PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScalingV2PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScalingV2PolicyInvoker) Invoke() (*model.UpdateScalingV2PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScalingV2PolicyResponse), nil
	}
}
