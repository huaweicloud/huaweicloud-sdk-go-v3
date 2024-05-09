package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/drs/v3/model"
)

type DrsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDrsClient(hcClient *httpclient.HcHttpClient) *DrsClient {
	return &DrsClient{HcClient: hcClient}
}

func DrsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchChangeData 批量数据加工
//
// 数据复制服务支持对同步的对象进行加工，即可以为选择的对象添加规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchChangeData(request *model.BatchChangeDataRequest) (*model.BatchChangeDataResponse, error) {
	requestDef := GenReqDefForBatchChangeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeDataResponse), nil
	}
}

// BatchChangeDataInvoker 批量数据加工
func (c *DrsClient) BatchChangeDataInvoker(request *model.BatchChangeDataRequest) *BatchChangeDataInvoker {
	requestDef := GenReqDefForBatchChangeData()
	return &BatchChangeDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckJobs 批量预检查
//
// 批量预检查，校验是否可进行迁移。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchCheckJobs(request *model.BatchCheckJobsRequest) (*model.BatchCheckJobsResponse, error) {
	requestDef := GenReqDefForBatchCheckJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckJobsResponse), nil
	}
}

// BatchCheckJobsInvoker 批量预检查
func (c *DrsClient) BatchCheckJobsInvoker(request *model.BatchCheckJobsRequest) *BatchCheckJobsInvoker {
	requestDef := GenReqDefForBatchCheckJobs()
	return &BatchCheckJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckResults 批量查询预检查结果
//
// 批量查询任务的预检查结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchCheckResults(request *model.BatchCheckResultsRequest) (*model.BatchCheckResultsResponse, error) {
	requestDef := GenReqDefForBatchCheckResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckResultsResponse), nil
	}
}

// BatchCheckResultsInvoker 批量查询预检查结果
func (c *DrsClient) BatchCheckResultsInvoker(request *model.BatchCheckResultsRequest) *BatchCheckResultsInvoker {
	requestDef := GenReqDefForBatchCheckResults()
	return &BatchCheckResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateJobs 批量创建任务
//
// 根据请求参数不同，可以批量创建实时迁移、实时同步、实时灾备任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchCreateJobs(request *model.BatchCreateJobsRequest) (*model.BatchCreateJobsResponse, error) {
	requestDef := GenReqDefForBatchCreateJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateJobsResponse), nil
	}
}

// BatchCreateJobsInvoker 批量创建任务
func (c *DrsClient) BatchCreateJobsInvoker(request *model.BatchCreateJobsRequest) *BatchCreateJobsInvoker {
	requestDef := GenReqDefForBatchCreateJobs()
	return &BatchCreateJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteJobs 批量结束任务或删除任务
//
// 批量结束任务或删除实时迁移、实时同步、实时灾备任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

// BatchDeleteJobsInvoker 批量结束任务或删除任务
func (c *DrsClient) BatchDeleteJobsInvoker(request *model.BatchDeleteJobsRequest) *BatchDeleteJobsInvoker {
	requestDef := GenReqDefForBatchDeleteJobs()
	return &BatchDeleteJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListJobDetails 批量查询任务详情
//
// 根据任务ID批量查询任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListJobDetails(request *model.BatchListJobDetailsRequest) (*model.BatchListJobDetailsResponse, error) {
	requestDef := GenReqDefForBatchListJobDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobDetailsResponse), nil
	}
}

// BatchListJobDetailsInvoker 批量查询任务详情
func (c *DrsClient) BatchListJobDetailsInvoker(request *model.BatchListJobDetailsRequest) *BatchListJobDetailsInvoker {
	requestDef := GenReqDefForBatchListJobDetails()
	return &BatchListJobDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListJobStatus 批量查询任务状态
//
// 根据任务ID批量查询任务状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListJobStatus(request *model.BatchListJobStatusRequest) (*model.BatchListJobStatusResponse, error) {
	requestDef := GenReqDefForBatchListJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobStatusResponse), nil
	}
}

