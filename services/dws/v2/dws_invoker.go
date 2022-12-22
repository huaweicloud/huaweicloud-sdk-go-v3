package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type AddWorkloadQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddWorkloadQueueInvoker) Invoke() (*model.AddWorkloadQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddWorkloadQueueResponse), nil
	}
}

type AssociateEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateEipInvoker) Invoke() (*model.AssociateEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateEipResponse), nil
	}
}

type AssociateElbInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateElbInvoker) Invoke() (*model.AssociateElbResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateElbResponse), nil
	}
}

type BatchCreateClusterCnInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateClusterCnInvoker) Invoke() (*model.BatchCreateClusterCnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateClusterCnResponse), nil
	}
}

type BatchCreateResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateResourceTagInvoker) Invoke() (*model.BatchCreateResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateResourceTagResponse), nil
	}
}

type BatchDeleteClusterCnInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteClusterCnInvoker) Invoke() (*model.BatchDeleteClusterCnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteClusterCnResponse), nil
	}
}

type BatchDeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourceTagInvoker) Invoke() (*model.BatchDeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourceTagResponse), nil
	}
}

type CancelReadonlyClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelReadonlyClusterInvoker) Invoke() (*model.CancelReadonlyClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelReadonlyClusterResponse), nil
	}
}

type CheckClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckClusterInvoker) Invoke() (*model.CheckClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckClusterResponse), nil
	}
}

type CopySnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopySnapshotInvoker) Invoke() (*model.CopySnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopySnapshotResponse), nil
	}
}

type CreateAlarmSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlarmSubInvoker) Invoke() (*model.CreateAlarmSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmSubResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateClusterDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterDnsInvoker) Invoke() (*model.CreateClusterDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterDnsResponse), nil
	}
}

type CreateClusterWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterWorkloadInvoker) Invoke() (*model.CreateClusterWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterWorkloadResponse), nil
	}
}

type CreateDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataSourceInvoker) Invoke() (*model.CreateDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataSourceResponse), nil
	}
}

type CreateDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDisasterRecoveryInvoker) Invoke() (*model.CreateDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDisasterRecoveryResponse), nil
	}
}

type CreateEventSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSubInvoker) Invoke() (*model.CreateEventSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSubResponse), nil
	}
}

type CreateSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotInvoker) Invoke() (*model.CreateSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotResponse), nil
	}
}

type CreateSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotPolicyInvoker) Invoke() (*model.CreateSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotPolicyResponse), nil
	}
}

type CreateWorkloadPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkloadPlanInvoker) Invoke() (*model.CreateWorkloadPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkloadPlanResponse), nil
	}
}

type DeleteAlarmSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmSubInvoker) Invoke() (*model.DeleteAlarmSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmSubResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
	}
}

type DeleteClusterDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterDnsInvoker) Invoke() (*model.DeleteClusterDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterDnsResponse), nil
	}
}

type DeleteDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataSourceInvoker) Invoke() (*model.DeleteDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataSourceResponse), nil
	}
}

type DeleteDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDisasterRecoveryInvoker) Invoke() (*model.DeleteDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDisasterRecoveryResponse), nil
	}
}

type DeleteEventSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSubInvoker) Invoke() (*model.DeleteEventSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSubResponse), nil
	}
}

type DeleteSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotInvoker) Invoke() (*model.DeleteSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotResponse), nil
	}
}

type DeleteSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotPolicyInvoker) Invoke() (*model.DeleteSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotPolicyResponse), nil
	}
}

type DeleteWorkloadQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkloadQueueInvoker) Invoke() (*model.DeleteWorkloadQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkloadQueueResponse), nil
	}
}

type DisassociateEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateEipInvoker) Invoke() (*model.DisassociateEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateEipResponse), nil
	}
}

type DisassociateElbInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateElbInvoker) Invoke() (*model.DisassociateElbResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateElbResponse), nil
	}
}

type ExecuteRedistributionClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteRedistributionClusterInvoker) Invoke() (*model.ExecuteRedistributionClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteRedistributionClusterResponse), nil
	}
}

type ExpandInstanceStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandInstanceStorageInvoker) Invoke() (*model.ExpandInstanceStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandInstanceStorageResponse), nil
	}
}

type ListAlarmConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmConfigsInvoker) Invoke() (*model.ListAlarmConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmConfigsResponse), nil
	}
}

type ListAlarmDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmDetailInvoker) Invoke() (*model.ListAlarmDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmDetailResponse), nil
	}
}

type ListAlarmStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmStatisticInvoker) Invoke() (*model.ListAlarmStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmStatisticResponse), nil
	}
}

type ListAlarmSubsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmSubsInvoker) Invoke() (*model.ListAlarmSubsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmSubsResponse), nil
	}
}

type ListAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditLogInvoker) Invoke() (*model.ListAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditLogResponse), nil
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

type ListClusterCnInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterCnInvoker) Invoke() (*model.ListClusterCnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterCnResponse), nil
	}
}

type ListClusterConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterConfigurationsInvoker) Invoke() (*model.ListClusterConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterConfigurationsResponse), nil
	}
}

type ListClusterConfigurationsParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterConfigurationsParameterInvoker) Invoke() (*model.ListClusterConfigurationsParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterConfigurationsParameterResponse), nil
	}
}

type ListClusterDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterDetailsInvoker) Invoke() (*model.ListClusterDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterDetailsResponse), nil
	}
}

type ListClusterScaleInNumbersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterScaleInNumbersInvoker) Invoke() (*model.ListClusterScaleInNumbersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterScaleInNumbersResponse), nil
	}
}

type ListClusterSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterSnapshotsInvoker) Invoke() (*model.ListClusterSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterSnapshotsResponse), nil
	}
}

type ListClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterTagsInvoker) Invoke() (*model.ListClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterTagsResponse), nil
	}
}

type ListClusterWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterWorkloadInvoker) Invoke() (*model.ListClusterWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterWorkloadResponse), nil
	}
}

type ListClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersInvoker) Invoke() (*model.ListClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersResponse), nil
	}
}

type ListDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataSourceInvoker) Invoke() (*model.ListDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataSourceResponse), nil
	}
}

type ListDisasterRecoverInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDisasterRecoverInvoker) Invoke() (*model.ListDisasterRecoverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDisasterRecoverResponse), nil
	}
}

type ListDssPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDssPoolsInvoker) Invoke() (*model.ListDssPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDssPoolsResponse), nil
	}
}

type ListElbsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListElbsInvoker) Invoke() (*model.ListElbsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListElbsResponse), nil
	}
}

type ListEventSpecsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSpecsInvoker) Invoke() (*model.ListEventSpecsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSpecsResponse), nil
	}
}

type ListEventSubsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSubsInvoker) Invoke() (*model.ListEventSubsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSubsResponse), nil
	}
}

type ListEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventsInvoker) Invoke() (*model.ListEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventsResponse), nil
	}
}

type ListHostDiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostDiskInvoker) Invoke() (*model.ListHostDiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostDiskResponse), nil
	}
}

type ListHostNetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostNetInvoker) Invoke() (*model.ListHostNetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostNetResponse), nil
	}
}

type ListHostOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostOverviewInvoker) Invoke() (*model.ListHostOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostOverviewResponse), nil
	}
}

type ListJobDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobDetailsInvoker) Invoke() (*model.ListJobDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobDetailsResponse), nil
	}
}

type ListNodeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeTypesInvoker) Invoke() (*model.ListNodeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeTypesResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSnapshotDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotDetailsInvoker) Invoke() (*model.ListSnapshotDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotDetailsResponse), nil
	}
}

type ListSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotPolicyInvoker) Invoke() (*model.ListSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotPolicyResponse), nil
	}
}

type ListSnapshotStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotStatisticsInvoker) Invoke() (*model.ListSnapshotStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotStatisticsResponse), nil
	}
}

type ListSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotsInvoker) Invoke() (*model.ListSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotsResponse), nil
	}
}

type ListStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatisticsInvoker) Invoke() (*model.ListStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatisticsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListWorkloadQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkloadQueueInvoker) Invoke() (*model.ListWorkloadQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkloadQueueResponse), nil
	}
}

type PauseDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseDisasterRecoveryInvoker) Invoke() (*model.PauseDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseDisasterRecoveryResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeClusterInvoker) Invoke() (*model.ResizeClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeClusterResponse), nil
	}
}

type RestartClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartClusterInvoker) Invoke() (*model.RestartClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartClusterResponse), nil
	}
}

type RestoreClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreClusterInvoker) Invoke() (*model.RestoreClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreClusterResponse), nil
	}
}

type RestoreDisasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreDisasterInvoker) Invoke() (*model.RestoreDisasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreDisasterResponse), nil
	}
}

type ShrinkClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkClusterInvoker) Invoke() (*model.ShrinkClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkClusterResponse), nil
	}
}

type StartDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDisasterRecoveryInvoker) Invoke() (*model.StartDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDisasterRecoveryResponse), nil
	}
}

type SwitchFailoverDisasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchFailoverDisasterInvoker) Invoke() (*model.SwitchFailoverDisasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchFailoverDisasterResponse), nil
	}
}

type SwitchOverClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchOverClusterInvoker) Invoke() (*model.SwitchOverClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchOverClusterResponse), nil
	}
}

type SwitchoverDisasterRecoveryInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchoverDisasterRecoveryInvoker) Invoke() (*model.SwitchoverDisasterRecoveryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchoverDisasterRecoveryResponse), nil
	}
}

type UpdateAlarmSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmSubInvoker) Invoke() (*model.UpdateAlarmSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmSubResponse), nil
	}
}

type UpdateClusterDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterDnsInvoker) Invoke() (*model.UpdateClusterDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterDnsResponse), nil
	}
}

type UpdateConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationInvoker) Invoke() (*model.UpdateConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationResponse), nil
	}
}

type UpdateDataSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataSourceInvoker) Invoke() (*model.UpdateDataSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataSourceResponse), nil
	}
}

type UpdateEventSubInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventSubInvoker) Invoke() (*model.UpdateEventSubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventSubResponse), nil
	}
}

type UpdateMaintenanceWindowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMaintenanceWindowInvoker) Invoke() (*model.UpdateMaintenanceWindowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMaintenanceWindowResponse), nil
	}
}
