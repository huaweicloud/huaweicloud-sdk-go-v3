package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/res/v1/model"
)

type ResClient struct {
	HcClient *http_client.HcHttpClient
}

func NewResClient(hcClient *http_client.HcHttpClient) *ResClient {
	return &ResClient{HcClient: hcClient}
}

func ResClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("basic.Credentials")
	return builder
}

// CreateResDatasource 创建数据源
//
// 在指定的工作空间下面创建一个新的数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResDatasource(request *model.CreateResDatasourceRequest) (*model.CreateResDatasourceResponse, error) {
	requestDef := GenReqDefForCreateResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResDatasourceResponse), nil
	}
}

// CreateResDatasourceInvoker 创建数据源
func (c *ResClient) CreateResDatasourceInvoker(request *model.CreateResDatasourceRequest) *CreateResDatasourceInvoker {
	requestDef := GenReqDefForCreateResDatasource()
	return &CreateResDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResIntelligentScene 创建智能场景
//
// 在指定工作空间下面创建智能场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResIntelligentScene(request *model.CreateResIntelligentSceneRequest) (*model.CreateResIntelligentSceneResponse, error) {
	requestDef := GenReqDefForCreateResIntelligentScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResIntelligentSceneResponse), nil
	}
}

// CreateResIntelligentSceneInvoker 创建智能场景
func (c *ResClient) CreateResIntelligentSceneInvoker(request *model.CreateResIntelligentSceneRequest) *CreateResIntelligentSceneInvoker {
	requestDef := GenReqDefForCreateResIntelligentScene()
	return &CreateResIntelligentSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResJob 新建训练作业
//
// 新建训练作业元数据，新建成功之后可手动执行此任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResJob(request *model.CreateResJobRequest) (*model.CreateResJobResponse, error) {
	requestDef := GenReqDefForCreateResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResJobResponse), nil
	}
}

// CreateResJobInvoker 新建训练作业
func (c *ResClient) CreateResJobInvoker(request *model.CreateResJobRequest) *CreateResJobInvoker {
	requestDef := GenReqDefForCreateResJob()
	return &CreateResJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResJobs 新建多个训练作业
//
// 批量新建作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResJobs(request *model.CreateResJobsRequest) (*model.CreateResJobsResponse, error) {
	requestDef := GenReqDefForCreateResJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResJobsResponse), nil
	}
}

// CreateResJobsInvoker 新建多个训练作业
func (c *ResClient) CreateResJobsInvoker(request *model.CreateResJobsRequest) *CreateResJobsInvoker {
	requestDef := GenReqDefForCreateResJobs()
	return &CreateResJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResOnlineInstance 新建在线服务
//
// 新建在线服务元数据，新建成功之后可手动发布此服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResOnlineInstance(request *model.CreateResOnlineInstanceRequest) (*model.CreateResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForCreateResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResOnlineInstanceResponse), nil
	}
}

// CreateResOnlineInstanceInvoker 新建在线服务
func (c *ResClient) CreateResOnlineInstanceInvoker(request *model.CreateResOnlineInstanceRequest) *CreateResOnlineInstanceInvoker {
	requestDef := GenReqDefForCreateResOnlineInstance()
	return &CreateResOnlineInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResScene 创建自定义场景
//
// 在指定工作空间下面创建自定义场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResScene(request *model.CreateResSceneRequest) (*model.CreateResSceneResponse, error) {
	requestDef := GenReqDefForCreateResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResSceneResponse), nil
	}
}

// CreateResSceneInvoker 创建自定义场景
func (c *ResClient) CreateResSceneInvoker(request *model.CreateResSceneRequest) *CreateResSceneInvoker {
	requestDef := GenReqDefForCreateResScene()
	return &CreateResSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResWorkspace 创建工作空间
//
// 用于在推荐系统下面创建独立的工作空间，用于资源的隔离，用户可以在工作空间下面继续创建数据源、场景以及推荐任务等。是否有工作空间的操作权限取决于用户是否属于当前工作空间绑定的企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) CreateResWorkspace(request *model.CreateResWorkspaceRequest) (*model.CreateResWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResWorkspaceResponse), nil
	}
}

