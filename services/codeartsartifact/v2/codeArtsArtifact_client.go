package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsartifact/v2/model"
)

type CodeArtsArtifactClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsArtifactClient(hcClient *httpclient.HcHttpClient) *CodeArtsArtifactClient {
	return &CodeArtsArtifactClient{HcClient: hcClient}
}

func CodeArtsArtifactClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchDeleteTrashes 批量删除回收站
//
// 批量删除回收站
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) BatchDeleteTrashes(request *model.BatchDeleteTrashesRequest) (*model.BatchDeleteTrashesResponse, error) {
	requestDef := GenReqDefForBatchDeleteTrashes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTrashesResponse), nil
	}
}

// BatchDeleteTrashesInvoker 批量删除回收站
func (c *CodeArtsArtifactClient) BatchDeleteTrashesInvoker(request *model.BatchDeleteTrashesRequest) *BatchDeleteTrashesInvoker {
	requestDef := GenReqDefForBatchDeleteTrashes()
	return &BatchDeleteTrashesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRestoreRepo 批量还原回收站
//
// 批量还原回收站
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) BatchRestoreRepo(request *model.BatchRestoreRepoRequest) (*model.BatchRestoreRepoResponse, error) {
	requestDef := GenReqDefForBatchRestoreRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestoreRepoResponse), nil
	}
}

// BatchRestoreRepoInvoker 批量还原回收站
func (c *CodeArtsArtifactClient) BatchRestoreRepoInvoker(request *model.BatchRestoreRepoRequest) *BatchRestoreRepoInvoker {
	requestDef := GenReqDefForBatchRestoreRepo()
	return &BatchRestoreRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateArtifactory 创建非maven仓库
//
// 创建非maven仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) CreateArtifactory(request *model.CreateArtifactoryRequest) (*model.CreateArtifactoryResponse, error) {
	requestDef := GenReqDefForCreateArtifactory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateArtifactoryResponse), nil
	}
}

// CreateArtifactoryInvoker 创建非maven仓库
func (c *CodeArtsArtifactClient) CreateArtifactoryInvoker(request *model.CreateArtifactoryRequest) *CreateArtifactoryInvoker {
	requestDef := GenReqDefForCreateArtifactory()
	return &CreateArtifactoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAttention 关注组件/取消关注组件
//
// 关注组件/取消关注组件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) CreateAttention(request *model.CreateAttentionRequest) (*model.CreateAttentionResponse, error) {
	requestDef := GenReqDefForCreateAttention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAttentionResponse), nil
	}
}

// CreateAttentionInvoker 关注组件/取消关注组件
func (c *CodeArtsArtifactClient) CreateAttentionInvoker(request *model.CreateAttentionRequest) *CreateAttentionInvoker {
	requestDef := GenReqDefForCreateAttention()
	return &CreateAttentionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDockerRepositories 创建docker仓库
//
// 创建docker仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) CreateDockerRepositories(request *model.CreateDockerRepositoriesRequest) (*model.CreateDockerRepositoriesResponse, error) {
	requestDef := GenReqDefForCreateDockerRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDockerRepositoriesResponse), nil
	}
}

// CreateDockerRepositoriesInvoker 创建docker仓库
func (c *CodeArtsArtifactClient) CreateDockerRepositoriesInvoker(request *model.CreateDockerRepositoriesRequest) *CreateDockerRepositoriesInvoker {
	requestDef := GenReqDefForCreateDockerRepositories()
	return &CreateDockerRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMavenRepo 创建maven仓库
//
// 创建maven仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) CreateMavenRepo(request *model.CreateMavenRepoRequest) (*model.CreateMavenRepoResponse, error) {
	requestDef := GenReqDefForCreateMavenRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMavenRepoResponse), nil
	}
}

// CreateMavenRepoInvoker 创建maven仓库
func (c *CodeArtsArtifactClient) CreateMavenRepoInvoker(request *model.CreateMavenRepoRequest) *CreateMavenRepoInvoker {
	requestDef := GenReqDefForCreateMavenRepo()
	return &CreateMavenRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectRelatedRepository 创建项目关联仓库
//
// 创建项目管理关联仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) CreateProjectRelatedRepository(request *model.CreateProjectRelatedRepositoryRequest) (*model.CreateProjectRelatedRepositoryResponse, error) {
	requestDef := GenReqDefForCreateProjectRelatedRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectRelatedRepositoryResponse), nil
	}
}

