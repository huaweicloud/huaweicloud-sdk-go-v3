package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/campusgo/v2/model"
)

type CampusGoClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCampusGoClient(hcClient *httpclient.HcHttpClient) *CampusGoClient {
	return &CampusGoClient{HcClient: hcClient}
}

func CampusGoClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateTasks 创建服务作业
//
// 该接口用于创建服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CampusGoClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

// CreateTasksInvoker 创建服务作业
func (c *CampusGoClient) CreateTasksInvoker(request *model.CreateTasksRequest) *CreateTasksInvoker {
	requestDef := GenReqDefForCreateTasks()
	return &CreateTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除服务作业
//
// 该接口用于删除服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CampusGoClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除服务作业
func (c *CampusGoClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasksDetails 获取服务作业列表
//
// 该接口用于获取服务作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CampusGoClient) ListTasksDetails(request *model.ListTasksDetailsRequest) (*model.ListTasksDetailsResponse, error) {
	requestDef := GenReqDefForListTasksDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksDetailsResponse), nil
	}
}

// ListTasksDetailsInvoker 获取服务作业列表
func (c *CampusGoClient) ListTasksDetailsInvoker(request *model.ListTasksDetailsRequest) *ListTasksDetailsInvoker {
	requestDef := GenReqDefForListTasksDetails()
	return &ListTasksDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 查询服务作业
//
// 该接口用于查询服务作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CampusGoClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 查询服务作业
func (c *CampusGoClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
