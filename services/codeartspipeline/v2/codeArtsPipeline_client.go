package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartspipeline/v2/model"
)

type CodeArtsPipelineClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsPipelineClient(hcClient *httpclient.HcHttpClient) *CodeArtsPipelineClient {
	return &CodeArtsPipelineClient{HcClient: hcClient}
}

func CodeArtsPipelineClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
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

// BatchMovePipelineToGroup 批量把流水线移动到分组下
//
// 批量把流水线移动到分组下
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) BatchMovePipelineToGroup(request *model.BatchMovePipelineToGroupRequest) (*model.BatchMovePipelineToGroupResponse, error) {
	requestDef := GenReqDefForBatchMovePipelineToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchMovePipelineToGroupResponse), nil
	}
}

// BatchMovePipelineToGroupInvoker 批量把流水线移动到分组下
func (c *CodeArtsPipelineClient) BatchMovePipelineToGroupInvoker(request *model.BatchMovePipelineToGroupRequest) *BatchMovePipelineToGroupInvoker {
	requestDef := GenReqDefForBatchMovePipelineToGroup()
	return &BatchMovePipelineToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateBasicPlugin 创建基础插件
//
// 创建基础插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreateBasicPlugin(request *model.CreateBasicPluginRequest) (*model.CreateBasicPluginResponse, error) {
	requestDef := GenReqDefForCreateBasicPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBasicPluginResponse), nil
	}
}

// CreateBasicPluginInvoker 创建基础插件
func (c *CodeArtsPipelineClient) CreateBasicPluginInvoker(request *model.CreateBasicPluginRequest) *CreateBasicPluginInvoker {
	requestDef := GenReqDefForCreateBasicPlugin()
	return &CreateBasicPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreatePipelineGroup 新建流水线分组
//
// 新建流水线分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePipelineGroup(request *model.CreatePipelineGroupRequest) (*model.CreatePipelineGroupResponse, error) {
	requestDef := GenReqDefForCreatePipelineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineGroupResponse), nil
	}
}

// CreatePipelineGroupInvoker 新建流水线分组
func (c *CodeArtsPipelineClient) CreatePipelineGroupInvoker(request *model.CreatePipelineGroupRequest) *CreatePipelineGroupInvoker {
	requestDef := GenReqDefForCreatePipelineGroup()
	return &CreatePipelineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreatePipelineTemplate 创建流水线模板
//
// 创建流水线模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePipelineTemplate(request *model.CreatePipelineTemplateRequest) (*model.CreatePipelineTemplateResponse, error) {
	requestDef := GenReqDefForCreatePipelineTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineTemplateResponse), nil
	}
}

// CreatePipelineTemplateInvoker 创建流水线模板
func (c *CodeArtsPipelineClient) CreatePipelineTemplateInvoker(request *model.CreatePipelineTemplateRequest) *CreatePipelineTemplateInvoker {
	requestDef := GenReqDefForCreatePipelineTemplate()
	return &CreatePipelineTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePluginDraft 创建插件草稿版本
//
// 创建插件草稿版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePluginDraft(request *model.CreatePluginDraftRequest) (*model.CreatePluginDraftResponse, error) {
	requestDef := GenReqDefForCreatePluginDraft()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePluginDraftResponse), nil
	}
}

// CreatePluginDraftInvoker 创建插件草稿版本
func (c *CodeArtsPipelineClient) CreatePluginDraftInvoker(request *model.CreatePluginDraftRequest) *CreatePluginDraftInvoker {
	requestDef := GenReqDefForCreatePluginDraft()
	return &CreatePluginDraftInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePluginVersion 创建插件版本
//
// 创建插件版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePluginVersion(request *model.CreatePluginVersionRequest) (*model.CreatePluginVersionResponse, error) {
	requestDef := GenReqDefForCreatePluginVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePluginVersionResponse), nil
	}
}

