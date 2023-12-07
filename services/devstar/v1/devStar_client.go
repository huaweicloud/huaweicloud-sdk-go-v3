package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// ShowApplicationReleaseRepositories 通过应用Id获取软件发布仓库列表
//
// 通过应用Id获取软件发布仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationReleaseRepositories(request *model.ShowApplicationReleaseRepositoriesRequest) (*model.ShowApplicationReleaseRepositoriesResponse, error) {
	requestDef := GenReqDefForShowApplicationReleaseRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationReleaseRepositoriesResponse), nil
	}
}

// ShowApplicationReleaseRepositoriesInvoker 通过应用Id获取软件发布仓库列表
func (c *DevStarClient) ShowApplicationReleaseRepositoriesInvoker(request *model.ShowApplicationReleaseRepositoriesRequest) *ShowApplicationReleaseRepositoriesInvoker {
	requestDef := GenReqDefForShowApplicationReleaseRepositories()
	return &ShowApplicationReleaseRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicationResDeleteStatus 查询应用关联资源删除状态
//
// 根据应用Id查询应用关联的代码仓、流水线删除状态 使用场景：用户删除应用关联的资源（如代码仓、流水线...）后，通过该接口实时查询代码仓、流水线删除状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationResDeleteStatus(request *model.ShowApplicationResDeleteStatusRequest) (*model.ShowApplicationResDeleteStatusResponse, error) {
	requestDef := GenReqDefForShowApplicationResDeleteStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationResDeleteStatusResponse), nil
	}
}

// ShowApplicationResDeleteStatusInvoker 查询应用关联资源删除状态
func (c *DevStarClient) ShowApplicationResDeleteStatusInvoker(request *model.ShowApplicationResDeleteStatusRequest) *ShowApplicationResDeleteStatusInvoker {
	requestDef := GenReqDefForShowApplicationResDeleteStatus()
	return &ShowApplicationResDeleteStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicationDependentResources 获取应用依赖元数据资源
//
// 根据应用Id获取依赖元数据资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationDependentResources(request *model.ShowApplicationDependentResourcesRequest) (*model.ShowApplicationDependentResourcesResponse, error) {
	requestDef := GenReqDefForShowApplicationDependentResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationDependentResourcesResponse), nil
	}
}

// ShowApplicationDependentResourcesInvoker 获取应用依赖元数据资源
func (c *DevStarClient) ShowApplicationDependentResourcesInvoker(request *model.ShowApplicationDependentResourcesRequest) *ShowApplicationDependentResourcesInvoker {
	requestDef := GenReqDefForShowApplicationDependentResources()
	return &ShowApplicationDependentResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicationV3 获取应用详情
//
// 根据应用Id获取应用详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowApplicationV3(request *model.ShowApplicationV3Request) (*model.ShowApplicationV3Response, error) {
	requestDef := GenReqDefForShowApplicationV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationV3Response), nil
	}
}

// ShowApplicationV3Invoker 获取应用详情
func (c *DevStarClient) ShowApplicationV3Invoker(request *model.ShowApplicationV3Request) *ShowApplicationV3Invoker {
	requestDef := GenReqDefForShowApplicationV3()
	return &ShowApplicationV3Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplication 更新应用信息
//
// 根据应用Id更新对应有权限的应用信息
// - 允许更新信息的信息包含
// name,description,icon
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

// UpdateApplicationInvoker 更新应用信息
func (c *DevStarClient) UpdateApplicationInvoker(request *model.UpdateApplicationRequest) *UpdateApplicationInvoker {
	requestDef := GenReqDefForUpdateApplication()
	return &UpdateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationV4 删除应用信息
//
// 根据应用Id删除应用，并可以选择删除其关联的代码仓、流水线资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) DeleteApplicationV4(request *model.DeleteApplicationV4Request) (*model.DeleteApplicationV4Response, error) {
	requestDef := GenReqDefForDeleteApplicationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationV4Response), nil
	}
}

// DeleteApplicationV4Invoker 删除应用信息
func (c *DevStarClient) DeleteApplicationV4Invoker(request *model.DeleteApplicationV4Request) *DeleteApplicationV4Invoker {
	requestDef := GenReqDefForDeleteApplicationV4()
	return &DeleteApplicationV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationsV6 获取应用列表
//
// 获取我创建的应用列表
// 当前只支持查询我创建的应用，其中请求参数is_created_by_self需为true
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListApplicationsV6(request *model.ListApplicationsV6Request) (*model.ListApplicationsV6Response, error) {
	requestDef := GenReqDefForListApplicationsV6()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsV6Response), nil
	}
}

