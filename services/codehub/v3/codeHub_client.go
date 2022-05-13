package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codehub/v3/model"
)

type CodeHubClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeHubClient(hcClient *http_client.HcHttpClient) *CodeHubClient {
	return &CodeHubClient{HcClient: hcClient}
}

func CodeHubClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 创建提交
//
// 能够一次提交位于不同目录的多个文件，目录不存在时，能自动创建目录。支持强制覆盖选项，当选择强制覆盖标志为true时，忽略冲突，强制提交。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) CreateCommit(request *model.CreateCommitRequest) (*model.CreateCommitResponse, error) {
	requestDef := GenReqDefForCreateCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommitResponse), nil
	}
}

// 查询某个仓库的提交信息
//
// 根据仓库短ID获取提交信息，支持根据文件路径，查询这个路径下所有的commits列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListCommits(request *model.ListCommitsRequest) (*model.ListCommitsResponse, error) {
	requestDef := GenReqDefForListCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitsResponse), nil
	}
}

// 查询某个仓库的提交差异信息
//
// 根据commit id查询提交差异信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowDiffCommit(request *model.ShowDiffCommitRequest) (*model.ShowDiffCommitResponse, error) {
	requestDef := GenReqDefForShowDiffCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiffCommitResponse), nil
	}
}

// 查询某个仓库的特定提交信息
//
// 获取由commit id或分支或标记的名称标识的特定提交。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowSingleCommit(request *model.ShowSingleCommitRequest) (*model.ShowSingleCommitResponse, error) {
	requestDef := GenReqDefForShowSingleCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleCommitResponse), nil
	}
}

// 查询某个仓库的文件信息
//
// 获取仓库中文件的信息，如名称、大小、内容。请注意，文件内容是Base64编码的。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowFile(request *model.ShowFileRequest) (*model.ShowFileResponse, error) {
	requestDef := GenReqDefForShowFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileResponse), nil
	}
}

// 获取项目下所有仓库信息
//
// 获取仓库列表 模糊查询支持范围,如果未传入project_id，则支持按仓库名或项目名模糊查询，否则，只按仓库名模糊匹配。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) GetAllRepositoryByProjectId(request *model.GetAllRepositoryByProjectIdRequest) (*model.GetAllRepositoryByProjectIdResponse, error) {
	requestDef := GenReqDefForGetAllRepositoryByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAllRepositoryByProjectIdResponse), nil
	}
}

// 获取一个项目下可以设置为公开状态的仓库列表
//
// 获取一个项目下可以设置为公开状态的仓库列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) GetProductTemplates(request *model.GetProductTemplatesRequest) (*model.GetProductTemplatesResponse, error) {
	requestDef := GenReqDefForGetProductTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetProductTemplatesResponse), nil
	}
}

// 获取一个项目下可以设置为公开状态的仓库列表
//
// 获取一个项目下可以设置为公开状态的仓库列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListProductTwoTemplates(request *model.ListProductTwoTemplatesRequest) (*model.ListProductTwoTemplatesResponse, error) {
	requestDef := GenReqDefForListProductTwoTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTwoTemplatesResponse), nil
	}
}

// 校验指定项目下的仓库名
//
// 一般创建仓库时调用。通过传入项目uuid,仓库名，调用CoudeHubAdapter接口，查询数据库来判断仓库是否重名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryNameExist(request *model.ShowRepositoryNameExistRequest) (*model.ShowRepositoryNameExistResponse, error) {
	requestDef := GenReqDefForShowRepositoryNameExist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryNameExistResponse), nil
	}
}

// 添加仓库成员
//
// 调用方codehubportal,添加仓库成员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddRepoMembers(request *model.AddRepoMembersRequest) (*model.AddRepoMembersResponse, error) {
	requestDef := GenReqDefForAddRepoMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepoMembersResponse), nil
	}
}

// 删除仓库成员
//
// 删除仓库成员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteRepoMember(request *model.DeleteRepoMemberRequest) (*model.DeleteRepoMemberResponse, error) {
	requestDef := GenReqDefForDeleteRepoMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoMemberResponse), nil
	}
}

// 获取仓库所有成员记录
//
// 获取仓库成员列表,可通过关键字搜索某成员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListRepoMembers(request *model.ListRepoMembersRequest) (*model.ListRepoMembersResponse, error) {
	requestDef := GenReqDefForListRepoMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoMembersResponse), nil
	}
}

// 设置成员在仓库中的角色
//
// 给仓库中成员设置仓库的操作权限，
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) SetRepoRole(request *model.SetRepoRoleRequest) (*model.SetRepoRoleResponse, error) {
	requestDef := GenReqDefForSetRepoRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRepoRoleResponse), nil
	}
}