// CreatePluginVersionInvoker 创建插件版本
func (c *CodeArtsPipelineClient) CreatePluginVersionInvoker(request *model.CreatePluginVersionRequest) *CreatePluginVersionInvoker {
	requestDef := GenReqDefForCreatePluginVersion()
	return &CreatePluginVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublisher 创建发布商
//
// 创建发布商
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreatePublisher(request *model.CreatePublisherRequest) (*model.CreatePublisherResponse, error) {
	requestDef := GenReqDefForCreatePublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublisherResponse), nil
	}
}

// CreatePublisherInvoker 创建发布商
func (c *CodeArtsPipelineClient) CreatePublisherInvoker(request *model.CreatePublisherRequest) *CreatePublisherInvoker {
	requestDef := GenReqDefForCreatePublisher()
	return &CreatePublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRule 创建规则
//
// 创建规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

// CreateRuleInvoker 创建规则
func (c *CodeArtsPipelineClient) CreateRuleInvoker(request *model.CreateRuleRequest) *CreateRuleInvoker {
	requestDef := GenReqDefForCreateRule()
	return &CreateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStrategy 创建策略
//
// 创建策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) CreateStrategy(request *model.CreateStrategyRequest) (*model.CreateStrategyResponse, error) {
	requestDef := GenReqDefForCreateStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStrategyResponse), nil
	}
}

// CreateStrategyInvoker 创建策略
func (c *CodeArtsPipelineClient) CreateStrategyInvoker(request *model.CreateStrategyRequest) *CreateStrategyInvoker {
	requestDef := GenReqDefForCreateStrategy()
	return &CreateStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBasicPlugin 删除基础插件
//
// 删除基础插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeleteBasicPlugin(request *model.DeleteBasicPluginRequest) (*model.DeleteBasicPluginResponse, error) {
	requestDef := GenReqDefForDeleteBasicPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBasicPluginResponse), nil
	}
}

// DeleteBasicPluginInvoker 删除基础插件
func (c *CodeArtsPipelineClient) DeleteBasicPluginInvoker(request *model.DeleteBasicPluginRequest) *DeleteBasicPluginInvoker {
	requestDef := GenReqDefForDeleteBasicPlugin()
	return &DeleteBasicPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeletePipelineGroup 删除流水线分组
//
// 删除流水线分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeletePipelineGroup(request *model.DeletePipelineGroupRequest) (*model.DeletePipelineGroupResponse, error) {
	requestDef := GenReqDefForDeletePipelineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipelineGroupResponse), nil
	}
}

// DeletePipelineGroupInvoker 删除流水线分组
func (c *CodeArtsPipelineClient) DeletePipelineGroupInvoker(request *model.DeletePipelineGroupRequest) *DeletePipelineGroupInvoker {
	requestDef := GenReqDefForDeletePipelineGroup()
	return &DeletePipelineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipelineTemplate 删除流水线模板
//
// 删除流水线模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeletePipelineTemplate(request *model.DeletePipelineTemplateRequest) (*model.DeletePipelineTemplateResponse, error) {
	requestDef := GenReqDefForDeletePipelineTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipelineTemplateResponse), nil
	}
}

// DeletePipelineTemplateInvoker 删除流水线模板
func (c *CodeArtsPipelineClient) DeletePipelineTemplateInvoker(request *model.DeletePipelineTemplateRequest) *DeletePipelineTemplateInvoker {
	requestDef := GenReqDefForDeletePipelineTemplate()
	return &DeletePipelineTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePluginDraft 删除插件草稿
//
// 删除插件草稿
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeletePluginDraft(request *model.DeletePluginDraftRequest) (*model.DeletePluginDraftResponse, error) {
	requestDef := GenReqDefForDeletePluginDraft()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePluginDraftResponse), nil
	}
}

// DeletePluginDraftInvoker 删除插件草稿
func (c *CodeArtsPipelineClient) DeletePluginDraftInvoker(request *model.DeletePluginDraftRequest) *DeletePluginDraftInvoker {
	requestDef := GenReqDefForDeletePluginDraft()
	return &DeletePluginDraftInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublisher 删除发布商
//
// 删除发布商
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeletePublisher(request *model.DeletePublisherRequest) (*model.DeletePublisherResponse, error) {
	requestDef := GenReqDefForDeletePublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublisherResponse), nil
	}
}

