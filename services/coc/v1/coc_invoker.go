package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/coc/v1/model"
)

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type ListInstanceCompliantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceCompliantInvoker) Invoke() (*model.ListInstanceCompliantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceCompliantResponse), nil
	}
}

type ShowInstancePatchItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstancePatchItemsInvoker) Invoke() (*model.ShowInstancePatchItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstancePatchItemsResponse), nil
	}
}

type CreateReportCustomEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportCustomEventInvoker) Invoke() (*model.CreateReportCustomEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportCustomEventResponse), nil
	}
}

type CreateReportPrometheusEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportPrometheusEventInvoker) Invoke() (*model.CreateReportPrometheusEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportPrometheusEventResponse), nil
	}
}

type CreateCocIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCocIncidentInvoker) Invoke() (*model.CreateCocIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCocIncidentResponse), nil
	}
}

type HandleCocIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleCocIncidentInvoker) Invoke() (*model.HandleCocIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleCocIncidentResponse), nil
	}
}

type ListCocTicketOperationHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCocTicketOperationHistoriesInvoker) Invoke() (*model.ListCocTicketOperationHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCocTicketOperationHistoriesResponse), nil
	}
}

type ShowCocIncidentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCocIncidentDetailInvoker) Invoke() (*model.ShowCocIncidentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCocIncidentDetailResponse), nil
	}
}

type CreateCocIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCocIssuesInvoker) Invoke() (*model.CreateCocIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCocIssuesResponse), nil
	}
}

type ShowCocIssuesDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCocIssuesDetailInvoker) Invoke() (*model.ShowCocIssuesDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCocIssuesDetailResponse), nil
	}
}

type ListAuthorizableTicketsExternalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorizableTicketsExternalInvoker) Invoke() (*model.ListAuthorizableTicketsExternalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizableTicketsExternalResponse), nil
	}
}

type ListResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInvoker) Invoke() (*model.ListResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceResponse), nil
	}
}

type SyncResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncResourceInvoker) Invoke() (*model.SyncResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncResourceResponse), nil
	}
}

type GetScriptJobBatchInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScriptJobBatchInvoker) Invoke() (*model.GetScriptJobBatchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetScriptJobBatchResponse), nil
	}
}

type GetScriptJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScriptJobInfoInvoker) Invoke() (*model.GetScriptJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetScriptJobInfoResponse), nil
	}
}

type GetScriptJobStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScriptJobStatisticsInvoker) Invoke() (*model.GetScriptJobStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetScriptJobStatisticsResponse), nil
	}
}

type ListScriptJobBatchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptJobBatchesInvoker) Invoke() (*model.ListScriptJobBatchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptJobBatchesResponse), nil
	}
}

type ListScriptJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptJobsInvoker) Invoke() (*model.ListScriptJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptJobsResponse), nil
	}
}

type OperateScriptJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *OperateScriptJobInvoker) Invoke() (*model.OperateScriptJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperateScriptJobResponse), nil
	}
}

type CreateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScriptInvoker) Invoke() (*model.CreateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScriptResponse), nil
	}
}

type DeleteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScriptInvoker) Invoke() (*model.DeleteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScriptResponse), nil
	}
}

type ExecuteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteScriptInvoker) Invoke() (*model.ExecuteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteScriptResponse), nil
	}
}

type GetScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScriptInvoker) Invoke() (*model.GetScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetScriptResponse), nil
	}
}

type ListScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptsInvoker) Invoke() (*model.ListScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptsResponse), nil
	}
}

type UpdateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScriptInvoker) Invoke() (*model.UpdateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScriptResponse), nil
	}
}

type ExecutePublicScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecutePublicScriptInvoker) Invoke() (*model.ExecutePublicScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecutePublicScriptResponse), nil
	}
}

type GetPublicScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetPublicScriptInvoker) Invoke() (*model.GetPublicScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPublicScriptResponse), nil
	}
}

type ListPublicScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicScriptsInvoker) Invoke() (*model.ListPublicScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicScriptsResponse), nil
	}
}

type CreateWarRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWarRoomInvoker) Invoke() (*model.CreateWarRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWarRoomResponse), nil
	}
}

type ListWarRoomsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWarRoomsInvoker) Invoke() (*model.ListWarRoomsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWarRoomsResponse), nil
	}
}
