package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/geip/v3/model"
)

type AddInternetBandwidthTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddInternetBandwidthTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddInternetBandwidthTagsInvoker) Invoke() (*model.AddInternetBandwidthTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddInternetBandwidthTagsResponse), nil
	}
}

type BatchCreateInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateInternetBandwidthInvoker) Invoke() (*model.BatchCreateInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateInternetBandwidthResponse), nil
	}
}

type BatchCreateInternetBandwidthTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateInternetBandwidthTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateInternetBandwidthTagsInvoker) Invoke() (*model.BatchCreateInternetBandwidthTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateInternetBandwidthTagsResponse), nil
	}
}

type BatchDeleteInternetBandwidthTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInternetBandwidthTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInternetBandwidthTagsInvoker) Invoke() (*model.BatchDeleteInternetBandwidthTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInternetBandwidthTagsResponse), nil
	}
}

type CountInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountInternetBandwidthInvoker) Invoke() (*model.CountInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountInternetBandwidthResponse), nil
	}
}

type CreateInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInternetBandwidthInvoker) Invoke() (*model.CreateInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInternetBandwidthResponse), nil
	}
}

type CreateUserDisclaimerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserDisclaimerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserDisclaimerInvoker) Invoke() (*model.CreateUserDisclaimerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserDisclaimerResponse), nil
	}
}

type DeleteInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInternetBandwidthInvoker) Invoke() (*model.DeleteInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInternetBandwidthResponse), nil
	}
}

type DeleteInternetBandwidthTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInternetBandwidthTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInternetBandwidthTagInvoker) Invoke() (*model.DeleteInternetBandwidthTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInternetBandwidthTagResponse), nil
	}
}

type DeleteUserDisclaimerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserDisclaimerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserDisclaimerInvoker) Invoke() (*model.DeleteUserDisclaimerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserDisclaimerResponse), nil
	}
}

type ListAccessSitesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessSitesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessSitesInvoker) Invoke() (*model.ListAccessSitesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessSitesResponse), nil
	}
}

type ListGeipResourceQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeipResourceQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeipResourceQuotasInvoker) Invoke() (*model.ListGeipResourceQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeipResourceQuotasResponse), nil
	}
}

type ListInternetBandwidthCountFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetBandwidthCountFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetBandwidthCountFilterTagsInvoker) Invoke() (*model.ListInternetBandwidthCountFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetBandwidthCountFilterTagsResponse), nil
	}
}

type ListInternetBandwidthDomainTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetBandwidthDomainTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetBandwidthDomainTagsInvoker) Invoke() (*model.ListInternetBandwidthDomainTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetBandwidthDomainTagsResponse), nil
	}
}

type ListInternetBandwidthFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetBandwidthFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetBandwidthFilterTagsInvoker) Invoke() (*model.ListInternetBandwidthFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetBandwidthFilterTagsResponse), nil
	}
}

type ListInternetBandwidthLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetBandwidthLimitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetBandwidthLimitsInvoker) Invoke() (*model.ListInternetBandwidthLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetBandwidthLimitsResponse), nil
	}
}

type ListInternetBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInternetBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInternetBandwidthsInvoker) Invoke() (*model.ListInternetBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInternetBandwidthsResponse), nil
	}
}

type ListSupportMasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSupportMasksInvoker) Invoke() (*model.ListSupportMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportMasksResponse), nil
	}
}

type ShowInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInternetBandwidthInvoker) Invoke() (*model.ShowInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInternetBandwidthResponse), nil
	}
}

type ShowInternetBandwidthTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInternetBandwidthTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInternetBandwidthTagsInvoker) Invoke() (*model.ShowInternetBandwidthTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInternetBandwidthTagsResponse), nil
	}
}

type ShowUserDisclaimerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserDisclaimerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserDisclaimerInvoker) Invoke() (*model.ShowUserDisclaimerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserDisclaimerResponse), nil
	}
}

type UpdateInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInternetBandwidthInvoker) Invoke() (*model.UpdateInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInternetBandwidthResponse), nil
	}
}

type AddGeipSegmentTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGeipSegmentTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddGeipSegmentTagsInvoker) Invoke() (*model.AddGeipSegmentTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGeipSegmentTagsResponse), nil
	}
}

type AddGlobalEipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGlobalEipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddGlobalEipTagsInvoker) Invoke() (*model.AddGlobalEipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGlobalEipTagsResponse), nil
	}
}

type AssociateGeipSegmentInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateGeipSegmentInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateGeipSegmentInstanceInvoker) Invoke() (*model.AssociateGeipSegmentInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateGeipSegmentInstanceResponse), nil
	}
}

type AssociateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateInstanceInvoker) Invoke() (*model.AssociateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateInstanceResponse), nil
	}
}

type AttachInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachInternetBandwidthInvoker) Invoke() (*model.AttachInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachInternetBandwidthResponse), nil
	}
}

type BatchAttachGeipSegmentInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAttachGeipSegmentInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAttachGeipSegmentInternetBandwidthInvoker) Invoke() (*model.BatchAttachGeipSegmentInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAttachGeipSegmentInternetBandwidthResponse), nil
	}
}

type BatchAttachInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAttachInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAttachInternetBandwidthInvoker) Invoke() (*model.BatchAttachInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAttachInternetBandwidthResponse), nil
	}
}

type BatchCreateGeipSegmentTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateGeipSegmentTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateGeipSegmentTagsInvoker) Invoke() (*model.BatchCreateGeipSegmentTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateGeipSegmentTagsResponse), nil
	}
}

type BatchCreateGlobalEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateGlobalEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateGlobalEipInvoker) Invoke() (*model.BatchCreateGlobalEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateGlobalEipResponse), nil
	}
}

type BatchCreateGlobalEipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateGlobalEipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateGlobalEipTagsInvoker) Invoke() (*model.BatchCreateGlobalEipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateGlobalEipTagsResponse), nil
	}
}

type BatchDeleteGeipSegmentTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteGeipSegmentTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteGeipSegmentTagsInvoker) Invoke() (*model.BatchDeleteGeipSegmentTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteGeipSegmentTagsResponse), nil
	}
}

type BatchDeleteGlobalEipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteGlobalEipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteGlobalEipTagsInvoker) Invoke() (*model.BatchDeleteGlobalEipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteGlobalEipTagsResponse), nil
	}
}

type BatchDetachGeipSegmentInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDetachGeipSegmentInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDetachGeipSegmentInternetBandwidthInvoker) Invoke() (*model.BatchDetachGeipSegmentInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDetachGeipSegmentInternetBandwidthResponse), nil
	}
}

type BatchDetachInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDetachInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDetachInternetBandwidthInvoker) Invoke() (*model.BatchDetachInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDetachInternetBandwidthResponse), nil
	}
}

type CountGlobalEipSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountGlobalEipSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountGlobalEipSegmentInvoker) Invoke() (*model.CountGlobalEipSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountGlobalEipSegmentResponse), nil
	}
}

type CountGlobalEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountGlobalEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountGlobalEipsInvoker) Invoke() (*model.CountGlobalEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountGlobalEipsResponse), nil
	}
}

type CreateGlobalEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGlobalEipInvoker) Invoke() (*model.CreateGlobalEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalEipResponse), nil
	}
}

type CreateGlobalEipSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalEipSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGlobalEipSegmentInvoker) Invoke() (*model.CreateGlobalEipSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalEipSegmentResponse), nil
	}
}

type DeleteGeipSegmentTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGeipSegmentTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGeipSegmentTagInvoker) Invoke() (*model.DeleteGeipSegmentTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGeipSegmentTagResponse), nil
	}
}

type DeleteGlobalEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGlobalEipInvoker) Invoke() (*model.DeleteGlobalEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalEipResponse), nil
	}
}

type DeleteGlobalEipSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalEipSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGlobalEipSegmentInvoker) Invoke() (*model.DeleteGlobalEipSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalEipSegmentResponse), nil
	}
}

type DeleteGlobalEipTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalEipTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGlobalEipTagInvoker) Invoke() (*model.DeleteGlobalEipTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalEipTagResponse), nil
	}
}

type DetachInternetBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachInternetBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachInternetBandwidthInvoker) Invoke() (*model.DetachInternetBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachInternetBandwidthResponse), nil
	}
}

type DisassociateGeipSegmentInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateGeipSegmentInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateGeipSegmentInstanceInvoker) Invoke() (*model.DisassociateGeipSegmentInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateGeipSegmentInstanceResponse), nil
	}
}

type DisassociateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateInstanceInvoker) Invoke() (*model.DisassociateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateInstanceResponse), nil
	}
}

type ListGeipPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeipPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeipPoolsInvoker) Invoke() (*model.ListGeipPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeipPoolsResponse), nil
	}
}

type ListGeipSegmentCountFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeipSegmentCountFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeipSegmentCountFilterTagsInvoker) Invoke() (*model.ListGeipSegmentCountFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeipSegmentCountFilterTagsResponse), nil
	}
}

type ListGeipSegmentDomainTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeipSegmentDomainTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeipSegmentDomainTagsInvoker) Invoke() (*model.ListGeipSegmentDomainTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeipSegmentDomainTagsResponse), nil
	}
}

type ListGeipSegmentFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeipSegmentFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeipSegmentFilterTagsInvoker) Invoke() (*model.ListGeipSegmentFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeipSegmentFilterTagsResponse), nil
	}
}

type ListGlobalEipCountFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalEipCountFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalEipCountFilterTagsInvoker) Invoke() (*model.ListGlobalEipCountFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalEipCountFilterTagsResponse), nil
	}
}

type ListGlobalEipDomainTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalEipDomainTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalEipDomainTagsInvoker) Invoke() (*model.ListGlobalEipDomainTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalEipDomainTagsResponse), nil
	}
}

type ListGlobalEipFilterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalEipFilterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalEipFilterTagsInvoker) Invoke() (*model.ListGlobalEipFilterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalEipFilterTagsResponse), nil
	}
}

type ListGlobalEipSegmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalEipSegmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalEipSegmentsInvoker) Invoke() (*model.ListGlobalEipSegmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalEipSegmentsResponse), nil
	}
}

type ListGlobalEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalEipsInvoker) Invoke() (*model.ListGlobalEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalEipsResponse), nil
	}
}

type ShowGeipSegmentTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGeipSegmentTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGeipSegmentTagsInvoker) Invoke() (*model.ShowGeipSegmentTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGeipSegmentTagsResponse), nil
	}
}

type ShowGlobalEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalEipInvoker) Invoke() (*model.ShowGlobalEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalEipResponse), nil
	}
}

type ShowGlobalEipSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalEipSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalEipSegmentInvoker) Invoke() (*model.ShowGlobalEipSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalEipSegmentResponse), nil
	}
}

type ShowGlobalEipTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalEipTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalEipTagsInvoker) Invoke() (*model.ShowGlobalEipTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalEipTagsResponse), nil
	}
}

type UpdateGlobalEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGlobalEipInvoker) Invoke() (*model.UpdateGlobalEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalEipResponse), nil
	}
}

type UpdateGlobalEipSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalEipSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGlobalEipSegmentInvoker) Invoke() (*model.UpdateGlobalEipSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalEipSegmentResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ShowJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobsInvoker) Invoke() (*model.ShowJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobsResponse), nil
	}
}

type ListSupportRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSupportRegionsInvoker) Invoke() (*model.ListSupportRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportRegionsResponse), nil
	}
}

type ListTenantGeipSupportInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantGeipSupportInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantGeipSupportInstancesInvoker) Invoke() (*model.ListTenantGeipSupportInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantGeipSupportInstancesResponse), nil
	}
}