// DeletePublisherInvoker 删除发布商
func (c *CodeArtsPipelineClient) DeletePublisherInvoker(request *model.DeletePublisherRequest) *DeletePublisherInvoker {
	requestDef := GenReqDefForDeletePublisher()
	return &DeletePublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRule 删除规则
//
// 删除规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// DeleteRuleInvoker 删除规则
func (c *CodeArtsPipelineClient) DeleteRuleInvoker(request *model.DeleteRuleRequest) *DeleteRuleInvoker {
	requestDef := GenReqDefForDeleteRule()
	return &DeleteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStrategy 删除策略
//
// 删除策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) DeleteStrategy(request *model.DeleteStrategyRequest) (*model.DeleteStrategyResponse, error) {
	requestDef := GenReqDefForDeleteStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStrategyResponse), nil
	}
}

// DeleteStrategyInvoker 删除策略
func (c *CodeArtsPipelineClient) DeleteStrategyInvoker(request *model.DeleteStrategyRequest) *DeleteStrategyInvoker {
	requestDef := GenReqDefForDeleteStrategy()
	return &DeleteStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailablePublisher 查询可用发布商
//
// 查询可用发布商
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListAvailablePublisher(request *model.ListAvailablePublisherRequest) (*model.ListAvailablePublisherResponse, error) {
	requestDef := GenReqDefForListAvailablePublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailablePublisherResponse), nil
	}
}

// ListAvailablePublisherInvoker 查询可用发布商
func (c *CodeArtsPipelineClient) ListAvailablePublisherInvoker(request *model.ListAvailablePublisherRequest) *ListAvailablePublisherInvoker {
	requestDef := GenReqDefForListAvailablePublisher()
	return &ListAvailablePublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBasePlugins 查询基础插件列表
//
// 查询基础插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListBasePlugins(request *model.ListBasePluginsRequest) (*model.ListBasePluginsResponse, error) {
	requestDef := GenReqDefForListBasePlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBasePluginsResponse), nil
	}
}

// ListBasePluginsInvoker 查询基础插件列表
func (c *CodeArtsPipelineClient) ListBasePluginsInvoker(request *model.ListBasePluginsRequest) *ListBasePluginsInvoker {
	requestDef := GenReqDefForListBasePlugins()
	return &ListBasePluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBasePluginsNewPost 分页查询可选插件列表
//
// 分页查询可选插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListBasePluginsNewPost(request *model.ListBasePluginsNewPostRequest) (*model.ListBasePluginsNewPostResponse, error) {
	requestDef := GenReqDefForListBasePluginsNewPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBasePluginsNewPostResponse), nil
	}
}

// ListBasePluginsNewPostInvoker 分页查询可选插件列表
func (c *CodeArtsPipelineClient) ListBasePluginsNewPostInvoker(request *model.ListBasePluginsNewPostRequest) *ListBasePluginsNewPostInvoker {
	requestDef := GenReqDefForListBasePluginsNewPost()
	return &ListBasePluginsNewPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPLuginVersion 查询插件所有版本信息
//
// 查询插件所有版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPLuginVersion(request *model.ListPLuginVersionRequest) (*model.ListPLuginVersionResponse, error) {
	requestDef := GenReqDefForListPLuginVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPLuginVersionResponse), nil
	}
}

// ListPLuginVersionInvoker 查询插件所有版本信息
func (c *CodeArtsPipelineClient) ListPLuginVersionInvoker(request *model.ListPLuginVersionRequest) *ListPLuginVersionInvoker {
	requestDef := GenReqDefForListPLuginVersion()
	return &ListPLuginVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPluginVersionNumber 查询插件版本号
//
// 查询插件版本号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPluginVersionNumber(request *model.ListPluginVersionNumberRequest) (*model.ListPluginVersionNumberResponse, error) {
	requestDef := GenReqDefForListPluginVersionNumber()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginVersionNumberResponse), nil
	}
}

