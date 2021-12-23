package v3

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/drs/v3/model"
)

type DrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDrsClient(hcClient *http_client.HcHttpClient) *DrsClient {
	return &DrsClient{HcClient: hcClient}
}

func DrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//数据复制服务支持对同步的对象进行加工，即可以为选择的对象添加规则。
func (c *DrsClient) BatchChangeData(request *model.BatchChangeDataRequest) (*model.BatchChangeDataResponse, error) {
	requestDef := GenReqDefForBatchChangeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeDataResponse), nil
	}
}

//批量预检查，校验是否可进行迁移。
func (c *DrsClient) BatchCheckJobs(request *model.BatchCheckJobsRequest) (*model.BatchCheckJobsResponse, error) {
	requestDef := GenReqDefForBatchCheckJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckJobsResponse), nil
	}
}

//批量查询任务的预检查结果。
func (c *DrsClient) BatchCheckResults(request *model.BatchCheckResultsRequest) (*model.BatchCheckResultsResponse, error) {
	requestDef := GenReqDefForBatchCheckResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckResultsResponse), nil
	}
}

//根据请求参数不同，可以批量创建实时迁移、实时同步、实时灾备任务。
func (c *DrsClient) BatchCreateJobs(request *model.BatchCreateJobsRequest) (*model.BatchCreateJobsResponse, error) {
	requestDef := GenReqDefForBatchCreateJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateJobsResponse), nil
	}
}

//批量结束任务或删除实时迁移、实时同步、实时灾备任务。
func (c *DrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

//根据任务ID批量查询任务详情。
func (c *DrsClient) BatchListJobDetails(request *model.BatchListJobDetailsRequest) (*model.BatchListJobDetailsResponse, error) {
	requestDef := GenReqDefForBatchListJobDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobDetailsResponse), nil
	}
}

//根据任务ID批量查询任务状态。
func (c *DrsClient) BatchListJobStatus(request *model.BatchListJobStatusRequest) (*model.BatchListJobStatusResponse, error) {
	requestDef := GenReqDefForBatchListJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobStatusResponse), nil
	}
}

//根据任务ID批量查询全量进度、增量时延信息。
func (c *DrsClient) BatchListProgresses(request *model.BatchListProgressesRequest) (*model.BatchListProgressesResponse, error) {
	requestDef := GenReqDefForBatchListProgresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListProgressesResponse), nil
	}
}

//批量查询RPO和RTO。
func (c *DrsClient) BatchListRposAndRtos(request *model.BatchListRposAndRtosRequest) (*model.BatchListRposAndRtosResponse, error) {
	requestDef := GenReqDefForBatchListRposAndRtos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListRposAndRtosResponse), nil
	}
}

//根据任务ID批量查询灾备初始化对象详情。
func (c *DrsClient) BatchListStructDetail(request *model.BatchListStructDetailRequest) (*model.BatchListStructDetailResponse, error) {
	requestDef := GenReqDefForBatchListStructDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructDetailResponse), nil
	}
}

//根据任务ID批量查询灾备初始化进度，虚拟id不支持查询。
func (c *DrsClient) BatchListStructProcess(request *model.BatchListStructProcessRequest) (*model.BatchListStructProcessResponse, error) {
	requestDef := GenReqDefForBatchListStructProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructProcessResponse), nil
	}
}

//任务启动之后需要修改源库/目标库密码时调用此接口。
func (c *DrsClient) BatchResetPassword(request *model.BatchResetPasswordRequest) (*model.BatchResetPasswordResponse, error) {
	requestDef := GenReqDefForBatchResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetPasswordResponse), nil
	}
}

//在迁移过程中由于不确定因素导致迁移任务失败，可通过重试功能，重新提交迁移任务。
func (c *DrsClient) BatchRestoreTask(request *model.BatchRestoreTaskRequest) (*model.BatchRestoreTaskResponse, error) {
	requestDef := GenReqDefForBatchRestoreTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestoreTaskResponse), nil
	}
}

//批量设置Definer迁移是否迁移到到该用户下。 - 选择是：迁移后，所有源数据库对象的Definer都会迁移至该用户下，其他用户需要授权后才具有数据库对象权限。 - 选择否：迁移后，将保持源数据库对象Definer定义不变，选择此选项，需要配合下一步用户权限迁移功能，将源数据库的用户全部迁移，这样才能保持源数据库的权限体系完全不变。
func (c *DrsClient) BatchSetDefiner(request *model.BatchSetDefinerRequest) (*model.BatchSetDefinerResponse, error) {
	requestDef := GenReqDefForBatchSetDefiner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetDefinerResponse), nil
	}
}

