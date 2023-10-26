package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartspipeline/v2/model"
)

type CodeArtsPipelineClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeArtsPipelineClient(hcClient *http_client.HcHttpClient) *CodeArtsPipelineClient {
	return &CodeArtsPipelineClient{HcClient: hcClient}
}

func CodeArtsPipelineClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AcceptManualReview 通过人工审核
//
// 通过人工审核
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) AcceptManualReview(request *model.AcceptManualReviewRequest) (*model.AcceptManualReviewResponse, error) {
	requestDef := GenReqDefForAcceptManualReview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptManualReviewResponse), nil
	}
}

// AcceptManualReviewInvoker 通过人工审核
func (c *CodeArtsPipelineClient) AcceptManualReviewInvoker(request *model.AcceptManualReviewRequest) *AcceptManualReviewInvoker {
	requestDef := GenReqDefForAcceptManualReview()
	return &AcceptManualReviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowPipelinesLatestStatus 批量获取流水线状态
//
// 批量获取流水线状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) BatchShowPipelinesLatestStatus(request *model.BatchShowPipelinesLatestStatusRequest) (*model.BatchShowPipelinesLatestStatusResponse, error) {
	requestDef := GenReqDefForBatchShowPipelinesLatestStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPipelinesLatestStatusResponse), nil
	}
}

// BatchShowPipelinesLatestStatusInvoker 批量获取流水线状态
func (c *CodeArtsPipelineClient) BatchShowPipelinesLatestStatusInvoker(request *model.BatchShowPipelinesLatestStatusRequest) *BatchShowPipelinesLatestStatusInvoker {
	requestDef := GenReqDefForBatchShowPipelinesLatestStatus()
	return &BatchShowPipelinesLatestStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowPipelinesStatus 批量获取流水线状态
//
// 批量获取流水线状态和阶段信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) BatchShowPipelinesStatus(request *model.BatchShowPipelinesStatusRequest) (*model.BatchShowPipelinesStatusResponse, error) {
	requestDef := GenReqDefForBatchShowPipelinesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPipelinesStatusResponse), nil
	}
}

// BatchShowPipelinesStatusInvoker 批量获取流水线状态
func (c *CodeArtsPipelineClient) BatchShowPipelinesStatusInvoker(request *model.BatchShowPipelinesStatusRequest) *BatchShowPipelinesStatusInvoker {
	requestDef := GenReqDefForBatchShowPipelinesStatus()
	return &BatchShowPipelinesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipelineByTemplate 基于模板快速创建流水线及流水线内任务
//
// 基于模板快速创建流水线及流水线内任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePipelineByTemplate(request *model.CreatePipelineByTemplateRequest) (*model.CreatePipelineByTemplateResponse, error) {
	requestDef := GenReqDefForCreatePipelineByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineByTemplateResponse), nil
	}
}

// CreatePipelineByTemplateInvoker 基于模板快速创建流水线及流水线内任务
func (c *CodeArtsPipelineClient) CreatePipelineByTemplateInvoker(request *model.CreatePipelineByTemplateRequest) *CreatePipelineByTemplateInvoker {
	requestDef := GenReqDefForCreatePipelineByTemplate()
	return &CreatePipelineByTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipelineByTemplateId 基于模板创建流水线
//
// 基于模板创建流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePipelineByTemplateId(request *model.CreatePipelineByTemplateIdRequest) (*model.CreatePipelineByTemplateIdResponse, error) {
	requestDef := GenReqDefForCreatePipelineByTemplateId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineByTemplateIdResponse), nil
	}
}

// CreatePipelineByTemplateIdInvoker 基于模板创建流水线
func (c *CodeArtsPipelineClient) CreatePipelineByTemplateIdInvoker(request *model.CreatePipelineByTemplateIdRequest) *CreatePipelineByTemplateIdInvoker {
	requestDef := GenReqDefForCreatePipelineByTemplateId()
	return &CreatePipelineByTemplateIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipelineNew 创建流水线
//
// 创建流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePipelineNew(request *model.CreatePipelineNewRequest) (*model.CreatePipelineNewResponse, error) {
	requestDef := GenReqDefForCreatePipelineNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineNewResponse), nil
	}
}

// CreatePipelineNewInvoker 创建流水线
func (c *CodeArtsPipelineClient) CreatePipelineNewInvoker(request *model.CreatePipelineNewRequest) *CreatePipelineNewInvoker {
	requestDef := GenReqDefForCreatePipelineNew()
	return &CreatePipelineNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipeline 删除流水线
//
// 删除流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeletePipeline(request *model.DeletePipelineRequest) (*model.DeletePipelineResponse, error) {
	requestDef := GenReqDefForDeletePipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipelineResponse), nil
	}
}

