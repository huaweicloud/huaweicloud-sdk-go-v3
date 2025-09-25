package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsdeploy/v2/model"
)

type CodeArtsDeployClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsDeployClient(hcClient *httpclient.HcHttpClient) *CodeArtsDeployClient {
	return &CodeArtsDeployClient{HcClient: hcClient}
}

func CodeArtsDeployClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAppGroups 创建分组
//
// 创建分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateAppGroups(request *model.CreateAppGroupsRequest) (*model.CreateAppGroupsResponse, error) {
	requestDef := GenReqDefForCreateAppGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppGroupsResponse), nil
	}
}

// CreateAppGroupsInvoker 创建分组
func (c *CodeArtsDeployClient) CreateAppGroupsInvoker(request *model.CreateAppGroupsRequest) *CreateAppGroupsInvoker {
	requestDef := GenReqDefForCreateAppGroups()
	return &CreateAppGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppGroups 删除分组
//
// 删除分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteAppGroups(request *model.DeleteAppGroupsRequest) (*model.DeleteAppGroupsResponse, error) {
	requestDef := GenReqDefForDeleteAppGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppGroupsResponse), nil
	}
}

// DeleteAppGroupsInvoker 删除分组
func (c *CodeArtsDeployClient) DeleteAppGroupsInvoker(request *model.DeleteAppGroupsRequest) *DeleteAppGroupsInvoker {
	requestDef := GenReqDefForDeleteAppGroups()
	return &DeleteAppGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppGroups 查询分组列表
//
// 查询分组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListAppGroups(request *model.ListAppGroupsRequest) (*model.ListAppGroupsResponse, error) {
	requestDef := GenReqDefForListAppGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppGroupsResponse), nil
	}
}

// ListAppGroupsInvoker 查询分组列表
func (c *CodeArtsDeployClient) ListAppGroupsInvoker(request *model.ListAppGroupsRequest) *ListAppGroupsInvoker {
	requestDef := GenReqDefForListAppGroups()
	return &ListAppGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MoveAppGroups 移动分组
//
// 往上或者往下移动单个分组,用来在页面上调整分组位置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) MoveAppGroups(request *model.MoveAppGroupsRequest) (*model.MoveAppGroupsResponse, error) {
	requestDef := GenReqDefForMoveAppGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MoveAppGroupsResponse), nil
	}
}

// MoveAppGroupsInvoker 移动分组
func (c *CodeArtsDeployClient) MoveAppGroupsInvoker(request *model.MoveAppGroupsRequest) *MoveAppGroupsInvoker {
	requestDef := GenReqDefForMoveAppGroups()
	return &MoveAppGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MoveAppToGroup 移动应用至指定分组
//
// 将应用移动至指定分组（支持批量）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) MoveAppToGroup(request *model.MoveAppToGroupRequest) (*model.MoveAppToGroupResponse, error) {
	requestDef := GenReqDefForMoveAppToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MoveAppToGroupResponse), nil
	}
}

// MoveAppToGroupInvoker 移动应用至指定分组
func (c *CodeArtsDeployClient) MoveAppToGroupInvoker(request *model.MoveAppToGroupRequest) *MoveAppToGroupInvoker {
	requestDef := GenReqDefForMoveAppToGroup()
	return &MoveAppToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppGroups 修改分组
//
// 修改分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateAppGroups(request *model.UpdateAppGroupsRequest) (*model.UpdateAppGroupsResponse, error) {
	requestDef := GenReqDefForUpdateAppGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppGroupsResponse), nil
	}
}

// UpdateAppGroupsInvoker 修改分组
func (c *CodeArtsDeployClient) UpdateAppGroupsInvoker(request *model.UpdateAppGroupsRequest) *UpdateAppGroupsInvoker {
	requestDef := GenReqDefForUpdateAppGroups()
	return &UpdateAppGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateApplicationPermissions 批量修改应用权限
//
// 批量修改应用权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) BatchUpdateApplicationPermissions(request *model.BatchUpdateApplicationPermissionsRequest) (*model.BatchUpdateApplicationPermissionsResponse, error) {
	requestDef := GenReqDefForBatchUpdateApplicationPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateApplicationPermissionsResponse), nil
	}
}

// BatchUpdateApplicationPermissionsInvoker 批量修改应用权限
func (c *CodeArtsDeployClient) BatchUpdateApplicationPermissionsInvoker(request *model.BatchUpdateApplicationPermissionsRequest) *BatchUpdateApplicationPermissionsInvoker {
	requestDef := GenReqDefForBatchUpdateApplicationPermissions()
	return &BatchUpdateApplicationPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdatePermissionLevel 批量配置应用下鉴权级别
//
// 批量配置应用下鉴权级别为项目级或实例级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) BatchUpdatePermissionLevel(request *model.BatchUpdatePermissionLevelRequest) (*model.BatchUpdatePermissionLevelResponse, error) {
	requestDef := GenReqDefForBatchUpdatePermissionLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePermissionLevelResponse), nil
	}
}

