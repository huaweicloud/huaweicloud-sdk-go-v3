package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vas/v2/model"
)

type VasClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVasClient(hcClient *http_client.HcHttpClient) *VasClient {
	return &VasClient{HcClient: hcClient}
}

func VasClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 创建服务作业
//
// 该接口用于创建服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

// 删除服务作业
//
// 该接口用于删除服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// 获取服务作业列表
//
// 该接口用于获取服务作业列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) ListTasksDetails(request *model.ListTasksDetailsRequest) (*model.ListTasksDetailsResponse, error) {
	requestDef := GenReqDefForListTasksDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksDetailsResponse), nil
	}
}

// 查询服务作业
//
// 该接口用于查询服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// 启动服务作业
//
// 该接口用于启动服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) StartTask(request *model.StartTaskRequest) (*model.StartTaskResponse, error) {
	requestDef := GenReqDefForStartTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartTaskResponse), nil
	}
}

// 停止服务作业
//
// 该接口用于停止服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) StopTask(request *model.StopTaskRequest) (*model.StopTaskResponse, error) {
	requestDef := GenReqDefForStopTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskResponse), nil
	}
}

// 更新服务作业
//
// 该接口用于更新服务作业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *VasClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}
