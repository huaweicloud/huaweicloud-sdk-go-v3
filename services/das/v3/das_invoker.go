package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/das/v3/model"
)

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

type CreateInstanceConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceConnectionInvoker) Invoke() (*model.CreateInstanceConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceConnectionResponse), nil
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

type ExecuteExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteExportTaskInvoker) Invoke() (*model.ExecuteExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteExportTaskResponse), nil
	}
}

type ExecuteImportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteImportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteImportTaskInvoker) Invoke() (*model.ExecuteImportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteImportTaskResponse), nil
	}
}

type ListConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectionsInvoker) Invoke() (*model.ListConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionsResponse), nil
	}
}

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

type AddTasksNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTasksNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddTasksNewInvoker) Invoke() (*model.AddTasksNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTasksNewResponse), nil
	}
}

type BatchAddFullSqlTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddFullSqlTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddFullSqlTasksInvoker) Invoke() (*model.BatchAddFullSqlTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddFullSqlTasksResponse), nil
	}
}

type BatchSetSqlSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetSqlSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetSqlSwitchNewInvoker) Invoke() (*model.BatchSetSqlSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetSqlSwitchNewResponse), nil
	}
}

type CancelConnectionProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelConnectionProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelConnectionProcessInvoker) Invoke() (*model.CancelConnectionProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelConnectionProcessResponse), nil
	}
}

type ChangeQuotaNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeQuotaNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeQuotaNewInvoker) Invoke() (*model.ChangeQuotaNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeQuotaNewResponse), nil
	}
}

type CheckHealthReportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckHealthReportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckHealthReportTaskInvoker) Invoke() (*model.CheckHealthReportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckHealthReportTaskResponse), nil
	}
}

type CreateBinlogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBinlogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBinlogTaskInvoker) Invoke() (*model.CreateBinlogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBinlogTaskResponse), nil
	}
}

type CreateDbsConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDbsConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDbsConnectionInvoker) Invoke() (*model.CreateDbsConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDbsConnectionResponse), nil
	}
}

type CreateFullSqlBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFullSqlBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFullSqlBucketInvoker) Invoke() (*model.CreateFullSqlBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFullSqlBucketResponse), nil
	}
}

type CreateIndexUsageExportTaskNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIndexUsageExportTaskNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIndexUsageExportTaskNewInvoker) Invoke() (*model.CreateIndexUsageExportTaskNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIndexUsageExportTaskNewResponse), nil
	}
}

type CreateInstanceHealthReportTaskNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceHealthReportTaskNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceHealthReportTaskNewInvoker) Invoke() (*model.CreateInstanceHealthReportTaskNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceHealthReportTaskNewResponse), nil
	}
}

type CreateWdrReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWdrReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWdrReportInvoker) Invoke() (*model.CreateWdrReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWdrReportResponse), nil
	}
}

type DeleteBinlogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBinlogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBinlogTaskInvoker) Invoke() (*model.DeleteBinlogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBinlogTaskResponse), nil
	}
}

type DeleteDbObjNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDbObjNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDbObjNewInvoker) Invoke() (*model.DeleteDbObjNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDbObjNewResponse), nil
	}
}

type DeleteExportTaskNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExportTaskNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteExportTaskNewInvoker) Invoke() (*model.DeleteExportTaskNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExportTaskNewResponse), nil
	}
}

type DeleteFullSqlExportTaskObsFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFullSqlExportTaskObsFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFullSqlExportTaskObsFileInvoker) Invoke() (*model.DeleteFullSqlExportTaskObsFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFullSqlExportTaskObsFileResponse), nil
	}
}

type EnableQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableQuotaInvoker) Invoke() (*model.EnableQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableQuotaResponse), nil
	}
}

type ExecuteFormatSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteFormatSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteFormatSqlInvoker) Invoke() (*model.ExecuteFormatSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteFormatSqlResponse), nil
	}
}

type ExecuteLoginConnectionNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteLoginConnectionNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteLoginConnectionNewInvoker) Invoke() (*model.ExecuteLoginConnectionNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteLoginConnectionNewResponse), nil
	}
}

type ExecuteSplitSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSplitSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteSplitSqlInvoker) Invoke() (*model.ExecuteSplitSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSplitSqlResponse), nil
	}
}

type ExecuteTestConnectionNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTestConnectionNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteTestConnectionNewInvoker) Invoke() (*model.ExecuteTestConnectionNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTestConnectionNewResponse), nil
	}
}

type ExecuteTuningInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTuningInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteTuningInvoker) Invoke() (*model.ExecuteTuningResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTuningResponse), nil
	}
}

type ExportInstanceListNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportInstanceListNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportInstanceListNewInvoker) Invoke() (*model.ExportInstanceListNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInstanceListNewResponse), nil
	}
}

type ImportExportObsObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportExportObsObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportExportObsObjectsInvoker) Invoke() (*model.ImportExportObsObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportExportObsObjectsResponse), nil
	}
}

type InvokeWdrReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *InvokeWdrReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InvokeWdrReportInvoker) Invoke() (*model.InvokeWdrReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvokeWdrReportResponse), nil
	}
}

type ListAllTypeInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTypeInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllTypeInstancesInvoker) Invoke() (*model.ListAllTypeInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTypeInstancesResponse), nil
	}
}

type ListBinlogExportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBinlogExportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBinlogExportsInvoker) Invoke() (*model.ListBinlogExportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBinlogExportsResponse), nil
	}
}

type ListBinlogFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBinlogFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBinlogFilesInvoker) Invoke() (*model.ListBinlogFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBinlogFilesResponse), nil
	}
}

type ListConnectionProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionProcessesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectionProcessesInvoker) Invoke() (*model.ListConnectionProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionProcessesResponse), nil
	}
}

type ListDatabaseObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseObjectsInvoker) Invoke() (*model.ListDatabaseObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseObjectsResponse), nil
	}
}

type ListDeadLockDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeadLockDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeadLockDatabasesInvoker) Invoke() (*model.ListDeadLockDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeadLockDatabasesResponse), nil
	}
}

type ListDeadLockDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeadLockDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeadLockDetailInvoker) Invoke() (*model.ListDeadLockDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeadLockDetailResponse), nil
	}
}

type ListFullDeadLocksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFullDeadLocksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFullDeadLocksInvoker) Invoke() (*model.ListFullDeadLocksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFullDeadLocksResponse), nil
	}
}

type ListFullSqlExportTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFullSqlExportTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFullSqlExportTasksInvoker) Invoke() (*model.ListFullSqlExportTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFullSqlExportTasksResponse), nil
	}
}

type ListInstanceHealthReportTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceHealthReportTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceHealthReportTasksInvoker) Invoke() (*model.ListInstanceHealthReportTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceHealthReportTasksResponse), nil
	}
}

type ListNotSetChargeModeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotSetChargeModeInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotSetChargeModeInstanceInvoker) Invoke() (*model.ListNotSetChargeModeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotSetChargeModeInstanceResponse), nil
	}
}

type ListSchemaNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSchemaNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSchemaNamesInvoker) Invoke() (*model.ListSchemaNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSchemaNamesResponse), nil
	}
}

type ListSharedConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharedConnectionsInvoker) Invoke() (*model.ListSharedConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedConnectionsResponse), nil
	}
}

type ListSmnTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmnTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmnTopicsInvoker) Invoke() (*model.ListSmnTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmnTopicsResponse), nil
	}
}

type ListSnapshots4ApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshots4ApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshots4ApiInvoker) Invoke() (*model.ListSnapshots4ApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshots4ApiResponse), nil
	}
}

type ListSqlLimitUserInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlLimitUserInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlLimitUserInstanceInvoker) Invoke() (*model.ListSqlLimitUserInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlLimitUserInstanceResponse), nil
	}
}

type ListSqlTemplateComparisonsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlTemplateComparisonsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlTemplateComparisonsInvoker) Invoke() (*model.ListSqlTemplateComparisonsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlTemplateComparisonsResponse), nil
	}
}

type ListSqlTemplateDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlTemplateDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlTemplateDatabasesInvoker) Invoke() (*model.ListSqlTemplateDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlTemplateDatabasesResponse), nil
	}
}

type ListSqlTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSqlTemplatesInvoker) Invoke() (*model.ListSqlTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlTemplatesResponse), nil
	}
}

type ListTasksByBatchIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksByBatchIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksByBatchIdInvoker) Invoke() (*model.ListTasksByBatchIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksByBatchIdResponse), nil
	}
}

type ListTasksBySqlTemplateIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksBySqlTemplateIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksBySqlTemplateIdInvoker) Invoke() (*model.ListTasksBySqlTemplateIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksBySqlTemplateIdResponse), nil
	}
}

type ListTasksByTaskIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksByTaskIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksByTaskIdInvoker) Invoke() (*model.ListTasksByTaskIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksByTaskIdResponse), nil
	}
}

type ListTemplateDatabaseComparisonsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateDatabaseComparisonsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateDatabaseComparisonsInvoker) Invoke() (*model.ListTemplateDatabaseComparisonsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateDatabaseComparisonsResponse), nil
	}
}

type ListUserInstanceListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserInstanceListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserInstanceListInvoker) Invoke() (*model.ListUserInstanceListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserInstanceListResponse), nil
	}
}

type RetryBinlogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryBinlogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryBinlogTaskInvoker) Invoke() (*model.RetryBinlogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryBinlogTaskResponse), nil
	}
}

type SearchBinlogParseInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchBinlogParseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchBinlogParseInvoker) Invoke() (*model.SearchBinlogParseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchBinlogParseResponse), nil
	}
}

type SearchErrorInfo4ApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchErrorInfo4ApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchErrorInfo4ApiInvoker) Invoke() (*model.SearchErrorInfo4ApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchErrorInfo4ApiResponse), nil
	}
}

type SearchErrorInfoSource4ApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchErrorInfoSource4ApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchErrorInfoSource4ApiInvoker) Invoke() (*model.SearchErrorInfoSource4ApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchErrorInfoSource4ApiResponse), nil
	}
}

type SearchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchNewInvoker) Invoke() (*model.SearchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchNewResponse), nil
	}
}

type SetSqlSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSqlSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSqlSwitchNewInvoker) Invoke() (*model.SetSqlSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSqlSwitchNewResponse), nil
	}
}

type ShowBinlogExportTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBinlogExportTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBinlogExportTaskInfoInvoker) Invoke() (*model.ShowBinlogExportTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBinlogExportTaskInfoResponse), nil
	}
}

type ShowBinlogParseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBinlogParseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBinlogParseInvoker) Invoke() (*model.ShowBinlogParseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBinlogParseResponse), nil
	}
}

type ShowBinlogTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBinlogTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBinlogTaskInfoInvoker) Invoke() (*model.ShowBinlogTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBinlogTaskInfoResponse), nil
	}
}

type ShowDdsConnectionStatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDdsConnectionStatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDdsConnectionStatInvoker) Invoke() (*model.ShowDdsConnectionStatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDdsConnectionStatResponse), nil
	}
}

type ShowDeadLockOriginDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockOriginDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockOriginDataInvoker) Invoke() (*model.ShowDeadLockOriginDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockOriginDataResponse), nil
	}
}

type ShowDeadLockRelationshipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockRelationshipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockRelationshipInvoker) Invoke() (*model.ShowDeadLockRelationshipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockRelationshipResponse), nil
	}
}

type ShowDeadLockStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockStatisticsInvoker) Invoke() (*model.ShowDeadLockStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockStatisticsResponse), nil
	}
}

type ShowDeadLockTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockTrendInvoker) Invoke() (*model.ShowDeadLockTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockTrendResponse), nil
	}
}

type ShowExecuteResultWithoutKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecuteResultWithoutKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExecuteResultWithoutKeyInvoker) Invoke() (*model.ShowExecuteResultWithoutKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecuteResultWithoutKeyResponse), nil
	}
}

type ShowExecuteResultWithoutKeyNoRetryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecuteResultWithoutKeyNoRetryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExecuteResultWithoutKeyNoRetryInvoker) Invoke() (*model.ShowExecuteResultWithoutKeyNoRetryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecuteResultWithoutKeyNoRetryResponse), nil
	}
}

type ShowExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecutionPlanInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExecutionPlanInvoker) Invoke() (*model.ShowExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecutionPlanResponse), nil
	}
}

type ShowExecutionTimeTemplateTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecutionTimeTemplateTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExecutionTimeTemplateTrendInvoker) Invoke() (*model.ShowExecutionTimeTemplateTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecutionTimeTemplateTrendResponse), nil
	}
}

type ShowFragmentSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFragmentSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFragmentSwitchInvoker) Invoke() (*model.ShowFragmentSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFragmentSwitchResponse), nil
	}
}

type ShowInstanceHealthReport4ApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceHealthReport4ApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceHealthReport4ApiInvoker) Invoke() (*model.ShowInstanceHealthReport4ApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceHealthReport4ApiResponse), nil
	}
}

type ShowInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInfoInvoker) Invoke() (*model.ShowInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceInfoResponse), nil
	}
}

type ShowInstanceLogUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceLogUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceLogUsageInvoker) Invoke() (*model.ShowInstanceLogUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceLogUsageResponse), nil
	}
}

type ShowInstanceMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceMetricInvoker) Invoke() (*model.ShowInstanceMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceMetricResponse), nil
	}
}

type ShowInstanceNodesInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceNodesInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceNodesInfoInvoker) Invoke() (*model.ShowInstanceNodesInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceNodesInfoResponse), nil
	}
}

type ShowIsSignedProtocolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIsSignedProtocolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIsSignedProtocolInvoker) Invoke() (*model.ShowIsSignedProtocolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIsSignedProtocolResponse), nil
	}
}

type ShowKillProcessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKillProcessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKillProcessTaskInvoker) Invoke() (*model.ShowKillProcessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKillProcessTaskResponse), nil
	}
}

type ShowLatestDeadLockSnapshot4ApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestDeadLockSnapshot4ApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestDeadLockSnapshot4ApiInvoker) Invoke() (*model.ShowLatestDeadLockSnapshot4ApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestDeadLockSnapshot4ApiResponse), nil
	}
}

type ShowMetaLockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetaLockInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetaLockInvoker) Invoke() (*model.ShowMetaLockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetaLockResponse), nil
	}
}

type ShowMetaLockSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetaLockSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetaLockSnapshotInvoker) Invoke() (*model.ShowMetaLockSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetaLockSnapshotResponse), nil
	}
}

type ShowOpeningInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOpeningInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOpeningInfoInvoker) Invoke() (*model.ShowOpeningInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOpeningInfoResponse), nil
	}
}

type ShowSingleTemplateTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleTemplateTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSingleTemplateTrendInvoker) Invoke() (*model.ShowSingleTemplateTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleTemplateTrendResponse), nil
	}
}

type ShowSqlTemplateTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlTemplateTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlTemplateTrendInvoker) Invoke() (*model.ShowSqlTemplateTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlTemplateTrendResponse), nil
	}
}

type ShowSupportKeyStringInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSupportKeyStringInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSupportKeyStringInvoker) Invoke() (*model.ShowSupportKeyStringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSupportKeyStringResponse), nil
	}
}

type ShowTuningResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTuningResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTuningResultInvoker) Invoke() (*model.ShowTuningResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTuningResultResponse), nil
	}
}

type ShowWaitingLocksSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWaitingLocksSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWaitingLocksSnapshotInvoker) Invoke() (*model.ShowWaitingLocksSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWaitingLocksSnapshotResponse), nil
	}
}

type ShowWdrSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWdrSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWdrSnapshotInvoker) Invoke() (*model.ShowWdrSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWdrSnapshotResponse), nil
	}
}

type SignProtocolNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignProtocolNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SignProtocolNewInvoker) Invoke() (*model.SignProtocolNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignProtocolNewResponse), nil
	}
}

type StopBinlogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBinlogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopBinlogTaskInvoker) Invoke() (*model.StopBinlogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBinlogTaskResponse), nil
	}
}

type SubscribeInstanceReportNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeInstanceReportNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeInstanceReportNewInvoker) Invoke() (*model.SubscribeInstanceReportNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeInstanceReportNewResponse), nil
	}
}

type SynchronizeInstanceListNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SynchronizeInstanceListNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SynchronizeInstanceListNewInvoker) Invoke() (*model.SynchronizeInstanceListNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SynchronizeInstanceListNewResponse), nil
	}
}

type UnsubscribeInstanceReportNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnsubscribeInstanceReportNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnsubscribeInstanceReportNewInvoker) Invoke() (*model.UnsubscribeInstanceReportNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribeInstanceReportNewResponse), nil
	}
}

type UpdateInstanceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceConfigInvoker) Invoke() (*model.UpdateInstanceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConfigResponse), nil
	}
}

type UpdateSearchPathFlagInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSearchPathFlagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSearchPathFlagInvoker) Invoke() (*model.UpdateSearchPathFlagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSearchPathFlagResponse), nil
	}
}

type UpdateSharedInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSharedInfoNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSharedInfoNewInvoker) Invoke() (*model.UpdateSharedInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSharedInfoNewResponse), nil
	}
}

type VerifyConnectionNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *VerifyConnectionNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *VerifyConnectionNewInvoker) Invoke() (*model.VerifyConnectionNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerifyConnectionNewResponse), nil
	}
}

type AddEmailTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddEmailTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddEmailTemplateInvoker) Invoke() (*model.AddEmailTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddEmailTemplateResponse), nil
	}
}

type AddFullSqlTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddFullSqlTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddFullSqlTaskInvoker) Invoke() (*model.AddFullSqlTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddFullSqlTaskResponse), nil
	}
}

type AddInstanceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddInstanceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddInstanceGroupInvoker) Invoke() (*model.AddInstanceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddInstanceGroupResponse), nil
	}
}

type AddInstanceToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddInstanceToGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddInstanceToGroupInvoker) Invoke() (*model.AddInstanceToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddInstanceToGroupResponse), nil
	}
}

type AddSqlLimitingRecordNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSqlLimitingRecordNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSqlLimitingRecordNewInvoker) Invoke() (*model.AddSqlLimitingRecordNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSqlLimitingRecordNewResponse), nil
	}
}

type BatchDeleteConnectionNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteConnectionNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteConnectionNewInvoker) Invoke() (*model.BatchDeleteConnectionNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteConnectionNewResponse), nil
	}
}

type BatchSendEmailInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSendEmailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSendEmailInvoker) Invoke() (*model.BatchSendEmailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSendEmailResponse), nil
	}
}

type BatchSubscribeReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSubscribeReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSubscribeReportInvoker) Invoke() (*model.BatchSubscribeReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSubscribeReportResponse), nil
	}
}

type CancelShareNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelShareNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelShareNewInvoker) Invoke() (*model.CancelShareNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelShareNewResponse), nil
	}
}

type ChangeChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeChargeModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeChargeModeInvoker) Invoke() (*model.ChangeChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeChargeModeResponse), nil
	}
}

type ChangeDeadLockSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDeadLockSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDeadLockSwitchNewInvoker) Invoke() (*model.ChangeDeadLockSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDeadLockSwitchNewResponse), nil
	}
}

type ChangeFullDeadLockSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeFullDeadLockSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeFullDeadLockSwitchInvoker) Invoke() (*model.ChangeFullDeadLockSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeFullDeadLockSwitchResponse), nil
	}
}

type ChangePaymentModeNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePaymentModeNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePaymentModeNewInvoker) Invoke() (*model.ChangePaymentModeNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePaymentModeNewResponse), nil
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

type CheckCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckCredentialInvoker) Invoke() (*model.CheckCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckCredentialResponse), nil
	}
}

type CheckCredentialForBatchInspectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckCredentialForBatchInspectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckCredentialForBatchInspectionInvoker) Invoke() (*model.CheckCredentialForBatchInspectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckCredentialForBatchInspectionResponse), nil
	}
}

type CreateHealthReportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHealthReportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHealthReportTaskInvoker) Invoke() (*model.CreateHealthReportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHealthReportTaskResponse), nil
	}
}

type CreateHistoryTransactionExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHistoryTransactionExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHistoryTransactionExportTaskInvoker) Invoke() (*model.CreateHistoryTransactionExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHistoryTransactionExportTaskResponse), nil
	}
}

type CreateSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSnapshotsInvoker) Invoke() (*model.CreateSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotsResponse), nil
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

type DeleteEmailTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEmailTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEmailTemplateInvoker) Invoke() (*model.DeleteEmailTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEmailTemplateResponse), nil
	}
}

type DeleteHistoryTransactionExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHistoryTransactionExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHistoryTransactionExportTaskInvoker) Invoke() (*model.DeleteHistoryTransactionExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHistoryTransactionExportTaskResponse), nil
	}
}

type DeleteInstanceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceGroupInvoker) Invoke() (*model.DeleteInstanceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceGroupResponse), nil
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

type ExportFullSqlDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportFullSqlDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportFullSqlDetailsInvoker) Invoke() (*model.ExportFullSqlDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportFullSqlDetailsResponse), nil
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

type ListAutoIncrementUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoIncrementUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoIncrementUsageInvoker) Invoke() (*model.ListAutoIncrementUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoIncrementUsageResponse), nil
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

type ListDbNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDbNamesInvoker) Invoke() (*model.ListDbNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbNamesResponse), nil
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

type ListEmailRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEmailRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEmailRecordInvoker) Invoke() (*model.ListEmailRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEmailRecordResponse), nil
	}
}

type ListEmailTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEmailTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEmailTemplateInvoker) Invoke() (*model.ListEmailTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEmailTemplateResponse), nil
	}
}

type ListFullSqlTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFullSqlTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFullSqlTasksInvoker) Invoke() (*model.ListFullSqlTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFullSqlTasksResponse), nil
	}
}

type ListHealthReportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHealthReportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHealthReportTaskInvoker) Invoke() (*model.ListHealthReportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHealthReportTaskResponse), nil
	}
}

type ListHistoryTransactionExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryTransactionExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryTransactionExportTaskInvoker) Invoke() (*model.ListHistoryTransactionExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryTransactionExportTaskResponse), nil
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

type ListInspectionReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInspectionReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInspectionReportInvoker) Invoke() (*model.ListInspectionReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInspectionReportResponse), nil
	}
}

type ListInstanceDistributionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceDistributionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceDistributionInvoker) Invoke() (*model.ListInstanceDistributionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceDistributionResponse), nil
	}
}

type ListInstanceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceGroupInvoker) Invoke() (*model.ListInstanceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceGroupResponse), nil
	}
}

type ListInstanceMultiNodesSingleMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceMultiNodesSingleMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceMultiNodesSingleMetricInvoker) Invoke() (*model.ListInstanceMultiNodesSingleMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceMultiNodesSingleMetricResponse), nil
	}
}

type ListInstanceNodesInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceNodesInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceNodesInfoInvoker) Invoke() (*model.ListInstanceNodesInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceNodesInfoResponse), nil
	}
}

type ListInstanceTopSlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTopSlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceTopSlowLogInvoker) Invoke() (*model.ListInstanceTopSlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTopSlowLogResponse), nil
	}
}

type ListLockBlockingDbInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLockBlockingDbInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLockBlockingDbInvoker) Invoke() (*model.ListLockBlockingDbResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLockBlockingDbResponse), nil
	}
}

type ListLockBlockingDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLockBlockingDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLockBlockingDetailInvoker) Invoke() (*model.ListLockBlockingDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLockBlockingDetailResponse), nil
	}
}

type ListLockBlockingRelationshipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLockBlockingRelationshipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLockBlockingRelationshipInvoker) Invoke() (*model.ListLockBlockingRelationshipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLockBlockingRelationshipResponse), nil
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

type ListRiskItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRiskItemsInvoker) Invoke() (*model.ListRiskItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskItemsResponse), nil
	}
}

type ListRiskTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRiskTrendInvoker) Invoke() (*model.ListRiskTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskTrendResponse), nil
	}
}

type ListSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotsInvoker) Invoke() (*model.ListSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotsResponse), nil
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

type ListTopSlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopSlowLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopSlowLogInvoker) Invoke() (*model.ListTopSlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopSlowLogResponse), nil
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

type LoginBuiltInAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *LoginBuiltInAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LoginBuiltInAccountInvoker) Invoke() (*model.LoginBuiltInAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginBuiltInAccountResponse), nil
	}
}

type LogoffBuiltInAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *LogoffBuiltInAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LogoffBuiltInAccountInvoker) Invoke() (*model.LogoffBuiltInAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogoffBuiltInAccountResponse), nil
	}
}

type ParseDeadLockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseDeadLockInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ParseDeadLockInvoker) Invoke() (*model.ParseDeadLockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseDeadLockResponse), nil
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

type SaveCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveCredentialInvoker) Invoke() (*model.SaveCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveCredentialResponse), nil
	}
}

type SaveCredentialForBatchInspectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveCredentialForBatchInspectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveCredentialForBatchInspectionInvoker) Invoke() (*model.SaveCredentialForBatchInspectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveCredentialForBatchInspectionResponse), nil
	}
}

type SetLockBlockingSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetLockBlockingSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetLockBlockingSwitchInvoker) Invoke() (*model.SetLockBlockingSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetLockBlockingSwitchResponse), nil
	}
}

type SetThresholdForMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetThresholdForMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetThresholdForMetricInvoker) Invoke() (*model.SetThresholdForMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetThresholdForMetricResponse), nil
	}
}

type ShowAnalysisSessionResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAnalysisSessionResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAnalysisSessionResultInvoker) Invoke() (*model.ShowAnalysisSessionResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAnalysisSessionResultResponse), nil
	}
}

type ShowAnalysisSessionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAnalysisSessionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAnalysisSessionStatusInvoker) Invoke() (*model.ShowAnalysisSessionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAnalysisSessionStatusResponse), nil
	}
}

type ShowClouddbaGetSearchPathFlagNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClouddbaGetSearchPathFlagNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClouddbaGetSearchPathFlagNewInvoker) Invoke() (*model.ShowClouddbaGetSearchPathFlagNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClouddbaGetSearchPathFlagNewResponse), nil
	}
}

type ShowCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCredentialInvoker) Invoke() (*model.ShowCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCredentialResponse), nil
	}
}

type ShowDasCloudDbaPriceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDasCloudDbaPriceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDasCloudDbaPriceInvoker) Invoke() (*model.ShowDasCloudDbaPriceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDasCloudDbaPriceResponse), nil
	}
}

type ShowDasRecommendSqlLimitRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDasRecommendSqlLimitRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDasRecommendSqlLimitRuleInvoker) Invoke() (*model.ShowDasRecommendSqlLimitRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDasRecommendSqlLimitRuleResponse), nil
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

type ShowDeadLockAnalysisResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockAnalysisResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockAnalysisResultInvoker) Invoke() (*model.ShowDeadLockAnalysisResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockAnalysisResultResponse), nil
	}
}

type ShowDeadLockTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockTopologyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockTopologyInvoker) Invoke() (*model.ShowDeadLockTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockTopologyResponse), nil
	}
}

type ShowExportTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExportTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExportTaskInfoInvoker) Invoke() (*model.ShowExportTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExportTaskInfoResponse), nil
	}
}

type ShowFullDeadLockListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFullDeadLockListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFullDeadLockListInvoker) Invoke() (*model.ShowFullDeadLockListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFullDeadLockListResponse), nil
	}
}

type ShowFullDeadLockSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFullDeadLockSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFullDeadLockSwitchInvoker) Invoke() (*model.ShowFullDeadLockSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFullDeadLockSwitchResponse), nil
	}
}

type ShowFullDeadLockSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFullDeadLockSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFullDeadLockSwitchNewInvoker) Invoke() (*model.ShowFullDeadLockSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFullDeadLockSwitchNewResponse), nil
	}
}

type ShowGlobalPrivacyNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalPrivacyNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalPrivacyNewInvoker) Invoke() (*model.ShowGlobalPrivacyNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalPrivacyNewResponse), nil
	}
}

type ShowHealthReportSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHealthReportSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHealthReportSettingsInvoker) Invoke() (*model.ShowHealthReportSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthReportSettingsResponse), nil
	}
}

type ShowHistoryTransactionExportTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryTransactionExportTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryTransactionExportTaskInfoInvoker) Invoke() (*model.ShowHistoryTransactionExportTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryTransactionExportTaskInfoResponse), nil
	}
}

type ShowHistoryTransactionSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryTransactionSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryTransactionSwitchNewInvoker) Invoke() (*model.ShowHistoryTransactionSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryTransactionSwitchNewResponse), nil
	}
}

type ShowIndexUsageSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIndexUsageSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIndexUsageSwitchNewInvoker) Invoke() (*model.ShowIndexUsageSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIndexUsageSwitchNewResponse), nil
	}
}

type ShowInstanceHealthReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceHealthReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceHealthReportInvoker) Invoke() (*model.ShowInstanceHealthReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceHealthReportResponse), nil
	}
}

type ShowKillProcessTaskSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKillProcessTaskSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKillProcessTaskSwitchInvoker) Invoke() (*model.ShowKillProcessTaskSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKillProcessTaskSwitchResponse), nil
	}
}

type ShowLatestDeadLockSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestDeadLockSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestDeadLockSnapshotInvoker) Invoke() (*model.ShowLatestDeadLockSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestDeadLockSnapshotResponse), nil
	}
}

type ShowLatestInstanceHealthReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestInstanceHealthReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestInstanceHealthReportInvoker) Invoke() (*model.ShowLatestInstanceHealthReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestInstanceHealthReportResponse), nil
	}
}

type ShowLockBlockingStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLockBlockingStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLockBlockingStatisticsInvoker) Invoke() (*model.ShowLockBlockingStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLockBlockingStatisticsResponse), nil
	}
}

type ShowLockBlockingSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLockBlockingSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLockBlockingSwitchInvoker) Invoke() (*model.ShowLockBlockingSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLockBlockingSwitchResponse), nil
	}
}

type ShowLockBlockingTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLockBlockingTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLockBlockingTrendInvoker) Invoke() (*model.ShowLockBlockingTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLockBlockingTrendResponse), nil
	}
}

type ShowLongHistoryTransactionSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLongHistoryTransactionSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLongHistoryTransactionSwitchNewInvoker) Invoke() (*model.ShowLongHistoryTransactionSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLongHistoryTransactionSwitchNewResponse), nil
	}
}

type ShowMetricNamesSupportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricNamesSupportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetricNamesSupportInvoker) Invoke() (*model.ShowMetricNamesSupportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricNamesSupportResponse), nil
	}
}

type ShowNameListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNameListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNameListInvoker) Invoke() (*model.ShowNameListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNameListResponse), nil
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

type ShowSlowLogSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowLogSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlowLogSwitchNewInvoker) Invoke() (*model.ShowSlowLogSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowLogSwitchNewResponse), nil
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

type ShowSqlLimitingSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlLimitingSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlLimitingSwitchNewInvoker) Invoke() (*model.ShowSqlLimitingSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlLimitingSwitchNewResponse), nil
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

type ShowSupportedEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSupportedEnginesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSupportedEnginesInvoker) Invoke() (*model.ShowSupportedEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSupportedEnginesResponse), nil
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

type ShowWhetherUseCloudDbaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWhetherUseCloudDbaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWhetherUseCloudDbaInvoker) Invoke() (*model.ShowWhetherUseCloudDbaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWhetherUseCloudDbaResponse), nil
	}
}

type StartAnalysisSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAnalysisSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAnalysisSessionInvoker) Invoke() (*model.StartAnalysisSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAnalysisSessionResponse), nil
	}
}

type SynchronizeInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SynchronizeInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SynchronizeInstancesInvoker) Invoke() (*model.SynchronizeInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SynchronizeInstancesResponse), nil
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

type UpdateEmailTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEmailTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEmailTemplateInvoker) Invoke() (*model.UpdateEmailTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEmailTemplateResponse), nil
	}
}

type UpdateFullSqlSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFullSqlSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFullSqlSwitchInvoker) Invoke() (*model.UpdateFullSqlSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFullSqlSwitchResponse), nil
	}
}

type UpdateHealthReportSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHealthReportSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHealthReportSettingsInvoker) Invoke() (*model.UpdateHealthReportSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHealthReportSettingsResponse), nil
	}
}

type UpdateInstanceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceGroupInvoker) Invoke() (*model.UpdateInstanceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceGroupResponse), nil
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

type ShowDeadLockSwitchNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeadLockSwitchNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeadLockSwitchNewInvoker) Invoke() (*model.ShowDeadLockSwitchNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeadLockSwitchNewResponse), nil
	}
}

type SwitchFullsqlSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchFullsqlSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchFullsqlSwitchInvoker) Invoke() (*model.SwitchFullsqlSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchFullsqlSwitchResponse), nil
	}
}