// BatchUpdatePermissionLevelInvoker 批量配置应用下鉴权级别
func (c *CodeArtsDeployClient) BatchUpdatePermissionLevelInvoker(request *model.BatchUpdatePermissionLevelRequest) *BatchUpdatePermissionLevelInvoker {
	requestDef := GenReqDefForBatchUpdatePermissionLevel()
	return &BatchUpdatePermissionLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckCanCreate 查询当前用户是否有项目下创建应用权限
//
// 查询当前用户是否有项目下创建应用权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CheckCanCreate(request *model.CheckCanCreateRequest) (*model.CheckCanCreateResponse, error) {
	requestDef := GenReqDefForCheckCanCreate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCanCreateResponse), nil
	}
}

// CheckCanCreateInvoker 查询当前用户是否有项目下创建应用权限
func (c *CodeArtsDeployClient) CheckCanCreateInvoker(request *model.CheckCanCreateRequest) *CheckCanCreateInvoker {
	requestDef := GenReqDefForCheckCanCreate()
	return &CheckCanCreateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationPermissions 查询应用实例级/项目级权限矩阵
//
// 查询应用实例级/项目级权限矩阵，传递app_id时，查询应用实例级权限矩阵；未传app_id，传递project_id时，查询应用项目级权限矩阵。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListApplicationPermissions(request *model.ListApplicationPermissionsRequest) (*model.ListApplicationPermissionsResponse, error) {
	requestDef := GenReqDefForListApplicationPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationPermissionsResponse), nil
	}
}

// ListApplicationPermissionsInvoker 查询应用实例级/项目级权限矩阵
func (c *CodeArtsDeployClient) ListApplicationPermissionsInvoker(request *model.ListApplicationPermissionsRequest) *ListApplicationPermissionsInvoker {
	requestDef := GenReqDefForListApplicationPermissions()
	return &ListApplicationPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteApp 批量删除项目下应用
//
// 批量删除项目下应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) BatchDeleteApp(request *model.BatchDeleteAppRequest) (*model.BatchDeleteAppResponse, error) {
	requestDef := GenReqDefForBatchDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppResponse), nil
	}
}

// BatchDeleteAppInvoker 批量删除项目下应用
func (c *CodeArtsDeployClient) BatchDeleteAppInvoker(request *model.BatchDeleteAppRequest) *BatchDeleteAppInvoker {
	requestDef := GenReqDefForBatchDeleteApp()
	return &BatchDeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckIsDuplicateAppName 查询项目下是否存在同名应用
//
// 查询项目下是否存在同名应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CheckIsDuplicateAppName(request *model.CheckIsDuplicateAppNameRequest) (*model.CheckIsDuplicateAppNameResponse, error) {
	requestDef := GenReqDefForCheckIsDuplicateAppName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckIsDuplicateAppNameResponse), nil
	}
}

// CheckIsDuplicateAppNameInvoker 查询项目下是否存在同名应用
func (c *CodeArtsDeployClient) CheckIsDuplicateAppNameInvoker(request *model.CheckIsDuplicateAppNameRequest) *CheckIsDuplicateAppNameInvoker {
	requestDef := GenReqDefForCheckIsDuplicateAppName()
	return &CheckIsDuplicateAppNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyApplication 复制应用
//
// 复制应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CopyApplication(request *model.CopyApplicationRequest) (*model.CopyApplicationResponse, error) {
	requestDef := GenReqDefForCopyApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyApplicationResponse), nil
	}
}

// CopyApplicationInvoker 复制应用
func (c *CodeArtsDeployClient) CopyApplicationInvoker(request *model.CopyApplicationRequest) *CopyApplicationInvoker {
	requestDef := GenReqDefForCopyApplication()
	return &CopyApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 新建应用
//
// 新建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 新建应用
func (c *CodeArtsDeployClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeployTaskByTemplate 通过模板新建应用
//
// 通过模板新建应用。该接口于2024年09月30日后不再维护，推荐使用新版CreateApp接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateDeployTaskByTemplate(request *model.CreateDeployTaskByTemplateRequest) (*model.CreateDeployTaskByTemplateResponse, error) {
	requestDef := GenReqDefForCreateDeployTaskByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeployTaskByTemplateResponse), nil
	}
}

// CreateDeployTaskByTemplateInvoker 通过模板新建应用
func (c *CodeArtsDeployClient) CreateDeployTaskByTemplateInvoker(request *model.CreateDeployTaskByTemplateRequest) *CreateDeployTaskByTemplateInvoker {
	requestDef := GenReqDefForCreateDeployTaskByTemplate()
	return &CreateDeployTaskByTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplication 删除应用
//
// 根据应用id删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// DeleteApplicationInvoker 删除应用
func (c *CodeArtsDeployClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployTask 删除应用
//
// 根据部署任务id删除应用。该接口于2024年09月30日后不再维护，推荐使用新版DeleteApplication接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteDeployTask(request *model.DeleteDeployTaskRequest) (*model.DeleteDeployTaskResponse, error) {
	requestDef := GenReqDefForDeleteDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployTaskResponse), nil
	}
}