// BatchListJobStatusInvoker 批量查询任务状态
func (c *DrsClient) BatchListJobStatusInvoker(request *model.BatchListJobStatusRequest) *BatchListJobStatusInvoker {
	requestDef := GenReqDefForBatchListJobStatus()
	return &BatchListJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListProgresses 批量查询任务进度
//
// 根据任务ID批量查询全量进度、增量时延信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListProgresses(request *model.BatchListProgressesRequest) (*model.BatchListProgressesResponse, error) {
	requestDef := GenReqDefForBatchListProgresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListProgressesResponse), nil
	}
}

// BatchListProgressesInvoker 批量查询任务进度
func (c *DrsClient) BatchListProgressesInvoker(request *model.BatchListProgressesRequest) *BatchListProgressesInvoker {
	requestDef := GenReqDefForBatchListProgresses()
	return &BatchListProgressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListRposAndRtos 批量查询RPO和RTO
//
// 批量查询RPO和RTO。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListRposAndRtos(request *model.BatchListRposAndRtosRequest) (*model.BatchListRposAndRtosResponse, error) {
	requestDef := GenReqDefForBatchListRposAndRtos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListRposAndRtosResponse), nil
	}
}

// BatchListRposAndRtosInvoker 批量查询RPO和RTO
func (c *DrsClient) BatchListRposAndRtosInvoker(request *model.BatchListRposAndRtosRequest) *BatchListRposAndRtosInvoker {
	requestDef := GenReqDefForBatchListRposAndRtos()
	return &BatchListRposAndRtosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListStructDetail 批量查询灾备初始化对象详情
//
// 根据任务ID批量查询灾备初始化对象详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListStructDetail(request *model.BatchListStructDetailRequest) (*model.BatchListStructDetailResponse, error) {
	requestDef := GenReqDefForBatchListStructDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructDetailResponse), nil
	}
}

// BatchListStructDetailInvoker 批量查询灾备初始化对象详情
func (c *DrsClient) BatchListStructDetailInvoker(request *model.BatchListStructDetailRequest) *BatchListStructDetailInvoker {
	requestDef := GenReqDefForBatchListStructDetail()
	return &BatchListStructDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListStructProcess 批量查询灾备初始化进度
//
// 根据任务ID批量查询灾备初始化进度，虚拟id不支持查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchListStructProcess(request *model.BatchListStructProcessRequest) (*model.BatchListStructProcessResponse, error) {
	requestDef := GenReqDefForBatchListStructProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructProcessResponse), nil
	}
}

// BatchListStructProcessInvoker 批量查询灾备初始化进度
func (c *DrsClient) BatchListStructProcessInvoker(request *model.BatchListStructProcessRequest) *BatchListStructProcessInvoker {
	requestDef := GenReqDefForBatchListStructProcess()
	return &BatchListStructProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResetPassword 批量修改源库/目标库密码
//
// 任务启动之后需要修改源库/目标库密码时调用此接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchResetPassword(request *model.BatchResetPasswordRequest) (*model.BatchResetPasswordResponse, error) {
	requestDef := GenReqDefForBatchResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetPasswordResponse), nil
	}
}

// BatchResetPasswordInvoker 批量修改源库/目标库密码
func (c *DrsClient) BatchResetPasswordInvoker(request *model.BatchResetPasswordRequest) *BatchResetPasswordInvoker {
	requestDef := GenReqDefForBatchResetPassword()
	return &BatchResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRestoreTask 批量续传/重试
//
// 在迁移过程中由于不确定因素导致迁移任务失败，可通过重试功能，重新提交迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchRestoreTask(request *model.BatchRestoreTaskRequest) (*model.BatchRestoreTaskResponse, error) {
	requestDef := GenReqDefForBatchRestoreTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestoreTaskResponse), nil
	}
}

