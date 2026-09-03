package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/das/v3/model"
)

type DasClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDasClient(hcClient *httpclient.HcHttpClient) *DasClient {
	return &DasClient{HcClient: hcClient}
}

func DasClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CancelShareConnections 删除共享链接
//
// 删除共享链接，
// 用于用户删除共享链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CancelShareConnections(request *model.CancelShareConnectionsRequest) (*model.CancelShareConnectionsResponse, error) {
	requestDef := GenReqDefForCancelShareConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelShareConnectionsResponse), nil
	}
}

// CancelShareConnectionsInvoker 删除共享链接
func (c *DasClient) CancelShareConnectionsInvoker(request *model.CancelShareConnectionsRequest) *CancelShareConnectionsInvoker {
	requestDef := GenReqDefForCancelShareConnections()
	return &CancelShareConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceConnection 创建实例连接
//
// 创建实例连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateInstanceConnection(request *model.CreateInstanceConnectionRequest) (*model.CreateInstanceConnectionResponse, error) {
	requestDef := GenReqDefForCreateInstanceConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceConnectionResponse), nil
	}
}

// CreateInstanceConnectionInvoker 创建实例连接
func (c *DasClient) CreateInstanceConnectionInvoker(request *model.CreateInstanceConnectionRequest) *CreateInstanceConnectionInvoker {
	requestDef := GenReqDefForCreateInstanceConnection()
	return &CreateInstanceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShareConnections 设置共享链接
//
// 设置共享链接，
// 用于用户添加共享链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateShareConnections(request *model.CreateShareConnectionsRequest) (*model.CreateShareConnectionsResponse, error) {
	requestDef := GenReqDefForCreateShareConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareConnectionsResponse), nil
	}
}

// CreateShareConnectionsInvoker 设置共享链接
func (c *DasClient) CreateShareConnectionsInvoker(request *model.CreateShareConnectionsRequest) *CreateShareConnectionsInvoker {
	requestDef := GenReqDefForCreateShareConnections()
	return &CreateShareConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteExportTask 立即执行导出任务
//
// 立即执行导出任务，
// 用于用户立即执行导出任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteExportTask(request *model.ExecuteExportTaskRequest) (*model.ExecuteExportTaskResponse, error) {
	requestDef := GenReqDefForExecuteExportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteExportTaskResponse), nil
	}
}

// ExecuteExportTaskInvoker 立即执行导出任务
func (c *DasClient) ExecuteExportTaskInvoker(request *model.ExecuteExportTaskRequest) *ExecuteExportTaskInvoker {
	requestDef := GenReqDefForExecuteExportTask()
	return &ExecuteExportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteImportTask 立即执行导入任务
//
// 立即执行导入任务，
// 用于用户立即执行导入任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteImportTask(request *model.ExecuteImportTaskRequest) (*model.ExecuteImportTaskResponse, error) {
	requestDef := GenReqDefForExecuteImportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteImportTaskResponse), nil
	}
}

// ExecuteImportTaskInvoker 立即执行导入任务
func (c *DasClient) ExecuteImportTaskInvoker(request *model.ExecuteImportTaskRequest) *ExecuteImportTaskInvoker {
	requestDef := GenReqDefForExecuteImportTask()
	return &ExecuteImportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 查询实例连接列表
//
// 查询实例连接列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 查询实例连接列表
func (c *DasClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本列表
//
// 查询API版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本列表
func (c *DasClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询指定的API版本信息
//
// 查询指定的API版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询指定的API版本信息
func (c *DasClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTasksNew 创建全量SQL明细解析任务
//
// 创建全量SQL明细解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddTasksNew(request *model.AddTasksNewRequest) (*model.AddTasksNewResponse, error) {
	requestDef := GenReqDefForAddTasksNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTasksNewResponse), nil
	}
}

// AddTasksNewInvoker 创建全量SQL明细解析任务
func (c *DasClient) AddTasksNewInvoker(request *model.AddTasksNewRequest) *AddTasksNewInvoker {
	requestDef := GenReqDefForAddTasksNew()
	return &AddTasksNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddFullSqlTasks 批量创建全量SQL明细解析任务
//
// 批量创建全量SQL明细解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) BatchAddFullSqlTasks(request *model.BatchAddFullSqlTasksRequest) (*model.BatchAddFullSqlTasksResponse, error) {
	requestDef := GenReqDefForBatchAddFullSqlTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddFullSqlTasksResponse), nil
	}
}

// BatchAddFullSqlTasksInvoker 批量创建全量SQL明细解析任务
func (c *DasClient) BatchAddFullSqlTasksInvoker(request *model.BatchAddFullSqlTasksRequest) *BatchAddFullSqlTasksInvoker {
	requestDef := GenReqDefForBatchAddFullSqlTasks()
	return &BatchAddFullSqlTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetSqlSwitchNew 批量设置SQL开关
//
// 批量设置SQL开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) BatchSetSqlSwitchNew(request *model.BatchSetSqlSwitchNewRequest) (*model.BatchSetSqlSwitchNewResponse, error) {
	requestDef := GenReqDefForBatchSetSqlSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSqlSwitchNewResponse), nil
	}
}

// BatchSetSqlSwitchNewInvoker 批量设置SQL开关
func (c *DasClient) BatchSetSqlSwitchNewInvoker(request *model.BatchSetSqlSwitchNewRequest) *BatchSetSqlSwitchNewInvoker {
	requestDef := GenReqDefForBatchSetSqlSwitchNew()
	return &BatchSetSqlSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelConnectionProcess Kill进程
//
// Kill进程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CancelConnectionProcess(request *model.CancelConnectionProcessRequest) (*model.CancelConnectionProcessResponse, error) {
	requestDef := GenReqDefForCancelConnectionProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelConnectionProcessResponse), nil
	}
}

// CancelConnectionProcessInvoker Kill进程
func (c *DasClient) CancelConnectionProcessInvoker(request *model.CancelConnectionProcessRequest) *CancelConnectionProcessInvoker {
	requestDef := GenReqDefForCancelConnectionProcess()
	return &CancelConnectionProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeQuotaNew 修改配额
//
// 修改配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeQuotaNew(request *model.ChangeQuotaNewRequest) (*model.ChangeQuotaNewResponse, error) {
	requestDef := GenReqDefForChangeQuotaNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeQuotaNewResponse), nil
	}
}

// ChangeQuotaNewInvoker 修改配额
func (c *DasClient) ChangeQuotaNewInvoker(request *model.ChangeQuotaNewRequest) *ChangeQuotaNewInvoker {
	requestDef := GenReqDefForChangeQuotaNew()
	return &ChangeQuotaNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckHealthReportTask 检查是否有健康报告任务
//
// 检查是否有健康报告任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CheckHealthReportTask(request *model.CheckHealthReportTaskRequest) (*model.CheckHealthReportTaskResponse, error) {
	requestDef := GenReqDefForCheckHealthReportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckHealthReportTaskResponse), nil
	}
}

// CheckHealthReportTaskInvoker 检查是否有健康报告任务
func (c *DasClient) CheckHealthReportTaskInvoker(request *model.CheckHealthReportTaskRequest) *CheckHealthReportTaskInvoker {
	requestDef := GenReqDefForCheckHealthReportTask()
	return &CheckHealthReportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBinlogTask 创建binlog解析任务
//
// 创建binlog解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateBinlogTask(request *model.CreateBinlogTaskRequest) (*model.CreateBinlogTaskResponse, error) {
	requestDef := GenReqDefForCreateBinlogTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBinlogTaskResponse), nil
	}
}

// CreateBinlogTaskInvoker 创建binlog解析任务
func (c *DasClient) CreateBinlogTaskInvoker(request *model.CreateBinlogTaskRequest) *CreateBinlogTaskInvoker {
	requestDef := GenReqDefForCreateBinlogTask()
	return &CreateBinlogTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDbsConnection DBS连接
//
// DBS连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateDbsConnection(request *model.CreateDbsConnectionRequest) (*model.CreateDbsConnectionResponse, error) {
	requestDef := GenReqDefForCreateDbsConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbsConnectionResponse), nil
	}
}

// CreateDbsConnectionInvoker DBS连接
func (c *DasClient) CreateDbsConnectionInvoker(request *model.CreateDbsConnectionRequest) *CreateDbsConnectionInvoker {
	requestDef := GenReqDefForCreateDbsConnection()
	return &CreateDbsConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFullSqlBucket 创建全量SQL桶
//
// 创建全量SQL桶
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateFullSqlBucket(request *model.CreateFullSqlBucketRequest) (*model.CreateFullSqlBucketResponse, error) {
	requestDef := GenReqDefForCreateFullSqlBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFullSqlBucketResponse), nil
	}
}

// CreateFullSqlBucketInvoker 创建全量SQL桶
func (c *DasClient) CreateFullSqlBucketInvoker(request *model.CreateFullSqlBucketRequest) *CreateFullSqlBucketInvoker {
	requestDef := GenReqDefForCreateFullSqlBucket()
	return &CreateFullSqlBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIndexUsageExportTaskNew 创建索引使用导出任务
//
// 创建索引使用导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateIndexUsageExportTaskNew(request *model.CreateIndexUsageExportTaskNewRequest) (*model.CreateIndexUsageExportTaskNewResponse, error) {
	requestDef := GenReqDefForCreateIndexUsageExportTaskNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndexUsageExportTaskNewResponse), nil
	}
}

// CreateIndexUsageExportTaskNewInvoker 创建索引使用导出任务
func (c *DasClient) CreateIndexUsageExportTaskNewInvoker(request *model.CreateIndexUsageExportTaskNewRequest) *CreateIndexUsageExportTaskNewInvoker {
	requestDef := GenReqDefForCreateIndexUsageExportTaskNew()
	return &CreateIndexUsageExportTaskNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceHealthReportTaskNew 创建实例健康报告任务
//
// 创建实例健康报告任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateInstanceHealthReportTaskNew(request *model.CreateInstanceHealthReportTaskNewRequest) (*model.CreateInstanceHealthReportTaskNewResponse, error) {
	requestDef := GenReqDefForCreateInstanceHealthReportTaskNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceHealthReportTaskNewResponse), nil
	}
}

// CreateInstanceHealthReportTaskNewInvoker 创建实例健康报告任务
func (c *DasClient) CreateInstanceHealthReportTaskNewInvoker(request *model.CreateInstanceHealthReportTaskNewRequest) *CreateInstanceHealthReportTaskNewInvoker {
	requestDef := GenReqDefForCreateInstanceHealthReportTaskNew()
	return &CreateInstanceHealthReportTaskNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWdrReport 触发WDR
//
// 触发WDR
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateWdrReport(request *model.CreateWdrReportRequest) (*model.CreateWdrReportResponse, error) {
	requestDef := GenReqDefForCreateWdrReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWdrReportResponse), nil
	}
}

// CreateWdrReportInvoker 触发WDR
func (c *DasClient) CreateWdrReportInvoker(request *model.CreateWdrReportRequest) *CreateWdrReportInvoker {
	requestDef := GenReqDefForCreateWdrReport()
	return &CreateWdrReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBinlogTask 删除binlog任务
//
// 删除binlog任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteBinlogTask(request *model.DeleteBinlogTaskRequest) (*model.DeleteBinlogTaskResponse, error) {
	requestDef := GenReqDefForDeleteBinlogTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBinlogTaskResponse), nil
	}
}

// DeleteBinlogTaskInvoker 删除binlog任务
func (c *DasClient) DeleteBinlogTaskInvoker(request *model.DeleteBinlogTaskRequest) *DeleteBinlogTaskInvoker {
	requestDef := GenReqDefForDeleteBinlogTask()
	return &DeleteBinlogTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDbObjNew 删除数据库对象
//
// 删除数据库对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteDbObjNew(request *model.DeleteDbObjNewRequest) (*model.DeleteDbObjNewResponse, error) {
	requestDef := GenReqDefForDeleteDbObjNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbObjNewResponse), nil
	}
}

// DeleteDbObjNewInvoker 删除数据库对象
func (c *DasClient) DeleteDbObjNewInvoker(request *model.DeleteDbObjNewRequest) *DeleteDbObjNewInvoker {
	requestDef := GenReqDefForDeleteDbObjNew()
	return &DeleteDbObjNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteExportTaskNew 删除binlog导出任务
//
// 删除binlog导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteExportTaskNew(request *model.DeleteExportTaskNewRequest) (*model.DeleteExportTaskNewResponse, error) {
	requestDef := GenReqDefForDeleteExportTaskNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExportTaskNewResponse), nil
	}
}

// DeleteExportTaskNewInvoker 删除binlog导出任务
func (c *DasClient) DeleteExportTaskNewInvoker(request *model.DeleteExportTaskNewRequest) *DeleteExportTaskNewInvoker {
	requestDef := GenReqDefForDeleteExportTaskNew()
	return &DeleteExportTaskNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFullSqlExportTaskObsFile 删除全量SQL导出任务OBS文件
//
// 删除全量SQL导出任务OBS文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteFullSqlExportTaskObsFile(request *model.DeleteFullSqlExportTaskObsFileRequest) (*model.DeleteFullSqlExportTaskObsFileResponse, error) {
	requestDef := GenReqDefForDeleteFullSqlExportTaskObsFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFullSqlExportTaskObsFileResponse), nil
	}
}

// DeleteFullSqlExportTaskObsFileInvoker 删除全量SQL导出任务OBS文件
func (c *DasClient) DeleteFullSqlExportTaskObsFileInvoker(request *model.DeleteFullSqlExportTaskObsFileRequest) *DeleteFullSqlExportTaskObsFileInvoker {
	requestDef := GenReqDefForDeleteFullSqlExportTaskObsFile()
	return &DeleteFullSqlExportTaskObsFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableQuota 开通配额
//
// 开通配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) EnableQuota(request *model.EnableQuotaRequest) (*model.EnableQuotaResponse, error) {
	requestDef := GenReqDefForEnableQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableQuotaResponse), nil
	}
}

// EnableQuotaInvoker 开通配额
func (c *DasClient) EnableQuotaInvoker(request *model.EnableQuotaRequest) *EnableQuotaInvoker {
	requestDef := GenReqDefForEnableQuota()
	return &EnableQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteFormatSql 格式化SQL
//
// 格式化SQL
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteFormatSql(request *model.ExecuteFormatSqlRequest) (*model.ExecuteFormatSqlResponse, error) {
	requestDef := GenReqDefForExecuteFormatSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteFormatSqlResponse), nil
	}
}

// ExecuteFormatSqlInvoker 格式化SQL
func (c *DasClient) ExecuteFormatSqlInvoker(request *model.ExecuteFormatSqlRequest) *ExecuteFormatSqlInvoker {
	requestDef := GenReqDefForExecuteFormatSql()
	return &ExecuteFormatSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteLoginConnectionNew 登录操作
//
// 登录操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteLoginConnectionNew(request *model.ExecuteLoginConnectionNewRequest) (*model.ExecuteLoginConnectionNewResponse, error) {
	requestDef := GenReqDefForExecuteLoginConnectionNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteLoginConnectionNewResponse), nil
	}
}

// ExecuteLoginConnectionNewInvoker 登录操作
func (c *DasClient) ExecuteLoginConnectionNewInvoker(request *model.ExecuteLoginConnectionNewRequest) *ExecuteLoginConnectionNewInvoker {
	requestDef := GenReqDefForExecuteLoginConnectionNew()
	return &ExecuteLoginConnectionNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteSplitSql 拆分SQL
//
// 拆分SQL
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteSplitSql(request *model.ExecuteSplitSqlRequest) (*model.ExecuteSplitSqlResponse, error) {
	requestDef := GenReqDefForExecuteSplitSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSplitSqlResponse), nil
	}
}

// ExecuteSplitSqlInvoker 拆分SQL
func (c *DasClient) ExecuteSplitSqlInvoker(request *model.ExecuteSplitSqlRequest) *ExecuteSplitSqlInvoker {
	requestDef := GenReqDefForExecuteSplitSql()
	return &ExecuteSplitSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTestConnectionNew 测试数据库实例连接
//
// 测试数据库实例连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteTestConnectionNew(request *model.ExecuteTestConnectionNewRequest) (*model.ExecuteTestConnectionNewResponse, error) {
	requestDef := GenReqDefForExecuteTestConnectionNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTestConnectionNewResponse), nil
	}
}

// ExecuteTestConnectionNewInvoker 测试数据库实例连接
func (c *DasClient) ExecuteTestConnectionNewInvoker(request *model.ExecuteTestConnectionNewRequest) *ExecuteTestConnectionNewInvoker {
	requestDef := GenReqDefForExecuteTestConnectionNew()
	return &ExecuteTestConnectionNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTuning 执行调优
//
// 执行调优
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExecuteTuning(request *model.ExecuteTuningRequest) (*model.ExecuteTuningResponse, error) {
	requestDef := GenReqDefForExecuteTuning()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTuningResponse), nil
	}
}

// ExecuteTuningInvoker 执行调优
func (c *DasClient) ExecuteTuningInvoker(request *model.ExecuteTuningRequest) *ExecuteTuningInvoker {
	requestDef := GenReqDefForExecuteTuning()
	return &ExecuteTuningInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportInstanceListNew 导出实例列表
//
// 导出实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportInstanceListNew(request *model.ExportInstanceListNewRequest) (*model.ExportInstanceListNewResponse, error) {
	requestDef := GenReqDefForExportInstanceListNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportInstanceListNewResponse), nil
	}
}

// ExportInstanceListNewInvoker 导出实例列表
func (c *DasClient) ExportInstanceListNewInvoker(request *model.ExportInstanceListNewRequest) *ExportInstanceListNewInvoker {
	requestDef := GenReqDefForExportInstanceListNew()
	return &ExportInstanceListNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportExportObsObjects 获取OBS对象列表
//
// 获取OBS对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ImportExportObsObjects(request *model.ImportExportObsObjectsRequest) (*model.ImportExportObsObjectsResponse, error) {
	requestDef := GenReqDefForImportExportObsObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportExportObsObjectsResponse), nil
	}
}

// ImportExportObsObjectsInvoker 获取OBS对象列表
func (c *DasClient) ImportExportObsObjectsInvoker(request *model.ImportExportObsObjectsRequest) *ImportExportObsObjectsInvoker {
	requestDef := GenReqDefForImportExportObsObjects()
	return &ImportExportObsObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InvokeWdrReport 获取WDR数据
//
// 获取WDR数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) InvokeWdrReport(request *model.InvokeWdrReportRequest) (*model.InvokeWdrReportResponse, error) {
	requestDef := GenReqDefForInvokeWdrReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeWdrReportResponse), nil
	}
}

