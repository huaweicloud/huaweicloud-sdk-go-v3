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

type EnableDefensePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDefensePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDefensePolicyInvoker) Invoke() (*model.EnableDefensePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDefensePolicyResponse), nil
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

type ListQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaInvoker) Invoke() (*model.ListQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaResponse), nil
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

type ShowLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLogConfigInvoker) Invoke() (*model.ShowLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogConfigResponse), nil
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

type UpdateLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLogConfigInvoker) Invoke() (*model.UpdateLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogConfigResponse), nil
	}
}