// BatchRestoreTaskInvoker 批量续传/重试
func (c *DrsClient) BatchRestoreTaskInvoker(request *model.BatchRestoreTaskRequest) *BatchRestoreTaskInvoker {
	requestDef := GenReqDefForBatchRestoreTask()
	return &BatchRestoreTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetDefiner 批量设置definer
//
// 批量设置Definer迁移是否迁移到到该用户下。
// - 选择是：迁移后，所有源数据库对象的Definer都会迁移至该用户下，其他用户需要授权后才具有数据库对象权限。
// - 选择否：迁移后，将保持源数据库对象Definer定义不变，选择此选项，需要配合下一步用户权限迁移功能，将源数据库的用户全部迁移，这样才能保持源数据库的权限体系完全不变。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSetDefiner(request *model.BatchSetDefinerRequest) (*model.BatchSetDefinerResponse, error) {
	requestDef := GenReqDefForBatchSetDefiner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetDefinerResponse), nil
	}
}

// BatchSetDefinerInvoker 批量设置definer
func (c *DrsClient) BatchSetDefinerInvoker(request *model.BatchSetDefinerRequest) *BatchSetDefinerInvoker {
	requestDef := GenReqDefForBatchSetDefiner()
	return &BatchSetDefinerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetObjects 批量数据库对象选择
//
// 迁移之前，选择需要迁移的数据库或者表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSetObjects(request *model.BatchSetObjectsRequest) (*model.BatchSetObjectsResponse, error) {
	requestDef := GenReqDefForBatchSetObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetObjectsResponse), nil
	}
}

// BatchSetObjectsInvoker 批量数据库对象选择
func (c *DrsClient) BatchSetObjectsInvoker(request *model.BatchSetObjectsRequest) *BatchSetObjectsInvoker {
	requestDef := GenReqDefForBatchSetObjects()
	return &BatchSetObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetPolicy 批量设置同步策略
//
// - 批量设置同步策略，包括冲突策略、过滤DROP Datase、对象同步范围。
// - 设置kafka同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSetPolicy(request *model.BatchSetPolicyRequest) (*model.BatchSetPolicyResponse, error) {
	requestDef := GenReqDefForBatchSetPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetPolicyResponse), nil
	}
}

// BatchSetPolicyInvoker 批量设置同步策略
func (c *DrsClient) BatchSetPolicyInvoker(request *model.BatchSetPolicyRequest) *BatchSetPolicyInvoker {
	requestDef := GenReqDefForBatchSetPolicy()
	return &BatchSetPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetSmn 批量配置异常通知
//
// 批量设置异常通知，已结束的任务不支持设置。
// - 支持选择已有的SMN主题和手动输入手机号、邮箱两种方式，具体根据自己使用情况选择
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSetSmn(request *model.BatchSetSmnRequest) (*model.BatchSetSmnResponse, error) {
	requestDef := GenReqDefForBatchSetSmn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSmnResponse), nil
	}
}

// BatchSetSmnInvoker 批量配置异常通知
func (c *DrsClient) BatchSetSmnInvoker(request *model.BatchSetSmnRequest) *BatchSetSmnInvoker {
	requestDef := GenReqDefForBatchSetSmn()
	return &BatchSetSmnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetSpeed 批量设置任务限速
//
// 批量设置任务限速，任务创建成功后默认不限速。
// - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。
// - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSetSpeed(request *model.BatchSetSpeedRequest) (*model.BatchSetSpeedResponse, error) {
	requestDef := GenReqDefForBatchSetSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSpeedResponse), nil
	}
}

