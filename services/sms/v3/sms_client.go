package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sms/v3/model"
)

type SmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSmsClient(hcClient *http_client.HcHttpClient) *SmsClient {
	return &SmsClient{HcClient: hcClient}
}

func SmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 新建迁移项目
//
// 新建迁移项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) CreateMigproject(request *model.CreateMigprojectRequest) (*model.CreateMigprojectResponse, error) {
	requestDef := GenReqDefForCreateMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigprojectResponse), nil
	}
}

// 创建迁移任务
//
// 根据源端服务器创建一个迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// 新增模板信息
//
// 新增源端模板信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// 删除迁移项目
//
// 删除指定ID的迁移项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteMigproject(request *model.DeleteMigprojectRequest) (*model.DeleteMigprojectResponse, error) {
	requestDef := GenReqDefForDeleteMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigprojectResponse), nil
	}
}

// 删除指定ID的源端服务器信息
//
// 从主机迁移服务界面上删除指定ID的源端服务器信息。一旦源端服务器信息被删除，则只能通过重启源端服务器上的迁移Agent来将源端服务器信息重新添加在主机迁移服务界面。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteServer(request *model.DeleteServerRequest) (*model.DeleteServerResponse, error) {
	requestDef := GenReqDefForDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerResponse), nil
	}
}

// 批量删除源端服务器信息
//
// 批量删除源端服务器信息。一旦源端服务器信息被删除，则只能通过重启源端服务器上的迁移Agent来将源端服务器信息重新添加在主机迁移服务界面。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteServers(request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	requestDef := GenReqDefForDeleteServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServersResponse), nil
	}
}

// 删除指定ID的迁移任务
//
// 删除指定ID的迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// 批量删除迁移任务
//
// 批量删除迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteTasks(request *model.DeleteTasksRequest) (*model.DeleteTasksResponse, error) {
	requestDef := GenReqDefForDeleteTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTasksResponse), nil
	}
}

// 删除指定ID的模板
//
// 删除指定ID的模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// 批量删除指定ID的模板
//
// 批量删除指定ID的模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) DeleteTemplates(request *model.DeleteTemplatesRequest) (*model.DeleteTemplatesResponse, error) {
	requestDef := GenReqDefForDeleteTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplatesResponse), nil
	}
}

// 查询待迁移源端的所有错误
//
// 主机迁移过程中可能发生错误，使用该接口可以批量查询迁移过程中出现错误的源端服务器信息，以及它们的错误信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ListErrorServers(request *model.ListErrorServersRequest) (*model.ListErrorServersResponse, error) {
	requestDef := GenReqDefForListErrorServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorServersResponse), nil
	}
}

// 获取项目列表
//
// 主机迁移服务中可以使用迁移项目来对源端进行项目管理，使用该接口获取当前账户下所有的迁移项目列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ListMigprojects(request *model.ListMigprojectsRequest) (*model.ListMigprojectsResponse, error) {
	requestDef := GenReqDefForListMigprojects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigprojectsResponse), nil
	}
}

// 查询源端服务器列表
//
// 用户在源端安装并成功启动Agent后，Agent会将源端服务器信息注册在主机迁移服务中，调用该接口查询已注册的源端服务器列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ListServers(request *model.ListServersRequest) (*model.ListServersResponse, error) {
	requestDef := GenReqDefForListServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersResponse), nil
	}
}

// 查询迁移任务列表
//
// 在设置目的端后，主机迁移服务会自动创建迁移任务，使用该接口可以查询迁移任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// 查询模板列表
//
// 查询弹性云服务器模板列表，迁移时选择“新建服务器”时可使用该模板创建弹性云服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// 上报源端服务器基本信息
//
// 上报源端服务器信息，上报成功后会在sms服务器列表中看到对应的源端服务器信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) RegisterServer(request *model.RegisterServerRequest) (*model.RegisterServerResponse, error) {
	requestDef := GenReqDefForRegisterServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterServerResponse), nil
	}
}

// 获取服务端命令
//
// 迁移Agent调用该接口从SMS服务端获取下发给指定源端迁移Agent的命令。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowCommand(request *model.ShowCommandRequest) (*model.ShowCommandResponse, error) {
	requestDef := GenReqDefForShowCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommandResponse), nil
	}
}

// 查询指定ID迁移项目详情
//
// 查询指定ID的迁移项目详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowMigproject(request *model.ShowMigprojectRequest) (*model.ShowMigprojectResponse, error) {
	requestDef := GenReqDefForShowMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigprojectResponse), nil
	}
}