// DeleteDeployTaskInvoker 删除应用
func (c *CodeArtsDeployClient) DeleteDeployTaskInvoker(request *model.DeleteDeployTaskRequest) *DeleteDeployTaskInvoker {
	requestDef := GenReqDefForDeleteDeployTask()
	return &DeleteDeployTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllApp 获取应用列表
//
// 查询项目下应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListAllApp(request *model.ListAllAppRequest) (*model.ListAllAppResponse, error) {
	requestDef := GenReqDefForListAllApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllAppResponse), nil
	}
}

// ListAllAppInvoker 获取应用列表
func (c *CodeArtsDeployClient) ListAllAppInvoker(request *model.ListAllAppRequest) *ListAllAppInvoker {
	requestDef := GenReqDefForListAllApp()
	return &ListAllAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployTaskHistoryByDate 根据开始时间和结束时间查询项目下指定应用的历史部署记录列表
//
// 根据开始时间和结束时间查询项目下指定应用的历史部署记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListDeployTaskHistoryByDate(request *model.ListDeployTaskHistoryByDateRequest) (*model.ListDeployTaskHistoryByDateResponse, error) {
	requestDef := GenReqDefForListDeployTaskHistoryByDate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeployTaskHistoryByDateResponse), nil
	}
}

// ListDeployTaskHistoryByDateInvoker 根据开始时间和结束时间查询项目下指定应用的历史部署记录列表
func (c *CodeArtsDeployClient) ListDeployTaskHistoryByDateInvoker(request *model.ListDeployTaskHistoryByDateRequest) *ListDeployTaskHistoryByDateInvoker {
	requestDef := GenReqDefForListDeployTaskHistoryByDate()
	return &ListDeployTaskHistoryByDateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployTasks 获取应用列表
//
// 查询项目下应用列表。该接口于2024年09月30日后不再维护，推荐使用新版ListAllApp接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListDeployTasks(request *model.ListDeployTasksRequest) (*model.ListDeployTasksResponse, error) {
	requestDef := GenReqDefForListDeployTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeployTasksResponse), nil
	}
}

// ListDeployTasksInvoker 获取应用列表
func (c *CodeArtsDeployClient) ListDeployTasksInvoker(request *model.ListDeployTasksRequest) *ListDeployTasksInvoker {
	requestDef := GenReqDefForListDeployTasks()
	return &ListDeployTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppDetailById 获取应用详情
//
// 根据应用id获取应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowAppDetailById(request *model.ShowAppDetailByIdRequest) (*model.ShowAppDetailByIdResponse, error) {
	requestDef := GenReqDefForShowAppDetailById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppDetailByIdResponse), nil
	}
}

// ShowAppDetailByIdInvoker 获取应用详情
func (c *CodeArtsDeployClient) ShowAppDetailByIdInvoker(request *model.ShowAppDetailByIdRequest) *ShowAppDetailByIdInvoker {
	requestDef := GenReqDefForShowAppDetailById()
	return &ShowAppDetailByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeployTaskDetail 获取应用详情
//
// 根据部署任务id获取应用详情。该接口于2024年09月30日后不再维护，推荐使用新版ShowAppDetailById接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowDeployTaskDetail(request *model.ShowDeployTaskDetailRequest) (*model.ShowDeployTaskDetailResponse, error) {
	requestDef := GenReqDefForShowDeployTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeployTaskDetailResponse), nil
	}
}

// ShowDeployTaskDetailInvoker 获取应用详情
func (c *CodeArtsDeployClient) ShowDeployTaskDetailInvoker(request *model.ShowDeployTaskDetailRequest) *ShowDeployTaskDetailInvoker {
	requestDef := GenReqDefForShowDeployTaskDetail()
	return &ShowDeployTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecutionParams 查询部署记录的执行参数
//
// 查询部署记录的执行参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowExecutionParams(request *model.ShowExecutionParamsRequest) (*model.ShowExecutionParamsResponse, error) {
	requestDef := GenReqDefForShowExecutionParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecutionParamsResponse), nil
	}
}

// ShowExecutionParamsInvoker 查询部署记录的执行参数
func (c *CodeArtsDeployClient) ShowExecutionParamsInvoker(request *model.ShowExecutionParamsRequest) *ShowExecutionParamsInvoker {
	requestDef := GenReqDefForShowExecutionParams()
	return &ShowExecutionParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDeployTask 部署应用
//
// 根据部署任务id部署应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) StartDeployTask(request *model.StartDeployTaskRequest) (*model.StartDeployTaskResponse, error) {
	requestDef := GenReqDefForStartDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDeployTaskResponse), nil
	}
}

// StartDeployTaskInvoker 部署应用
func (c *CodeArtsDeployClient) StartDeployTaskInvoker(request *model.StartDeployTaskRequest) *StartDeployTaskInvoker {
	requestDef := GenReqDefForStartDeployTask()
	return &StartDeployTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppDisableStatus 禁用/取消禁用应用
//
// 禁用/取消禁用应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateAppDisableStatus(request *model.UpdateAppDisableStatusRequest) (*model.UpdateAppDisableStatusResponse, error) {
	requestDef := GenReqDefForUpdateAppDisableStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppDisableStatusResponse), nil
	}
}