// BatchSetSpeedInvoker 批量设置任务限速
func (c *DrsClient) BatchSetSpeedInvoker(request *model.BatchSetSpeedRequest) *BatchSetSpeedInvoker {
	requestDef := GenReqDefForBatchSetSpeed()
	return &BatchSetSpeedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowParams 批量获取数据库参数
//
// 在进行数据库迁移时，为了确保迁移成功后业务应用的使用不受影响，数据复制服务提供了参数对比功能帮助您进行源库和目标库参数一致性对比，此接口可以获取源库和目标库的数据库参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchShowParams(request *model.BatchShowParamsRequest) (*model.BatchShowParamsResponse, error) {
	requestDef := GenReqDefForBatchShowParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowParamsResponse), nil
	}
}

// BatchShowParamsInvoker 批量获取数据库参数
func (c *DrsClient) BatchShowParamsInvoker(request *model.BatchShowParamsRequest) *BatchShowParamsInvoker {
	requestDef := GenReqDefForBatchShowParams()
	return &BatchShowParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartJobs 批量启动任务
//
// 批量启动实时迁移、同步、灾备任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchStartJobs(request *model.BatchStartJobsRequest) (*model.BatchStartJobsResponse, error) {
	requestDef := GenReqDefForBatchStartJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartJobsResponse), nil
	}
}

// BatchStartJobsInvoker 批量启动任务
func (c *DrsClient) BatchStartJobsInvoker(request *model.BatchStartJobsRequest) *BatchStartJobsInvoker {
	requestDef := GenReqDefForBatchStartJobs()
	return &BatchStartJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopJobs 批量暂停任务
//
// 批量暂停任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchStopJobs(request *model.BatchStopJobsRequest) (*model.BatchStopJobsResponse, error) {
	requestDef := GenReqDefForBatchStopJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopJobsResponse), nil
	}
}

// BatchStopJobsInvoker 批量暂停任务
func (c *DrsClient) BatchStopJobsInvoker(request *model.BatchStopJobsRequest) *BatchStopJobsInvoker {
	requestDef := GenReqDefForBatchStopJobs()
	return &BatchStopJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSwitchover 批量主备倒换
//
// 批量主备倒换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchSwitchover(request *model.BatchSwitchoverRequest) (*model.BatchSwitchoverResponse, error) {
	requestDef := GenReqDefForBatchSwitchover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSwitchoverResponse), nil
	}
}

// BatchSwitchoverInvoker 批量主备倒换
func (c *DrsClient) BatchSwitchoverInvoker(request *model.BatchSwitchoverRequest) *BatchSwitchoverInvoker {
	requestDef := GenReqDefForBatchSwitchover()
	return &BatchSwitchoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateJob 批量修改任务
//
// 批量修改任务名称或描述，设置异常通知信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchUpdateJob(request *model.BatchUpdateJobRequest) (*model.BatchUpdateJobResponse, error) {
	requestDef := GenReqDefForBatchUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateJobResponse), nil
	}
}

// BatchUpdateJobInvoker 批量修改任务
func (c *DrsClient) BatchUpdateJobInvoker(request *model.BatchUpdateJobRequest) *BatchUpdateJobInvoker {
	requestDef := GenReqDefForBatchUpdateJob()
	return &BatchUpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateUser 批量更新迁移用户信息
//
// 数据库的迁移过程中，迁移用户需要进行单独处理，该接口可以批量设置需要迁移的用户和角色。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchUpdateUser(request *model.BatchUpdateUserRequest) (*model.BatchUpdateUserResponse, error) {
	requestDef := GenReqDefForBatchUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUserResponse), nil
	}
}

// BatchUpdateUserInvoker 批量更新迁移用户信息
func (c *DrsClient) BatchUpdateUserInvoker(request *model.BatchUpdateUserRequest) *BatchUpdateUserInvoker {
	requestDef := GenReqDefForBatchUpdateUser()
	return &BatchUpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchValidateClustersConnections 批量测试连接-集群模式
//
// - 批量测试连接（集群模式）。
// - 主备任务测试连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchValidateClustersConnections(request *model.BatchValidateClustersConnectionsRequest) (*model.BatchValidateClustersConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateClustersConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateClustersConnectionsResponse), nil
	}
}