// ListPluginVersionNumberInvoker 查询插件版本号
func (c *CodeArtsPipelineClient) ListPluginVersionNumberInvoker(request *model.ListPluginVersionNumberRequest) *ListPluginVersionNumberInvoker {
	requestDef := GenReqDefForListPluginVersionNumber()
	return &ListPluginVersionNumberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlugins 查询插件列表
//
// 查询插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPlugins(request *model.ListPluginsRequest) (*model.ListPluginsResponse, error) {
	requestDef := GenReqDefForListPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginsResponse), nil
	}
}

// ListPluginsInvoker 查询插件列表
func (c *CodeArtsPipelineClient) ListPluginsInvoker(request *model.ListPluginsRequest) *ListPluginsInvoker {
	requestDef := GenReqDefForListPlugins()
	return &ListPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectStrategy 获取规则模板实例列表
//
// 获取策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListProjectStrategy(request *model.ListProjectStrategyRequest) (*model.ListProjectStrategyResponse, error) {
	requestDef := GenReqDefForListProjectStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectStrategyResponse), nil
	}
}

// ListProjectStrategyInvoker 获取规则模板实例列表
func (c *CodeArtsPipelineClient) ListProjectStrategyInvoker(request *model.ListProjectStrategyRequest) *ListProjectStrategyInvoker {
	requestDef := GenReqDefForListProjectStrategy()
	return &ListProjectStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublisher 查询发布商列表
//
// 查询发布商列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListPublisher(request *model.ListPublisherRequest) (*model.ListPublisherResponse, error) {
	requestDef := GenReqDefForListPublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublisherResponse), nil
	}
}

// ListPublisherInvoker 查询发布商列表
func (c *CodeArtsPipelineClient) ListPublisherInvoker(request *model.ListPublisherRequest) *ListPublisherInvoker {
	requestDef := GenReqDefForListPublisher()
	return &ListPublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRule 分页获取规则列表
//
// 分页获取规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListRule(request *model.ListRuleRequest) (*model.ListRuleResponse, error) {
	requestDef := GenReqDefForListRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleResponse), nil
	}
}

// ListRuleInvoker 分页获取规则列表
func (c *CodeArtsPipelineClient) ListRuleInvoker(request *model.ListRuleRequest) *ListRuleInvoker {
	requestDef := GenReqDefForListRule()
	return &ListRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStagePlugins 查询可选插件列表
//
// 查询可选插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListStagePlugins(request *model.ListStagePluginsRequest) (*model.ListStagePluginsResponse, error) {
	requestDef := GenReqDefForListStagePlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStagePluginsResponse), nil
	}
}

// ListStagePluginsInvoker 查询可选插件列表
func (c *CodeArtsPipelineClient) ListStagePluginsInvoker(request *model.ListStagePluginsRequest) *ListStagePluginsInvoker {
	requestDef := GenReqDefForListStagePlugins()
	return &ListStagePluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStrategy 获取策略列表
//
// 获取策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ListStrategy(request *model.ListStrategyRequest) (*model.ListStrategyResponse, error) {
	requestDef := GenReqDefForListStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStrategyResponse), nil
	}
}

// ListStrategyInvoker 获取策略列表
func (c *CodeArtsPipelineClient) ListStrategyInvoker(request *model.ListStrategyRequest) *ListStrategyInvoker {
	requestDef := GenReqDefForListStrategy()
	return &ListStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// PublishPlugin 发布插件
//
// 发布插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) PublishPlugin(request *model.PublishPluginRequest) (*model.PublishPluginResponse, error) {
	requestDef := GenReqDefForPublishPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishPluginResponse), nil
	}
}

// PublishPluginInvoker 发布插件
func (c *CodeArtsPipelineClient) PublishPluginInvoker(request *model.PublishPluginRequest) *PublishPluginInvoker {
	requestDef := GenReqDefForPublishPlugin()
	return &PublishPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishPluginBind 插件绑定发布商
//
// 插件绑定发布商
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) PublishPluginBind(request *model.PublishPluginBindRequest) (*model.PublishPluginBindResponse, error) {
	requestDef := GenReqDefForPublishPluginBind()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishPluginBindResponse), nil
	}
}

