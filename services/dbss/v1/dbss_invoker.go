package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type AddRdsNoAgentDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRdsNoAgentDatabaseInvoker) Invoke() (*model.AddRdsNoAgentDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsNoAgentDatabaseResponse), nil
	}
}

type BatchAddResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddResourceTagInvoker) Invoke() (*model.BatchAddResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddResourceTagResponse), nil
	}
}

type BatchDeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourceTagInvoker) Invoke() (*model.BatchDeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourceTagResponse), nil
	}
}

type CountResourceInstanceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourceInstanceByTagInvoker) Invoke() (*model.CountResourceInstanceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourceInstanceByTagResponse), nil
	}
}

type CreateInstancesPeriodOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstancesPeriodOrderInvoker) Invoke() (*model.CreateInstancesPeriodOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstancesPeriodOrderResponse), nil
	}
}

type ListAuditDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditDatabasesInvoker) Invoke() (*model.ListAuditDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditDatabasesResponse), nil
	}
}

type ListAuditInstanceJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditInstanceJobsInvoker) Invoke() (*model.ListAuditInstanceJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstanceJobsResponse), nil
	}
}

type ListAuditInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditInstancesInvoker) Invoke() (*model.ListAuditInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstancesResponse), nil
	}
}

type ListAuditOperateLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditOperateLogsInvoker) Invoke() (*model.ListAuditOperateLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditOperateLogsResponse), nil
	}
}

type ListAuditRuleRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditRuleRisksInvoker) Invoke() (*model.ListAuditRuleRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleRisksResponse), nil
	}
}

type ListAuditRuleScopesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditRuleScopesInvoker) Invoke() (*model.ListAuditRuleScopesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleScopesResponse), nil
	}
}

type ListAuditSensitiveMasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditSensitiveMasksInvoker) Invoke() (*model.ListAuditSensitiveMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSensitiveMasksResponse), nil
	}
}

type ListAvailabilityZoneInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZoneInfosInvoker) Invoke() (*model.ListAvailabilityZoneInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneInfosResponse), nil
	}
}

type ListEcsSpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcsSpecificationInvoker) Invoke() (*model.ListEcsSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcsSpecificationResponse), nil
	}
}

type ListProjectResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectResourceTagsInvoker) Invoke() (*model.ListProjectResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectResourceTagsResponse), nil
	}
}

type ListResourceInstanceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstanceByTagInvoker) Invoke() (*model.ListResourceInstanceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstanceByTagResponse), nil
	}
}

type ListSqlInjectionRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlInjectionRulesInvoker) Invoke() (*model.ListSqlInjectionRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlInjectionRulesResponse), nil
	}
}

type ShowAuditQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditQuotaInvoker) Invoke() (*model.ShowAuditQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditQuotaResponse), nil
	}
}

type ShowAuditRuleRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditRuleRiskInvoker) Invoke() (*model.ShowAuditRuleRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditRuleRiskResponse), nil
	}
}

type SwitchAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAgentInvoker) Invoke() (*model.SwitchAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAgentResponse), nil
	}
}

type SwitchRiskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchRiskRuleInvoker) Invoke() (*model.SwitchRiskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRiskRuleResponse), nil
	}
}

type UpdateAuditSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditSecurityGroupInvoker) Invoke() (*model.UpdateAuditSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditSecurityGroupResponse), nil
	}
}
