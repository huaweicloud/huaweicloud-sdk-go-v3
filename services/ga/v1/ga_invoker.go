package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ga/v1/model"
)

type CreateAcceleratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAcceleratorInvoker) Invoke() (*model.CreateAcceleratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAcceleratorResponse), nil
	}
}

type DeleteAcceleratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAcceleratorInvoker) Invoke() (*model.DeleteAcceleratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAcceleratorResponse), nil
	}
}

type ListAcceleratorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAcceleratorsInvoker) Invoke() (*model.ListAcceleratorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAcceleratorsResponse), nil
	}
}

type ShowAcceleratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAcceleratorInvoker) Invoke() (*model.ShowAcceleratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAcceleratorResponse), nil
	}
}

type UpdateAcceleratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAcceleratorInvoker) Invoke() (*model.UpdateAcceleratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAcceleratorResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ShowEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndpointInvoker) Invoke() (*model.ShowEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndpointResponse), nil
	}
}

type UpdateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointInvoker) Invoke() (*model.UpdateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointResponse), nil
	}
}

type CreateEndpointGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointGroupInvoker) Invoke() (*model.CreateEndpointGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointGroupResponse), nil
	}
}

type DeleteEndpointGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointGroupInvoker) Invoke() (*model.DeleteEndpointGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointGroupResponse), nil
	}
}

type ListEndpointGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointGroupsInvoker) Invoke() (*model.ListEndpointGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointGroupsResponse), nil
	}
}

type ShowEndpointGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndpointGroupInvoker) Invoke() (*model.ShowEndpointGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndpointGroupResponse), nil
	}
}

type UpdateEndpointGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointGroupInvoker) Invoke() (*model.UpdateEndpointGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointGroupResponse), nil
	}
}

type CreateHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHealthCheckInvoker) Invoke() (*model.CreateHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHealthCheckResponse), nil
	}
}

type DeleteHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHealthCheckInvoker) Invoke() (*model.DeleteHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHealthCheckResponse), nil
	}
}

type ListHealthChecksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHealthChecksInvoker) Invoke() (*model.ListHealthChecksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHealthChecksResponse), nil
	}
}

type ShowHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHealthCheckInvoker) Invoke() (*model.ShowHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthCheckResponse), nil
	}
}

type UpdateHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHealthCheckInvoker) Invoke() (*model.UpdateHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHealthCheckResponse), nil
	}
}

type CreateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateListenerInvoker) Invoke() (*model.CreateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateListenerResponse), nil
	}
}

type DeleteListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteListenerInvoker) Invoke() (*model.DeleteListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteListenerResponse), nil
	}
}

type ListListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenersInvoker) Invoke() (*model.ListListenersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenersResponse), nil
	}
}

type ShowListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListenerInvoker) Invoke() (*model.ShowListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerResponse), nil
	}
}

type UpdateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateListenerInvoker) Invoke() (*model.UpdateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateListenerResponse), nil
	}
}

type ListRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegionsInvoker) Invoke() (*model.ListRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegionsResponse), nil
	}
}
