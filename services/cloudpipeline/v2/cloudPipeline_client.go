package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpipeline/v2/model"
)

type CloudPipelineClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudPipelineClient(hcClient *http_client.HcHttpClient) *CloudPipelineClient {
	return &CloudPipelineClient{HcClient: hcClient}
}

func CloudPipelineClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchShowPipelinesStatus 批量获取流水线状态
//
// 批量获取流水线状态和阶段信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) BatchShowPipelinesStatus(request *model.BatchShowPipelinesStatusRequest) (*model.BatchShowPipelinesStatusResponse, error) {
	requestDef := GenReqDefForBatchShowPipelinesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPipelinesStatusResponse), nil
	}
}

// BatchShowPipelinesStatusInvoker 批量获取流水线状态
func (c *CloudPipelineClient) BatchShowPipelinesStatusInvoker(request *model.BatchShowPipelinesStatusRequest) *BatchShowPipelinesStatusInvoker {
	requestDef := GenReqDefForBatchShowPipelinesStatus()
	return &BatchShowPipelinesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipelineByTemplate 基于模板快速创建流水线及流水线内任务
//
// 基于模板快速创建流水线及流水线内任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) CreatePipelineByTemplate(request *model.CreatePipelineByTemplateRequest) (*model.CreatePipelineByTemplateResponse, error) {
	requestDef := GenReqDefForCreatePipelineByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineByTemplateResponse), nil
	}
}

// CreatePipelineByTemplateInvoker 基于模板快速创建流水线及流水线内任务
func (c *CloudPipelineClient) CreatePipelineByTemplateInvoker(request *model.CreatePipelineByTemplateRequest) *CreatePipelineByTemplateInvoker {
	requestDef := GenReqDefForCreatePipelineByTemplate()
	return &CreatePipelineByTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineSimpleInfo 获取流水线列表接口
//
// 获取流水线列表接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ListPipelineSimpleInfo(request *model.ListPipelineSimpleInfoRequest) (*model.ListPipelineSimpleInfoResponse, error) {
	requestDef := GenReqDefForListPipelineSimpleInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineSimpleInfoResponse), nil
	}
}

// ListPipelineSimpleInfoInvoker 获取流水线列表接口
func (c *CloudPipelineClient) ListPipelineSimpleInfoInvoker(request *model.ListPipelineSimpleInfoRequest) *ListPipelineSimpleInfoInvoker {
	requestDef := GenReqDefForListPipelineSimpleInfo()
	return &ListPipelineSimpleInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipleineBuildResult 获取项目下流水线执行状况
//
// 获取项目下流水线执行状况
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ListPipleineBuildResult(request *model.ListPipleineBuildResultRequest) (*model.ListPipleineBuildResultResponse, error) {
	requestDef := GenReqDefForListPipleineBuildResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipleineBuildResultResponse), nil
	}
}

// ListPipleineBuildResultInvoker 获取项目下流水线执行状况
func (c *CloudPipelineClient) ListPipleineBuildResultInvoker(request *model.ListPipleineBuildResultRequest) *ListPipleineBuildResultInvoker {
	requestDef := GenReqDefForListPipleineBuildResult()
	return &ListPipleineBuildResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 查询模板列表
//
// 查询模板列表，支持分页查询,支持模板名字模糊查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 查询模板列表
func (c *CloudPipelineClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemovePipeline 删除流水线
//
// 根据id删除流水线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) RemovePipeline(request *model.RemovePipelineRequest) (*model.RemovePipelineResponse, error) {
	requestDef := GenReqDefForRemovePipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemovePipelineResponse), nil
	}
}

// RemovePipelineInvoker 删除流水线
func (c *CloudPipelineClient) RemovePipelineInvoker(request *model.RemovePipelineRequest) *RemovePipelineInvoker {
	requestDef := GenReqDefForRemovePipeline()
	return &RemovePipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceStatus 检查流水线创建状态
//
// 检查流水线创建状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

// ShowInstanceStatusInvoker 检查流水线创建状态
func (c *CloudPipelineClient) ShowInstanceStatusInvoker(request *model.ShowInstanceStatusRequest) *ShowInstanceStatusInvoker {
	requestDef := GenReqDefForShowInstanceStatus()
	return &ShowInstanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipleineStatus 获取流水线状态
//
// 获取流水线状态,阶段及任务信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ShowPipleineStatus(request *model.ShowPipleineStatusRequest) (*model.ShowPipleineStatusResponse, error) {
	requestDef := GenReqDefForShowPipleineStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipleineStatusResponse), nil
	}
}

// ShowPipleineStatusInvoker 获取流水线状态
func (c *CloudPipelineClient) ShowPipleineStatusInvoker(request *model.ShowPipleineStatusRequest) *ShowPipleineStatusInvoker {
	requestDef := GenReqDefForShowPipleineStatus()
	return &ShowPipleineStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateDetail 查询模板详情
//
// 查询模板详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}

// ShowTemplateDetailInvoker 查询模板详情
func (c *CloudPipelineClient) ShowTemplateDetailInvoker(request *model.ShowTemplateDetailRequest) *ShowTemplateDetailInvoker {
	requestDef := GenReqDefForShowTemplateDetail()
	return &ShowTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartNewPipeline 启动流水线
//
// 启动流水线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) StartNewPipeline(request *model.StartNewPipelineRequest) (*model.StartNewPipelineResponse, error) {
	requestDef := GenReqDefForStartNewPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNewPipelineResponse), nil
	}
}

// StartNewPipelineInvoker 启动流水线
func (c *CloudPipelineClient) StartNewPipelineInvoker(request *model.StartNewPipelineRequest) *StartNewPipelineInvoker {
	requestDef := GenReqDefForStartNewPipeline()
	return &StartNewPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopPipelineNew 停止流水线
//
// 停止流水线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudPipelineClient) StopPipelineNew(request *model.StopPipelineNewRequest) (*model.StopPipelineNewResponse, error) {
	requestDef := GenReqDefForStopPipelineNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineNewResponse), nil
	}
}

// StopPipelineNewInvoker 停止流水线
func (c *CloudPipelineClient) StopPipelineNewInvoker(request *model.StopPipelineNewRequest) *StopPipelineNewInvoker {
	requestDef := GenReqDefForStopPipelineNew()
	return &StopPipelineNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
