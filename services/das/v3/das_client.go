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
