package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
)

type AssociateAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateAutoSnapshotPolicyInvoker) Invoke() (*model.AssociateAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateAutoSnapshotPolicyResponse), nil
	}
}

type BatchCreateSnapshotGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSnapshotGroupTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateSnapshotGroupTagInvoker) Invoke() (*model.BatchCreateSnapshotGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSnapshotGroupTagResponse), nil
	}
}

type BatchCreateSnapshotTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSnapshotTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateSnapshotTagsInvoker) Invoke() (*model.BatchCreateSnapshotTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSnapshotTagsResponse), nil
	}
}

type BatchCreateVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateVolumeTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateVolumeTagsInvoker) Invoke() (*model.BatchCreateVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

type BatchDeleteSnapshotGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSnapshotGroupTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSnapshotGroupTagInvoker) Invoke() (*model.BatchDeleteSnapshotGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSnapshotGroupTagResponse), nil
	}
}

type BatchDeleteSnapshotTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSnapshotTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSnapshotTagsInvoker) Invoke() (*model.BatchDeleteSnapshotTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSnapshotTagsResponse), nil
	}
}

type BatchDeleteVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteVolumeTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteVolumeTagsInvoker) Invoke() (*model.BatchDeleteVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

type BatchResizeVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResizeVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchResizeVolumesInvoker) Invoke() (*model.BatchResizeVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResizeVolumesResponse), nil
	}
}

type ChangeVolumeChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVolumeChargeModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeVolumeChargeModeInvoker) Invoke() (*model.ChangeVolumeChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVolumeChargeModeResponse), nil
	}
}

type CinderAcceptVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderAcceptVolumeTransferInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderAcceptVolumeTransferInvoker) Invoke() (*model.CinderAcceptVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

type CinderCreateVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderCreateVolumeTransferInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderCreateVolumeTransferInvoker) Invoke() (*model.CinderCreateVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

type CinderDeleteVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderDeleteVolumeTransferInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderDeleteVolumeTransferInvoker) Invoke() (*model.CinderDeleteVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

type CinderListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListAvailabilityZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderListAvailabilityZonesInvoker) Invoke() (*model.CinderListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

type CinderListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderListQuotasInvoker) Invoke() (*model.CinderListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListQuotasResponse), nil
	}
}

type CinderListVolumeTransfersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTransfersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderListVolumeTransfersInvoker) Invoke() (*model.CinderListVolumeTransfersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTransfersResponse), nil
	}
}

type CinderListVolumeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderListVolumeTypesInvoker) Invoke() (*model.CinderListVolumeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTypesResponse), nil
	}
}

type CinderShowVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderShowVolumeTransferInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CinderShowVolumeTransferInvoker) Invoke() (*model.CinderShowVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderShowVolumeTransferResponse), nil
	}
}

type CountAutoSnapshotPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountAutoSnapshotPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountAutoSnapshotPoliciesInvoker) Invoke() (*model.CountAutoSnapshotPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountAutoSnapshotPoliciesResponse), nil
	}
}

type CountSnapshotChainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSnapshotChainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSnapshotChainsInvoker) Invoke() (*model.CountSnapshotChainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSnapshotChainsResponse), nil
	}
}

type CountSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSnapshotsInvoker) Invoke() (*model.CountSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSnapshotsResponse), nil
	}
}

type CountSnapshotsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSnapshotsGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSnapshotsGroupInvoker) Invoke() (*model.CountSnapshotsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSnapshotsGroupResponse), nil
	}
}

type CreateAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAutoSnapshotPolicyInvoker) Invoke() (*model.CreateAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoSnapshotPolicyResponse), nil
	}
}

type CreateSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSnapshotInvoker) Invoke() (*model.CreateSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotResponse), nil
	}
}

type CreateSnapshotGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSnapshotGroupInvoker) Invoke() (*model.CreateSnapshotGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotGroupResponse), nil
	}
}

type CreateStandardSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStandardSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStandardSnapshotInvoker) Invoke() (*model.CreateStandardSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStandardSnapshotResponse), nil
	}
}

type CreateVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVolumeInvoker) Invoke() (*model.CreateVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVolumeResponse), nil
	}
}

type DeleteAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAutoSnapshotPolicyInvoker) Invoke() (*model.DeleteAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAutoSnapshotPolicyResponse), nil
	}
}

type DeleteSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSnapshotInvoker) Invoke() (*model.DeleteSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotResponse), nil
	}
}

type DeleteSnapshotGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSnapshotGroupInvoker) Invoke() (*model.DeleteSnapshotGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotGroupResponse), nil
	}
}

type DeleteStandardSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStandardSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStandardSnapshotInvoker) Invoke() (*model.DeleteStandardSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStandardSnapshotResponse), nil
	}
}

type DeleteVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVolumeInvoker) Invoke() (*model.DeleteVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVolumeResponse), nil
	}
}

type DeleteVolumeInRecycleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVolumeInRecycleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVolumeInRecycleInvoker) Invoke() (*model.DeleteVolumeInRecycleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVolumeInRecycleResponse), nil
	}
}

type DisassociateAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateAutoSnapshotPolicyInvoker) Invoke() (*model.DisassociateAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateAutoSnapshotPolicyResponse), nil
	}
}

type ListAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoSnapshotPolicyInvoker) Invoke() (*model.ListAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoSnapshotPolicyResponse), nil
	}
}

type ListSnapshotChainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotChainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotChainsInvoker) Invoke() (*model.ListSnapshotChainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotChainsResponse), nil
	}
}

type ListSnapshotGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotGroupInvoker) Invoke() (*model.ListSnapshotGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotGroupResponse), nil
	}
}

type ListSnapshotGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotGroupTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotGroupTagsInvoker) Invoke() (*model.ListSnapshotGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotGroupTagsResponse), nil
	}
}

type ListSnapshotTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotTagsInvoker) Invoke() (*model.ListSnapshotTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotTagsResponse), nil
	}
}

type ListSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotsInvoker) Invoke() (*model.ListSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotsResponse), nil
	}
}

type ListSnapshotsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotsDetailsInvoker) Invoke() (*model.ListSnapshotsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotsDetailsResponse), nil
	}
}

type ListVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumeTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumeTagsInvoker) Invoke() (*model.ListVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumeTagsResponse), nil
	}
}

type ListVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumesInvoker) Invoke() (*model.ListVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesResponse), nil
	}
}

type ListVolumesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumesByTagsInvoker) Invoke() (*model.ListVolumesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesByTagsResponse), nil
	}
}

type ListVolumesInRecycleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesInRecycleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumesInRecycleInvoker) Invoke() (*model.ListVolumesInRecycleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesInRecycleResponse), nil
	}
}

type ModifyVolumeQoSInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyVolumeQoSInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyVolumeQoSInvoker) Invoke() (*model.ModifyVolumeQoSResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyVolumeQoSResponse), nil
	}
}

type ResizeVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeVolumeInvoker) Invoke() (*model.ResizeVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeVolumeResponse), nil
	}
}

type RetypeVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetypeVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetypeVolumeInvoker) Invoke() (*model.RetypeVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetypeVolumeResponse), nil
	}
}

type RevertVolumeInRecycleInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevertVolumeInRecycleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RevertVolumeInRecycleInvoker) Invoke() (*model.RevertVolumeInRecycleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevertVolumeInRecycleResponse), nil
	}
}

type RollbackSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollbackSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RollbackSnapshotInvoker) Invoke() (*model.RollbackSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollbackSnapshotResponse), nil
	}
}

type RollbackStandardSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollbackStandardSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RollbackStandardSnapshotInvoker) Invoke() (*model.RollbackStandardSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollbackStandardSnapshotResponse), nil
	}
}

type ShowAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSnapshotPolicyInvoker) Invoke() (*model.ShowAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSnapshotPolicyResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) Invoke() (*model.ShowRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecyclePolicyResponse), nil
	}
}

type ShowSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSnapshotInvoker) Invoke() (*model.ShowSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSnapshotResponse), nil
	}
}

type ShowSnapshotGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSnapshotGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSnapshotGroupInvoker) Invoke() (*model.ShowSnapshotGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSnapshotGroupResponse), nil
	}
}

type ShowStandardSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStandardSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStandardSnapshotInvoker) Invoke() (*model.ShowStandardSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStandardSnapshotResponse), nil
	}
}

type ShowVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVolumeInvoker) Invoke() (*model.ShowVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeResponse), nil
	}
}

type ShowVolumeInRecycleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeInRecycleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVolumeInRecycleInvoker) Invoke() (*model.ShowVolumeInRecycleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeInRecycleResponse), nil
	}
}

type ShowVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVolumeTagsInvoker) Invoke() (*model.ShowVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeTagsResponse), nil
	}
}

type UnsubscribePostpaidVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnsubscribePostpaidVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnsubscribePostpaidVolumeInvoker) Invoke() (*model.UnsubscribePostpaidVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribePostpaidVolumeResponse), nil
	}
}

type UpdateAutoSnapshotPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoSnapshotPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAutoSnapshotPolicyInvoker) Invoke() (*model.UpdateAutoSnapshotPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoSnapshotPolicyResponse), nil
	}
}

type UpdateRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecyclePolicyInvoker) Invoke() (*model.UpdateRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecyclePolicyResponse), nil
	}
}

type UpdateSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSnapshotInvoker) Invoke() (*model.UpdateSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSnapshotResponse), nil
	}
}

type UpdateSnapshotGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSnapshotGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSnapshotGroupInvoker) Invoke() (*model.UpdateSnapshotGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSnapshotGroupResponse), nil
	}
}

type UpdateStandardSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStandardSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStandardSnapshotInvoker) Invoke() (*model.UpdateStandardSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStandardSnapshotResponse), nil
	}
}

type UpdateVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVolumeInvoker) Invoke() (*model.UpdateVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVolumeResponse), nil
	}
}

type ListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}
