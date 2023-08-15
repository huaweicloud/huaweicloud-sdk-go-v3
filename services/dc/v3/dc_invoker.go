package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dc/v3/model"
)

type CreateHostedDirectConnectInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateHostedDirectConnectInvoker) Invoke() (*model.UpdateHostedDirectConnectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostedDirectConnectResponse), nil
	}
}

type BatchCreateResourceTagsInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateVirtualGatewayInvoker) Invoke() (*model.UpdateVirtualGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVirtualGatewayResponse), nil
	}
}

type CreateVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVirtualInterfaceInvoker) Invoke() (*model.CreateVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVirtualInterfaceResponse), nil
	}
}

type DeleteVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVirtualInterfaceInvoker) Invoke() (*model.DeleteVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVirtualInterfaceResponse), nil
	}
}

type ListVirtualInterfacesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowVirtualInterfaceInvoker) Invoke() (*model.ShowVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVirtualInterfaceResponse), nil
	}
}

type UpdateVirtualInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVirtualInterfaceInvoker) Invoke() (*model.UpdateVirtualInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVirtualInterfaceResponse), nil
	}
}