// InvokeWdrReportInvoker 获取WDR数据
func (c *DasClient) InvokeWdrReportInvoker(request *model.InvokeWdrReportRequest) *InvokeWdrReportInvoker {
	requestDef := GenReqDefForInvokeWdrReport()
	return &InvokeWdrReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTypeInstances 查询所有类型实例列表
//
// 查询所有类型实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListAllTypeInstances(request *model.ListAllTypeInstancesRequest) (*model.ListAllTypeInstancesResponse, error) {
	requestDef := GenReqDefForListAllTypeInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTypeInstancesResponse), nil
	}
}

// ListAllTypeInstancesInvoker 查询所有类型实例列表
func (c *DasClient) ListAllTypeInstancesInvoker(request *model.ListAllTypeInstancesRequest) *ListAllTypeInstancesInvoker {
	requestDef := GenReqDefForListAllTypeInstances()
	return &ListAllTypeInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBinlogExports 导出binlog任务列表
//
// 导出binlog任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListBinlogExports(request *model.ListBinlogExportsRequest) (*model.ListBinlogExportsResponse, error) {
	requestDef := GenReqDefForListBinlogExports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBinlogExportsResponse), nil
	}
}

// ListBinlogExportsInvoker 导出binlog任务列表
func (c *DasClient) ListBinlogExportsInvoker(request *model.ListBinlogExportsRequest) *ListBinlogExportsInvoker {
	requestDef := GenReqDefForListBinlogExports()
	return &ListBinlogExportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBinlogFiles 查询binlog文件列表
//
// 查询binlog文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListBinlogFiles(request *model.ListBinlogFilesRequest) (*model.ListBinlogFilesResponse, error) {
	requestDef := GenReqDefForListBinlogFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBinlogFilesResponse), nil
	}
}

// ListBinlogFilesInvoker 查询binlog文件列表
func (c *DasClient) ListBinlogFilesInvoker(request *model.ListBinlogFilesRequest) *ListBinlogFilesInvoker {
	requestDef := GenReqDefForListBinlogFiles()
	return &ListBinlogFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnectionProcesses 查询实例会话
//
// 查询实例会话
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListConnectionProcesses(request *model.ListConnectionProcessesRequest) (*model.ListConnectionProcessesResponse, error) {
	requestDef := GenReqDefForListConnectionProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionProcessesResponse), nil
	}
}

// ListConnectionProcessesInvoker 查询实例会话
func (c *DasClient) ListConnectionProcessesInvoker(request *model.ListConnectionProcessesRequest) *ListConnectionProcessesInvoker {
	requestDef := GenReqDefForListConnectionProcesses()
	return &ListConnectionProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseObjects 查询数据库对象列表
//
// 查询数据库对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListDatabaseObjects(request *model.ListDatabaseObjectsRequest) (*model.ListDatabaseObjectsResponse, error) {
	requestDef := GenReqDefForListDatabaseObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseObjectsResponse), nil
	}
}

// ListDatabaseObjectsInvoker 查询数据库对象列表
func (c *DasClient) ListDatabaseObjectsInvoker(request *model.ListDatabaseObjectsRequest) *ListDatabaseObjectsInvoker {
	requestDef := GenReqDefForListDatabaseObjects()
	return &ListDatabaseObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeadLockDatabases 获取死锁数据库列表
//
// 获取死锁数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListDeadLockDatabases(request *model.ListDeadLockDatabasesRequest) (*model.ListDeadLockDatabasesResponse, error) {
	requestDef := GenReqDefForListDeadLockDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeadLockDatabasesResponse), nil
	}
}

// ListDeadLockDatabasesInvoker 获取死锁数据库列表
func (c *DasClient) ListDeadLockDatabasesInvoker(request *model.ListDeadLockDatabasesRequest) *ListDeadLockDatabasesInvoker {
	requestDef := GenReqDefForListDeadLockDatabases()
	return &ListDeadLockDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeadLockDetail 获取死锁详情列表
//
// 获取死锁详情列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListDeadLockDetail(request *model.ListDeadLockDetailRequest) (*model.ListDeadLockDetailResponse, error) {
	requestDef := GenReqDefForListDeadLockDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeadLockDetailResponse), nil
	}
}

// ListDeadLockDetailInvoker 获取死锁详情列表
func (c *DasClient) ListDeadLockDetailInvoker(request *model.ListDeadLockDetailRequest) *ListDeadLockDetailInvoker {
	requestDef := GenReqDefForListDeadLockDetail()
	return &ListDeadLockDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFullDeadLocks 获取完整死锁列表
//
// 获取完整死锁列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListFullDeadLocks(request *model.ListFullDeadLocksRequest) (*model.ListFullDeadLocksResponse, error) {
	requestDef := GenReqDefForListFullDeadLocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFullDeadLocksResponse), nil
	}
}

// ListFullDeadLocksInvoker 获取完整死锁列表
func (c *DasClient) ListFullDeadLocksInvoker(request *model.ListFullDeadLocksRequest) *ListFullDeadLocksInvoker {
	requestDef := GenReqDefForListFullDeadLocks()
	return &ListFullDeadLocksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFullSqlExportTasks 获取全量SQL导出任务列表
//
// 获取全量SQL导出任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListFullSqlExportTasks(request *model.ListFullSqlExportTasksRequest) (*model.ListFullSqlExportTasksResponse, error) {
	requestDef := GenReqDefForListFullSqlExportTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFullSqlExportTasksResponse), nil
	}
}

// ListFullSqlExportTasksInvoker 获取全量SQL导出任务列表
func (c *DasClient) ListFullSqlExportTasksInvoker(request *model.ListFullSqlExportTasksRequest) *ListFullSqlExportTasksInvoker {
	requestDef := GenReqDefForListFullSqlExportTasks()
	return &ListFullSqlExportTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceHealthReportTasks 获取实例健康报告任务列表
//
// 获取实例健康报告任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceHealthReportTasks(request *model.ListInstanceHealthReportTasksRequest) (*model.ListInstanceHealthReportTasksResponse, error) {
	requestDef := GenReqDefForListInstanceHealthReportTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceHealthReportTasksResponse), nil
	}
}

// ListInstanceHealthReportTasksInvoker 获取实例健康报告任务列表
func (c *DasClient) ListInstanceHealthReportTasksInvoker(request *model.ListInstanceHealthReportTasksRequest) *ListInstanceHealthReportTasksInvoker {
	requestDef := GenReqDefForListInstanceHealthReportTasks()
	return &ListInstanceHealthReportTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotSetChargeModeInstance 获取未设置付费的实例列表
//
// 获取未设置付费的实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListNotSetChargeModeInstance(request *model.ListNotSetChargeModeInstanceRequest) (*model.ListNotSetChargeModeInstanceResponse, error) {
	requestDef := GenReqDefForListNotSetChargeModeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotSetChargeModeInstanceResponse), nil
	}
}

// ListNotSetChargeModeInstanceInvoker 获取未设置付费的实例列表
func (c *DasClient) ListNotSetChargeModeInstanceInvoker(request *model.ListNotSetChargeModeInstanceRequest) *ListNotSetChargeModeInstanceInvoker {
	requestDef := GenReqDefForListNotSetChargeModeInstance()
	return &ListNotSetChargeModeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchemaNames 获取schema名称列表
//
// 获取schema名称列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSchemaNames(request *model.ListSchemaNamesRequest) (*model.ListSchemaNamesResponse, error) {
	requestDef := GenReqDefForListSchemaNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemaNamesResponse), nil
	}
}

// ListSchemaNamesInvoker 获取schema名称列表
func (c *DasClient) ListSchemaNamesInvoker(request *model.ListSchemaNamesRequest) *ListSchemaNamesInvoker {
	requestDef := GenReqDefForListSchemaNames()
	return &ListSchemaNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharedConnections 查询共享列表
//
// 查询共享列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSharedConnections(request *model.ListSharedConnectionsRequest) (*model.ListSharedConnectionsResponse, error) {
	requestDef := GenReqDefForListSharedConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedConnectionsResponse), nil
	}
}

// ListSharedConnectionsInvoker 查询共享列表
func (c *DasClient) ListSharedConnectionsInvoker(request *model.ListSharedConnectionsRequest) *ListSharedConnectionsInvoker {
	requestDef := GenReqDefForListSharedConnections()
	return &ListSharedConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmnTopics 获取SMN主题列表
//
// 获取SMN主题列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSmnTopics(request *model.ListSmnTopicsRequest) (*model.ListSmnTopicsResponse, error) {
	requestDef := GenReqDefForListSmnTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmnTopicsResponse), nil
	}
}

// ListSmnTopicsInvoker 获取SMN主题列表
func (c *DasClient) ListSmnTopicsInvoker(request *model.ListSmnTopicsRequest) *ListSmnTopicsInvoker {
	requestDef := GenReqDefForListSmnTopics()
	return &ListSmnTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshots4Api 查询快照
//
// 查询快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSnapshots4Api(request *model.ListSnapshots4ApiRequest) (*model.ListSnapshots4ApiResponse, error) {
	requestDef := GenReqDefForListSnapshots4Api()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshots4ApiResponse), nil
	}
}

// ListSnapshots4ApiInvoker 查询快照
func (c *DasClient) ListSnapshots4ApiInvoker(request *model.ListSnapshots4ApiRequest) *ListSnapshots4ApiInvoker {
	requestDef := GenReqDefForListSnapshots4Api()
	return &ListSnapshots4ApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlLimitUserInstance 获取用户实例
//
// 获取用户实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSqlLimitUserInstance(request *model.ListSqlLimitUserInstanceRequest) (*model.ListSqlLimitUserInstanceResponse, error) {
	requestDef := GenReqDefForListSqlLimitUserInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlLimitUserInstanceResponse), nil
	}
}

// ListSqlLimitUserInstanceInvoker 获取用户实例
func (c *DasClient) ListSqlLimitUserInstanceInvoker(request *model.ListSqlLimitUserInstanceRequest) *ListSqlLimitUserInstanceInvoker {
	requestDef := GenReqDefForListSqlLimitUserInstance()
	return &ListSqlLimitUserInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlTemplateComparisons 查询SQL模板对比列表
//
// 查询SQL模板对比列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSqlTemplateComparisons(request *model.ListSqlTemplateComparisonsRequest) (*model.ListSqlTemplateComparisonsResponse, error) {
	requestDef := GenReqDefForListSqlTemplateComparisons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlTemplateComparisonsResponse), nil
	}
}

// ListSqlTemplateComparisonsInvoker 查询SQL模板对比列表
func (c *DasClient) ListSqlTemplateComparisonsInvoker(request *model.ListSqlTemplateComparisonsRequest) *ListSqlTemplateComparisonsInvoker {
	requestDef := GenReqDefForListSqlTemplateComparisons()
	return &ListSqlTemplateComparisonsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlTemplateDatabases 查询SQL模板数据库列表
//
// 查询SQL模板数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSqlTemplateDatabases(request *model.ListSqlTemplateDatabasesRequest) (*model.ListSqlTemplateDatabasesResponse, error) {
	requestDef := GenReqDefForListSqlTemplateDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlTemplateDatabasesResponse), nil
	}
}

// ListSqlTemplateDatabasesInvoker 查询SQL模板数据库列表
func (c *DasClient) ListSqlTemplateDatabasesInvoker(request *model.ListSqlTemplateDatabasesRequest) *ListSqlTemplateDatabasesInvoker {
	requestDef := GenReqDefForListSqlTemplateDatabases()
	return &ListSqlTemplateDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlTemplates 查询SQL模板列表
//
// 查询SQL模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSqlTemplates(request *model.ListSqlTemplatesRequest) (*model.ListSqlTemplatesResponse, error) {
	requestDef := GenReqDefForListSqlTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlTemplatesResponse), nil
	}
}

// ListSqlTemplatesInvoker 查询SQL模板列表
func (c *DasClient) ListSqlTemplatesInvoker(request *model.ListSqlTemplatesRequest) *ListSqlTemplatesInvoker {
	requestDef := GenReqDefForListSqlTemplates()
	return &ListSqlTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasksByBatchId 按批次ID查询全量SQL任务
//
// 按批次ID查询全量SQL任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTasksByBatchId(request *model.ListTasksByBatchIdRequest) (*model.ListTasksByBatchIdResponse, error) {
	requestDef := GenReqDefForListTasksByBatchId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksByBatchIdResponse), nil
	}
}

// ListTasksByBatchIdInvoker 按批次ID查询全量SQL任务
func (c *DasClient) ListTasksByBatchIdInvoker(request *model.ListTasksByBatchIdRequest) *ListTasksByBatchIdInvoker {
	requestDef := GenReqDefForListTasksByBatchId()
	return &ListTasksByBatchIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasksBySqlTemplateId 按SQL模板ID查询全量SQL任务
//
// 按SQL模板ID查询全量SQL任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTasksBySqlTemplateId(request *model.ListTasksBySqlTemplateIdRequest) (*model.ListTasksBySqlTemplateIdResponse, error) {
	requestDef := GenReqDefForListTasksBySqlTemplateId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksBySqlTemplateIdResponse), nil
	}
}

// ListTasksBySqlTemplateIdInvoker 按SQL模板ID查询全量SQL任务
func (c *DasClient) ListTasksBySqlTemplateIdInvoker(request *model.ListTasksBySqlTemplateIdRequest) *ListTasksBySqlTemplateIdInvoker {
	requestDef := GenReqDefForListTasksBySqlTemplateId()
	return &ListTasksBySqlTemplateIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasksByTaskId 按任务ID查询全量SQL任务
//
// 按任务ID查询全量SQL任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTasksByTaskId(request *model.ListTasksByTaskIdRequest) (*model.ListTasksByTaskIdResponse, error) {
	requestDef := GenReqDefForListTasksByTaskId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksByTaskIdResponse), nil
	}
}

// ListTasksByTaskIdInvoker 按任务ID查询全量SQL任务
func (c *DasClient) ListTasksByTaskIdInvoker(request *model.ListTasksByTaskIdRequest) *ListTasksByTaskIdInvoker {
	requestDef := GenReqDefForListTasksByTaskId()
	return &ListTasksByTaskIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateDatabaseComparisons 查询模板数据库对比列表
//
// 查询模板数据库对比列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTemplateDatabaseComparisons(request *model.ListTemplateDatabaseComparisonsRequest) (*model.ListTemplateDatabaseComparisonsResponse, error) {
	requestDef := GenReqDefForListTemplateDatabaseComparisons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateDatabaseComparisonsResponse), nil
	}
}

// ListTemplateDatabaseComparisonsInvoker 查询模板数据库对比列表
func (c *DasClient) ListTemplateDatabaseComparisonsInvoker(request *model.ListTemplateDatabaseComparisonsRequest) *ListTemplateDatabaseComparisonsInvoker {
	requestDef := GenReqDefForListTemplateDatabaseComparisons()
	return &ListTemplateDatabaseComparisonsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserInstanceList 获取用户实例列表
//
// 获取用户实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListUserInstanceList(request *model.ListUserInstanceListRequest) (*model.ListUserInstanceListResponse, error) {
	requestDef := GenReqDefForListUserInstanceList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserInstanceListResponse), nil
	}
}

// ListUserInstanceListInvoker 获取用户实例列表
func (c *DasClient) ListUserInstanceListInvoker(request *model.ListUserInstanceListRequest) *ListUserInstanceListInvoker {
	requestDef := GenReqDefForListUserInstanceList()
	return &ListUserInstanceListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryBinlogTask 重试binlog解析任务
//
// 重试binlog解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) RetryBinlogTask(request *model.RetryBinlogTaskRequest) (*model.RetryBinlogTaskResponse, error) {
	requestDef := GenReqDefForRetryBinlogTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryBinlogTaskResponse), nil
	}
}

// RetryBinlogTaskInvoker 重试binlog解析任务
func (c *DasClient) RetryBinlogTaskInvoker(request *model.RetryBinlogTaskRequest) *RetryBinlogTaskInvoker {
	requestDef := GenReqDefForRetryBinlogTask()
	return &RetryBinlogTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchBinlogParse 查看binlog解析详情
//
// 查看binlog解析详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SearchBinlogParse(request *model.SearchBinlogParseRequest) (*model.SearchBinlogParseResponse, error) {
	requestDef := GenReqDefForSearchBinlogParse()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchBinlogParseResponse), nil
	}
}

// SearchBinlogParseInvoker 查看binlog解析详情
func (c *DasClient) SearchBinlogParseInvoker(request *model.SearchBinlogParseRequest) *SearchBinlogParseInvoker {
	requestDef := GenReqDefForSearchBinlogParse()
	return &SearchBinlogParseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchErrorInfo4Api 查看binlog解析错误信息
//
// 查看binlog解析错误信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SearchErrorInfo4Api(request *model.SearchErrorInfo4ApiRequest) (*model.SearchErrorInfo4ApiResponse, error) {
	requestDef := GenReqDefForSearchErrorInfo4Api()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchErrorInfo4ApiResponse), nil
	}
}

// SearchErrorInfo4ApiInvoker 查看binlog解析错误信息
func (c *DasClient) SearchErrorInfo4ApiInvoker(request *model.SearchErrorInfo4ApiRequest) *SearchErrorInfo4ApiInvoker {
	requestDef := GenReqDefForSearchErrorInfo4Api()
	return &SearchErrorInfo4ApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchErrorInfoSource4Api 查看binlog解析错误信息条件
//
// 查看binlog解析错误信息条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SearchErrorInfoSource4Api(request *model.SearchErrorInfoSource4ApiRequest) (*model.SearchErrorInfoSource4ApiResponse, error) {
	requestDef := GenReqDefForSearchErrorInfoSource4Api()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchErrorInfoSource4ApiResponse), nil
	}
}

// SearchErrorInfoSource4ApiInvoker 查看binlog解析错误信息条件
func (c *DasClient) SearchErrorInfoSource4ApiInvoker(request *model.SearchErrorInfoSource4ApiRequest) *SearchErrorInfoSource4ApiInvoker {
	requestDef := GenReqDefForSearchErrorInfoSource4Api()
	return &SearchErrorInfoSource4ApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchNew 全量SQL搜索
//
// 全量SQL搜索
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SearchNew(request *model.SearchNewRequest) (*model.SearchNewResponse, error) {
	requestDef := GenReqDefForSearchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchNewResponse), nil
	}
}

