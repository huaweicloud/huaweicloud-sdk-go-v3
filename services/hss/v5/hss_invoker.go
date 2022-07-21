package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

type ListHostStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostStatusInvoker) Invoke() (*model.ListHostStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostStatusResponse), nil
	}
}

type ListPasswordComplexityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPasswordComplexityInvoker) Invoke() (*model.ListPasswordComplexityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPasswordComplexityResponse), nil
	}
}

type ListRiskConfigCheckRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigCheckRulesInvoker) Invoke() (*model.ListRiskConfigCheckRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigCheckRulesResponse), nil
	}
}

type ListRiskConfigHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigHostsInvoker) Invoke() (*model.ListRiskConfigHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigHostsResponse), nil
	}
}

type ListRiskConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigsInvoker) Invoke() (*model.ListRiskConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigsResponse), nil
	}
}

type ListSecurityEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityEventsInvoker) Invoke() (*model.ListSecurityEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityEventsResponse), nil
	}
}

type ListVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) Invoke() (*model.ListVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulnerabilitiesResponse), nil
	}
}

type ListWeakPasswordUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWeakPasswordUsersInvoker) Invoke() (*model.ListWeakPasswordUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWeakPasswordUsersResponse), nil
	}
}

type ShowCheckRuleDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCheckRuleDetailInvoker) Invoke() (*model.ShowCheckRuleDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckRuleDetailResponse), nil
	}
}

type ShowRiskConfigDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRiskConfigDetailInvoker) Invoke() (*model.ShowRiskConfigDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRiskConfigDetailResponse), nil
	}
}
