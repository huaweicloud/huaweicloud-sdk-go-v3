package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsdeploy/v2/model"
)

type CodeArtsDeployClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeArtsDeployClient(hcClient *http_client.HcHttpClient) *CodeArtsDeployClient {
	return &CodeArtsDeployClient{HcClient: hcClient}
}

func CodeArtsDeployClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
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
// 查询部署记录的执行参数
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