// SearchNewInvoker 全量SQL搜索
func (c *DasClient) SearchNewInvoker(request *model.SearchNewRequest) *SearchNewInvoker {
	requestDef := GenReqDefForSearchNew()
	return &SearchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSqlSwitchNew 设置SQL开关
//
// 设置SQL开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SetSqlSwitchNew(request *model.SetSqlSwitchNewRequest) (*model.SetSqlSwitchNewResponse, error) {
	requestDef := GenReqDefForSetSqlSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSqlSwitchNewResponse), nil
	}
}

// SetSqlSwitchNewInvoker 设置SQL开关
func (c *DasClient) SetSqlSwitchNewInvoker(request *model.SetSqlSwitchNewRequest) *SetSqlSwitchNewInvoker {
	requestDef := GenReqDefForSetSqlSwitchNew()
	return &SetSqlSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBinlogExportTaskInfo 查询binlog导出任务信息
//
// 查询binlog导出任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowBinlogExportTaskInfo(request *model.ShowBinlogExportTaskInfoRequest) (*model.ShowBinlogExportTaskInfoResponse, error) {
	requestDef := GenReqDefForShowBinlogExportTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBinlogExportTaskInfoResponse), nil
	}
}

// ShowBinlogExportTaskInfoInvoker 查询binlog导出任务信息
func (c *DasClient) ShowBinlogExportTaskInfoInvoker(request *model.ShowBinlogExportTaskInfoRequest) *ShowBinlogExportTaskInfoInvoker {
	requestDef := GenReqDefForShowBinlogExportTaskInfo()
	return &ShowBinlogExportTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBinlogParse 查看binlog概览
//
// 查看binlog概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowBinlogParse(request *model.ShowBinlogParseRequest) (*model.ShowBinlogParseResponse, error) {
	requestDef := GenReqDefForShowBinlogParse()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBinlogParseResponse), nil
	}
}

// ShowBinlogParseInvoker 查看binlog概览
func (c *DasClient) ShowBinlogParseInvoker(request *model.ShowBinlogParseRequest) *ShowBinlogParseInvoker {
	requestDef := GenReqDefForShowBinlogParse()
	return &ShowBinlogParseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBinlogTaskInfo 查看binlog解析任务详情
//
// 查看binlog解析任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowBinlogTaskInfo(request *model.ShowBinlogTaskInfoRequest) (*model.ShowBinlogTaskInfoResponse, error) {
	requestDef := GenReqDefForShowBinlogTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBinlogTaskInfoResponse), nil
	}
}

// ShowBinlogTaskInfoInvoker 查看binlog解析任务详情
func (c *DasClient) ShowBinlogTaskInfoInvoker(request *model.ShowBinlogTaskInfoRequest) *ShowBinlogTaskInfoInvoker {
	requestDef := GenReqDefForShowBinlogTaskInfo()
	return &ShowBinlogTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdsConnectionStat DDS连接统计
//
// DDS连接统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDdsConnectionStat(request *model.ShowDdsConnectionStatRequest) (*model.ShowDdsConnectionStatResponse, error) {
	requestDef := GenReqDefForShowDdsConnectionStat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdsConnectionStatResponse), nil
	}
}

// ShowDdsConnectionStatInvoker DDS连接统计
func (c *DasClient) ShowDdsConnectionStatInvoker(request *model.ShowDdsConnectionStatRequest) *ShowDdsConnectionStatInvoker {
	requestDef := GenReqDefForShowDdsConnectionStat()
	return &ShowDdsConnectionStatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockOriginData 获取死锁原始数据
//
// 获取死锁原始数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockOriginData(request *model.ShowDeadLockOriginDataRequest) (*model.ShowDeadLockOriginDataResponse, error) {
	requestDef := GenReqDefForShowDeadLockOriginData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockOriginDataResponse), nil
	}
}

// ShowDeadLockOriginDataInvoker 获取死锁原始数据
func (c *DasClient) ShowDeadLockOriginDataInvoker(request *model.ShowDeadLockOriginDataRequest) *ShowDeadLockOriginDataInvoker {
	requestDef := GenReqDefForShowDeadLockOriginData()
	return &ShowDeadLockOriginDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockRelationship 获取死锁关系
//
// 获取死锁关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockRelationship(request *model.ShowDeadLockRelationshipRequest) (*model.ShowDeadLockRelationshipResponse, error) {
	requestDef := GenReqDefForShowDeadLockRelationship()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockRelationshipResponse), nil
	}
}

// ShowDeadLockRelationshipInvoker 获取死锁关系
func (c *DasClient) ShowDeadLockRelationshipInvoker(request *model.ShowDeadLockRelationshipRequest) *ShowDeadLockRelationshipInvoker {
	requestDef := GenReqDefForShowDeadLockRelationship()
	return &ShowDeadLockRelationshipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockStatistics 获取死锁统计
//
// 获取死锁统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockStatistics(request *model.ShowDeadLockStatisticsRequest) (*model.ShowDeadLockStatisticsResponse, error) {
	requestDef := GenReqDefForShowDeadLockStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockStatisticsResponse), nil
	}
}

// ShowDeadLockStatisticsInvoker 获取死锁统计
func (c *DasClient) ShowDeadLockStatisticsInvoker(request *model.ShowDeadLockStatisticsRequest) *ShowDeadLockStatisticsInvoker {
	requestDef := GenReqDefForShowDeadLockStatistics()
	return &ShowDeadLockStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockTrend 获取死锁趋势
//
// 获取死锁趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockTrend(request *model.ShowDeadLockTrendRequest) (*model.ShowDeadLockTrendResponse, error) {
	requestDef := GenReqDefForShowDeadLockTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockTrendResponse), nil
	}
}

// ShowDeadLockTrendInvoker 获取死锁趋势
func (c *DasClient) ShowDeadLockTrendInvoker(request *model.ShowDeadLockTrendRequest) *ShowDeadLockTrendInvoker {
	requestDef := GenReqDefForShowDeadLockTrend()
	return &ShowDeadLockTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecuteResultWithoutKey 查询SQL执行结果
//
// 查询SQL执行结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowExecuteResultWithoutKey(request *model.ShowExecuteResultWithoutKeyRequest) (*model.ShowExecuteResultWithoutKeyResponse, error) {
	requestDef := GenReqDefForShowExecuteResultWithoutKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecuteResultWithoutKeyResponse), nil
	}
}

// ShowExecuteResultWithoutKeyInvoker 查询SQL执行结果
func (c *DasClient) ShowExecuteResultWithoutKeyInvoker(request *model.ShowExecuteResultWithoutKeyRequest) *ShowExecuteResultWithoutKeyInvoker {
	requestDef := GenReqDefForShowExecuteResultWithoutKey()
	return &ShowExecuteResultWithoutKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecuteResultWithoutKeyNoRetry 查询SQL执行结果（POST）
//
// 查询SQL执行结果（POST）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowExecuteResultWithoutKeyNoRetry(request *model.ShowExecuteResultWithoutKeyNoRetryRequest) (*model.ShowExecuteResultWithoutKeyNoRetryResponse, error) {
	requestDef := GenReqDefForShowExecuteResultWithoutKeyNoRetry()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecuteResultWithoutKeyNoRetryResponse), nil
	}
}

// ShowExecuteResultWithoutKeyNoRetryInvoker 查询SQL执行结果（POST）
func (c *DasClient) ShowExecuteResultWithoutKeyNoRetryInvoker(request *model.ShowExecuteResultWithoutKeyNoRetryRequest) *ShowExecuteResultWithoutKeyNoRetryInvoker {
	requestDef := GenReqDefForShowExecuteResultWithoutKeyNoRetry()
	return &ShowExecuteResultWithoutKeyNoRetryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecutionPlan 获取执行计划
//
// 获取执行计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowExecutionPlan(request *model.ShowExecutionPlanRequest) (*model.ShowExecutionPlanResponse, error) {
	requestDef := GenReqDefForShowExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecutionPlanResponse), nil
	}
}

// ShowExecutionPlanInvoker 获取执行计划
func (c *DasClient) ShowExecutionPlanInvoker(request *model.ShowExecutionPlanRequest) *ShowExecutionPlanInvoker {
	requestDef := GenReqDefForShowExecutionPlan()
	return &ShowExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecutionTimeTemplateTrend 查询执行时间模板趋势
//
// 查询执行时间模板趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowExecutionTimeTemplateTrend(request *model.ShowExecutionTimeTemplateTrendRequest) (*model.ShowExecutionTimeTemplateTrendResponse, error) {
	requestDef := GenReqDefForShowExecutionTimeTemplateTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecutionTimeTemplateTrendResponse), nil
	}
}

// ShowExecutionTimeTemplateTrendInvoker 查询执行时间模板趋势
func (c *DasClient) ShowExecutionTimeTemplateTrendInvoker(request *model.ShowExecutionTimeTemplateTrendRequest) *ShowExecutionTimeTemplateTrendInvoker {
	requestDef := GenReqDefForShowExecutionTimeTemplateTrend()
	return &ShowExecutionTimeTemplateTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFragmentSwitch 是否展示fragment任务
//
// 是否展示fragment任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowFragmentSwitch(request *model.ShowFragmentSwitchRequest) (*model.ShowFragmentSwitchResponse, error) {
	requestDef := GenReqDefForShowFragmentSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFragmentSwitchResponse), nil
	}
}

// ShowFragmentSwitchInvoker 是否展示fragment任务
func (c *DasClient) ShowFragmentSwitchInvoker(request *model.ShowFragmentSwitchRequest) *ShowFragmentSwitchInvoker {
	requestDef := GenReqDefForShowFragmentSwitch()
	return &ShowFragmentSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceHealthReport4Api 获取实例健康报告
//
// 获取实例健康报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceHealthReport4Api(request *model.ShowInstanceHealthReport4ApiRequest) (*model.ShowInstanceHealthReport4ApiResponse, error) {
	requestDef := GenReqDefForShowInstanceHealthReport4Api()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceHealthReport4ApiResponse), nil
	}
}

// ShowInstanceHealthReport4ApiInvoker 获取实例健康报告
func (c *DasClient) ShowInstanceHealthReport4ApiInvoker(request *model.ShowInstanceHealthReport4ApiRequest) *ShowInstanceHealthReport4ApiInvoker {
	requestDef := GenReqDefForShowInstanceHealthReport4Api()
	return &ShowInstanceHealthReport4ApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceInfo 获取实例信息
//
// 获取实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceInfo(request *model.ShowInstanceInfoRequest) (*model.ShowInstanceInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceInfoResponse), nil
	}
}

// ShowInstanceInfoInvoker 获取实例信息
func (c *DasClient) ShowInstanceInfoInvoker(request *model.ShowInstanceInfoRequest) *ShowInstanceInfoInvoker {
	requestDef := GenReqDefForShowInstanceInfo()
	return &ShowInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceLogUsage 查看实例日志存储使用量
//
// 查看实例日志存储使用量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceLogUsage(request *model.ShowInstanceLogUsageRequest) (*model.ShowInstanceLogUsageResponse, error) {
	requestDef := GenReqDefForShowInstanceLogUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceLogUsageResponse), nil
	}
}

// ShowInstanceLogUsageInvoker 查看实例日志存储使用量
func (c *DasClient) ShowInstanceLogUsageInvoker(request *model.ShowInstanceLogUsageRequest) *ShowInstanceLogUsageInvoker {
	requestDef := GenReqDefForShowInstanceLogUsage()
	return &ShowInstanceLogUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceMetric 查询实例指标
//
// 查询实例指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceMetric(request *model.ShowInstanceMetricRequest) (*model.ShowInstanceMetricResponse, error) {
	requestDef := GenReqDefForShowInstanceMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMetricResponse), nil
	}
}

// ShowInstanceMetricInvoker 查询实例指标
func (c *DasClient) ShowInstanceMetricInvoker(request *model.ShowInstanceMetricRequest) *ShowInstanceMetricInvoker {
	requestDef := GenReqDefForShowInstanceMetric()
	return &ShowInstanceMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceNodesInfo 获取实例节点信息
//
// 获取实例节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceNodesInfo(request *model.ShowInstanceNodesInfoRequest) (*model.ShowInstanceNodesInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceNodesInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceNodesInfoResponse), nil
	}
}

// ShowInstanceNodesInfoInvoker 获取实例节点信息
func (c *DasClient) ShowInstanceNodesInfoInvoker(request *model.ShowInstanceNodesInfoRequest) *ShowInstanceNodesInfoInvoker {
	requestDef := GenReqDefForShowInstanceNodesInfo()
	return &ShowInstanceNodesInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIsSignedProtocol 是否签署数据安全协议
//
// 是否签署数据安全协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowIsSignedProtocol(request *model.ShowIsSignedProtocolRequest) (*model.ShowIsSignedProtocolResponse, error) {
	requestDef := GenReqDefForShowIsSignedProtocol()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIsSignedProtocolResponse), nil
	}
}

// ShowIsSignedProtocolInvoker 是否签署数据安全协议
func (c *DasClient) ShowIsSignedProtocolInvoker(request *model.ShowIsSignedProtocolRequest) *ShowIsSignedProtocolInvoker {
	requestDef := GenReqDefForShowIsSignedProtocol()
	return &ShowIsSignedProtocolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKillProcessTask 查询Kill进程任务
//
// 查询Kill进程任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowKillProcessTask(request *model.ShowKillProcessTaskRequest) (*model.ShowKillProcessTaskResponse, error) {
	requestDef := GenReqDefForShowKillProcessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKillProcessTaskResponse), nil
	}
}

// ShowKillProcessTaskInvoker 查询Kill进程任务
func (c *DasClient) ShowKillProcessTaskInvoker(request *model.ShowKillProcessTaskRequest) *ShowKillProcessTaskInvoker {
	requestDef := GenReqDefForShowKillProcessTask()
	return &ShowKillProcessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestDeadLockSnapshot4Api 查询最新死锁快照
//
// 查询最新死锁快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLatestDeadLockSnapshot4Api(request *model.ShowLatestDeadLockSnapshot4ApiRequest) (*model.ShowLatestDeadLockSnapshot4ApiResponse, error) {
	requestDef := GenReqDefForShowLatestDeadLockSnapshot4Api()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestDeadLockSnapshot4ApiResponse), nil
	}
}

// ShowLatestDeadLockSnapshot4ApiInvoker 查询最新死锁快照
func (c *DasClient) ShowLatestDeadLockSnapshot4ApiInvoker(request *model.ShowLatestDeadLockSnapshot4ApiRequest) *ShowLatestDeadLockSnapshot4ApiInvoker {
	requestDef := GenReqDefForShowLatestDeadLockSnapshot4Api()
	return &ShowLatestDeadLockSnapshot4ApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetaLock 查询元数据锁
//
// 查询元数据锁
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowMetaLock(request *model.ShowMetaLockRequest) (*model.ShowMetaLockResponse, error) {
	requestDef := GenReqDefForShowMetaLock()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetaLockResponse), nil
	}
}

// ShowMetaLockInvoker 查询元数据锁
func (c *DasClient) ShowMetaLockInvoker(request *model.ShowMetaLockRequest) *ShowMetaLockInvoker {
	requestDef := GenReqDefForShowMetaLock()
	return &ShowMetaLockInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetaLockSnapshot 查询元数据锁快照
//
// 查询元数据锁快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowMetaLockSnapshot(request *model.ShowMetaLockSnapshotRequest) (*model.ShowMetaLockSnapshotResponse, error) {
	requestDef := GenReqDefForShowMetaLockSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetaLockSnapshotResponse), nil
	}
}

// ShowMetaLockSnapshotInvoker 查询元数据锁快照
func (c *DasClient) ShowMetaLockSnapshotInvoker(request *model.ShowMetaLockSnapshotRequest) *ShowMetaLockSnapshotInvoker {
	requestDef := GenReqDefForShowMetaLockSnapshot()
	return &ShowMetaLockSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOpeningInfo 获取开通信息
//
// 获取开通信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowOpeningInfo(request *model.ShowOpeningInfoRequest) (*model.ShowOpeningInfoResponse, error) {
	requestDef := GenReqDefForShowOpeningInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOpeningInfoResponse), nil
	}
}

// ShowOpeningInfoInvoker 获取开通信息
func (c *DasClient) ShowOpeningInfoInvoker(request *model.ShowOpeningInfoRequest) *ShowOpeningInfoInvoker {
	requestDef := GenReqDefForShowOpeningInfo()
	return &ShowOpeningInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSingleTemplateTrend 查询单个模板趋势
//
// 查询单个模板趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSingleTemplateTrend(request *model.ShowSingleTemplateTrendRequest) (*model.ShowSingleTemplateTrendResponse, error) {
	requestDef := GenReqDefForShowSingleTemplateTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleTemplateTrendResponse), nil
	}
}

// ShowSingleTemplateTrendInvoker 查询单个模板趋势
func (c *DasClient) ShowSingleTemplateTrendInvoker(request *model.ShowSingleTemplateTrendRequest) *ShowSingleTemplateTrendInvoker {
	requestDef := GenReqDefForShowSingleTemplateTrend()
	return &ShowSingleTemplateTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlTemplateTrend 查询SQL模板趋势
//
// 查询SQL模板趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlTemplateTrend(request *model.ShowSqlTemplateTrendRequest) (*model.ShowSqlTemplateTrendResponse, error) {
	requestDef := GenReqDefForShowSqlTemplateTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlTemplateTrendResponse), nil
	}
}

// ShowSqlTemplateTrendInvoker 查询SQL模板趋势
func (c *DasClient) ShowSqlTemplateTrendInvoker(request *model.ShowSqlTemplateTrendRequest) *ShowSqlTemplateTrendInvoker {
	requestDef := GenReqDefForShowSqlTemplateTrend()
	return &ShowSqlTemplateTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSupportKeyString 支持的关键字
//
// 支持的关键字
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSupportKeyString(request *model.ShowSupportKeyStringRequest) (*model.ShowSupportKeyStringResponse, error) {
	requestDef := GenReqDefForShowSupportKeyString()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSupportKeyStringResponse), nil
	}
}

// ShowSupportKeyStringInvoker 支持的关键字
func (c *DasClient) ShowSupportKeyStringInvoker(request *model.ShowSupportKeyStringRequest) *ShowSupportKeyStringInvoker {
	requestDef := GenReqDefForShowSupportKeyString()
	return &ShowSupportKeyStringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTuningResult 获取调优结果
//
// 获取调优结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowTuningResult(request *model.ShowTuningResultRequest) (*model.ShowTuningResultResponse, error) {
	requestDef := GenReqDefForShowTuningResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTuningResultResponse), nil
	}
}

