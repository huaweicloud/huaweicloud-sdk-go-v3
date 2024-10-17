package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

type AddPublicipsIntoSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPublicipsIntoSharedBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddPublicipsIntoSharedBandwidthInvoker) Invoke() (*model.AddPublicipsIntoSharedBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPublicipsIntoSharedBandwidthResponse), nil
	}
}

type BatchCreateSharedBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSharedBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateSharedBandwidthsInvoker) Invoke() (*model.BatchCreateSharedBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSharedBandwidthsResponse), nil
	}
}

type BatchModifyBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchModifyBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchModifyBandwidthInvoker) Invoke() (*model.BatchModifyBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchModifyBandwidthResponse), nil
	}
}

type ChangeBandwidthToPeriodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeBandwidthToPeriodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeBandwidthToPeriodInvoker) Invoke() (*model.ChangeBandwidthToPeriodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeBandwidthToPeriodResponse), nil
	}
}

type CreateSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSharedBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSharedBandwidthInvoker) Invoke() (*model.CreateSharedBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSharedBandwidthResponse), nil
	}
}

type DeleteSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSharedBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSharedBandwidthInvoker) Invoke() (*model.DeleteSharedBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSharedBandwidthResponse), nil
	}
}

type ListBandwidthPkgInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthPkgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthPkgInvoker) Invoke() (*model.ListBandwidthPkgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthPkgResponse), nil
	}
}

type ListBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthsInvoker) Invoke() (*model.ListBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthsResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type RemovePublicipsFromSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemovePublicipsFromSharedBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemovePublicipsFromSharedBandwidthInvoker) Invoke() (*model.RemovePublicipsFromSharedBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemovePublicipsFromSharedBandwidthResponse), nil
	}
}

type ShowBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBandwidthInvoker) Invoke() (*model.ShowBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthResponse), nil
	}
}

type UpdateBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBandwidthInvoker) Invoke() (*model.UpdateBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBandwidthResponse), nil
	}
}

type UpdatePrePaidBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrePaidBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrePaidBandwidthInvoker) Invoke() (*model.UpdatePrePaidBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrePaidBandwidthResponse), nil
	}
}

type BatchCreatePublicipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreatePublicipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreatePublicipTagsInvoker) Invoke() (*model.BatchCreatePublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePublicipTagsResponse), nil
	}
}

type BatchCreatePublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreatePublicipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreatePublicipsInvoker) Invoke() (*model.BatchCreatePublicipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePublicipsResponse), nil
	}
}

type BatchDeletePublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePublicIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePublicIpInvoker) Invoke() (*model.BatchDeletePublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePublicIpResponse), nil
	}
}

type BatchDeletePublicipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePublicipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePublicipTagsInvoker) Invoke() (*model.BatchDeletePublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePublicipTagsResponse), nil
	}
}

type BatchDisassociatePublicipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociatePublicipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisassociatePublicipsInvoker) Invoke() (*model.BatchDisassociatePublicipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociatePublicipsResponse), nil
	}
}

type ChangePublicipToPeriodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePublicipToPeriodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePublicipToPeriodInvoker) Invoke() (*model.ChangePublicipToPeriodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePublicipToPeriodResponse), nil
	}
}

type CountPublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountPublicIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountPublicIpInvoker) Invoke() (*model.CountPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountPublicIpResponse), nil
	}
}

type CountPublicIpInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountPublicIpInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountPublicIpInstanceInvoker) Invoke() (*model.CountPublicIpInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountPublicIpInstanceResponse), nil
	}
}

type CreatePrePaidPublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrePaidPublicipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrePaidPublicipInvoker) Invoke() (*model.CreatePrePaidPublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrePaidPublicipResponse), nil
	}
}

type CreatePublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePublicipInvoker) Invoke() (*model.CreatePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicipResponse), nil
	}
}

type CreatePublicipTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicipTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePublicipTagInvoker) Invoke() (*model.CreatePublicipTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicipTagResponse), nil
	}
}

type DeletePublicipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePublicipInvoker) Invoke() (*model.DeletePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicipResponse), nil
	}
}

type DeletePublicipTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicipTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePublicipTagInvoker) Invoke() (*model.DeletePublicipTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicipTagResponse), nil
	}
}

type ListPublicipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublicipTagsInvoker) Invoke() (*model.ListPublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicipTagsResponse), nil
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

type ListPublicipsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicipsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublicipsByTagsInvoker) Invoke() (*model.ListPublicipsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicipsByTagsResponse), nil
	}
}

type ShowPublicIpTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicIpTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicIpTypeInvoker) Invoke() (*model.ShowPublicIpTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicIpTypeResponse), nil
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

type ShowPublicipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicipTagsInvoker) Invoke() (*model.ShowPublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicipTagsResponse), nil
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

type ShowResourcesJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourcesJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourcesJobDetailInvoker) Invoke() (*model.ShowResourcesJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourcesJobDetailResponse), nil
	}
}

type NeutronCreateFloatingIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateFloatingIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NeutronCreateFloatingIpInvoker) Invoke() (*model.NeutronCreateFloatingIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateFloatingIpResponse), nil
	}
}

type NeutronDeleteFloatingIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteFloatingIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NeutronDeleteFloatingIpInvoker) Invoke() (*model.NeutronDeleteFloatingIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteFloatingIpResponse), nil
	}
}

type NeutronListFloatingIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListFloatingIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NeutronListFloatingIpsInvoker) Invoke() (*model.NeutronListFloatingIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListFloatingIpsResponse), nil
	}
}

type NeutronShowFloatingIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowFloatingIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NeutronShowFloatingIpInvoker) Invoke() (*model.NeutronShowFloatingIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowFloatingIpResponse), nil
	}
}

type NeutronUpdateFloatingIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateFloatingIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NeutronUpdateFloatingIpInvoker) Invoke() (*model.NeutronUpdateFloatingIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateFloatingIpResponse), nil
	}
}