// PublishPluginBindInvoker 插件绑定发布商
func (c *CodeArtsPipelineClient) PublishPluginBindInvoker(request *model.PublishPluginBindRequest) *PublishPluginBindInvoker {
	requestDef := GenReqDefForPublishPluginBind()
	return &PublishPluginBindInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishPluginDraft 发布插件草稿
//
// 发布插件草稿
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) PublishPluginDraft(request *model.PublishPluginDraftRequest) (*model.PublishPluginDraftResponse, error) {
	requestDef := GenReqDefForPublishPluginDraft()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishPluginDraftResponse), nil
	}
}

// PublishPluginDraftInvoker 发布插件草稿
func (c *CodeArtsPipelineClient) PublishPluginDraftInvoker(request *model.PublishPluginDraftRequest) *PublishPluginDraftInvoker {
	requestDef := GenReqDefForPublishPluginDraft()
	return &PublishPluginDraftInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowBasicPlugin 查询基础插件详情
//
// 查询基础插件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowBasicPlugin(request *model.ShowBasicPluginRequest) (*model.ShowBasicPluginResponse, error) {
	requestDef := GenReqDefForShowBasicPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBasicPluginResponse), nil
	}
}

// ShowBasicPluginInvoker 查询基础插件详情
func (c *CodeArtsPipelineClient) ShowBasicPluginInvoker(request *model.ShowBasicPluginRequest) *ShowBasicPluginInvoker {
	requestDef := GenReqDefForShowBasicPlugin()
	return &ShowBasicPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowPipelineArtifacts 查询流水线上的构建产物
//
// 查询流水线上的构建产物
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineArtifacts(request *model.ShowPipelineArtifactsRequest) (*model.ShowPipelineArtifactsResponse, error) {
	requestDef := GenReqDefForShowPipelineArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineArtifactsResponse), nil
	}
}

// ShowPipelineArtifactsInvoker 查询流水线上的构建产物
func (c *CodeArtsPipelineClient) ShowPipelineArtifactsInvoker(request *model.ShowPipelineArtifactsRequest) *ShowPipelineArtifactsInvoker {
	requestDef := GenReqDefForShowPipelineArtifacts()
	return &ShowPipelineArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineDetail 查询流水线详情
//
// 查询流水线详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineDetail(request *model.ShowPipelineDetailRequest) (*model.ShowPipelineDetailResponse, error) {
	requestDef := GenReqDefForShowPipelineDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineDetailResponse), nil
	}
}

// ShowPipelineDetailInvoker 查询流水线详情
func (c *CodeArtsPipelineClient) ShowPipelineDetailInvoker(request *model.ShowPipelineDetailRequest) *ShowPipelineDetailInvoker {
	requestDef := GenReqDefForShowPipelineDetail()
	return &ShowPipelineDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineGroupTree 查询流水线分组树
//
// 查询流水线分组树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineGroupTree(request *model.ShowPipelineGroupTreeRequest) (*model.ShowPipelineGroupTreeResponse, error) {
	requestDef := GenReqDefForShowPipelineGroupTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineGroupTreeResponse), nil
	}
}

// ShowPipelineGroupTreeInvoker 查询流水线分组树
func (c *CodeArtsPipelineClient) ShowPipelineGroupTreeInvoker(request *model.ShowPipelineGroupTreeRequest) *ShowPipelineGroupTreeInvoker {
	requestDef := GenReqDefForShowPipelineGroupTree()
	return &ShowPipelineGroupTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineLog 查询流水线日志
//
// 查询流水线日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPipelineLog(request *model.ShowPipelineLogRequest) (*model.ShowPipelineLogResponse, error) {
	requestDef := GenReqDefForShowPipelineLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineLogResponse), nil
	}
}

// ShowPipelineLogInvoker 查询流水线日志
func (c *CodeArtsPipelineClient) ShowPipelineLogInvoker(request *model.ShowPipelineLogRequest) *ShowPipelineLogInvoker {
	requestDef := GenReqDefForShowPipelineLog()
	return &ShowPipelineLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowPluginInputs 查询插件输入配置
//
// 查询插件输入配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPluginInputs(request *model.ShowPluginInputsRequest) (*model.ShowPluginInputsResponse, error) {
	requestDef := GenReqDefForShowPluginInputs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginInputsResponse), nil
	}
}