// ListApplicationsV6Invoker 获取应用列表
func (c *DevStarClient) ListApplicationsV6Invoker(request *model.ListApplicationsV6Request) *ListApplicationsV6Invoker {
	requestDef := GenReqDefForListApplicationsV6()
	return &ListApplicationsV6Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadApplicationCode 下载模板产物
//
// 下载模板产物。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) DownloadApplicationCode(request *model.DownloadApplicationCodeRequest) (*model.DownloadApplicationCodeResponse, error) {
	requestDef := GenReqDefForDownloadApplicationCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadApplicationCodeResponse), nil
	}
}

// DownloadApplicationCodeInvoker 下载模板产物
func (c *DevStarClient) DownloadApplicationCodeInvoker(request *model.DownloadApplicationCodeRequest) *DownloadApplicationCodeInvoker {
	requestDef := GenReqDefForDownloadApplicationCode()
	return &DownloadApplicationCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmDeploymentJob 部署任务执行变更人工审核
//
// 部署任务执行变更人工审核，终止或者继续部署任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ConfirmDeploymentJob(request *model.ConfirmDeploymentJobRequest) (*model.ConfirmDeploymentJobResponse, error) {
	requestDef := GenReqDefForConfirmDeploymentJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmDeploymentJobResponse), nil
	}
}

// ConfirmDeploymentJobInvoker 部署任务执行变更人工审核
func (c *DevStarClient) ConfirmDeploymentJobInvoker(request *model.ConfirmDeploymentJobRequest) *ConfirmDeploymentJobInvoker {
	requestDef := GenReqDefForConfirmDeploymentJob()
	return &ConfirmDeploymentJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeploymentJobs 创建部署任务
//
// 创建部署任务，并触发任务执行，当前只支持函数部署。
// 其中，报文中file_id为查询软件版本包接口返回版本包id;
// handler为在函数部署方式下，入口函数名称，从应用代码中获取，格式为“包名.类名.函数名称”，例如：com.example.demo.APIGTrigger.handler。
// 也可以从应用详情接口返回结构template_deployment中直接获取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) CreateDeploymentJobs(request *model.CreateDeploymentJobsRequest) (*model.CreateDeploymentJobsResponse, error) {
	requestDef := GenReqDefForCreateDeploymentJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentJobsResponse), nil
	}
}

// CreateDeploymentJobsInvoker 创建部署任务
func (c *DevStarClient) CreateDeploymentJobsInvoker(request *model.CreateDeploymentJobsRequest) *CreateDeploymentJobsInvoker {
	requestDef := GenReqDefForCreateDeploymentJobs()
	return &CreateDeploymentJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentJobs 查询应用环境部署任务详情
//
// 查询应用环境部署任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowDeploymentJobs(request *model.ShowDeploymentJobsRequest) (*model.ShowDeploymentJobsResponse, error) {
	requestDef := GenReqDefForShowDeploymentJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentJobsResponse), nil
	}
}

// ShowDeploymentJobsInvoker 查询应用环境部署任务详情
func (c *DevStarClient) ShowDeploymentJobsInvoker(request *model.ShowDeploymentJobsRequest) *ShowDeploymentJobsInvoker {
	requestDef := GenReqDefForShowDeploymentJobs()
	return &ShowDeploymentJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCodehubTemplateJob CodeHub 模板生成代码
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
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) RunCodehubTemplateJob(request *model.RunCodehubTemplateJobRequest) (*model.RunCodehubTemplateJobResponse, error) {
	requestDef := GenReqDefForRunCodehubTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCodehubTemplateJobResponse), nil
	}
}

// RunCodehubTemplateJobInvoker CodeHub 模板生成代码
func (c *DevStarClient) RunCodehubTemplateJobInvoker(request *model.RunCodehubTemplateJobRequest) *RunCodehubTemplateJobInvoker {
	requestDef := GenReqDefForRunCodehubTemplateJob()
	return &RunCodehubTemplateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunDevstarTemplateJob Devstar 模板生成代码
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
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) RunDevstarTemplateJob(request *model.RunDevstarTemplateJobRequest) (*model.RunDevstarTemplateJobResponse, error) {
	requestDef := GenReqDefForRunDevstarTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDevstarTemplateJobResponse), nil
	}
}

// RunDevstarTemplateJobInvoker Devstar 模板生成代码
func (c *DevStarClient) RunDevstarTemplateJobInvoker(request *model.RunDevstarTemplateJobRequest) *RunDevstarTemplateJobInvoker {
	requestDef := GenReqDefForRunDevstarTemplateJob()
	return &RunDevstarTemplateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 查询任务详情
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
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 查询任务详情
func (c *DevStarClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineTemplates 流水线模板列表查询
//
// 流水线模板列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListPipelineTemplates(request *model.ListPipelineTemplatesRequest) (*model.ListPipelineTemplatesResponse, error) {
	requestDef := GenReqDefForListPipelineTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineTemplatesResponse), nil
	}
}