// CreateResWorkspaceInvoker 创建工作空间
func (c *ResClient) CreateResWorkspaceInvoker(request *model.CreateResWorkspaceRequest) *CreateResWorkspaceInvoker {
	requestDef := GenReqDefForCreateResWorkspace()
	return &CreateResWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResDatasource 删除数据源
//
// 删除数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) DeleteResDatasource(request *model.DeleteResDatasourceRequest) (*model.DeleteResDatasourceResponse, error) {
	requestDef := GenReqDefForDeleteResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResDatasourceResponse), nil
	}
}

// DeleteResDatasourceInvoker 删除数据源
func (c *ResClient) DeleteResDatasourceInvoker(request *model.DeleteResDatasourceRequest) *DeleteResDatasourceInvoker {
	requestDef := GenReqDefForDeleteResDatasource()
	return &DeleteResDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResJob 删除训练作业
//
// 删除指定作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) DeleteResJob(request *model.DeleteResJobRequest) (*model.DeleteResJobResponse, error) {
	requestDef := GenReqDefForDeleteResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResJobResponse), nil
	}
}

// DeleteResJobInvoker 删除训练作业
func (c *ResClient) DeleteResJobInvoker(request *model.DeleteResJobRequest) *DeleteResJobInvoker {
	requestDef := GenReqDefForDeleteResJob()
	return &DeleteResJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResOnlineInstance 删除在线服务
//
// 删除在线服务实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) DeleteResOnlineInstance(request *model.DeleteResOnlineInstanceRequest) (*model.DeleteResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForDeleteResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResOnlineInstanceResponse), nil
	}
}

// DeleteResOnlineInstanceInvoker 删除在线服务
func (c *ResClient) DeleteResOnlineInstanceInvoker(request *model.DeleteResOnlineInstanceRequest) *DeleteResOnlineInstanceInvoker {
	requestDef := GenReqDefForDeleteResOnlineInstance()
	return &DeleteResOnlineInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResScene 删除场景
//
// 该接口用于删除场景，删除之后不能恢复，请您谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) DeleteResScene(request *model.DeleteResSceneRequest) (*model.DeleteResSceneResponse, error) {
	requestDef := GenReqDefForDeleteResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResSceneResponse), nil
	}
}

// DeleteResSceneInvoker 删除场景
func (c *ResClient) DeleteResSceneInvoker(request *model.DeleteResSceneRequest) *DeleteResSceneInvoker {
	requestDef := GenReqDefForDeleteResScene()
	return &DeleteResSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResWorkspace 删除工作空间
//
// 删除指定工作空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) DeleteResWorkspace(request *model.DeleteResWorkspaceRequest) (*model.DeleteResWorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResWorkspaceResponse), nil
	}
}

// DeleteResWorkspaceInvoker 删除工作空间
func (c *ResClient) DeleteResWorkspaceInvoker(request *model.DeleteResWorkspaceRequest) *DeleteResWorkspaceInvoker {
	requestDef := GenReqDefForDeleteResWorkspace()
	return &DeleteResWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResDatasources 查询数据源列表
//
// 查询当前工作空间下的数据源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResDatasources(request *model.ListResDatasourcesRequest) (*model.ListResDatasourcesResponse, error) {
	requestDef := GenReqDefForListResDatasources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResDatasourcesResponse), nil
	}
}

// ListResDatasourcesInvoker 查询数据源列表
func (c *ResClient) ListResDatasourcesInvoker(request *model.ListResDatasourcesRequest) *ListResDatasourcesInvoker {
	requestDef := GenReqDefForListResDatasources()
	return &ListResDatasourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResEnterprises 查询企业项目列表
//
// 查询用户在当前项目id下的企业项目列表。在创建工作空间时需要提供企业项目id。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResEnterprises(request *model.ListResEnterprisesRequest) (*model.ListResEnterprisesResponse, error) {
	requestDef := GenReqDefForListResEnterprises()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResEnterprisesResponse), nil
	}
}