// UpdateAppDisableStatusInvoker 禁用/取消禁用应用
func (c *CodeArtsDeployClient) UpdateAppDisableStatusInvoker(request *model.UpdateAppDisableStatusRequest) *UpdateAppDisableStatusInvoker {
	requestDef := GenReqDefForUpdateAppDisableStatus()
	return &UpdateAppDisableStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppInfo 更新应用
//
// 更新应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateAppInfo(request *model.UpdateAppInfoRequest) (*model.UpdateAppInfoResponse, error) {
	requestDef := GenReqDefForUpdateAppInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppInfoResponse), nil
	}
}

// UpdateAppInfoInvoker 更新应用
func (c *CodeArtsDeployClient) UpdateAppInfoInvoker(request *model.UpdateAppInfoRequest) *UpdateAppInfoInvoker {
	requestDef := GenReqDefForUpdateAppInfo()
	return &UpdateAppInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironment 应用下创建环境
//
// 应用下创建环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateEnvironment(request *model.CreateEnvironmentRequest) (*model.CreateEnvironmentResponse, error) {
	requestDef := GenReqDefForCreateEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentResponse), nil
	}
}

// CreateEnvironmentInvoker 应用下创建环境
func (c *CodeArtsDeployClient) CreateEnvironmentInvoker(request *model.CreateEnvironmentRequest) *CreateEnvironmentInvoker {
	requestDef := GenReqDefForCreateEnvironment()
	return &CreateEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironment 删除应用下的环境
//
// 删除应用下的环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteEnvironment(request *model.DeleteEnvironmentRequest) (*model.DeleteEnvironmentResponse, error) {
	requestDef := GenReqDefForDeleteEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentResponse), nil
	}
}

// DeleteEnvironmentInvoker 删除应用下的环境
func (c *CodeArtsDeployClient) DeleteEnvironmentInvoker(request *model.DeleteEnvironmentRequest) *DeleteEnvironmentInvoker {
	requestDef := GenReqDefForDeleteEnvironment()
	return &DeleteEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHostFromEnvironment 环境下删除主机
//
// 环境下删除主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteHostFromEnvironment(request *model.DeleteHostFromEnvironmentRequest) (*model.DeleteHostFromEnvironmentResponse, error) {
	requestDef := GenReqDefForDeleteHostFromEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostFromEnvironmentResponse), nil
	}
}

// DeleteHostFromEnvironmentInvoker 环境下删除主机
func (c *CodeArtsDeployClient) DeleteHostFromEnvironmentInvoker(request *model.DeleteHostFromEnvironmentRequest) *DeleteHostFromEnvironmentInvoker {
	requestDef := GenReqDefForDeleteHostFromEnvironment()
	return &DeleteHostFromEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportHostToEnvironment 环境下导入主机
//
// 环境下导入主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ImportHostToEnvironment(request *model.ImportHostToEnvironmentRequest) (*model.ImportHostToEnvironmentResponse, error) {
	requestDef := GenReqDefForImportHostToEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportHostToEnvironmentResponse), nil
	}
}

// ImportHostToEnvironmentInvoker 环境下导入主机
func (c *CodeArtsDeployClient) ImportHostToEnvironmentInvoker(request *model.ImportHostToEnvironmentRequest) *ImportHostToEnvironmentInvoker {
	requestDef := GenReqDefForImportHostToEnvironment()
	return &ImportHostToEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentHosts 查询环境内的主机列表
//
// 查询环境内的主机列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListEnvironmentHosts(request *model.ListEnvironmentHostsRequest) (*model.ListEnvironmentHostsResponse, error) {
	requestDef := GenReqDefForListEnvironmentHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentHostsResponse), nil
	}
}

// ListEnvironmentHostsInvoker 查询环境内的主机列表
func (c *CodeArtsDeployClient) ListEnvironmentHostsInvoker(request *model.ListEnvironmentHostsRequest) *ListEnvironmentHostsInvoker {
	requestDef := GenReqDefForListEnvironmentHosts()
	return &ListEnvironmentHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironments 查询应用下环境列表
//
// 查询应用下环境列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}

// ListEnvironmentsInvoker 查询应用下环境列表
func (c *CodeArtsDeployClient) ListEnvironmentsInvoker(request *model.ListEnvironmentsRequest) *ListEnvironmentsInvoker {
	requestDef := GenReqDefForListEnvironments()
	return &ListEnvironmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnvironmentDetail 查询环境详情
//
// 查询环境详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowEnvironmentDetail(request *model.ShowEnvironmentDetailRequest) (*model.ShowEnvironmentDetailResponse, error) {
	requestDef := GenReqDefForShowEnvironmentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvironmentDetailResponse), nil
	}
}

// ShowEnvironmentDetailInvoker 查询环境详情
func (c *CodeArtsDeployClient) ShowEnvironmentDetailInvoker(request *model.ShowEnvironmentDetailRequest) *ShowEnvironmentDetailInvoker {
	requestDef := GenReqDefForShowEnvironmentDetail()
	return &ShowEnvironmentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironment 应用下编辑环境
//
// 应用下编辑环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateEnvironment(request *model.UpdateEnvironmentRequest) (*model.UpdateEnvironmentResponse, error) {
	requestDef := GenReqDefForUpdateEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentResponse), nil
	}
}

