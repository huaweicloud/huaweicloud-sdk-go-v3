package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 创建数据源
//
// 在指定的工作空间下面创建一个新的数据源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResDatasource(request *model.CreateResDatasourceRequest) (*model.CreateResDatasourceResponse, error) {
	requestDef := GenReqDefForCreateResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResDatasourceResponse), nil
	}
}

// 创建智能场景
//
// 在指定工作空间下面创建智能场景。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResIntelligentScene(request *model.CreateResIntelligentSceneRequest) (*model.CreateResIntelligentSceneResponse, error) {
	requestDef := GenReqDefForCreateResIntelligentScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResIntelligentSceneResponse), nil
	}
}

// 新建训练作业
//
// 新建训练作业元数据，新建成功之后可手动执行此任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResJob(request *model.CreateResJobRequest) (*model.CreateResJobResponse, error) {
	requestDef := GenReqDefForCreateResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResJobResponse), nil
	}
}

// 新建多个训练作业
//
// 批量新建作业。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResJobs(request *model.CreateResJobsRequest) (*model.CreateResJobsResponse, error) {
	requestDef := GenReqDefForCreateResJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResJobsResponse), nil
	}
}

// 新建在线服务
//
// 新建在线服务元数据，新建成功之后可手动发布此服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResOnlineInstance(request *model.CreateResOnlineInstanceRequest) (*model.CreateResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForCreateResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResOnlineInstanceResponse), nil
	}
}

// 创建自定义场景
//
// 在指定工作空间下面创建自定义场景。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResScene(request *model.CreateResSceneRequest) (*model.CreateResSceneResponse, error) {
	requestDef := GenReqDefForCreateResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResSceneResponse), nil
	}
}

// 创建工作空间
//
// 用于在推荐系统下面创建独立的工作空间，用于资源的隔离，用户可以在工作空间下面继续创建数据源、场景以及推荐任务等。是否有工作空间的操作权限取决于用户是否属于当前工作空间绑定的企业项目。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) CreateResWorkspace(request *model.CreateResWorkspaceRequest) (*model.CreateResWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResWorkspaceResponse), nil
	}
}

// 删除数据源
//
// 删除数据源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) DeleteResDatasource(request *model.DeleteResDatasourceRequest) (*model.DeleteResDatasourceResponse, error) {
	requestDef := GenReqDefForDeleteResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResDatasourceResponse), nil
	}
}

// 删除训练作业
//
// 删除指定作业。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) DeleteResJob(request *model.DeleteResJobRequest) (*model.DeleteResJobResponse, error) {
	requestDef := GenReqDefForDeleteResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResJobResponse), nil
	}
}

// 删除在线服务
//
// 删除在线服务实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) DeleteResOnlineInstance(request *model.DeleteResOnlineInstanceRequest) (*model.DeleteResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForDeleteResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResOnlineInstanceResponse), nil
	}
}

// 删除场景
//
// 该接口用于删除场景，删除之后不能恢复，请您谨慎操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) DeleteResScene(request *model.DeleteResSceneRequest) (*model.DeleteResSceneResponse, error) {
	requestDef := GenReqDefForDeleteResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResSceneResponse), nil
	}
}

// 删除工作空间
//
// 删除指定工作空间。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) DeleteResWorkspace(request *model.DeleteResWorkspaceRequest) (*model.DeleteResWorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResWorkspaceResponse), nil
	}
}

// 查询数据源列表
//
// 查询当前工作空间下的数据源列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResDatasources(request *model.ListResDatasourcesRequest) (*model.ListResDatasourcesResponse, error) {
	requestDef := GenReqDefForListResDatasources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResDatasourcesResponse), nil
	}
}

// 查询企业项目列表
//
// 查询用户在当前项目id下的企业项目列表。在创建工作空间时需要提供企业项目id。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResEnterprises(request *model.ListResEnterprisesRequest) (*model.ListResEnterprisesResponse, error) {
	requestDef := GenReqDefForListResEnterprises()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResEnterprisesResponse), nil
	}
}

// 查询在线服务详情
//
// 根据给定的workspace_id和resource_id及category查询在线服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResOnlineServiceDetails(request *model.ListResOnlineServiceDetailsRequest) (*model.ListResOnlineServiceDetailsResponse, error) {
	requestDef := GenReqDefForListResOnlineServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResOnlineServiceDetailsResponse), nil
	}
}

// 查询训练规格
//
// 查询当前推荐系统所提供的离线计算规格，实时计算规格和排序模型训练规格。在创建数据源和场景时，需要提供此信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResResourceSpec(request *model.ListResResourceSpecRequest) (*model.ListResResourceSpecResponse, error) {
	requestDef := GenReqDefForListResResourceSpec()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResResourceSpecResponse), nil
	}
}

