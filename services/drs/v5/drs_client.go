package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/drs/v5/model"
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

// BatchCreateJobsAsync 批量异步创建任务
//
// 批量异步创建任务，根据请求参数不同，可以批量异步创建实时迁移、实时同步、实时灾备等任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchCreateJobsAsync(request *model.BatchCreateJobsAsyncRequest) (*model.BatchCreateJobsAsyncResponse, error) {
	requestDef := GenReqDefForBatchCreateJobsAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateJobsAsyncResponse), nil
	}
}

// BatchCreateJobsAsyncInvoker 批量异步创建任务
func (c *DrsClient) BatchCreateJobsAsyncInvoker(request *model.BatchCreateJobsAsyncRequest) *BatchCreateJobsAsyncInvoker {
	requestDef := GenReqDefForBatchCreateJobsAsync()
	return &BatchCreateJobsAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateTags 批量添加资源标签
//
// 批量添加资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量添加资源标签
func (c *DrsClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteJobsById 批量删除任务
//
// 批量删除租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchDeleteJobsById(request *model.BatchDeleteJobsByIdRequest) (*model.BatchDeleteJobsByIdResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobsById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsByIdResponse), nil
	}
}

// BatchDeleteJobsByIdInvoker 批量删除任务
func (c *DrsClient) BatchDeleteJobsByIdInvoker(request *model.BatchDeleteJobsByIdRequest) *BatchDeleteJobsByIdInvoker {
	requestDef := GenReqDefForBatchDeleteJobsById()
	return &BatchDeleteJobsByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTags 批量删除资源标签
//
// 为指定实例批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchDeleteTags(request *model.BatchDeleteTagsRequest) (*model.BatchDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagsResponse), nil
	}
}

// BatchDeleteTagsInvoker 批量删除资源标签
func (c *DrsClient) BatchDeleteTagsInvoker(request *model.BatchDeleteTagsRequest) *BatchDeleteTagsInvoker {
	requestDef := GenReqDefForBatchDeleteTags()
	return &BatchDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchExecuteJobActions 批量操作指定ID任务
//
// 批量操作租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchExecuteJobActions(request *model.BatchExecuteJobActionsRequest) (*model.BatchExecuteJobActionsResponse, error) {
	requestDef := GenReqDefForBatchExecuteJobActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchExecuteJobActionsResponse), nil
	}
}

// BatchExecuteJobActionsInvoker 批量操作指定ID任务
func (c *DrsClient) BatchExecuteJobActionsInvoker(request *model.BatchExecuteJobActionsRequest) *BatchExecuteJobActionsInvoker {
	requestDef := GenReqDefForBatchExecuteJobActions()
	return &BatchExecuteJobActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopJobsAction 批量结束任务
//
// 批量结束租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchStopJobsAction(request *model.BatchStopJobsActionRequest) (*model.BatchStopJobsActionResponse, error) {
	requestDef := GenReqDefForBatchStopJobsAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopJobsActionResponse), nil
	}
}

// BatchStopJobsActionInvoker 批量结束任务
func (c *DrsClient) BatchStopJobsActionInvoker(request *model.BatchStopJobsActionRequest) *BatchStopJobsActionInvoker {
	requestDef := GenReqDefForBatchStopJobsAction()
	return &BatchStopJobsActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchTagAction 批量添加或删除资源标签
//
// 批量添加删除资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

// BatchTagActionInvoker 批量添加或删除资源标签
func (c *DrsClient) BatchTagActionInvoker(request *model.BatchTagActionRequest) *BatchTagActionInvoker {
	requestDef := GenReqDefForBatchTagAction()
	return &BatchTagActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeToPeriod 按需转包周期
//
// DRS同步和灾备任务按需计费转包周期计费。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ChangeToPeriod(request *model.ChangeToPeriodRequest) (*model.ChangeToPeriodResponse, error) {
	requestDef := GenReqDefForChangeToPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeToPeriodResponse), nil
	}
}

// ChangeToPeriodInvoker 按需转包周期
func (c *DrsClient) ChangeToPeriodInvoker(request *model.ChangeToPeriodRequest) *ChangeToPeriodInvoker {
	requestDef := GenReqDefForChangeToPeriod()
	return &ChangeToPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDataFilter 数据过滤规则校验
//
// 数据过滤规则校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CheckDataFilter(request *model.CheckDataFilterRequest) (*model.CheckDataFilterResponse, error) {
	requestDef := GenReqDefForCheckDataFilter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDataFilterResponse), nil
	}
}

// CheckDataFilterInvoker 数据过滤规则校验
func (c *DrsClient) CheckDataFilterInvoker(request *model.CheckDataFilterRequest) *CheckDataFilterInvoker {
	requestDef := GenReqDefForCheckDataFilter()
	return &CheckDataFilterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CleanAlarms 清除DDL告警
//
// 清除DDL告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CleanAlarms(request *model.CleanAlarmsRequest) (*model.CleanAlarmsResponse, error) {
	requestDef := GenReqDefForCleanAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CleanAlarmsResponse), nil
	}
}

// CleanAlarmsInvoker 清除DDL告警
func (c *DrsClient) CleanAlarmsInvoker(request *model.CleanAlarmsRequest) *CleanAlarmsInvoker {
	requestDef := GenReqDefForCleanAlarms()
	return &CleanAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectColumns 采集指定数据库表的列信息
//
// 采集指定数据库表的列信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CollectColumns(request *model.CollectColumnsRequest) (*model.CollectColumnsResponse, error) {
	requestDef := GenReqDefForCollectColumns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectColumnsResponse), nil
	}
}

// CollectColumnsInvoker 采集指定数据库表的列信息
func (c *DrsClient) CollectColumnsInvoker(request *model.CollectColumnsRequest) *CollectColumnsInvoker {
	requestDef := GenReqDefForCollectColumns()
	return &CollectColumnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectDbObjectsAsync 提交查询数据库对象信息
//
// 提交查询数据库对象信息。例如：
// - 当type取值为source时，表示查询源库库表信息。
// - 当源库库表信息有变化时，则type取值为source，is_refresh取值为true。
// - 当已同步到目标库的库表信息过大，需要提前将数据加载到缓存中时，type取值为synchronized。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CollectDbObjectsAsync(request *model.CollectDbObjectsAsyncRequest) (*model.CollectDbObjectsAsyncResponse, error) {
	requestDef := GenReqDefForCollectDbObjectsAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectDbObjectsAsyncResponse), nil
	}
}