// ShowTuningResultInvoker 获取调优结果
func (c *DasClient) ShowTuningResultInvoker(request *model.ShowTuningResultRequest) *ShowTuningResultInvoker {
	requestDef := GenReqDefForShowTuningResult()
	return &ShowTuningResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWaitingLocksSnapshot 查询InnoDB锁等待快照
//
// 查询InnoDB锁等待快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowWaitingLocksSnapshot(request *model.ShowWaitingLocksSnapshotRequest) (*model.ShowWaitingLocksSnapshotResponse, error) {
	requestDef := GenReqDefForShowWaitingLocksSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWaitingLocksSnapshotResponse), nil
	}
}

// ShowWaitingLocksSnapshotInvoker 查询InnoDB锁等待快照
func (c *DasClient) ShowWaitingLocksSnapshotInvoker(request *model.ShowWaitingLocksSnapshotRequest) *ShowWaitingLocksSnapshotInvoker {
	requestDef := GenReqDefForShowWaitingLocksSnapshot()
	return &ShowWaitingLocksSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWdrSnapshot 获取WDR快照列表
//
// 获取WDR快照列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowWdrSnapshot(request *model.ShowWdrSnapshotRequest) (*model.ShowWdrSnapshotResponse, error) {
	requestDef := GenReqDefForShowWdrSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWdrSnapshotResponse), nil
	}
}

// ShowWdrSnapshotInvoker 获取WDR快照列表
func (c *DasClient) ShowWdrSnapshotInvoker(request *model.ShowWdrSnapshotRequest) *ShowWdrSnapshotInvoker {
	requestDef := GenReqDefForShowWdrSnapshot()
	return &ShowWdrSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SignProtocolNew 签署数据安全协议
//
// 签署数据安全协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SignProtocolNew(request *model.SignProtocolNewRequest) (*model.SignProtocolNewResponse, error) {
	requestDef := GenReqDefForSignProtocolNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignProtocolNewResponse), nil
	}
}

// SignProtocolNewInvoker 签署数据安全协议
func (c *DasClient) SignProtocolNewInvoker(request *model.SignProtocolNewRequest) *SignProtocolNewInvoker {
	requestDef := GenReqDefForSignProtocolNew()
	return &SignProtocolNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopBinlogTask 停止binlog解析任务
//
// 停止binlog解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) StopBinlogTask(request *model.StopBinlogTaskRequest) (*model.StopBinlogTaskResponse, error) {
	requestDef := GenReqDefForStopBinlogTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopBinlogTaskResponse), nil
	}
}

// StopBinlogTaskInvoker 停止binlog解析任务
func (c *DasClient) StopBinlogTaskInvoker(request *model.StopBinlogTaskRequest) *StopBinlogTaskInvoker {
	requestDef := GenReqDefForStopBinlogTask()
	return &StopBinlogTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeInstanceReportNew 订阅实例报告
//
// 订阅实例报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SubscribeInstanceReportNew(request *model.SubscribeInstanceReportNewRequest) (*model.SubscribeInstanceReportNewResponse, error) {
	requestDef := GenReqDefForSubscribeInstanceReportNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeInstanceReportNewResponse), nil
	}
}

// SubscribeInstanceReportNewInvoker 订阅实例报告
func (c *DasClient) SubscribeInstanceReportNewInvoker(request *model.SubscribeInstanceReportNewRequest) *SubscribeInstanceReportNewInvoker {
	requestDef := GenReqDefForSubscribeInstanceReportNew()
	return &SubscribeInstanceReportNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SynchronizeInstanceListNew 同步实例列表
//
// 同步实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SynchronizeInstanceListNew(request *model.SynchronizeInstanceListNewRequest) (*model.SynchronizeInstanceListNewResponse, error) {
	requestDef := GenReqDefForSynchronizeInstanceListNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SynchronizeInstanceListNewResponse), nil
	}
}

// SynchronizeInstanceListNewInvoker 同步实例列表
func (c *DasClient) SynchronizeInstanceListNewInvoker(request *model.SynchronizeInstanceListNewRequest) *SynchronizeInstanceListNewInvoker {
	requestDef := GenReqDefForSynchronizeInstanceListNew()
	return &SynchronizeInstanceListNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnsubscribeInstanceReportNew 取消订阅实例报告
//
// 取消订阅实例报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UnsubscribeInstanceReportNew(request *model.UnsubscribeInstanceReportNewRequest) (*model.UnsubscribeInstanceReportNewResponse, error) {
	requestDef := GenReqDefForUnsubscribeInstanceReportNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnsubscribeInstanceReportNewResponse), nil
	}
}

// UnsubscribeInstanceReportNewInvoker 取消订阅实例报告
func (c *DasClient) UnsubscribeInstanceReportNewInvoker(request *model.UnsubscribeInstanceReportNewRequest) *UnsubscribeInstanceReportNewInvoker {
	requestDef := GenReqDefForUnsubscribeInstanceReportNew()
	return &UnsubscribeInstanceReportNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfig 设置实例配置
//
// Space Set Config New
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateInstanceConfig(request *model.UpdateInstanceConfigRequest) (*model.UpdateInstanceConfigResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigResponse), nil
	}
}

// UpdateInstanceConfigInvoker 设置实例配置
func (c *DasClient) UpdateInstanceConfigInvoker(request *model.UpdateInstanceConfigRequest) *UpdateInstanceConfigInvoker {
	requestDef := GenReqDefForUpdateInstanceConfig()
	return &UpdateInstanceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSearchPathFlag 设置searchpath开关
//
// 设置searchpath开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateSearchPathFlag(request *model.UpdateSearchPathFlagRequest) (*model.UpdateSearchPathFlagResponse, error) {
	requestDef := GenReqDefForUpdateSearchPathFlag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSearchPathFlagResponse), nil
	}
}

// UpdateSearchPathFlagInvoker 设置searchpath开关
func (c *DasClient) UpdateSearchPathFlagInvoker(request *model.UpdateSearchPathFlagRequest) *UpdateSearchPathFlagInvoker {
	requestDef := GenReqDefForUpdateSearchPathFlag()
	return &UpdateSearchPathFlagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSharedInfoNew 更新共享信息
//
// 更新共享信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateSharedInfoNew(request *model.UpdateSharedInfoNewRequest) (*model.UpdateSharedInfoNewResponse, error) {
	requestDef := GenReqDefForUpdateSharedInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSharedInfoNewResponse), nil
	}
}

// UpdateSharedInfoNewInvoker 更新共享信息
func (c *DasClient) UpdateSharedInfoNewInvoker(request *model.UpdateSharedInfoNewRequest) *UpdateSharedInfoNewInvoker {
	requestDef := GenReqDefForUpdateSharedInfoNew()
	return &UpdateSharedInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// VerifyConnectionNew 验证数据库实例连接
//
// 验证数据库实例连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) VerifyConnectionNew(request *model.VerifyConnectionNewRequest) (*model.VerifyConnectionNewResponse, error) {
	requestDef := GenReqDefForVerifyConnectionNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.VerifyConnectionNewResponse), nil
	}
}

// VerifyConnectionNewInvoker 验证数据库实例连接
func (c *DasClient) VerifyConnectionNewInvoker(request *model.VerifyConnectionNewRequest) *VerifyConnectionNewInvoker {
	requestDef := GenReqDefForVerifyConnectionNew()
	return &VerifyConnectionNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEmailTemplate 新增邮件模板
//
// 新增邮件模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddEmailTemplate(request *model.AddEmailTemplateRequest) (*model.AddEmailTemplateResponse, error) {
	requestDef := GenReqDefForAddEmailTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEmailTemplateResponse), nil
	}
}

// AddEmailTemplateInvoker 新增邮件模板
func (c *DasClient) AddEmailTemplateInvoker(request *model.AddEmailTemplateRequest) *AddEmailTemplateInvoker {
	requestDef := GenReqDefForAddEmailTemplate()
	return &AddEmailTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddFullSqlTask 创建全量SQL明细解析任务
//
// 创建全量SQL明细解析任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddFullSqlTask(request *model.AddFullSqlTaskRequest) (*model.AddFullSqlTaskResponse, error) {
	requestDef := GenReqDefForAddFullSqlTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFullSqlTaskResponse), nil
	}
}

// AddFullSqlTaskInvoker 创建全量SQL明细解析任务
func (c *DasClient) AddFullSqlTaskInvoker(request *model.AddFullSqlTaskRequest) *AddFullSqlTaskInvoker {
	requestDef := GenReqDefForAddFullSqlTask()
	return &AddFullSqlTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddInstanceGroup 新增实例组
//
// 新增实例组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddInstanceGroup(request *model.AddInstanceGroupRequest) (*model.AddInstanceGroupResponse, error) {
	requestDef := GenReqDefForAddInstanceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddInstanceGroupResponse), nil
	}
}

// AddInstanceGroupInvoker 新增实例组
func (c *DasClient) AddInstanceGroupInvoker(request *model.AddInstanceGroupRequest) *AddInstanceGroupInvoker {
	requestDef := GenReqDefForAddInstanceGroup()
	return &AddInstanceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddInstanceToGroup 将实例添加到实例组
//
// 将实例添加到实例组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddInstanceToGroup(request *model.AddInstanceToGroupRequest) (*model.AddInstanceToGroupResponse, error) {
	requestDef := GenReqDefForAddInstanceToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddInstanceToGroupResponse), nil
	}
}

// AddInstanceToGroupInvoker 将实例添加到实例组
func (c *DasClient) AddInstanceToGroupInvoker(request *model.AddInstanceToGroupRequest) *AddInstanceToGroupInvoker {
	requestDef := GenReqDefForAddInstanceToGroup()
	return &AddInstanceToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSqlLimitingRecordNew 新增SQL限流规则
//
// 新增SQL限流规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) AddSqlLimitingRecordNew(request *model.AddSqlLimitingRecordNewRequest) (*model.AddSqlLimitingRecordNewResponse, error) {
	requestDef := GenReqDefForAddSqlLimitingRecordNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSqlLimitingRecordNewResponse), nil
	}
}

// AddSqlLimitingRecordNewInvoker 新增SQL限流规则
func (c *DasClient) AddSqlLimitingRecordNewInvoker(request *model.AddSqlLimitingRecordNewRequest) *AddSqlLimitingRecordNewInvoker {
	requestDef := GenReqDefForAddSqlLimitingRecordNew()
	return &AddSqlLimitingRecordNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteConnectionNew 批量删除连接
//
// 批量删除连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) BatchDeleteConnectionNew(request *model.BatchDeleteConnectionNewRequest) (*model.BatchDeleteConnectionNewResponse, error) {
	requestDef := GenReqDefForBatchDeleteConnectionNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteConnectionNewResponse), nil
	}
}

// BatchDeleteConnectionNewInvoker 批量删除连接
func (c *DasClient) BatchDeleteConnectionNewInvoker(request *model.BatchDeleteConnectionNewRequest) *BatchDeleteConnectionNewInvoker {
	requestDef := GenReqDefForBatchDeleteConnectionNew()
	return &BatchDeleteConnectionNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSendEmail 批量发送邮件
//
// 批量发送邮件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) BatchSendEmail(request *model.BatchSendEmailRequest) (*model.BatchSendEmailResponse, error) {
	requestDef := GenReqDefForBatchSendEmail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSendEmailResponse), nil
	}
}

// BatchSendEmailInvoker 批量发送邮件
func (c *DasClient) BatchSendEmailInvoker(request *model.BatchSendEmailRequest) *BatchSendEmailInvoker {
	requestDef := GenReqDefForBatchSendEmail()
	return &BatchSendEmailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSubscribeReport 批量订阅/取消订阅
//
// 批量订阅/取消订阅
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) BatchSubscribeReport(request *model.BatchSubscribeReportRequest) (*model.BatchSubscribeReportResponse, error) {
	requestDef := GenReqDefForBatchSubscribeReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSubscribeReportResponse), nil
	}
}

// BatchSubscribeReportInvoker 批量订阅/取消订阅
func (c *DasClient) BatchSubscribeReportInvoker(request *model.BatchSubscribeReportRequest) *BatchSubscribeReportInvoker {
	requestDef := GenReqDefForBatchSubscribeReport()
	return &BatchSubscribeReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelShareNew 取消共享链接
//
// 取消共享链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CancelShareNew(request *model.CancelShareNewRequest) (*model.CancelShareNewResponse, error) {
	requestDef := GenReqDefForCancelShareNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelShareNewResponse), nil
	}
}

// CancelShareNewInvoker 取消共享链接
func (c *DasClient) CancelShareNewInvoker(request *model.CancelShareNewRequest) *CancelShareNewInvoker {
	requestDef := GenReqDefForCancelShareNew()
	return &CancelShareNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeChargeMode 设置实例付费/免费模式
//
// 设置实例付费/免费模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeChargeMode(request *model.ChangeChargeModeRequest) (*model.ChangeChargeModeResponse, error) {
	requestDef := GenReqDefForChangeChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeChargeModeResponse), nil
	}
}

// ChangeChargeModeInvoker 设置实例付费/免费模式
func (c *DasClient) ChangeChargeModeInvoker(request *model.ChangeChargeModeRequest) *ChangeChargeModeInvoker {
	requestDef := GenReqDefForChangeChargeMode()
	return &ChangeChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDeadLockSwitchNew 修改死锁开关
//
// 修改死锁开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeDeadLockSwitchNew(request *model.ChangeDeadLockSwitchNewRequest) (*model.ChangeDeadLockSwitchNewResponse, error) {
	requestDef := GenReqDefForChangeDeadLockSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDeadLockSwitchNewResponse), nil
	}
}

// ChangeDeadLockSwitchNewInvoker 修改死锁开关
func (c *DasClient) ChangeDeadLockSwitchNewInvoker(request *model.ChangeDeadLockSwitchNewRequest) *ChangeDeadLockSwitchNewInvoker {
	requestDef := GenReqDefForChangeDeadLockSwitchNew()
	return &ChangeDeadLockSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeFullDeadLockSwitch 设置全量死锁开关
//
// 设置全量死锁开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeFullDeadLockSwitch(request *model.ChangeFullDeadLockSwitchRequest) (*model.ChangeFullDeadLockSwitchResponse, error) {
	requestDef := GenReqDefForChangeFullDeadLockSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeFullDeadLockSwitchResponse), nil
	}
}

// ChangeFullDeadLockSwitchInvoker 设置全量死锁开关
func (c *DasClient) ChangeFullDeadLockSwitchInvoker(request *model.ChangeFullDeadLockSwitchRequest) *ChangeFullDeadLockSwitchInvoker {
	requestDef := GenReqDefForChangeFullDeadLockSwitch()
	return &ChangeFullDeadLockSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePaymentModeNew 设置实例付费/免费模式
//
// 设置实例付费/免费模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangePaymentModeNew(request *model.ChangePaymentModeNewRequest) (*model.ChangePaymentModeNewResponse, error) {
	requestDef := GenReqDefForChangePaymentModeNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePaymentModeNewResponse), nil
	}
}

// ChangePaymentModeNewInvoker 设置实例付费/免费模式
func (c *DasClient) ChangePaymentModeNewInvoker(request *model.ChangePaymentModeNewRequest) *ChangePaymentModeNewInvoker {
	requestDef := GenReqDefForChangePaymentModeNew()
	return &ChangePaymentModeNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSqlLimitSwitchStatus 设置SQL限流开关状态
//
// 设置SQL限流开关状态。目前仅支持MySQL数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeSqlLimitSwitchStatus(request *model.ChangeSqlLimitSwitchStatusRequest) (*model.ChangeSqlLimitSwitchStatusResponse, error) {
	requestDef := GenReqDefForChangeSqlLimitSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSqlLimitSwitchStatusResponse), nil
	}
}

// ChangeSqlLimitSwitchStatusInvoker 设置SQL限流开关状态
func (c *DasClient) ChangeSqlLimitSwitchStatusInvoker(request *model.ChangeSqlLimitSwitchStatusRequest) *ChangeSqlLimitSwitchStatusInvoker {
	requestDef := GenReqDefForChangeSqlLimitSwitchStatus()
	return &ChangeSqlLimitSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSqlSwitch 开启/关闭全量SQL、慢SQL开关
//
// 打开或者关闭DAS收集全量SQL开关，开启后，实例的性能损耗在5%以内。开启全量SQL后，本服务会对SQL的文本内容进行存储，以便进行分析。用户可自行设置全量SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。
// 打开或者关闭DAS收集慢SQL开关。开启慢SQL后，本服务会对慢SQL的文本内容进行存储，以便进行分析。用户可自行设置慢SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。该功能仅支持付费实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeSqlSwitch(request *model.ChangeSqlSwitchRequest) (*model.ChangeSqlSwitchResponse, error) {
	requestDef := GenReqDefForChangeSqlSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSqlSwitchResponse), nil
	}
}

// ChangeSqlSwitchInvoker 开启/关闭全量SQL、慢SQL开关
func (c *DasClient) ChangeSqlSwitchInvoker(request *model.ChangeSqlSwitchRequest) *ChangeSqlSwitchInvoker {
	requestDef := GenReqDefForChangeSqlSwitch()
	return &ChangeSqlSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeTransactionSwitchStatus 开启/关闭历史事务开关
//
// 开启/关闭历史事务开关，仅支持MySQL引擎，并且依赖开启全量SQL或者慢SQL功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ChangeTransactionSwitchStatus(request *model.ChangeTransactionSwitchStatusRequest) (*model.ChangeTransactionSwitchStatusResponse, error) {
	requestDef := GenReqDefForChangeTransactionSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeTransactionSwitchStatusResponse), nil
	}
}

// ChangeTransactionSwitchStatusInvoker 开启/关闭历史事务开关
func (c *DasClient) ChangeTransactionSwitchStatusInvoker(request *model.ChangeTransactionSwitchStatusRequest) *ChangeTransactionSwitchStatusInvoker {
	requestDef := GenReqDefForChangeTransactionSwitchStatus()
	return &ChangeTransactionSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckCredential 测试AK/SK
//
// 测试AK/SK，测试用户AK/SK能否正常访问OBS桶。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CheckCredential(request *model.CheckCredentialRequest) (*model.CheckCredentialResponse, error) {
	requestDef := GenReqDefForCheckCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCredentialResponse), nil
	}
}

