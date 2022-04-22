package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/devstar/v1/model"
)

type DevStarClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDevStarClient(hcClient *http_client.HcHttpClient) *DevStarClient {
	return &DevStarClient{HcClient: hcClient}
}

func DevStarClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// 通过应用Id获取软件发布仓库列表
//
// 通过应用Id获取软件发布仓库列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationReleaseRepositories(request *model.ShowApplicationReleaseRepositoriesRequest) (*model.ShowApplicationReleaseRepositoriesResponse, error) {
	requestDef := GenReqDefForShowApplicationReleaseRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationReleaseRepositoriesResponse), nil
	}
}

// 查询应用关联资源删除状态
//
// 根据应用Id查询应用关联的代码仓、流水线删除状态 使用场景：用户删除应用关联的资源（如代码仓、流水线...）后，通过该接口实时查询代码仓、流水线删除状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationResDeleteStatus(request *model.ShowApplicationResDeleteStatusRequest) (*model.ShowApplicationResDeleteStatusResponse, error) {
	requestDef := GenReqDefForShowApplicationResDeleteStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationResDeleteStatusResponse), nil
	}
}

// 获取应用依赖元数据资源
//
// 根据应用Id获取依赖元数据资源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationDependentResources(request *model.ShowApplicationDependentResourcesRequest) (*model.ShowApplicationDependentResourcesResponse, error) {
	requestDef := GenReqDefForShowApplicationDependentResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationDependentResourcesResponse), nil
	}
}

// 获取应用详情
//
// 根据应用Id获取应用详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationV3(request *model.ShowApplicationV3Request) (*model.ShowApplicationV3Response, error) {
	requestDef := GenReqDefForShowApplicationV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationV3Response), nil
	}
}

// 更新应用信息
//
// 根据应用Id更新对应有权限的应用信息
// - 允许更新信息的信息包含
// name,description,icon
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

// 删除应用信息
//
// 根据应用Id删除应用，并可以选择删除其关联的代码仓、流水线资源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) DeleteApplicationV4(request *model.DeleteApplicationV4Request) (*model.DeleteApplicationV4Response, error) {
	requestDef := GenReqDefForDeleteApplicationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationV4Response), nil
	}
}

// 获取应用列表
//
// 获取我创建的应用列表
// 当前只支持查询我创建的应用，其中请求参数is_created_by_self需为true
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListApplicationsV6(request *model.ListApplicationsV6Request) (*model.ListApplicationsV6Response, error) {
	requestDef := GenReqDefForListApplicationsV6()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsV6Response), nil
	}
}

// 下载模板产物
//
// 下载模板产物。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) DownloadApplicationCode(request *model.DownloadApplicationCodeRequest) (*model.DownloadApplicationCodeResponse, error) {
	requestDef := GenReqDefForDownloadApplicationCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadApplicationCodeResponse), nil
	}
}

// 创建部署任务
//
// 创建部署任务，并触发任务执行，当前只支持函数部署。
// 其中，报文中file_id为查询软件版本包接口返回版本包id;
// handler为在函数部署方式下，入口函数名称，从应用代码中获取，格式为“包名.类名.函数名称”，例如：com.example.demo.APIGTrigger.handler。
// 也可以从应用详情接口返回结构template_deployment中直接获取。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) CreateDeploymentJobs(request *model.CreateDeploymentJobsRequest) (*model.CreateDeploymentJobsResponse, error) {
	requestDef := GenReqDefForCreateDeploymentJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentJobsResponse), nil
	}
}

// 查询应用环境部署任务详情
//
// 查询应用环境部署任务详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowDeploymentJobs(request *model.ShowDeploymentJobsRequest) (*model.ShowDeploymentJobsResponse, error) {
	requestDef := GenReqDefForShowDeploymentJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentJobsResponse), nil
	}
}

// CodeHub 模板生成代码
//
// 使用CodeHub模板创建应用代码。
//
// 通过 Codehub 模板创建生成应用代码的任务，并将应用代码存储于指定的 CodeHub 仓库中或者生成代码压缩包，可以通过返回的任务 ID 查询相关任务状态。
//
// - 接口鉴权方式
// 通过华为云服务获取的用户token。
//
// - 代码生成位置
// 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) RunCodehubTemplateJob(request *model.RunCodehubTemplateJobRequest) (*model.RunCodehubTemplateJobResponse, error) {
	requestDef := GenReqDefForRunCodehubTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCodehubTemplateJobResponse), nil
	}
}

// Devstar 模板生成代码
//
// 使用DevStar的模板创建应用代码。
//
// 通过 DevStar 模板创建生成应用代码的任务，并将应用代码存储于指定的 CodeHub 仓库中，可以通过返回的任务 ID 查询相关任务状态。
//
// - 接口鉴权方式
// 通过华为云服务获取的用户token。
//
// - 代码生成位置
// 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) RunDevstarTemplateJob(request *model.RunDevstarTemplateJobRequest) (*model.RunDevstarTemplateJobResponse, error) {
	requestDef := GenReqDefForRunDevstarTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDevstarTemplateJobResponse), nil
	}
}