// DeletePipelineInvoker 删除流水线
func (c *CodeArtsPipelineClient) DeletePipelineInvoker(request *model.DeletePipelineRequest) *DeletePipelineInvoker {
	requestDef := GenReqDefForDeletePipeline()
	return &DeletePipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineRuns 获取流水线执行记录
//
// 获取流水线执行记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPipelineRuns(request *model.ListPipelineRunsRequest) (*model.ListPipelineRunsResponse, error) {
	requestDef := GenReqDefForListPipelineRuns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineRunsResponse), nil
	}
}

// ListPipelineRunsInvoker 获取流水线执行记录
func (c *CodeArtsPipelineClient) ListPipelineRunsInvoker(request *model.ListPipelineRunsRequest) *ListPipelineRunsInvoker {
	requestDef := GenReqDefForListPipelineRuns()
	return &ListPipelineRunsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineSimpleInfo 获取流水线列表接口
//
// 获取流水线列表接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPipelineSimpleInfo(request *model.ListPipelineSimpleInfoRequest) (*model.ListPipelineSimpleInfoResponse, error) {
	requestDef := GenReqDefForListPipelineSimpleInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineSimpleInfoResponse), nil
	}
}

// ListPipelineSimpleInfoInvoker 获取流水线列表接口
func (c *CodeArtsPipelineClient) ListPipelineSimpleInfoInvoker(request *model.ListPipelineSimpleInfoRequest) *ListPipelineSimpleInfoInvoker {
	requestDef := GenReqDefForListPipelineSimpleInfo()
	return &ListPipelineSimpleInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineTemplates 查询模板列表
//
// 查询流水线模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPipelineTemplates(request *model.ListPipelineTemplatesRequest) (*model.ListPipelineTemplatesResponse, error) {
	requestDef := GenReqDefForListPipelineTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineTemplatesResponse), nil
	}
}

// ListPipelineTemplatesInvoker 查询模板列表
func (c *CodeArtsPipelineClient) ListPipelineTemplatesInvoker(request *model.ListPipelineTemplatesRequest) *ListPipelineTemplatesInvoker {
	requestDef := GenReqDefForListPipelineTemplates()
	return &ListPipelineTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelines 获取流水线列表/获取项目下流水线执行状况
//
// 获取流水线列表/获取项目下流水线执行状况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPipelines(request *model.ListPipelinesRequest) (*model.ListPipelinesResponse, error) {
	requestDef := GenReqDefForListPipelines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelinesResponse), nil
	}
}

// ListPipelinesInvoker 获取流水线列表/获取项目下流水线执行状况
func (c *CodeArtsPipelineClient) ListPipelinesInvoker(request *model.ListPipelinesRequest) *ListPipelinesInvoker {
	requestDef := GenReqDefForListPipelines()
	return &ListPipelinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipleineBuildResult 获取项目下流水线执行状况
//
// 获取项目下流水线执行状况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPipleineBuildResult(request *model.ListPipleineBuildResultRequest) (*model.ListPipleineBuildResultResponse, error) {
	requestDef := GenReqDefForListPipleineBuildResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipleineBuildResultResponse), nil
	}
}

// ListPipleineBuildResultInvoker 获取项目下流水线执行状况
func (c *CodeArtsPipelineClient) ListPipleineBuildResultInvoker(request *model.ListPipleineBuildResultRequest) *ListPipleineBuildResultInvoker {
	requestDef := GenReqDefForListPipleineBuildResult()
	return &ListPipleineBuildResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 查询模板列表
//
// 查询模板列表，支持分页查询,支持模板名字模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 查询模板列表
func (c *CodeArtsPipelineClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RejectManualReview 驳回人工审核
//
// 驳回人工审核
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) RejectManualReview(request *model.RejectManualReviewRequest) (*model.RejectManualReviewResponse, error) {
	requestDef := GenReqDefForRejectManualReview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RejectManualReviewResponse), nil
	}
}

// RejectManualReviewInvoker 驳回人工审核
func (c *CodeArtsPipelineClient) RejectManualReviewInvoker(request *model.RejectManualReviewRequest) *RejectManualReviewInvoker {
	requestDef := GenReqDefForRejectManualReview()
	return &RejectManualReviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemovePipeline 删除流水线
//
// 根据id删除流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) RemovePipeline(request *model.RemovePipelineRequest) (*model.RemovePipelineResponse, error) {
	requestDef := GenReqDefForRemovePipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemovePipelineResponse), nil
	}
}

// RemovePipelineInvoker 删除流水线
func (c *CodeArtsPipelineClient) RemovePipelineInvoker(request *model.RemovePipelineRequest) *RemovePipelineInvoker {
	requestDef := GenReqDefForRemovePipeline()
	return &RemovePipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryPipelineRun 重试运行流水线
//
// 重试运行流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) RetryPipelineRun(request *model.RetryPipelineRunRequest) (*model.RetryPipelineRunResponse, error) {
	requestDef := GenReqDefForRetryPipelineRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryPipelineRunResponse), nil
	}
}