// CheckCredentialInvoker 测试AK/SK
func (c *DasClient) CheckCredentialInvoker(request *model.CheckCredentialRequest) *CheckCredentialInvoker {
	requestDef := GenReqDefForCheckCredential()
	return &CheckCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckCredentialForBatchInspection 测试AK/SK
//
// 测试AK/SK，测试用户AK/SK能否正常访问OBS桶。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CheckCredentialForBatchInspection(request *model.CheckCredentialForBatchInspectionRequest) (*model.CheckCredentialForBatchInspectionResponse, error) {
	requestDef := GenReqDefForCheckCredentialForBatchInspection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCredentialForBatchInspectionResponse), nil
	}
}

// CheckCredentialForBatchInspectionInvoker 测试AK/SK
func (c *DasClient) CheckCredentialForBatchInspectionInvoker(request *model.CheckCredentialForBatchInspectionRequest) *CheckCredentialForBatchInspectionInvoker {
	requestDef := GenReqDefForCheckCredentialForBatchInspection()
	return &CheckCredentialForBatchInspectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHealthReportTask 创建实例健康诊断任务
//
// 创建实例健康诊断任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateHealthReportTask(request *model.CreateHealthReportTaskRequest) (*model.CreateHealthReportTaskResponse, error) {
	requestDef := GenReqDefForCreateHealthReportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthReportTaskResponse), nil
	}
}

// CreateHealthReportTaskInvoker 创建实例健康诊断任务
func (c *DasClient) CreateHealthReportTaskInvoker(request *model.CreateHealthReportTaskRequest) *CreateHealthReportTaskInvoker {
	requestDef := GenReqDefForCreateHealthReportTask()
	return &CreateHealthReportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHistoryTransactionExportTask 创建导出历史事务任务
//
// DAS收集历史事务开关打开后，支持创建一次性导出指定时间范围内的历史事务数据任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateHistoryTransactionExportTask(request *model.CreateHistoryTransactionExportTaskRequest) (*model.CreateHistoryTransactionExportTaskResponse, error) {
	requestDef := GenReqDefForCreateHistoryTransactionExportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHistoryTransactionExportTaskResponse), nil
	}
}

// CreateHistoryTransactionExportTaskInvoker 创建导出历史事务任务
func (c *DasClient) CreateHistoryTransactionExportTaskInvoker(request *model.CreateHistoryTransactionExportTaskRequest) *CreateHistoryTransactionExportTaskInvoker {
	requestDef := GenReqDefForCreateHistoryTransactionExportTask()
	return &CreateHistoryTransactionExportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSnapshots 创建快照
//
// 创建快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateSnapshots(request *model.CreateSnapshotsRequest) (*model.CreateSnapshotsResponse, error) {
	requestDef := GenReqDefForCreateSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotsResponse), nil
	}
}

// CreateSnapshotsInvoker 创建快照
func (c *DasClient) CreateSnapshotsInvoker(request *model.CreateSnapshotsRequest) *CreateSnapshotsInvoker {
	requestDef := GenReqDefForCreateSnapshots()
	return &CreateSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSpaceAnalysisTask 创建空间分析任务
//
// 创建空间分析任务，如触发重新分析，支持MySQL和GaussDB(for MySQL)引擎
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateSpaceAnalysisTask(request *model.CreateSpaceAnalysisTaskRequest) (*model.CreateSpaceAnalysisTaskResponse, error) {
	requestDef := GenReqDefForCreateSpaceAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpaceAnalysisTaskResponse), nil
	}
}

// CreateSpaceAnalysisTaskInvoker 创建空间分析任务
func (c *DasClient) CreateSpaceAnalysisTaskInvoker(request *model.CreateSpaceAnalysisTaskRequest) *CreateSpaceAnalysisTaskInvoker {
	requestDef := GenReqDefForCreateSpaceAnalysisTask()
	return &CreateSpaceAnalysisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlLimitRules 创建SQL限流规则
//
// 添加SQL限流规则。目前仅支持MySQL和PostgreSQL数据库。
// MySQL使用限制如下：
// 1.规则举例详细说明：例如关键字是\&quot;select~a\&quot;, 含义为：select以及a为该并发控制所包含的两个关键字，~为关键字间隔符，即若执行SQL命令包含select与a两个关键字视为命中此条并发控制规则。
// 2.当SQL语句匹配多条限流规则时，优先生效最新添加的规则，之前的规则不再生效。
// 3.限流规则关键字有顺序要求，只会按顺序匹配。如：a~and~b 只会匹配 xxx a&gt;1 and b&gt;2，而不会匹配 xxx b&gt;2 and a&gt;1。
// 4.关键字可能大小写敏感，请执行 \&quot;show variables like &#39;rds_sqlfilter_case_sensitive&#39;或者到实例参数设置页面进行确认。
// 5.部分版本只读实例不允许设置限流规则，如果要设置限流规则，请到主实例上进行添加。
// 6.系统表不限制、不涉及数据查询的不限制、root账号在特定版本下不限制。
// PostgreSQL使用限制如下：
// 1.无法添加相同QUERY_ID或SQL语句的规则。
// 2.使用SQL语句添加规则时，需要确保存在数据库表，如：select * from test，需要确保数据库中有test表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateSqlLimitRules(request *model.CreateSqlLimitRulesRequest) (*model.CreateSqlLimitRulesResponse, error) {
	requestDef := GenReqDefForCreateSqlLimitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlLimitRulesResponse), nil
	}
}

// CreateSqlLimitRulesInvoker 创建SQL限流规则
func (c *DasClient) CreateSqlLimitRulesInvoker(request *model.CreateSqlLimitRulesRequest) *CreateSqlLimitRulesInvoker {
	requestDef := GenReqDefForCreateSqlLimitRules()
	return &CreateSqlLimitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTuning 执行SQL诊断
//
// 执行SQL诊断，
// 用于用户执行SQL诊断。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) CreateTuning(request *model.CreateTuningRequest) (*model.CreateTuningResponse, error) {
	requestDef := GenReqDefForCreateTuning()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTuningResponse), nil
	}
}

// CreateTuningInvoker 执行SQL诊断
func (c *DasClient) CreateTuningInvoker(request *model.CreateTuningRequest) *CreateTuningInvoker {
	requestDef := GenReqDefForCreateTuning()
	return &CreateTuningInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDbUser 删除数据库用户
//
// 删除注册在DAS里的数据库用户。此接口只是将注册的数据库用户在DAS系统里删除，不会真正删除数据库用户对象。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteDbUser(request *model.DeleteDbUserRequest) (*model.DeleteDbUserResponse, error) {
	requestDef := GenReqDefForDeleteDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbUserResponse), nil
	}
}

// DeleteDbUserInvoker 删除数据库用户
func (c *DasClient) DeleteDbUserInvoker(request *model.DeleteDbUserRequest) *DeleteDbUserInvoker {
	requestDef := GenReqDefForDeleteDbUser()
	return &DeleteDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEmailTemplate 删除邮件模板
//
// 删除邮件模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteEmailTemplate(request *model.DeleteEmailTemplateRequest) (*model.DeleteEmailTemplateResponse, error) {
	requestDef := GenReqDefForDeleteEmailTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEmailTemplateResponse), nil
	}
}

// DeleteEmailTemplateInvoker 删除邮件模板
func (c *DasClient) DeleteEmailTemplateInvoker(request *model.DeleteEmailTemplateRequest) *DeleteEmailTemplateInvoker {
	requestDef := GenReqDefForDeleteEmailTemplate()
	return &DeleteEmailTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHistoryTransactionExportTask 删除导出历史事务任务
//
// DAS收集历史事务开关打开后，删除历史事务导出任务记录对应的OBS文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteHistoryTransactionExportTask(request *model.DeleteHistoryTransactionExportTaskRequest) (*model.DeleteHistoryTransactionExportTaskResponse, error) {
	requestDef := GenReqDefForDeleteHistoryTransactionExportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHistoryTransactionExportTaskResponse), nil
	}
}

// DeleteHistoryTransactionExportTaskInvoker 删除导出历史事务任务
func (c *DasClient) DeleteHistoryTransactionExportTaskInvoker(request *model.DeleteHistoryTransactionExportTaskRequest) *DeleteHistoryTransactionExportTaskInvoker {
	requestDef := GenReqDefForDeleteHistoryTransactionExportTask()
	return &DeleteHistoryTransactionExportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceGroup 删除实例组
//
// 删除实例组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteInstanceGroup(request *model.DeleteInstanceGroupRequest) (*model.DeleteInstanceGroupResponse, error) {
	requestDef := GenReqDefForDeleteInstanceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceGroupResponse), nil
	}
}

// DeleteInstanceGroupInvoker 删除实例组
func (c *DasClient) DeleteInstanceGroupInvoker(request *model.DeleteInstanceGroupRequest) *DeleteInstanceGroupInvoker {
	requestDef := GenReqDefForDeleteInstanceGroup()
	return &DeleteInstanceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProcess 查杀会话
//
// 查杀会话。支持按照用户、数据库、会话列表查杀会话，三个条件至少指定一个。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteProcess(request *model.DeleteProcessRequest) (*model.DeleteProcessResponse, error) {
	requestDef := GenReqDefForDeleteProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProcessResponse), nil
	}
}

// DeleteProcessInvoker 查杀会话
func (c *DasClient) DeleteProcessInvoker(request *model.DeleteProcessRequest) *DeleteProcessInvoker {
	requestDef := GenReqDefForDeleteProcess()
	return &DeleteProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlLimitRules 删除SQL限流规则
//
// 删除SQL限流规则。目前仅支持MySQL和PostgreSQL数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) DeleteSqlLimitRules(request *model.DeleteSqlLimitRulesRequest) (*model.DeleteSqlLimitRulesResponse, error) {
	requestDef := GenReqDefForDeleteSqlLimitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlLimitRulesResponse), nil
	}
}

// DeleteSqlLimitRulesInvoker 删除SQL限流规则
func (c *DasClient) DeleteSqlLimitRulesInvoker(request *model.DeleteSqlLimitRulesRequest) *DeleteSqlLimitRulesInvoker {
	requestDef := GenReqDefForDeleteSqlLimitRules()
	return &DeleteSqlLimitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportFullSqlDetails 导出全量SQL明细
//
// 全量SQL开关打开后，创建SQL洞察任务，支持按节点、用户名、数据库、操作类型等导出全量SQL明细数据。该功能仅支持付费实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportFullSqlDetails(request *model.ExportFullSqlDetailsRequest) (*model.ExportFullSqlDetailsResponse, error) {
	requestDef := GenReqDefForExportFullSqlDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFullSqlDetailsResponse), nil
	}
}

// ExportFullSqlDetailsInvoker 导出全量SQL明细
func (c *DasClient) ExportFullSqlDetailsInvoker(request *model.ExportFullSqlDetailsRequest) *ExportFullSqlDetailsInvoker {
	requestDef := GenReqDefForExportFullSqlDetails()
	return &ExportFullSqlDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSlowQueryLogs 导出慢SQL数据
//
// DAS收集慢SQL开关打开后，一次性导出指定时间范围内的慢SQL数据，支持分页滚动获取。免费实例仅支持查看最近一小时数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportSlowQueryLogs(request *model.ExportSlowQueryLogsRequest) (*model.ExportSlowQueryLogsResponse, error) {
	requestDef := GenReqDefForExportSlowQueryLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowQueryLogsResponse), nil
	}
}

// ExportSlowQueryLogsInvoker 导出慢SQL数据
func (c *DasClient) ExportSlowQueryLogsInvoker(request *model.ExportSlowQueryLogsRequest) *ExportSlowQueryLogsInvoker {
	requestDef := GenReqDefForExportSlowQueryLogs()
	return &ExportSlowQueryLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSlowSqlStatistics 导出慢SQL统计数据
//
// 慢SQL开关打开后，导出慢SQL统计数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportSlowSqlStatistics(request *model.ExportSlowSqlStatisticsRequest) (*model.ExportSlowSqlStatisticsResponse, error) {
	requestDef := GenReqDefForExportSlowSqlStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowSqlStatisticsResponse), nil
	}
}

// ExportSlowSqlStatisticsInvoker 导出慢SQL统计数据
func (c *DasClient) ExportSlowSqlStatisticsInvoker(request *model.ExportSlowSqlStatisticsRequest) *ExportSlowSqlStatisticsInvoker {
	requestDef := GenReqDefForExportSlowSqlStatistics()
	return &ExportSlowSqlStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSlowSqlTemplatesDetails 导出慢SQL模板列表
//
// 慢SQL开关打开后，导出慢SQL模板列表。免费实例仅支持查看最近一小时数据。查询时间间隔最长一天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportSlowSqlTemplatesDetails(request *model.ExportSlowSqlTemplatesDetailsRequest) (*model.ExportSlowSqlTemplatesDetailsResponse, error) {
	requestDef := GenReqDefForExportSlowSqlTemplatesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowSqlTemplatesDetailsResponse), nil
	}
}

// ExportSlowSqlTemplatesDetailsInvoker 导出慢SQL模板列表
func (c *DasClient) ExportSlowSqlTemplatesDetailsInvoker(request *model.ExportSlowSqlTemplatesDetailsRequest) *ExportSlowSqlTemplatesDetailsInvoker {
	requestDef := GenReqDefForExportSlowSqlTemplatesDetails()
	return &ExportSlowSqlTemplatesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSlowSqlTrendDetails 导出慢SQL数量趋势
//
// 慢SQL开关打开后，导出慢SQL数量趋势。免费实例仅支持查看最近一小时数据。查询时间间隔最长一天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportSlowSqlTrendDetails(request *model.ExportSlowSqlTrendDetailsRequest) (*model.ExportSlowSqlTrendDetailsResponse, error) {
	requestDef := GenReqDefForExportSlowSqlTrendDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowSqlTrendDetailsResponse), nil
	}
}

// ExportSlowSqlTrendDetailsInvoker 导出慢SQL数量趋势
func (c *DasClient) ExportSlowSqlTrendDetailsInvoker(request *model.ExportSlowSqlTrendDetailsRequest) *ExportSlowSqlTrendDetailsInvoker {
	requestDef := GenReqDefForExportSlowSqlTrendDetails()
	return &ExportSlowSqlTrendDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSqlStatements 导出全量SQL
//
// 全量SQL开关打开后，一次性导出指定时间范围内的全量SQL数据，支持分页滚动获取。该功能仅支持付费实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportSqlStatements(request *model.ExportSqlStatementsRequest) (*model.ExportSqlStatementsResponse, error) {
	requestDef := GenReqDefForExportSqlStatements()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSqlStatementsResponse), nil
	}
}

// ExportSqlStatementsInvoker 导出全量SQL
func (c *DasClient) ExportSqlStatementsInvoker(request *model.ExportSqlStatementsRequest) *ExportSqlStatementsInvoker {
	requestDef := GenReqDefForExportSqlStatements()
	return &ExportSqlStatementsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportTopRiskInstances 导出TOP风险实例列表
//
// 导出TOP风险实例列表，支持查看最近24小时数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportTopRiskInstances(request *model.ExportTopRiskInstancesRequest) (*model.ExportTopRiskInstancesResponse, error) {
	requestDef := GenReqDefForExportTopRiskInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportTopRiskInstancesResponse), nil
	}
}

// ExportTopRiskInstancesInvoker 导出TOP风险实例列表
func (c *DasClient) ExportTopRiskInstancesInvoker(request *model.ExportTopRiskInstancesRequest) *ExportTopRiskInstancesInvoker {
	requestDef := GenReqDefForExportTopRiskInstances()
	return &ExportTopRiskInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportTopSqlTemplatesDetails 导出TopSQL模板列表
//
// TopSQL开关打开后，导出TopSQL模板列表。该功能仅支持付费实例。查询时间间隔最长一小时。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportTopSqlTemplatesDetails(request *model.ExportTopSqlTemplatesDetailsRequest) (*model.ExportTopSqlTemplatesDetailsResponse, error) {
	requestDef := GenReqDefForExportTopSqlTemplatesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportTopSqlTemplatesDetailsResponse), nil
	}
}

// ExportTopSqlTemplatesDetailsInvoker 导出TopSQL模板列表
func (c *DasClient) ExportTopSqlTemplatesDetailsInvoker(request *model.ExportTopSqlTemplatesDetailsRequest) *ExportTopSqlTemplatesDetailsInvoker {
	requestDef := GenReqDefForExportTopSqlTemplatesDetails()
	return &ExportTopSqlTemplatesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportTopSqlTrendDetails 导出SQL执行耗时区间数据
//
// TopSQL开关打开后，导出SQL执行耗时区间数据。该功能仅支持付费实例。查询时间间隔最长六小时。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ExportTopSqlTrendDetails(request *model.ExportTopSqlTrendDetailsRequest) (*model.ExportTopSqlTrendDetailsResponse, error) {
	requestDef := GenReqDefForExportTopSqlTrendDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportTopSqlTrendDetailsResponse), nil
	}
}

// ExportTopSqlTrendDetailsInvoker 导出SQL执行耗时区间数据
func (c *DasClient) ExportTopSqlTrendDetailsInvoker(request *model.ExportTopSqlTrendDetailsRequest) *ExportTopSqlTrendDetailsInvoker {
	requestDef := GenReqDefForExportTopSqlTrendDetails()
	return &ExportTopSqlTrendDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoIncrementUsage 查询自增配额
//
// 查询自增配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListAutoIncrementUsage(request *model.ListAutoIncrementUsageRequest) (*model.ListAutoIncrementUsageResponse, error) {
	requestDef := GenReqDefForListAutoIncrementUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoIncrementUsageResponse), nil
	}
}

// ListAutoIncrementUsageInvoker 查询自增配额
func (c *DasClient) ListAutoIncrementUsageInvoker(request *model.ListAutoIncrementUsageRequest) *ListAutoIncrementUsageInvoker {
	requestDef := GenReqDefForListAutoIncrementUsage()
	return &ListAutoIncrementUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudDbaInstances 获取DAS云DBA实例列表
//
// 获取DAS云DBA实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListCloudDbaInstances(request *model.ListCloudDbaInstancesRequest) (*model.ListCloudDbaInstancesResponse, error) {
	requestDef := GenReqDefForListCloudDbaInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudDbaInstancesResponse), nil
	}
}

// ListCloudDbaInstancesInvoker 获取DAS云DBA实例列表
func (c *DasClient) ListCloudDbaInstancesInvoker(request *model.ListCloudDbaInstancesRequest) *ListCloudDbaInstancesInvoker {
	requestDef := GenReqDefForListCloudDbaInstances()
	return &ListCloudDbaInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbNames 获取库名列表
//
// 获取库名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListDbNames(request *model.ListDbNamesRequest) (*model.ListDbNamesResponse, error) {
	requestDef := GenReqDefForListDbNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbNamesResponse), nil
	}
}