// 添加部署密钥
//
// 添加部署密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddDeployKey(request *model.AddDeployKeyRequest) (*model.AddDeployKeyResponse, error) {
	requestDef := GenReqDefForAddDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeployKeyResponse), nil
	}
}

// 添加部署密钥
//
// 添加部署密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddDeployKeyV2(request *model.AddDeployKeyV2Request) (*model.AddDeployKeyV2Response, error) {
	requestDef := GenReqDefForAddDeployKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeployKeyV2Response), nil
	}
}

// 新建保护分支
//
// 新建保护分支
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddProtectBranchV2(request *model.AddProtectBranchV2Request) (*model.AddProtectBranchV2Response, error) {
	requestDef := GenReqDefForAddProtectBranchV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProtectBranchV2Response), nil
	}
}

// 新建标签
//
// 新建标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddTagV2(request *model.AddTagV2Request) (*model.AddTagV2Response, error) {
	requestDef := GenReqDefForAddTagV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTagV2Response), nil
	}
}

// 创建仓库
//
// 用指定的名称在指定项目上创建仓库。传入参数：仓库名、模板id、是否导入项目成员、归属项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) CreateRepository(request *model.CreateRepositoryRequest) (*model.CreateRepositoryResponse, error) {
	requestDef := GenReqDefForCreateRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositoryResponse), nil
	}
}

// 删除仓库部署密钥
//
// 删除仓库部署密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteDeployKey(request *model.DeleteDeployKeyRequest) (*model.DeleteDeployKeyResponse, error) {
	requestDef := GenReqDefForDeleteDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployKeyResponse), nil
	}
}

// 删除仓库部署密钥
//
// 删除仓库部署密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteDeployKeyV2(request *model.DeleteDeployKeyV2Request) (*model.DeleteDeployKeyV2Response, error) {
	requestDef := GenReqDefForDeleteDeployKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployKeyV2Response), nil
	}
}

// 删除仓库
//
// 根据仓库32位uuid删除指定的仓库
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteRepository(request *model.DeleteRepositoryRequest) (*model.DeleteRepositoryResponse, error) {
	requestDef := GenReqDefForDeleteRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepositoryResponse), nil
	}
}

// 查询项目下的某个仓库
//
// 不建议再使用,建议使用/{repository_uuid}/status
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) GetRepositoryByProjectId(request *model.GetRepositoryByProjectIdRequest) (*model.GetRepositoryByProjectIdResponse, error) {
	requestDef := GenReqDefForGetRepositoryByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetRepositoryByProjectIdResponse), nil
	}
}

// 获取公开示例模板列表
//
// 获取公开示例模板列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) GetTemplates(request *model.GetTemplatesRequest) (*model.GetTemplatesResponse, error) {
	requestDef := GenReqDefForGetTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetTemplatesResponse), nil
	}
}

// 获取仓库上一次的提交统计信息
//
// 获取仓库上一次的提交统计信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListCommitStatistics(request *model.ListCommitStatisticsRequest) (*model.ListCommitStatisticsResponse, error) {
	requestDef := GenReqDefForListCommitStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitStatisticsResponse), nil
	}
}

// 获取一个仓库下特定分支指定文件内容
//
// 获取一个仓库下特定分支指定文件内容
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListFiles(request *model.ListFilesRequest) (*model.ListFilesResponse, error) {
	requestDef := GenReqDefForListFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFilesResponse), nil
	}
}

// 查看仓库的创建状态
//
// 获取仓库状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryStatus(request *model.ListRepositoryStatusRequest) (*model.ListRepositoryStatusResponse, error) {
	requestDef := GenReqDefForListRepositoryStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryStatusResponse), nil
	}
}

// 获取分支目录下的文件
//
// 获取分支目录下的文件
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListSubfiles(request *model.ListSubfilesRequest) (*model.ListSubfilesResponse, error) {
	requestDef := GenReqDefForListSubfiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubfilesResponse), nil
	}
}

// 设置仓库是公开状态还是私有状态
//
// 设置仓库是公开状态还是私有状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListTemplatesTwo(request *model.ListTemplatesTwoRequest) (*model.ListTemplatesTwoResponse, error) {
	requestDef := GenReqDefForListTemplatesTwo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesTwoResponse), nil
	}
}

// 获取公开示例模板列表
//
// 获取公开示例模板列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListTwoTemplates(request *model.ListTwoTemplatesRequest) (*model.ListTwoTemplatesResponse, error) {
	requestDef := GenReqDefForListTwoTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTwoTemplatesResponse), nil
	}
}