// CreateProjectRelatedRepositoryInvoker 创建项目关联仓库
func (c *CodeArtsArtifactClient) CreateProjectRelatedRepositoryInvoker(request *model.CreateProjectRelatedRepositoryRequest) *CreateProjectRelatedRepositoryInvoker {
	requestDef := GenReqDefForCreateProjectRelatedRepository()
	return &CreateProjectRelatedRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteArtifactFile 非maven删除文件
//
// 非maven删除文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) DeleteArtifactFile(request *model.DeleteArtifactFileRequest) (*model.DeleteArtifactFileResponse, error) {
	requestDef := GenReqDefForDeleteArtifactFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteArtifactFileResponse), nil
	}
}

// DeleteArtifactFileInvoker 非maven删除文件
func (c *CodeArtsArtifactClient) DeleteArtifactFileInvoker(request *model.DeleteArtifactFileRequest) *DeleteArtifactFileInvoker {
	requestDef := GenReqDefForDeleteArtifactFile()
	return &DeleteArtifactFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCompletelyUpdateFileState 彻底删除文件/文件夹
//
// 根据文件ID彻底删除文件或文件夹，删除后不能恢复，支持批量删除。可使用该接口清理不再需要的文件或文件夹以释放存储空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) DeleteCompletelyUpdateFileState(request *model.DeleteCompletelyUpdateFileStateRequest) (*model.DeleteCompletelyUpdateFileStateResponse, error) {
	requestDef := GenReqDefForDeleteCompletelyUpdateFileState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCompletelyUpdateFileStateResponse), nil
	}
}

// DeleteCompletelyUpdateFileStateInvoker 彻底删除文件/文件夹
func (c *CodeArtsArtifactClient) DeleteCompletelyUpdateFileStateInvoker(request *model.DeleteCompletelyUpdateFileStateRequest) *DeleteCompletelyUpdateFileStateInvoker {
	requestDef := GenReqDefForDeleteCompletelyUpdateFileState()
	return &DeleteCompletelyUpdateFileStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepository 删除仓库到回收站
//
// 删除仓库到回收站
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) DeleteRepository(request *model.DeleteRepositoryRequest) (*model.DeleteRepositoryResponse, error) {
	requestDef := GenReqDefForDeleteRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepositoryResponse), nil
	}
}

// DeleteRepositoryInvoker 删除仓库到回收站
func (c *CodeArtsArtifactClient) DeleteRepositoryInvoker(request *model.DeleteRepositoryRequest) *DeleteRepositoryInvoker {
	requestDef := GenReqDefForDeleteRepository()
	return &DeleteRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllRepositories 查询仓库详情，不会去统计仓库下的制品数量
//
// 查询仓库详情，不会去统计仓库下的制品数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListAllRepositories(request *model.ListAllRepositoriesRequest) (*model.ListAllRepositoriesResponse, error) {
	requestDef := GenReqDefForListAllRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllRepositoriesResponse), nil
	}
}

// ListAllRepositoriesInvoker 查询仓库详情，不会去统计仓库下的制品数量
func (c *CodeArtsArtifactClient) ListAllRepositoriesInvoker(request *model.ListAllRepositoriesRequest) *ListAllRepositoriesInvoker {
	requestDef := GenReqDefForListAllRepositories()
	return &ListAllRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArtifactoryComponent 查询仓库文件详情
//
// 查询仓库文件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListArtifactoryComponent(request *model.ListArtifactoryComponentRequest) (*model.ListArtifactoryComponentResponse, error) {
	requestDef := GenReqDefForListArtifactoryComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArtifactoryComponentResponse), nil
	}
}

// ListArtifactoryComponentInvoker 查询仓库文件详情
func (c *CodeArtsArtifactClient) ListArtifactoryComponentInvoker(request *model.ListArtifactoryComponentRequest) *ListArtifactoryComponentInvoker {
	requestDef := GenReqDefForListArtifactoryComponent()
	return &ListArtifactoryComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArtifactoryStorageStatistic 查询存储容量趋势
//
// 查询存储容量趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListArtifactoryStorageStatistic(request *model.ListArtifactoryStorageStatisticRequest) (*model.ListArtifactoryStorageStatisticResponse, error) {
	requestDef := GenReqDefForListArtifactoryStorageStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArtifactoryStorageStatisticResponse), nil
	}
}

// ListArtifactoryStorageStatisticInvoker 查询存储容量趋势
func (c *CodeArtsArtifactClient) ListArtifactoryStorageStatisticInvoker(request *model.ListArtifactoryStorageStatisticRequest) *ListArtifactoryStorageStatisticInvoker {
	requestDef := GenReqDefForListArtifactoryStorageStatistic()
	return &ListArtifactoryStorageStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttentions 查询关注列表
//
// 查询关注列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListAttentions(request *model.ListAttentionsRequest) (*model.ListAttentionsResponse, error) {
	requestDef := GenReqDefForListAttentions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttentionsResponse), nil
	}
}

