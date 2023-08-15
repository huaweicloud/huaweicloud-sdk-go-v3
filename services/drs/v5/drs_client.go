package v5

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/drs/v5/model"
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

// ShowProgressData 查询数据级流式对比列表
//
// 查询不同迁移对象类型的迁移进度。
// 说明：
// - 目前仅MySQL-&gt;MySQL、MySQL-&gt;GaussDB(for MySQL)、MongoDB-&gt;DDS、DDS-&gt;MongoDB的迁移支持查看迁移明细。
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