// RetryPipelineRunInvoker 重试运行流水线
func (c *CodeArtsPipelineClient) RetryPipelineRunInvoker(request *model.RetryPipelineRunRequest) *RetryPipelineRunInvoker {
	requestDef := GenReqDefForRetryPipelineRun()
	return &RetryPipelineRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunPipeline 启动流水线
//
// 启动流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) RunPipeline(request *model.RunPipelineRequest) (*model.RunPipelineResponse, error) {
	requestDef := GenReqDefForRunPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunPipelineResponse), nil
	}
}

// RunPipelineInvoker 启动流水线
func (c *CodeArtsPipelineClient) RunPipelineInvoker(request *model.RunPipelineRequest) *RunPipelineInvoker {
	requestDef := GenReqDefForRunPipeline()
	return &RunPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceStatus 检查流水线创建状态
//
// 检查流水线创建状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

// ShowInstanceStatusInvoker 检查流水线创建状态
func (c *CodeArtsPipelineClient) ShowInstanceStatusInvoker(request *model.ShowInstanceStatusRequest) *ShowInstanceStatusInvoker {
	requestDef := GenReqDefForShowInstanceStatus()
	return &ShowInstanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineRunDetail 获取流水线状态/获取流水线执行详情
//
// 获取流水线状态/获取流水线执行详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineRunDetail(request *model.ShowPipelineRunDetailRequest) (*model.ShowPipelineRunDetailResponse, error) {
	requestDef := GenReqDefForShowPipelineRunDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineRunDetailResponse), nil
	}
}

// ShowPipelineRunDetailInvoker 获取流水线状态/获取流水线执行详情
func (c *CodeArtsPipelineClient) ShowPipelineRunDetailInvoker(request *model.ShowPipelineRunDetailRequest) *ShowPipelineRunDetailInvoker {
	requestDef := GenReqDefForShowPipelineRunDetail()
	return &ShowPipelineRunDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineTemplateDetail 查询模板详情
//
// 查询模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineTemplateDetail(request *model.ShowPipelineTemplateDetailRequest) (*model.ShowPipelineTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowPipelineTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineTemplateDetailResponse), nil
	}
}

// ShowPipelineTemplateDetailInvoker 查询模板详情
func (c *CodeArtsPipelineClient) ShowPipelineTemplateDetailInvoker(request *model.ShowPipelineTemplateDetailRequest) *ShowPipelineTemplateDetailInvoker {
	requestDef := GenReqDefForShowPipelineTemplateDetail()
	return &ShowPipelineTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipleineStatus 获取流水线状态
//
// 获取流水线状态,阶段及任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipleineStatus(request *model.ShowPipleineStatusRequest) (*model.ShowPipleineStatusResponse, error) {
	requestDef := GenReqDefForShowPipleineStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipleineStatusResponse), nil
	}
}

// ShowPipleineStatusInvoker 获取流水线状态
func (c *CodeArtsPipelineClient) ShowPipleineStatusInvoker(request *model.ShowPipleineStatusRequest) *ShowPipleineStatusInvoker {
	requestDef := GenReqDefForShowPipleineStatus()
	return &ShowPipleineStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateDetail 查询模板详情
//
// 查询模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}

// ShowTemplateDetailInvoker 查询模板详情
func (c *CodeArtsPipelineClient) ShowTemplateDetailInvoker(request *model.ShowTemplateDetailRequest) *ShowTemplateDetailInvoker {
	requestDef := GenReqDefForShowTemplateDetail()
	return &ShowTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartNewPipeline 启动流水线
//
// 启动流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) StartNewPipeline(request *model.StartNewPipelineRequest) (*model.StartNewPipelineResponse, error) {
	requestDef := GenReqDefForStartNewPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNewPipelineResponse), nil
	}
}

// StartNewPipelineInvoker 启动流水线
func (c *CodeArtsPipelineClient) StartNewPipelineInvoker(request *model.StartNewPipelineRequest) *StartNewPipelineInvoker {
	requestDef := GenReqDefForStartNewPipeline()
	return &StartNewPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopPipelineNew 停止流水线
//
// 停止流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) StopPipelineNew(request *model.StopPipelineNewRequest) (*model.StopPipelineNewResponse, error) {
	requestDef := GenReqDefForStopPipelineNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineNewResponse), nil
	}
}

// StopPipelineNewInvoker 停止流水线
func (c *CodeArtsPipelineClient) StopPipelineNewInvoker(request *model.StopPipelineNewRequest) *StopPipelineNewInvoker {
	requestDef := GenReqDefForStopPipelineNew()
	return &StopPipelineNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopPipelineRun 停止流水线
//
// 停止流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) StopPipelineRun(request *model.StopPipelineRunRequest) (*model.StopPipelineRunResponse, error) {
	requestDef := GenReqDefForStopPipelineRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineRunResponse), nil
	}
}

// StopPipelineRunInvoker 停止流水线
func (c *CodeArtsPipelineClient) StopPipelineRunInvoker(request *model.StopPipelineRunRequest) *StopPipelineRunInvoker {
	requestDef := GenReqDefForStopPipelineRun()
	return &StopPipelineRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
