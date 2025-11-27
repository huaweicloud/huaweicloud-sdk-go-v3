package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/coc/v1/model"
)

type CreatePasswordChangePlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePasswordChangePlanInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePasswordChangePlanInvoker) Invoke() (*model.CreatePasswordChangePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePasswordChangePlanResponse), nil
	}
}

type ResetAccountPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetAccountPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetAccountPasswordInvoker) Invoke() (*model.ResetAccountPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetAccountPasswordResponse), nil
	}
}

type UpdateAccountPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccountPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccountPasswordInvoker) Invoke() (*model.UpdateAccountPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccountPasswordResponse), nil
	}
}

type ClearAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *ClearAlarmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClearAlarmInvoker) Invoke() (*model.ClearAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearAlarmResponse), nil
	}
}

type HandlerAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandlerAlarmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandlerAlarmInvoker) Invoke() (*model.HandlerAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandlerAlarmResponse), nil
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

type TransferAlarmToIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *TransferAlarmToIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TransferAlarmToIncidentInvoker) Invoke() (*model.TransferAlarmToIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TransferAlarmToIncidentResponse), nil
	}
}

type CreateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationInvoker) Invoke() (*model.CreateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type UpdateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInvoker) Invoke() (*model.UpdateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationResponse), nil
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

type ListApplicationViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationViewInvoker) Invoke() (*model.ListApplicationViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationViewResponse), nil
	}
}

type CreateAssessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssessTaskInvoker) Invoke() (*model.CreateAssessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssessTaskResponse), nil
	}
}

type ListAssessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssessTaskInvoker) Invoke() (*model.ListAssessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssessTaskResponse), nil
	}
}

type UpdateChangeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateChangeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateChangeInvoker) Invoke() (*model.UpdateChangeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateChangeResponse), nil
	}
}

type HandleIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleIncidentInvoker) Invoke() (*model.HandleIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleIncidentResponse), nil
	}
}

type ListIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIncidentsInvoker) Invoke() (*model.ListIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncidentsResponse), nil
	}
}

type ListIncidentsHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncidentsHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIncidentsHistoriesInvoker) Invoke() (*model.ListIncidentsHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncidentsHistoriesResponse), nil
	}
}

type ShowIncidentTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIncidentTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIncidentTaskInvoker) Invoke() (*model.ShowIncidentTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIncidentTaskResponse), nil
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

type CreateApplicationComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationComponentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationComponentsInvoker) Invoke() (*model.CreateApplicationComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationComponentsResponse), nil
	}
}

type DeleteApplicationComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationComponentInvoker) Invoke() (*model.DeleteApplicationComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationComponentResponse), nil
	}
}

type ListApplicationComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationComponentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationComponentsInvoker) Invoke() (*model.ListApplicationComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationComponentsResponse), nil
	}
}

type UpdateApplicationComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationComponentInvoker) Invoke() (*model.UpdateApplicationComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationComponentResponse), nil
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

type CountOtherResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountOtherResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountOtherResourceInvoker) Invoke() (*model.CountOtherResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountOtherResourceResponse), nil
	}
}

type ImportOtherResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportOtherResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportOtherResourceInvoker) Invoke() (*model.ImportOtherResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportOtherResourceResponse), nil
	}
}

type CancelDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelDiagnosisTaskInvoker) Invoke() (*model.CancelDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelDiagnosisTaskResponse), nil
	}
}

type CreateDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDiagnosisTaskInvoker) Invoke() (*model.CreateDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnosisTaskResponse), nil
	}
}

type ListDiagnosisTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnosisTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnosisTasksInvoker) Invoke() (*model.ListDiagnosisTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnosisTasksResponse), nil
	}
}

type RetryDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryDiagnosisTaskInvoker) Invoke() (*model.RetryDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryDiagnosisTaskResponse), nil
	}
}

type ShowDiagnosisNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisNodeInvoker) Invoke() (*model.ShowDiagnosisNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisNodeResponse), nil
	}
}

type ShowDiagnosisSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisSummaryInvoker) Invoke() (*model.ShowDiagnosisSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisSummaryResponse), nil
	}
}

type ShowDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisTaskInvoker) Invoke() (*model.ShowDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisTaskResponse), nil
	}
}

type CreateDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDocumentInvoker) Invoke() (*model.CreateDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocumentResponse), nil
	}
}

type DeleteDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDocumentInvoker) Invoke() (*model.DeleteDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDocumentResponse), nil
	}
}

type ExecuteDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteDocumentInvoker) Invoke() (*model.ExecuteDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDocumentResponse), nil
	}
}

type GetDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDocumentInvoker) Invoke() (*model.GetDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDocumentResponse), nil
	}
}

type GetDocumentAtomicInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDocumentAtomicInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDocumentAtomicInfoInvoker) Invoke() (*model.GetDocumentAtomicInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDocumentAtomicInfoResponse), nil
	}
}

type ListDocumentAtomicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDocumentAtomicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDocumentAtomicsInvoker) Invoke() (*model.ListDocumentAtomicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDocumentAtomicsResponse), nil
	}
}

type ListDocumentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDocumentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDocumentsInvoker) Invoke() (*model.ListDocumentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDocumentsResponse), nil
	}
}

type UpdateDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDocumentInvoker) Invoke() (*model.UpdateDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDocumentResponse), nil
	}
}

type ListEnterpriseProjectCollectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseProjectCollectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseProjectCollectInvoker) Invoke() (*model.ListEnterpriseProjectCollectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseProjectCollectResponse), nil
	}
}

type UpdateEnterpriseProjectCollectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnterpriseProjectCollectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnterpriseProjectCollectInvoker) Invoke() (*model.UpdateEnterpriseProjectCollectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnterpriseProjectCollectResponse), nil
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

type GetExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetExecutionInvoker) Invoke() (*model.GetExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetExecutionResponse), nil
	}
}

type ListExecutionInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExecutionInstancesInvoker) Invoke() (*model.ListExecutionInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionInstancesResponse), nil
	}
}

type ListExecutionStepsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionStepsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExecutionStepsInvoker) Invoke() (*model.ListExecutionStepsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionStepsResponse), nil
	}
}

type ListExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExecutionsInvoker) Invoke() (*model.ListExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionsResponse), nil
	}
}

type OperateExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *OperateExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *OperateExecutionInvoker) Invoke() (*model.OperateExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperateExecutionResponse), nil
	}
}

type ListSubTicketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubTicketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubTicketsInvoker) Invoke() (*model.ListSubTicketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubTicketsResponse), nil
	}
}

type ExecuteTicketActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTicketActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteTicketActionInvoker) Invoke() (*model.ExecuteTicketActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTicketActionResponse), nil
	}
}

type ListTicketOperationHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTicketOperationHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTicketOperationHistoriesInvoker) Invoke() (*model.ListTicketOperationHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTicketOperationHistoriesResponse), nil
	}
}

type ListTicketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTicketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTicketsInvoker) Invoke() (*model.ListTicketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTicketsResponse), nil
	}
}

type ShowTicketInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTicketInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTicketInfoInvoker) Invoke() (*model.ShowTicketInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTicketInfoResponse), nil
	}
}

type DeleteTicketInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTicketInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTicketInfoInvoker) Invoke() (*model.DeleteTicketInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTicketInfoResponse), nil
	}
}

type UpdateTicketInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTicketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTicketInvoker) Invoke() (*model.UpdateTicketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTicketResponse), nil
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

type CreateExternalCocAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExternalCocAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExternalCocAttachmentInvoker) Invoke() (*model.CreateExternalCocAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExternalCocAttachmentResponse), nil
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

type ListIncidentSimpleTicketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncidentSimpleTicketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIncidentSimpleTicketsInvoker) Invoke() (*model.ListIncidentSimpleTicketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncidentSimpleTicketsResponse), nil
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

type CreateAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAttachmentInvoker) Invoke() (*model.CreateAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAttachmentResponse), nil
	}
}

type CreateTicketInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTicketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTicketInvoker) Invoke() (*model.CreateTicketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTicketResponse), nil
	}
}

type DownloadAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAttachmentInvoker) Invoke() (*model.DownloadAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAttachmentResponse), nil
	}
}

type CreateApplicationGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationGroupInvoker) Invoke() (*model.CreateApplicationGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationGroupResponse), nil
	}
}

type DeleteApplicationGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationGroupInvoker) Invoke() (*model.DeleteApplicationGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationGroupResponse), nil
	}
}

type ListApplicationGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationGroupsInvoker) Invoke() (*model.ListApplicationGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationGroupsResponse), nil
	}
}

type SyncGroupResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncGroupResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncGroupResourceInvoker) Invoke() (*model.SyncGroupResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncGroupResourceResponse), nil
	}
}

type UpdateApplicationGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationGroupInvoker) Invoke() (*model.UpdateApplicationGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationGroupResponse), nil
	}
}

type ListGroupAliResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupAliResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupAliResourceRelationsInvoker) Invoke() (*model.ListGroupAliResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupAliResourceRelationsResponse), nil
	}
}

type ListGroupOtherResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupOtherResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupOtherResourceRelationsInvoker) Invoke() (*model.ListGroupOtherResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupOtherResourceRelationsResponse), nil
	}
}

type CountGroupRmsResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountGroupRmsResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountGroupRmsResourceRelationsInvoker) Invoke() (*model.CountGroupRmsResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountGroupRmsResourceRelationsResponse), nil
	}
}

type CreateGroupRmsResourceRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupRmsResourceRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupRmsResourceRelationInvoker) Invoke() (*model.CreateGroupRmsResourceRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupRmsResourceRelationResponse), nil
	}
}

type DeleteGroupRmsResourceRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupRmsResourceRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupRmsResourceRelationInvoker) Invoke() (*model.DeleteGroupRmsResourceRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupRmsResourceRelationResponse), nil
	}
}

type ListCapacityOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCapacityOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCapacityOrderInvoker) Invoke() (*model.ListCapacityOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCapacityOrderResponse), nil
	}
}

type ListCapacityViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCapacityViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCapacityViewInvoker) Invoke() (*model.ListCapacityViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCapacityViewResponse), nil
	}
}

type ListGroupRmsResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupRmsResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupRmsResourceRelationsInvoker) Invoke() (*model.ListGroupRmsResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupRmsResourceRelationsResponse), nil
	}
}

type UpdateGroupRmsResourceRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupRmsResourceRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupRmsResourceRelationInvoker) Invoke() (*model.UpdateGroupRmsResourceRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupRmsResourceRelationResponse), nil
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

type CountMultiCloudResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountMultiCloudResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountMultiCloudResourcesInvoker) Invoke() (*model.CountMultiCloudResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountMultiCloudResourcesResponse), nil
	}
}

type SyncMultiCloudResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncMultiCloudResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncMultiCloudResourceInvoker) Invoke() (*model.SyncMultiCloudResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncMultiCloudResourceResponse), nil
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

type CountResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountResourcesInvoker) Invoke() (*model.CountResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourcesResponse), nil
	}
}

type ListResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
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

type SyncResourceAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncResourceAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncResourceAgentInvoker) Invoke() (*model.SyncResourceAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncResourceAgentResponse), nil
	}
}

type ListScriptResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScriptResourceTagsInvoker) Invoke() (*model.ListScriptResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptResourceTagsResponse), nil
	}
}

type UpdateResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceTagsInvoker) Invoke() (*model.UpdateResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceTagsResponse), nil
	}
}

type CreateResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceTagsInvoker) Invoke() (*model.CreateResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceTagsResponse), nil
	}
}

type ListResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type CountResourcesOfResourceViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourcesOfResourceViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountResourcesOfResourceViewInvoker) Invoke() (*model.CountResourcesOfResourceViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourcesOfResourceViewResponse), nil
	}
}

type CreateResourceViewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceViewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceViewsInvoker) Invoke() (*model.CreateResourceViewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceViewsResponse), nil
	}
}

type DeleteResourceViewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceViewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceViewsInvoker) Invoke() (*model.DeleteResourceViewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceViewsResponse), nil
	}
}

type ListResourceViewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceViewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceViewsInvoker) Invoke() (*model.ListResourceViewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceViewsResponse), nil
	}
}

type ListResourcesOfResourceViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesOfResourceViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesOfResourceViewInvoker) Invoke() (*model.ListResourcesOfResourceViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesOfResourceViewResponse), nil
	}
}

type SyncResourcesOfResourceViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncResourcesOfResourceViewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncResourcesOfResourceViewInvoker) Invoke() (*model.SyncResourcesOfResourceViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncResourcesOfResourceViewResponse), nil
	}
}

type UpdateResourceViewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceViewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceViewsInvoker) Invoke() (*model.UpdateResourceViewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceViewsResponse), nil
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

type AcceptScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptScriptInvoker) Invoke() (*model.AcceptScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptScriptResponse), nil
	}
}

type CheckScriptRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckScriptRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckScriptRiskInvoker) Invoke() (*model.CheckScriptRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckScriptRiskResponse), nil
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

type ListInstancesBatchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesBatchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesBatchInvoker) Invoke() (*model.ListInstancesBatchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesBatchResponse), nil
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

type CreateVendorAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVendorAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVendorAccountInvoker) Invoke() (*model.CreateVendorAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVendorAccountResponse), nil
	}
}

type DeleteVendorAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVendorAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVendorAccountInvoker) Invoke() (*model.DeleteVendorAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVendorAccountResponse), nil
	}
}

type ListVendorAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVendorAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVendorAccountInvoker) Invoke() (*model.ListVendorAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVendorAccountResponse), nil
	}
}

type UpdateVendorAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVendorAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVendorAccountInvoker) Invoke() (*model.UpdateVendorAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVendorAccountResponse), nil
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
