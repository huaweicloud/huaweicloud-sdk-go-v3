package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/devsecurity/v1/model"
)

type DevSecurityClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDevSecurityClient(hcClient *http_client.HcHttpClient) *DevSecurityClient {
	return &DevSecurityClient{HcClient: hcClient}
}

func DevSecurityClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateSecAppTask 创建移动应用安全任务并启动
//
// 创建移动应用安全任务并启动
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevSecurityClient) CreateSecAppTask(request *model.CreateSecAppTaskRequest) (*model.CreateSecAppTaskResponse, error) {
	requestDef := GenReqDefForCreateSecAppTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecAppTaskResponse), nil
	}
}

// CreateSecAppTaskInvoker 创建移动应用安全任务并启动
func (c *DevSecurityClient) CreateSecAppTaskInvoker(request *model.CreateSecAppTaskRequest) *CreateSecAppTaskInvoker {
	requestDef := GenReqDefForCreateSecAppTask()
	return &CreateSecAppTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecAppTask 删除移动应用安全任务
//
// 删除移动应用安全任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevSecurityClient) DeleteSecAppTask(request *model.DeleteSecAppTaskRequest) (*model.DeleteSecAppTaskResponse, error) {
	requestDef := GenReqDefForDeleteSecAppTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecAppTaskResponse), nil
	}
}

// DeleteSecAppTaskInvoker 删除移动应用安全任务
func (c *DevSecurityClient) DeleteSecAppTaskInvoker(request *model.DeleteSecAppTaskRequest) *DeleteSecAppTaskInvoker {
	requestDef := GenReqDefForDeleteSecAppTask()
	return &DeleteSecAppTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecAppTaskResult 获取移动应用安全任务结果
//
// 获取移动应用安全任务结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevSecurityClient) ShowSecAppTaskResult(request *model.ShowSecAppTaskResultRequest) (*model.ShowSecAppTaskResultResponse, error) {
	requestDef := GenReqDefForShowSecAppTaskResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecAppTaskResultResponse), nil
	}
}

// ShowSecAppTaskResultInvoker 获取移动应用安全任务结果
func (c *DevSecurityClient) ShowSecAppTaskResultInvoker(request *model.ShowSecAppTaskResultRequest) *ShowSecAppTaskResultInvoker {
	requestDef := GenReqDefForShowSecAppTaskResult()
	return &ShowSecAppTaskResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecAppTaskStatus 查询移动应用安全任务状态
//
// 查询移动应用安全任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevSecurityClient) ShowSecAppTaskStatus(request *model.ShowSecAppTaskStatusRequest) (*model.ShowSecAppTaskStatusResponse, error) {
	requestDef := GenReqDefForShowSecAppTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecAppTaskStatusResponse), nil
	}
}

// ShowSecAppTaskStatusInvoker 查询移动应用安全任务状态
func (c *DevSecurityClient) ShowSecAppTaskStatusInvoker(request *model.ShowSecAppTaskStatusRequest) *ShowSecAppTaskStatusInvoker {
	requestDef := GenReqDefForShowSecAppTaskStatus()
	return &ShowSecAppTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