// ListPipelineTemplatesInvoker 流水线模板列表查询
func (c *DevStarClient) ListPipelineTemplatesInvoker(request *model.ListPipelineTemplatesRequest) *ListPipelineTemplatesInvoker {
	requestDef := GenReqDefForListPipelineTemplates()
	return &ListPipelineTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipelineLastStatusV2 查询流水线最近一次运行状态查询接口
//
// 查询应用流水线最近一次运行状态查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowPipelineLastStatusV2(request *model.ShowPipelineLastStatusV2Request) (*model.ShowPipelineLastStatusV2Response, error) {
	requestDef := GenReqDefForShowPipelineLastStatusV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipelineLastStatusV2Response), nil
	}
}

// ShowPipelineLastStatusV2Invoker 查询流水线最近一次运行状态查询接口
func (c *DevStarClient) ShowPipelineLastStatusV2Invoker(request *model.ShowPipelineLastStatusV2Request) *ShowPipelineLastStatusV2Invoker {
	requestDef := GenReqDefForShowPipelineLastStatusV2()
	return &ShowPipelineLastStatusV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartPipeline 根据流水线Id操作流水线启动
//
// 根据流水线Id操作流水线启动
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) StartPipeline(request *model.StartPipelineRequest) (*model.StartPipelineResponse, error) {
	requestDef := GenReqDefForStartPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineResponse), nil
	}
}

// StartPipelineInvoker 根据流水线Id操作流水线启动
func (c *DevStarClient) StartPipelineInvoker(request *model.StartPipelineRequest) *StartPipelineInvoker {
	requestDef := GenReqDefForStartPipeline()
	return &StartPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectsV4 获取用户有权限的DevStar存量DevCloud项目列表
//
// 获取用户有权限的DevStar存量DevCloud项目列表。
// 来源包括：1.DevStar创建的DevCloud项目；2.DevStar应用有关联DevCloud项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListProjectsV4(request *model.ListProjectsV4Request) (*model.ListProjectsV4Response, error) {
	requestDef := GenReqDefForListProjectsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectsV4Response), nil
	}
}

// ListProjectsV4Invoker 获取用户有权限的DevStar存量DevCloud项目列表
func (c *DevStarClient) ListProjectsV4Invoker(request *model.ListProjectsV4Request) *ListProjectsV4Invoker {
	requestDef := GenReqDefForListProjectsV4()
	return &ListProjectsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRepositoryDuplicateName 检查仓库名称是否重名
//
// 检查仓库名称是否重名
//   - 校验规则
//     同一个项目下的仓库名称不能存在重复,当结果为true时,校验通过,仓库名称可用,否则,校验不通过,当前项目下的仓库名称已存在,不可用
//   - 必传参数
//     project_id,name,region_id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) CheckRepositoryDuplicateName(request *model.CheckRepositoryDuplicateNameRequest) (*model.CheckRepositoryDuplicateNameResponse, error) {
	requestDef := GenReqDefForCheckRepositoryDuplicateName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRepositoryDuplicateNameResponse), nil
	}
}

// CheckRepositoryDuplicateNameInvoker 检查仓库名称是否重名
func (c *DevStarClient) CheckRepositoryDuplicateNameInvoker(request *model.CheckRepositoryDuplicateNameRequest) *CheckRepositoryDuplicateNameInvoker {
	requestDef := GenReqDefForCheckRepositoryDuplicateName()
	return &CheckRepositoryDuplicateNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryByCloudIde 使用 CloudIDE 实例打开应用代码
//
// 使用 CloudIDE 实例打开应用代码。CloudIDE会保存用户项目数据，相同用户使用同一个CloudIDE，使用要求：
// - 用户需为登录状态。
// - 拥有仓库权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowRepositoryByCloudIde(request *model.ShowRepositoryByCloudIdeRequest) (*model.ShowRepositoryByCloudIdeResponse, error) {
	requestDef := GenReqDefForShowRepositoryByCloudIde()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryByCloudIdeResponse), nil
	}
}

// ShowRepositoryByCloudIdeInvoker 使用 CloudIDE 实例打开应用代码
func (c *DevStarClient) ShowRepositoryByCloudIdeInvoker(request *model.ShowRepositoryByCloudIdeRequest) *ShowRepositoryByCloudIdeInvoker {
	requestDef := GenReqDefForShowRepositoryByCloudIde()
	return &ShowRepositoryByCloudIdeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatisticalDataV2 应用代码仓库统计信息
//
// 查询代码仓库的统计信息,包括代码仓的名称,代码行数等信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowRepositoryStatisticalDataV2(request *model.ShowRepositoryStatisticalDataV2Request) (*model.ShowRepositoryStatisticalDataV2Response, error) {
	requestDef := GenReqDefForShowRepositoryStatisticalDataV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticalDataV2Response), nil
	}
}