// CollectDbObjectsAsyncInvoker 提交查询数据库对象信息
func (c *DrsClient) CollectDbObjectsAsyncInvoker(request *model.CollectDbObjectsAsyncRequest) *CollectDbObjectsAsyncInvoker {
	requestDef := GenReqDefForCollectDbObjectsAsync()
	return &CollectDbObjectsAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectDbObjectsInfo 提交查询数据库对象信息
//
// 提交查询数据库对象信息。例如：
// - 当type取值为source时，表示查询源库库表信息。
// - 当源库库表信息有变化时，则type取值为source，is_refresh取值为true。
// - 当已同步到目标库的库表信息过大，需要提前将数据加载到缓存中时，type取值为synchronized。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CollectDbObjectsInfo(request *model.CollectDbObjectsInfoRequest) (*model.CollectDbObjectsInfoResponse, error) {
	requestDef := GenReqDefForCollectDbObjectsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectDbObjectsInfoResponse), nil
	}
}

// CollectDbObjectsInfoInvoker 提交查询数据库对象信息
func (c *DrsClient) CollectDbObjectsInfoInvoker(request *model.CollectDbObjectsInfoRequest) *CollectDbObjectsInfoInvoker {
	requestDef := GenReqDefForCollectDbObjectsInfo()
	return &CollectDbObjectsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectPositionAsync 采集数据库位点信息
//
// 采集数据库位点信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CollectPositionAsync(request *model.CollectPositionAsyncRequest) (*model.CollectPositionAsyncResponse, error) {
	requestDef := GenReqDefForCollectPositionAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectPositionAsyncResponse), nil
	}
}

// CollectPositionAsyncInvoker 采集数据库位点信息
func (c *DrsClient) CollectPositionAsyncInvoker(request *model.CollectPositionAsyncRequest) *CollectPositionAsyncInvoker {
	requestDef := GenReqDefForCollectPositionAsync()
	return &CollectPositionAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitAsyncJob 提交批量创建异步任务
//
// 提交批量创建异步任务，当批量异步任务创建或更新参数后，系统会自动开始进行参数校验，待所有任务成功完成参数校验后并且无报错时，可调用此接口开始创建DRS任务实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CommitAsyncJob(request *model.CommitAsyncJobRequest) (*model.CommitAsyncJobResponse, error) {
	requestDef := GenReqDefForCommitAsyncJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitAsyncJobResponse), nil
	}
}

// CommitAsyncJobInvoker 提交批量创建异步任务
func (c *DrsClient) CommitAsyncJobInvoker(request *model.CommitAsyncJobRequest) *CommitAsyncJobInvoker {
	requestDef := GenReqDefForCommitAsyncJob()
	return &CommitAsyncJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyJob 克隆任务
//
// DRS支持通过克隆功能，快速复制现有同步任务的配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CopyJob(request *model.CopyJobRequest) (*model.CopyJobResponse, error) {
	requestDef := GenReqDefForCopyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyJobResponse), nil
	}
}

// CopyJobInvoker 克隆任务
func (c *DrsClient) CopyJobInvoker(request *model.CopyJobRequest) *CopyJobInvoker {
	requestDef := GenReqDefForCopyJob()
	return &CopyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountInstanceByTags 查询资源实例数量
//
// 查询资源实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CountInstanceByTags(request *model.CountInstanceByTagsRequest) (*model.CountInstanceByTagsResponse, error) {
	requestDef := GenReqDefForCountInstanceByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountInstanceByTagsResponse), nil
	}
}

// CountInstanceByTagsInvoker 查询资源实例数量
func (c *DrsClient) CountInstanceByTagsInvoker(request *model.CountInstanceByTagsRequest) *CountInstanceByTagsInvoker {
	requestDef := GenReqDefForCountInstanceByTags()
	return &CountInstanceByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnection 创建连接
//
// 创建单个连接，该连接可以为线下自建库或云上RDS等，目前支持的数据库引擎包括MySQL、PostgreSQL、Oracle和MongoDB。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

// CreateConnectionInvoker 创建连接
func (c *DrsClient) CreateConnectionInvoker(request *model.CreateConnectionRequest) *CreateConnectionInvoker {
	requestDef := GenReqDefForCreateConnection()
	return &CreateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateJob 创建任务
//
// 创建单个任务，根据请求参数不同，可以创建单个实时迁移、实时同步、实时灾备等任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateJob(request *model.CreateJobRequest) (*model.CreateJobResponse, error) {
	requestDef := GenReqDefForCreateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobResponse), nil
	}
}

// CreateJobInvoker 创建任务
func (c *DrsClient) CreateJobInvoker(request *model.CreateJobRequest) *CreateJobInvoker {
	requestDef := GenReqDefForCreateJob()
	return &CreateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReplicationJob 创建备份迁移任务
//
// 该接口主要用于三种常见场景下备份迁移任务的配置。
// 备份迁移支持如下的常见场景：
// - 通过OBS桶备份文件进行全量数据迁移。
// - 通过OBS桶备份文件进行全量+增量数据迁移。
// - 通过RDS全量备份进行全量数据迁移。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) CreateReplicationJob(request *model.CreateReplicationJobRequest) (*model.CreateReplicationJobResponse, error) {
	requestDef := GenReqDefForCreateReplicationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReplicationJobResponse), nil
	}
}

// CreateReplicationJobInvoker 创建备份迁移任务
func (c *DrsClient) CreateReplicationJobInvoker(request *model.CreateReplicationJobRequest) *CreateReplicationJobInvoker {
	requestDef := GenReqDefForCreateReplicationJob()
	return &CreateReplicationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnection 删除连接
//
// 删除租户指定的连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteConnection(request *model.DeleteConnectionRequest) (*model.DeleteConnectionResponse, error) {
	requestDef := GenReqDefForDeleteConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionResponse), nil
	}
}

// DeleteConnectionInvoker 删除连接
func (c *DrsClient) DeleteConnectionInvoker(request *model.DeleteConnectionRequest) *DeleteConnectionInvoker {
	requestDef := GenReqDefForDeleteConnection()
	return &DeleteConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJdbcDriver 删除驱动文件（不再推广）
//
// 删除驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteJdbcDriver(request *model.DeleteJdbcDriverRequest) (*model.DeleteJdbcDriverResponse, error) {
	requestDef := GenReqDefForDeleteJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJdbcDriverResponse), nil
	}
}

// DeleteJdbcDriverInvoker 删除驱动文件（不再推广）
func (c *DrsClient) DeleteJdbcDriverInvoker(request *model.DeleteJdbcDriverRequest) *DeleteJdbcDriverInvoker {
	requestDef := GenReqDefForDeleteJdbcDriver()
	return &DeleteJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除指定ID任务
//
// 删除租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除指定ID任务
func (c *DrsClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteReplicationJob 删除备份迁移任务
//
// 对于已经完成的备份迁移任务，可以选择删除迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteReplicationJob(request *model.DeleteReplicationJobRequest) (*model.DeleteReplicationJobResponse, error) {
	requestDef := GenReqDefForDeleteReplicationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReplicationJobResponse), nil
	}
}

