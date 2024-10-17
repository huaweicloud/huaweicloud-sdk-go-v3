package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sdrs/v1/model"
)

type AddProtectedInstanceNicInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddProtectedInstanceNicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddProtectedInstanceNicInvoker) Invoke() (*model.AddProtectedInstanceNicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProtectedInstanceNicResponse), nil
	}
}

type AddProtectedInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddProtectedInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddProtectedInstanceTagsInvoker) Invoke() (*model.AddProtectedInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProtectedInstanceTagsResponse), nil
	}
}

type AttachProtectedInstanceReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachProtectedInstanceReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachProtectedInstanceReplicationInvoker) Invoke() (*model.AttachProtectedInstanceReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachProtectedInstanceReplicationResponse), nil
	}
}

type BatchAddTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddTagsInvoker) Invoke() (*model.BatchAddTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddTagsResponse), nil
	}
}

type BatchCreateProtectedInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateProtectedInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateProtectedInstancesInvoker) Invoke() (*model.BatchCreateProtectedInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateProtectedInstancesResponse), nil
	}
}

type BatchDeleteProtectedInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteProtectedInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteProtectedInstancesInvoker) Invoke() (*model.BatchDeleteProtectedInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteProtectedInstancesResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type CreateDisasterRecoveryDrillInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDisasterRecoveryDrillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDisasterRecoveryDrillInvoker) Invoke() (*model.CreateDisasterRecoveryDrillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDisasterRecoveryDrillResponse), nil
	}
}

type CreateProtectedInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProtectedInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProtectedInstanceInvoker) Invoke() (*model.CreateProtectedInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProtectedInstanceResponse), nil
	}
}

type CreateProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProtectionGroupInvoker) Invoke() (*model.CreateProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProtectionGroupResponse), nil
	}
}

type CreateReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReplicationInvoker) Invoke() (*model.CreateReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReplicationResponse), nil
	}
}

type DeleteAllServerGroupFailureJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAllServerGroupFailureJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAllServerGroupFailureJobsInvoker) Invoke() (*model.DeleteAllServerGroupFailureJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAllServerGroupFailureJobsResponse), nil
	}
}

type DeleteDisasterRecoveryDrillInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDisasterRecoveryDrillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDisasterRecoveryDrillInvoker) Invoke() (*model.DeleteDisasterRecoveryDrillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDisasterRecoveryDrillResponse), nil
	}
}

type DeleteFailureJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFailureJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFailureJobInvoker) Invoke() (*model.DeleteFailureJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFailureJobResponse), nil
	}
}

type DeleteProtectedInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectedInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectedInstanceInvoker) Invoke() (*model.DeleteProtectedInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectedInstanceResponse), nil
	}
}

type DeleteProtectedInstanceNicInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectedInstanceNicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectedInstanceNicInvoker) Invoke() (*model.DeleteProtectedInstanceNicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectedInstanceNicResponse), nil
	}
}

type DeleteProtectedInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectedInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectedInstanceTagInvoker) Invoke() (*model.DeleteProtectedInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectedInstanceTagResponse), nil
	}
}

type DeleteProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectionGroupInvoker) Invoke() (*model.DeleteProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectionGroupResponse), nil
	}
}

type DeleteReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteReplicationInvoker) Invoke() (*model.DeleteReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteReplicationResponse), nil
	}
}

type DeleteServerGroupFailureJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerGroupFailureJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServerGroupFailureJobsInvoker) Invoke() (*model.DeleteServerGroupFailureJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerGroupFailureJobsResponse), nil
	}
}

type DetachProtectedInstanceReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachProtectedInstanceReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachProtectedInstanceReplicationInvoker) Invoke() (*model.DetachProtectedInstanceReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachProtectedInstanceReplicationResponse), nil
	}
}

type ExpandReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandReplicationInvoker) Invoke() (*model.ExpandReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandReplicationResponse), nil
	}
}

type ListActiveActiveDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListActiveActiveDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListActiveActiveDomainsInvoker) Invoke() (*model.ListActiveActiveDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListActiveActiveDomainsResponse), nil
	}
}

type ListDisasterRecoveryDrillsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDisasterRecoveryDrillsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDisasterRecoveryDrillsInvoker) Invoke() (*model.ListDisasterRecoveryDrillsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDisasterRecoveryDrillsResponse), nil
	}
}

type ListFailureJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFailureJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFailureJobsInvoker) Invoke() (*model.ListFailureJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFailureJobsResponse), nil
	}
}

type ListProtectedInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedInstanceTagsInvoker) Invoke() (*model.ListProtectedInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedInstanceTagsResponse), nil
	}
}

type ListProtectedInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedInstancesInvoker) Invoke() (*model.ListProtectedInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedInstancesResponse), nil
	}
}

type ListProtectedInstancesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedInstancesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedInstancesByTagsInvoker) Invoke() (*model.ListProtectedInstancesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedInstancesByTagsResponse), nil
	}
}

type ListProtectedInstancesProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedInstancesProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedInstancesProjectTagsInvoker) Invoke() (*model.ListProtectedInstancesProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedInstancesProjectTagsResponse), nil
	}
}

type ListProtectionGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectionGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectionGroupsInvoker) Invoke() (*model.ListProtectionGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectionGroupsResponse), nil
	}
}

type ListReplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReplicationsInvoker) Invoke() (*model.ListReplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReplicationsResponse), nil
	}
}

type ListRpoStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRpoStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRpoStatisticsInvoker) Invoke() (*model.ListRpoStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRpoStatisticsResponse), nil
	}
}

type ResizeProtectedInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeProtectedInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeProtectedInstanceInvoker) Invoke() (*model.ResizeProtectedInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeProtectedInstanceResponse), nil
	}
}

type ShowDisasterRecoveryDrillInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDisasterRecoveryDrillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDisasterRecoveryDrillInvoker) Invoke() (*model.ShowDisasterRecoveryDrillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDisasterRecoveryDrillResponse), nil
	}
}

type ShowProtectedInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectedInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProtectedInstanceInvoker) Invoke() (*model.ShowProtectedInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectedInstanceResponse), nil
	}
}

type ShowProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProtectionGroupInvoker) Invoke() (*model.ShowProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectionGroupResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type ShowReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReplicationInvoker) Invoke() (*model.ShowReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReplicationResponse), nil
	}
}

type StartFailoverProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartFailoverProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartFailoverProtectionGroupInvoker) Invoke() (*model.StartFailoverProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartFailoverProtectionGroupResponse), nil
	}
}

type StartProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartProtectionGroupInvoker) Invoke() (*model.StartProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartProtectionGroupResponse), nil
	}
}

type StartReverseProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartReverseProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartReverseProtectionGroupInvoker) Invoke() (*model.StartReverseProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartReverseProtectionGroupResponse), nil
	}
}

type StopProtectionGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopProtectionGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopProtectionGroupInvoker) Invoke() (*model.StopProtectionGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopProtectionGroupResponse), nil
	}
}

type UpdateDisasterRecoveryDrillNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDisasterRecoveryDrillNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDisasterRecoveryDrillNameInvoker) Invoke() (*model.UpdateDisasterRecoveryDrillNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDisasterRecoveryDrillNameResponse), nil
	}
}

type UpdateProtectedInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProtectedInstanceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProtectedInstanceNameInvoker) Invoke() (*model.UpdateProtectedInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectedInstanceNameResponse), nil
	}
}

type UpdateProtectionGroupNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProtectionGroupNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProtectionGroupNameInvoker) Invoke() (*model.UpdateProtectionGroupNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectionGroupNameResponse), nil
	}
}

type UpdateReplicationNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateReplicationNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateReplicationNameInvoker) Invoke() (*model.UpdateReplicationNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateReplicationNameResponse), nil
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

type ShowSpecifiedApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpecifiedApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSpecifiedApiVersionInvoker) Invoke() (*model.ShowSpecifiedApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpecifiedApiVersionResponse), nil
	}
}

type ShowJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobStatusInvoker) Invoke() (*model.ShowJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobStatusResponse), nil
	}
}
