package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type AddAlarmRuleResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAlarmRuleResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAlarmRuleResourcesInvoker) Invoke() (*model.AddAlarmRuleResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAlarmRuleResourcesResponse), nil
	}
}

type BatchCreateResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateResourcesInvoker) Invoke() (*model.BatchCreateResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateResourcesResponse), nil
	}
}

type BatchDeleteAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAlarmRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAlarmRulesInvoker) Invoke() (*model.BatchDeleteAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAlarmRulesResponse), nil
	}
}

type BatchDeleteAlarmTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAlarmTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAlarmTemplatesInvoker) Invoke() (*model.BatchDeleteAlarmTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAlarmTemplatesResponse), nil
	}
}

type BatchDeleteNotificationMasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteNotificationMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteNotificationMasksInvoker) Invoke() (*model.BatchDeleteNotificationMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteNotificationMasksResponse), nil
	}
}

type BatchDeleteOneClickAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteOneClickAlarmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteOneClickAlarmsInvoker) Invoke() (*model.BatchDeleteOneClickAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteOneClickAlarmsResponse), nil
	}
}

type BatchDeleteResourceGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourceGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteResourceGroupsInvoker) Invoke() (*model.BatchDeleteResourceGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourceGroupsResponse), nil
	}
}

type BatchDeleteResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteResourcesInvoker) Invoke() (*model.BatchDeleteResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourcesResponse), nil
	}
}

type BatchEnableAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableAlarmRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableAlarmRulesInvoker) Invoke() (*model.BatchEnableAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableAlarmRulesResponse), nil
	}
}

type BatchUpdateNotificationMaskTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNotificationMaskTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateNotificationMaskTimeInvoker) Invoke() (*model.BatchUpdateNotificationMaskTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNotificationMaskTimeResponse), nil
	}
}

type BatchUpdateNotificationMasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNotificationMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateNotificationMasksInvoker) Invoke() (*model.BatchUpdateNotificationMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNotificationMasksResponse), nil
	}
}

type BatchUpdateOneClickAlarmPoliciesEnabledStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateOneClickAlarmPoliciesEnabledStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateOneClickAlarmPoliciesEnabledStateInvoker) Invoke() (*model.BatchUpdateOneClickAlarmPoliciesEnabledStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateOneClickAlarmPoliciesEnabledStateResponse), nil
	}
}

type BatchUpdateOneClickAlarmsEnabledStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateOneClickAlarmsEnabledStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateOneClickAlarmsEnabledStateInvoker) Invoke() (*model.BatchUpdateOneClickAlarmsEnabledStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateOneClickAlarmsEnabledStateResponse), nil
	}
}

type BatchUpdateWidgetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateWidgetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateWidgetsInvoker) Invoke() (*model.BatchUpdateWidgetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateWidgetsResponse), nil
	}
}

type CreateAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlarmRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlarmRulesInvoker) Invoke() (*model.CreateAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmRulesResponse), nil
	}
}

type CreateAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlarmTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlarmTemplateInvoker) Invoke() (*model.CreateAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmTemplateResponse), nil
	}
}

type CreateDashboardWidgetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDashboardWidgetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDashboardWidgetsInvoker) Invoke() (*model.CreateDashboardWidgetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDashboardWidgetsResponse), nil
	}
}

type CreateOneClickAlarmInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOneClickAlarmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOneClickAlarmInvoker) Invoke() (*model.CreateOneClickAlarmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOneClickAlarmResponse), nil
	}
}

type CreateOneDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOneDashboardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOneDashboardInvoker) Invoke() (*model.CreateOneDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOneDashboardResponse), nil
	}
}

type CreateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceGroupInvoker) Invoke() (*model.CreateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceGroupResponse), nil
	}
}

type DeleteAlarmRuleResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmRuleResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlarmRuleResourcesInvoker) Invoke() (*model.DeleteAlarmRuleResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmRuleResourcesResponse), nil
	}
}

type DeleteDashboardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDashboardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDashboardsInvoker) Invoke() (*model.DeleteDashboardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDashboardsResponse), nil
	}
}

type DeleteOneWidgetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOneWidgetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOneWidgetInvoker) Invoke() (*model.DeleteOneWidgetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOneWidgetResponse), nil
	}
}

type ListAgentDimensionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentDimensionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentDimensionInfoInvoker) Invoke() (*model.ListAgentDimensionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentDimensionInfoResponse), nil
	}
}

type ListAlarmHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmHistoriesInvoker) Invoke() (*model.ListAlarmHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmHistoriesResponse), nil
	}
}

type ListAlarmRulePoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmRulePoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmRulePoliciesInvoker) Invoke() (*model.ListAlarmRulePoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmRulePoliciesResponse), nil
	}
}

type ListAlarmRuleResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmRuleResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmRuleResourcesInvoker) Invoke() (*model.ListAlarmRuleResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmRuleResourcesResponse), nil
	}
}

type ListAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmRulesInvoker) Invoke() (*model.ListAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmRulesResponse), nil
	}
}

type ListAlarmTemplateAssociationAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmTemplateAssociationAlarmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmTemplateAssociationAlarmsInvoker) Invoke() (*model.ListAlarmTemplateAssociationAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTemplateAssociationAlarmsResponse), nil
	}
}

type ListAlarmTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmTemplatesInvoker) Invoke() (*model.ListAlarmTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTemplatesResponse), nil
	}
}

type ListCesTargetProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCesTargetProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCesTargetProjectTagsInvoker) Invoke() (*model.ListCesTargetProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCesTargetProjectTagsResponse), nil
	}
}

type ListDashboardInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDashboardInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDashboardInfosInvoker) Invoke() (*model.ListDashboardInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDashboardInfosResponse), nil
	}
}

type ListDashboardWidgetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDashboardWidgetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDashboardWidgetsInvoker) Invoke() (*model.ListDashboardWidgetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDashboardWidgetsResponse), nil
	}
}

type ListNotificationMaskResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationMaskResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationMaskResourcesInvoker) Invoke() (*model.ListNotificationMaskResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationMaskResourcesResponse), nil
	}
}

type ListNotificationMasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationMasksInvoker) Invoke() (*model.ListNotificationMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationMasksResponse), nil
	}
}

type ListOneClickAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOneClickAlarmRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOneClickAlarmRulesInvoker) Invoke() (*model.ListOneClickAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOneClickAlarmRulesResponse), nil
	}
}

type ListOneClickAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOneClickAlarmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOneClickAlarmsInvoker) Invoke() (*model.ListOneClickAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOneClickAlarmsResponse), nil
	}
}

type ListResourceGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceGroupsInvoker) Invoke() (*model.ListResourceGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceGroupsResponse), nil
	}
}

type ListResourceGroupsServicesResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceGroupsServicesResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceGroupsServicesResourcesInvoker) Invoke() (*model.ListResourceGroupsServicesResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceGroupsServicesResourcesResponse), nil
	}
}

type ShowAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmTemplateInvoker) Invoke() (*model.ShowAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmTemplateResponse), nil
	}
}

type ShowResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceGroupInvoker) Invoke() (*model.ShowResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceGroupResponse), nil
	}
}

type ShowWidgetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWidgetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWidgetInvoker) Invoke() (*model.ShowWidgetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWidgetResponse), nil
	}
}

type UpdateAlarmNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlarmNotificationsInvoker) Invoke() (*model.UpdateAlarmNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmNotificationsResponse), nil
	}
}

type UpdateAlarmRulePoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmRulePoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlarmRulePoliciesInvoker) Invoke() (*model.UpdateAlarmRulePoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmRulePoliciesResponse), nil
	}
}

type UpdateAlarmTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlarmTemplateInvoker) Invoke() (*model.UpdateAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmTemplateResponse), nil
	}
}

type UpdateDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDashboardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDashboardInvoker) Invoke() (*model.UpdateDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDashboardResponse), nil
	}
}

type UpdateNotificationMaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotificationMaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotificationMaskInvoker) Invoke() (*model.UpdateNotificationMaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotificationMaskResponse), nil
	}
}

type UpdateOneClickAlarmNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOneClickAlarmNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOneClickAlarmNotificationsInvoker) Invoke() (*model.UpdateOneClickAlarmNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOneClickAlarmNotificationsResponse), nil
	}
}

type UpdateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceGroupInvoker) Invoke() (*model.UpdateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceGroupResponse), nil
	}
}
