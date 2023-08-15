package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/ces/v1/model"
)

type BatchListMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListMetricDataInvoker) Invoke() (*model.BatchListMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListMetricDataResponse), nil
	}
}

type CreateAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlarmInvoker) Invoke() (*model.CreateAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmResponse), nil
	}
}

type CreateAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlarmTemplateInvoker) Invoke() (*model.CreateAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmTemplateResponse), nil
	}
}

type CreateEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventsInvoker) Invoke() (*model.CreateEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventsResponse), nil
	}
}

type CreateMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMetricDataInvoker) Invoke() (*model.CreateMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMetricDataResponse), nil
	}
}

type CreateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceGroupInvoker) Invoke() (*model.CreateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceGroupResponse), nil
	}
}

type DeleteAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmInvoker) Invoke() (*model.DeleteAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmResponse), nil
	}
}

type DeleteAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmTemplateInvoker) Invoke() (*model.DeleteAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmTemplateResponse), nil
	}
}

type DeleteResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceGroupInvoker) Invoke() (*model.DeleteResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceGroupResponse), nil
	}
}

type ListAlarmHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmHistoriesInvoker) Invoke() (*model.ListAlarmHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmHistoriesResponse), nil
	}
}

type ListAlarmTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmTemplatesInvoker) Invoke() (*model.ListAlarmTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTemplatesResponse), nil
	}
}

type ListAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmsInvoker) Invoke() (*model.ListAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmsResponse), nil
	}
}

type ListEventDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventDetailInvoker) Invoke() (*model.ListEventDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventDetailResponse), nil
	}
}

type ListEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventsInvoker) Invoke() (*model.ListEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventsResponse), nil
	}
}

type ListMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsInvoker) Invoke() (*model.ListMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsResponse), nil
	}
}

type ListResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceGroupInvoker) Invoke() (*model.ListResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceGroupResponse), nil
	}
}

type ShowAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmInvoker) Invoke() (*model.ShowAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmResponse), nil
	}
}

type ShowEventDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventDataInvoker) Invoke() (*model.ShowEventDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventDataResponse), nil
	}
}

type ShowMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricDataInvoker) Invoke() (*model.ShowMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricDataResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type ShowResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceGroupInvoker) Invoke() (*model.ShowResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceGroupResponse), nil
	}
}

type UpdateAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmInvoker) Invoke() (*model.UpdateAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmResponse), nil
	}
}

type UpdateAlarmActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmActionInvoker) Invoke() (*model.UpdateAlarmActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmActionResponse), nil
	}
}

type UpdateAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmTemplateInvoker) Invoke() (*model.UpdateAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmTemplateResponse), nil
	}
}

type UpdateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceGroupInvoker) Invoke() (*model.UpdateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceGroupResponse), nil
	}
}
