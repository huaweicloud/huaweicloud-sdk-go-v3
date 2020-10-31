package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpipeline/v2/model"
)

type DevcloudpipelineClient struct {
	hcClient *http_client.HcHttpClient
}

func NewDevcloudpipelineClient(hcClient *http_client.HcHttpClient) *DevcloudpipelineClient {
	return &DevcloudpipelineClient{hcClient: hcClient}
}

func DevcloudpipelineClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//批量获取流水线状态和阶段信息
func (c *DevcloudpipelineClient) BatchShowPipelinesStatus(request *model.BatchShowPipelinesStatusRequest) (*model.BatchShowPipelinesStatusResponse, error) {
	requestDef := GenReqDefForBatchShowPipelinesStatus()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPipelinesStatusResponse), nil
	}
}

//基于模板快速创建流水线及流水线内任务
func (c *DevcloudpipelineClient) CreatePipelineByTemplate(request *model.CreatePipelineByTemplateRequest) (*model.CreatePipelineByTemplateResponse, error) {
	requestDef := GenReqDefForCreatePipelineByTemplate()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipelineByTemplateResponse), nil
	}
}

//查询模板列表，支持分页查询,支持模板名字模糊查询
func (c *DevcloudpipelineClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

//注册创建Slave接口
func (c *DevcloudpipelineClient) RegisterAgent(request *model.RegisterAgentRequest) (*model.RegisterAgentResponse, error) {
	requestDef := GenReqDefForRegisterAgent()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterAgentResponse), nil
	}
}

//根据id删除流水线
func (c *DevcloudpipelineClient) RemovePipeline(request *model.RemovePipelineRequest) (*model.RemovePipelineResponse, error) {
	requestDef := GenReqDefForRemovePipeline()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemovePipelineResponse), nil
	}
}

//Agent状态查询
func (c *DevcloudpipelineClient) ShowAgentStatus(request *model.ShowAgentStatusRequest) (*model.ShowAgentStatusResponse, error) {
	requestDef := GenReqDefForShowAgentStatus()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentStatusResponse), nil
	}
}

//检查流水线创建状态
func (c *DevcloudpipelineClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

//获取流水线状态,阶段及任务信息
func (c *DevcloudpipelineClient) ShowPipleineStatus(request *model.ShowPipleineStatusRequest) (*model.ShowPipleineStatusResponse, error) {
	requestDef := GenReqDefForShowPipleineStatus()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipleineStatusResponse), nil
	}
}

//查询模板详情
func (c *DevcloudpipelineClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}

//启动流水线
func (c *DevcloudpipelineClient) StartNewPipeline(request *model.StartNewPipelineRequest) (*model.StartNewPipelineResponse, error) {
	requestDef := GenReqDefForStartNewPipeline()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNewPipelineResponse), nil
	}
}

//执行流水线
func (c *DevcloudpipelineClient) StartPipeline(request *model.StartPipelineRequest) (*model.StartPipelineResponse, error) {
	requestDef := GenReqDefForStartPipeline()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineResponse), nil
	}
}

//停止流水线
func (c *DevcloudpipelineClient) StopPipeline(request *model.StopPipelineRequest) (*model.StopPipelineResponse, error) {
	requestDef := GenReqDefForStopPipeline()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineResponse), nil
	}
}