//迁移之前，选择需要迁移的数据库或者表。
func (c *DrsClient) BatchSetObjects(request *model.BatchSetObjectsRequest) (*model.BatchSetObjectsResponse, error) {
	requestDef := GenReqDefForBatchSetObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetObjectsResponse), nil
	}
}

//批量设置MySQL同步策略，包括冲突策略、过滤DROP Datase、对象同步范围。
func (c *DrsClient) BatchSetPolicy(request *model.BatchSetPolicyRequest) (*model.BatchSetPolicyResponse, error) {
	requestDef := GenReqDefForBatchSetPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetPolicyResponse), nil
	}
}

//批量设置任务限速，任务创建成功后默认不限速。 - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。 - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
func (c *DrsClient) BatchSetSpeed(request *model.BatchSetSpeedRequest) (*model.BatchSetSpeedResponse, error) {
	requestDef := GenReqDefForBatchSetSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSpeedResponse), nil
	}
}

//在进行数据库迁移时，为了确保迁移成功后业务应用的使用不受影响，数据复制服务提供了参数对比功能帮助您进行源库和目标库参数一致性对比，此接口可以获取源库和目标库的数据库参数。
func (c *DrsClient) BatchShowParams(request *model.BatchShowParamsRequest) (*model.BatchShowParamsResponse, error) {
	requestDef := GenReqDefForBatchShowParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowParamsResponse), nil
	}
}

//批量启动实时迁移、同步、灾备任务。
func (c *DrsClient) BatchStartJobs(request *model.BatchStartJobsRequest) (*model.BatchStartJobsResponse, error) {
	requestDef := GenReqDefForBatchStartJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartJobsResponse), nil
	}
}

//批量暂停任务。
func (c *DrsClient) BatchStopJobs(request *model.BatchStopJobsRequest) (*model.BatchStopJobsResponse, error) {
	requestDef := GenReqDefForBatchStopJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopJobsResponse), nil
	}
}

//批量主备倒换。
func (c *DrsClient) BatchSwitchover(request *model.BatchSwitchoverRequest) (*model.BatchSwitchoverResponse, error) {
	requestDef := GenReqDefForBatchSwitchover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSwitchoverResponse), nil
	}
}

//批量修改任务名称或描述，设置异常通知信息。
func (c *DrsClient) BatchUpdateJob(request *model.BatchUpdateJobRequest) (*model.BatchUpdateJobResponse, error) {
	requestDef := GenReqDefForBatchUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateJobResponse), nil
	}
}

//数据库的迁移过程中，迁移用户需要进行单独处理，该接口可以批量设置需要迁移的用户和角色。
func (c *DrsClient) BatchUpdateUser(request *model.BatchUpdateUserRequest) (*model.BatchUpdateUserResponse, error) {
	requestDef := GenReqDefForBatchUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUserResponse), nil
	}
}

//批量测试连接（集群模式）。
func (c *DrsClient) BatchValidateClustersConnections(request *model.BatchValidateClustersConnectionsRequest) (*model.BatchValidateClustersConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateClustersConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateClustersConnectionsResponse), nil
	}
}

//批量测试连接。
func (c *DrsClient) BatchValidateConnections(request *model.BatchValidateConnectionsRequest) (*model.BatchValidateConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateConnectionsResponse), nil
	}
}

//创建对比任务。
func (c *DrsClient) CreateCompareTask(request *model.CreateCompareTaskRequest) (*model.CreateCompareTaskResponse, error) {
	requestDef := GenReqDefForCreateCompareTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompareTaskResponse), nil
	}
}

//查询对比结果。
func (c *DrsClient) ListCompareResult(request *model.ListCompareResultRequest) (*model.ListCompareResultResponse, error) {
	requestDef := GenReqDefForListCompareResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompareResultResponse), nil
	}
}

//数据库的迁移过程中，迁移用户需要进行单独处理，该接口可以查询源库的用户信息。
func (c *DrsClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

//查询租户任务列表，可以根据引擎类型，网络类型，任务状态，任务名称，任务ID进行查询。
func (c *DrsClient) ShowJobList(request *model.ShowJobListRequest) (*model.ShowJobListResponse, error) {
	requestDef := GenReqDefForShowJobList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobListResponse), nil
	}
}

//根据任务ID查询容灾监控数据。
func (c *DrsClient) ShowMonitoringData(request *model.ShowMonitoringDataRequest) (*model.ShowMonitoringDataResponse, error) {
	requestDef := GenReqDefForShowMonitoringData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitoringDataResponse), nil
	}
}

//查询单租户在DRS服务下的配额信息。
func (c *DrsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

//修改数据库参数。
func (c *DrsClient) UpdateParams(request *model.UpdateParamsRequest) (*model.UpdateParamsResponse, error) {
	requestDef := GenReqDefForUpdateParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateParamsResponse), nil
	}
}
