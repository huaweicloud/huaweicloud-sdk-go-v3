package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/secmaster/v2/model"
)

type BatchSearchMetricHitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSearchMetricHitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSearchMetricHitsInvoker) Invoke() (*model.BatchSearchMetricHitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSearchMetricHitsResponse), nil
	}
}

type ChangeAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAlertInvoker) Invoke() (*model.ChangeAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAlertResponse), nil
	}
}

type ChangeIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIncidentInvoker) Invoke() (*model.ChangeIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIncidentResponse), nil
	}
}

type ChangePlaybookInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePlaybookInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePlaybookInstanceInvoker) Invoke() (*model.ChangePlaybookInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePlaybookInstanceResponse), nil
	}
}

type CopyPlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyPlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyPlaybookVersionInvoker) Invoke() (*model.CopyPlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyPlaybookVersionResponse), nil
	}
}

type CreateAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertInvoker) Invoke() (*model.CreateAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertResponse), nil
	}
}

type CreateAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertRuleInvoker) Invoke() (*model.CreateAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertRuleResponse), nil
	}
}

type CreateAlertRuleSimulationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertRuleSimulationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertRuleSimulationInvoker) Invoke() (*model.CreateAlertRuleSimulationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertRuleSimulationResponse), nil
	}
}

type CreateBatchOrderAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchOrderAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBatchOrderAlertsInvoker) Invoke() (*model.CreateBatchOrderAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

type CreateDataobjectRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataobjectRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataobjectRelationsInvoker) Invoke() (*model.CreateDataobjectRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataobjectRelationsResponse), nil
	}
}

type CreateDataspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataspaceInvoker) Invoke() (*model.CreateDataspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataspaceResponse), nil
	}
}

type CreateIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIncidentInvoker) Invoke() (*model.CreateIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIncidentResponse), nil
	}
}

type CreateIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIndicatorInvoker) Invoke() (*model.CreateIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIndicatorResponse), nil
	}
}

type CreatePipeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePipeInvoker) Invoke() (*model.CreatePipeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipeResponse), nil
	}
}

type CreatePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookInvoker) Invoke() (*model.CreatePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookResponse), nil
	}
}

type CreatePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookActionInvoker) Invoke() (*model.CreatePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookActionResponse), nil
	}
}

type CreatePlaybookApproveInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookApproveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookApproveInvoker) Invoke() (*model.CreatePlaybookApproveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookApproveResponse), nil
	}
}

type CreatePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookRuleInvoker) Invoke() (*model.CreatePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookRuleResponse), nil
	}
}

type CreatePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookVersionInvoker) Invoke() (*model.CreatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookVersionResponse), nil
	}
}

type CreatePostPaidOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePostPaidOrderInvoker) Invoke() (*model.CreatePostPaidOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidOrderResponse), nil
	}
}

type CreateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkspaceInvoker) Invoke() (*model.CreateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkspaceResponse), nil
	}
}

type DeleteAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlertInvoker) Invoke() (*model.DeleteAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertResponse), nil
	}
}

type DeleteAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlertRuleInvoker) Invoke() (*model.DeleteAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertRuleResponse), nil
	}
}

type DeleteDataobjectRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataobjectRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataobjectRelationsInvoker) Invoke() (*model.DeleteDataobjectRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataobjectRelationsResponse), nil
	}
}

type DeleteIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIncidentInvoker) Invoke() (*model.DeleteIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIncidentResponse), nil
	}
}

type DeleteIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIndicatorInvoker) Invoke() (*model.DeleteIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIndicatorResponse), nil
	}
}

type DeletePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookInvoker) Invoke() (*model.DeletePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookResponse), nil
	}
}

type DeletePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookActionInvoker) Invoke() (*model.DeletePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookActionResponse), nil
	}
}

type DeletePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookRuleInvoker) Invoke() (*model.DeletePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookRuleResponse), nil
	}
}

type DeletePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookVersionInvoker) Invoke() (*model.DeletePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookVersionResponse), nil
	}
}

type DisableAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableAlertRuleInvoker) Invoke() (*model.DisableAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableAlertRuleResponse), nil
	}
}

type EnableAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableAlertRuleInvoker) Invoke() (*model.EnableAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableAlertRuleResponse), nil
	}
}

type ListAlertRuleMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRuleMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRuleMetricsInvoker) Invoke() (*model.ListAlertRuleMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRuleMetricsResponse), nil
	}
}

type ListAlertRuleTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRuleTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRuleTemplatesInvoker) Invoke() (*model.ListAlertRuleTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

type ListAlertRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRulesInvoker) Invoke() (*model.ListAlertRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRulesResponse), nil
	}
}

type ListAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertsInvoker) Invoke() (*model.ListAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertsResponse), nil
	}
}

type ListDataclassInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataclassInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataclassInvoker) Invoke() (*model.ListDataclassResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataclassResponse), nil
	}
}

type ListDataclassFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataclassFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataclassFieldsInvoker) Invoke() (*model.ListDataclassFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataclassFieldsResponse), nil
	}
}

type ListDataobjectRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataobjectRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataobjectRelationsInvoker) Invoke() (*model.ListDataobjectRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataobjectRelationsResponse), nil
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

type ListIndicatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIndicatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIndicatorsInvoker) Invoke() (*model.ListIndicatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIndicatorsResponse), nil
	}
}

type ListPlaybookActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookActionsInvoker) Invoke() (*model.ListPlaybookActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookActionsResponse), nil
	}
}

type ListPlaybookApprovesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookApprovesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookApprovesInvoker) Invoke() (*model.ListPlaybookApprovesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookApprovesResponse), nil
	}
}

type ListPlaybookAuditLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookAuditLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookAuditLogsInvoker) Invoke() (*model.ListPlaybookAuditLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookAuditLogsResponse), nil
	}
}

type ListPlaybookInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookInstancesInvoker) Invoke() (*model.ListPlaybookInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookInstancesResponse), nil
	}
}

type ListPlaybookVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookVersionsInvoker) Invoke() (*model.ListPlaybookVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookVersionsResponse), nil
	}
}

type ListPlaybooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybooksInvoker) Invoke() (*model.ListPlaybooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybooksResponse), nil
	}
}

type ListWorkflowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowsInvoker) Invoke() (*model.ListWorkflowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowsResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type SearchBaselineInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchBaselineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchBaselineInvoker) Invoke() (*model.SearchBaselineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchBaselineResponse), nil
	}
}

type ShowAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertInvoker) Invoke() (*model.ShowAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertResponse), nil
	}
}

type ShowAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertRuleInvoker) Invoke() (*model.ShowAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertRuleResponse), nil
	}
}

type ShowAlertRuleTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertRuleTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertRuleTemplateInvoker) Invoke() (*model.ShowAlertRuleTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

type ShowIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIncidentInvoker) Invoke() (*model.ShowIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIncidentResponse), nil
	}
}

type ShowIndicatorDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIndicatorDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIndicatorDetailInvoker) Invoke() (*model.ShowIndicatorDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIndicatorDetailResponse), nil
	}
}

type ShowPlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookInvoker) Invoke() (*model.ShowPlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookResponse), nil
	}
}

type ShowPlaybookInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookInstanceInvoker) Invoke() (*model.ShowPlaybookInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookInstanceResponse), nil
	}
}

type ShowPlaybookMonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookMonitorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookMonitorsInvoker) Invoke() (*model.ShowPlaybookMonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookMonitorsResponse), nil
	}
}

type ShowPlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookRuleInvoker) Invoke() (*model.ShowPlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookRuleResponse), nil
	}
}

type ShowPlaybookStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookStatisticsInvoker) Invoke() (*model.ShowPlaybookStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookStatisticsResponse), nil
	}
}

type ShowPlaybookTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookTopologyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookTopologyInvoker) Invoke() (*model.ShowPlaybookTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookTopologyResponse), nil
	}
}

type ShowPlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookVersionInvoker) Invoke() (*model.ShowPlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookVersionResponse), nil
	}
}

type UpdateAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlertRuleInvoker) Invoke() (*model.UpdateAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlertRuleResponse), nil
	}
}

type UpdateIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIndicatorInvoker) Invoke() (*model.UpdateIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIndicatorResponse), nil
	}
}

type UpdatePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookInvoker) Invoke() (*model.UpdatePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookResponse), nil
	}
}

type UpdatePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookActionInvoker) Invoke() (*model.UpdatePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookActionResponse), nil
	}
}

type UpdatePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookRuleInvoker) Invoke() (*model.UpdatePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookRuleResponse), nil
	}
}

type UpdatePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookVersionInvoker) Invoke() (*model.UpdatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookVersionResponse), nil
	}
}