// DeleteReplicationJobInvoker 删除备份迁移任务
func (c *DrsClient) DeleteReplicationJobInvoker(request *model.DeleteReplicationJobRequest) *DeleteReplicationJobInvoker {
	requestDef := GenReqDefForDeleteReplicationJob()
	return &DeleteReplicationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserJdbcDriver 删除驱动文件
//
// 删除驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DeleteUserJdbcDriver(request *model.DeleteUserJdbcDriverRequest) (*model.DeleteUserJdbcDriverResponse, error) {
	requestDef := GenReqDefForDeleteUserJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserJdbcDriverResponse), nil
	}
}

// DeleteUserJdbcDriverInvoker 删除驱动文件
func (c *DrsClient) DeleteUserJdbcDriverInvoker(request *model.DeleteUserJdbcDriverRequest) *DeleteUserJdbcDriverInvoker {
	requestDef := GenReqDefForDeleteUserJdbcDriver()
	return &DeleteUserJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBatchCreateTemplate 下载批量导入任务模板
//
// 下载批量导入任务模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DownloadBatchCreateTemplate(request *model.DownloadBatchCreateTemplateRequest) (*model.DownloadBatchCreateTemplateResponse, error) {
	requestDef := GenReqDefForDownloadBatchCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBatchCreateTemplateResponse), nil
	}
}

// DownloadBatchCreateTemplateInvoker 下载批量导入任务模板
func (c *DrsClient) DownloadBatchCreateTemplateInvoker(request *model.DownloadBatchCreateTemplateRequest) *DownloadBatchCreateTemplateInvoker {
	requestDef := GenReqDefForDownloadBatchCreateTemplate()
	return &DownloadBatchCreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadDbObjectTemplate 对象选择（文件导入 - 模板下载）
//
// 对象选择（文件导入 - 模板下载）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) DownloadDbObjectTemplate(request *model.DownloadDbObjectTemplateRequest) (*model.DownloadDbObjectTemplateResponse, error) {
	requestDef := GenReqDefForDownloadDbObjectTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDbObjectTemplateResponse), nil
	}
}

// DownloadDbObjectTemplateInvoker 对象选择（文件导入 - 模板下载）
func (c *DrsClient) DownloadDbObjectTemplateInvoker(request *model.DownloadDbObjectTemplateRequest) *DownloadDbObjectTemplateInvoker {
	requestDef := GenReqDefForDownloadDbObjectTemplate()
	return &DownloadDbObjectTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteJobAction 操作指定ID任务
//
// 操作租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ExecuteJobAction(request *model.ExecuteJobActionRequest) (*model.ExecuteJobActionResponse, error) {
	requestDef := GenReqDefForExecuteJobAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteJobActionResponse), nil
	}
}

// ExecuteJobActionInvoker 操作指定ID任务
func (c *DrsClient) ExecuteJobActionInvoker(request *model.ExecuteJobActionRequest) *ExecuteJobActionInvoker {
	requestDef := GenReqDefForExecuteJobAction()
	return &ExecuteJobActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportOperationInfo 导出任务操作统计信息
//
// 导出指定任务操作统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ExportOperationInfo(request *model.ExportOperationInfoRequest) (*model.ExportOperationInfoResponse, error) {
	requestDef := GenReqDefForExportOperationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportOperationInfoResponse), nil
	}
}

// ExportOperationInfoInvoker 导出任务操作统计信息
func (c *DrsClient) ExportOperationInfoInvoker(request *model.ExportOperationInfoRequest) *ExportOperationInfoInvoker {
	requestDef := GenReqDefForExportOperationInfo()
	return &ExportOperationInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportBatchCreateJobs 批量导入任务
//
// 批量导入任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ImportBatchCreateJobs(request *model.ImportBatchCreateJobsRequest) (*model.ImportBatchCreateJobsResponse, error) {
	requestDef := GenReqDefForImportBatchCreateJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportBatchCreateJobsResponse), nil
	}
}

// ImportBatchCreateJobsInvoker 批量导入任务
func (c *DrsClient) ImportBatchCreateJobsInvoker(request *model.ImportBatchCreateJobsRequest) *ImportBatchCreateJobsInvoker {
	requestDef := GenReqDefForImportBatchCreateJobs()
	return &ImportBatchCreateJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAsyncJobDetail 查询指定ID批量异步任务详情
//
// 查询租户指定ID批量异步任务详情，默认为任务的“create_time”降序排序获取结果，支持分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListAsyncJobDetail(request *model.ListAsyncJobDetailRequest) (*model.ListAsyncJobDetailResponse, error) {
	requestDef := GenReqDefForListAsyncJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAsyncJobDetailResponse), nil
	}
}

// ListAsyncJobDetailInvoker 查询指定ID批量异步任务详情
func (c *DrsClient) ListAsyncJobDetailInvoker(request *model.ListAsyncJobDetailRequest) *ListAsyncJobDetailInvoker {
	requestDef := GenReqDefForListAsyncJobDetail()
	return &ListAsyncJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAsyncJobs 查询批量异步创建的任务列表
//
// 查询租户批量异步创建的任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListAsyncJobs(request *model.ListAsyncJobsRequest) (*model.ListAsyncJobsResponse, error) {
	requestDef := GenReqDefForListAsyncJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAsyncJobsResponse), nil
	}
}

// ListAsyncJobsInvoker 查询批量异步创建的任务列表
func (c *DrsClient) ListAsyncJobsInvoker(request *model.ListAsyncJobsRequest) *ListAsyncJobsInvoker {
	requestDef := GenReqDefForListAsyncJobs()
	return &ListAsyncJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 查询连接列表
//
// 查询连接列表，可根据连接类型进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 查询连接列表
func (c *DrsClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbObjects 查询数据库对象信息
//
// 查询数据库对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListDbObjects(request *model.ListDbObjectsRequest) (*model.ListDbObjectsResponse, error) {
	requestDef := GenReqDefForListDbObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbObjectsResponse), nil
	}
}

// ListDbObjectsInvoker 查询数据库对象信息
func (c *DrsClient) ListDbObjectsInvoker(request *model.ListDbObjectsRequest) *ListDbObjectsInvoker {
	requestDef := GenReqDefForListDbObjects()
	return &ListDbObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceByTags 查询资源实例列表
//
// 查询资源实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListInstanceByTags(request *model.ListInstanceByTagsRequest) (*model.ListInstanceByTagsResponse, error) {
	requestDef := GenReqDefForListInstanceByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceByTagsResponse), nil
	}
}

// ListInstanceByTagsInvoker 查询资源实例列表
func (c *DrsClient) ListInstanceByTagsInvoker(request *model.ListInstanceByTagsRequest) *ListInstanceByTagsInvoker {
	requestDef := GenReqDefForListInstanceByTags()
	return &ListInstanceByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询资源标签
func (c *DrsClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJdbcDrivers 查询驱动文件列表（不再推广）
//
// 查询驱动文件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListJdbcDrivers(request *model.ListJdbcDriversRequest) (*model.ListJdbcDriversResponse, error) {
	requestDef := GenReqDefForListJdbcDrivers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJdbcDriversResponse), nil
	}
}

