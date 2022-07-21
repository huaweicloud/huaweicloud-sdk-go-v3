package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
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
