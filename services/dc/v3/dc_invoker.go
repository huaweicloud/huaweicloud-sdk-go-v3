package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dc/v3/model"
)

type BindGlobalEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindGlobalEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindGlobalEipsInvoker) Invoke() (*model.BindGlobalEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindGlobalEipsResponse), nil
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

type UnbindGlobalEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnbindGlobalEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnbindGlobalEipsInvoker) Invoke() (*model.UnbindGlobalEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnbindGlobalEipsResponse), nil
	}
}

type CreateConnectGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectGatewayInvoker) Invoke() (*model.CreateConnectGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectGatewayResponse), nil
	}
}

type DeleteConnectGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectGatewayInvoker) Invoke() (*model.DeleteConnectGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectGatewayResponse), nil
	}
}

type ListConnectGatewaysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectGatewaysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectGatewaysInvoker) Invoke() (*model.ListConnectGatewaysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectGatewaysResponse), nil
	}
}

type ShowConnectGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConnectGatewayInvoker) Invoke() (*model.ShowConnectGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectGatewayResponse), nil
	}
}

type UpdateConnectGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConnectGatewayInvoker) Invoke() (*model.UpdateConnectGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectGatewayResponse), nil
	}
}

type CreateHostedDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostedDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHostedDirectConnectInvoker) Invoke() (*model.CreateHostedDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostedDirectConnectResponse), nil
	}
}

type DeleteDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDirectConnectInvoker) Invoke() (*model.DeleteDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDirectConnectResponse), nil
	}
}

type DeleteHostedDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostedDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostedDirectConnectInvoker) Invoke() (*model.DeleteHostedDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostedDirectConnectResponse), nil
	}
}

type ListDirectConnectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDirectConnectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDirectConnectsInvoker) Invoke() (*model.ListDirectConnectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDirectConnectsResponse), nil
	}
}

type ListHostedDirectConnectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostedDirectConnectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostedDirectConnectsInvoker) Invoke() (*model.ListHostedDirectConnectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostedDirectConnectsResponse), nil
	}
}

type ShowDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDirectConnectInvoker) Invoke() (*model.ShowDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDirectConnectResponse), nil
	}
}

type ShowHostedDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostedDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostedDirectConnectInvoker) Invoke() (*model.ShowHostedDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostedDirectConnectResponse), nil
	}
}

type UpdateDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDirectConnectInvoker) Invoke() (*model.UpdateDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDirectConnectResponse), nil
	}
}

type UpdateHostedDirectConnectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostedDirectConnectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostedDirectConnectInvoker) Invoke() (*model.UpdateHostedDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostedDirectConnectResponse), nil
	}
}

type ListGdgwRouteTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGdgwRouteTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGdgwRouteTablesInvoker) Invoke() (*model.ListGdgwRouteTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGdgwRouteTablesResponse), nil
	}
}

type UpdateGdgwRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGdgwRouteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGdgwRouteTableInvoker) Invoke() (*model.UpdateGdgwRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGdgwRouteTableResponse), nil
	}
}

type CreateGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGlobalDcGatewayInvoker) Invoke() (*model.CreateGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalDcGatewayResponse), nil
	}
}

type CreatePeerLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePeerLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePeerLinkInvoker) Invoke() (*model.CreatePeerLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePeerLinkResponse), nil
	}
}

type DeleteGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGlobalDcGatewayInvoker) Invoke() (*model.DeleteGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalDcGatewayResponse), nil
	}
}

type DeletePeerLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePeerLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePeerLinkInvoker) Invoke() (*model.DeletePeerLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePeerLinkResponse), nil
	}
}

type ListGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalDcGatewayInvoker) Invoke() (*model.ListGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalDcGatewayResponse), nil
	}
}

type ListPeerLinksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPeerLinksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPeerLinksInvoker) Invoke() (*model.ListPeerLinksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPeerLinksResponse), nil
	}
}

type ListRmsGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRmsGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRmsGlobalDcGatewayInvoker) Invoke() (*model.ListRmsGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRmsGlobalDcGatewayResponse), nil
	}
}

type ShowGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalDcGatewayInvoker) Invoke() (*model.ShowGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalDcGatewayResponse), nil
	}
}

type ShowPeerLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPeerLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPeerLinkInvoker) Invoke() (*model.ShowPeerLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPeerLinkResponse), nil
	}
}

type ShowRmsGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRmsGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRmsGlobalDcGatewayInvoker) Invoke() (*model.ShowRmsGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRmsGlobalDcGatewayResponse), nil
	}
}

type UpdateGlobalDcGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalDcGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGlobalDcGatewayInvoker) Invoke() (*model.UpdateGlobalDcGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalDcGatewayResponse), nil
	}
}

type UpdatePeerLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePeerLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePeerLinkInvoker) Invoke() (*model.UpdatePeerLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePeerLinkResponse), nil
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

type ListTagResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagResourceInstancesInvoker) Invoke() (*model.ListTagResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagResourceInstancesResponse), nil
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

type CreateVirtualGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVirtualGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVirtualGatewayInvoker) Invoke() (*model.CreateVirtualGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVirtualGatewayResponse), nil
	}
}

type DeleteVirtualGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVirtualGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVirtualGatewayInvoker) Invoke() (*model.DeleteVirtualGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVirtualGatewayResponse), nil
	}
}

type ListVirtualGatewaysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVirtualGatewaysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVirtualGatewaysInvoker) Invoke() (*model.ListVirtualGatewaysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVirtualGatewaysResponse), nil
	}
}

type ShowVirtualGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVirtualGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVirtualGatewayInvoker) Invoke() (*model.ShowVirtualGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVirtualGatewayResponse), nil
	}
}

type UpdateVirtualGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVirtualGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVirtualGatewayInvoker) Invoke() (*model.UpdateVirtualGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVirtualGatewayResponse), nil
	}
}

type CreateVifPeerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVifPeerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVifPeerInvoker) Invoke() (*model.CreateVifPeerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVifPeerResponse), nil
	}
}

type CreateVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVirtualInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVirtualInterfaceInvoker) Invoke() (*model.CreateVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVirtualInterfaceResponse), nil
	}
}

type DeleteVifPeerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVifPeerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVifPeerInvoker) Invoke() (*model.DeleteVifPeerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVifPeerResponse), nil
	}
}

type DeleteVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVirtualInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVirtualInterfaceInvoker) Invoke() (*model.DeleteVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVirtualInterfaceResponse), nil
	}
}

type ListSwitchoverTestRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSwitchoverTestRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSwitchoverTestRecordsInvoker) Invoke() (*model.ListSwitchoverTestRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSwitchoverTestRecordsResponse), nil
	}
}

type ListVirtualInterfacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVirtualInterfacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVirtualInterfacesInvoker) Invoke() (*model.ListVirtualInterfacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVirtualInterfacesResponse), nil
	}
}

type ShowVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVirtualInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVirtualInterfaceInvoker) Invoke() (*model.ShowVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVirtualInterfaceResponse), nil
	}
}

type SwitchoverTestInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchoverTestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchoverTestInvoker) Invoke() (*model.SwitchoverTestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchoverTestResponse), nil
	}
}

type UpdateVifPeerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVifPeerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVifPeerInvoker) Invoke() (*model.UpdateVifPeerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVifPeerResponse), nil
	}
}

type UpdateVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVirtualInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVirtualInterfaceInvoker) Invoke() (*model.UpdateVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVirtualInterfaceResponse), nil
	}
}
