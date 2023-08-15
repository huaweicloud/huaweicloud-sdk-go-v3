package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/secmaster/v2/model"
)

type ChangeAlertInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateBatchOrderAlertsInvoker) Invoke() (*model.CreateBatchOrderAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

type CreateDataobjectRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataobjectRelationInvoker) Invoke() (*model.CreateDataobjectRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataobjectRelationResponse), nil
	}
}

type CreateIncidentInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateIndicatorInvoker) Invoke() (*model.CreateIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIndicatorResponse), nil
	}
}

type CreatePlaybookInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreatePlaybookVersionInvoker) Invoke() (*model.CreatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookVersionResponse), nil
	}
}

type DeleteAlertInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteAlertRuleInvoker) Invoke() (*model.DeleteAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertRuleResponse), nil
	}
}

type DeleteDataobjectRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataobjectRelationInvoker) Invoke() (*model.DeleteDataobjectRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataobjectRelationResponse), nil
	}
}

type DeleteIncidentInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAlertsInvoker) Invoke() (*model.ListAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertsResponse), nil
	}
}

type ListDataobjectRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataobjectRelationInvoker) Invoke() (*model.ListDataobjectRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataobjectRelationResponse), nil
	}
}

type ListIncidentTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncidentTypesInvoker) Invoke() (*model.ListIncidentTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncidentTypesResponse), nil
	}
}

type ListIncidentsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPlaybooksInvoker) Invoke() (*model.ListPlaybooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybooksResponse), nil
	}
}

type ShowAlertInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdatePlaybookVersionInvoker) Invoke() (*model.UpdatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookVersionResponse), nil
	}
}
