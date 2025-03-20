package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/coc/v1/model"
)

type ShowAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccountInvoker) Invoke() (*model.ShowAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccountResponse), nil
	}
}

type ListAlarmHandleHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmHandleHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmHandleHistoriesInvoker) Invoke() (*model.ListAlarmHandleHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmHandleHistoriesResponse), nil
	}
}

type ShowAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmInvoker) Invoke() (*model.ShowAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type ListApplicationModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationModelInvoker) Invoke() (*model.ListApplicationModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationModelResponse), nil
	}
}

type BatchCreateApplicationViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateApplicationViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateApplicationViewInvoker) Invoke() (*model.BatchCreateApplicationViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateApplicationViewResponse), nil
	}
}

type ShowPatchBaselineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPatchBaselineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPatchBaselineInvoker) Invoke() (*model.ShowPatchBaselineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPatchBaselineResponse), nil
	}
}

type ListCceCompliantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCceCompliantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCceCompliantInvoker) Invoke() (*model.ListCceCompliantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCceCompliantResponse), nil
	}
}

type ListInstanceCompliantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceCompliantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowInstancePatchItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateReportCustomEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateReportPrometheusEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateCocIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *HandleCocIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCocTicketOperationHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCocIncidentDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateCocIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCocIssuesDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuthorizableTicketsExternalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuthorizableTicketsExternalInvoker) Invoke() (*model.ListAuthorizableTicketsExternalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizableTicketsExternalResponse), nil
	}
}

type ListMultiCloudResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMultiCloudResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMultiCloudResourcesInvoker) Invoke() (*model.ListMultiCloudResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMultiCloudResourcesResponse), nil
	}
}

type ListPersonnelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPersonnelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPersonnelInvoker) Invoke() (*model.ListPersonnelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPersonnelResponse), nil
	}
}

type SyncAddPersonnelInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncAddPersonnelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncAddPersonnelInvoker) Invoke() (*model.SyncAddPersonnelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncAddPersonnelResponse), nil
	}
}

type CountMultiResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountMultiResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountMultiResourcesInvoker) Invoke() (*model.CountMultiResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountMultiResourcesResponse), nil
	}
}

type ListResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SyncResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncResourceInvoker) Invoke() (*model.SyncResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncResourceResponse), nil
	}
}

type CreateScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScheduledTaskInvoker) Invoke() (*model.CreateScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduledTaskResponse), nil
	}
}

type DeleteScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduledTaskInvoker) Invoke() (*model.DeleteScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduledTaskResponse), nil
	}
}

type DisableScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableScheduledTaskInvoker) Invoke() (*model.DisableScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableScheduledTaskResponse), nil
	}
}

type EnableScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableScheduledTaskInvoker) Invoke() (*model.EnableScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableScheduledTaskResponse), nil
	}
}

type ListScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTaskInvoker) Invoke() (*model.ListScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTaskResponse), nil
	}
}

type ListScheduledTaskHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTaskHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTaskHistoryInvoker) Invoke() (*model.ListScheduledTaskHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTaskHistoryResponse), nil
	}
}

type ShowScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScheduledTaskInvoker) Invoke() (*model.ShowScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduledTaskResponse), nil
	}
}

type UpdateScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScheduledTaskInvoker) Invoke() (*model.UpdateScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduledTaskResponse), nil
	}
}

type GetScriptJobBatchInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScriptJobBatchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *GetScriptJobInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *GetScriptJobStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListScriptJobBatchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListScriptJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *OperateScriptJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExecuteScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *GetScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExecutePublicScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *GetPublicScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPublicScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublicScriptsInvoker) Invoke() (*model.ListPublicScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicScriptsResponse), nil
	}
}

type ShowSlaCustomizedTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlaCustomizedTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlaCustomizedTemplateInvoker) Invoke() (*model.ShowSlaCustomizedTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlaCustomizedTemplateResponse), nil
	}
}

type ShowSlaOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlaOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlaOrderInvoker) Invoke() (*model.ShowSlaOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlaOrderResponse), nil
	}
}

type ListInterruptRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInterruptRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInterruptRecordsInvoker) Invoke() (*model.ListInterruptRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInterruptRecordsResponse), nil
	}
}

type ShowSloDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSloDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSloDetailInvoker) Invoke() (*model.ShowSloDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSloDetailResponse), nil
	}
}

type CreateWarRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWarRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListWarRoomsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWarRoomsInvoker) Invoke() (*model.ListWarRoomsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWarRoomsResponse), nil
	}
}
