package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/das/v3/model"
)

type DasClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDasClient(hcClient *http_client.HcHttpClient) *DasClient {
	return &DasClient{HcClient: hcClient}
}

func DasClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListApiVersions 查询API版本列表
//
// 查询API版本列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ChangeSqlSwitch 开启/关闭全量SQL、慢SQL开关
//
// 打开或者关闭DAS收集全量SQL开关，开启后，实例的性能损耗在5%以内。开启全量SQL后，本服务会对SQL的文本内容进行存储，以便进行分析。用户可自行设置全量SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。
// 打开或者关闭DAS收集慢SQL开关。开启慢SQL后，本服务会对慢SQL的文本内容进行存储，以便进行分析。用户可自行设置慢SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。该功能仅支持付费实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateSpaceAnalysisTask 创建空间分析任务
//
// 创建空间分析任务，如触发重新分析，支持MySQL和GaussDB(for MySQL)引擎
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteDbUser 删除数据库用户
//
// 删除注册在DAS里的数据库用户。此接口只是将注册的数据库用户在DAS系统里删除，不会真正删除数据库用户对象。
// 目前仅支持MySQL实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ExportSlowQueryLogs 导出慢SQL数据
//
// DAS收集慢SQL开关打开后，一次性导出指定时间范围内的慢SQL数据，支持分页滚动获取。该功能仅支持付费实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ExportSqlStatements 导出全量SQL
//
// 全量SQL开关打开后，一次性导出指定时间范围内的全量SQL数据，支持分页滚动获取。该功能仅支持付费实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListDbUsers 查询数据库用户列表
//
// 查询注册在DAS里的数据库用户列表，后续调用其他接口时(如查询实例会话列表接口)需要用到此接口返回的db_user_id。此接口不会返回数据库实例上的数据库用户对象。
// 目前仅支持MySQL实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 获取空间分析数据列表。实例级别数据来源于文件系统，库级别和表级别数据来源于information_schema.tables表。空间&amp;元数据分析最多分析10000张表，若缺少库表空间数据，可能是因为数据库实例表个数过多或者账号未保存密码。如果为保存密码，请使用用户管理接口或页面录入数据库账号。支持MySQL和GaussDB(for MySQL)引擎
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// RegisterDbUser 注册数据库用户
//
// 此接口是将数据库用户和密码注册进DAS系统，同时会返回一个数据库用户ID ，后续调用其他接口时（如查询实例会话列表接口）需要用到此数据库用户ID。密码为加密存储，且仅用于DAS API相关功能。此接口不会在数据库实例上创建数据库用户对象。请确保输入的用户名和密码是已经存在并且是正确的。
// 目前仅支持MySQL实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowSqlSwitchStatus 查询全量SQL和慢SQL的开关状态。
//
// 查询DAS收集全量SQL和慢SQL的开关状态。该功能仅支持付费实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DasClient) ShowSqlSwitchStatus(request *model.ShowSqlSwitchStatusRequest) (*model.ShowSqlSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowSqlSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlSwitchStatusResponse), nil
	}
}

// ShowSqlSwitchStatusInvoker 查询全量SQL和慢SQL的开关状态。
func (c *DasClient) ShowSqlSwitchStatusInvoker(request *model.ShowSqlSwitchStatusRequest) *ShowSqlSwitchStatusInvoker {
	requestDef := GenReqDefForShowSqlSwitchStatus()
	return &ShowSqlSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDbUser 修改数据库用户
//
// 修改注册在DAS里的数据库用户名和密码。此接口不会修改数据库实例上的数据库用户对象的用户名和密码。请确保输入的用户名和密码是已经存在并且是正确的。
// 目前仅支持MySQL实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