// BatchValidateClustersConnectionsInvoker 批量测试连接-集群模式
func (c *DrsClient) BatchValidateClustersConnectionsInvoker(request *model.BatchValidateClustersConnectionsRequest) *BatchValidateClustersConnectionsInvoker {
	requestDef := GenReqDefForBatchValidateClustersConnections()
	return &BatchValidateClustersConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchValidateConnections 批量测试连接
//
// 批量测试连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchValidateConnections(request *model.BatchValidateConnectionsRequest) (*model.BatchValidateConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateConnectionsResponse), nil
	}
}

// BatchValidateConnectionsInvoker 批量测试连接
func (c *DrsClient) BatchValidateConnectionsInvoker(request *model.BatchValidateConnectionsRequest) *BatchValidateConnectionsInvoker {
	requestDef := GenReqDefForBatchValidateConnections()
	return &BatchValidateConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCompareResultFile 导出对比任务结果文件
//
// 导出对比任务结果文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateCompareResultFile(request *model.CreateCompareResultFileRequest) (*model.CreateCompareResultFileResponse, error) {
	requestDef := GenReqDefForCreateCompareResultFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompareResultFileResponse), nil
	}
}

// CreateCompareResultFileInvoker 导出对比任务结果文件
func (c *DrsClient) CreateCompareResultFileInvoker(request *model.CreateCompareResultFileRequest) *CreateCompareResultFileInvoker {
	requestDef := GenReqDefForCreateCompareResultFile()
	return &CreateCompareResultFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCompareTask 创建对比任务
//
// 创建对比任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateCompareTask(request *model.CreateCompareTaskRequest) (*model.CreateCompareTaskResponse, error) {
	requestDef := GenReqDefForCreateCompareTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompareTaskResponse), nil
	}
}

// CreateCompareTaskInvoker 创建对比任务
func (c *DrsClient) CreateCompareTaskInvoker(request *model.CreateCompareTaskRequest) *CreateCompareTaskInvoker {
	requestDef := GenReqDefForCreateCompareTask()
	return &CreateCompareTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataLevelTableCompareJob 创建数据级表对比任务
//
// 创建数据级表对比任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateDataLevelTableCompareJob(request *model.CreateDataLevelTableCompareJobRequest) (*model.CreateDataLevelTableCompareJobResponse, error) {
	requestDef := GenReqDefForCreateDataLevelTableCompareJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataLevelTableCompareJobResponse), nil
	}
}

// CreateDataLevelTableCompareJobInvoker 创建数据级表对比任务
func (c *DrsClient) CreateDataLevelTableCompareJobInvoker(request *model.CreateDataLevelTableCompareJobRequest) *CreateDataLevelTableCompareJobInvoker {
	requestDef := GenReqDefForCreateDataLevelTableCompareJob()
	return &CreateDataLevelTableCompareJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateObjectLevelCompareJob 创建对象级对比任务
//
// 创建对象级对比任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateObjectLevelCompareJob(request *model.CreateObjectLevelCompareJobRequest) (*model.CreateObjectLevelCompareJobResponse, error) {
	requestDef := GenReqDefForCreateObjectLevelCompareJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateObjectLevelCompareJobResponse), nil
	}
}

// CreateObjectLevelCompareJobInvoker 创建对象级对比任务
func (c *DrsClient) CreateObjectLevelCompareJobInvoker(request *model.CreateObjectLevelCompareJobRequest) *CreateObjectLevelCompareJobInvoker {
	requestDef := GenReqDefForCreateObjectLevelCompareJob()
	return &CreateObjectLevelCompareJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCompareJob 取消对比任务
//
// 取消对比任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteCompareJob(request *model.DeleteCompareJobRequest) (*model.DeleteCompareJobResponse, error) {
	requestDef := GenReqDefForDeleteCompareJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCompareJobResponse), nil
	}
}