// UpdateEnvironmentInvoker 应用下编辑环境
func (c *CodeArtsDeployClient) UpdateEnvironmentInvoker(request *model.UpdateEnvironmentRequest) *UpdateEnvironmentInvoker {
	requestDef := GenReqDefForUpdateEnvironment()
	return &UpdateEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentPermissions 查询环境权限
//
// 查询环境权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListEnvironmentPermissions(request *model.ListEnvironmentPermissionsRequest) (*model.ListEnvironmentPermissionsResponse, error) {
	requestDef := GenReqDefForListEnvironmentPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentPermissionsResponse), nil
	}
}

// ListEnvironmentPermissionsInvoker 查询环境权限
func (c *CodeArtsDeployClient) ListEnvironmentPermissionsInvoker(request *model.ListEnvironmentPermissionsRequest) *ListEnvironmentPermissionsInvoker {
	requestDef := GenReqDefForListEnvironmentPermissions()
	return &ListEnvironmentPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironmentPermission 编辑环境权限
//
// 编辑环境权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateEnvironmentPermission(request *model.UpdateEnvironmentPermissionRequest) (*model.UpdateEnvironmentPermissionResponse, error) {
	requestDef := GenReqDefForUpdateEnvironmentPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentPermissionResponse), nil
	}
}

// UpdateEnvironmentPermissionInvoker 编辑环境权限
func (c *CodeArtsDeployClient) UpdateEnvironmentPermissionInvoker(request *model.UpdateEnvironmentPermissionRequest) *UpdateEnvironmentPermissionInvoker {
	requestDef := GenReqDefForUpdateEnvironmentPermission()
	return &UpdateEnvironmentPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteHosts 批量删除主机集群下的主机
//
// 批量删除主机集群下的主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) BatchDeleteHosts(request *model.BatchDeleteHostsRequest) (*model.BatchDeleteHostsResponse, error) {
	requestDef := GenReqDefForBatchDeleteHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteHostsResponse), nil
	}
}

// BatchDeleteHostsInvoker 批量删除主机集群下的主机
func (c *CodeArtsDeployClient) BatchDeleteHostsInvoker(request *model.BatchDeleteHostsRequest) *BatchDeleteHostsInvoker {
	requestDef := GenReqDefForBatchDeleteHosts()
	return &BatchDeleteHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyHostsToTarget 批量复制主机至目标主机集群
//
// 批量复制主机至目标主机集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CopyHostsToTarget(request *model.CopyHostsToTargetRequest) (*model.CopyHostsToTargetResponse, error) {
	requestDef := GenReqDefForCopyHostsToTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyHostsToTargetResponse), nil
	}
}

// CopyHostsToTargetInvoker 批量复制主机至目标主机集群
func (c *CodeArtsDeployClient) CopyHostsToTargetInvoker(request *model.CopyHostsToTargetRequest) *CopyHostsToTargetInvoker {
	requestDef := GenReqDefForCopyHostsToTarget()
	return &CopyHostsToTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeploymentHost 新建主机
//
// 在指定主机集群下新建主机。该接口于2024年09月30日后不再维护，推荐使用新版CreateHost接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateDeploymentHost(request *model.CreateDeploymentHostRequest) (*model.CreateDeploymentHostResponse, error) {
	requestDef := GenReqDefForCreateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentHostResponse), nil
	}
}

// CreateDeploymentHostInvoker 新建主机
func (c *CodeArtsDeployClient) CreateDeploymentHostInvoker(request *model.CreateDeploymentHostRequest) *CreateDeploymentHostInvoker {
	requestDef := GenReqDefForCreateDeploymentHost()
	return &CreateDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHost 新建主机
//
// 在指定主机集群下新建主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateHost(request *model.CreateHostRequest) (*model.CreateHostResponse, error) {
	requestDef := GenReqDefForCreateHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostResponse), nil
	}
}

// CreateHostInvoker 新建主机
func (c *CodeArtsDeployClient) CreateHostInvoker(request *model.CreateHostRequest) *CreateHostInvoker {
	requestDef := GenReqDefForCreateHost()
	return &CreateHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeploymentHost 删除主机
//
// 根据主机id删除主机。该接口于2024年9月30日后不再维护。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteDeploymentHost(request *model.DeleteDeploymentHostRequest) (*model.DeleteDeploymentHostResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentHostResponse), nil
	}
}

// DeleteDeploymentHostInvoker 删除主机
func (c *CodeArtsDeployClient) DeleteDeploymentHostInvoker(request *model.DeleteDeploymentHostRequest) *DeleteDeploymentHostInvoker {
	requestDef := GenReqDefForDeleteDeploymentHost()
	return &DeleteDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHost 删除主机集群下主机
//
// 根据主机id删除主机集群下主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteHost(request *model.DeleteHostRequest) (*model.DeleteHostResponse, error) {
	requestDef := GenReqDefForDeleteHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostResponse), nil
	}
}

