package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vas/v2/model"
)

type VasClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewVasClient(hcClient *httpclient.HcHttpClient) *VasClient {
	return &VasClient{HcClient: hcClient}
}

func VasClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateTasks 创建服务作业
//
// 该接口用于创建服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

// CreateTasksInvoker 创建服务作业
func (c *VasClient) CreateTasksInvoker(request *model.CreateTasksRequest) *CreateTasksInvoker {
	requestDef := GenReqDefForCreateTasks()
	return &CreateTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除服务作业
//
// 该接口用于删除服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除服务作业
func (c *VasClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasksDetails 获取服务作业列表
//
// 该接口用于获取服务作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) ListTasksDetails(request *model.ListTasksDetailsRequest) (*model.ListTasksDetailsResponse, error) {
	requestDef := GenReqDefForListTasksDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksDetailsResponse), nil
	}
}

// ListTasksDetailsInvoker 获取服务作业列表
func (c *VasClient) ListTasksDetailsInvoker(request *model.ListTasksDetailsRequest) *ListTasksDetailsInvoker {
	requestDef := GenReqDefForListTasksDetails()
	return &ListTasksDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 查询服务作业
//
// 该接口用于查询服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 查询服务作业
func (c *VasClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartTask 启动服务作业
//
// 该接口用于启动服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) StartTask(request *model.StartTaskRequest) (*model.StartTaskResponse, error) {
	requestDef := GenReqDefForStartTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartTaskResponse), nil
	}
}

// StartTaskInvoker 启动服务作业
func (c *VasClient) StartTaskInvoker(request *model.StartTaskRequest) *StartTaskInvoker {
	requestDef := GenReqDefForStartTask()
	return &StartTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopTask 停止服务作业
//
// 该接口用于停止服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) StopTask(request *model.StopTaskRequest) (*model.StopTaskResponse, error) {
	requestDef := GenReqDefForStopTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskResponse), nil
	}
}

// StopTaskInvoker 停止服务作业
func (c *VasClient) StopTaskInvoker(request *model.StopTaskRequest) *StopTaskInvoker {
	requestDef := GenReqDefForStopTask()
	return &StopTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTask 更新服务作业
//
// 该接口用于更新服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VasClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

// UpdateTaskInvoker 更新服务作业
func (c *VasClient) UpdateTaskInvoker(request *model.UpdateTaskRequest) *UpdateTaskInvoker {
	requestDef := GenReqDefForUpdateTask()
	return &UpdateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
