package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

type CreateCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCloudConnectionInvoker) Invoke() (*model.CreateCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCloudConnectionResponse), nil
	}
}

type CreateNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNetworkInstanceInvoker) Invoke() (*model.CreateNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNetworkInstanceResponse), nil
	}
}

type DeleteCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudConnectionInvoker) Invoke() (*model.DeleteCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudConnectionResponse), nil
	}
}

type DeleteNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNetworkInstanceInvoker) Invoke() (*model.DeleteNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNetworkInstanceResponse), nil
	}
}

type ListCloudConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionRoutesInvoker) Invoke() (*model.ListCloudConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionRoutesResponse), nil
	}
}

type ListCloudConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionsInvoker) Invoke() (*model.ListCloudConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionsResponse), nil
	}
}

type ListNetworkInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkInstancesInvoker) Invoke() (*model.ListNetworkInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkInstancesResponse), nil
	}
}

type ShowCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudConnectionInvoker) Invoke() (*model.ShowCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudConnectionResponse), nil
	}
}

type ShowCloudConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudConnectionRoutesInvoker) Invoke() (*model.ShowCloudConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudConnectionRoutesResponse), nil
	}
}

type ShowNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkInstanceInvoker) Invoke() (*model.ShowNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkInstanceResponse), nil
	}
}

type UpdateCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCloudConnectionInvoker) Invoke() (*model.UpdateCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCloudConnectionResponse), nil
	}
}

type UpdateNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNetworkInstanceInvoker) Invoke() (*model.UpdateNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNetworkInstanceResponse), nil
	}
}