// 查询场景列表
//
// 查询当前工作空间下的场景列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResScenes(request *model.ListResScenesRequest) (*model.ListResScenesResponse, error) {
	requestDef := GenReqDefForListResScenes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResScenesResponse), nil
	}
}

// 查询工作空间列表
//
// 用于查询当前用户具有操作权限的工作空间列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ListResWorkspaces(request *model.ListResWorkspacesRequest) (*model.ListResWorkspacesResponse, error) {
	requestDef := GenReqDefForListResWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResWorkspacesResponse), nil
	}
}

// 查询数据源详情
//
// 查询指定数据源的详情信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResDatasource(request *model.ShowResDatasourceRequest) (*model.ShowResDatasourceResponse, error) {
	requestDef := GenReqDefForShowResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResDatasourceResponse), nil
	}
}

// 查询数据源任务结果
//
// 查询指定数据源下离线任务的结果。其中包括数据格式，数据检测、数据探索及效果评估的内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResDatasourceWorkDetail(request *model.ShowResDatasourceWorkDetailRequest) (*model.ShowResDatasourceWorkDetailResponse, error) {
	requestDef := GenReqDefForShowResDatasourceWorkDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResDatasourceWorkDetailResponse), nil
	}
}

// 查询训练作业
//
// 查询resource_id（数据源id或场景id）下的指定类型的作业。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResJob(request *model.ShowResJobRequest) (*model.ShowResJobResponse, error) {
	requestDef := GenReqDefForShowResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResJobResponse), nil
	}
}

// 查询训练作业候选集
//
// 查询给定workspaces_id和指定resource_id下的候选集。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResRecallSet(request *model.ShowResRecallSetRequest) (*model.ShowResRecallSetResponse, error) {
	requestDef := GenReqDefForShowResRecallSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResRecallSetResponse), nil
	}
}

// 查询场景详情
//
// 查询指定场景的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResScene(request *model.ShowResSceneRequest) (*model.ShowResSceneResponse, error) {
	requestDef := GenReqDefForShowResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResSceneResponse), nil
	}
}

// 查询工作空间详情
//
// 查询指定工作空间的具体信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) ShowResWrokspace(request *model.ShowResWrokspaceRequest) (*model.ShowResWrokspaceResponse, error) {
	requestDef := GenReqDefForShowResWrokspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResWrokspaceResponse), nil
	}
}

// 执行作业
//
// 执行独立的作业。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) StartResJob(request *model.StartResJobRequest) (*model.StartResJobResponse, error) {
	requestDef := GenReqDefForStartResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResJobResponse), nil
	}
}

// 执行场景
//
// 执行场景下面的所有作业和服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) StartResSceneJobs(request *model.StartResSceneJobsRequest) (*model.StartResSceneJobsResponse, error) {
	requestDef := GenReqDefForStartResSceneJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResSceneJobsResponse), nil
	}
}

// 修改数据源内容
//
// 修改指定数据源的配置内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResDatasource(request *model.UpdateResDatasourceRequest) (*model.UpdateResDatasourceResponse, error) {
	requestDef := GenReqDefForUpdateResDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResDatasourceResponse), nil
	}
}

// 修改数据源特征
//
// 修改数据源中的特征。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResDatastruct(request *model.UpdateResDatastructRequest) (*model.UpdateResDatastructResponse, error) {
	requestDef := GenReqDefForUpdateResDatastruct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResDatastructResponse), nil
	}
}

// 更新智能场景内容
//
// 更新智能场景的内容信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResIntelligentScene(request *model.UpdateResIntelligentSceneRequest) (*model.UpdateResIntelligentSceneResponse, error) {
	requestDef := GenReqDefForUpdateResIntelligentScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResIntelligentSceneResponse), nil
	}
}

// 修改训练作业参数
//
// 修改指定作业的元数据信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResJob(request *model.UpdateResJobRequest) (*model.UpdateResJobResponse, error) {
	requestDef := GenReqDefForUpdateResJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResJobResponse), nil
	}
}

// 修改在线服务参数
//
// 修改指定在线服务的元数据内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResOnlineInstance(request *model.UpdateResOnlineInstanceRequest) (*model.UpdateResOnlineInstanceResponse, error) {
	requestDef := GenReqDefForUpdateResOnlineInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResOnlineInstanceResponse), nil
	}
}

// 更新自定义场景内容
//
// 更新自定义场景的内容信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResScene(request *model.UpdateResSceneRequest) (*model.UpdateResSceneResponse, error) {
	requestDef := GenReqDefForUpdateResScene()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResSceneResponse), nil
	}
}

// 更新工作空间
//
// 更新工作空间信息, 只允许更新描述信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ResClient) UpdateResWorkspace(request *model.UpdateResWorkspaceRequest) (*model.UpdateResWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateResWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResWorkspaceResponse), nil
	}
}