// DeleteCompareJobInvoker 取消对比任务
func (c *DrsClient) DeleteCompareJobInvoker(request *model.DeleteCompareJobRequest) *DeleteCompareJobInvoker {
	requestDef := GenReqDefForDeleteCompareJob()
	return &DeleteCompareJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadCompareResultFile 下载对比任务结果文件
//
// 下载对比任务结果文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DownloadCompareResultFile(request *model.DownloadCompareResultFileRequest) (*model.DownloadCompareResultFileResponse, error) {
	requestDef := GenReqDefForDownloadCompareResultFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadCompareResultFileResponse), nil
	}
}

// DownloadCompareResultFileInvoker 下载对比任务结果文件
func (c *DrsClient) DownloadCompareResultFileInvoker(request *model.DownloadCompareResultFileRequest) *DownloadCompareResultFileInvoker {
	requestDef := GenReqDefForDownloadCompareResultFile()
	return &DownloadCompareResultFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableNodeTypes 查询可用的Node规格
//
// 查询可用的Node规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListAvailableNodeTypes(request *model.ListAvailableNodeTypesRequest) (*model.ListAvailableNodeTypesResponse, error) {
	requestDef := GenReqDefForListAvailableNodeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableNodeTypesResponse), nil
	}
}

// ListAvailableNodeTypesInvoker 查询可用的Node规格
func (c *DrsClient) ListAvailableNodeTypesInvoker(request *model.ListAvailableNodeTypesRequest) *ListAvailableNodeTypesInvoker {
	requestDef := GenReqDefForListAvailableNodeTypes()
	return &ListAvailableNodeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZone 查询规格未售罄的可用区
//
// 查询规格未售罄的可用区
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListAvailableZone(request *model.ListAvailableZoneRequest) (*model.ListAvailableZoneResponse, error) {
	requestDef := GenReqDefForListAvailableZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZoneResponse), nil
	}
}

// ListAvailableZoneInvoker 查询规格未售罄的可用区
func (c *DrsClient) ListAvailableZoneInvoker(request *model.ListAvailableZoneRequest) *ListAvailableZoneInvoker {
	requestDef := GenReqDefForListAvailableZone()
	return &ListAvailableZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCompareResult 查询对比结果
//
// 查询对比结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListCompareResult(request *model.ListCompareResultRequest) (*model.ListCompareResultResponse, error) {
	requestDef := GenReqDefForListCompareResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompareResultResponse), nil
	}
}

// ListCompareResultInvoker 查询对比结果
func (c *DrsClient) ListCompareResultInvoker(request *model.ListCompareResultRequest) *ListCompareResultInvoker {
	requestDef := GenReqDefForListCompareResult()
	return &ListCompareResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContentCompareDetail 查询内容对比详情
//
// 查询内容对比详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListContentCompareDetail(request *model.ListContentCompareDetailRequest) (*model.ListContentCompareDetailResponse, error) {
	requestDef := GenReqDefForListContentCompareDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContentCompareDetailResponse), nil
	}
}

// ListContentCompareDetailInvoker 查询内容对比详情
func (c *DrsClient) ListContentCompareDetailInvoker(request *model.ListContentCompareDetailRequest) *ListContentCompareDetailInvoker {
	requestDef := GenReqDefForListContentCompareDetail()
	return &ListContentCompareDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContentCompareDifference 查询内容对比差异
//
// 查询内容对比差异。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListContentCompareDifference(request *model.ListContentCompareDifferenceRequest) (*model.ListContentCompareDifferenceResponse, error) {
	requestDef := GenReqDefForListContentCompareDifference()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContentCompareDifferenceResponse), nil
	}
}

