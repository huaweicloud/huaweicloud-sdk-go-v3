package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabricep/v1/model"
)

type DeleteServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceInstanceInvoker) Invoke() (*model.DeleteServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceInstanceResponse), nil
	}
}

type DeployServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeployServiceInstanceInvoker) Invoke() (*model.DeployServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployServiceInstanceResponse), nil
	}
}

type ListServicesInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicesInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicesInstancesInvoker) Invoke() (*model.ListServicesInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicesInstancesResponse), nil
	}
}

type ShowServiceInstanceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServiceInstanceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServiceInstanceDetailInvoker) Invoke() (*model.ShowServiceInstanceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServiceInstanceDetailResponse), nil
	}
}

type UpdateServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServiceInstanceInvoker) Invoke() (*model.UpdateServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceInstanceResponse), nil
	}
}
