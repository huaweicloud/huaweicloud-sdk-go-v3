package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

type AddPublicipsIntoSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchCreateSharedBandwidthsInvoker) Invoke() (*model.BatchCreateSharedBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSharedBandwidthsResponse), nil
	}
}

type CreateSharedBandwidthInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteSharedBandwidthInvoker) Invoke() (*model.DeleteSharedBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSharedBandwidthResponse), nil
	}
}

type ListBandwidthsInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchCreatePublicipTagsInvoker) Invoke() (*model.BatchCreatePublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePublicipTagsResponse), nil
	}
}

type BatchDeletePublicipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePublicipTagsInvoker) Invoke() (*model.BatchDeletePublicipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePublicipTagsResponse), nil
	}
}

type CreatePrePaidPublicipInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPublicipsByTagsInvoker) Invoke() (*model.ListPublicipsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicipsByTagsResponse), nil
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

type ShowPublicipTagsInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdatePublicipInvoker) Invoke() (*model.UpdatePublicipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicipResponse), nil
	}
}

type NeutronCreateFloatingIpInvoker struct {
	*invoker.BaseInvoker
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

func (i *NeutronUpdateFloatingIpInvoker) Invoke() (*model.NeutronUpdateFloatingIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateFloatingIpResponse), nil
	}
}