// ListJdbcDriversInvoker 查询驱动文件列表（不再推广）
func (c *DrsClient) ListJdbcDriversInvoker(request *model.ListJdbcDriversRequest) *ListJdbcDriversInvoker {
	requestDef := GenReqDefForListJdbcDrivers()
	return &ListJdbcDriversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobDdls 查询增量DDL列表
//
// 查询增量DDL列表，可根据status查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListJobDdls(request *model.ListJobDdlsRequest) (*model.ListJobDdlsResponse, error) {
	requestDef := GenReqDefForListJobDdls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobDdlsResponse), nil
	}
}

// ListJobDdlsInvoker 查询增量DDL列表
func (c *DrsClient) ListJobDdlsInvoker(request *model.ListJobDdlsRequest) *ListJobDdlsInvoker {
	requestDef := GenReqDefForListJobDdls()
	return &ListJobDdlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobHistoryParameters 查询任务的参数配置修改历史
//
// 查询任务的参数配置修改历史
// - 仅engine_type为mysql、mysql-to-pgl、mysql-to-gaussdbv5、mysql-to-gaussdbv5ha、mysql-to-dws、mysql-to-taurus、mysql-to-kafka、mysql-to-elasticsearch、mysql-to-oracle且任务状态只能为配置中、全量中、增量中、全量失败、增量失败、暂停中的实时同步任务支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListJobHistoryParameters(request *model.ListJobHistoryParametersRequest) (*model.ListJobHistoryParametersResponse, error) {
	requestDef := GenReqDefForListJobHistoryParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobHistoryParametersResponse), nil
	}
}

// ListJobHistoryParametersInvoker 查询任务的参数配置修改历史
func (c *DrsClient) ListJobHistoryParametersInvoker(request *model.ListJobHistoryParametersRequest) *ListJobHistoryParametersInvoker {
	requestDef := GenReqDefForListJobHistoryParameters()
	return &ListJobHistoryParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobParameters 查询任务参数配置列表
//
// 查询任务的参数配置列表信息
// - 仅engine_type为mysql、mysql-to-pgl、mysql-to-gaussdbv5、mysql-to-gaussdbv5ha、mysql-to-dws、mysql-to-taurus、mysql-to-kafka、mysql-to-elasticsearch、mysql-to-oracle且任务状态只能为配置中、全量中、增量中、全量失败、增量失败、暂停中的实时同步任务支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListJobParameters(request *model.ListJobParametersRequest) (*model.ListJobParametersResponse, error) {
	requestDef := GenReqDefForListJobParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobParametersResponse), nil
	}
}

// ListJobParametersInvoker 查询任务参数配置列表
func (c *DrsClient) ListJobParametersInvoker(request *model.ListJobParametersRequest) *ListJobParametersInvoker {
	requestDef := GenReqDefForListJobParameters()
	return &ListJobParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询任务列表
//
// 查询租户任务列表，可以根据企业项目，引擎类型，网络类型，任务状态，任务名称，任务ID进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询任务列表
func (c *DrsClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLinks 查询可用链路信息
//
// 根据参数不同，可查询实时迁移、实时同步、实时灾备等可用链路信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListLinks(request *model.ListLinksRequest) (*model.ListLinksResponse, error) {
	requestDef := GenReqDefForListLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLinksResponse), nil
	}
}

// ListLinksInvoker 查询可用链路信息
func (c *DrsClient) ListLinksInvoker(request *model.ListLinksRequest) *ListLinksInvoker {
	requestDef := GenReqDefForListLinks()
	return &ListLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询指定project ID下不同任务类型的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *DrsClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReplicationJobs 查询备份迁移任务列表
//
// 获取当前备份迁移任务列表，不包含已删除的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListReplicationJobs(request *model.ListReplicationJobsRequest) (*model.ListReplicationJobsResponse, error) {
	requestDef := GenReqDefForListReplicationJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReplicationJobsResponse), nil
	}
}

// ListReplicationJobsInvoker 查询备份迁移任务列表
func (c *DrsClient) ListReplicationJobsInvoker(request *model.ListReplicationJobsRequest) *ListReplicationJobsInvoker {
	requestDef := GenReqDefForListReplicationJobs()
	return &ListReplicationJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询项目标签
//
// 查询租户在指定Project中实例类型的所有资源标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询项目标签
func (c *DrsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserJdbcDrivers 查询驱动文件列表
//
// 查询驱动文件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListUserJdbcDrivers(request *model.ListUserJdbcDriversRequest) (*model.ListUserJdbcDriversResponse, error) {
	requestDef := GenReqDefForListUserJdbcDrivers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserJdbcDriversResponse), nil
	}
}

// ListUserJdbcDriversInvoker 查询驱动文件列表
func (c *DrsClient) ListUserJdbcDriversInvoker(request *model.ListUserJdbcDriversRequest) *ListUserJdbcDriversInvoker {
	requestDef := GenReqDefForListUserJdbcDrivers()
	return &ListUserJdbcDriversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListsAgencyPermissions 查询委托的权限列表
//
// 根据源库类型，目标库类型，是否自建，获取委托所需要的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ListsAgencyPermissions(request *model.ListsAgencyPermissionsRequest) (*model.ListsAgencyPermissionsResponse, error) {
	requestDef := GenReqDefForListsAgencyPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListsAgencyPermissionsResponse), nil
	}
}

// ListsAgencyPermissionsInvoker 查询委托的权限列表
func (c *DrsClient) ListsAgencyPermissionsInvoker(request *model.ListsAgencyPermissionsRequest) *ListsAgencyPermissionsInvoker {
	requestDef := GenReqDefForListsAgencyPermissions()
	return &ListsAgencyPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyConnection 修改连接
//
// 修改创建的连接信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ModifyConnection(request *model.ModifyConnectionRequest) (*model.ModifyConnectionResponse, error) {
	requestDef := GenReqDefForModifyConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyConnectionResponse), nil
	}
}

// ModifyConnectionInvoker 修改连接
func (c *DrsClient) ModifyConnectionInvoker(request *model.ModifyConnectionRequest) *ModifyConnectionInvoker {
	requestDef := GenReqDefForModifyConnection()
	return &ModifyConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowActions 获取指定任务操作信息
//
// 获取指定任务允许、不允许、当前操作信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowActions(request *model.ShowActionsRequest) (*model.ShowActionsResponse, error) {
	requestDef := GenReqDefForShowActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowActionsResponse), nil
	}
}

