package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type AddAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddAuditAgentInvoker) Invoke() (*model.AddAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditAgentResponse), nil
	}
}

type AddAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddAuditDatabaseInvoker) Invoke() (*model.AddAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditDatabaseResponse), nil
	}
}

type AddAuditDatabaseNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAuditDatabaseNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAuditDatabaseNewInvoker) Invoke() (*model.AddAuditDatabaseNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditDatabaseNewResponse), nil
	}
}

type AddRdsDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddRdsDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddRdsDatabaseInvoker) Invoke() (*model.AddRdsDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsDatabaseResponse), nil
	}
}

type AddRdsDatabaseNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRdsDatabaseNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRdsDatabaseNewInvoker) Invoke() (*model.AddRdsDatabaseNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsDatabaseNewResponse), nil
	}
}

type AddRdsNoAgentDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddRdsNoAgentDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddRdsNoAgentDatabaseInvoker) Invoke() (*model.AddRdsNoAgentDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRdsNoAgentDatabaseResponse), nil
	}
}

type BatchDeleteAuditScopeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAuditScopeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAuditScopeInvoker) Invoke() (*model.BatchDeleteAuditScopeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAuditScopeResponse), nil
	}
}

type BatchSetAuditAlarmLogStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetAuditAlarmLogStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetAuditAlarmLogStatusInvoker) Invoke() (*model.BatchSetAuditAlarmLogStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetAuditAlarmLogStatusResponse), nil
	}
}

type BindDbEncryptEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindDbEncryptEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindDbEncryptEipInvoker) Invoke() (*model.BindDbEncryptEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindDbEncryptEipResponse), nil
	}
}

type BindDbOmEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindDbOmEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindDbOmEipInvoker) Invoke() (*model.BindDbOmEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindDbOmEipResponse), nil
	}
}

type ChangeDbEncryptSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDbEncryptSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDbEncryptSecurityGroupInvoker) Invoke() (*model.ChangeDbEncryptSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDbEncryptSecurityGroupResponse), nil
	}
}

type ChangeDbOmSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDbOmSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDbOmSecurityGroupInvoker) Invoke() (*model.ChangeDbOmSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDbOmSecurityGroupResponse), nil
	}
}

type ConfirmUpgradeAuditInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmUpgradeAuditInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmUpgradeAuditInvoker) Invoke() (*model.ConfirmUpgradeAuditResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmUpgradeAuditResponse), nil
	}
}

type CountDbAccountSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountDbAccountSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountDbAccountSessionInvoker) Invoke() (*model.CountDbAccountSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountDbAccountSessionResponse), nil
	}
}

type CountDbClientSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountDbClientSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountDbClientSessionInvoker) Invoke() (*model.CountDbClientSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountDbClientSessionResponse), nil
	}
}

type CountInjectionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountInjectionStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountInjectionStatisticsInvoker) Invoke() (*model.CountInjectionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountInjectionStatisticsResponse), nil
	}
}

type CountOperationStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountOperationStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountOperationStatisticsInvoker) Invoke() (*model.CountOperationStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountOperationStatisticsResponse), nil
	}
}

type CountRiskTrendStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountRiskTrendStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountRiskTrendStatisticsInvoker) Invoke() (*model.CountRiskTrendStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountRiskTrendStatisticsResponse), nil
	}
}

type CountSessionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSessionStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSessionStatisticsInvoker) Invoke() (*model.CountSessionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSessionStatisticsResponse), nil
	}
}

type CountSqlStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSqlStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSqlStatisticsInvoker) Invoke() (*model.CountSqlStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSqlStatisticsResponse), nil
	}
}

type CountSqlTrendStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSqlTrendStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSqlTrendStatisticsInvoker) Invoke() (*model.CountSqlTrendStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSqlTrendStatisticsResponse), nil
	}
}

type CreateAuditRiskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuditRiskRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuditRiskRuleInvoker) Invoke() (*model.CreateAuditRiskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuditRiskRuleResponse), nil
	}
}

type CreateAuditScopeRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuditScopeRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuditScopeRuleInvoker) Invoke() (*model.CreateAuditScopeRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuditScopeRuleResponse), nil
	}
}

type CreateDbEncryptInstancePeriodInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbEncryptInstancePeriodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDbEncryptInstancePeriodInvoker) Invoke() (*model.CreateDbEncryptInstancePeriodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbEncryptInstancePeriodResponse), nil
	}
}

type CreateDbOmInstancePeriodInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbOmInstancePeriodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDbOmInstancePeriodInvoker) Invoke() (*model.CreateDbOmInstancePeriodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbOmInstancePeriodResponse), nil
	}
}

type CreateInstancesPeriodOrderInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateInstancesPeriodOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateInstancesPeriodOrderInvoker) Invoke() (*model.CreateInstancesPeriodOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstancesPeriodOrderResponse), nil
	}
}

type CreateInstancesPeriodOrderNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstancesPeriodOrderNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstancesPeriodOrderNewInvoker) Invoke() (*model.CreateInstancesPeriodOrderNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstancesPeriodOrderNewResponse), nil
	}
}

type CreateReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReportInvoker) Invoke() (*model.CreateReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportResponse), nil
	}
}

type CreateSensitiveMaskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSensitiveMaskRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSensitiveMaskRuleInvoker) Invoke() (*model.CreateSensitiveMaskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSensitiveMaskRuleResponse), nil
	}
}

type DeleteAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteAuditAgentInvoker) Invoke() (*model.DeleteAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditAgentResponse), nil
	}
}

type DeleteAuditAlarmLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditAlarmLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditAlarmLogInvoker) Invoke() (*model.DeleteAuditAlarmLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditAlarmLogResponse), nil
	}
}

type DeleteAuditBackupTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditBackupTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditBackupTaskInvoker) Invoke() (*model.DeleteAuditBackupTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditBackupTaskResponse), nil
	}
}

type DeleteAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteAuditDatabaseInvoker) Invoke() (*model.DeleteAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditDatabaseResponse), nil
	}
}

type DeleteAuditDatabaseNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditDatabaseNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditDatabaseNewInvoker) Invoke() (*model.DeleteAuditDatabaseNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditDatabaseNewResponse), nil
	}
}

type DeleteAuditReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditReportInvoker) Invoke() (*model.DeleteAuditReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditReportResponse), nil
	}
}

type DeleteAuditRuleRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditRuleRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditRuleRiskInvoker) Invoke() (*model.DeleteAuditRuleRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditRuleRiskResponse), nil
	}
}

type DeleteAuditScopeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditScopeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditScopeInvoker) Invoke() (*model.DeleteAuditScopeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditScopeResponse), nil
	}
}

type DeleteDbEncryptInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDbEncryptInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDbEncryptInstanceInvoker) Invoke() (*model.DeleteDbEncryptInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDbEncryptInstanceResponse), nil
	}
}

type DeleteDbOmInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDbOmInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDbOmInstanceInvoker) Invoke() (*model.DeleteDbOmInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDbOmInstanceResponse), nil
	}
}

type DeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteInstancesInvoker) Invoke() (*model.DeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesResponse), nil
	}
}

type DeleteInstancesNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstancesNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstancesNewInvoker) Invoke() (*model.DeleteInstancesNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesNewResponse), nil
	}
}

type DeleteSensitiveRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSensitiveRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSensitiveRulesInvoker) Invoke() (*model.DeleteSensitiveRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSensitiveRulesResponse), nil
	}
}

type DownloadAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadAuditAgentInvoker) Invoke() (*model.DownloadAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAuditAgentResponse), nil
	}
}

type DownloadAuditReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAuditReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAuditReportInvoker) Invoke() (*model.DownloadAuditReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAuditReportResponse), nil
	}
}

type ListAlarmTopicConfigInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAlarmTopicConfigInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAlarmTopicConfigInfoInvoker) Invoke() (*model.ListAlarmTopicConfigInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTopicConfigInfoResponse), nil
	}
}

type ListAlarmTopicConfigInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmTopicConfigInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmTopicConfigInfoNewInvoker) Invoke() (*model.ListAlarmTopicConfigInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmTopicConfigInfoNewResponse), nil
	}
}

type ListAuditAgentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditAgentInvoker) Invoke() (*model.ListAuditAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAgentResponse), nil
	}
}

type ListAuditAlarmLogInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditAlarmLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditAlarmLogInvoker) Invoke() (*model.ListAuditAlarmLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAlarmLogResponse), nil
	}
}

type ListAuditAlarmLogNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditAlarmLogNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditAlarmLogNewInvoker) Invoke() (*model.ListAuditAlarmLogNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAlarmLogNewResponse), nil
	}
}

type ListAuditBackupInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditBackupInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditBackupInfoInvoker) Invoke() (*model.ListAuditBackupInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditBackupInfoResponse), nil
	}
}

type ListAuditBackupRiskTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditBackupRiskTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditBackupRiskTemplatesInvoker) Invoke() (*model.ListAuditBackupRiskTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditBackupRiskTemplatesResponse), nil
	}
}

type ListAuditDatabasesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditDatabasesInvoker) Invoke() (*model.ListAuditDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditDatabasesResponse), nil
	}
}

type ListAuditDatabasesNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditDatabasesNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditDatabasesNewInvoker) Invoke() (*model.ListAuditDatabasesNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditDatabasesNewResponse), nil
	}
}

type ListAuditInstanceJobsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditInstanceJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditInstanceJobsInvoker) Invoke() (*model.ListAuditInstanceJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstanceJobsResponse), nil
	}
}

type ListAuditInstanceJobsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditInstanceJobsNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditInstanceJobsNewInvoker) Invoke() (*model.ListAuditInstanceJobsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstanceJobsNewResponse), nil
	}
}

type ListAuditInstancesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditInstancesInvoker) Invoke() (*model.ListAuditInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstancesResponse), nil
	}
}

type ListAuditInstancesNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditInstancesNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditInstancesNewInvoker) Invoke() (*model.ListAuditInstancesNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditInstancesNewResponse), nil
	}
}

type ListAuditObsBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditObsBucketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditObsBucketsInvoker) Invoke() (*model.ListAuditObsBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditObsBucketsResponse), nil
	}
}

type ListAuditOperateLogsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditOperateLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditOperateLogsInvoker) Invoke() (*model.ListAuditOperateLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditOperateLogsResponse), nil
	}
}

type ListAuditOperateLogsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditOperateLogsNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditOperateLogsNewInvoker) Invoke() (*model.ListAuditOperateLogsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditOperateLogsNewResponse), nil
	}
}

type ListAuditReportTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditReportTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditReportTemplatesInvoker) Invoke() (*model.ListAuditReportTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditReportTemplatesResponse), nil
	}
}

type ListAuditReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditReportsInvoker) Invoke() (*model.ListAuditReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditReportsResponse), nil
	}
}

type ListAuditRuleRisksInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditRuleRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditRuleRisksInvoker) Invoke() (*model.ListAuditRuleRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleRisksResponse), nil
	}
}

type ListAuditRuleRisksNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditRuleRisksNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditRuleRisksNewInvoker) Invoke() (*model.ListAuditRuleRisksNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleRisksNewResponse), nil
	}
}

type ListAuditRuleScopesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditRuleScopesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditRuleScopesInvoker) Invoke() (*model.ListAuditRuleScopesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleScopesResponse), nil
	}
}

type ListAuditRuleScopesNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditRuleScopesNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditRuleScopesNewInvoker) Invoke() (*model.ListAuditRuleScopesNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditRuleScopesNewResponse), nil
	}
}

type ListAuditSensitiveMasksInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditSensitiveMasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditSensitiveMasksInvoker) Invoke() (*model.ListAuditSensitiveMasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSensitiveMasksResponse), nil
	}
}

type ListAuditSensitiveMasksNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditSensitiveMasksNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditSensitiveMasksNewInvoker) Invoke() (*model.ListAuditSensitiveMasksNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSensitiveMasksNewResponse), nil
	}
}

type ListAuditSqlsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditSqlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuditSqlsInvoker) Invoke() (*model.ListAuditSqlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSqlsResponse), nil
	}
}

type ListAuditSqlsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditSqlsNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditSqlsNewInvoker) Invoke() (*model.ListAuditSqlsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditSqlsNewResponse), nil
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

type ListAuditTrendHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditTrendHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditTrendHistoryInvoker) Invoke() (*model.ListAuditTrendHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditTrendHistoryResponse), nil
	}
}

type ListAvailabilityZoneInfosInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAvailabilityZoneInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAvailabilityZoneInfosInvoker) Invoke() (*model.ListAvailabilityZoneInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneInfosResponse), nil
	}
}

type ListAvailabilityZoneInfosNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZoneInfosNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailabilityZoneInfosNewInvoker) Invoke() (*model.ListAvailabilityZoneInfosNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneInfosNewResponse), nil
	}
}

type ListDbEncryptInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbEncryptInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDbEncryptInstancesInvoker) Invoke() (*model.ListDbEncryptInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbEncryptInstancesResponse), nil
	}
}

type ListEcsSpecificationInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListEcsSpecificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListEcsSpecificationInvoker) Invoke() (*model.ListEcsSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcsSpecificationResponse), nil
	}
}

type ListEcsSpecificationNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcsSpecificationNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcsSpecificationNewInvoker) Invoke() (*model.ListEcsSpecificationNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcsSpecificationNewResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListSqlInjectionRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *RebootAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *RebootAuditInstanceInvoker) Invoke() (*model.RebootAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootAuditInstanceResponse), nil
	}
}

type RebootAuditInstanceNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootAuditInstanceNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootAuditInstanceNewInvoker) Invoke() (*model.RebootAuditInstanceNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootAuditInstanceNewResponse), nil
	}
}

type RebootDbEncryptInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootDbEncryptInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootDbEncryptInstanceInvoker) Invoke() (*model.RebootDbEncryptInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootDbEncryptInstanceResponse), nil
	}
}

type RebootDbOmInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootDbOmInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootDbOmInstanceInvoker) Invoke() (*model.RebootDbOmInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootDbOmInstanceResponse), nil
	}
}

type ResetDbEncryptPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetDbEncryptPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetDbEncryptPasswordInvoker) Invoke() (*model.ResetDbEncryptPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetDbEncryptPasswordResponse), nil
	}
}

type ResetDbOmPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetDbOmPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetDbOmPasswordInvoker) Invoke() (*model.ResetDbOmPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetDbOmPasswordResponse), nil
	}
}

type RestoreAuditBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreAuditBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreAuditBackupInvoker) Invoke() (*model.RestoreAuditBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreAuditBackupResponse), nil
	}
}

type RetryAuditBackupTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryAuditBackupTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryAuditBackupTaskInvoker) Invoke() (*model.RetryAuditBackupTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryAuditBackupTaskResponse), nil
	}
}

type SetAlarmTopicConfigInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SetAlarmTopicConfigInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SetAlarmTopicConfigInfoInvoker) Invoke() (*model.SetAlarmTopicConfigInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAlarmTopicConfigInfoResponse), nil
	}
}

type SetAlarmTopicConfigInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAlarmTopicConfigInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAlarmTopicConfigInfoNewInvoker) Invoke() (*model.SetAlarmTopicConfigInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAlarmTopicConfigInfoNewResponse), nil
	}
}

type SetAuditAlarmLogStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditAlarmLogStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditAlarmLogStatusInvoker) Invoke() (*model.SetAuditAlarmLogStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditAlarmLogStatusResponse), nil
	}
}

type SetAuditAutoBackupTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditAutoBackupTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditAutoBackupTemplateInvoker) Invoke() (*model.SetAuditAutoBackupTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditAutoBackupTemplateResponse), nil
	}
}

type SetAuditBackupRiskSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditBackupRiskSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditBackupRiskSwitchInvoker) Invoke() (*model.SetAuditBackupRiskSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditBackupRiskSwitchResponse), nil
	}
}

type SetAuditBackupSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditBackupSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditBackupSwitchInvoker) Invoke() (*model.SetAuditBackupSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditBackupSwitchResponse), nil
	}
}

type SetAuditInstanceTimeZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditInstanceTimeZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditInstanceTimeZoneInvoker) Invoke() (*model.SetAuditInstanceTimeZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditInstanceTimeZoneResponse), nil
	}
}

type SetAuditScopeRuleSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditScopeRuleSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditScopeRuleSwitchInvoker) Invoke() (*model.SetAuditScopeRuleSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditScopeRuleSwitchResponse), nil
	}
}

type SetRiskOperationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRiskOperationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRiskOperationPolicyInvoker) Invoke() (*model.SetRiskOperationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRiskOperationPolicyResponse), nil
	}
}

type SetRiskRuleRankInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRiskRuleRankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRiskRuleRankInvoker) Invoke() (*model.SetRiskRuleRankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRiskRuleRankResponse), nil
	}
}

type SetSensitiveMaskRuleSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSensitiveMaskRuleSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSensitiveMaskRuleSwitchInvoker) Invoke() (*model.SetSensitiveMaskRuleSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSensitiveMaskRuleSwitchResponse), nil
	}
}

type SetSensitiveResultSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSensitiveResultSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSensitiveResultSwitchInvoker) Invoke() (*model.SetSensitiveResultSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSensitiveResultSwitchResponse), nil
	}
}

type SetSensitiveSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSensitiveSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSensitiveSwitchInvoker) Invoke() (*model.SetSensitiveSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSensitiveSwitchResponse), nil
	}
}

type ShowAuditBackRiskTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditBackRiskTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditBackRiskTemplateInvoker) Invoke() (*model.ShowAuditBackRiskTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditBackRiskTemplateResponse), nil
	}
}

type ShowAuditBackupStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditBackupStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditBackupStatusInvoker) Invoke() (*model.ShowAuditBackupStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditBackupStatusResponse), nil
	}
}

type ShowAuditQuotaInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowAuditQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowAuditQuotaInvoker) Invoke() (*model.ShowAuditQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditQuotaResponse), nil
	}
}

type ShowAuditQuotaNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditQuotaNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditQuotaNewInvoker) Invoke() (*model.ShowAuditQuotaNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditQuotaNewResponse), nil
	}
}

type ShowAuditRuleRiskInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowAuditRuleRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowAuditRuleRiskInvoker) Invoke() (*model.ShowAuditRuleRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditRuleRiskResponse), nil
	}
}

type ShowAuditRuleRiskNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditRuleRiskNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditRuleRiskNewInvoker) Invoke() (*model.ShowAuditRuleRiskNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditRuleRiskNewResponse), nil
	}
}

type ShowAuditStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditStatisticsInvoker) Invoke() (*model.ShowAuditStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditStatisticsResponse), nil
	}
}

type ShowAuditTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditTaskStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditTaskStatusInvoker) Invoke() (*model.ShowAuditTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditTaskStatusResponse), nil
	}
}

type ShowAuditTopicReportSchedulerConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditTopicReportSchedulerConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditTopicReportSchedulerConfigInvoker) Invoke() (*model.ShowAuditTopicReportSchedulerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditTopicReportSchedulerConfigResponse), nil
	}
}

type ShowAuditUpgradeStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditUpgradeStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditUpgradeStatusInvoker) Invoke() (*model.ShowAuditUpgradeStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditUpgradeStatusResponse), nil
	}
}

type ShowBackupRiskBucketPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupRiskBucketPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackupRiskBucketPathInvoker) Invoke() (*model.ShowBackupRiskBucketPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupRiskBucketPathResponse), nil
	}
}

type ShowInstanceMonitorInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMonitorInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceMonitorInfoInvoker) Invoke() (*model.ShowInstanceMonitorInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceMonitorInfoResponse), nil
	}
}

type ShowInstanceQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceQuotaInvoker) Invoke() (*model.ShowInstanceQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceQuotaResponse), nil
	}
}

type ShowSensitiveMaskSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSensitiveMaskSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSensitiveMaskSwitchInvoker) Invoke() (*model.ShowSensitiveMaskSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSensitiveMaskSwitchResponse), nil
	}
}

type ShowSensitiveResultSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSensitiveResultSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSensitiveResultSwitchInvoker) Invoke() (*model.ShowSensitiveResultSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSensitiveResultSwitchResponse), nil
	}
}

type ShowServerVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerVersionInvoker) Invoke() (*model.ShowServerVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerVersionResponse), nil
	}
}

type ShowSqlDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlDetailInvoker) Invoke() (*model.ShowSqlDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlDetailResponse), nil
	}
}

type StartAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *StartAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *StartAuditInstanceInvoker) Invoke() (*model.StartAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAuditInstanceResponse), nil
	}
}

type StartAuditInstanceNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAuditInstanceNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAuditInstanceNewInvoker) Invoke() (*model.StartAuditInstanceNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAuditInstanceNewResponse), nil
	}
}

type StartDbEncryptInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDbEncryptInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDbEncryptInstanceInvoker) Invoke() (*model.StartDbEncryptInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDbEncryptInstanceResponse), nil
	}
}

type StartDbOmInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDbOmInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDbOmInstanceInvoker) Invoke() (*model.StartDbOmInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDbOmInstanceResponse), nil
	}
}

type StopAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *StopAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *StopAuditInstanceInvoker) Invoke() (*model.StopAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAuditInstanceResponse), nil
	}
}

type StopAuditInstanceNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAuditInstanceNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopAuditInstanceNewInvoker) Invoke() (*model.StopAuditInstanceNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAuditInstanceNewResponse), nil
	}
}

type StopDbEncryptInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopDbEncryptInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopDbEncryptInstanceInvoker) Invoke() (*model.StopDbEncryptInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopDbEncryptInstanceResponse), nil
	}
}

type StopDbOmInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopDbOmInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopDbOmInstanceInvoker) Invoke() (*model.StopDbOmInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopDbOmInstanceResponse), nil
	}
}

type SwitchAgentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchAgentInvoker) Invoke() (*model.SwitchAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAgentResponse), nil
	}
}

type SwitchAuditDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchAuditDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchAuditDatabaseInvoker) Invoke() (*model.SwitchAuditDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAuditDatabaseResponse), nil
	}
}

type SwitchAuditDatabaseNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAuditDatabaseNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAuditDatabaseNewInvoker) Invoke() (*model.SwitchAuditDatabaseNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAuditDatabaseNewResponse), nil
	}
}

type SwitchRiskRuleInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchRiskRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SwitchRiskRuleInvoker) Invoke() (*model.SwitchRiskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRiskRuleResponse), nil
	}
}

type SwitchRiskRuleNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchRiskRuleNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchRiskRuleNewInvoker) Invoke() (*model.SwitchRiskRuleNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRiskRuleNewResponse), nil
	}
}

type UnbindDbEncryptEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnbindDbEncryptEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnbindDbEncryptEipInvoker) Invoke() (*model.UnbindDbEncryptEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnbindDbEncryptEipResponse), nil
	}
}

type UnbindDbOmEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnbindDbOmEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnbindDbOmEipInvoker) Invoke() (*model.UnbindDbOmEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnbindDbOmEipResponse), nil
	}
}

type UpdateAuditInstanceInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateAuditInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateAuditInstanceInvoker) Invoke() (*model.UpdateAuditInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditInstanceResponse), nil
	}
}

type UpdateAuditInstanceNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditInstanceNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditInstanceNewInvoker) Invoke() (*model.UpdateAuditInstanceNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditInstanceNewResponse), nil
	}
}

type UpdateAuditRiskBucketPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditRiskBucketPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditRiskBucketPathInvoker) Invoke() (*model.UpdateAuditRiskBucketPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditRiskBucketPathResponse), nil
	}
}

type UpdateAuditScopeRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditScopeRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditScopeRuleInvoker) Invoke() (*model.UpdateAuditScopeRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditScopeRuleResponse), nil
	}
}

type UpdateAuditSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateAuditSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateAuditSecurityGroupInvoker) Invoke() (*model.UpdateAuditSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditSecurityGroupResponse), nil
	}
}

type UpdateAuditSecurityGroupNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditSecurityGroupNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditSecurityGroupNewInvoker) Invoke() (*model.UpdateAuditSecurityGroupNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditSecurityGroupNewResponse), nil
	}
}

type UpdateAuditTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditTaskStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditTaskStatusInvoker) Invoke() (*model.UpdateAuditTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditTaskStatusResponse), nil
	}
}

type UpdateAuditTopicReportSchedulerConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditTopicReportSchedulerConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditTopicReportSchedulerConfigInvoker) Invoke() (*model.UpdateAuditTopicReportSchedulerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditTopicReportSchedulerConfigResponse), nil
	}
}

type UpdateDbEncryptInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDbEncryptInstanceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDbEncryptInstanceNameInvoker) Invoke() (*model.UpdateDbEncryptInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDbEncryptInstanceNameResponse), nil
	}
}

type UpdateDbOmInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDbOmInstanceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDbOmInstanceNameInvoker) Invoke() (*model.UpdateDbOmInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDbOmInstanceNameResponse), nil
	}
}

type UpdateSensitiveMaskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSensitiveMaskRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSensitiveMaskRuleInvoker) Invoke() (*model.UpdateSensitiveMaskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSensitiveMaskRuleResponse), nil
	}
}

type UploadAuditDbConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAuditDbConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAuditDbConfigInvoker) Invoke() (*model.UploadAuditDbConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAuditDbConfigResponse), nil
	}
}

type AddAuditAgentNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAuditAgentNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAuditAgentNewInvoker) Invoke() (*model.AddAuditAgentNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAuditAgentNewResponse), nil
	}
}

type CreateAuditDbAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuditDbAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuditDbAgentInvoker) Invoke() (*model.CreateAuditDbAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuditDbAgentResponse), nil
	}
}

type DeleteAuditAgentNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditAgentNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditAgentNewInvoker) Invoke() (*model.DeleteAuditAgentNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditAgentNewResponse), nil
	}
}

type DownloadAuditAgentNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAuditAgentNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAuditAgentNewInvoker) Invoke() (*model.DownloadAuditAgentNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAuditAgentNewResponse), nil
	}
}

type ListAuditAgentNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditAgentNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditAgentNewInvoker) Invoke() (*model.ListAuditAgentNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditAgentNewResponse), nil
	}
}

type SwitchAgentNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAgentNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAgentNewInvoker) Invoke() (*model.SwitchAgentNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAgentNewResponse), nil
	}
}

type BatchAddAuditWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddAuditWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddAuditWhitelistInvoker) Invoke() (*model.BatchAddAuditWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddAuditWhitelistResponse), nil
	}
}

type CreateAuditSqlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuditSqlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuditSqlRuleInvoker) Invoke() (*model.CreateAuditSqlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuditSqlRuleResponse), nil
	}
}

type DeleteAuditRuleSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditRuleSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditRuleSqlInvoker) Invoke() (*model.DeleteAuditRuleSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditRuleSqlResponse), nil
	}
}

type DeleteAuditWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuditWhitelistInvoker) Invoke() (*model.DeleteAuditWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditWhitelistResponse), nil
	}
}

type ListSqlInjectionRulesNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlInjectionRulesNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlInjectionRulesNewInvoker) Invoke() (*model.ListSqlInjectionRulesNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlInjectionRulesNewResponse), nil
	}
}

type ListWhitelistsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWhitelistsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWhitelistsInvoker) Invoke() (*model.ListWhitelistsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWhitelistsResponse), nil
	}
}

type SetAuditSqlRuleSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditSqlRuleSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAuditSqlRuleSwitchInvoker) Invoke() (*model.SetAuditSqlRuleSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditSqlRuleSwitchResponse), nil
	}
}

type SetSqlRuleRankInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSqlRuleRankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSqlRuleRankInvoker) Invoke() (*model.SetSqlRuleRankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSqlRuleRankResponse), nil
	}
}

type UpdateAuditSqlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditSqlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditSqlRuleInvoker) Invoke() (*model.UpdateAuditSqlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditSqlRuleResponse), nil
	}
}

type UpdateAuditWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuditWhitelistInvoker) Invoke() (*model.UpdateAuditWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditWhitelistResponse), nil
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

type ListAuditTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditTagsInvoker) Invoke() (*model.ListAuditTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditTagsResponse), nil
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
