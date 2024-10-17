package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/er/v3/model"
)

type AssociateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRouteTableInvoker) Invoke() (*model.AssociateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRouteTableResponse), nil
	}
}

type DisassociateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateRouteTableInvoker) Invoke() (*model.DisassociateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRouteTableResponse), nil
	}
}

type ListAssociationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociationsInvoker) Invoke() (*model.ListAssociationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociationsResponse), nil
	}
}

type AcceptAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptAttachmentInvoker) Invoke() (*model.AcceptAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptAttachmentResponse), nil
	}
}

type ListAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttachmentsInvoker) Invoke() (*model.ListAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachmentsResponse), nil
	}
}

type RejectAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RejectAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RejectAttachmentInvoker) Invoke() (*model.RejectAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RejectAttachmentResponse), nil
	}
}

type ShowAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttachmentInvoker) Invoke() (*model.ShowAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttachmentResponse), nil
	}
}

type UpdateAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAttachmentInvoker) Invoke() (*model.UpdateAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAttachmentResponse), nil
	}
}

type ListAvailabilityZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailabilityZoneInvoker) Invoke() (*model.ListAvailabilityZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneResponse), nil
	}
}

type ChangeAvailabilityZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAvailabilityZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAvailabilityZoneInvoker) Invoke() (*model.ChangeAvailabilityZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAvailabilityZoneResponse), nil
	}
}

type CreateEnterpriseRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnterpriseRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnterpriseRouterInvoker) Invoke() (*model.CreateEnterpriseRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnterpriseRouterResponse), nil
	}
}

type DeleteEnterpriseRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnterpriseRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEnterpriseRouterInvoker) Invoke() (*model.DeleteEnterpriseRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnterpriseRouterResponse), nil
	}
}

type ListEnterpriseRoutersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseRoutersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseRoutersInvoker) Invoke() (*model.ListEnterpriseRoutersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseRoutersResponse), nil
	}
}

type ShowEnterpriseRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnterpriseRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEnterpriseRouterInvoker) Invoke() (*model.ShowEnterpriseRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseRouterResponse), nil
	}
}

type UpdateEnterpriseRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnterpriseRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnterpriseRouterInvoker) Invoke() (*model.UpdateEnterpriseRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnterpriseRouterResponse), nil
	}
}

type CreateFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlowLogInvoker) Invoke() (*model.CreateFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowLogResponse), nil
	}
}

type DeleteFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlowLogInvoker) Invoke() (*model.DeleteFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlowLogResponse), nil
	}
}

type DisableFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableFlowLogInvoker) Invoke() (*model.DisableFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableFlowLogResponse), nil
	}
}

type EnableFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableFlowLogInvoker) Invoke() (*model.EnableFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableFlowLogResponse), nil
	}
}

type ListFlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlowLogsInvoker) Invoke() (*model.ListFlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowLogsResponse), nil
	}
}

type ShowFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowLogInvoker) Invoke() (*model.ShowFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowLogResponse), nil
	}
}

type UpdateFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFlowLogInvoker) Invoke() (*model.UpdateFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlowLogResponse), nil
	}
}

type DisablePropagationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisablePropagationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisablePropagationInvoker) Invoke() (*model.DisablePropagationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisablePropagationResponse), nil
	}
}

type EnablePropagationInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnablePropagationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnablePropagationInvoker) Invoke() (*model.EnablePropagationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnablePropagationResponse), nil
	}
}

type ListPropagationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPropagationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPropagationsInvoker) Invoke() (*model.ListPropagationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPropagationsResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type CreateStaticRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStaticRouteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStaticRouteInvoker) Invoke() (*model.CreateStaticRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStaticRouteResponse), nil
	}
}

type DeleteStaticRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStaticRouteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStaticRouteInvoker) Invoke() (*model.DeleteStaticRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStaticRouteResponse), nil
	}
}

type ListEffectiveRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEffectiveRoutesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEffectiveRoutesInvoker) Invoke() (*model.ListEffectiveRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEffectiveRoutesResponse), nil
	}
}

type ListStaticRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStaticRoutesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStaticRoutesInvoker) Invoke() (*model.ListStaticRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStaticRoutesResponse), nil
	}
}

type ShowStaticRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStaticRouteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStaticRouteInvoker) Invoke() (*model.ShowStaticRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStaticRouteResponse), nil
	}
}

type UpdateStaticRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStaticRouteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStaticRouteInvoker) Invoke() (*model.UpdateStaticRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStaticRouteResponse), nil
	}
}

type CreateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRouteTableInvoker) Invoke() (*model.CreateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRouteTableResponse), nil
	}
}

type DeleteRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRouteTableInvoker) Invoke() (*model.DeleteRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRouteTableResponse), nil
	}
}

type ListRouteTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRouteTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRouteTablesInvoker) Invoke() (*model.ListRouteTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRouteTablesResponse), nil
	}
}

type ShowRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRouteTableInvoker) Invoke() (*model.ShowRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRouteTableResponse), nil
	}
}

type UpdateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRouteTableInvoker) Invoke() (*model.UpdateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRouteTableResponse), nil
	}
}

type BatchCreateResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateResourceTagsInvoker) Invoke() (*model.BatchCreateResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateResourceTagsResponse), nil
	}
}

type CreateResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceTagInvoker) Invoke() (*model.CreateResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceTagResponse), nil
	}
}

type DeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceTagInvoker) Invoke() (*model.DeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ShowResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceTagInvoker) Invoke() (*model.ShowResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagResponse), nil
	}
}

type CreateVpcAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVpcAttachmentInvoker) Invoke() (*model.CreateVpcAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcAttachmentResponse), nil
	}
}

type DeleteVpcAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVpcAttachmentInvoker) Invoke() (*model.DeleteVpcAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcAttachmentResponse), nil
	}
}

type ListVpcAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVpcAttachmentsInvoker) Invoke() (*model.ListVpcAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcAttachmentsResponse), nil
	}
}

type ShowVpcAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVpcAttachmentInvoker) Invoke() (*model.ShowVpcAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcAttachmentResponse), nil
	}
}

type UpdateVpcAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVpcAttachmentInvoker) Invoke() (*model.UpdateVpcAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcAttachmentResponse), nil
	}
}