// ShowPluginInputsInvoker 查询插件输入配置
func (c *CodeArtsPipelineClient) ShowPluginInputsInvoker(request *model.ShowPluginInputsRequest) *ShowPluginInputsInvoker {
	requestDef := GenReqDefForShowPluginInputs()
	return &ShowPluginInputsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginMetrics 查询插件指标配置
//
// 查询插件指标配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPluginMetrics(request *model.ShowPluginMetricsRequest) (*model.ShowPluginMetricsResponse, error) {
	requestDef := GenReqDefForShowPluginMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginMetricsResponse), nil
	}
}

// ShowPluginMetricsInvoker 查询插件指标配置
func (c *CodeArtsPipelineClient) ShowPluginMetricsInvoker(request *model.ShowPluginMetricsRequest) *ShowPluginMetricsInvoker {
	requestDef := GenReqDefForShowPluginMetrics()
	return &ShowPluginMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginOutputs 查询插件输出配置
//
// 查询插件输出配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPluginOutputs(request *model.ShowPluginOutputsRequest) (*model.ShowPluginOutputsResponse, error) {
	requestDef := GenReqDefForShowPluginOutputs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginOutputsResponse), nil
	}
}

// ShowPluginOutputsInvoker 查询插件输出配置
func (c *CodeArtsPipelineClient) ShowPluginOutputsInvoker(request *model.ShowPluginOutputsRequest) *ShowPluginOutputsInvoker {
	requestDef := GenReqDefForShowPluginOutputs()
	return &ShowPluginOutputsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginVersion 查询插件版本详情
//
// 查询插件版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPluginVersion(request *model.ShowPluginVersionRequest) (*model.ShowPluginVersionResponse, error) {
	requestDef := GenReqDefForShowPluginVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginVersionResponse), nil
	}
}

// ShowPluginVersionInvoker 查询插件版本详情
func (c *CodeArtsPipelineClient) ShowPluginVersionInvoker(request *model.ShowPluginVersionRequest) *ShowPluginVersionInvoker {
	requestDef := GenReqDefForShowPluginVersion()
	return &ShowPluginVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectStrategy
//
// 查询项目级策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowProjectStrategy(request *model.ShowProjectStrategyRequest) (*model.ShowProjectStrategyResponse, error) {
	requestDef := GenReqDefForShowProjectStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectStrategyResponse), nil
	}
}

// ShowProjectStrategyInvoker
func (c *CodeArtsPipelineClient) ShowProjectStrategyInvoker(request *model.ShowProjectStrategyRequest) *ShowProjectStrategyInvoker {
	requestDef := GenReqDefForShowProjectStrategy()
	return &ShowProjectStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublisher 查询发布商详情
//
// 查询发布商详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowPublisher(request *model.ShowPublisherRequest) (*model.ShowPublisherResponse, error) {
	requestDef := GenReqDefForShowPublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublisherResponse), nil
	}
}

// ShowPublisherInvoker 查询发布商详情
func (c *CodeArtsPipelineClient) ShowPublisherInvoker(request *model.ShowPublisherRequest) *ShowPublisherInvoker {
	requestDef := GenReqDefForShowPublisher()
	return &ShowPublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRule 获取单条规则详情
//
// 获取单条规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowRule(request *model.ShowRuleRequest) (*model.ShowRuleResponse, error) {
	requestDef := GenReqDefForShowRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleResponse), nil
	}
}

// ShowRuleInvoker 获取单条规则详情
func (c *CodeArtsPipelineClient) ShowRuleInvoker(request *model.ShowRuleRequest) *ShowRuleInvoker {
	requestDef := GenReqDefForShowRule()
	return &ShowRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStepOutputs 获取流水线步骤执行输出
//
// 获取流水线步骤执行输出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowStepOutputs(request *model.ShowStepOutputsRequest) (*model.ShowStepOutputsResponse, error) {
	requestDef := GenReqDefForShowStepOutputs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStepOutputsResponse), nil
	}
}