// ShowRepositoryStatisticalDataV2Invoker 应用代码仓库统计信息
func (c *DevStarClient) ShowRepositoryStatisticalDataV2Invoker(request *model.ShowRepositoryStatisticalDataV2Request) *ShowRepositoryStatisticalDataV2Invoker {
	requestDef := GenReqDefForShowRepositoryStatisticalDataV2()
	return &ShowRepositoryStatisticalDataV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateFile 读取模板文件
//
// 该接口可以用于模板作者或模板维护人读取模板文件内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateFile(request *model.ShowTemplateFileRequest) (*model.ShowTemplateFileResponse, error) {
	requestDef := GenReqDefForShowTemplateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateFileResponse), nil
	}
}

// ShowTemplateFileInvoker 读取模板文件
func (c *DevStarClient) ShowTemplateFileInvoker(request *model.ShowTemplateFileRequest) *ShowTemplateFileInvoker {
	requestDef := GenReqDefForShowTemplateFile()
	return &ShowTemplateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplateViewHistories 同步模板浏览记录
//
// 未登录状态下，将用户浏览过的模板缓存在浏览器中，登录时，调用该接口同步模板浏览记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) CreateTemplateViewHistories(request *model.CreateTemplateViewHistoriesRequest) (*model.CreateTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForCreateTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateViewHistoriesResponse), nil
	}
}

// CreateTemplateViewHistoriesInvoker 同步模板浏览记录
func (c *DevStarClient) CreateTemplateViewHistoriesInvoker(request *model.CreateTemplateViewHistoriesRequest) *CreateTemplateViewHistoriesInvoker {
	requestDef := GenReqDefForCreateTemplateViewHistories()
	return &CreateTemplateViewHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListPublishedTemplates 查询模板列表（V1）
//
// 查询模板列表，推荐使用/v1/templates/query接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListPublishedTemplates(request *model.ListPublishedTemplatesRequest) (*model.ListPublishedTemplatesResponse, error) {
	requestDef := GenReqDefForListPublishedTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublishedTemplatesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListPublishedTemplatesInvoker 查询模板列表（V1）
func (c *DevStarClient) ListPublishedTemplatesInvoker(request *model.ListPublishedTemplatesRequest) *ListPublishedTemplatesInvoker {
	requestDef := GenReqDefForListPublishedTemplates()
	return &ListPublishedTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateViewHistories 我浏览的模板记录
//
// 查询DevStar或者CodeLabs登录用户浏览过的模板（只返回最近浏览的5个模板）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListTemplateViewHistories(request *model.ListTemplateViewHistoriesRequest) (*model.ListTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForListTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateViewHistoriesResponse), nil
	}
}

// ListTemplateViewHistoriesInvoker 我浏览的模板记录
func (c *DevStarClient) ListTemplateViewHistoriesInvoker(request *model.ListTemplateViewHistoriesRequest) *ListTemplateViewHistoriesInvoker {
	requestDef := GenReqDefForListTemplateViewHistories()
	return &ListTemplateViewHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 查询模板列表
//
// 查询模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 查询模板列表
func (c *DevStarClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTemplatesV2 查询模板列表（V2）
//
// 查询模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ListTemplatesV2(request *model.ListTemplatesV2Request) (*model.ListTemplatesV2Response, error) {
	requestDef := GenReqDefForListTemplatesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesV2Response), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTemplatesV2Invoker 查询模板列表（V2）
func (c *DevStarClient) ListTemplatesV2Invoker(request *model.ListTemplatesV2Request) *ListTemplatesV2Invoker {
	requestDef := GenReqDefForListTemplatesV2()
	return &ListTemplatesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateV3 查询模板详情（V3）
//
// 获取指定模板详情，包括模板id、名称、描述、作者、标签、上架时间等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateV3(request *model.ShowTemplateV3Request) (*model.ShowTemplateV3Response, error) {
	requestDef := GenReqDefForShowTemplateV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateV3Response), nil
	}
}

// ShowTemplateV3Invoker 查询模板详情（V3）
func (c *DevStarClient) ShowTemplateV3Invoker(request *model.ShowTemplateV3Request) *ShowTemplateV3Invoker {
	requestDef := GenReqDefForShowTemplateV3()
	return &ShowTemplateV3Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateDetail 查询模板详情（V1）
//
// 查询模板详情，推荐使用V3版本接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DevStarClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}

// ShowTemplateDetailInvoker 查询模板详情（V1）
func (c *DevStarClient) ShowTemplateDetailInvoker(request *model.ShowTemplateDetailRequest) *ShowTemplateDetailInvoker {
	requestDef := GenReqDefForShowTemplateDetail()
	return &ShowTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