// DeleteHostInvoker 删除主机集群下主机
func (c *CodeArtsDeployClient) DeleteHostInvoker(request *model.DeleteHostRequest) *DeleteHostInvoker {
	requestDef := GenReqDefForDeleteHost()
	return &DeleteHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHosts 查询主机列表
//
// 根据主机集群id查询指定主机集群下的主机列表。该接口于2024年09月30日后不再维护，推荐使用新版ListNewHosts接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

// ListHostsInvoker 查询主机列表
func (c *CodeArtsDeployClient) ListHostsInvoker(request *model.ListHostsRequest) *ListHostsInvoker {
	requestDef := GenReqDefForListHosts()
	return &ListHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNewHosts 查询主机列表
//
// 根据主机集群id查询指定主机集群下的主机列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListNewHosts(request *model.ListNewHostsRequest) (*model.ListNewHostsResponse, error) {
	requestDef := GenReqDefForListNewHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNewHostsResponse), nil
	}
}

// ListNewHostsInvoker 查询主机列表
func (c *CodeArtsDeployClient) ListNewHostsInvoker(request *model.ListNewHostsRequest) *ListNewHostsInvoker {
	requestDef := GenReqDefForListNewHosts()
	return &ListNewHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentHostDetail 查询主机详情
//
// 根据主机id查询主机详情。该接口于2024年09月30日后不再维护，推荐使用新版ShowHostDetail接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowDeploymentHostDetail(request *model.ShowDeploymentHostDetailRequest) (*model.ShowDeploymentHostDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentHostDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentHostDetailResponse), nil
	}
}

// ShowDeploymentHostDetailInvoker 查询主机详情
func (c *CodeArtsDeployClient) ShowDeploymentHostDetailInvoker(request *model.ShowDeploymentHostDetailRequest) *ShowDeploymentHostDetailInvoker {
	requestDef := GenReqDefForShowDeploymentHostDetail()
	return &ShowDeploymentHostDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostDetail 查询主机详情
//
// 根据主机id查询主机详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowHostDetail(request *model.ShowHostDetailRequest) (*model.ShowHostDetailResponse, error) {
	requestDef := GenReqDefForShowHostDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostDetailResponse), nil
	}
}

// ShowHostDetailInvoker 查询主机详情
func (c *CodeArtsDeployClient) ShowHostDetailInvoker(request *model.ShowHostDetailRequest) *ShowHostDetailInvoker {
	requestDef := GenReqDefForShowHostDetail()
	return &ShowHostDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeploymentHost 修改主机
//
// 根据主机id修改主机信息。该接口于2024年9月30日后不再维护。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateDeploymentHost(request *model.UpdateDeploymentHostRequest) (*model.UpdateDeploymentHostResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentHostResponse), nil
	}
}

// UpdateDeploymentHostInvoker 修改主机
func (c *CodeArtsDeployClient) UpdateDeploymentHostInvoker(request *model.UpdateDeploymentHostRequest) *UpdateDeploymentHostInvoker {
	requestDef := GenReqDefForUpdateDeploymentHost()
	return &UpdateDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostInfo 编辑主机集群下主机信息
//
// 根据主机id编辑主机集群下主机信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateHostInfo(request *model.UpdateHostInfoRequest) (*model.UpdateHostInfoResponse, error) {
	requestDef := GenReqDefForUpdateHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostInfoResponse), nil
	}
}

// UpdateHostInfoInvoker 编辑主机集群下主机信息
func (c *CodeArtsDeployClient) UpdateHostInfoInvoker(request *model.UpdateHostInfoRequest) *UpdateHostInfoInvoker {
	requestDef := GenReqDefForUpdateHostInfo()
	return &UpdateHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeploymentGroup 新建主机集群
//
// 在项目下新建主机集群。该接口于2024年09月30日后不再维护，推荐使用新版CreateHostCluster接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateDeploymentGroup(request *model.CreateDeploymentGroupRequest) (*model.CreateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForCreateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentGroupResponse), nil
	}
}

// CreateDeploymentGroupInvoker 新建主机集群
func (c *CodeArtsDeployClient) CreateDeploymentGroupInvoker(request *model.CreateDeploymentGroupRequest) *CreateDeploymentGroupInvoker {
	requestDef := GenReqDefForCreateDeploymentGroup()
	return &CreateDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHostCluster 新建主机集群
//
// 在项目下新建主机集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CreateHostCluster(request *model.CreateHostClusterRequest) (*model.CreateHostClusterResponse, error) {
	requestDef := GenReqDefForCreateHostCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostClusterResponse), nil
	}
}

// CreateHostClusterInvoker 新建主机集群
func (c *CodeArtsDeployClient) CreateHostClusterInvoker(request *model.CreateHostClusterRequest) *CreateHostClusterInvoker {
	requestDef := GenReqDefForCreateHostCluster()
	return &CreateHostClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeploymentGroup 删除主机集群
//
// 根据主机集群id删除主机集群。该接口于2024年9月30日后不再维护。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteDeploymentGroup(request *model.DeleteDeploymentGroupRequest) (*model.DeleteDeploymentGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentGroupResponse), nil
	}
}