// ShowStepOutputsInvoker 获取流水线步骤执行输出
func (c *CodeArtsPipelineClient) ShowStepOutputsInvoker(request *model.ShowStepOutputsRequest) *ShowStepOutputsInvoker {
	requestDef := GenReqDefForShowStepOutputs()
	return &ShowStepOutputsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStrategy 获取策略详情
//
// 获取策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) ShowStrategy(request *model.ShowStrategyRequest) (*model.ShowStrategyResponse, error) {
	requestDef := GenReqDefForShowStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStrategyResponse), nil
	}
}

// ShowStrategyInvoker 获取策略详情
func (c *CodeArtsPipelineClient) ShowStrategyInvoker(request *model.ShowStrategyRequest) *ShowStrategyInvoker {
	requestDef := GenReqDefForShowStrategy()
	return &ShowStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// SwitchStrategy 开关策略
//
// 修改策略状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) SwitchStrategy(request *model.SwitchStrategyRequest) (*model.SwitchStrategyResponse, error) {
	requestDef := GenReqDefForSwitchStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchStrategyResponse), nil
	}
}

// SwitchStrategyInvoker 开关策略
func (c *CodeArtsPipelineClient) SwitchStrategyInvoker(request *model.SwitchStrategyRequest) *SwitchStrategyInvoker {
	requestDef := GenReqDefForSwitchStrategy()
	return &SwitchStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBasicPlugin 更新基础插件
//
// 更新基础插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdateBasicPlugin(request *model.UpdateBasicPluginRequest) (*model.UpdateBasicPluginResponse, error) {
	requestDef := GenReqDefForUpdateBasicPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBasicPluginResponse), nil
	}
}

// UpdateBasicPluginInvoker 更新基础插件
func (c *CodeArtsPipelineClient) UpdateBasicPluginInvoker(request *model.UpdateBasicPluginRequest) *UpdateBasicPluginInvoker {
	requestDef := GenReqDefForUpdateBasicPlugin()
	return &UpdateBasicPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipelineGroup 更新流水线分组
//
// 更新流水线分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdatePipelineGroup(request *model.UpdatePipelineGroupRequest) (*model.UpdatePipelineGroupResponse, error) {
	requestDef := GenReqDefForUpdatePipelineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipelineGroupResponse), nil
	}
}

// UpdatePipelineGroupInvoker 更新流水线分组
func (c *CodeArtsPipelineClient) UpdatePipelineGroupInvoker(request *model.UpdatePipelineGroupRequest) *UpdatePipelineGroupInvoker {
	requestDef := GenReqDefForUpdatePipelineGroup()
	return &UpdatePipelineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipelineInfo 修改流水线信息
//
// 修改流水线信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdatePipelineInfo(request *model.UpdatePipelineInfoRequest) (*model.UpdatePipelineInfoResponse, error) {
	requestDef := GenReqDefForUpdatePipelineInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipelineInfoResponse), nil
	}
}

// UpdatePipelineInfoInvoker 修改流水线信息
func (c *CodeArtsPipelineClient) UpdatePipelineInfoInvoker(request *model.UpdatePipelineInfoRequest) *UpdatePipelineInfoInvoker {
	requestDef := GenReqDefForUpdatePipelineInfo()
	return &UpdatePipelineInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipelineTemplate 更新流水线模板
//
// 更新流水线模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdatePipelineTemplate(request *model.UpdatePipelineTemplateRequest) (*model.UpdatePipelineTemplateResponse, error) {
	requestDef := GenReqDefForUpdatePipelineTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipelineTemplateResponse), nil
	}
}

