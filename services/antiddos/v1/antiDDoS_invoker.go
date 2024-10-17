package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/antiddos/v1/model"
)

type CreateDefaultConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDefaultConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDefaultConfigInvoker) Invoke() (*model.CreateDefaultConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDefaultConfigResponse), nil
	}
}

type DeleteDefaultConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDefaultConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDefaultConfigInvoker) Invoke() (*model.DeleteDefaultConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDefaultConfigResponse), nil
	}
}

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

type ShowDefaultConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDefaultConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDefaultConfigInvoker) Invoke() (*model.ShowDefaultConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDefaultConfigResponse), nil
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

type ListDailyLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDailyLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDailyLogInvoker) Invoke() (*model.ListDailyLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDailyLogResponse), nil
	}
}

type ListDailyReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDailyReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDailyReportInvoker) Invoke() (*model.ListDailyReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDailyReportResponse), nil
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

type ListWeeklyReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWeeklyReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWeeklyReportsInvoker) Invoke() (*model.ListWeeklyReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWeeklyReportsResponse), nil
	}
}

type ShowDDosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDDosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDDosInvoker) Invoke() (*model.ShowDDosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDDosResponse), nil
	}
}

type ShowDDosStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDDosStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDDosStatusInvoker) Invoke() (*model.ShowDDosStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDDosStatusResponse), nil
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

type UpdateDDosInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDDosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDDosInvoker) Invoke() (*model.UpdateDDosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDDosResponse), nil
	}
}