// ListContentCompareDifferenceInvoker 查询内容对比差异
func (c *DrsClient) ListContentCompareDifferenceInvoker(request *model.ListContentCompareDifferenceRequest) *ListContentCompareDifferenceInvoker {
	requestDef := GenReqDefForListContentCompareDifference()
	return &ListContentCompareDifferenceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContentCompareOverview 查询内容对比总览
//
// 查询内容对比总览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListContentCompareOverview(request *model.ListContentCompareOverviewRequest) (*model.ListContentCompareOverviewResponse, error) {
	requestDef := GenReqDefForListContentCompareOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContentCompareOverviewResponse), nil
	}
}

// ListContentCompareOverviewInvoker 查询内容对比总览
func (c *DrsClient) ListContentCompareOverviewInvoker(request *model.ListContentCompareOverviewRequest) *ListContentCompareOverviewInvoker {
	requestDef := GenReqDefForListContentCompareOverview()
	return &ListContentCompareOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataCompareDetail 查询行数对比详情
//
// 查询行数对比详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListDataCompareDetail(request *model.ListDataCompareDetailRequest) (*model.ListDataCompareDetailResponse, error) {
	requestDef := GenReqDefForListDataCompareDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataCompareDetailResponse), nil
	}
}

// ListDataCompareDetailInvoker 查询行数对比详情
func (c *DrsClient) ListDataCompareDetailInvoker(request *model.ListDataCompareDetailRequest) *ListDataCompareDetailInvoker {
	requestDef := GenReqDefForListDataCompareDetail()
	return &ListDataCompareDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataCompareOverview 查询行数对比总览
//
// 查询行数对比总览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListDataCompareOverview(request *model.ListDataCompareOverviewRequest) (*model.ListDataCompareOverviewResponse, error) {
	requestDef := GenReqDefForListDataCompareOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataCompareOverviewResponse), nil
	}
}

// ListDataCompareOverviewInvoker 查询行数对比总览
func (c *DrsClient) ListDataCompareOverviewInvoker(request *model.ListDataCompareOverviewRequest) *ListDataCompareOverviewInvoker {
	requestDef := GenReqDefForListDataCompareOverview()
	return &ListDataCompareOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataLevelTableCompareJobs 查询数据级表对比任务列表
//
// 查询数据级表对比任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListDataLevelTableCompareJobs(request *model.ListDataLevelTableCompareJobsRequest) (*model.ListDataLevelTableCompareJobsResponse, error) {
	requestDef := GenReqDefForListDataLevelTableCompareJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataLevelTableCompareJobsResponse), nil
	}
}

// ListDataLevelTableCompareJobsInvoker 查询数据级表对比任务列表
func (c *DrsClient) ListDataLevelTableCompareJobsInvoker(request *model.ListDataLevelTableCompareJobsRequest) *ListDataLevelTableCompareJobsInvoker {
	requestDef := GenReqDefForListDataLevelTableCompareJobs()
	return &ListDataLevelTableCompareJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObejectLevelCompareDetail 查询对象对比任务详情
//
// 查询对象对比任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListObejectLevelCompareDetail(request *model.ListObejectLevelCompareDetailRequest) (*model.ListObejectLevelCompareDetailResponse, error) {
	requestDef := GenReqDefForListObejectLevelCompareDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObejectLevelCompareDetailResponse), nil
	}
}

// ListObejectLevelCompareDetailInvoker 查询对象对比任务详情
func (c *DrsClient) ListObejectLevelCompareDetailInvoker(request *model.ListObejectLevelCompareDetailRequest) *ListObejectLevelCompareDetailInvoker {
	requestDef := GenReqDefForListObejectLevelCompareDetail()
	return &ListObejectLevelCompareDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObejectLevelCompareOverview 查询对象对比任务概览
//
// 查询对象对比任务概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListObejectLevelCompareOverview(request *model.ListObejectLevelCompareOverviewRequest) (*model.ListObejectLevelCompareOverviewResponse, error) {
	requestDef := GenReqDefForListObejectLevelCompareOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObejectLevelCompareOverviewResponse), nil
	}
}

