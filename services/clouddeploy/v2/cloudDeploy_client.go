package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/clouddeploy/v2/model"
)

type CloudDeployClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudDeployClient(hcClient *http_client.HcHttpClient) *CloudDeployClient {
	return &CloudDeployClient{HcClient: hcClient}
}

func CloudDeployClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateDeployTaskByTemplate 通过模板新建部署任务
//
// 通过模板新建部署任务cloudpipeline流水线调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) CreateDeployTaskByTemplate(request *model.CreateDeployTaskByTemplateRequest) (*model.CreateDeployTaskByTemplateResponse, error) {
	requestDef := GenReqDefForCreateDeployTaskByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeployTaskByTemplateResponse), nil
	}
}

// CreateDeployTaskByTemplateInvoker 通过模板新建部署任务
func (c *CloudDeployClient) CreateDeployTaskByTemplateInvoker(request *model.CreateDeployTaskByTemplateRequest) *CreateDeployTaskByTemplateInvoker {
	requestDef := GenReqDefForCreateDeployTaskByTemplate()
	return &CreateDeployTaskByTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployTask 删除部署任务
//
// 根据部署任务id删除部署任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) DeleteDeployTask(request *model.DeleteDeployTaskRequest) (*model.DeleteDeployTaskResponse, error) {
	requestDef := GenReqDefForDeleteDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployTaskResponse), nil
	}
}

// DeleteDeployTaskInvoker 删除部署任务
func (c *CloudDeployClient) DeleteDeployTaskInvoker(request *model.DeleteDeployTaskRequest) *DeleteDeployTaskInvoker {
	requestDef := GenReqDefForDeleteDeployTask()
	return &DeleteDeployTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployTaskHistoryByDate 根据开始时间和结束时间查询项目下指定任务的历史执行记录列表
//
// 根据开始时间和结束时间查询项目下指定任务的历史执行记录列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ListDeployTaskHistoryByDate(request *model.ListDeployTaskHistoryByDateRequest) (*model.ListDeployTaskHistoryByDateResponse, error) {
	requestDef := GenReqDefForListDeployTaskHistoryByDate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeployTaskHistoryByDateResponse), nil
	}
}

// ListDeployTaskHistoryByDateInvoker 根据开始时间和结束时间查询项目下指定任务的历史执行记录列表
func (c *CloudDeployClient) ListDeployTaskHistoryByDateInvoker(request *model.ListDeployTaskHistoryByDateRequest) *ListDeployTaskHistoryByDateInvoker {
	requestDef := GenReqDefForListDeployTaskHistoryByDate()
	return &ListDeployTaskHistoryByDateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployTasks 获取部署任务列表
//
// 查询项目下部署任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ListDeployTasks(request *model.ListDeployTasksRequest) (*model.ListDeployTasksResponse, error) {
	requestDef := GenReqDefForListDeployTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeployTasksResponse), nil
	}
}

// ListDeployTasksInvoker 获取部署任务列表
func (c *CloudDeployClient) ListDeployTasksInvoker(request *model.ListDeployTasksRequest) *ListDeployTasksInvoker {
	requestDef := GenReqDefForListDeployTasks()
	return &ListDeployTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeployTaskDetail 获取部署任务详情
//
// 根据部署任务id获取部署任务详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ShowDeployTaskDetail(request *model.ShowDeployTaskDetailRequest) (*model.ShowDeployTaskDetailResponse, error) {
	requestDef := GenReqDefForShowDeployTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeployTaskDetailResponse), nil
	}
}

// ShowDeployTaskDetailInvoker 获取部署任务详情
func (c *CloudDeployClient) ShowDeployTaskDetailInvoker(request *model.ShowDeployTaskDetailRequest) *ShowDeployTaskDetailInvoker {
	requestDef := GenReqDefForShowDeployTaskDetail()
	return &ShowDeployTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDeployTask 启动部署任务
//
// 根据部署任务id启动部署任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) StartDeployTask(request *model.StartDeployTaskRequest) (*model.StartDeployTaskResponse, error) {
	requestDef := GenReqDefForStartDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDeployTaskResponse), nil
	}
}

// StartDeployTaskInvoker 启动部署任务
func (c *CloudDeployClient) StartDeployTaskInvoker(request *model.StartDeployTaskRequest) *StartDeployTaskInvoker {
	requestDef := GenReqDefForStartDeployTask()
	return &StartDeployTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeploymentHost 新建主机
//
// 在指定主机组下新建主机。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) CreateDeploymentHost(request *model.CreateDeploymentHostRequest) (*model.CreateDeploymentHostResponse, error) {
	requestDef := GenReqDefForCreateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentHostResponse), nil
	}
}

// CreateDeploymentHostInvoker 新建主机
func (c *CloudDeployClient) CreateDeploymentHostInvoker(request *model.CreateDeploymentHostRequest) *CreateDeploymentHostInvoker {
	requestDef := GenReqDefForCreateDeploymentHost()
	return &CreateDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeploymentHost 删除主机
//
// 根据主机id删除主机。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) DeleteDeploymentHost(request *model.DeleteDeploymentHostRequest) (*model.DeleteDeploymentHostResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentHostResponse), nil
	}
}

