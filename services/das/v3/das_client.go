package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

//查询API版本列表
func (c *DasClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

//查询指定的API版本信息
func (c *DasClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

//打开或者关闭DAS收集全量SQL开关，开启后，实例的性能损耗在5%以内。开启全量SQL后，本服务会对SQL的文本内容进行存储，以便进行分析。用户可自行设置全量SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。 打开或者关闭DAS收集慢SQL开关。开启慢SQL后，本服务会对慢SQL的文本内容进行存储，以便进行分析。用户可自行设置慢SQL的保存时间范围，到期后会自动删除；如果未设置，数据默认保留7天。该功能仅支持付费实例。
func (c *DasClient) ChangeSqlSwitch(request *model.ChangeSqlSwitchRequest) (*model.ChangeSqlSwitchResponse, error) {
	requestDef := GenReqDefForChangeSqlSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSqlSwitchResponse), nil
	}
}

//创建空间分析任务，如触发重新分析，支持MySQL和GaussDB(for MySQL)引擎
func (c *DasClient) CreateSpaceAnalysisTask(request *model.CreateSpaceAnalysisTaskRequest) (*model.CreateSpaceAnalysisTaskResponse, error) {
	requestDef := GenReqDefForCreateSpaceAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpaceAnalysisTaskResponse), nil
	}
}

//删除注册在DAS里的数据库用户。此接口只是将注册的数据库用户在DAS系统里删除，不会真正删除数据库用户对象。 目前仅支持MySQL实例。
func (c *DasClient) DeleteDbUser(request *model.DeleteDbUserRequest) (*model.DeleteDbUserResponse, error) {
	requestDef := GenReqDefForDeleteDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbUserResponse), nil
	}
}

//查杀会话。支持按照用户、数据库、会话列表查杀会话，三个条件至少指定一个。 目前仅支持MySQL实例。
func (c *DasClient) DeleteProcess(request *model.DeleteProcessRequest) (*model.DeleteProcessResponse, error) {
	requestDef := GenReqDefForDeleteProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProcessResponse), nil
	}
}

//DAS收集慢SQL开关打开后，一次性导出指定时间范围内的慢SQL数据，支持分页滚动获取。该功能仅支持付费实例。
func (c *DasClient) ExportSlowQueryLogs(request *model.ExportSlowQueryLogsRequest) (*model.ExportSlowQueryLogsResponse, error) {
	requestDef := GenReqDefForExportSlowQueryLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowQueryLogsResponse), nil
	}
}

//全量SQL开关打开后，一次性导出指定时间范围内的全量SQL数据，支持分页滚动获取。该功能仅支持付费实例。
func (c *DasClient) ExportSqlStatements(request *model.ExportSqlStatementsRequest) (*model.ExportSqlStatementsResponse, error) {
	requestDef := GenReqDefForExportSqlStatements()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSqlStatementsResponse), nil
	}
}

//查询注册在DAS里的数据库用户列表，后续调用其他接口时(如查询实例会话列表接口)需要用到此接口返回的db_user_id。此接口不会返回数据库实例上的数据库用户对象。 目前仅支持MySQL实例。
func (c *DasClient) ListDbUsers(request *model.ListDbUsersRequest) (*model.ListDbUsersResponse, error) {
	requestDef := GenReqDefForListDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbUsersResponse), nil
	}
}

//查询InnoDB锁等待列表。 目前仅支持MySQL实例。
func (c *DasClient) ListInnodbLocks(request *model.ListInnodbLocksRequest) (*model.ListInnodbLocksResponse, error) {
	requestDef := GenReqDefForListInnodbLocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInnodbLocksResponse), nil
	}
}

//查询元数据锁列表。 目前仅支持MySQL实例。
func (c *DasClient) ListMetadataLocks(request *model.ListMetadataLocksRequest) (*model.ListMetadataLocksResponse, error) {
	requestDef := GenReqDefForListMetadataLocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetadataLocksResponse), nil
	}
}

//支持根据数据库、用户查询实例会话列表。 目前仅支持MySQL实例。
func (c *DasClient) ListProcesses(request *model.ListProcessesRequest) (*model.ListProcessesResponse, error) {
	requestDef := GenReqDefForListProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProcessesResponse), nil
	}
}

//获取空间分析数据列表。实例级别数据来源于文件系统，库级别和表级别数据来源于information_schema.tables表。空间&元数据分析最多分析10000张表，若缺少库表空间数据，可能是因为数据库实例表个数过多或者账号未保存密码。如果为保存密码，请使用用户管理接口或页面录入数据库账号。支持MySQL和GaussDB(for MySQL)引擎
func (c *DasClient) ListSpaceAnalysis(request *model.ListSpaceAnalysisRequest) (*model.ListSpaceAnalysisResponse, error) {
	requestDef := GenReqDefForListSpaceAnalysis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpaceAnalysisResponse), nil
	}
}

//此接口是将数据库用户和密码注册进DAS系统，同时会返回一个数据库用户ID ，后续调用其他接口时（如查询实例会话列表接口）需要用到此数据库用户ID。密码为加密存储，且仅用于DAS API相关功能。此接口不会在数据库实例上创建数据库用户对象。请确保输入的用户名和密码是已经存在并且是正确的。 目前仅支持MySQL实例。
func (c *DasClient) RegisterDbUser(request *model.RegisterDbUserRequest) (*model.RegisterDbUserResponse, error) {
	requestDef := GenReqDefForRegisterDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDbUserResponse), nil
	}
}

//查询注册在DAS里的数据库用户信息。此接口不能查询数据库实例上的数据库用户对象。 目前仅支持MySQL实例。
func (c *DasClient) ShowDbUser(request *model.ShowDbUserRequest) (*model.ShowDbUserResponse, error) {
	requestDef := GenReqDefForShowDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbUserResponse), nil
	}
}

//查询云DBA配额
func (c *DasClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

//查询SQL执行计划。 目前仅支持MySQL实例。
func (c *DasClient) ShowSqlExecutionPlan(request *model.ShowSqlExecutionPlanRequest) (*model.ShowSqlExecutionPlanResponse, error) {
	requestDef := GenReqDefForShowSqlExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlExecutionPlanResponse), nil
	}
}

//查询DAS收集全量SQL和慢SQL的开关状态。该功能仅支持付费实例。
func (c *DasClient) ShowSqlSwitchStatus(request *model.ShowSqlSwitchStatusRequest) (*model.ShowSqlSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowSqlSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlSwitchStatusResponse), nil
	}
}

//修改注册在DAS里的数据库用户名和密码。此接口不会修改数据库实例上的数据库用户对象的用户名和密码。请确保输入的用户名和密码是已经存在并且是正确的。 目前仅支持MySQL实例。
func (c *DasClient) UpdateDbUser(request *model.UpdateDbUserRequest) (*model.UpdateDbUserResponse, error) {
	requestDef := GenReqDefForUpdateDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbUserResponse), nil
	}
}