// ListResEnterprisesInvoker 查询企业项目列表
func (c *ResClient) ListResEnterprisesInvoker(request *model.ListResEnterprisesRequest) *ListResEnterprisesInvoker {
	requestDef := GenReqDefForListResEnterprises()
	return &ListResEnterprisesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResOnlineServiceDetails 查询在线服务详情
//
// 根据给定的workspace_id和resource_id及category查询在线服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResOnlineServiceDetails(request *model.ListResOnlineServiceDetailsRequest) (*model.ListResOnlineServiceDetailsResponse, error) {
	requestDef := GenReqDefForListResOnlineServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResOnlineServiceDetailsResponse), nil
	}
}

// ListResOnlineServiceDetailsInvoker 查询在线服务详情
func (c *ResClient) ListResOnlineServiceDetailsInvoker(request *model.ListResOnlineServiceDetailsRequest) *ListResOnlineServiceDetailsInvoker {
	requestDef := GenReqDefForListResOnlineServiceDetails()
	return &ListResOnlineServiceDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResResourceSpec 查询训练规格
//
// 查询当前推荐系统所提供的离线计算规格，实时计算规格和排序模型训练规格。在创建数据源和场景时，需要提供此信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResResourceSpec(request *model.ListResResourceSpecRequest) (*model.ListResResourceSpecResponse, error) {
	requestDef := GenReqDefForListResResourceSpec()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResResourceSpecResponse), nil
	}
}

// ListResResourceSpecInvoker 查询训练规格
func (c *ResClient) ListResResourceSpecInvoker(request *model.ListResResourceSpecRequest) *ListResResourceSpecInvoker {
	requestDef := GenReqDefForListResResourceSpec()
	return &ListResResourceSpecInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResScenes 查询场景列表
//
// 查询当前工作空间下的场景列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResScenes(request *model.ListResScenesRequest) (*model.ListResScenesResponse, error) {
	requestDef := GenReqDefForListResScenes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResScenesResponse), nil
	}
}

// ListResScenesInvoker 查询场景列表
func (c *ResClient) ListResScenesInvoker(request *model.ListResScenesRequest) *ListResScenesInvoker {
	requestDef := GenReqDefForListResScenes()
	return &ListResScenesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResWorkspaces 查询工作空间列表
//
// 用于查询当前用户具有操作权限的工作空间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ListResWorkspaces(request *model.ListResWorkspacesRequest) (*model.ListResWorkspacesResponse, error) {
	requestDef := GenReqDefForListResWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResWorkspacesResponse), nil
	}
}

// ListResWorkspacesInvoker 查询工作空间列表
func (c *ResClient) ListResWorkspacesInvoker(request *model.ListResWorkspacesRequest) *ListResWorkspacesInvoker {
	requestDef := GenReqDefForListResWorkspaces()
	return &ListResWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResDatasource 查询数据源详情
//
// 查询指定数据源的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResDatasource(request *model.ShowResDatasourceRequest) (*model.ShowResDatasourceResponse, error) {
	requestDef := GenReqDefForShowResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResDatasourceResponse), nil
	}
}

// ShowResDatasourceInvoker 查询数据源详情
func (c *ResClient) ShowResDatasourceInvoker(request *model.ShowResDatasourceRequest) *ShowResDatasourceInvoker {
	requestDef := GenReqDefForShowResDatasource()
	return &ShowResDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResDatasourceWorkDetail 查询数据源任务结果
//
// 查询指定数据源下离线任务的结果。其中包括数据格式，数据检测、数据探索及效果评估的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResDatasourceWorkDetail(request *model.ShowResDatasourceWorkDetailRequest) (*model.ShowResDatasourceWorkDetailResponse, error) {
	requestDef := GenReqDefForShowResDatasourceWorkDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResDatasourceWorkDetailResponse), nil
	}
}