// ListAttentionsInvoker 查询关注列表
func (c *CodeArtsArtifactClient) ListAttentionsInvoker(request *model.ListAttentionsRequest) *ListAttentionsInvoker {
	requestDef := GenReqDefForListAttentions()
	return &ListAttentionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCapacityMessageSetting 查询租户容量消息通知设置信息
//
// 查询租户容量消息通知设置，包含容量阈值和是否启用邮件通知等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListCapacityMessageSetting(request *model.ListCapacityMessageSettingRequest) (*model.ListCapacityMessageSettingResponse, error) {
	requestDef := GenReqDefForListCapacityMessageSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCapacityMessageSettingResponse), nil
	}
}

// ListCapacityMessageSettingInvoker 查询租户容量消息通知设置信息
func (c *CodeArtsArtifactClient) ListCapacityMessageSettingInvoker(request *model.ListCapacityMessageSettingRequest) *ListCapacityMessageSettingInvoker {
	requestDef := GenReqDefForListCapacityMessageSetting()
	return &ListCapacityMessageSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListChildProxyRepositoriesList 获取聚合仓下的仓库代理列表
//
// 根据仓库ID获取指定聚合仓的仓库代理列表，包含仓库状态、类型、地址和访问路径白名单等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListChildProxyRepositoriesList(request *model.ListChildProxyRepositoriesListRequest) (*model.ListChildProxyRepositoriesListResponse, error) {
	requestDef := GenReqDefForListChildProxyRepositoriesList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChildProxyRepositoriesListResponse), nil
	}
}

// ListChildProxyRepositoriesListInvoker 获取聚合仓下的仓库代理列表
func (c *CodeArtsArtifactClient) ListChildProxyRepositoriesListInvoker(request *model.ListChildProxyRepositoriesListRequest) *ListChildProxyRepositoriesListInvoker {
	requestDef := GenReqDefForListChildProxyRepositoriesList()
	return &ListChildProxyRepositoriesListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainIpConfig 查询租户级IP白名单
//
// 查询租户级IP白名单列表。在IP白名单的IP才能访问制品仓库，未配置IP白名单时，默认所有IP都可访问。该功能可用于保障制品仓库的安全，对访问IP进行严格控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListDomainIpConfig(request *model.ListDomainIpConfigRequest) (*model.ListDomainIpConfigResponse, error) {
	requestDef := GenReqDefForListDomainIpConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainIpConfigResponse), nil
	}
}

// ListDomainIpConfigInvoker 查询租户级IP白名单
func (c *CodeArtsArtifactClient) ListDomainIpConfigInvoker(request *model.ListDomainIpConfigRequest) *ListDomainIpConfigInvoker {
	requestDef := GenReqDefForListDomainIpConfig()
	return &ListDomainIpConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileBuildArchives 分页查询构建归档包列表
//
// 当归档包数量庞大时，分页查询构建归档包列表，包含文件名、文件大小、下载地址、MD5校验和、构建地址、构建代码分支等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListFileBuildArchives(request *model.ListFileBuildArchivesRequest) (*model.ListFileBuildArchivesResponse, error) {
	requestDef := GenReqDefForListFileBuildArchives()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileBuildArchivesResponse), nil
	}
}

// ListFileBuildArchivesInvoker 分页查询构建归档包列表
func (c *CodeArtsArtifactClient) ListFileBuildArchivesInvoker(request *model.ListFileBuildArchivesRequest) *ListFileBuildArchivesInvoker {
	requestDef := GenReqDefForListFileBuildArchives()
	return &ListFileBuildArchivesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFiles 查询文件/项目列表
//
// 当项目或文件数量庞大时，分页查询项目或文件列表。可根据文件名、文件状态和文件的发布状态等参数进行过滤，从而快速找到所需文件，包含文件名、文件大小和下载地址等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListFiles(request *model.ListFilesRequest) (*model.ListFilesResponse, error) {
	requestDef := GenReqDefForListFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFilesResponse), nil
	}
}

// ListFilesInvoker 查询文件/项目列表
func (c *CodeArtsArtifactClient) ListFilesInvoker(request *model.ListFilesRequest) *ListFilesInvoker {
	requestDef := GenReqDefForListFiles()
	return &ListFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatestVersionFiles 查询项目下所有文件的最新版本
//
// 当项目文件数量庞大时，通过该接口可以分页查询项目下所有文件的最新版本，包含文件名、文件大小、文件状态和发布状态等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListLatestVersionFiles(request *model.ListLatestVersionFilesRequest) (*model.ListLatestVersionFilesResponse, error) {
	requestDef := GenReqDefForListLatestVersionFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatestVersionFilesResponse), nil
	}
}

