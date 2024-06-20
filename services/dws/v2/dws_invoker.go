package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type AddQueueUserListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddQueueUserListInvoker) Invoke() (*model.AddQueueUserListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddQueueUserListResponse), nil
	}
}

type AddSnapshotCrossRegionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSnapshotCrossRegionPolicyInvoker) Invoke() (*model.AddSnapshotCrossRegionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSnapshotCrossRegionPolicyResponse), nil
	}
}

type AddWorkloadPlanStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddWorkloadPlanStageInvoker) Invoke() (*model.AddWorkloadPlanStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddWorkloadPlanStageResponse), nil
	}
}

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

type ChangeSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) Invoke() (*model.ChangeSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSecurityGroupResponse), nil
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

type CheckDisasterNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDisasterNameInvoker) Invoke() (*model.CheckDisasterNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDisasterNameResponse), nil
	}
}

type CheckTableRestoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTableRestoreInvoker) Invoke() (*model.CheckTableRestoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTableRestoreResponse), nil
	}
}

type ConvertToLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConvertToLogicalClusterInvoker) Invoke() (*model.ConvertToLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConvertToLogicalClusterResponse), nil
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

type CreateClusterV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterV2Invoker) Invoke() (*model.CreateClusterV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterV2Response), nil
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

type CreateLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogicalClusterInvoker) Invoke() (*model.CreateLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogicalClusterResponse), nil
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

type DeleteClusterNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterNodesInvoker) Invoke() (*model.DeleteClusterNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterNodesResponse), nil
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

type DeleteLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogicalClusterInvoker) Invoke() (*model.DeleteLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogicalClusterResponse), nil
	}
}

type DeleteQueueUserListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueueUserListInvoker) Invoke() (*model.DeleteQueueUserListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueueUserListResponse), nil
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

type DeleteSnapshotCrossRegionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotCrossRegionPolicyInvoker) Invoke() (*model.DeleteSnapshotCrossRegionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotCrossRegionPolicyResponse), nil
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

type DeleteWorkloadPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkloadPlanInvoker) Invoke() (*model.DeleteWorkloadPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkloadPlanResponse), nil
	}
}

type DeleteWorkloadPlanStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkloadPlanStageInvoker) Invoke() (*model.DeleteWorkloadPlanStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkloadPlanStageResponse), nil
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

type DisableLtsLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableLtsLogsInvoker) Invoke() (*model.DisableLtsLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableLtsLogsResponse), nil
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

type EnableLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableLogicalClusterInvoker) Invoke() (*model.EnableLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableLogicalClusterResponse), nil
	}
}

type EnableLtsLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableLtsLogsInvoker) Invoke() (*model.EnableLtsLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableLtsLogsResponse), nil
	}
}

type ExecuteClusterUpgradeActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteClusterUpgradeActionInvoker) Invoke() (*model.ExecuteClusterUpgradeActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteClusterUpgradeActionResponse), nil
	}
}

type ExecuteDatabaseOmUserActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDatabaseOmUserActionInvoker) Invoke() (*model.ExecuteDatabaseOmUserActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDatabaseOmUserActionResponse), nil
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

type ListAvailableDisasterClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableDisasterClustersInvoker) Invoke() (*model.ListAvailableDisasterClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableDisasterClustersResponse), nil
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

type ListClusterNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterNodesInvoker) Invoke() (*model.ListClusterNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterNodesResponse), nil
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

type ListConfigurationsAuditRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsAuditRecordsInvoker) Invoke() (*model.ListConfigurationsAuditRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsAuditRecordsResponse), nil
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

type ListDatabaseUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseUsersInvoker) Invoke() (*model.ListDatabaseUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseUsersResponse), nil
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

type ListLogicalClusterRingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogicalClusterRingsInvoker) Invoke() (*model.ListLogicalClusterRingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogicalClusterRingsResponse), nil
	}
}

type ListLogicalClusterTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogicalClusterTasksInvoker) Invoke() (*model.ListLogicalClusterTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogicalClusterTasksResponse), nil
	}
}

type ListLogicalClusterVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogicalClusterVolumesInvoker) Invoke() (*model.ListLogicalClusterVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogicalClusterVolumesResponse), nil
	}
}

type ListLogicalClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogicalClustersInvoker) Invoke() (*model.ListLogicalClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogicalClustersResponse), nil
	}
}

type ListLtsLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsLogsInvoker) Invoke() (*model.ListLtsLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsLogsResponse), nil
	}
}

type ListMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsInvoker) Invoke() (*model.ListMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsResponse), nil
	}
}

type ListMetricsDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsDataInvoker) Invoke() (*model.ListMetricsDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsDataResponse), nil
	}
}

type ListMonitorIndicatorDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitorIndicatorDataInvoker) Invoke() (*model.ListMonitorIndicatorDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitorIndicatorDataResponse), nil
	}
}

type ListMonitorIndicatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitorIndicatorsInvoker) Invoke() (*model.ListMonitorIndicatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitorIndicatorsResponse), nil
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

type ListPlanExecLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlanExecLogsInvoker) Invoke() (*model.ListPlanExecLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlanExecLogsResponse), nil
	}
}

type ListQueriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueriesInvoker) Invoke() (*model.ListQueriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueriesResponse), nil
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

type ListSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSchemasInvoker) Invoke() (*model.ListSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSchemasResponse), nil
	}
}

type ListSnapshotCrossRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotCrossRegionInvoker) Invoke() (*model.ListSnapshotCrossRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotCrossRegionResponse), nil
	}
}

type ListSnapshotCrossRegionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotCrossRegionPolicyInvoker) Invoke() (*model.ListSnapshotCrossRegionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotCrossRegionPolicyResponse), nil
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

type ListTablesStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTablesStatisticInvoker) Invoke() (*model.ListTablesStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTablesStatisticResponse), nil
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

type ListTopoRingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopoRingsInvoker) Invoke() (*model.ListTopoRingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopoRingsResponse), nil
	}
}

type ListUpdatableVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUpdatableVersionInvoker) Invoke() (*model.ListUpdatableVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUpdatableVersionResponse), nil
	}
}

type ListUpdateRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUpdateRecordInvoker) Invoke() (*model.ListUpdateRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUpdateRecordResponse), nil
	}
}

type ListWorkloadPlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkloadPlansInvoker) Invoke() (*model.ListWorkloadPlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkloadPlansResponse), nil
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

type ListWorkloadQueueUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkloadQueueUsersInvoker) Invoke() (*model.ListWorkloadQueueUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkloadQueueUsersResponse), nil
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

type ResizeClusterWithExistedNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeClusterWithExistedNodesInvoker) Invoke() (*model.ResizeClusterWithExistedNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeClusterWithExistedNodesResponse), nil
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

type RestartLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartLogicalClusterInvoker) Invoke() (*model.RestartLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartLogicalClusterResponse), nil
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

type RestoreRedistributionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreRedistributionInvoker) Invoke() (*model.RestoreRedistributionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreRedistributionResponse), nil
	}
}

type RestoreTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreTableInvoker) Invoke() (*model.RestoreTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreTableResponse), nil
	}
}

type SaveClusterDescriptionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveClusterDescriptionInfoInvoker) Invoke() (*model.SaveClusterDescriptionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveClusterDescriptionInfoResponse), nil
	}
}

type ShowClusterFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterFlavorInvoker) Invoke() (*model.ShowClusterFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterFlavorResponse), nil
	}
}

type ShowClusterRedistributionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterRedistributionInvoker) Invoke() (*model.ShowClusterRedistributionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterRedistributionResponse), nil
	}
}

type ShowDatabaseAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseAuthorityInvoker) Invoke() (*model.ShowDatabaseAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseAuthorityResponse), nil
	}
}

type ShowDatabaseOmUserStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseOmUserStatusInvoker) Invoke() (*model.ShowDatabaseOmUserStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseOmUserStatusResponse), nil
	}
}

type ShowDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseUserInvoker) Invoke() (*model.ShowDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseUserResponse), nil
	}
}

type ShowDisasterDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDisasterDetailInvoker) Invoke() (*model.ShowDisasterDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDisasterDetailResponse), nil
	}
}

type ShowDisasterProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDisasterProgressInvoker) Invoke() (*model.ShowDisasterProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDisasterProgressResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowQueryDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueryDetailInvoker) Invoke() (*model.ShowQueryDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueryDetailResponse), nil
	}
}

type ShowResourceStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceStatisticsInvoker) Invoke() (*model.ShowResourceStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceStatisticsResponse), nil
	}
}

type ShowWorkloadPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkloadPlanInvoker) Invoke() (*model.ShowWorkloadPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkloadPlanResponse), nil
	}
}

type ShowWorkloadPlanStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkloadPlanStageInvoker) Invoke() (*model.ShowWorkloadPlanStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkloadPlanStageResponse), nil
	}
}

type ShowWorkloadQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkloadQueueInvoker) Invoke() (*model.ShowWorkloadQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkloadQueueResponse), nil
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

type StartWorkloadPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartWorkloadPlanInvoker) Invoke() (*model.StartWorkloadPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartWorkloadPlanResponse), nil
	}
}

type StopRedistributionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopRedistributionInvoker) Invoke() (*model.StopRedistributionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopRedistributionResponse), nil
	}
}

type StopWorkloadPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopWorkloadPlanInvoker) Invoke() (*model.StopWorkloadPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopWorkloadPlanResponse), nil
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

type SwitchPlanStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchPlanStageInvoker) Invoke() (*model.SwitchPlanStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchPlanStageResponse), nil
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

type SyncIamUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncIamUsersInvoker) Invoke() (*model.SyncIamUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncIamUsersResponse), nil
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

type UpdateDatabaseAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseAuthorityInvoker) Invoke() (*model.UpdateDatabaseAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseAuthorityResponse), nil
	}
}

type UpdateDatabaseUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseUserInfoInvoker) Invoke() (*model.UpdateDatabaseUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseUserInfoResponse), nil
	}
}

type UpdateDisasterInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDisasterInfoInvoker) Invoke() (*model.UpdateDisasterInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDisasterInfoResponse), nil
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

type UpdateLogicalClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogicalClusterInvoker) Invoke() (*model.UpdateLogicalClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogicalClusterResponse), nil
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

type UpdateQueueResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateQueueResourcesInvoker) Invoke() (*model.UpdateQueueResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQueueResourcesResponse), nil
	}
}

type UpdateSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSchemasInvoker) Invoke() (*model.UpdateSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSchemasResponse), nil
	}
}
