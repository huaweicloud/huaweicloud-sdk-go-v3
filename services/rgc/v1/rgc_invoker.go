package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type DisableControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableControlInvoker) Invoke() (*model.DisableControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableControlResponse), nil
	}
}

type EnableControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableControlInvoker) Invoke() (*model.EnableControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableControlResponse), nil
	}
}

type ListControlsForOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlsForOrganizationalUnitInvoker) Invoke() (*model.ListControlsForOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsForOrganizationalUnitResponse), nil
	}
}

type ShowControlOperateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlOperateInvoker) Invoke() (*model.ShowControlOperateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlOperateResponse), nil
	}
}
