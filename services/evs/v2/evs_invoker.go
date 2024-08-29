package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
)

type BatchCreateVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateVolumeTagsInvoker) Invoke() (*model.BatchCreateVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

type BatchDeleteVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteVolumeTagsInvoker) Invoke() (*model.BatchDeleteVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

type ChangeVolumeChargeModeInvoker struct {
	*invoker.BaseInvoker
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

func (i *CinderShowVolumeTransferInvoker) Invoke() (*model.CinderShowVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderShowVolumeTransferResponse), nil
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

type CreateVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVolumeInvoker) Invoke() (*model.CreateVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVolumeResponse), nil
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

type DeleteVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVolumeInvoker) Invoke() (*model.DeleteVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVolumeResponse), nil
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

type ListVolumeTagsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListVolumesByTagsInvoker) Invoke() (*model.ListVolumesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesByTagsResponse), nil
	}
}

type ModifyVolumeQoSInvoker struct {
	*invoker.BaseInvoker
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

func (i *RetypeVolumeInvoker) Invoke() (*model.RetypeVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetypeVolumeResponse), nil
	}
}

type RollbackSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollbackSnapshotInvoker) Invoke() (*model.RollbackSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollbackSnapshotResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSnapshotInvoker) Invoke() (*model.ShowSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSnapshotResponse), nil
	}
}

type ShowVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeInvoker) Invoke() (*model.ShowVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeResponse), nil
	}
}

type ShowVolumeTagsInvoker struct {
	*invoker.BaseInvoker
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

func (i *UnsubscribePostpaidVolumeInvoker) Invoke() (*model.UnsubscribePostpaidVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribePostpaidVolumeResponse), nil
	}
}

type UpdateSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSnapshotInvoker) Invoke() (*model.UpdateSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSnapshotResponse), nil
	}
}

type UpdateVolumeInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}