// ListLatestVersionFilesInvoker 查询项目下所有文件的最新版本
func (c *CodeArtsArtifactClient) ListLatestVersionFilesInvoker(request *model.ListLatestVersionFilesRequest) *ListLatestVersionFilesInvoker {
	requestDef := GenReqDefForListLatestVersionFiles()
	return &ListLatestVersionFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMavenList 查询Maven仓库列表
//
// 查询Maven仓库列表，包含仓库状态、类型、地址和访问路径白名单等信息。支持根据项目ID和仓库ID等参数进行过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListMavenList(request *model.ListMavenListRequest) (*model.ListMavenListResponse, error) {
	requestDef := GenReqDefForListMavenList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMavenListResponse), nil
	}
}

// ListMavenListInvoker 查询Maven仓库列表
func (c *CodeArtsArtifactClient) ListMavenListInvoker(request *model.ListMavenListRequest) *ListMavenListInvoker {
	requestDef := GenReqDefForListMavenList()
	return &ListMavenListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMavenUser 查询私有库用户列表
//
// 分页查询私有库用户列表，包含用户名和用户是否启用等信息。可根据用户名进行过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListMavenUser(request *model.ListMavenUserRequest) (*model.ListMavenUserResponse, error) {
	requestDef := GenReqDefForListMavenUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMavenUserResponse), nil
	}
}

// ListMavenUserInvoker 查询私有库用户列表
func (c *CodeArtsArtifactClient) ListMavenUserInvoker(request *model.ListMavenUserRequest) *ListMavenUserInvoker {
	requestDef := GenReqDefForListMavenUser()
	return &ListMavenUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetProxy 查询网络代理列表
//
// 查询网络代理列表，返回代理源的访问地址及认证信息等，用于代理外部公开的制品资源。通过网络代理，开发人员可以安全、高效地访问所需的外部资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListNetProxy(request *model.ListNetProxyRequest) (*model.ListNetProxyResponse, error) {
	requestDef := GenReqDefForListNetProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetProxyResponse), nil
	}
}

// ListNetProxyInvoker 查询网络代理列表
func (c *CodeArtsArtifactClient) ListNetProxyInvoker(request *model.ListNetProxyRequest) *ListNetProxyInvoker {
	requestDef := GenReqDefForListNetProxy()
	return &ListNetProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectRolePermissions 查看项目的角色权限设置
//
// 查看项目的角色权限设置，包含上传下载、创建文件夹、清空或者还原回收站和更改软件包状态等设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListProjectRolePermissions(request *model.ListProjectRolePermissionsRequest) (*model.ListProjectRolePermissionsResponse, error) {
	requestDef := GenReqDefForListProjectRolePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectRolePermissionsResponse), nil
	}
}

// ListProjectRolePermissionsInvoker 查看项目的角色权限设置
func (c *CodeArtsArtifactClient) ListProjectRolePermissionsInvoker(request *model.ListProjectRolePermissionsRequest) *ListProjectRolePermissionsInvoker {
	requestDef := GenReqDefForListProjectRolePermissions()
	return &ListProjectRolePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectUsers 查询项目下的用户
//
// 当项目中的用户数量较多时，分页查询指定项目下的用户列表，包含用户名和角色等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListProjectUsers(request *model.ListProjectUsersRequest) (*model.ListProjectUsersResponse, error) {
	requestDef := GenReqDefForListProjectUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectUsersResponse), nil
	}
}

// ListProjectUsersInvoker 查询项目下的用户
func (c *CodeArtsArtifactClient) ListProjectUsersInvoker(request *model.ListProjectUsersRequest) *ListProjectUsersInvoker {
	requestDef := GenReqDefForListProjectUsers()
	return &ListProjectUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecGuardList 查询制品安全扫描任务列表
//
// 分页查询制品安全扫描任务列表，包含扫描制品数量、漏洞数量、病毒文件数量和恶意代码文件数量等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ListSecGuardList(request *model.ListSecGuardListRequest) (*model.ListSecGuardListResponse, error) {
	requestDef := GenReqDefForListSecGuardList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecGuardListResponse), nil
	}
}

// ListSecGuardListInvoker 查询制品安全扫描任务列表
func (c *CodeArtsArtifactClient) ListSecGuardListInvoker(request *model.ListSecGuardListRequest) *ListSecGuardListInvoker {
	requestDef := GenReqDefForListSecGuardList()
	return &ListSecGuardListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyRepository 编辑仓库
//
// 编辑仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ModifyRepository(request *model.ModifyRepositoryRequest) (*model.ModifyRepositoryResponse, error) {
	requestDef := GenReqDefForModifyRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyRepositoryResponse), nil
	}
}

