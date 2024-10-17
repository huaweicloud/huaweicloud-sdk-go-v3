package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/model"
)

type ListBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthInvoker) Invoke() (*model.ListBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthResponse), nil
	}
}

type ListBandwidthsLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthsLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthsLimitInvoker) Invoke() (*model.ListBandwidthsLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthsLimitResponse), nil
	}
}

type ListCommonPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommonPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommonPoolsInvoker) Invoke() (*model.ListCommonPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommonPoolsResponse), nil
	}
}

type ListEipBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEipBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEipBandwidthsInvoker) Invoke() (*model.ListEipBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEipBandwidthsResponse), nil
	}
}

type ListPublicBorderGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicBorderGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPublicipPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListShareBandwidthTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPublicipPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicipPoolInvoker) Invoke() (*model.ShowPublicipPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicipPoolResponse), nil
	}
}

type ListProjectGeipBindingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectGeipBindingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectGeipBindingsInvoker) Invoke() (*model.ListProjectGeipBindingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectGeipBindingsResponse), nil
	}
}

type CreateTenantVpcIgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTenantVpcIgwInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTenantVpcIgwInvoker) Invoke() (*model.CreateTenantVpcIgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTenantVpcIgwResponse), nil
	}
}

type DeleteTenantVpcIgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTenantVpcIgwInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTenantVpcIgwInvoker) Invoke() (*model.DeleteTenantVpcIgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTenantVpcIgwResponse), nil
	}
}

type ListTenantVpcIgwsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantVpcIgwsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantVpcIgwsInvoker) Invoke() (*model.ListTenantVpcIgwsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantVpcIgwsResponse), nil
	}
}

type ShowInternalVpcIgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInternalVpcIgwInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInternalVpcIgwInvoker) Invoke() (*model.ShowInternalVpcIgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInternalVpcIgwResponse), nil
	}
}

type UpdateTenantVpcIgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantVpcIgwInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTenantVpcIgwInvoker) Invoke() (*model.UpdateTenantVpcIgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantVpcIgwResponse), nil
	}
}

type AssociatePublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociatePublicipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AttachBatchPublicIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AttachShareBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CountEipAvailableResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DetachBatchPublicIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DetachShareBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DisableNat64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DisassociatePublicipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *EnableNat64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPublicipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPublicipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicipInvoker) Invoke() (*model.ShowPublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicipResponse), nil
	}
}

type UpdatePublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePublicipInvoker) Invoke() (*model.UpdatePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicipResponse), nil
	}
}
