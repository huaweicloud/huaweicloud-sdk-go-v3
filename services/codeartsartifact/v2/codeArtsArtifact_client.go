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