// ModifyRepositoryInvoker 编辑仓库
func (c *CodeArtsArtifactClient) ModifyRepositoryInvoker(request *model.ModifyRepositoryRequest) *ModifyRepositoryInvoker {
	requestDef := GenReqDefForModifyRepository()
	return &ModifyRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetUserPassword 重置用户密码
//
// 重置用户密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ResetUserPassword(request *model.ResetUserPasswordRequest) (*model.ResetUserPasswordResponse, error) {
	requestDef := GenReqDefForResetUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetUserPasswordResponse), nil
	}
}

// ResetUserPasswordInvoker 重置用户密码
func (c *CodeArtsArtifactClient) ResetUserPasswordInvoker(request *model.ResetUserPasswordRequest) *ResetUserPasswordInvoker {
	requestDef := GenReqDefForResetUserPassword()
	return &ResetUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchArtifacts 统筹搜索
//
// 统筹搜索
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) SearchArtifacts(request *model.SearchArtifactsRequest) (*model.SearchArtifactsResponse, error) {
	requestDef := GenReqDefForSearchArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchArtifactsResponse), nil
	}
}

// SearchArtifactsInvoker 统筹搜索
func (c *CodeArtsArtifactClient) SearchArtifactsInvoker(request *model.SearchArtifactsRequest) *SearchArtifactsInvoker {
	requestDef := GenReqDefForSearchArtifacts()
	return &SearchArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchByChecksum 通过checksum搜索文件
//
// 通过checksum搜索文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) SearchByChecksum(request *model.SearchByChecksumRequest) (*model.SearchByChecksumResponse, error) {
	requestDef := GenReqDefForSearchByChecksum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchByChecksumResponse), nil
	}
}

// SearchByChecksumInvoker 通过checksum搜索文件
func (c *CodeArtsArtifactClient) SearchByChecksumInvoker(request *model.SearchByChecksumRequest) *SearchByChecksumInvoker {
	requestDef := GenReqDefForSearchByChecksum()
	return &SearchByChecksumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAudit 查询仓库或文件的审计日志信息
//
// 查询仓库或文件的审计日志信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowAudit(request *model.ShowAuditRequest) (*model.ShowAuditResponse, error) {
	requestDef := GenReqDefForShowAudit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditResponse), nil
	}
}

// ShowAuditInvoker 查询仓库或文件的审计日志信息
func (c *CodeArtsArtifactClient) ShowAuditInvoker(request *model.ShowAuditRequest) *ShowAuditInvoker {
	requestDef := GenReqDefForShowAudit()
	return &ShowAuditInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoDeleteJobSettings 查询项目自动删除任务设置
//
// 查询项目自动删除任务设置，包含文件的过期自动删除时间及删除规则。自动删除任务可以自动执行文件清理任务，减少管理员的工作负担，确保存储资源的有效利用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowAutoDeleteJobSettings(request *model.ShowAutoDeleteJobSettingsRequest) (*model.ShowAutoDeleteJobSettingsResponse, error) {
	requestDef := GenReqDefForShowAutoDeleteJobSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoDeleteJobSettingsResponse), nil
	}
}

// ShowAutoDeleteJobSettingsInvoker 查询项目自动删除任务设置
func (c *CodeArtsArtifactClient) ShowAutoDeleteJobSettingsInvoker(request *model.ShowAutoDeleteJobSettingsRequest) *ShowAutoDeleteJobSettingsInvoker {
	requestDef := GenReqDefForShowAutoDeleteJobSettings()
	return &ShowAutoDeleteJobSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainReleaseRepoStorage 查询租户发布库存储容量
//
// 查询租户发布库存储容量，包含已使用存储容量、最大存储容量、套餐类型和仓库数量等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowDomainReleaseRepoStorage(request *model.ShowDomainReleaseRepoStorageRequest) (*model.ShowDomainReleaseRepoStorageResponse, error) {
	requestDef := GenReqDefForShowDomainReleaseRepoStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainReleaseRepoStorageResponse), nil
	}
}

// ShowDomainReleaseRepoStorageInvoker 查询租户发布库存储容量
func (c *CodeArtsArtifactClient) ShowDomainReleaseRepoStorageInvoker(request *model.ShowDomainReleaseRepoStorageRequest) *ShowDomainReleaseRepoStorageInvoker {
	requestDef := GenReqDefForShowDomainReleaseRepoStorage()
	return &ShowDomainReleaseRepoStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileDetail 根据文件ID查询文件详情
//
// 在日常数据管理工作中，根据文件ID查询指定文件详情，包含文件名、文件大小、下载地址、校验和、文件状态和发布状态等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowFileDetail(request *model.ShowFileDetailRequest) (*model.ShowFileDetailResponse, error) {
	requestDef := GenReqDefForShowFileDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileDetailResponse), nil
	}
}

