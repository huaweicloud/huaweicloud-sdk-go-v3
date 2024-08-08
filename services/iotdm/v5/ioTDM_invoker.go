package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotdm/v5/model"
)

type BindInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindInstanceTagsInvoker) Invoke() (*model.BindInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindInstanceTagsResponse), nil
	}
}

type ChangeInstanceChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstanceChargeModeInvoker) Invoke() (*model.ChangeInstanceChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceChargeModeResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type UnbindInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnbindInstanceTagsInvoker) Invoke() (*model.UnbindInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnbindInstanceTagsResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}