// 设置仓库是公开状态还是私有状态
//
// 设置仓库是公开状态还是私有状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShareTemplates(request *model.ShareTemplatesRequest) (*model.ShareTemplatesResponse, error) {
	requestDef := GenReqDefForShareTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShareTemplatesResponse), nil
	}
}

// 查询某仓库对应的分支
//
// 根据仓库id获取指定仓库的分支列表.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowBranchesByRepositoryId(request *model.ShowBranchesByRepositoryIdRequest) (*model.ShowBranchesByRepositoryIdResponse, error) {
	requestDef := GenReqDefForShowBranchesByRepositoryId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchesByRepositoryIdResponse), nil
	}
}

// 查询某仓库的标签列表
//
// 查询指定仓库对应的分支。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowBranchesByTwoRepositoryId(request *model.ShowBranchesByTwoRepositoryIdRequest) (*model.ShowBranchesByTwoRepositoryIdResponse, error) {
	requestDef := GenReqDefForShowBranchesByTwoRepositoryId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchesByTwoRepositoryIdResponse), nil
	}
}

// 根据组名和仓库名查询某仓库某分支对应的提交
//
// 根据仓库组名、仓库名和分支获取提交列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitsByBranch(request *model.ShowCommitsByBranchRequest) (*model.ShowCommitsByBranchResponse, error) {
	requestDef := GenReqDefForShowCommitsByBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitsByBranchResponse), nil
	}
}

// 根据仓库id查询仓库某分支对应的提交，提供更多可选参数
//
// 根据仓库id查询仓库某分支对应的提交.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitsByRepoId(request *model.ShowCommitsByRepoIdRequest) (*model.ShowCommitsByRepoIdResponse, error) {
	requestDef := GenReqDefForShowCommitsByRepoId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitsByRepoIdResponse), nil
	}
}

// 修改被流水线引用的仓库状态
//
// 修改被流水线引用的仓库状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowHasPipeline(request *model.ShowHasPipelineRequest) (*model.ShowHasPipelineResponse, error) {
	requestDef := GenReqDefForShowHasPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHasPipelineResponse), nil
	}
}

// 获取一个仓库下特定分支的图片文件
//
// 获取一个仓库下特定分支的图片文件
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowImageBlob(request *model.ShowImageBlobRequest) (*model.ShowImageBlobResponse, error) {
	requestDef := GenReqDefForShowImageBlob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageBlobResponse), nil
	}
}

// 判断用户是否有仓库的管理员权限
//
// 判断用户是否有仓库的管理员权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowMaster(request *model.ShowMasterRequest) (*model.ShowMasterResponse, error) {
	requestDef := GenReqDefForShowMaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterResponse), nil
	}
}

// 根据仓库名组名获取仓库短id，用以拼接与commitid对应提交详情页面url
//
// 获取仓库短id,用于获取仓库详情页面url
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowRepoId(request *model.ShowRepoIdRequest) (*model.ShowRepoIdResponse, error) {
	requestDef := GenReqDefForShowRepoId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepoIdResponse), nil
	}
}

// 下载仓库
//
// 按照指定格式下载仓库
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryArchive(request *model.ShowRepositoryArchiveRequest) (*model.ShowRepositoryArchiveResponse, error) {
	requestDef := GenReqDefForShowRepositoryArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryArchiveResponse), nil
	}
}

// 查询某个仓库的详细信息
//
// 根据仓库UUID获取仓库信息仓库信息。返回 包含id，name，组名，仓库访问URL。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryByUuid(request *model.ShowRepositoryByUuidRequest) (*model.ShowRepositoryByUuidResponse, error) {
	requestDef := GenReqDefForShowRepositoryByUuid()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryByUuidResponse), nil
	}
}

// 仓库统计
//
// 根据仓库短id，查询仓库的代码提交记录统计
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryStatistics(request *model.ShowRepositoryStatisticsRequest) (*model.ShowRepositoryStatisticsResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsResponse), nil
	}
}

// 获取代码提交行数
//
// 获取指定日期内代码仓指定分支的代码提交行数
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticCommit(request *model.ShowStatisticCommitRequest) (*model.ShowStatisticCommitResponse, error) {
	requestDef := GenReqDefForShowStatisticCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticCommitResponse), nil
	}
}

// 获取代码提交行数
//
// 获取指定日期内代码仓指定分支的代码提交行数
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticCommitV3(request *model.ShowStatisticCommitV3Request) (*model.ShowStatisticCommitV3Response, error) {
	requestDef := GenReqDefForShowStatisticCommitV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticCommitV3Response), nil
	}
}