// ShowFileDetailInvoker 根据文件ID查询文件详情
func (c *CodeArtsArtifactClient) ShowFileDetailInvoker(request *model.ShowFileDetailRequest) *ShowFileDetailInvoker {
	requestDef := GenReqDefForShowFileDetail()
	return &ShowFileDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileDetailByFullName 根据文件完整路径查询文件详情
//
// 在日常数据管理工作中，根据文件完整路径查询指定文件详情，包含文件名、文件大小、下载地址、校验和、文件状态和发布状态等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowFileDetailByFullName(request *model.ShowFileDetailByFullNameRequest) (*model.ShowFileDetailByFullNameResponse, error) {
	requestDef := GenReqDefForShowFileDetailByFullName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileDetailByFullNameResponse), nil
	}
}

// ShowFileDetailByFullNameInvoker 根据文件完整路径查询文件详情
func (c *CodeArtsArtifactClient) ShowFileDetailByFullNameInvoker(request *model.ShowFileDetailByFullNameRequest) *ShowFileDetailByFullNameInvoker {
	requestDef := GenReqDefForShowFileDetailByFullName()
	return &ShowFileDetailByFullNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileTree 查询仓库文件夹目录
//
// 查询仓库文件夹目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowFileTree(request *model.ShowFileTreeRequest) (*model.ShowFileTreeResponse, error) {
	requestDef := GenReqDefForShowFileTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileTreeResponse), nil
	}
}

// ShowFileTreeInvoker 查询仓库文件夹目录
func (c *CodeArtsArtifactClient) ShowFileTreeInvoker(request *model.ShowFileTreeRequest) *ShowFileTreeInvoker {
	requestDef := GenReqDefForShowFileTree()
	return &ShowFileTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestVersionFilesCount 查询项目下所有文件的数量
//
// 查询项目下所有文件的数量，该接口会识别所有文件的最新版本，从而提供准确的文件数量统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowLatestVersionFilesCount(request *model.ShowLatestVersionFilesCountRequest) (*model.ShowLatestVersionFilesCountResponse, error) {
	requestDef := GenReqDefForShowLatestVersionFilesCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestVersionFilesCountResponse), nil
	}
}

// ShowLatestVersionFilesCountInvoker 查询项目下所有文件的数量
func (c *CodeArtsArtifactClient) ShowLatestVersionFilesCountInvoker(request *model.ShowLatestVersionFilesCountRequest) *ShowLatestVersionFilesCountInvoker {
	requestDef := GenReqDefForShowLatestVersionFilesCount()
	return &ShowLatestVersionFilesCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMavenInfo 查询租户Maven仓库列表和账号密码
//
// 查询租户Maven仓库列表和账号密码，支持跨租户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowMavenInfo(request *model.ShowMavenInfoRequest) (*model.ShowMavenInfoResponse, error) {
	requestDef := GenReqDefForShowMavenInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMavenInfoResponse), nil
	}
}

// ShowMavenInfoInvoker 查询租户Maven仓库列表和账号密码
func (c *CodeArtsArtifactClient) ShowMavenInfoInvoker(request *model.ShowMavenInfoRequest) *ShowMavenInfoInvoker {
	requestDef := GenReqDefForShowMavenInfo()
	return &ShowMavenInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMultiRolesUserPermissions 查询多角色用户权限
//
// 查询多角色用户权限，包含上传下载、创建文件夹、清空或者还原回收站和更改软件包状态等权限信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowMultiRolesUserPermissions(request *model.ShowMultiRolesUserPermissionsRequest) (*model.ShowMultiRolesUserPermissionsResponse, error) {
	requestDef := GenReqDefForShowMultiRolesUserPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiRolesUserPermissionsResponse), nil
	}
}

// ShowMultiRolesUserPermissionsInvoker 查询多角色用户权限
func (c *CodeArtsArtifactClient) ShowMultiRolesUserPermissionsInvoker(request *model.ShowMultiRolesUserPermissionsRequest) *ShowMultiRolesUserPermissionsInvoker {
	requestDef := GenReqDefForShowMultiRolesUserPermissions()
	return &ShowMultiRolesUserPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOpenSourceEnabled 查询中心仓是否启用
//
// 查询中心仓是否启用，用于确定当前局点是否具备中心仓功能，从而确保业务流程的顺利进行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowOpenSourceEnabled(request *model.ShowOpenSourceEnabledRequest) (*model.ShowOpenSourceEnabledResponse, error) {
	requestDef := GenReqDefForShowOpenSourceEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOpenSourceEnabledResponse), nil
	}
}