// ListDbNamesInvoker 获取库名列表
func (c *DasClient) ListDbNamesInvoker(request *model.ListDbNamesRequest) *ListDbNamesInvoker {
	requestDef := GenReqDefForListDbNames()
	return &ListDbNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbUsers 查询数据库用户列表
//
// 查询注册在DAS里的数据库用户列表，后续调用其他接口时(如查询实例会话列表接口)需要用到此接口返回的db_user_id。此接口不会返回数据库实例上的数据库用户对象。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListDbUsers(request *model.ListDbUsersRequest) (*model.ListDbUsersResponse, error) {
	requestDef := GenReqDefForListDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbUsersResponse), nil
	}
}

// ListDbUsersInvoker 查询数据库用户列表
func (c *DasClient) ListDbUsersInvoker(request *model.ListDbUsersRequest) *ListDbUsersInvoker {
	requestDef := GenReqDefForListDbUsers()
	return &ListDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEmailRecord 查询邮件推送记录
//
// 查询邮件推送记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListEmailRecord(request *model.ListEmailRecordRequest) (*model.ListEmailRecordResponse, error) {
	requestDef := GenReqDefForListEmailRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEmailRecordResponse), nil
	}
}

// ListEmailRecordInvoker 查询邮件推送记录
func (c *DasClient) ListEmailRecordInvoker(request *model.ListEmailRecordRequest) *ListEmailRecordInvoker {
	requestDef := GenReqDefForListEmailRecord()
	return &ListEmailRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEmailTemplate 查询邮件模板列表
//
// 查询邮件模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListEmailTemplate(request *model.ListEmailTemplateRequest) (*model.ListEmailTemplateResponse, error) {
	requestDef := GenReqDefForListEmailTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEmailTemplateResponse), nil
	}
}

// ListEmailTemplateInvoker 查询邮件模板列表
func (c *DasClient) ListEmailTemplateInvoker(request *model.ListEmailTemplateRequest) *ListEmailTemplateInvoker {
	requestDef := GenReqDefForListEmailTemplate()
	return &ListEmailTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFullSqlTasks 查询SQL洞察任务列表
//
// 全量SQL开关打开后，查询SQL洞察任务列表。该功能仅支持付费实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListFullSqlTasks(request *model.ListFullSqlTasksRequest) (*model.ListFullSqlTasksResponse, error) {
	requestDef := GenReqDefForListFullSqlTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFullSqlTasksResponse), nil
	}
}

// ListFullSqlTasksInvoker 查询SQL洞察任务列表
func (c *DasClient) ListFullSqlTasksInvoker(request *model.ListFullSqlTasksRequest) *ListFullSqlTasksInvoker {
	requestDef := GenReqDefForListFullSqlTasks()
	return &ListFullSqlTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHealthReportTask 查询实例健康诊断报告列表
//
// 查询实例健康诊断报告列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListHealthReportTask(request *model.ListHealthReportTaskRequest) (*model.ListHealthReportTaskResponse, error) {
	requestDef := GenReqDefForListHealthReportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthReportTaskResponse), nil
	}
}

// ListHealthReportTaskInvoker 查询实例健康诊断报告列表
func (c *DasClient) ListHealthReportTaskInvoker(request *model.ListHealthReportTaskRequest) *ListHealthReportTaskInvoker {
	requestDef := GenReqDefForListHealthReportTask()
	return &ListHealthReportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryTransactionExportTask 查询历史事务导出任务列表
//
// DAS收集历史事务开关打开后，查询历史事务导出任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListHistoryTransactionExportTask(request *model.ListHistoryTransactionExportTaskRequest) (*model.ListHistoryTransactionExportTaskResponse, error) {
	requestDef := GenReqDefForListHistoryTransactionExportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryTransactionExportTaskResponse), nil
	}
}

// ListHistoryTransactionExportTaskInvoker 查询历史事务导出任务列表
func (c *DasClient) ListHistoryTransactionExportTaskInvoker(request *model.ListHistoryTransactionExportTaskRequest) *ListHistoryTransactionExportTaskInvoker {
	requestDef := GenReqDefForListHistoryTransactionExportTask()
	return &ListHistoryTransactionExportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInnodbLocks 查询InnoDB锁等待列表
//
// 查询InnoDB锁等待列表。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInnodbLocks(request *model.ListInnodbLocksRequest) (*model.ListInnodbLocksResponse, error) {
	requestDef := GenReqDefForListInnodbLocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInnodbLocksResponse), nil
	}
}

// ListInnodbLocksInvoker 查询InnoDB锁等待列表
func (c *DasClient) ListInnodbLocksInvoker(request *model.ListInnodbLocksRequest) *ListInnodbLocksInvoker {
	requestDef := GenReqDefForListInnodbLocks()
	return &ListInnodbLocksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInspectionReport 查询巡检报告列表
//
// 查询巡检报告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInspectionReport(request *model.ListInspectionReportRequest) (*model.ListInspectionReportResponse, error) {
	requestDef := GenReqDefForListInspectionReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInspectionReportResponse), nil
	}
}

// ListInspectionReportInvoker 查询巡检报告列表
func (c *DasClient) ListInspectionReportInvoker(request *model.ListInspectionReportRequest) *ListInspectionReportInvoker {
	requestDef := GenReqDefForListInspectionReport()
	return &ListInspectionReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceDistribution 查询实例分布情况
//
// 查询实例分布情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceDistribution(request *model.ListInstanceDistributionRequest) (*model.ListInstanceDistributionResponse, error) {
	requestDef := GenReqDefForListInstanceDistribution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceDistributionResponse), nil
	}
}

// ListInstanceDistributionInvoker 查询实例分布情况
func (c *DasClient) ListInstanceDistributionInvoker(request *model.ListInstanceDistributionRequest) *ListInstanceDistributionInvoker {
	requestDef := GenReqDefForListInstanceDistribution()
	return &ListInstanceDistributionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceGroup 查询实例组列表
//
// 查询实例组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceGroup(request *model.ListInstanceGroupRequest) (*model.ListInstanceGroupResponse, error) {
	requestDef := GenReqDefForListInstanceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceGroupResponse), nil
	}
}

// ListInstanceGroupInvoker 查询实例组列表
func (c *DasClient) ListInstanceGroupInvoker(request *model.ListInstanceGroupRequest) *ListInstanceGroupInvoker {
	requestDef := GenReqDefForListInstanceGroup()
	return &ListInstanceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceMultiNodesSingleMetric 获取多节点单指标数据
//
// 获取多节点单指标数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceMultiNodesSingleMetric(request *model.ListInstanceMultiNodesSingleMetricRequest) (*model.ListInstanceMultiNodesSingleMetricResponse, error) {
	requestDef := GenReqDefForListInstanceMultiNodesSingleMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceMultiNodesSingleMetricResponse), nil
	}
}

// ListInstanceMultiNodesSingleMetricInvoker 获取多节点单指标数据
func (c *DasClient) ListInstanceMultiNodesSingleMetricInvoker(request *model.ListInstanceMultiNodesSingleMetricRequest) *ListInstanceMultiNodesSingleMetricInvoker {
	requestDef := GenReqDefForListInstanceMultiNodesSingleMetric()
	return &ListInstanceMultiNodesSingleMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceNodesInfo 获取单个实例节点信息
//
// 获取单个实例节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceNodesInfo(request *model.ListInstanceNodesInfoRequest) (*model.ListInstanceNodesInfoResponse, error) {
	requestDef := GenReqDefForListInstanceNodesInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceNodesInfoResponse), nil
	}
}

// ListInstanceNodesInfoInvoker 获取单个实例节点信息
func (c *DasClient) ListInstanceNodesInfoInvoker(request *model.ListInstanceNodesInfoRequest) *ListInstanceNodesInfoInvoker {
	requestDef := GenReqDefForListInstanceNodesInfo()
	return &ListInstanceNodesInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTopSlowLog 查询实例的TOP慢SQL列表
//
// 查询实例的TOP慢SQL列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListInstanceTopSlowLog(request *model.ListInstanceTopSlowLogRequest) (*model.ListInstanceTopSlowLogResponse, error) {
	requestDef := GenReqDefForListInstanceTopSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTopSlowLogResponse), nil
	}
}

// ListInstanceTopSlowLogInvoker 查询实例的TOP慢SQL列表
func (c *DasClient) ListInstanceTopSlowLogInvoker(request *model.ListInstanceTopSlowLogRequest) *ListInstanceTopSlowLogInvoker {
	requestDef := GenReqDefForListInstanceTopSlowLog()
	return &ListInstanceTopSlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLockBlockingDb 查询锁阻塞数据库名列表
//
// 查询锁阻塞数据库名列表。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListLockBlockingDb(request *model.ListLockBlockingDbRequest) (*model.ListLockBlockingDbResponse, error) {
	requestDef := GenReqDefForListLockBlockingDb()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLockBlockingDbResponse), nil
	}
}

// ListLockBlockingDbInvoker 查询锁阻塞数据库名列表
func (c *DasClient) ListLockBlockingDbInvoker(request *model.ListLockBlockingDbRequest) *ListLockBlockingDbInvoker {
	requestDef := GenReqDefForListLockBlockingDb()
	return &ListLockBlockingDbInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLockBlockingDetail 查询锁阻塞明细列表
//
// 查询锁阻塞明细列表。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListLockBlockingDetail(request *model.ListLockBlockingDetailRequest) (*model.ListLockBlockingDetailResponse, error) {
	requestDef := GenReqDefForListLockBlockingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLockBlockingDetailResponse), nil
	}
}

// ListLockBlockingDetailInvoker 查询锁阻塞明细列表
func (c *DasClient) ListLockBlockingDetailInvoker(request *model.ListLockBlockingDetailRequest) *ListLockBlockingDetailInvoker {
	requestDef := GenReqDefForListLockBlockingDetail()
	return &ListLockBlockingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLockBlockingRelationship 查询锁阻塞关系
//
// 查询锁阻塞关系。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListLockBlockingRelationship(request *model.ListLockBlockingRelationshipRequest) (*model.ListLockBlockingRelationshipResponse, error) {
	requestDef := GenReqDefForListLockBlockingRelationship()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLockBlockingRelationshipResponse), nil
	}
}

// ListLockBlockingRelationshipInvoker 查询锁阻塞关系
func (c *DasClient) ListLockBlockingRelationshipInvoker(request *model.ListLockBlockingRelationshipRequest) *ListLockBlockingRelationshipInvoker {
	requestDef := GenReqDefForListLockBlockingRelationship()
	return &ListLockBlockingRelationshipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetadataLocks 查询元数据锁列表
//
// 查询元数据锁列表。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListMetadataLocks(request *model.ListMetadataLocksRequest) (*model.ListMetadataLocksResponse, error) {
	requestDef := GenReqDefForListMetadataLocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetadataLocksResponse), nil
	}
}

// ListMetadataLocksInvoker 查询元数据锁列表
func (c *DasClient) ListMetadataLocksInvoker(request *model.ListMetadataLocksRequest) *ListMetadataLocksInvoker {
	requestDef := GenReqDefForListMetadataLocks()
	return &ListMetadataLocksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProcesses 查询实例会话列表
//
// 支持根据数据库、用户查询实例会话列表。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListProcesses(request *model.ListProcessesRequest) (*model.ListProcessesResponse, error) {
	requestDef := GenReqDefForListProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProcessesResponse), nil
	}
}

// ListProcessesInvoker 查询实例会话列表
func (c *DasClient) ListProcessesInvoker(request *model.ListProcessesRequest) *ListProcessesInvoker {
	requestDef := GenReqDefForListProcesses()
	return &ListProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskItems 查询资源风险实例风险项
//
// 查询资源风险实例风险项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListRiskItems(request *model.ListRiskItemsRequest) (*model.ListRiskItemsResponse, error) {
	requestDef := GenReqDefForListRiskItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskItemsResponse), nil
	}
}

// ListRiskItemsInvoker 查询资源风险实例风险项
func (c *DasClient) ListRiskItemsInvoker(request *model.ListRiskItemsRequest) *ListRiskItemsInvoker {
	requestDef := GenReqDefForListRiskItems()
	return &ListRiskItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskTrend 查询资源风险实例风险趋势
//
// 查询资源风险实例风险趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListRiskTrend(request *model.ListRiskTrendRequest) (*model.ListRiskTrendResponse, error) {
	requestDef := GenReqDefForListRiskTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskTrendResponse), nil
	}
}

// ListRiskTrendInvoker 查询资源风险实例风险趋势
func (c *DasClient) ListRiskTrendInvoker(request *model.ListRiskTrendRequest) *ListRiskTrendInvoker {
	requestDef := GenReqDefForListRiskTrend()
	return &ListRiskTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshots 查询快照列表
//
// 查询快照列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// ListSnapshotsInvoker 查询快照列表
func (c *DasClient) ListSnapshotsInvoker(request *model.ListSnapshotsRequest) *ListSnapshotsInvoker {
	requestDef := GenReqDefForListSnapshots()
	return &ListSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpaceAnalysis 获取空间分析数据列表
//
// 获取空间分析数据列表。实例级别数据来源于文件系统，库级别和表级别数据来源于information_schema.tables表。空间&amp;元数据分析最多分析10000张表，若缺少库表空间数据，可能是因为数据库实例表个数过多或者账号未保存密码。如果为保存密码，请使用用户管理接口或页面录入数据库账号。 支持MySQL、GaussDB(for MySQL)和SQLServer引擎。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSpaceAnalysis(request *model.ListSpaceAnalysisRequest) (*model.ListSpaceAnalysisResponse, error) {
	requestDef := GenReqDefForListSpaceAnalysis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpaceAnalysisResponse), nil
	}
}

// ListSpaceAnalysisInvoker 获取空间分析数据列表
func (c *DasClient) ListSpaceAnalysisInvoker(request *model.ListSpaceAnalysisRequest) *ListSpaceAnalysisInvoker {
	requestDef := GenReqDefForListSpaceAnalysis()
	return &ListSpaceAnalysisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlLimitRules 查询SQL限流规则列表
//
// 查询SQL限流规则。目前仅支持MySQL和PostgreSQL数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListSqlLimitRules(request *model.ListSqlLimitRulesRequest) (*model.ListSqlLimitRulesResponse, error) {
	requestDef := GenReqDefForListSqlLimitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlLimitRulesResponse), nil
	}
}

// ListSqlLimitRulesInvoker 查询SQL限流规则列表
func (c *DasClient) ListSqlLimitRulesInvoker(request *model.ListSqlLimitRulesRequest) *ListSqlLimitRulesInvoker {
	requestDef := GenReqDefForListSqlLimitRules()
	return &ListSqlLimitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopSlowLog 查询TOP慢SQL列表
//
// 查询TOP慢SQL列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTopSlowLog(request *model.ListTopSlowLogRequest) (*model.ListTopSlowLogResponse, error) {
	requestDef := GenReqDefForListTopSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopSlowLogResponse), nil
	}
}

// ListTopSlowLogInvoker 查询TOP慢SQL列表
func (c *DasClient) ListTopSlowLogInvoker(request *model.ListTopSlowLogRequest) *ListTopSlowLogInvoker {
	requestDef := GenReqDefForListTopSlowLog()
	return &ListTopSlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransactions 查询历史事务列表
//
// 查询历史事务列表。
// 目前仅支持MySQL实例，仅支持查看最近7天的历史事务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ListTransactions(request *model.ListTransactionsRequest) (*model.ListTransactionsResponse, error) {
	requestDef := GenReqDefForListTransactions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransactionsResponse), nil
	}
}

// ListTransactionsInvoker 查询历史事务列表
func (c *DasClient) ListTransactionsInvoker(request *model.ListTransactionsRequest) *ListTransactionsInvoker {
	requestDef := GenReqDefForListTransactions()
	return &ListTransactionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LoginBuiltInAccount 内置账号登录
//
// 内置账号登录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) LoginBuiltInAccount(request *model.LoginBuiltInAccountRequest) (*model.LoginBuiltInAccountResponse, error) {
	requestDef := GenReqDefForLoginBuiltInAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LoginBuiltInAccountResponse), nil
	}
}

// LoginBuiltInAccountInvoker 内置账号登录
func (c *DasClient) LoginBuiltInAccountInvoker(request *model.LoginBuiltInAccountRequest) *LoginBuiltInAccountInvoker {
	requestDef := GenReqDefForLoginBuiltInAccount()
	return &LoginBuiltInAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LogoffBuiltInAccount 内置账号登出
//
// 内置账号登出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) LogoffBuiltInAccount(request *model.LogoffBuiltInAccountRequest) (*model.LogoffBuiltInAccountResponse, error) {
	requestDef := GenReqDefForLogoffBuiltInAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LogoffBuiltInAccountResponse), nil
	}
}

// LogoffBuiltInAccountInvoker 内置账号登出
func (c *DasClient) LogoffBuiltInAccountInvoker(request *model.LogoffBuiltInAccountRequest) *LogoffBuiltInAccountInvoker {
	requestDef := GenReqDefForLogoffBuiltInAccount()
	return &LogoffBuiltInAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseDeadLock 一键分析死锁日志
//
// 一键分析死锁日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ParseDeadLock(request *model.ParseDeadLockRequest) (*model.ParseDeadLockResponse, error) {
	requestDef := GenReqDefForParseDeadLock()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseDeadLockResponse), nil
	}
}

// ParseDeadLockInvoker 一键分析死锁日志
func (c *DasClient) ParseDeadLockInvoker(request *model.ParseDeadLockRequest) *ParseDeadLockInvoker {
	requestDef := GenReqDefForParseDeadLock()
	return &ParseDeadLockInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseSqlLimitRules 根据原始SQL生成SQL限流关键字
//
// 根据原始SQL生成SQL限流关键字，目前支持MySQL、MariaDB、GaussDB(for MySQL)三种引擎。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ParseSqlLimitRules(request *model.ParseSqlLimitRulesRequest) (*model.ParseSqlLimitRulesResponse, error) {
	requestDef := GenReqDefForParseSqlLimitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseSqlLimitRulesResponse), nil
	}
}