// ListObejectLevelCompareOverviewInvoker 查询对象对比任务概览
func (c *DrsClient) ListObejectLevelCompareOverviewInvoker(request *model.ListObejectLevelCompareOverviewRequest) *ListObejectLevelCompareOverviewInvoker {
	requestDef := GenReqDefForListObejectLevelCompareOverview()
	return &ListObejectLevelCompareOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 获取源库迁移用户列表
//
// 数据库的迁移过程中，迁移用户需要进行单独处理，该接口可以查询源库的用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 获取源库迁移用户列表
func (c *DrsClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobList 查询租户任务列表
//
// 查询租户任务列表，可以根据引擎类型，网络类型，任务状态，任务名称，任务ID进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowJobList(request *model.ShowJobListRequest) (*model.ShowJobListResponse, error) {
	requestDef := GenReqDefForShowJobList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobListResponse), nil
	}
}

// ShowJobListInvoker 查询租户任务列表
func (c *DrsClient) ShowJobListInvoker(request *model.ShowJobListRequest) *ShowJobListInvoker {
	requestDef := GenReqDefForShowJobList()
	return &ShowJobListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitoringData 查询容灾监控数据
//
// 根据任务ID查询容灾监控数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowMonitoringData(request *model.ShowMonitoringDataRequest) (*model.ShowMonitoringDataResponse, error) {
	requestDef := GenReqDefForShowMonitoringData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitoringDataResponse), nil
	}
}

// ShowMonitoringDataInvoker 查询容灾监控数据
func (c *DrsClient) ShowMonitoringDataInvoker(request *model.ShowMonitoringDataRequest) *ShowMonitoringDataInvoker {
	requestDef := GenReqDefForShowMonitoringData()
	return &ShowMonitoringDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询单租户在某一项目下DRS服务下的配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *DrsClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartPromptlyDataLevelTableCompareJob 立即启动数据级表对比任务
//
// 立即启动数据级表对比任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) StartPromptlyDataLevelTableCompareJob(request *model.StartPromptlyDataLevelTableCompareJobRequest) (*model.StartPromptlyDataLevelTableCompareJobResponse, error) {
	requestDef := GenReqDefForStartPromptlyDataLevelTableCompareJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPromptlyDataLevelTableCompareJobResponse), nil
	}
}

// StartPromptlyDataLevelTableCompareJobInvoker 立即启动数据级表对比任务
func (c *DrsClient) StartPromptlyDataLevelTableCompareJobInvoker(request *model.StartPromptlyDataLevelTableCompareJobRequest) *StartPromptlyDataLevelTableCompareJobInvoker {
	requestDef := GenReqDefForStartPromptlyDataLevelTableCompareJob()
	return &StartPromptlyDataLevelTableCompareJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateParams 修改数据库参数
//
// 修改数据库参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateParams(request *model.UpdateParamsRequest) (*model.UpdateParamsResponse, error) {
	requestDef := GenReqDefForUpdateParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateParamsResponse), nil
	}
}

// UpdateParamsInvoker 修改数据库参数
func (c *DrsClient) UpdateParamsInvoker(request *model.UpdateParamsRequest) *UpdateParamsInvoker {
	requestDef := GenReqDefForUpdateParams()
	return &UpdateParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTuningParams 高级设置
//
// 修改调优参数的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateTuningParams(request *model.UpdateTuningParamsRequest) (*model.UpdateTuningParamsResponse, error) {
	requestDef := GenReqDefForUpdateTuningParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTuningParamsResponse), nil
	}
}

// UpdateTuningParamsInvoker 高级设置
func (c *DrsClient) UpdateTuningParamsInvoker(request *model.UpdateTuningParamsRequest) *UpdateTuningParamsInvoker {
	requestDef := GenReqDefForUpdateTuningParams()
	return &UpdateTuningParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