// ShowResDatasourceWorkDetailInvoker 查询数据源任务结果
func (c *ResClient) ShowResDatasourceWorkDetailInvoker(request *model.ShowResDatasourceWorkDetailRequest) *ShowResDatasourceWorkDetailInvoker {
	requestDef := GenReqDefForShowResDatasourceWorkDetail()
	return &ShowResDatasourceWorkDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResJob 查询训练作业
//
// 查询resource_id（数据源id或场景id）下的指定类型的作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResJob(request *model.ShowResJobRequest) (*model.ShowResJobResponse, error) {
	requestDef := GenReqDefForShowResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResJobResponse), nil
	}
}

// ShowResJobInvoker 查询训练作业
func (c *ResClient) ShowResJobInvoker(request *model.ShowResJobRequest) *ShowResJobInvoker {
	requestDef := GenReqDefForShowResJob()
	return &ShowResJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResRecallSet 查询训练作业候选集
//
// 查询给定workspaces_id和指定resource_id下的候选集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResRecallSet(request *model.ShowResRecallSetRequest) (*model.ShowResRecallSetResponse, error) {
	requestDef := GenReqDefForShowResRecallSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResRecallSetResponse), nil
	}
}

// ShowResRecallSetInvoker 查询训练作业候选集
func (c *ResClient) ShowResRecallSetInvoker(request *model.ShowResRecallSetRequest) *ShowResRecallSetInvoker {
	requestDef := GenReqDefForShowResRecallSet()
	return &ShowResRecallSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResScene 查询场景详情
//
// 查询指定场景的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResScene(request *model.ShowResSceneRequest) (*model.ShowResSceneResponse, error) {
	requestDef := GenReqDefForShowResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResSceneResponse), nil
	}
}

// ShowResSceneInvoker 查询场景详情
func (c *ResClient) ShowResSceneInvoker(request *model.ShowResSceneRequest) *ShowResSceneInvoker {
	requestDef := GenReqDefForShowResScene()
	return &ShowResSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResWrokspace 查询工作空间详情
//
// 查询指定工作空间的具体信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) ShowResWrokspace(request *model.ShowResWrokspaceRequest) (*model.ShowResWrokspaceResponse, error) {
	requestDef := GenReqDefForShowResWrokspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResWrokspaceResponse), nil
	}
}

// ShowResWrokspaceInvoker 查询工作空间详情
func (c *ResClient) ShowResWrokspaceInvoker(request *model.ShowResWrokspaceRequest) *ShowResWrokspaceInvoker {
	requestDef := GenReqDefForShowResWrokspace()
	return &ShowResWrokspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartResJob 执行作业
//
// 执行独立的作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) StartResJob(request *model.StartResJobRequest) (*model.StartResJobResponse, error) {
	requestDef := GenReqDefForStartResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResJobResponse), nil
	}
}

// StartResJobInvoker 执行作业
func (c *ResClient) StartResJobInvoker(request *model.StartResJobRequest) *StartResJobInvoker {
	requestDef := GenReqDefForStartResJob()
	return &StartResJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartResSceneJobs 执行场景
//
// 执行场景下面的所有作业和服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) StartResSceneJobs(request *model.StartResSceneJobsRequest) (*model.StartResSceneJobsResponse, error) {
	requestDef := GenReqDefForStartResSceneJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResSceneJobsResponse), nil
	}
}

// StartResSceneJobsInvoker 执行场景
func (c *ResClient) StartResSceneJobsInvoker(request *model.StartResSceneJobsRequest) *StartResSceneJobsInvoker {
	requestDef := GenReqDefForStartResSceneJobs()
	return &StartResSceneJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResDatasource 修改数据源内容
//
// 修改指定数据源的配置内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResDatasource(request *model.UpdateResDatasourceRequest) (*model.UpdateResDatasourceResponse, error) {
	requestDef := GenReqDefForUpdateResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResDatasourceResponse), nil
	}
}

