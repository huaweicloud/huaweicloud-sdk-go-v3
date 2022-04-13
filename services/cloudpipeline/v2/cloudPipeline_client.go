package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cloudpipeline/v2/model"
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

//批量获取流水线状态和阶段信息
func (c *CloudPipelineClient) BatchShowPipelinesStatus(request *model.BatchShowPipelinesStatusRequest) (*model.BatchShowPipelinesStatusResponse, error) {
	requestDef := GenReqDefForBatchShowPipelinesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPipelinesStatusResponse), nil
	}
}

//基于模板快速创建流水线及流水线内任务
func (c *CloudPipelineClient) CreatePipelineByTemplate(request *model.CreatePipelineByTemplateRequest) (*model.CreatePipelineByTemplateResponse, error) {
	requestDef := GenReqDefForCreatePipelineByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineByTemplateResponse), nil
	}
}

//获取流水线列表接口
func (c *CloudPipelineClient) ListPipelineSimpleInfo(request *model.ListPipelineSimpleInfoRequest) (*model.ListPipelineSimpleInfoResponse, error) {
	requestDef := GenReqDefForListPipelineSimpleInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineSimpleInfoResponse), nil
	}
}

//获取项目下流水线执行状况
func (c *CloudPipelineClient) ListPipleineBuildResult(request *model.ListPipleineBuildResultRequest) (*model.ListPipleineBuildResultResponse, error) {
	requestDef := GenReqDefForListPipleineBuildResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipleineBuildResultResponse), nil
	}
}

//查询模板列表，支持分页查询,支持模板名字模糊查询
func (c *CloudPipelineClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

//注册创建Slave接口
func (c *CloudPipelineClient) RegisterAgent(request *model.RegisterAgentRequest) (*model.RegisterAgentResponse, error) {
	requestDef := GenReqDefForRegisterAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterAgentResponse), nil
	}
}

//根据id删除流水线
func (c *CloudPipelineClient) RemovePipeline(request *model.RemovePipelineRequest) (*model.RemovePipelineResponse, error) {
	requestDef := GenReqDefForRemovePipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemovePipelineResponse), nil
	}
}

//Agent状态查询
func (c *CloudPipelineClient) ShowAgentStatus(request *model.ShowAgentStatusRequest) (*model.ShowAgentStatusResponse, error) {
	requestDef := GenReqDefForShowAgentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentStatusResponse), nil
	}
}

//检查流水线创建状态
func (c *CloudPipelineClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

//获取流水线状态,阶段及任务信息
func (c *CloudPipelineClient) ShowPipleineStatus(request *model.ShowPipleineStatusRequest) (*model.ShowPipleineStatusResponse, error) {
	requestDef := GenReqDefForShowPipleineStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipleineStatusResponse), nil
	}
}

//查询模板详情
func (c *CloudPipelineClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}

//启动流水线
func (c *CloudPipelineClient) StartNewPipeline(request *model.StartNewPipelineRequest) (*model.StartNewPipelineResponse, error) {
	requestDef := GenReqDefForStartNewPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNewPipelineResponse), nil
	}
}

//停止流水线
func (c *CloudPipelineClient) StopPipelineNew(request *model.StopPipelineNewRequest) (*model.StopPipelineNewResponse, error) {
	requestDef := GenReqDefForStopPipelineNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineNewResponse), nil
	}
}
