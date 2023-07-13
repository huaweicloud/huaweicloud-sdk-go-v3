package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/oroas/v1/model"
)

type OroasClient struct {
	HcClient *http_client.HcHttpClient
}

func NewOroasClient(hcClient *http_client.HcHttpClient) *OroasClient {
	return &OroasClient{HcClient: hcClient}
}

func OroasClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateTask 创建任务
//
// 创建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OroasClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// CreateTaskInvoker 创建任务
func (c *OroasClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除任务
//
// 删除任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OroasClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除任务
func (c *OroasClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTask 查询任务列表
//
// 查询任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OroasClient) ListTask(request *model.ListTaskRequest) (*model.ListTaskResponse, error) {
	requestDef := GenReqDefForListTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskResponse), nil
	}
}

// ListTaskInvoker 查询任务列表
func (c *OroasClient) ListTaskInvoker(request *model.ListTaskRequest) *ListTaskInvoker {
	requestDef := GenReqDefForListTask()
	return &ListTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 获取任务详情
//
// 获取任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OroasClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 获取任务详情
func (c *OroasClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