// 获取服务器总览
//
// 获取服务器总览
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowOverview(request *model.ShowOverviewRequest) (*model.ShowOverviewResponse, error) {
	requestDef := GenReqDefForShowOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOverviewResponse), nil
	}
}

// 查询指定ID的源端服务器
//
// 迁移Agent将源端服务器信息上报到主机迁移服务后，主机迁移服务会对迁移的可行性进行检测，该接口返回源端服务器的基本信息和检查结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

// 查询指定ID的迁移任务
//
// 查询指定ID的迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// 查询指定ID模板信息
//
// 查询指定ID的弹性云服务器模板信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowTemplate(request *model.ShowTemplateRequest) (*model.ShowTemplateResponse, error) {
	requestDef := GenReqDefForShowTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateResponse), nil
	}
}

// 查询任务限速规则
//
// 按时间段查询迁移任务的迁移速率
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) ShowsSpeedLimits(request *model.ShowsSpeedLimitsRequest) (*model.ShowsSpeedLimitsResponse, error) {
	requestDef := GenReqDefForShowsSpeedLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowsSpeedLimitsResponse), nil
	}
}

// 上报服务端命令执行结果
//
// 迁移Agent调用该接口向SMS服务端反馈指定指令的执行结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateCommandResult(request *model.UpdateCommandResultRequest) (*model.UpdateCommandResultResponse, error) {
	requestDef := GenReqDefForUpdateCommandResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCommandResultResponse), nil
	}
}

// 更新任务对应源端复制状态
//
// 更新任务对应源端复制状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateCopyState(request *model.UpdateCopyStateRequest) (*model.UpdateCopyStateResponse, error) {
	requestDef := GenReqDefForUpdateCopyState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCopyStateResponse), nil
	}
}

// 更新默认迁移项目
//
// 更改默认迁移项目，注册源端会注册在当前的默认项目下。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateDefaultMigproject(request *model.UpdateDefaultMigprojectRequest) (*model.UpdateDefaultMigprojectResponse, error) {
	requestDef := GenReqDefForUpdateDefaultMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDefaultMigprojectResponse), nil
	}
}

// 更新磁盘信息
//
// 更新服务器的磁盘信息，此接口会把服务器原有磁盘信息清空，然后更新成新磁盘信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateDiskInfo(request *model.UpdateDiskInfoRequest) (*model.UpdateDiskInfoResponse, error) {
	requestDef := GenReqDefForUpdateDiskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDiskInfoResponse), nil
	}
}

// 更新迁移项目信息
//
// 更新迁移项目的信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateMigproject(request *model.UpdateMigprojectRequest) (*model.UpdateMigprojectResponse, error) {
	requestDef := GenReqDefForUpdateMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMigprojectResponse), nil
	}
}

// 修改指定ID的源端服务器名称
//
// 该功能用来修改SMS服务端的源端名称，方便用户对源端进行管理。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateServerName(request *model.UpdateServerNameRequest) (*model.UpdateServerNameResponse, error) {
	requestDef := GenReqDefForUpdateServerName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerNameResponse), nil
	}
}

// 设置迁移限速规则
//
// 设置迁移任务的迁移速率。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateSpeed(request *model.UpdateSpeedRequest) (*model.UpdateSpeedResponse, error) {
	requestDef := GenReqDefForUpdateSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpeedResponse), nil
	}
}

// 更新指定ID的迁移任务
//
// 更新指定ID的迁移任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

// 上报数据迁移进度和速率
//
// 此接口由安装在源端服务器上的迁移Agent在数据迁移阶段调用，用来将迁移的具体进度上报给SMS服务端。
//
// 迁移Agent自动调用此接口用于上报数据迁移进度，您无需调用此接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateTaskSpeed(request *model.UpdateTaskSpeedRequest) (*model.UpdateTaskSpeedResponse, error) {
	requestDef := GenReqDefForUpdateTaskSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskSpeedResponse), nil
	}
}

// 管理迁移任务
//
// 管理迁移任务，包括启动任务，暂停任务，同步任务，日志上传，回滚失败迁移任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateTaskStatus(request *model.UpdateTaskStatusRequest) (*model.UpdateTaskStatusResponse, error) {
	requestDef := GenReqDefForUpdateTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskStatusResponse), nil
	}
}

// 修改模板信息
//
// 修改源端模板信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SmsClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}
