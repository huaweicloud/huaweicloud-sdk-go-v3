package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type AddAlarmRuleResourcesInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchDeleteAlarmTemplatesInvoker) Invoke() (*model.BatchDeleteAlarmTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAlarmTemplatesResponse), nil
	}
}

type BatchDeleteResourceGroupsInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchEnableAlarmRulesInvoker) Invoke() (*model.BatchEnableAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableAlarmRulesResponse), nil
	}
}

type CreateAlarmRulesInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateAlarmTemplateInvoker) Invoke() (*model.CreateAlarmTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlarmTemplateResponse), nil
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

type DeleteAlarmRuleResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmRuleResourcesInvoker) Invoke() (*model.DeleteAlarmRuleResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmRuleResourcesResponse), nil
	}
}

type ListAgentDimensionInfoInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAlarmTemplatesInvoker) Invoke() (*model.ListAlarmTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTemplatesResponse), nil
	}
}

type ListResourceGroupsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowResourceGroupInvoker) Invoke() (*model.ShowResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceGroupResponse), nil
	}
}

type UpdateAlarmRulePoliciesInvoker struct {
	*invoker.BaseInvoker
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