// ShowOpenSourceEnabledInvoker 查询中心仓是否启用
func (c *CodeArtsArtifactClient) ShowOpenSourceEnabledInvoker(request *model.ShowOpenSourceEnabledRequest) *ShowOpenSourceEnabledInvoker {
	requestDef := GenReqDefForShowOpenSourceEnabled()
	return &ShowOpenSourceEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPackageDataDetail 获取当前用户的套餐信息
//
// 获取当前用户的套餐信息，包含总存储容量和已使用存储容量等信息，以便合理规划资源使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowPackageDataDetail(request *model.ShowPackageDataDetailRequest) (*model.ShowPackageDataDetailResponse, error) {
	requestDef := GenReqDefForShowPackageDataDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPackageDataDetailResponse), nil
	}
}

// ShowPackageDataDetailInvoker 获取当前用户的套餐信息
func (c *CodeArtsArtifactClient) ShowPackageDataDetailInvoker(request *model.ShowPackageDataDetailRequest) *ShowPackageDataDetailInvoker {
	requestDef := GenReqDefForShowPackageDataDetail()
	return &ShowPackageDataDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPackageInfo 获取当前用户的套餐状态
//
// 获取当前用户的套餐状态，包含套餐扩展包的容量和流量规格，如资源类型、套餐状态、剩余天数等信息，帮助用户高效管理云资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowPackageInfo(request *model.ShowPackageInfoRequest) (*model.ShowPackageInfoResponse, error) {
	requestDef := GenReqDefForShowPackageInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPackageInfoResponse), nil
	}
}

// ShowPackageInfoInvoker 获取当前用户的套餐状态
func (c *CodeArtsArtifactClient) ShowPackageInfoInvoker(request *model.ShowPackageInfoRequest) *ShowPackageInfoInvoker {
	requestDef := GenReqDefForShowPackageInfo()
	return &ShowPackageInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectList 查询项目管理关联仓库
//
// 查询项目管理关联仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowProjectList(request *model.ShowProjectListRequest) (*model.ShowProjectListResponse, error) {
	requestDef := GenReqDefForShowProjectList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectListResponse), nil
	}
}

// ShowProjectListInvoker 查询项目管理关联仓库
func (c *CodeArtsArtifactClient) ShowProjectListInvoker(request *model.ShowProjectListRequest) *ShowProjectListInvoker {
	requestDef := GenReqDefForShowProjectList()
	return &ShowProjectListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectReleaseFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowProjectReleaseFiles(request *model.ShowProjectReleaseFilesRequest) (*model.ShowProjectReleaseFilesResponse, error) {
	requestDef := GenReqDefForShowProjectReleaseFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectReleaseFilesResponse), nil
	}
}

// ShowProjectReleaseFilesInvoker 获取项目下文件版本信息列表
func (c *CodeArtsArtifactClient) ShowProjectReleaseFilesInvoker(request *model.ShowProjectReleaseFilesRequest) *ShowProjectReleaseFilesInvoker {
	requestDef := GenReqDefForShowProjectReleaseFiles()
	return &ShowProjectReleaseFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectStorageInfo 查询项目下的制品存储容量信息
//
// 查询项目下的制品存储容量，包含已使用存储容量和文件数量等信息。在项目管理中，可以使用该接口监控项目下的制品存储情况，以确保资源的有效利用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowProjectStorageInfo(request *model.ShowProjectStorageInfoRequest) (*model.ShowProjectStorageInfoResponse, error) {
	requestDef := GenReqDefForShowProjectStorageInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectStorageInfoResponse), nil
	}
}

// ShowProjectStorageInfoInvoker 查询项目下的制品存储容量信息
func (c *CodeArtsArtifactClient) ShowProjectStorageInfoInvoker(request *model.ShowProjectStorageInfoRequest) *ShowProjectStorageInfoInvoker {
	requestDef := GenReqDefForShowProjectStorageInfo()
	return &ShowProjectStorageInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowReleaseProjectFiles 获取项目下文件版本信息列表
//
// 获取项目下文件版本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowReleaseProjectFiles(request *model.ShowReleaseProjectFilesRequest) (*model.ShowReleaseProjectFilesResponse, error) {
	requestDef := GenReqDefForShowReleaseProjectFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReleaseProjectFilesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowReleaseProjectFilesInvoker 获取项目下文件版本信息列表
func (c *CodeArtsArtifactClient) ShowReleaseProjectFilesInvoker(request *model.ShowReleaseProjectFilesRequest) *ShowReleaseProjectFilesInvoker {
	requestDef := GenReqDefForShowReleaseProjectFiles()
	return &ShowReleaseProjectFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepoUserInfo 查询租户私有依赖库的账号密码
//
// 在自动化构建场景，用户可调用该接口查询租户私有依赖库的账号密码，以便进行后续的上传下载操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowRepoUserInfo(request *model.ShowRepoUserInfoRequest) (*model.ShowRepoUserInfoResponse, error) {
	requestDef := GenReqDefForShowRepoUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepoUserInfoResponse), nil
	}
}