// UpdateResDatasourceInvoker 修改数据源内容
func (c *ResClient) UpdateResDatasourceInvoker(request *model.UpdateResDatasourceRequest) *UpdateResDatasourceInvoker {
	requestDef := GenReqDefForUpdateResDatasource()
	return &UpdateResDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResDatastruct 修改数据源特征
//
// 修改数据源中的特征。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResDatastruct(request *model.UpdateResDatastructRequest) (*model.UpdateResDatastructResponse, error) {
	requestDef := GenReqDefForUpdateResDatastruct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResDatastructResponse), nil
	}
}

// UpdateResDatastructInvoker 修改数据源特征
func (c *ResClient) UpdateResDatastructInvoker(request *model.UpdateResDatastructRequest) *UpdateResDatastructInvoker {
	requestDef := GenReqDefForUpdateResDatastruct()
	return &UpdateResDatastructInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResIntelligentScene 更新智能场景内容
//
// 更新智能场景的内容信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResIntelligentScene(request *model.UpdateResIntelligentSceneRequest) (*model.UpdateResIntelligentSceneResponse, error) {
	requestDef := GenReqDefForUpdateResIntelligentScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResIntelligentSceneResponse), nil
	}
}

// UpdateResIntelligentSceneInvoker 更新智能场景内容
func (c *ResClient) UpdateResIntelligentSceneInvoker(request *model.UpdateResIntelligentSceneRequest) *UpdateResIntelligentSceneInvoker {
	requestDef := GenReqDefForUpdateResIntelligentScene()
	return &UpdateResIntelligentSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResJob 修改训练作业参数
//
// 修改指定作业的元数据信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResJob(request *model.UpdateResJobRequest) (*model.UpdateResJobResponse, error) {
	requestDef := GenReqDefForUpdateResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResJobResponse), nil
	}
}

// UpdateResJobInvoker 修改训练作业参数
func (c *ResClient) UpdateResJobInvoker(request *model.UpdateResJobRequest) *UpdateResJobInvoker {
	requestDef := GenReqDefForUpdateResJob()
	return &UpdateResJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResOnlineInstance 修改在线服务参数
//
// 修改指定在线服务的元数据内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResOnlineInstance(request *model.UpdateResOnlineInstanceRequest) (*model.UpdateResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForUpdateResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResOnlineInstanceResponse), nil
	}
}

// UpdateResOnlineInstanceInvoker 修改在线服务参数
func (c *ResClient) UpdateResOnlineInstanceInvoker(request *model.UpdateResOnlineInstanceRequest) *UpdateResOnlineInstanceInvoker {
	requestDef := GenReqDefForUpdateResOnlineInstance()
	return &UpdateResOnlineInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResScene 更新自定义场景内容
//
// 更新自定义场景的内容信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResScene(request *model.UpdateResSceneRequest) (*model.UpdateResSceneResponse, error) {
	requestDef := GenReqDefForUpdateResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResSceneResponse), nil
	}
}

// UpdateResSceneInvoker 更新自定义场景内容
func (c *ResClient) UpdateResSceneInvoker(request *model.UpdateResSceneRequest) *UpdateResSceneInvoker {
	requestDef := GenReqDefForUpdateResScene()
	return &UpdateResSceneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResWorkspace 更新工作空间
//
// 更新工作空间信息, 只允许更新描述信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ResClient) UpdateResWorkspace(request *model.UpdateResWorkspaceRequest) (*model.UpdateResWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResWorkspaceResponse), nil
	}
}

// UpdateResWorkspaceInvoker 更新工作空间
func (c *ResClient) UpdateResWorkspaceInvoker(request *model.UpdateResWorkspaceRequest) *UpdateResWorkspaceInvoker {
	requestDef := GenReqDefForUpdateResWorkspace()
	return &UpdateResWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