// ShowActionsInvoker 获取指定任务操作信息
func (c *DrsClient) ShowActionsInvoker(request *model.ShowActionsRequest) *ShowActionsInvoker {
	requestDef := GenReqDefForShowActions()
	return &ShowActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowColumnInfoResult 获取指定数据库表列信息
//
// 获取指定数据库表列信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowColumnInfoResult(request *model.ShowColumnInfoResultRequest) (*model.ShowColumnInfoResultResponse, error) {
	requestDef := GenReqDefForShowColumnInfoResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowColumnInfoResultResponse), nil
	}
}

// ShowColumnInfoResultInvoker 获取指定数据库表列信息
func (c *DrsClient) ShowColumnInfoResultInvoker(request *model.ShowColumnInfoResultRequest) *ShowColumnInfoResultInvoker {
	requestDef := GenReqDefForShowColumnInfoResult()
	return &ShowColumnInfoResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComparePolicy 查询对比策略
//
// 查询对比策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowComparePolicy(request *model.ShowComparePolicyRequest) (*model.ShowComparePolicyResponse, error) {
	requestDef := GenReqDefForShowComparePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComparePolicyResponse), nil
	}
}

// ShowComparePolicyInvoker 查询对比策略
func (c *DrsClient) ShowComparePolicyInvoker(request *model.ShowComparePolicyRequest) *ShowComparePolicyInvoker {
	requestDef := GenReqDefForShowComparePolicy()
	return &ShowComparePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataFilteringResult 获取数据过滤校验结果
//
// 获取数据过滤校验结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDataFilteringResult(request *model.ShowDataFilteringResultRequest) (*model.ShowDataFilteringResultResponse, error) {
	requestDef := GenReqDefForShowDataFilteringResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataFilteringResultResponse), nil
	}
}

// ShowDataFilteringResultInvoker 获取数据过滤校验结果
func (c *DrsClient) ShowDataFilteringResultInvoker(request *model.ShowDataFilteringResultRequest) *ShowDataFilteringResultInvoker {
	requestDef := GenReqDefForShowDataFilteringResult()
	return &ShowDataFilteringResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataProcessingRulesResult 获取指定任务数据加工规则更新结果
//
// 获取指定任务数据加工规则更新结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDataProcessingRulesResult(request *model.ShowDataProcessingRulesResultRequest) (*model.ShowDataProcessingRulesResultResponse, error) {
	requestDef := GenReqDefForShowDataProcessingRulesResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataProcessingRulesResultResponse), nil
	}
}

// ShowDataProcessingRulesResultInvoker 获取指定任务数据加工规则更新结果
func (c *DrsClient) ShowDataProcessingRulesResultInvoker(request *model.ShowDataProcessingRulesResultRequest) *ShowDataProcessingRulesResultInvoker {
	requestDef := GenReqDefForShowDataProcessingRulesResult()
	return &ShowDataProcessingRulesResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataProgress 查询数据加工规则
//
// 查询数据加工规则:包含数据库表的映射信息、列信息、数据过滤信息、附加列信息、DDL以及DML信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDataProgress(request *model.ShowDataProgressRequest) (*model.ShowDataProgressResponse, error) {
	requestDef := GenReqDefForShowDataProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataProgressResponse), nil
	}
}

// ShowDataProgressInvoker 查询数据加工规则
func (c *DrsClient) ShowDataProgressInvoker(request *model.ShowDataProgressRequest) *ShowDataProgressInvoker {
	requestDef := GenReqDefForShowDataProgress()
	return &ShowDataProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDbObjectCollectionStatus 获取提交查询数据库对象信息的结果
//
// 获取提交查询数据库对象信息的结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDbObjectCollectionStatus(request *model.ShowDbObjectCollectionStatusRequest) (*model.ShowDbObjectCollectionStatusResponse, error) {
	requestDef := GenReqDefForShowDbObjectCollectionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbObjectCollectionStatusResponse), nil
	}
}

// ShowDbObjectCollectionStatusInvoker 获取提交查询数据库对象信息的结果
func (c *DrsClient) ShowDbObjectCollectionStatusInvoker(request *model.ShowDbObjectCollectionStatusRequest) *ShowDbObjectCollectionStatusInvoker {
	requestDef := GenReqDefForShowDbObjectCollectionStatus()
	return &ShowDbObjectCollectionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDbObjectTemplateProgress 对象选择（文件导入 - 进度查询）
//
// 对象选择（文件导入 - 进度查询）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDbObjectTemplateProgress(request *model.ShowDbObjectTemplateProgressRequest) (*model.ShowDbObjectTemplateProgressResponse, error) {
	requestDef := GenReqDefForShowDbObjectTemplateProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbObjectTemplateProgressResponse), nil
	}
}

// ShowDbObjectTemplateProgressInvoker 对象选择（文件导入 - 进度查询）
func (c *DrsClient) ShowDbObjectTemplateProgressInvoker(request *model.ShowDbObjectTemplateProgressRequest) *ShowDbObjectTemplateProgressInvoker {
	requestDef := GenReqDefForShowDbObjectTemplateProgress()
	return &ShowDbObjectTemplateProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDbObjectTemplateResult 对象选择（文件导入 - 获取导入结果）
//
// 对象选择（文件导入 - 获取导入结果）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDbObjectTemplateResult(request *model.ShowDbObjectTemplateResultRequest) (*model.ShowDbObjectTemplateResultResponse, error) {
	requestDef := GenReqDefForShowDbObjectTemplateResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbObjectTemplateResultResponse), nil
	}
}

// ShowDbObjectTemplateResultInvoker 对象选择（文件导入 - 获取导入结果）
func (c *DrsClient) ShowDbObjectTemplateResultInvoker(request *model.ShowDbObjectTemplateResultRequest) *ShowDbObjectTemplateResultInvoker {
	requestDef := GenReqDefForShowDbObjectTemplateResult()
	return &ShowDbObjectTemplateResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDbObjectsList 查询数据库对象信息
//
// 查询数据库对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDbObjectsList(request *model.ShowDbObjectsListRequest) (*model.ShowDbObjectsListResponse, error) {
	requestDef := GenReqDefForShowDbObjectsList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDbObjectsListResponse), nil
	}
}

// ShowDbObjectsListInvoker 查询数据库对象信息
func (c *DrsClient) ShowDbObjectsListInvoker(request *model.ShowDbObjectsListRequest) *ShowDbObjectsListInvoker {
	requestDef := GenReqDefForShowDbObjectsList()
	return &ShowDbObjectsListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDirtyData 查询异常数据列表
//
// 查询异常数据列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowDirtyData(request *model.ShowDirtyDataRequest) (*model.ShowDirtyDataResponse, error) {
	requestDef := GenReqDefForShowDirtyData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirtyDataResponse), nil
	}
}

// ShowDirtyDataInvoker 查询异常数据列表
func (c *DrsClient) ShowDirtyDataInvoker(request *model.ShowDirtyDataRequest) *ShowDirtyDataInvoker {
	requestDef := GenReqDefForShowDirtyData()
	return &ShowDirtyDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnterpriseProject 查询企业项目列表
//
// 查询企业项目列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowEnterpriseProject(request *model.ShowEnterpriseProjectRequest) (*model.ShowEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForShowEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnterpriseProjectResponse), nil
	}
}