// ParseSqlLimitRulesInvoker 根据原始SQL生成SQL限流关键字
func (c *DasClient) ParseSqlLimitRulesInvoker(request *model.ParseSqlLimitRulesRequest) *ParseSqlLimitRulesInvoker {
	requestDef := GenReqDefForParseSqlLimitRules()
	return &ParseSqlLimitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterDbUser 注册数据库用户
//
// 此接口是将数据库用户和密码注册进DAS系统，同时会返回一个数据库用户ID ，后续调用其他接口时（如查询实例会话列表接口）需要用到此数据库用户ID。密码为加密存储，且仅用于DAS API相关功能。此接口不会在数据库实例上创建数据库用户对象。请确保输入的用户名和密码是已经存在并且是正确的。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) RegisterDbUser(request *model.RegisterDbUserRequest) (*model.RegisterDbUserResponse, error) {
	requestDef := GenReqDefForRegisterDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDbUserResponse), nil
	}
}

// RegisterDbUserInvoker 注册数据库用户
func (c *DasClient) RegisterDbUserInvoker(request *model.RegisterDbUserRequest) *RegisterDbUserInvoker {
	requestDef := GenReqDefForRegisterDbUser()
	return &RegisterDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveCredential 保存AK/SK
//
// 保存AK/SK，用于后台任务访问OBS上传实例诊断报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SaveCredential(request *model.SaveCredentialRequest) (*model.SaveCredentialResponse, error) {
	requestDef := GenReqDefForSaveCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveCredentialResponse), nil
	}
}

// SaveCredentialInvoker 保存AK/SK
func (c *DasClient) SaveCredentialInvoker(request *model.SaveCredentialRequest) *SaveCredentialInvoker {
	requestDef := GenReqDefForSaveCredential()
	return &SaveCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveCredentialForBatchInspection 保存AK/SK
//
// 保存AK/SK，用于后台任务访问OBS上传实例诊断报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SaveCredentialForBatchInspection(request *model.SaveCredentialForBatchInspectionRequest) (*model.SaveCredentialForBatchInspectionResponse, error) {
	requestDef := GenReqDefForSaveCredentialForBatchInspection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveCredentialForBatchInspectionResponse), nil
	}
}

// SaveCredentialForBatchInspectionInvoker 保存AK/SK
func (c *DasClient) SaveCredentialForBatchInspectionInvoker(request *model.SaveCredentialForBatchInspectionRequest) *SaveCredentialForBatchInspectionInvoker {
	requestDef := GenReqDefForSaveCredentialForBatchInspection()
	return &SaveCredentialForBatchInspectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetLockBlockingSwitch 设置锁阻塞开关和保存时长
//
// 设置锁阻塞开关和保存时长，仅支持SQLServer引擎
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SetLockBlockingSwitch(request *model.SetLockBlockingSwitchRequest) (*model.SetLockBlockingSwitchResponse, error) {
	requestDef := GenReqDefForSetLockBlockingSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetLockBlockingSwitchResponse), nil
	}
}

// SetLockBlockingSwitchInvoker 设置锁阻塞开关和保存时长
func (c *DasClient) SetLockBlockingSwitchInvoker(request *model.SetLockBlockingSwitchRequest) *SetLockBlockingSwitchInvoker {
	requestDef := GenReqDefForSetLockBlockingSwitch()
	return &SetLockBlockingSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetThresholdForMetric 设置指标阈值
//
// 设置指标阈值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SetThresholdForMetric(request *model.SetThresholdForMetricRequest) (*model.SetThresholdForMetricResponse, error) {
	requestDef := GenReqDefForSetThresholdForMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetThresholdForMetricResponse), nil
	}
}

// SetThresholdForMetricInvoker 设置指标阈值
func (c *DasClient) SetThresholdForMetricInvoker(request *model.SetThresholdForMetricRequest) *SetThresholdForMetricInvoker {
	requestDef := GenReqDefForSetThresholdForMetric()
	return &SetThresholdForMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAnalysisSessionResult 查询会话分析结果
//
// 查询会话分析结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowAnalysisSessionResult(request *model.ShowAnalysisSessionResultRequest) (*model.ShowAnalysisSessionResultResponse, error) {
	requestDef := GenReqDefForShowAnalysisSessionResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAnalysisSessionResultResponse), nil
	}
}

// ShowAnalysisSessionResultInvoker 查询会话分析结果
func (c *DasClient) ShowAnalysisSessionResultInvoker(request *model.ShowAnalysisSessionResultRequest) *ShowAnalysisSessionResultInvoker {
	requestDef := GenReqDefForShowAnalysisSessionResult()
	return &ShowAnalysisSessionResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAnalysisSessionStatus 查询会话分析状态
//
// 查询会话分析状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowAnalysisSessionStatus(request *model.ShowAnalysisSessionStatusRequest) (*model.ShowAnalysisSessionStatusResponse, error) {
	requestDef := GenReqDefForShowAnalysisSessionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAnalysisSessionStatusResponse), nil
	}
}

// ShowAnalysisSessionStatusInvoker 查询会话分析状态
func (c *DasClient) ShowAnalysisSessionStatusInvoker(request *model.ShowAnalysisSessionStatusRequest) *ShowAnalysisSessionStatusInvoker {
	requestDef := GenReqDefForShowAnalysisSessionStatus()
	return &ShowAnalysisSessionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClouddbaGetSearchPathFlagNew 查询searchpath开关状态
//
// 查询searchpath开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowClouddbaGetSearchPathFlagNew(request *model.ShowClouddbaGetSearchPathFlagNewRequest) (*model.ShowClouddbaGetSearchPathFlagNewResponse, error) {
	requestDef := GenReqDefForShowClouddbaGetSearchPathFlagNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClouddbaGetSearchPathFlagNewResponse), nil
	}
}

// ShowClouddbaGetSearchPathFlagNewInvoker 查询searchpath开关状态
func (c *DasClient) ShowClouddbaGetSearchPathFlagNewInvoker(request *model.ShowClouddbaGetSearchPathFlagNewRequest) *ShowClouddbaGetSearchPathFlagNewInvoker {
	requestDef := GenReqDefForShowClouddbaGetSearchPathFlagNew()
	return &ShowClouddbaGetSearchPathFlagNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCredential 查询AK/SK
//
// 查询AK/SK。用于判断是否已保存AK/SK
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowCredential(request *model.ShowCredentialRequest) (*model.ShowCredentialResponse, error) {
	requestDef := GenReqDefForShowCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCredentialResponse), nil
	}
}

// ShowCredentialInvoker 查询AK/SK
func (c *DasClient) ShowCredentialInvoker(request *model.ShowCredentialRequest) *ShowCredentialInvoker {
	requestDef := GenReqDefForShowCredential()
	return &ShowCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDasCloudDbaPrice 开通配额询价
//
// 开通配额询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDasCloudDbaPrice(request *model.ShowDasCloudDbaPriceRequest) (*model.ShowDasCloudDbaPriceResponse, error) {
	requestDef := GenReqDefForShowDasCloudDbaPrice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDasCloudDbaPriceResponse), nil
	}
}

// ShowDasCloudDbaPriceInvoker 开通配额询价
func (c *DasClient) ShowDasCloudDbaPriceInvoker(request *model.ShowDasCloudDbaPriceRequest) *ShowDasCloudDbaPriceInvoker {
	requestDef := GenReqDefForShowDasCloudDbaPrice()
	return &ShowDasCloudDbaPriceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDasRecommendSqlLimitRule 自动推荐SQL限流规则
//
// 根据条件（包括模板所代表的sql平均时长，条数，最大执行时长，前三者混合）自动推荐SQL限流规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDasRecommendSqlLimitRule(request *model.ShowDasRecommendSqlLimitRuleRequest) (*model.ShowDasRecommendSqlLimitRuleResponse, error) {
	requestDef := GenReqDefForShowDasRecommendSqlLimitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDasRecommendSqlLimitRuleResponse), nil
	}
}

// ShowDasRecommendSqlLimitRuleInvoker 自动推荐SQL限流规则
func (c *DasClient) ShowDasRecommendSqlLimitRuleInvoker(request *model.ShowDasRecommendSqlLimitRuleRequest) *ShowDasRecommendSqlLimitRuleInvoker {
	requestDef := GenReqDefForShowDasRecommendSqlLimitRule()
	return &ShowDasRecommendSqlLimitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDbUser 查询数据库用户信息
//
// 查询注册在DAS里的数据库用户信息。此接口不能查询数据库实例上的数据库用户对象。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDbUser(request *model.ShowDbUserRequest) (*model.ShowDbUserResponse, error) {
	requestDef := GenReqDefForShowDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbUserResponse), nil
	}
}

// ShowDbUserInvoker 查询数据库用户信息
func (c *DasClient) ShowDbUserInvoker(request *model.ShowDbUserRequest) *ShowDbUserInvoker {
	requestDef := GenReqDefForShowDbUser()
	return &ShowDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockAnalysisResult 查询死锁日志分析结果
//
// 查询死锁日志分析结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockAnalysisResult(request *model.ShowDeadLockAnalysisResultRequest) (*model.ShowDeadLockAnalysisResultResponse, error) {
	requestDef := GenReqDefForShowDeadLockAnalysisResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockAnalysisResultResponse), nil
	}
}

// ShowDeadLockAnalysisResultInvoker 查询死锁日志分析结果
func (c *DasClient) ShowDeadLockAnalysisResultInvoker(request *model.ShowDeadLockAnalysisResultRequest) *ShowDeadLockAnalysisResultInvoker {
	requestDef := GenReqDefForShowDeadLockAnalysisResult()
	return &ShowDeadLockAnalysisResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockTopology 获取死锁拓扑图数据
//
// 获取死锁拓扑图数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockTopology(request *model.ShowDeadLockTopologyRequest) (*model.ShowDeadLockTopologyResponse, error) {
	requestDef := GenReqDefForShowDeadLockTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockTopologyResponse), nil
	}
}

// ShowDeadLockTopologyInvoker 获取死锁拓扑图数据
func (c *DasClient) ShowDeadLockTopologyInvoker(request *model.ShowDeadLockTopologyRequest) *ShowDeadLockTopologyInvoker {
	requestDef := GenReqDefForShowDeadLockTopology()
	return &ShowDeadLockTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExportTaskInfo 查看全量SQL导出任务详情
//
// 查看全量SQL导出任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowExportTaskInfo(request *model.ShowExportTaskInfoRequest) (*model.ShowExportTaskInfoResponse, error) {
	requestDef := GenReqDefForShowExportTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExportTaskInfoResponse), nil
	}
}

// ShowExportTaskInfoInvoker 查看全量SQL导出任务详情
func (c *DasClient) ShowExportTaskInfoInvoker(request *model.ShowExportTaskInfoRequest) *ShowExportTaskInfoInvoker {
	requestDef := GenReqDefForShowExportTaskInfo()
	return &ShowExportTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFullDeadLockList 获取全量死锁信息
//
// 获取全量死锁信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowFullDeadLockList(request *model.ShowFullDeadLockListRequest) (*model.ShowFullDeadLockListResponse, error) {
	requestDef := GenReqDefForShowFullDeadLockList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFullDeadLockListResponse), nil
	}
}

// ShowFullDeadLockListInvoker 获取全量死锁信息
func (c *DasClient) ShowFullDeadLockListInvoker(request *model.ShowFullDeadLockListRequest) *ShowFullDeadLockListInvoker {
	requestDef := GenReqDefForShowFullDeadLockList()
	return &ShowFullDeadLockListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFullDeadLockSwitch 获取全量死锁开关
//
// 获取全量死锁开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowFullDeadLockSwitch(request *model.ShowFullDeadLockSwitchRequest) (*model.ShowFullDeadLockSwitchResponse, error) {
	requestDef := GenReqDefForShowFullDeadLockSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFullDeadLockSwitchResponse), nil
	}
}

// ShowFullDeadLockSwitchInvoker 获取全量死锁开关
func (c *DasClient) ShowFullDeadLockSwitchInvoker(request *model.ShowFullDeadLockSwitchRequest) *ShowFullDeadLockSwitchInvoker {
	requestDef := GenReqDefForShowFullDeadLockSwitch()
	return &ShowFullDeadLockSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFullDeadLockSwitchNew 获取全量死锁开关
//
// 获取全量死锁开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowFullDeadLockSwitchNew(request *model.ShowFullDeadLockSwitchNewRequest) (*model.ShowFullDeadLockSwitchNewResponse, error) {
	requestDef := GenReqDefForShowFullDeadLockSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFullDeadLockSwitchNewResponse), nil
	}
}

// ShowFullDeadLockSwitchNewInvoker 获取全量死锁开关
func (c *DasClient) ShowFullDeadLockSwitchNewInvoker(request *model.ShowFullDeadLockSwitchNewRequest) *ShowFullDeadLockSwitchNewInvoker {
	requestDef := GenReqDefForShowFullDeadLockSwitchNew()
	return &ShowFullDeadLockSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalPrivacyNew 获取产品级别的安全协议
//
// 获取产品级别的安全协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowGlobalPrivacyNew(request *model.ShowGlobalPrivacyNewRequest) (*model.ShowGlobalPrivacyNewResponse, error) {
	requestDef := GenReqDefForShowGlobalPrivacyNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalPrivacyNewResponse), nil
	}
}

// ShowGlobalPrivacyNewInvoker 获取产品级别的安全协议
func (c *DasClient) ShowGlobalPrivacyNewInvoker(request *model.ShowGlobalPrivacyNewRequest) *ShowGlobalPrivacyNewInvoker {
	requestDef := GenReqDefForShowGlobalPrivacyNew()
	return &ShowGlobalPrivacyNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthReportSettings 查看实例诊断报告设置
//
// 查看实例诊断报告设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowHealthReportSettings(request *model.ShowHealthReportSettingsRequest) (*model.ShowHealthReportSettingsResponse, error) {
	requestDef := GenReqDefForShowHealthReportSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthReportSettingsResponse), nil
	}
}

// ShowHealthReportSettingsInvoker 查看实例诊断报告设置
func (c *DasClient) ShowHealthReportSettingsInvoker(request *model.ShowHealthReportSettingsRequest) *ShowHealthReportSettingsInvoker {
	requestDef := GenReqDefForShowHealthReportSettings()
	return &ShowHealthReportSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHistoryTransactionExportTaskInfo 查询历史事务导出任务详情
//
// DAS收集历史事务开关打开后，查询历史事务导出任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowHistoryTransactionExportTaskInfo(request *model.ShowHistoryTransactionExportTaskInfoRequest) (*model.ShowHistoryTransactionExportTaskInfoResponse, error) {
	requestDef := GenReqDefForShowHistoryTransactionExportTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTransactionExportTaskInfoResponse), nil
	}
}

// ShowHistoryTransactionExportTaskInfoInvoker 查询历史事务导出任务详情
func (c *DasClient) ShowHistoryTransactionExportTaskInfoInvoker(request *model.ShowHistoryTransactionExportTaskInfoRequest) *ShowHistoryTransactionExportTaskInfoInvoker {
	requestDef := GenReqDefForShowHistoryTransactionExportTaskInfo()
	return &ShowHistoryTransactionExportTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHistoryTransactionSwitchNew 查询历史事务开关
//
// 查询历史事务开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowHistoryTransactionSwitchNew(request *model.ShowHistoryTransactionSwitchNewRequest) (*model.ShowHistoryTransactionSwitchNewResponse, error) {
	requestDef := GenReqDefForShowHistoryTransactionSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTransactionSwitchNewResponse), nil
	}
}

// ShowHistoryTransactionSwitchNewInvoker 查询历史事务开关
func (c *DasClient) ShowHistoryTransactionSwitchNewInvoker(request *model.ShowHistoryTransactionSwitchNewRequest) *ShowHistoryTransactionSwitchNewInvoker {
	requestDef := GenReqDefForShowHistoryTransactionSwitchNew()
	return &ShowHistoryTransactionSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIndexUsageSwitchNew 查询索引使用开关
//
// 查询索引使用开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowIndexUsageSwitchNew(request *model.ShowIndexUsageSwitchNewRequest) (*model.ShowIndexUsageSwitchNewResponse, error) {
	requestDef := GenReqDefForShowIndexUsageSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndexUsageSwitchNewResponse), nil
	}
}

// ShowIndexUsageSwitchNewInvoker 查询索引使用开关
func (c *DasClient) ShowIndexUsageSwitchNewInvoker(request *model.ShowIndexUsageSwitchNewRequest) *ShowIndexUsageSwitchNewInvoker {
	requestDef := GenReqDefForShowIndexUsageSwitchNew()
	return &ShowIndexUsageSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceHealthReport 获取实例健康诊断报告内容
//
// 获取实例健康诊断报告内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowInstanceHealthReport(request *model.ShowInstanceHealthReportRequest) (*model.ShowInstanceHealthReportResponse, error) {
	requestDef := GenReqDefForShowInstanceHealthReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceHealthReportResponse), nil
	}
}

// ShowInstanceHealthReportInvoker 获取实例健康诊断报告内容
func (c *DasClient) ShowInstanceHealthReportInvoker(request *model.ShowInstanceHealthReportRequest) *ShowInstanceHealthReportInvoker {
	requestDef := GenReqDefForShowInstanceHealthReport()
	return &ShowInstanceHealthReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKillProcessTaskSwitch 查询自治限流开关
//
// 查询自治限流开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowKillProcessTaskSwitch(request *model.ShowKillProcessTaskSwitchRequest) (*model.ShowKillProcessTaskSwitchResponse, error) {
	requestDef := GenReqDefForShowKillProcessTaskSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKillProcessTaskSwitchResponse), nil
	}
}

// ShowKillProcessTaskSwitchInvoker 查询自治限流开关
func (c *DasClient) ShowKillProcessTaskSwitchInvoker(request *model.ShowKillProcessTaskSwitchRequest) *ShowKillProcessTaskSwitchInvoker {
	requestDef := GenReqDefForShowKillProcessTaskSwitch()
	return &ShowKillProcessTaskSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestDeadLockSnapshot 获取死锁的快照信息
//
// 获取死锁的快照信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLatestDeadLockSnapshot(request *model.ShowLatestDeadLockSnapshotRequest) (*model.ShowLatestDeadLockSnapshotResponse, error) {
	requestDef := GenReqDefForShowLatestDeadLockSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestDeadLockSnapshotResponse), nil
	}
}

// ShowLatestDeadLockSnapshotInvoker 获取死锁的快照信息
func (c *DasClient) ShowLatestDeadLockSnapshotInvoker(request *model.ShowLatestDeadLockSnapshotRequest) *ShowLatestDeadLockSnapshotInvoker {
	requestDef := GenReqDefForShowLatestDeadLockSnapshot()
	return &ShowLatestDeadLockSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestInstanceHealthReport 获取最新的数据库健康日报内容
//
// 获取最新的数据库健康日报内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLatestInstanceHealthReport(request *model.ShowLatestInstanceHealthReportRequest) (*model.ShowLatestInstanceHealthReportResponse, error) {
	requestDef := GenReqDefForShowLatestInstanceHealthReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestInstanceHealthReportResponse), nil
	}
}