// 获取仓库统计数据
//
// 获取仓库统计数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticalData(request *model.ShowStatisticalDataRequest) (*model.ShowStatisticalDataResponse, error) {
	requestDef := GenReqDefForShowStatisticalData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticalDataResponse), nil
	}
}

// 添加ssh key
//
// 添加ssh key
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddSshKey(request *model.AddSshKeyRequest) (*model.AddSshKeyResponse, error) {
	requestDef := GenReqDefForAddSshKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSshKeyResponse), nil
	}
}

// 删除用户公钥
//
// 调用gitlab原生接口删除用户公钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteSShkey(request *model.DeleteSShkeyRequest) (*model.DeleteSShkeyResponse, error) {
	requestDef := GenReqDefForDeleteSShkey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSShkeyResponse), nil
	}
}

// 获取ssh key列表
//
// 获取ssh key列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListSshKeys(request *model.ListSshKeysRequest) (*model.ListSshKeysResponse, error) {
	requestDef := GenReqDefForListSshKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSshKeysResponse), nil
	}
}

// 检验私钥是否有拉取代码的权限
//
// 检验私钥是否有拉取代码的权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowPrivateKeyVerify(request *model.ShowPrivateKeyVerifyRequest) (*model.ShowPrivateKeyVerifyResponse, error) {
	requestDef := GenReqDefForShowPrivateKeyVerify()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateKeyVerifyResponse), nil
	}
}

//  https账号密码校验
//
// 调用 gitlab 接口判断用户使用 https 上传/下载代码时输入的用户名和密码是否合法。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ValidateHttpsInfo(request *model.ValidateHttpsInfoRequest) (*model.ValidateHttpsInfoResponse, error) {
	requestDef := GenReqDefForValidateHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateHttpsInfoResponse), nil
	}
}

//  https账号密码校验
//
// 调用 gitlab 接口判断用户使用 https 上传/下载代码时输入的用户名和密码是否合法。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ValidateHttpsInfoV2(request *model.ValidateHttpsInfoV2Request) (*model.ValidateHttpsInfoV2Response, error) {
	requestDef := GenReqDefForValidateHttpsInfoV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateHttpsInfoV2Response), nil
	}
}

// 创建项目、仓库
//
// 创建项目后，创建仓库组由后台生成方式 传入参数：仓库名、模板id、是否导入项目成员、归属项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectAndRepositories(request *model.CreateProjectAndRepositoriesRequest) (*model.CreateProjectAndRepositoriesResponse, error) {
	requestDef := GenReqDefForCreateProjectAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectAndRepositoriesResponse), nil
	}
}

// 创建项目并fork仓库
//
// 创建仓库后fork仓库 传入参数：仓库名、是否导入项目成员、归属项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectAndforkRepositories(request *model.CreateProjectAndforkRepositoriesRequest) (*model.CreateProjectAndforkRepositoriesResponse, error) {
	requestDef := GenReqDefForCreateProjectAndforkRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectAndforkRepositoriesResponse), nil
	}
}

// 查询用户的所有仓库
//
// 获取用户的所有仓库信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListUserAllRepositories(request *model.ListUserAllRepositoriesRequest) (*model.ListUserAllRepositoriesResponse, error) {
	requestDef := GenReqDefForListUserAllRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserAllRepositoriesResponse), nil
	}
}

// 查询项目下的所有仓库
//
// 获取仓库列表,模糊查询支持范围,如果未传入project uuid，则支持按仓库名或项目名模糊查询，否则，只按仓库名模糊匹配
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ShowAllRepositoryByTwoProjectId(request *model.ShowAllRepositoryByTwoProjectIdRequest) (*model.ShowAllRepositoryByTwoProjectIdResponse, error) {
	requestDef := GenReqDefForShowAllRepositoryByTwoProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllRepositoryByTwoProjectIdResponse), nil
	}
}

// 为指定仓库添加hook
//
// 提交代码自动触发编译构建，添加仓库钩子
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) AddHooks(request *model.AddHooksRequest) (*model.AddHooksResponse, error) {
	requestDef := GenReqDefForAddHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddHooksResponse), nil
	}
}

// 删除指定仓库的 hook
//
// 提交代码自动触发编译构建，删除仓库钩子
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) DeleteHooks(request *model.DeleteHooksRequest) (*model.DeleteHooksResponse, error) {
	requestDef := GenReqDefForDeleteHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHooksResponse), nil
	}
}

// 查询指定仓库的webhook
//
// 获取仓库webhook
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CodeHubClient) ListHooks(request *model.ListHooksRequest) (*model.ListHooksResponse, error) {
	requestDef := GenReqDefForListHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHooksResponse), nil
	}
}
