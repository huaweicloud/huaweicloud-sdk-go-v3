package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/antiddos/v2/model"
)

type ShowAlertConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertConfigInvoker) Invoke() (*model.ShowAlertConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertConfigResponse), nil
	}
}

type UpdateAlertConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlertConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlertConfigInvoker) Invoke() (*model.UpdateAlertConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlertConfigResponse), nil
	}
}

type ListDDosStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDDosStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDDosStatusInvoker) Invoke() (*model.ListDDosStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDDosStatusResponse), nil
	}
}

type ListNewConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNewConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNewConfigsInvoker) Invoke() (*model.ListNewConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNewConfigsResponse), nil
	}
}

type ShowNewTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNewTaskStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNewTaskStatusInvoker) Invoke() (*model.ShowNewTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNewTaskStatusResponse), nil
	}
}
