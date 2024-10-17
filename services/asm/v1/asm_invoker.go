package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/asm/v1/model"
)

type CreateMeshInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMeshInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMeshInvoker) Invoke() (*model.CreateMeshResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMeshResponse), nil
	}
}

type DeleteMeshInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMeshInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMeshInvoker) Invoke() (*model.DeleteMeshResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMeshResponse), nil
	}
}

type ListMeshesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMeshesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMeshesInvoker) Invoke() (*model.ListMeshesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMeshesResponse), nil
	}
}

type ShowMeshInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMeshInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMeshInvoker) Invoke() (*model.ShowMeshResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMeshResponse), nil
	}
}