// DeleteDeploymentGroupInvoker 删除主机集群
func (c *CodeArtsDeployClient) DeleteDeploymentGroupInvoker(request *model.DeleteDeploymentGroupRequest) *DeleteDeploymentGroupInvoker {
	requestDef := GenReqDefForDeleteDeploymentGroup()
	return &DeleteDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHostCluster 删除主机集群
//
// 根据主机集群id删除主机集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) DeleteHostCluster(request *model.DeleteHostClusterRequest) (*model.DeleteHostClusterResponse, error) {
	requestDef := GenReqDefForDeleteHostCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostClusterResponse), nil
	}
}

// DeleteHostClusterInvoker 删除主机集群
func (c *CodeArtsDeployClient) DeleteHostClusterInvoker(request *model.DeleteHostClusterRequest) *DeleteHostClusterInvoker {
	requestDef := GenReqDefForDeleteHostCluster()
	return &DeleteHostClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssociateEnvironmentsInfos 查询主机集群关联环境信息
//
// 查询主机集群关联环境信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListAssociateEnvironmentsInfos(request *model.ListAssociateEnvironmentsInfosRequest) (*model.ListAssociateEnvironmentsInfosResponse, error) {
	requestDef := GenReqDefForListAssociateEnvironmentsInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssociateEnvironmentsInfosResponse), nil
	}
}

// ListAssociateEnvironmentsInfosInvoker 查询主机集群关联环境信息
func (c *CodeArtsDeployClient) ListAssociateEnvironmentsInfosInvoker(request *model.ListAssociateEnvironmentsInfosRequest) *ListAssociateEnvironmentsInfosInvoker {
	requestDef := GenReqDefForListAssociateEnvironmentsInfos()
	return &ListAssociateEnvironmentsInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostClusters 查询主机集群列表
//
// 按条件查询主机集群列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListHostClusters(request *model.ListHostClustersRequest) (*model.ListHostClustersResponse, error) {
	requestDef := GenReqDefForListHostClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostClustersResponse), nil
	}
}

// ListHostClustersInvoker 查询主机集群列表
func (c *CodeArtsDeployClient) ListHostClustersInvoker(request *model.ListHostClustersRequest) *ListHostClustersInvoker {
	requestDef := GenReqDefForListHostClusters()
	return &ListHostClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroupBaseInfos 查询应用下环境基本信息列表
//
// 查询应用下环境基本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListHostGroupBaseInfos(request *model.ListHostGroupBaseInfosRequest) (*model.ListHostGroupBaseInfosResponse, error) {
	requestDef := GenReqDefForListHostGroupBaseInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupBaseInfosResponse), nil
	}
}

// ListHostGroupBaseInfosInvoker 查询应用下环境基本信息列表
func (c *CodeArtsDeployClient) ListHostGroupBaseInfosInvoker(request *model.ListHostGroupBaseInfosRequest) *ListHostGroupBaseInfosInvoker {
	requestDef := GenReqDefForListHostGroupBaseInfos()
	return &ListHostGroupBaseInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroups 查询主机集群列表
//
// 按条件查询主机集群列表。该接口于2024年09月30日后不再维护，推荐使用新版ListHostClusters接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListHostGroups(request *model.ListHostGroupsRequest) (*model.ListHostGroupsResponse, error) {
	requestDef := GenReqDefForListHostGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupsResponse), nil
	}
}

// ListHostGroupsInvoker 查询主机集群列表
func (c *CodeArtsDeployClient) ListHostGroupsInvoker(request *model.ListHostGroupsRequest) *ListHostGroupsInvoker {
	requestDef := GenReqDefForListHostGroups()
	return &ListHostGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentGroupDetail 查询主机集群详情
//
// 根据主机集群id查询主机集群详情。该接口于2024年09月30日后不再维护，推荐使用新版ShowHostClusterDetail接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowDeploymentGroupDetail(request *model.ShowDeploymentGroupDetailRequest) (*model.ShowDeploymentGroupDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentGroupDetailResponse), nil
	}
}

// ShowDeploymentGroupDetailInvoker 查询主机集群详情
func (c *CodeArtsDeployClient) ShowDeploymentGroupDetailInvoker(request *model.ShowDeploymentGroupDetailRequest) *ShowDeploymentGroupDetailInvoker {
	requestDef := GenReqDefForShowDeploymentGroupDetail()
	return &ShowDeploymentGroupDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostClusterDetail 查询主机集群详情
//
// 根据主机集群id查询主机集群详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowHostClusterDetail(request *model.ShowHostClusterDetailRequest) (*model.ShowHostClusterDetailResponse, error) {
	requestDef := GenReqDefForShowHostClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostClusterDetailResponse), nil
	}
}

// ShowHostClusterDetailInvoker 查询主机集群详情
func (c *CodeArtsDeployClient) ShowHostClusterDetailInvoker(request *model.ShowHostClusterDetailRequest) *ShowHostClusterDetailInvoker {
	requestDef := GenReqDefForShowHostClusterDetail()
	return &ShowHostClusterDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeploymentGroup 修改主机集群
//
// 根据主机集群id修改主机集群信息。该接口于2024年9月30日后不再维护。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateDeploymentGroup(request *model.UpdateDeploymentGroupRequest) (*model.UpdateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentGroupResponse), nil
	}
}