// ShowLatestInstanceHealthReportInvoker 获取最新的数据库健康日报内容
func (c *DasClient) ShowLatestInstanceHealthReportInvoker(request *model.ShowLatestInstanceHealthReportRequest) *ShowLatestInstanceHealthReportInvoker {
	requestDef := GenReqDefForShowLatestInstanceHealthReport()
	return &ShowLatestInstanceHealthReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLockBlockingStatistics 查询锁阻塞数量统计
//
// 查询锁阻塞数量统计。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLockBlockingStatistics(request *model.ShowLockBlockingStatisticsRequest) (*model.ShowLockBlockingStatisticsResponse, error) {
	requestDef := GenReqDefForShowLockBlockingStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLockBlockingStatisticsResponse), nil
	}
}

// ShowLockBlockingStatisticsInvoker 查询锁阻塞数量统计
func (c *DasClient) ShowLockBlockingStatisticsInvoker(request *model.ShowLockBlockingStatisticsRequest) *ShowLockBlockingStatisticsInvoker {
	requestDef := GenReqDefForShowLockBlockingStatistics()
	return &ShowLockBlockingStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLockBlockingSwitch 查询锁阻塞开关和保存时长
//
// 查询锁阻塞开关和保存时长。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLockBlockingSwitch(request *model.ShowLockBlockingSwitchRequest) (*model.ShowLockBlockingSwitchResponse, error) {
	requestDef := GenReqDefForShowLockBlockingSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLockBlockingSwitchResponse), nil
	}
}

// ShowLockBlockingSwitchInvoker 查询锁阻塞开关和保存时长
func (c *DasClient) ShowLockBlockingSwitchInvoker(request *model.ShowLockBlockingSwitchRequest) *ShowLockBlockingSwitchInvoker {
	requestDef := GenReqDefForShowLockBlockingSwitch()
	return &ShowLockBlockingSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLockBlockingTrend 查询锁阻塞趋势列表
//
// 查询锁阻塞趋势列表。
// 仅支持SQLServer实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLockBlockingTrend(request *model.ShowLockBlockingTrendRequest) (*model.ShowLockBlockingTrendResponse, error) {
	requestDef := GenReqDefForShowLockBlockingTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLockBlockingTrendResponse), nil
	}
}

// ShowLockBlockingTrendInvoker 查询锁阻塞趋势列表
func (c *DasClient) ShowLockBlockingTrendInvoker(request *model.ShowLockBlockingTrendRequest) *ShowLockBlockingTrendInvoker {
	requestDef := GenReqDefForShowLockBlockingTrend()
	return &ShowLockBlockingTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLongHistoryTransactionSwitchNew 查询长事务开关
//
// 查询长事务开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowLongHistoryTransactionSwitchNew(request *model.ShowLongHistoryTransactionSwitchNewRequest) (*model.ShowLongHistoryTransactionSwitchNewResponse, error) {
	requestDef := GenReqDefForShowLongHistoryTransactionSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLongHistoryTransactionSwitchNewResponse), nil
	}
}

// ShowLongHistoryTransactionSwitchNewInvoker 查询长事务开关
func (c *DasClient) ShowLongHistoryTransactionSwitchNewInvoker(request *model.ShowLongHistoryTransactionSwitchNewRequest) *ShowLongHistoryTransactionSwitchNewInvoker {
	requestDef := GenReqDefForShowLongHistoryTransactionSwitchNew()
	return &ShowLongHistoryTransactionSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricNamesSupport 多节点单指标支持指标信息
//
// 多节点单指标支持指标信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowMetricNamesSupport(request *model.ShowMetricNamesSupportRequest) (*model.ShowMetricNamesSupportResponse, error) {
	requestDef := GenReqDefForShowMetricNamesSupport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricNamesSupportResponse), nil
	}
}

// ShowMetricNamesSupportInvoker 多节点单指标支持指标信息
func (c *DasClient) ShowMetricNamesSupportInvoker(request *model.ShowMetricNamesSupportRequest) *ShowMetricNamesSupportInvoker {
	requestDef := GenReqDefForShowMetricNamesSupport()
	return &ShowMetricNamesSupportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNameList 查看库名列表
//
// 查看库名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowNameList(request *model.ShowNameListRequest) (*model.ShowNameListResponse, error) {
	requestDef := GenReqDefForShowNameList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNameListResponse), nil
	}
}

// ShowNameListInvoker 查看库名列表
func (c *DasClient) ShowNameListInvoker(request *model.ShowNameListRequest) *ShowNameListInvoker {
	requestDef := GenReqDefForShowNameList()
	return &ShowNameListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询云DBA配额
//
// 查询云DBA配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询云DBA配额
func (c *DasClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSlowLogSwitchNew 查询慢日志开关
//
// 查询慢日志开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSlowLogSwitchNew(request *model.ShowSlowLogSwitchNewRequest) (*model.ShowSlowLogSwitchNewResponse, error) {
	requestDef := GenReqDefForShowSlowLogSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowLogSwitchNewResponse), nil
	}
}

// ShowSlowLogSwitchNewInvoker 查询慢日志开关
func (c *DasClient) ShowSlowLogSwitchNewInvoker(request *model.ShowSlowLogSwitchNewRequest) *ShowSlowLogSwitchNewInvoker {
	requestDef := GenReqDefForShowSlowLogSwitchNew()
	return &ShowSlowLogSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlExecutionPlan 查询SQL执行计划
//
// 查询SQL执行计划。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlExecutionPlan(request *model.ShowSqlExecutionPlanRequest) (*model.ShowSqlExecutionPlanResponse, error) {
	requestDef := GenReqDefForShowSqlExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlExecutionPlanResponse), nil
	}
}

// ShowSqlExecutionPlanInvoker 查询SQL执行计划
func (c *DasClient) ShowSqlExecutionPlanInvoker(request *model.ShowSqlExecutionPlanRequest) *ShowSqlExecutionPlanInvoker {
	requestDef := GenReqDefForShowSqlExecutionPlan()
	return &ShowSqlExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlExplain 查询SQL执行计划
//
// 查询SQL执行计划。
// 目前仅支持MySQL实例。
// 补充GET请求，处理超长SQL
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlExplain(request *model.ShowSqlExplainRequest) (*model.ShowSqlExplainResponse, error) {
	requestDef := GenReqDefForShowSqlExplain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlExplainResponse), nil
	}
}

// ShowSqlExplainInvoker 查询SQL执行计划
func (c *DasClient) ShowSqlExplainInvoker(request *model.ShowSqlExplainRequest) *ShowSqlExplainInvoker {
	requestDef := GenReqDefForShowSqlExplain()
	return &ShowSqlExplainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlLimitJobInfo 查询SQL限流任务
//
// 查询指定ID的SQL限流任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlLimitJobInfo(request *model.ShowSqlLimitJobInfoRequest) (*model.ShowSqlLimitJobInfoResponse, error) {
	requestDef := GenReqDefForShowSqlLimitJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlLimitJobInfoResponse), nil
	}
}

// ShowSqlLimitJobInfoInvoker 查询SQL限流任务
func (c *DasClient) ShowSqlLimitJobInfoInvoker(request *model.ShowSqlLimitJobInfoRequest) *ShowSqlLimitJobInfoInvoker {
	requestDef := GenReqDefForShowSqlLimitJobInfo()
	return &ShowSqlLimitJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlLimitSwitchStatus 查看SQL限流开关状态
//
// 查询SQL限流的开关状态。目前仅支持MySQL实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlLimitSwitchStatus(request *model.ShowSqlLimitSwitchStatusRequest) (*model.ShowSqlLimitSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowSqlLimitSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlLimitSwitchStatusResponse), nil
	}
}

// ShowSqlLimitSwitchStatusInvoker 查看SQL限流开关状态
func (c *DasClient) ShowSqlLimitSwitchStatusInvoker(request *model.ShowSqlLimitSwitchStatusRequest) *ShowSqlLimitSwitchStatusInvoker {
	requestDef := GenReqDefForShowSqlLimitSwitchStatus()
	return &ShowSqlLimitSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlLimitingSwitchNew 查询SQL限流开关
//
// 查询SQL限流开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlLimitingSwitchNew(request *model.ShowSqlLimitingSwitchNewRequest) (*model.ShowSqlLimitingSwitchNewResponse, error) {
	requestDef := GenReqDefForShowSqlLimitingSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlLimitingSwitchNewResponse), nil
	}
}

// ShowSqlLimitingSwitchNewInvoker 查询SQL限流开关
func (c *DasClient) ShowSqlLimitingSwitchNewInvoker(request *model.ShowSqlLimitingSwitchNewRequest) *ShowSqlLimitingSwitchNewInvoker {
	requestDef := GenReqDefForShowSqlLimitingSwitchNew()
	return &ShowSqlLimitingSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlSwitchStatus 查询全量SQL和慢SQL的开关状态
//
// 查询DAS收集全量SQL和慢SQL的开关状态。该功能仅支持付费实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSqlSwitchStatus(request *model.ShowSqlSwitchStatusRequest) (*model.ShowSqlSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowSqlSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlSwitchStatusResponse), nil
	}
}

// ShowSqlSwitchStatusInvoker 查询全量SQL和慢SQL的开关状态
func (c *DasClient) ShowSqlSwitchStatusInvoker(request *model.ShowSqlSwitchStatusRequest) *ShowSqlSwitchStatusInvoker {
	requestDef := GenReqDefForShowSqlSwitchStatus()
	return &ShowSqlSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSupportedEngines 查看支持的引擎类型
//
// 查看支持的引擎类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowSupportedEngines(request *model.ShowSupportedEnginesRequest) (*model.ShowSupportedEnginesResponse, error) {
	requestDef := GenReqDefForShowSupportedEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSupportedEnginesResponse), nil
	}
}

// ShowSupportedEnginesInvoker 查看支持的引擎类型
func (c *DasClient) ShowSupportedEnginesInvoker(request *model.ShowSupportedEnginesRequest) *ShowSupportedEnginesInvoker {
	requestDef := GenReqDefForShowSupportedEngines()
	return &ShowSupportedEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransactionSwitchStatus 查询历史事务开关
//
// 查询历史事务开关。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowTransactionSwitchStatus(request *model.ShowTransactionSwitchStatusRequest) (*model.ShowTransactionSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowTransactionSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransactionSwitchStatusResponse), nil
	}
}

// ShowTransactionSwitchStatusInvoker 查询历史事务开关
func (c *DasClient) ShowTransactionSwitchStatusInvoker(request *model.ShowTransactionSwitchStatusRequest) *ShowTransactionSwitchStatusInvoker {
	requestDef := GenReqDefForShowTransactionSwitchStatus()
	return &ShowTransactionSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTuning 获取诊断结果
//
// 获取诊断结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowTuning(request *model.ShowTuningRequest) (*model.ShowTuningResponse, error) {
	requestDef := GenReqDefForShowTuning()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTuningResponse), nil
	}
}

// ShowTuningInvoker 获取诊断结果
func (c *DasClient) ShowTuningInvoker(request *model.ShowTuningRequest) *ShowTuningInvoker {
	requestDef := GenReqDefForShowTuning()
	return &ShowTuningInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWhetherUseCloudDba 判断该实例能否使用云DBA功能
//
// 判断该实例能否使用云DBA功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowWhetherUseCloudDba(request *model.ShowWhetherUseCloudDbaRequest) (*model.ShowWhetherUseCloudDbaResponse, error) {
	requestDef := GenReqDefForShowWhetherUseCloudDba()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWhetherUseCloudDbaResponse), nil
	}
}

// ShowWhetherUseCloudDbaInvoker 判断该实例能否使用云DBA功能
func (c *DasClient) ShowWhetherUseCloudDbaInvoker(request *model.ShowWhetherUseCloudDbaRequest) *ShowWhetherUseCloudDbaInvoker {
	requestDef := GenReqDefForShowWhetherUseCloudDba()
	return &ShowWhetherUseCloudDbaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAnalysisSession 开始会话分析
//
// 开始会话分析
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) StartAnalysisSession(request *model.StartAnalysisSessionRequest) (*model.StartAnalysisSessionResponse, error) {
	requestDef := GenReqDefForStartAnalysisSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAnalysisSessionResponse), nil
	}
}

// StartAnalysisSessionInvoker 开始会话分析
func (c *DasClient) StartAnalysisSessionInvoker(request *model.StartAnalysisSessionRequest) *StartAnalysisSessionInvoker {
	requestDef := GenReqDefForStartAnalysisSession()
	return &StartAnalysisSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SynchronizeInstances 同步实例列表
//
// 同步实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SynchronizeInstances(request *model.SynchronizeInstancesRequest) (*model.SynchronizeInstancesResponse, error) {
	requestDef := GenReqDefForSynchronizeInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SynchronizeInstancesResponse), nil
	}
}

// SynchronizeInstancesInvoker 同步实例列表
func (c *DasClient) SynchronizeInstancesInvoker(request *model.SynchronizeInstancesRequest) *SynchronizeInstancesInvoker {
	requestDef := GenReqDefForSynchronizeInstances()
	return &SynchronizeInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDbUser 修改数据库用户
//
// 修改注册在DAS里的数据库用户名和密码。此接口不会修改数据库实例上的数据库用户对象的用户名和密码。请确保输入的用户名和密码是已经存在并且是正确的。
// 目前仅支持MySQL实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateDbUser(request *model.UpdateDbUserRequest) (*model.UpdateDbUserResponse, error) {
	requestDef := GenReqDefForUpdateDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbUserResponse), nil
	}
}

// UpdateDbUserInvoker 修改数据库用户
func (c *DasClient) UpdateDbUserInvoker(request *model.UpdateDbUserRequest) *UpdateDbUserInvoker {
	requestDef := GenReqDefForUpdateDbUser()
	return &UpdateDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEmailTemplate 修改邮件模板
//
// 修改邮件模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateEmailTemplate(request *model.UpdateEmailTemplateRequest) (*model.UpdateEmailTemplateResponse, error) {
	requestDef := GenReqDefForUpdateEmailTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEmailTemplateResponse), nil
	}
}

// UpdateEmailTemplateInvoker 修改邮件模板
func (c *DasClient) UpdateEmailTemplateInvoker(request *model.UpdateEmailTemplateRequest) *UpdateEmailTemplateInvoker {
	requestDef := GenReqDefForUpdateEmailTemplate()
	return &UpdateEmailTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFullSqlSwitch 全量SQL开关
//
// 全量SQL开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateFullSqlSwitch(request *model.UpdateFullSqlSwitchRequest) (*model.UpdateFullSqlSwitchResponse, error) {
	requestDef := GenReqDefForUpdateFullSqlSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFullSqlSwitchResponse), nil
	}
}

// UpdateFullSqlSwitchInvoker 全量SQL开关
func (c *DasClient) UpdateFullSqlSwitchInvoker(request *model.UpdateFullSqlSwitchRequest) *UpdateFullSqlSwitchInvoker {
	requestDef := GenReqDefForUpdateFullSqlSwitch()
	return &UpdateFullSqlSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthReportSettings 更新实例诊断报告设置
//
// 更新实例诊断报告设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateHealthReportSettings(request *model.UpdateHealthReportSettingsRequest) (*model.UpdateHealthReportSettingsResponse, error) {
	requestDef := GenReqDefForUpdateHealthReportSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthReportSettingsResponse), nil
	}
}

// UpdateHealthReportSettingsInvoker 更新实例诊断报告设置
func (c *DasClient) UpdateHealthReportSettingsInvoker(request *model.UpdateHealthReportSettingsRequest) *UpdateHealthReportSettingsInvoker {
	requestDef := GenReqDefForUpdateHealthReportSettings()
	return &UpdateHealthReportSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceGroup 修改实例组
//
// 修改实例组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateInstanceGroup(request *model.UpdateInstanceGroupRequest) (*model.UpdateInstanceGroupResponse, error) {
	requestDef := GenReqDefForUpdateInstanceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceGroupResponse), nil
	}
}

// UpdateInstanceGroupInvoker 修改实例组
func (c *DasClient) UpdateInstanceGroupInvoker(request *model.UpdateInstanceGroupRequest) *UpdateInstanceGroupInvoker {
	requestDef := GenReqDefForUpdateInstanceGroup()
	return &UpdateInstanceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSqlLimitRules 修改SQL限流规则
//
// 修改SQL限流规则。目前仅支持PostgreSQL数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) UpdateSqlLimitRules(request *model.UpdateSqlLimitRulesRequest) (*model.UpdateSqlLimitRulesResponse, error) {
	requestDef := GenReqDefForUpdateSqlLimitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlLimitRulesResponse), nil
	}
}

// UpdateSqlLimitRulesInvoker 修改SQL限流规则
func (c *DasClient) UpdateSqlLimitRulesInvoker(request *model.UpdateSqlLimitRulesRequest) *UpdateSqlLimitRulesInvoker {
	requestDef := GenReqDefForUpdateSqlLimitRules()
	return &UpdateSqlLimitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeadLockSwitchNew 查询死锁开关状态
//
// 查询死锁开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) ShowDeadLockSwitchNew(request *model.ShowDeadLockSwitchNewRequest) (*model.ShowDeadLockSwitchNewResponse, error) {
	requestDef := GenReqDefForShowDeadLockSwitchNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeadLockSwitchNewResponse), nil
	}
}

// ShowDeadLockSwitchNewInvoker 查询死锁开关状态
func (c *DasClient) ShowDeadLockSwitchNewInvoker(request *model.ShowDeadLockSwitchNewRequest) *ShowDeadLockSwitchNewInvoker {
	requestDef := GenReqDefForShowDeadLockSwitchNew()
	return &ShowDeadLockSwitchNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchFullsqlSwitch 开启/关闭全量SQL开关
//
// 开启/关闭全量SQL开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DasClient) SwitchFullsqlSwitch(request *model.SwitchFullsqlSwitchRequest) (*model.SwitchFullsqlSwitchResponse, error) {
	requestDef := GenReqDefForSwitchFullsqlSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchFullsqlSwitchResponse), nil
	}
}

// SwitchFullsqlSwitchInvoker 开启/关闭全量SQL开关
func (c *DasClient) SwitchFullsqlSwitchInvoker(request *model.SwitchFullsqlSwitchRequest) *SwitchFullsqlSwitchInvoker {
	requestDef := GenReqDefForSwitchFullsqlSwitch()
	return &SwitchFullsqlSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
