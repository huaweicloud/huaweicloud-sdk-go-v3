package v2

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/oms/v2/model"
)

type OmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewOmsClient(hcClient *http_client.HcHttpClient) *OmsClient {
	return &OmsClient{HcClient: hcClient}
}

func OmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//源端有对象需要进行同步时，调用该接口创建一个同步事件，系统将根据同步事件中包含的对象名称进行同步
func (c *OmsClient) CreateSyncEvents(request *model.CreateSyncEventsRequest) (*model.CreateSyncEventsResponse, error) {
	requestDef := GenReqDefForCreateSyncEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSyncEventsResponse), nil
	}
}

//创建迁移任务，创建成功后，任务会被自动启动，不需要额外调用启动任务命令。
func (c *OmsClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

//调用该接口删除迁移任务。 正在运行的任务不允许删除，如果删除会返回失败；若要删除，请先行暂停任务。
func (c *OmsClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

//查询用户账户下的所有任务信息。
func (c *OmsClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

//查询指定ID的任务详情。
func (c *OmsClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

//迁移任务暂停或失败后，调用该接口以启动任务。
func (c *OmsClient) StartTask(request *model.StartTaskRequest) (*model.StartTaskResponse, error) {
	requestDef := GenReqDefForStartTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartTaskResponse), nil
	}
}

//当迁移任务处于迁移中时，调用该接口停止任务。
func (c *OmsClient) StopTask(request *model.StopTaskRequest) (*model.StopTaskResponse, error) {
	requestDef := GenReqDefForStopTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskResponse), nil
	}
}

//当迁移任务未执行完成时，修改迁移任务的流量控制策略。
func (c *OmsClient) UpdateBandwidthPolicy(request *model.UpdateBandwidthPolicyRequest) (*model.UpdateBandwidthPolicyResponse, error) {
	requestDef := GenReqDefForUpdateBandwidthPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBandwidthPolicyResponse), nil
	}
}

//查询对象存储迁移服务的API版本信息。
func (c *OmsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

//查询对象存储迁移服务指定API版本信息。
func (c *OmsClient) ShowApiInfo(request *model.ShowApiInfoRequest) (*model.ShowApiInfoResponse, error) {
	requestDef := GenReqDefForShowApiInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiInfoResponse), nil
	}
}
