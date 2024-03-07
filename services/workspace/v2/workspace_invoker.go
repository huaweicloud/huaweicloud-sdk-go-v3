package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

type BatchDeleteAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAccessPoliciesInvoker) Invoke() (*model.BatchDeleteAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAccessPoliciesResponse), nil
	}
}

type CreateAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessPolicyInvoker) Invoke() (*model.CreateAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessPolicyResponse), nil
	}
}

type ListAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPoliciesInvoker) Invoke() (*model.ListAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPoliciesResponse), nil
	}
}

type ListAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPolicyObjectsInvoker) Invoke() (*model.ListAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPolicyObjectsResponse), nil
	}
}

type UpdateAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessPolicyObjectsInvoker) Invoke() (*model.UpdateAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessPolicyObjectsResponse), nil
	}
}

type CreateAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgenciesInvoker) Invoke() (*model.CreateAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgenciesResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type ShowAssistAuthConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssistAuthConfigInvoker) Invoke() (*model.ShowAssistAuthConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssistAuthConfigResponse), nil
	}
}

type UpdateAssistAuthMethodConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssistAuthMethodConfigInvoker) Invoke() (*model.UpdateAssistAuthMethodConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssistAuthMethodConfigResponse), nil
	}
}

type ListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZonesInvoker) Invoke() (*model.ListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZonesResponse), nil
	}
}

type ExportUserLoginInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserLoginInfoNewInvoker) Invoke() (*model.ExportUserLoginInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserLoginInfoNewResponse), nil
	}
}

type ListHistoryOnlineInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryOnlineInfoNewInvoker) Invoke() (*model.ListHistoryOnlineInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryOnlineInfoNewResponse), nil
	}
}

type ListLoginRecordsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginRecordsNewInvoker) Invoke() (*model.ListLoginRecordsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginRecordsNewResponse), nil
	}
}

type AttachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachInstancesInvoker) Invoke() (*model.AttachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachInstancesResponse), nil
	}
}

type BatchDeleteDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopsInvoker) Invoke() (*model.BatchDeleteDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopsResponse), nil
	}
}

type BatchLogoffDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchLogoffDesktopsInvoker) Invoke() (*model.BatchLogoffDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchLogoffDesktopsResponse), nil
	}
}

type BatchRebuildDesktopsSystemDiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebuildDesktopsSystemDiskInvoker) Invoke() (*model.BatchRebuildDesktopsSystemDiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebuildDesktopsSystemDiskResponse), nil
	}
}

type BatchRunDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRunDesktopsInvoker) Invoke() (*model.BatchRunDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRunDesktopsResponse), nil
	}
}

type ChangeDesktopNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDesktopNetworkInvoker) Invoke() (*model.ChangeDesktopNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDesktopNetworkResponse), nil
	}
}

type CreateDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopInvoker) Invoke() (*model.CreateDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopResponse), nil
	}
}

type DeleteDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopInvoker) Invoke() (*model.DeleteDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopResponse), nil
	}
}

type DetachInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachInstancesInvoker) Invoke() (*model.DetachInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachInstancesResponse), nil
	}
}

type ListDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsInvoker) Invoke() (*model.ListDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsResponse), nil
	}
}

type ListDesktopsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsDetailInvoker) Invoke() (*model.ListDesktopsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsDetailResponse), nil
	}
}

type ResizeDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeDesktopInvoker) Invoke() (*model.ResizeDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeDesktopResponse), nil
	}
}

type ShowDesktopDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopDetailInvoker) Invoke() (*model.ShowDesktopDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopDetailResponse), nil
	}
}

type ShowDesktopNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopNetworkInvoker) Invoke() (*model.ShowDesktopNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopNetworkResponse), nil
	}
}

type BatchDeleteDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopNamePolicyInvoker) Invoke() (*model.BatchDeleteDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopNamePolicyResponse), nil
	}
}

type CreateDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopNamePolicyInvoker) Invoke() (*model.CreateDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopNamePolicyResponse), nil
	}
}

type ListDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopNamePolicyInvoker) Invoke() (*model.ListDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopNamePolicyResponse), nil
	}
}

type UpdateDesktopNamePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesktopNamePolicyInvoker) Invoke() (*model.UpdateDesktopNamePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesktopNamePolicyResponse), nil
	}
}

type ListUnusedDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnusedDesktopsInvoker) Invoke() (*model.ListUnusedDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnusedDesktopsResponse), nil
	}
}

type ListUsedDesktopInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsedDesktopInfoInvoker) Invoke() (*model.ListUsedDesktopInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsedDesktopInfoResponse), nil
	}
}

type BatchAddDesktopsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDesktopsTagsInvoker) Invoke() (*model.BatchAddDesktopsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDesktopsTagsResponse), nil
	}
}

type BatchChangeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeTagsInvoker) Invoke() (*model.BatchChangeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeTagsResponse), nil
	}
}

type BatchDeleteDesktopsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopsTagsInvoker) Invoke() (*model.BatchDeleteDesktopsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopsTagsResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListDesktopByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopByTagsInvoker) Invoke() (*model.ListDesktopByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopByTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ShowTagByDesktopIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagByDesktopIdInvoker) Invoke() (*model.ShowTagByDesktopIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagByDesktopIdResponse), nil
	}
}

type BatchDeleteUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteUserGroupsInvoker) Invoke() (*model.BatchDeleteUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteUserGroupsResponse), nil
	}
}

type CreateUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserGroupInvoker) Invoke() (*model.CreateUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserGroupResponse), nil
	}
}

type DeleteUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserGroupInvoker) Invoke() (*model.DeleteUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserGroupResponse), nil
	}
}

type ListUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserGroupsInvoker) Invoke() (*model.ListUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserGroupsResponse), nil
	}
}

type ListUsersOfGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersOfGroupInvoker) Invoke() (*model.ListUsersOfGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersOfGroupResponse), nil
	}
}

type RunActionsOnGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunActionsOnGroupInvoker) Invoke() (*model.RunActionsOnGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunActionsOnGroupResponse), nil
	}
}

type UpdateUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserGroupInvoker) Invoke() (*model.UpdateUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserGroupResponse), nil
	}
}

type ListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesInvoker) Invoke() (*model.ListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesResponse), nil
	}
}

type ListItaSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListItaSubJobsInvoker) Invoke() (*model.ListItaSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListItaSubJobsResponse), nil
	}
}

type ApplyDesktopsInternetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyDesktopsInternetInvoker) Invoke() (*model.ApplyDesktopsInternetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyDesktopsInternetResponse), nil
	}
}

type AssociateDesktopsEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateDesktopsEipInvoker) Invoke() (*model.AssociateDesktopsEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateDesktopsEipResponse), nil
	}
}

type BatchDisassociateDesktopsEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociateDesktopsEipInvoker) Invoke() (*model.BatchDisassociateDesktopsEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociateDesktopsEipResponse), nil
	}
}

type ListDesktopsEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsEipsInvoker) Invoke() (*model.ListDesktopsEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsEipsResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type BatchDeleteScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScheduledTasksInvoker) Invoke() (*model.BatchDeleteScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScheduledTasksResponse), nil
	}
}

type CreateScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduledTasksInvoker) Invoke() (*model.CreateScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduledTasksResponse), nil
	}
}

type DeleteScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduledTasksInvoker) Invoke() (*model.DeleteScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduledTasksResponse), nil
	}
}

type ListFutureExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFutureExecutionsInvoker) Invoke() (*model.ListFutureExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFutureExecutionsResponse), nil
	}
}

type ListScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksInvoker) Invoke() (*model.ListScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksResponse), nil
	}
}

type ListScheduledTasksRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksRecordsInvoker) Invoke() (*model.ListScheduledTasksRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksRecordsResponse), nil
	}
}

type ListScheduledTasksRecordsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksRecordsDetailsInvoker) Invoke() (*model.ListScheduledTasksRecordsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksRecordsDetailsResponse), nil
	}
}

type ShowScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduledTasksInvoker) Invoke() (*model.ShowScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduledTasksResponse), nil
	}
}

type UpdateScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduledTasksInvoker) Invoke() (*model.UpdateScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduledTasksResponse), nil
	}
}

type AddMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddMetricNotifyRuleInvoker) Invoke() (*model.AddMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddMetricNotifyRuleResponse), nil
	}
}

type DeleteMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMetricNotifyRuleInvoker) Invoke() (*model.DeleteMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMetricNotifyRuleResponse), nil
	}
}

type ListDesktopUsageMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopUsageMetricInvoker) Invoke() (*model.ListDesktopUsageMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopUsageMetricResponse), nil
	}
}

type ListMetricNotifyRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricNotifyRecordInvoker) Invoke() (*model.ListMetricNotifyRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricNotifyRecordResponse), nil
	}
}

type ListMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricNotifyRuleInvoker) Invoke() (*model.ListMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricNotifyRuleResponse), nil
	}
}

type ListUserUsageMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserUsageMetricInvoker) Invoke() (*model.ListUserUsageMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserUsageMetricResponse), nil
	}
}

type UpdateMetricNotifyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMetricNotifyRuleInvoker) Invoke() (*model.UpdateMetricNotifyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMetricNotifyRuleResponse), nil
	}
}

type CreateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTerminalsBindingDesktopsInvoker) Invoke() (*model.CreateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTerminalsBindingDesktopsResponse), nil
	}
}

type DeleteTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTerminalsBindingDesktopsInvoker) Invoke() (*model.DeleteTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTerminalsBindingDesktopsResponse), nil
	}
}

type ListTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsInvoker) Invoke() (*model.ListTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsResponse), nil
	}
}

type ListTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.ListTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsConfigResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsConfigResponse), nil
	}
}

type BatchCreateUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateUsersInvoker) Invoke() (*model.BatchCreateUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateUsersResponse), nil
	}
}

type BatchDeleteOtpDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteOtpDevicesInvoker) Invoke() (*model.BatchDeleteOtpDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteOtpDevicesResponse), nil
	}
}

type ChangeUserStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeUserStatusInvoker) Invoke() (*model.ChangeUserStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeUserStatusResponse), nil
	}
}

type CreateDesktopUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopUserInvoker) Invoke() (*model.CreateDesktopUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ListOtpDevicesByUserIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOtpDevicesByUserIdInvoker) Invoke() (*model.ListOtpDevicesByUserIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOtpDevicesByUserIdResponse), nil
	}
}

type ListUserDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserDetailInvoker) Invoke() (*model.ListUserDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserDetailResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ResetRandomPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetRandomPasswordInvoker) Invoke() (*model.ResetRandomPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetRandomPasswordResponse), nil
	}
}

type UpdateUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInfoInvoker) Invoke() (*model.UpdateUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserInfoResponse), nil
	}
}

type AddVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVolumesInvoker) Invoke() (*model.AddVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVolumesResponse), nil
	}
}

type DeleteDesktopVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopVolumesInvoker) Invoke() (*model.DeleteDesktopVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopVolumesResponse), nil
	}
}

type ExpandVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandVolumesInvoker) Invoke() (*model.ExpandVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandVolumesResponse), nil
	}
}

type ApplyWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyWorkspaceInvoker) Invoke() (*model.ApplyWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyWorkspaceResponse), nil
	}
}

type CancelWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelWorkspaceInvoker) Invoke() (*model.CancelWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelWorkspaceResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type ShowWorkspaceLockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceLockInvoker) Invoke() (*model.ShowWorkspaceLockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceLockResponse), nil
	}
}

type UnlockWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockWorkspaceInvoker) Invoke() (*model.UnlockWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockWorkspaceResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}
