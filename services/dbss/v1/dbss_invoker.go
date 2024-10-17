package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type AddAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAuditDatabaseInvoker) Invoke() (*model.AddAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditDatabaseResponse), nil
	}
}

type AddRdsDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRdsDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRdsDatabaseInvoker) Invoke() (*model.AddRdsDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsDatabaseResponse), nil
	}
}

type AddRdsNoAgentDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRdsNoAgentDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRdsNoAgentDatabaseInvoker) Invoke() (*model.AddRdsNoAgentDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsNoAgentDatabaseResponse), nil
	}
}

type CreateInstancesPeriodOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstancesPeriodOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstancesPeriodOrderInvoker) Invoke() (*model.CreateInstancesPeriodOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstancesPeriodOrderResponse), nil
	}
}

type DeleteAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditDatabaseInvoker) Invoke() (*model.DeleteAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditDatabaseResponse), nil
	}
}

type DeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstancesInvoker) Invoke() (*model.DeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesResponse), nil
	}
}

type ListAuditAlarmLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditAlarmLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditAlarmLogInvoker) Invoke() (*model.ListAuditAlarmLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAlarmLogResponse), nil
	}
}

type ListAuditDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditInstanceJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditOperateLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditRuleRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditRuleScopesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditSensitiveMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditSensitiveMasksInvoker) Invoke() (*model.ListAuditSensitiveMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSensitiveMasksResponse), nil
	}
}

type ListAuditSqlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditSqlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditSqlsInvoker) Invoke() (*model.ListAuditSqlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSqlsResponse), nil
	}
}

type ListAuditSummaryInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditSummaryInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditSummaryInfosInvoker) Invoke() (*model.ListAuditSummaryInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSummaryInfosResponse), nil
	}
}

type ListAvailabilityZoneInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZoneInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListEcsSpecificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcsSpecificationInvoker) Invoke() (*model.ListEcsSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcsSpecificationResponse), nil
	}
}

type ListRdsDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRdsDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRdsDatabasesInvoker) Invoke() (*model.ListRdsDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRdsDatabasesResponse), nil
	}
}

type ListSqlInjectionRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlInjectionRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlInjectionRulesInvoker) Invoke() (*model.ListSqlInjectionRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlInjectionRulesResponse), nil
	}
}

type RebootAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootAuditInstanceInvoker) Invoke() (*model.RebootAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootAuditInstanceResponse), nil
	}
}

type ShowAuditQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowAuditRuleRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditRuleRiskInvoker) Invoke() (*model.ShowAuditRuleRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditRuleRiskResponse), nil
	}
}

type StartAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAuditInstanceInvoker) Invoke() (*model.StartAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAuditInstanceResponse), nil
	}
}

type StopAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopAuditInstanceInvoker) Invoke() (*model.StopAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAuditInstanceResponse), nil
	}
}

type SwitchAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAuditDatabaseInvoker) Invoke() (*model.SwitchAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAuditDatabaseResponse), nil
	}
}

type SwitchRiskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchRiskRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchRiskRuleInvoker) Invoke() (*model.SwitchRiskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRiskRuleResponse), nil
	}
}

type UpdateAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditInstanceInvoker) Invoke() (*model.UpdateAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditInstanceResponse), nil
	}
}

type UpdateAuditSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditSecurityGroupInvoker) Invoke() (*model.UpdateAuditSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditSecurityGroupResponse), nil
	}
}

type AddAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAuditAgentInvoker) Invoke() (*model.AddAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditAgentResponse), nil
	}
}

type DeleteAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditAgentInvoker) Invoke() (*model.DeleteAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditAgentResponse), nil
	}
}

type DownloadAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAuditAgentInvoker) Invoke() (*model.DownloadAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAuditAgentResponse), nil
	}
}

type ListAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditAgentInvoker) Invoke() (*model.ListAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAgentResponse), nil
	}
}

type SwitchAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAgentInvoker) Invoke() (*model.SwitchAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAgentResponse), nil
	}
}

type BatchAddResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchDeleteResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CountResourceInstanceByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountResourceInstanceByTagInvoker) Invoke() (*model.CountResourceInstanceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourceInstanceByTagResponse), nil
	}
}

type ListProjectResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListResourceInstanceByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstanceByTagInvoker) Invoke() (*model.ListResourceInstanceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstanceByTagResponse), nil
	}
}