// 查询任务详情
//
// 查询任务的详情。
//
// 通过任务ID可以查看任务的状态
// 当任务结束时返回应用代码存放的位置。
//
// - 接口鉴权方式
// 通过华为云服务获取的用户token。
//
// - 代码生成位置
// 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// 流水线模板列表查询
//
// 流水线模板列表查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListPipelineTemplates(request *model.ListPipelineTemplatesRequest) (*model.ListPipelineTemplatesResponse, error) {
	requestDef := GenReqDefForListPipelineTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineTemplatesResponse), nil
	}
}

// 查询流水线最近一次运行状态查询接口
//
// 查询应用流水线最近一次运行状态查询接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowPipelineLastStatusV2(request *model.ShowPipelineLastStatusV2Request) (*model.ShowPipelineLastStatusV2Response, error) {
	requestDef := GenReqDefForShowPipelineLastStatusV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineLastStatusV2Response), nil
	}
}

// 根据流水线Id操作流水线启动
//
// 根据流水线Id操作流水线启动
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) StartPipeline(request *model.StartPipelineRequest) (*model.StartPipelineResponse, error) {
	requestDef := GenReqDefForStartPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineResponse), nil
	}
}

// 获取用户有权限的DevStar存量DevCloud项目列表
//
// 获取用户有权限的DevStar存量DevCloud项目列表。
// 来源包括：1.DevStar创建的DevCloud项目；2.DevStar应用有关联DevCloud项目。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListProjectsV4(request *model.ListProjectsV4Request) (*model.ListProjectsV4Response, error) {
	requestDef := GenReqDefForListProjectsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectsV4Response), nil
	}
}

// 检查仓库名称是否重名
//
// 检查仓库名称是否重名
// - 校验规则
//     同一个项目下的仓库名称不能存在重复,当结果为true时,校验通过,仓库名称可用,否则,校验不通过,当前项目下的仓库名称已存在,不可用
// - 必传参数
//     project_id,name,region_id
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) CheckRepositoryDuplicateName(request *model.CheckRepositoryDuplicateNameRequest) (*model.CheckRepositoryDuplicateNameResponse, error) {
	requestDef := GenReqDefForCheckRepositoryDuplicateName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRepositoryDuplicateNameResponse), nil
	}
}

// 使用 CloudIDE 实例打开应用代码
//
// 使用 CloudIDE 实例打开应用代码。CloudIDE会保存用户项目数据，相同用户使用同一个CloudIDE，使用要求：
// - 用户需为登录状态。
// - 拥有仓库权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowRepositoryByCloudIde(request *model.ShowRepositoryByCloudIdeRequest) (*model.ShowRepositoryByCloudIdeResponse, error) {
	requestDef := GenReqDefForShowRepositoryByCloudIde()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryByCloudIdeResponse), nil
	}
}

// 应用代码仓库统计信息
//
// 查询代码仓库的统计信息,包括代码仓的名称,代码行数等信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowRepositoryStatisticalDataV2(request *model.ShowRepositoryStatisticalDataV2Request) (*model.ShowRepositoryStatisticalDataV2Response, error) {
	requestDef := GenReqDefForShowRepositoryStatisticalDataV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticalDataV2Response), nil
	}
}

// 读取模板文件
//
// 该接口可以用于模板作者或模板维护人读取模板文件内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateFile(request *model.ShowTemplateFileRequest) (*model.ShowTemplateFileResponse, error) {
	requestDef := GenReqDefForShowTemplateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateFileResponse), nil
	}
}

// 同步模板浏览记录
//
// 未登录状态下，将用户浏览过的模板缓存在浏览器中，登录时，调用该接口同步模板浏览记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) CreateTemplateViewHistories(request *model.CreateTemplateViewHistoriesRequest) (*model.CreateTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForCreateTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateViewHistoriesResponse), nil
	}
}

// 查询模板列表（V1）
//
// 查询模板列表，推荐使用/v1/templates/query接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListPublishedTemplates(request *model.ListPublishedTemplatesRequest) (*model.ListPublishedTemplatesResponse, error) {
	requestDef := GenReqDefForListPublishedTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublishedTemplatesResponse), nil
	}
}

// 我浏览的模板记录
//
// 查询DevStar或者CodeLabs登录用户浏览过的模板（只返回最近浏览的5个模板）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListTemplateViewHistories(request *model.ListTemplateViewHistoriesRequest) (*model.ListTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForListTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateViewHistoriesResponse), nil
	}
}

// 查询模板列表
//
// 查询模板列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// 查询模板列表（V2）
//
// 查询模板列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ListTemplatesV2(request *model.ListTemplatesV2Request) (*model.ListTemplatesV2Response, error) {
	requestDef := GenReqDefForListTemplatesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesV2Response), nil
	}
}

// 查询模板详情（V3）
//
// 获取指定模板详情，包括模板id、名称、描述、作者、标签、上架时间等信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateV3(request *model.ShowTemplateV3Request) (*model.ShowTemplateV3Response, error) {
	requestDef := GenReqDefForShowTemplateV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateV3Response), nil
	}
}

// 查询模板详情（V1）
//
// 查询模板详情，推荐使用V3版本接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}