// UpdateDeploymentGroupInvoker 修改主机集群
func (c *CodeArtsDeployClient) UpdateDeploymentGroupInvoker(request *model.UpdateDeploymentGroupRequest) *UpdateDeploymentGroupInvoker {
	requestDef := GenReqDefForUpdateDeploymentGroup()
	return &UpdateDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostCluster 编辑主机集群
//
// 编辑主机集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateHostCluster(request *model.UpdateHostClusterRequest) (*model.UpdateHostClusterResponse, error) {
	requestDef := GenReqDefForUpdateHostCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostClusterResponse), nil
	}
}

// UpdateHostClusterInvoker 编辑主机集群
func (c *CodeArtsDeployClient) UpdateHostClusterInvoker(request *model.UpdateHostClusterRequest) *UpdateHostClusterInvoker {
	requestDef := GenReqDefForUpdateHostCluster()
	return &UpdateHostClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckWhetherHostGroupCanBeCreated 判断当前用户在项目下是否有权限创建主机集群
//
// 判断当前用户在项目下是否有权限创建主机集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) CheckWhetherHostGroupCanBeCreated(request *model.CheckWhetherHostGroupCanBeCreatedRequest) (*model.CheckWhetherHostGroupCanBeCreatedResponse, error) {
	requestDef := GenReqDefForCheckWhetherHostGroupCanBeCreated()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWhetherHostGroupCanBeCreatedResponse), nil
	}
}

// CheckWhetherHostGroupCanBeCreatedInvoker 判断当前用户在项目下是否有权限创建主机集群
func (c *CodeArtsDeployClient) CheckWhetherHostGroupCanBeCreatedInvoker(request *model.CheckWhetherHostGroupCanBeCreatedRequest) *CheckWhetherHostGroupCanBeCreatedInvoker {
	requestDef := GenReqDefForCheckWhetherHostGroupCanBeCreated()
	return &CheckWhetherHostGroupCanBeCreatedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroupPermissions 查询主机集群权限矩阵
//
// 根据主机集群id查询主机集群权限矩阵。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListHostGroupPermissions(request *model.ListHostGroupPermissionsRequest) (*model.ListHostGroupPermissionsResponse, error) {
	requestDef := GenReqDefForListHostGroupPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupPermissionsResponse), nil
	}
}

// ListHostGroupPermissionsInvoker 查询主机集群权限矩阵
func (c *CodeArtsDeployClient) ListHostGroupPermissionsInvoker(request *model.ListHostGroupPermissionsRequest) *ListHostGroupPermissionsInvoker {
	requestDef := GenReqDefForListHostGroupPermissions()
	return &ListHostGroupPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostGroupPermissions 修改主机集群权限矩阵
//
// 根据主机集群id修改主机集群权限矩阵。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) UpdateHostGroupPermissions(request *model.UpdateHostGroupPermissionsRequest) (*model.UpdateHostGroupPermissionsResponse, error) {
	requestDef := GenReqDefForUpdateHostGroupPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostGroupPermissionsResponse), nil
	}
}

// UpdateHostGroupPermissionsInvoker 修改主机集群权限矩阵
func (c *CodeArtsDeployClient) UpdateHostGroupPermissionsInvoker(request *model.UpdateHostGroupPermissionsRequest) *UpdateHostGroupPermissionsInvoker {
	requestDef := GenReqDefForUpdateHostGroupPermissions()
	return &UpdateHostGroupPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskSuccessRate 获取指定应用的应用部署成功率
//
// 获取指定应用的应用部署成功率。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ListTaskSuccessRate(request *model.ListTaskSuccessRateRequest) (*model.ListTaskSuccessRateResponse, error) {
	requestDef := GenReqDefForListTaskSuccessRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskSuccessRateResponse), nil
	}
}

// ListTaskSuccessRateInvoker 获取指定应用的应用部署成功率
func (c *CodeArtsDeployClient) ListTaskSuccessRateInvoker(request *model.ListTaskSuccessRateRequest) *ListTaskSuccessRateInvoker {
	requestDef := GenReqDefForListTaskSuccessRate()
	return &ListTaskSuccessRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectSuccessRate 获取指定项目的应用部署成功率
//
// 获取指定项目的应用部署成功率。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsDeployClient) ShowProjectSuccessRate(request *model.ShowProjectSuccessRateRequest) (*model.ShowProjectSuccessRateResponse, error) {
	requestDef := GenReqDefForShowProjectSuccessRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectSuccessRateResponse), nil
	}
}

// ShowProjectSuccessRateInvoker 获取指定项目的应用部署成功率
func (c *CodeArtsDeployClient) ShowProjectSuccessRateInvoker(request *model.ShowProjectSuccessRateRequest) *ShowProjectSuccessRateInvoker {
	requestDef := GenReqDefForShowProjectSuccessRate()
	return &ShowProjectSuccessRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