// ShowEnterpriseProjectInvoker 查询企业项目列表
func (c *DrsClient) ShowEnterpriseProjectInvoker(request *model.ShowEnterpriseProjectRequest) *ShowEnterpriseProjectInvoker {
	requestDef := GenReqDefForShowEnterpriseProject()
	return &ShowEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthCompareJobDetail 查询健康对比任务详情
//
// 查询健康对比任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowHealthCompareJobDetail(request *model.ShowHealthCompareJobDetailRequest) (*model.ShowHealthCompareJobDetailResponse, error) {
	requestDef := GenReqDefForShowHealthCompareJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthCompareJobDetailResponse), nil
	}
}

// ShowHealthCompareJobDetailInvoker 查询健康对比任务详情
func (c *DrsClient) ShowHealthCompareJobDetailInvoker(request *model.ShowHealthCompareJobDetailRequest) *ShowHealthCompareJobDetailInvoker {
	requestDef := GenReqDefForShowHealthCompareJobDetail()
	return &ShowHealthCompareJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthCompareJobList 查询健康对比列表
//
// 查询健康对比列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowHealthCompareJobList(request *model.ShowHealthCompareJobListRequest) (*model.ShowHealthCompareJobListResponse, error) {
	requestDef := GenReqDefForShowHealthCompareJobList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthCompareJobListResponse), nil
	}
}

// ShowHealthCompareJobListInvoker 查询健康对比列表
func (c *DrsClient) ShowHealthCompareJobListInvoker(request *model.ShowHealthCompareJobListRequest) *ShowHealthCompareJobListInvoker {
	requestDef := GenReqDefForShowHealthCompareJobList()
	return &ShowHealthCompareJobListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthObjectCompareJobOverview 获取健康对比对象级对比概览
//
// 获取健康对比对象级对比概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowHealthObjectCompareJobOverview(request *model.ShowHealthObjectCompareJobOverviewRequest) (*model.ShowHealthObjectCompareJobOverviewResponse, error) {
	requestDef := GenReqDefForShowHealthObjectCompareJobOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthObjectCompareJobOverviewResponse), nil
	}
}

// ShowHealthObjectCompareJobOverviewInvoker 获取健康对比对象级对比概览
func (c *DrsClient) ShowHealthObjectCompareJobOverviewInvoker(request *model.ShowHealthObjectCompareJobOverviewRequest) *ShowHealthObjectCompareJobOverviewInvoker {
	requestDef := GenReqDefForShowHealthObjectCompareJobOverview()
	return &ShowHealthObjectCompareJobOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIncrementComponentsDetail 查询增量组件详情
//
// 查询任务同步的增量组件的详细信息，实时同步任务，任务模式为增量或者全量+增量才支持。具体介绍可以参考：[查询同步进度](https://support.huaweicloud.com/realtimesyn-drs/drs_10_0007.html)
// - 支持的引擎：oracle-to-gaussdbv5，oracle-to-gaussdbv5ha，gaussdbv5，gaussdbv5-to-mysql，gaussdbv5-to-gaussdbv5ha，gaussdbv5ha，gaussdbv5ha-to-gaussdbv5，gaussdbv5-to-dws，gaussdbv5ha-to-dws，gaussdbv5-to-oracle，gaussdbv5ha-to-oracle，oracle-to-dws，oracle-to-mysql
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowIncrementComponentsDetail(request *model.ShowIncrementComponentsDetailRequest) (*model.ShowIncrementComponentsDetailResponse, error) {
	requestDef := GenReqDefForShowIncrementComponentsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIncrementComponentsDetailResponse), nil
	}
}

// ShowIncrementComponentsDetailInvoker 查询增量组件详情
func (c *DrsClient) ShowIncrementComponentsDetailInvoker(request *model.ShowIncrementComponentsDetailRequest) *ShowIncrementComponentsDetailInvoker {
	requestDef := GenReqDefForShowIncrementComponentsDetail()
	return &ShowIncrementComponentsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowInstanceTags(request *model.ShowInstanceTagsRequest) (*model.ShowInstanceTagsResponse, error) {
	requestDef := GenReqDefForShowInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceTagsResponse), nil
	}
}

// ShowInstanceTagsInvoker 查询资源标签
func (c *DrsClient) ShowInstanceTagsInvoker(request *model.ShowInstanceTagsRequest) *ShowInstanceTagsInvoker {
	requestDef := GenReqDefForShowInstanceTags()
	return &ShowInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 查询任务详情
//
// 查询任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 查询任务详情
func (c *DrsClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetering 获取任务价格信息
//
// 获取询价接口的参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowMetering(request *model.ShowMeteringRequest) (*model.ShowMeteringResponse, error) {
	requestDef := GenReqDefForShowMetering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeteringResponse), nil
	}
}

// ShowMeteringInvoker 获取任务价格信息
func (c *DrsClient) ShowMeteringInvoker(request *model.ShowMeteringRequest) *ShowMeteringInvoker {
	requestDef := GenReqDefForShowMetering()
	return &ShowMeteringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitorData 查询监控数据
//
// 获取任务监控数据。
// - Cassandra灾备不支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowMonitorData(request *model.ShowMonitorDataRequest) (*model.ShowMonitorDataResponse, error) {
	requestDef := GenReqDefForShowMonitorData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitorDataResponse), nil
	}
}

// ShowMonitorDataInvoker 查询监控数据
func (c *DrsClient) ShowMonitorDataInvoker(request *model.ShowMonitorDataRequest) *ShowMonitorDataInvoker {
	requestDef := GenReqDefForShowMonitorData()
	return &ShowMonitorDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowObjectMapping 查询同步映射列表
//
// 查询实时同步映射关系包括对象选择时的库映射、schema映射、表映射和数据加工时的列映射。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowObjectMapping(request *model.ShowObjectMappingRequest) (*model.ShowObjectMappingResponse, error) {
	requestDef := GenReqDefForShowObjectMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowObjectMappingResponse), nil
	}
}

// ShowObjectMappingInvoker 查询同步映射列表
func (c *DrsClient) ShowObjectMappingInvoker(request *model.ShowObjectMappingRequest) *ShowObjectMappingInvoker {
	requestDef := GenReqDefForShowObjectMapping()
	return &ShowObjectMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPositionResult 获取查询数据库位点的结果
//
// 获取查询数据库位点的结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowPositionResult(request *model.ShowPositionResultRequest) (*model.ShowPositionResultResponse, error) {
	requestDef := GenReqDefForShowPositionResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPositionResultResponse), nil
	}
}