// ShowRepoUserInfoInvoker 查询租户私有依赖库的账号密码
func (c *CodeArtsArtifactClient) ShowRepoUserInfoInvoker(request *model.ShowRepoUserInfoRequest) *ShowRepoUserInfoInvoker {
	requestDef := GenReqDefForShowRepoUserInfo()
	return &ShowRepoUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepository 查询单个仓库详细信息，会去统计仓库下的制品数量
//
// 查询单个仓库详细信息，会去统计仓库下的制品数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowRepository(request *model.ShowRepositoryRequest) (*model.ShowRepositoryResponse, error) {
	requestDef := GenReqDefForShowRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryResponse), nil
	}
}

// ShowRepositoryInvoker 查询单个仓库详细信息，会去统计仓库下的制品数量
func (c *CodeArtsArtifactClient) ShowRepositoryInvoker(request *model.ShowRepositoryRequest) *ShowRepositoryInvoker {
	requestDef := GenReqDefForShowRepository()
	return &ShowRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryInfo 查看仓库信息
//
// 查看仓库信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowRepositoryInfo(request *model.ShowRepositoryInfoRequest) (*model.ShowRepositoryInfoResponse, error) {
	requestDef := GenReqDefForShowRepositoryInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryInfoResponse), nil
	}
}

// ShowRepositoryInfoInvoker 查看仓库信息
func (c *CodeArtsArtifactClient) ShowRepositoryInfoInvoker(request *model.ShowRepositoryInfoRequest) *ShowRepositoryInfoInvoker {
	requestDef := GenReqDefForShowRepositoryInfo()
	return &ShowRepositoryInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStorage 仓库用量查询
//
// 仓库用量查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowStorage(request *model.ShowStorageRequest) (*model.ShowStorageResponse, error) {
	requestDef := GenReqDefForShowStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStorageResponse), nil
	}
}

// ShowStorageInvoker 仓库用量查询
func (c *CodeArtsArtifactClient) ShowStorageInvoker(request *model.ShowStorageRequest) *ShowStorageInvoker {
	requestDef := GenReqDefForShowStorage()
	return &ShowStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserPrivileges 查询用户在项目下的权限
//
// 查询用户在项目下的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowUserPrivileges(request *model.ShowUserPrivilegesRequest) (*model.ShowUserPrivilegesResponse, error) {
	requestDef := GenReqDefForShowUserPrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserPrivilegesResponse), nil
	}
}

// ShowUserPrivilegesInvoker 查询用户在项目下的权限
func (c *CodeArtsArtifactClient) ShowUserPrivilegesInvoker(request *model.ShowUserPrivilegesRequest) *ShowUserPrivilegesInvoker {
	requestDef := GenReqDefForShowUserPrivileges()
	return &ShowUserPrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserTicket 查询用户凭证
//
// 查询用户凭证，该凭证为IDC用户下载制品时使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) ShowUserTicket(request *model.ShowUserTicketRequest) (*model.ShowUserTicketResponse, error) {
	requestDef := GenReqDefForShowUserTicket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserTicketResponse), nil
	}
}

// ShowUserTicketInvoker 查询用户凭证
func (c *CodeArtsArtifactClient) ShowUserTicketInvoker(request *model.ShowUserTicketRequest) *ShowUserTicketInvoker {
	requestDef := GenReqDefForShowUserTicket()
	return &ShowUserTicketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateArtifactory 编辑非maven仓库信息
//
// 编辑非maven仓库信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsArtifactClient) UpdateArtifactory(request *model.UpdateArtifactoryRequest) (*model.UpdateArtifactoryResponse, error) {
	requestDef := GenReqDefForUpdateArtifactory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateArtifactoryResponse), nil
	}
}

// UpdateArtifactoryInvoker 编辑非maven仓库信息
func (c *CodeArtsArtifactClient) UpdateArtifactoryInvoker(request *model.UpdateArtifactoryRequest) *UpdateArtifactoryInvoker {
	requestDef := GenReqDefForUpdateArtifactory()
	return &UpdateArtifactoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