// DeleteDeploymentHostInvoker 删除主机
func (c *CloudDeployClient) DeleteDeploymentHostInvoker(request *model.DeleteDeploymentHostRequest) *DeleteDeploymentHostInvoker {
	requestDef := GenReqDefForDeleteDeploymentHost()
	return &DeleteDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHosts 查询主机列表
//
// 根据主机组id查询指定主机组下的主机列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

// ListHostsInvoker 查询主机列表
func (c *CloudDeployClient) ListHostsInvoker(request *model.ListHostsRequest) *ListHostsInvoker {
	requestDef := GenReqDefForListHosts()
	return &ListHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentHostDetail 查询主机详情
//
// 根据主机id查询主机详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ShowDeploymentHostDetail(request *model.ShowDeploymentHostDetailRequest) (*model.ShowDeploymentHostDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentHostDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentHostDetailResponse), nil
	}
}

// ShowDeploymentHostDetailInvoker 查询主机详情
func (c *CloudDeployClient) ShowDeploymentHostDetailInvoker(request *model.ShowDeploymentHostDetailRequest) *ShowDeploymentHostDetailInvoker {
	requestDef := GenReqDefForShowDeploymentHostDetail()
	return &ShowDeploymentHostDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeploymentHost 修改主机
//
// 根据主机id修改主机信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) UpdateDeploymentHost(request *model.UpdateDeploymentHostRequest) (*model.UpdateDeploymentHostResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentHostResponse), nil
	}
}

// UpdateDeploymentHostInvoker 修改主机
func (c *CloudDeployClient) UpdateDeploymentHostInvoker(request *model.UpdateDeploymentHostRequest) *UpdateDeploymentHostInvoker {
	requestDef := GenReqDefForUpdateDeploymentHost()
	return &UpdateDeploymentHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeploymentGroup 新建主机组
//
// 在项目下新建主机组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) CreateDeploymentGroup(request *model.CreateDeploymentGroupRequest) (*model.CreateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForCreateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentGroupResponse), nil
	}
}

// CreateDeploymentGroupInvoker 新建主机组
func (c *CloudDeployClient) CreateDeploymentGroupInvoker(request *model.CreateDeploymentGroupRequest) *CreateDeploymentGroupInvoker {
	requestDef := GenReqDefForCreateDeploymentGroup()
	return &CreateDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeploymentGroup 删除主机组
//
// 根据主机组id删除主机组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) DeleteDeploymentGroup(request *model.DeleteDeploymentGroupRequest) (*model.DeleteDeploymentGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentGroupResponse), nil
	}
}

// DeleteDeploymentGroupInvoker 删除主机组
func (c *CloudDeployClient) DeleteDeploymentGroupInvoker(request *model.DeleteDeploymentGroupRequest) *DeleteDeploymentGroupInvoker {
	requestDef := GenReqDefForDeleteDeploymentGroup()
	return &DeleteDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroups 查询主机组列表
//
// 按条件查询主机组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ListHostGroups(request *model.ListHostGroupsRequest) (*model.ListHostGroupsResponse, error) {
	requestDef := GenReqDefForListHostGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupsResponse), nil
	}
}

// ListHostGroupsInvoker 查询主机组列表
func (c *CloudDeployClient) ListHostGroupsInvoker(request *model.ListHostGroupsRequest) *ListHostGroupsInvoker {
	requestDef := GenReqDefForListHostGroups()
	return &ListHostGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentGroupDetail 查询主机组
//
// 根据主机组id查询主机组详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ShowDeploymentGroupDetail(request *model.ShowDeploymentGroupDetailRequest) (*model.ShowDeploymentGroupDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentGroupDetailResponse), nil
	}
}

// ShowDeploymentGroupDetailInvoker 查询主机组
func (c *CloudDeployClient) ShowDeploymentGroupDetailInvoker(request *model.ShowDeploymentGroupDetailRequest) *ShowDeploymentGroupDetailInvoker {
	requestDef := GenReqDefForShowDeploymentGroupDetail()
	return &ShowDeploymentGroupDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeploymentGroup 修改主机组
//
// 根据主机组id修改主机组信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) UpdateDeploymentGroup(request *model.UpdateDeploymentGroupRequest) (*model.UpdateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentGroupResponse), nil
	}
}

// UpdateDeploymentGroupInvoker 修改主机组
func (c *CloudDeployClient) UpdateDeploymentGroupInvoker(request *model.UpdateDeploymentGroupRequest) *UpdateDeploymentGroupInvoker {
	requestDef := GenReqDefForUpdateDeploymentGroup()
	return &UpdateDeploymentGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskSuccessRate 获取指定任务的部署任务执行成功率
//
// 获取指定任务的部署任务执行成功率
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ListTaskSuccessRate(request *model.ListTaskSuccessRateRequest) (*model.ListTaskSuccessRateResponse, error) {
	requestDef := GenReqDefForListTaskSuccessRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskSuccessRateResponse), nil
	}
}

// ListTaskSuccessRateInvoker 获取指定任务的部署任务执行成功率
func (c *CloudDeployClient) ListTaskSuccessRateInvoker(request *model.ListTaskSuccessRateRequest) *ListTaskSuccessRateInvoker {
	requestDef := GenReqDefForListTaskSuccessRate()
	return &ListTaskSuccessRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectSuccessRate 获取指定项目的部署任务执行成功率
//
// 获取指定项目的部署任务执行成功率
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudDeployClient) ShowProjectSuccessRate(request *model.ShowProjectSuccessRateRequest) (*model.ShowProjectSuccessRateResponse, error) {
	requestDef := GenReqDefForShowProjectSuccessRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectSuccessRateResponse), nil
	}
}

// ShowProjectSuccessRateInvoker 获取指定项目的部署任务执行成功率
func (c *CloudDeployClient) ShowProjectSuccessRateInvoker(request *model.ShowProjectSuccessRateRequest) *ShowProjectSuccessRateInvoker {
	requestDef := GenReqDefForShowProjectSuccessRate()
	return &ShowProjectSuccessRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