// ShowPositionResultInvoker 获取查询数据库位点的结果
func (c *DrsClient) ShowPositionResultInvoker(request *model.ShowPositionResultRequest) *ShowPositionResultInvoker {
	requestDef := GenReqDefForShowPositionResult()
	return &ShowPositionResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProgressData 查询数据级流式对比列表
//
// 查询不同迁移对象类型的迁移进度。
// 说明：
// - 目前仅MySQL-&gt;MySQL、MySQL-&gt;GaussDB(for MySQL)、MongoDB-&gt;DDS、DDS-&gt;MongoDB的实时迁移和所有实时同步链路支持查看迁移明细。
// - 在任务未结束前，不能修改源库和目标库的所有用户、密码和用户权限等。
// - 全量、增量完成不代表任务结束，如果存在触发器和事件将会进行迁移。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowProgressData(request *model.ShowProgressDataRequest) (*model.ShowProgressDataResponse, error) {
	requestDef := GenReqDefForShowProgressData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgressDataResponse), nil
	}
}

// ShowProgressDataInvoker 查询数据级流式对比列表
func (c *DrsClient) ShowProgressDataInvoker(request *model.ShowProgressDataRequest) *ShowProgressDataInvoker {
	requestDef := GenReqDefForShowProgressData()
	return &ShowProgressDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReplayResults 查询录制回放结果
//
// 获取录制回放结果数据，包括：回放基于时间维度统计信息，异常SQL及统计结果、慢SQL及统计结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowReplayResults(request *model.ShowReplayResultsRequest) (*model.ShowReplayResultsResponse, error) {
	requestDef := GenReqDefForShowReplayResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplayResultsResponse), nil
	}
}

// ShowReplayResultsInvoker 查询录制回放结果
func (c *DrsClient) ShowReplayResultsInvoker(request *model.ShowReplayResultsRequest) *ShowReplayResultsInvoker {
	requestDef := GenReqDefForShowReplayResults()
	return &ShowReplayResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReplicationJob 查询备份迁移任务详细信息
//
// 获取指定备份迁移任务详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowReplicationJob(request *model.ShowReplicationJobRequest) (*model.ShowReplicationJobResponse, error) {
	requestDef := GenReqDefForShowReplicationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplicationJobResponse), nil
	}
}

// ShowReplicationJobInvoker 查询备份迁移任务详细信息
func (c *DrsClient) ShowReplicationJobInvoker(request *model.ShowReplicationJobRequest) *ShowReplicationJobInvoker {
	requestDef := GenReqDefForShowReplicationJob()
	return &ShowReplicationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSupportObjectType 查询是否支持对象选择和列映射
//
// 查询任务支持的对象选择类型、列映射、支持搜索的对象类型等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowSupportObjectType(request *model.ShowSupportObjectTypeRequest) (*model.ShowSupportObjectTypeResponse, error) {
	requestDef := GenReqDefForShowSupportObjectType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSupportObjectTypeResponse), nil
	}
}

// ShowSupportObjectTypeInvoker 查询是否支持对象选择和列映射
func (c *DrsClient) ShowSupportObjectTypeInvoker(request *model.ShowSupportObjectTypeRequest) *ShowSupportObjectTypeInvoker {
	requestDef := GenReqDefForShowSupportObjectType()
	return &ShowSupportObjectTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpdateObjectSavingStatus 获取对象保存进度
//
// 获取对象保存进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ShowUpdateObjectSavingStatus(request *model.ShowUpdateObjectSavingStatusRequest) (*model.ShowUpdateObjectSavingStatusResponse, error) {
	requestDef := GenReqDefForShowUpdateObjectSavingStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpdateObjectSavingStatusResponse), nil
	}
}

// ShowUpdateObjectSavingStatusInvoker 获取对象保存进度
func (c *DrsClient) ShowUpdateObjectSavingStatusInvoker(request *model.ShowUpdateObjectSavingStatusRequest) *ShowUpdateObjectSavingStatusInvoker {
	requestDef := GenReqDefForShowUpdateObjectSavingStatus()
	return &ShowUpdateObjectSavingStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJobAction 结束任务
//
// 结束租户指定ID任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) StopJobAction(request *model.StopJobActionRequest) (*model.StopJobActionResponse, error) {
	requestDef := GenReqDefForStopJobAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobActionResponse), nil
	}
}

// StopJobActionInvoker 结束任务
func (c *DrsClient) StopJobActionInvoker(request *model.StopJobActionRequest) *StopJobActionInvoker {
	requestDef := GenReqDefForStopJobAction()
	return &StopJobActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncJdbcDriver 同步驱动文件（不再推广）
//
// 同步驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) SyncJdbcDriver(request *model.SyncJdbcDriverRequest) (*model.SyncJdbcDriverResponse, error) {
	requestDef := GenReqDefForSyncJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncJdbcDriverResponse), nil
	}
}

// SyncJdbcDriverInvoker 同步驱动文件（不再推广）
func (c *DrsClient) SyncJdbcDriverInvoker(request *model.SyncJdbcDriverRequest) *SyncJdbcDriverInvoker {
	requestDef := GenReqDefForSyncJdbcDriver()
	return &SyncJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncUserJdbcDriver 同步驱动文件
//
// 同步驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) SyncUserJdbcDriver(request *model.SyncUserJdbcDriverRequest) (*model.SyncUserJdbcDriverResponse, error) {
	requestDef := GenReqDefForSyncUserJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncUserJdbcDriverResponse), nil
	}
}

// SyncUserJdbcDriverInvoker 同步驱动文件
func (c *DrsClient) SyncUserJdbcDriverInvoker(request *model.SyncUserJdbcDriverRequest) *SyncUserJdbcDriverInvoker {
	requestDef := GenReqDefForSyncUserJdbcDriver()
	return &SyncUserJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBatchAsyncJobs 更新指定ID批量异步任务详情
//
// 更新租户指定ID批量异步任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateBatchAsyncJobs(request *model.UpdateBatchAsyncJobsRequest) (*model.UpdateBatchAsyncJobsResponse, error) {
	requestDef := GenReqDefForUpdateBatchAsyncJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBatchAsyncJobsResponse), nil
	}
}

// UpdateBatchAsyncJobsInvoker 更新指定ID批量异步任务详情
func (c *DrsClient) UpdateBatchAsyncJobsInvoker(request *model.UpdateBatchAsyncJobsRequest) *UpdateBatchAsyncJobsInvoker {
	requestDef := GenReqDefForUpdateBatchAsyncJobs()
	return &UpdateBatchAsyncJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComparePolicy 修改对比策略
//
// 修改周期性对比的对比策略，目前仅MySQL-&gt;MySQL、MySQL-&gt;GaussDB(for MySQL)、MySQL-&gt;GaussDB(DWS)、GaussDB(for MySQL)-&gt;MySQL同步任务，MySQL-&gt;MySQL、MySQL-&gt;GaussDB(for MySQL)迁移任务，MySQL-&gt;MySQL、MySQL-&gt;GaussDB(for MySQL)、GaussDB(for MySQL)-&gt;GaussDB(for MySQL)、DDM-&gt;DDM、DDS-DDS灾备任务支持对比策略设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateComparePolicy(request *model.UpdateComparePolicyRequest) (*model.UpdateComparePolicyResponse, error) {
	requestDef := GenReqDefForUpdateComparePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComparePolicyResponse), nil
	}
}

