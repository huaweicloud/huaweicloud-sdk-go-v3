package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/eip/v3/model"
)

type ListBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthInvoker) Invoke() (*model.ListBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthResponse), nil
	}
}

type ListCommonPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommonPoolsInvoker) Invoke() (*model.ListCommonPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommonPoolsResponse), nil
	}
}

type ListPublicBorderGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicBorderGroupsInvoker) Invoke() (*model.ListPublicBorderGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicBorderGroupsResponse), nil
	}
}

type ListPublicipPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicipPoolInvoker) Invoke() (*model.ListPublicipPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicipPoolResponse), nil
	}
}

type ListShareBandwidthTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShareBandwidthTypesInvoker) Invoke() (*model.ListShareBandwidthTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShareBandwidthTypesResponse), nil
	}
}

type ShowPublicipPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicipPoolInvoker) Invoke() (*model.ShowPublicipPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicipPoolResponse), nil
	}
}

type AssociatePublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociatePublicipsInvoker) Invoke() (*model.AssociatePublicipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociatePublicipsResponse), nil
	}
}

type AttachBatchPublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachBatchPublicIpInvoker) Invoke() (*model.AttachBatchPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachBatchPublicIpResponse), nil
	}
}

type AttachShareBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachShareBandwidthInvoker) Invoke() (*model.AttachShareBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachShareBandwidthResponse), nil
	}
}

type CountEipAvailableResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountEipAvailableResourcesInvoker) Invoke() (*model.CountEipAvailableResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountEipAvailableResourcesResponse), nil
	}
}

type DetachBatchPublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachBatchPublicIpInvoker) Invoke() (*model.DetachBatchPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachBatchPublicIpResponse), nil
	}
}

type DetachShareBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachShareBandwidthInvoker) Invoke() (*model.DetachShareBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachShareBandwidthResponse), nil
	}
}

type DisableNat64Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisableNat64Invoker) Invoke() (*model.DisableNat64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableNat64Response), nil
	}
}

type DisassociatePublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociatePublicipsInvoker) Invoke() (*model.DisassociatePublicipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociatePublicipsResponse), nil
	}
}

type EnableNat64Invoker struct {
	*invoker.BaseInvoker
}

func (i *EnableNat64Invoker) Invoke() (*model.EnableNat64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableNat64Response), nil
	}
}

type ListPublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicipsInvoker) Invoke() (*model.ListPublicipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicipsResponse), nil
	}
}

type ShowPublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicipInvoker) Invoke() (*model.ShowPublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicipResponse), nil
	}
}

type UpdateAssociatePublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssociatePublicipInvoker) Invoke() (*model.UpdateAssociatePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssociatePublicipResponse), nil
	}
}

type UpdateDisassociatePublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDisassociatePublicipInvoker) Invoke() (*model.UpdateDisassociatePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDisassociatePublicipResponse), nil
	}
}