// UpdatePipelineTemplateInvoker 更新流水线模板
func (c *CodeArtsPipelineClient) UpdatePipelineTemplateInvoker(request *model.UpdatePipelineTemplateRequest) *UpdatePipelineTemplateInvoker {
	requestDef := GenReqDefForUpdatePipelineTemplate()
	return &UpdatePipelineTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePluginBaseInfo 更新插件基本信息
//
// 更新插件基本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdatePluginBaseInfo(request *model.UpdatePluginBaseInfoRequest) (*model.UpdatePluginBaseInfoResponse, error) {
	requestDef := GenReqDefForUpdatePluginBaseInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePluginBaseInfoResponse), nil
	}
}

// UpdatePluginBaseInfoInvoker 更新插件基本信息
func (c *CodeArtsPipelineClient) UpdatePluginBaseInfoInvoker(request *model.UpdatePluginBaseInfoRequest) *UpdatePluginBaseInfoInvoker {
	requestDef := GenReqDefForUpdatePluginBaseInfo()
	return &UpdatePluginBaseInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePluginDraft 更新插件草稿
//
// 更新插件草稿
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdatePluginDraft(request *model.UpdatePluginDraftRequest) (*model.UpdatePluginDraftResponse, error) {
	requestDef := GenReqDefForUpdatePluginDraft()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePluginDraftResponse), nil
	}
}

// UpdatePluginDraftInvoker 更新插件草稿
func (c *CodeArtsPipelineClient) UpdatePluginDraftInvoker(request *model.UpdatePluginDraftRequest) *UpdatePluginDraftInvoker {
	requestDef := GenReqDefForUpdatePluginDraft()
	return &UpdatePluginDraftInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRule 更新规则
//
// 更新规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
	requestDef := GenReqDefForUpdateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleResponse), nil
	}
}

// UpdateRuleInvoker 更新规则
func (c *CodeArtsPipelineClient) UpdateRuleInvoker(request *model.UpdateRuleRequest) *UpdateRuleInvoker {
	requestDef := GenReqDefForUpdateRule()
	return &UpdateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStrategy 修改策略
//
// 修改策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UpdateStrategy(request *model.UpdateStrategyRequest) (*model.UpdateStrategyResponse, error) {
	requestDef := GenReqDefForUpdateStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStrategyResponse), nil
	}
}

// UpdateStrategyInvoker 修改策略
func (c *CodeArtsPipelineClient) UpdateStrategyInvoker(request *model.UpdateStrategyRequest) *UpdateStrategyInvoker {
	requestDef := GenReqDefForUpdateStrategy()
	return &UpdateStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadBasicPlugin 上传基础插件
//
// 上传基础插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UploadBasicPlugin(request *model.UploadBasicPluginRequest) (*model.UploadBasicPluginResponse, error) {
	requestDef := GenReqDefForUploadBasicPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadBasicPluginResponse), nil
	}
}

// UploadBasicPluginInvoker 上传基础插件
func (c *CodeArtsPipelineClient) UploadBasicPluginInvoker(request *model.UploadBasicPluginRequest) *UploadBasicPluginInvoker {
	requestDef := GenReqDefForUploadBasicPlugin()
	return &UploadBasicPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadPluginIcon 更新插件图标
//
// 更新插件图标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UploadPluginIcon(request *model.UploadPluginIconRequest) (*model.UploadPluginIconResponse, error) {
	requestDef := GenReqDefForUploadPluginIcon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadPluginIconResponse), nil
	}
}

// UploadPluginIconInvoker 更新插件图标
func (c *CodeArtsPipelineClient) UploadPluginIconInvoker(request *model.UploadPluginIconRequest) *UploadPluginIconInvoker {
	requestDef := GenReqDefForUploadPluginIcon()
	return &UploadPluginIconInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadPublisherIcon 更新发布商图标
//
// 更新发布商图标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsPipelineClient) UploadPublisherIcon(request *model.UploadPublisherIconRequest) (*model.UploadPublisherIconResponse, error) {
	requestDef := GenReqDefForUploadPublisherIcon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadPublisherIconResponse), nil
	}
}

// UploadPublisherIconInvoker 更新发布商图标
func (c *CodeArtsPipelineClient) UploadPublisherIconInvoker(request *model.UploadPublisherIconRequest) *UploadPublisherIconInvoker {
	requestDef := GenReqDefForUploadPublisherIcon()
	return &UploadPublisherIconInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