// UpdateComparePolicyInvoker 修改对比策略
func (c *DrsClient) UpdateComparePolicyInvoker(request *model.UpdateComparePolicyRequest) *UpdateComparePolicyInvoker {
	requestDef := GenReqDefForUpdateComparePolicy()
	return &UpdateComparePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataProgress 更新指定任务数据加工规则
//
// 更新指定任务数据加工规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateDataProgress(request *model.UpdateDataProgressRequest) (*model.UpdateDataProgressResponse, error) {
	requestDef := GenReqDefForUpdateDataProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataProgressResponse), nil
	}
}

// UpdateDataProgressInvoker 更新指定任务数据加工规则
func (c *DrsClient) UpdateDataProgressInvoker(request *model.UpdateDataProgressRequest) *UpdateDataProgressInvoker {
	requestDef := GenReqDefForUpdateDataProgress()
	return &UpdateDataProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJob 更新指定ID任务详情
//
// 更新租户指定ID任务详情。
// 当type取值为db_object， 进行异步处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateJob(request *model.UpdateJobRequest) (*model.UpdateJobResponse, error) {
	requestDef := GenReqDefForUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResponse), nil
	}
}

// UpdateJobInvoker 更新指定ID任务详情
func (c *DrsClient) UpdateJobInvoker(request *model.UpdateJobRequest) *UpdateJobInvoker {
	requestDef := GenReqDefForUpdateJob()
	return &UpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJobConfigurations 更新任务的参数信息
//
// 更新任务的参数信息。
// - 仅engine_type为mysql、mysql-to-pgl、mysql-to-gaussdbv5、mysql-to-gaussdbv5ha、mysql-to-dws、mysql-to-taurus、mysql-to-kafka、mysql-to-elasticsearch、mysql-to-oracle且任务状态只能为配置中、全量中、增量中、全量失败、增量失败、暂停中的实时同步任务支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateJobConfigurations(request *model.UpdateJobConfigurationsRequest) (*model.UpdateJobConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateJobConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobConfigurationsResponse), nil
	}
}

// UpdateJobConfigurationsInvoker 更新任务的参数信息
func (c *DrsClient) UpdateJobConfigurationsInvoker(request *model.UpdateJobConfigurationsRequest) *UpdateJobConfigurationsInvoker {
	requestDef := GenReqDefForUpdateJobConfigurations()
	return &UpdateJobConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReplicationJob 修改备份迁移任务信息
//
// 修改指定备份迁移任务信息，任务名与任务描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateReplicationJob(request *model.UpdateReplicationJobRequest) (*model.UpdateReplicationJobResponse, error) {
	requestDef := GenReqDefForUpdateReplicationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReplicationJobResponse), nil
	}
}

// UpdateReplicationJobInvoker 修改备份迁移任务信息
func (c *DrsClient) UpdateReplicationJobInvoker(request *model.UpdateReplicationJobRequest) *UpdateReplicationJobInvoker {
	requestDef := GenReqDefForUpdateReplicationJob()
	return &UpdateReplicationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStartPosition 更新增量任务启动位点
//
// 更新增量任务的启动位点。
// - 仅engine_type为mysql,mysql-to-dws,mysql-to-taurus,taurus,mysql-to-oracle,taurus-to-oracle,taurus-to-mysql,mysql-to-kafka,taurus-to-kafka,mongodb-to-kafka,mongodb且为单增量实时同步任务支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UpdateStartPosition(request *model.UpdateStartPositionRequest) (*model.UpdateStartPositionResponse, error) {
	requestDef := GenReqDefForUpdateStartPosition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStartPositionResponse), nil
	}
}

// UpdateStartPositionInvoker 更新增量任务启动位点
func (c *DrsClient) UpdateStartPositionInvoker(request *model.UpdateStartPositionRequest) *UpdateStartPositionInvoker {
	requestDef := GenReqDefForUpdateStartPosition()
	return &UpdateStartPositionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadDbObjectTemplate 对象选择（文件导入 - 模板上传）
//
// 对象选择（文件导入 - 模板上传）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UploadDbObjectTemplate(request *model.UploadDbObjectTemplateRequest) (*model.UploadDbObjectTemplateResponse, error) {
	requestDef := GenReqDefForUploadDbObjectTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadDbObjectTemplateResponse), nil
	}
}

// UploadDbObjectTemplateInvoker 对象选择（文件导入 - 模板上传）
func (c *DrsClient) UploadDbObjectTemplateInvoker(request *model.UploadDbObjectTemplateRequest) *UploadDbObjectTemplateInvoker {
	requestDef := GenReqDefForUploadDbObjectTemplate()
	return &UploadDbObjectTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadJdbcDriver 上传驱动文件（不再推广）
//
// 上传驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UploadJdbcDriver(request *model.UploadJdbcDriverRequest) (*model.UploadJdbcDriverResponse, error) {
	requestDef := GenReqDefForUploadJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJdbcDriverResponse), nil
	}
}

// UploadJdbcDriverInvoker 上传驱动文件（不再推广）
func (c *DrsClient) UploadJdbcDriverInvoker(request *model.UploadJdbcDriverRequest) *UploadJdbcDriverInvoker {
	requestDef := GenReqDefForUploadJdbcDriver()
	return &UploadJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadUserJdbcDriver 上传驱动文件
//
// 上传驱动文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) UploadUserJdbcDriver(request *model.UploadUserJdbcDriverRequest) (*model.UploadUserJdbcDriverResponse, error) {
	requestDef := GenReqDefForUploadUserJdbcDriver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadUserJdbcDriverResponse), nil
	}
}

// UploadUserJdbcDriverInvoker 上传驱动文件
func (c *DrsClient) UploadUserJdbcDriverInvoker(request *model.UploadUserJdbcDriverRequest) *UploadUserJdbcDriverInvoker {
	requestDef := GenReqDefForUploadUserJdbcDriver()
	return &UploadUserJdbcDriverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateJobName 任务名称校验
//
// 创建任务时对任务名称进行校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DrsClient) ValidateJobName(request *model.ValidateJobNameRequest) (*model.ValidateJobNameResponse, error) {
	requestDef := GenReqDefForValidateJobName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateJobNameResponse), nil
	}
}

// ValidateJobNameInvoker 任务名称校验
func (c *DrsClient) ValidateJobNameInvoker(request *model.ValidateJobNameRequest) *ValidateJobNameInvoker {
	requestDef := GenReqDefForValidateJobName()
	return &ValidateJobNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
