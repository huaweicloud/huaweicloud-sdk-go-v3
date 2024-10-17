package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/das/v3/model"
)

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}

type CancelShareConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelShareConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelShareConnectionsInvoker) Invoke() (*model.CancelShareConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelShareConnectionsResponse), nil
	}
}

type ChangeSqlLimitSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSqlLimitSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSqlLimitSwitchStatusInvoker) Invoke() (*model.ChangeSqlLimitSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSqlLimitSwitchStatusResponse), nil
	}
}

type ChangeSqlSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSqlSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSqlSwitchInvoker) Invoke() (*model.ChangeSqlSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSqlSwitchResponse), nil
	}
}

type ChangeTransactionSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeTransactionSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeTransactionSwitchStatusInvoker) Invoke() (*model.ChangeTransactionSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeTransactionSwitchStatusResponse), nil
	}
}

type CreateShareConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShareConnectionsInvoker) Invoke() (*model.CreateShareConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareConnectionsResponse), nil
	}
}

type CreateSpaceAnalysisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSpaceAnalysisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSpaceAnalysisTaskInvoker) Invoke() (*model.CreateSpaceAnalysisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSpaceAnalysisTaskResponse), nil
	}
}

type CreateSqlLimitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlLimitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSqlLimitRulesInvoker) Invoke() (*model.CreateSqlLimitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlLimitRulesResponse), nil
	}
}

type CreateTuningInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTuningInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTuningInvoker) Invoke() (*model.CreateTuningResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTuningResponse), nil
	}
}

type DeleteDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDbUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDbUserInvoker) Invoke() (*model.DeleteDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDbUserResponse), nil
	}
}

type DeleteProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProcessInvoker) Invoke() (*model.DeleteProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProcessResponse), nil
	}
}

type DeleteSqlLimitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlLimitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSqlLimitRulesInvoker) Invoke() (*model.DeleteSqlLimitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlLimitRulesResponse), nil
	}
}

type ExportSlowQueryLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSlowQueryLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSlowQueryLogsInvoker) Invoke() (*model.ExportSlowQueryLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSlowQueryLogsResponse), nil
	}
}

type ExportSlowSqlStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSlowSqlStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSlowSqlStatisticsInvoker) Invoke() (*model.ExportSlowSqlStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSlowSqlStatisticsResponse), nil
	}
}

type ExportSlowSqlTemplatesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSlowSqlTemplatesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSlowSqlTemplatesDetailsInvoker) Invoke() (*model.ExportSlowSqlTemplatesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSlowSqlTemplatesDetailsResponse), nil
	}
}

type ExportSlowSqlTrendDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSlowSqlTrendDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSlowSqlTrendDetailsInvoker) Invoke() (*model.ExportSlowSqlTrendDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSlowSqlTrendDetailsResponse), nil
	}
}

type ExportSqlStatementsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSqlStatementsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSqlStatementsInvoker) Invoke() (*model.ExportSqlStatementsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSqlStatementsResponse), nil
	}
}

type ExportTopRiskInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTopRiskInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTopRiskInstancesInvoker) Invoke() (*model.ExportTopRiskInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTopRiskInstancesResponse), nil
	}
}

type ExportTopSqlTemplatesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTopSqlTemplatesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTopSqlTemplatesDetailsInvoker) Invoke() (*model.ExportTopSqlTemplatesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTopSqlTemplatesDetailsResponse), nil
	}
}

type ExportTopSqlTrendDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTopSqlTrendDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTopSqlTrendDetailsInvoker) Invoke() (*model.ExportTopSqlTrendDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTopSqlTrendDetailsResponse), nil
	}
}

type ListCloudDbaInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudDbaInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudDbaInstancesInvoker) Invoke() (*model.ListCloudDbaInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudDbaInstancesResponse), nil
	}
}

type ListDbUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDbUsersInvoker) Invoke() (*model.ListDbUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbUsersResponse), nil
	}
}

type ListInnodbLocksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInnodbLocksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInnodbLocksInvoker) Invoke() (*model.ListInnodbLocksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInnodbLocksResponse), nil
	}
}

type ListMetadataLocksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetadataLocksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetadataLocksInvoker) Invoke() (*model.ListMetadataLocksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetadataLocksResponse), nil
	}
}

type ListProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProcessesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProcessesInvoker) Invoke() (*model.ListProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProcessesResponse), nil
	}
}

type ListSpaceAnalysisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpaceAnalysisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpaceAnalysisInvoker) Invoke() (*model.ListSpaceAnalysisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpaceAnalysisResponse), nil
	}
}

type ListSqlLimitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlLimitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlLimitRulesInvoker) Invoke() (*model.ListSqlLimitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlLimitRulesResponse), nil
	}
}

type ListTransactionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransactionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransactionsInvoker) Invoke() (*model.ListTransactionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransactionsResponse), nil
	}
}

type ParseSqlLimitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseSqlLimitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ParseSqlLimitRulesInvoker) Invoke() (*model.ParseSqlLimitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseSqlLimitRulesResponse), nil
	}
}

type RegisterDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterDbUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterDbUserInvoker) Invoke() (*model.RegisterDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterDbUserResponse), nil
	}
}

type ShowDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDbUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDbUserInvoker) Invoke() (*model.ShowDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDbUserResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type ShowSqlExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlExecutionPlanInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlExecutionPlanInvoker) Invoke() (*model.ShowSqlExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlExecutionPlanResponse), nil
	}
}

type ShowSqlExplainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlExplainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlExplainInvoker) Invoke() (*model.ShowSqlExplainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlExplainResponse), nil
	}
}

type ShowSqlLimitJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlLimitJobInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlLimitJobInfoInvoker) Invoke() (*model.ShowSqlLimitJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlLimitJobInfoResponse), nil
	}
}

type ShowSqlLimitSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlLimitSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlLimitSwitchStatusInvoker) Invoke() (*model.ShowSqlLimitSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlLimitSwitchStatusResponse), nil
	}
}

type ShowSqlSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlSwitchStatusInvoker) Invoke() (*model.ShowSqlSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlSwitchStatusResponse), nil
	}
}

type ShowTransactionSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTransactionSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTransactionSwitchStatusInvoker) Invoke() (*model.ShowTransactionSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTransactionSwitchStatusResponse), nil
	}
}

type ShowTuningInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTuningInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTuningInvoker) Invoke() (*model.ShowTuningResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTuningResponse), nil
	}
}

type UpdateDbUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDbUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDbUserInvoker) Invoke() (*model.UpdateDbUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDbUserResponse), nil
	}
}

type UpdateSqlLimitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlLimitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSqlLimitRulesInvoker) Invoke() (*model.UpdateSqlLimitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlLimitRulesResponse), nil
	}
}
