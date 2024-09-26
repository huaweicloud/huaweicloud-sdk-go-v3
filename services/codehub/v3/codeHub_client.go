package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codehub/v3/model"
)

type CodeHubClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeHubClient(hcClient *httpclient.HcHttpClient) *CodeHubClient {
	return &CodeHubClient{HcClient: hcClient}
}

func CodeHubClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateCommit 创建提交
//
// 能够一次提交位于不同目录的多个文件，目录不存在时，能自动创建目录。支持强制覆盖选项，当选择强制覆盖标志为true时，忽略冲突，强制提交。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateCommit(request *model.CreateCommitRequest) (*model.CreateCommitResponse, error) {
	requestDef := GenReqDefForCreateCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommitResponse), nil
	}
}

// CreateCommitInvoker 创建提交
func (c *CodeHubClient) CreateCommitInvoker(request *model.CreateCommitRequest) *CreateCommitInvoker {
	requestDef := GenReqDefForCreateCommit()
	return &CreateCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommits 查询某个仓库的提交信息
//
// 根据仓库短ID获取提交信息，支持根据文件路径，查询这个路径下所有的commits列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListCommits(request *model.ListCommitsRequest) (*model.ListCommitsResponse, error) {
	requestDef := GenReqDefForListCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitsResponse), nil
	}
}

// ListCommitsInvoker 查询某个仓库的提交信息
func (c *CodeHubClient) ListCommitsInvoker(request *model.ListCommitsRequest) *ListCommitsInvoker {
	requestDef := GenReqDefForListCommits()
	return &ListCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiffCommit 查询某个仓库的提交差异信息
//
// 根据commit id查询提交差异信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowDiffCommit(request *model.ShowDiffCommitRequest) (*model.ShowDiffCommitResponse, error) {
	requestDef := GenReqDefForShowDiffCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiffCommitResponse), nil
	}
}

// ShowDiffCommitInvoker 查询某个仓库的提交差异信息
func (c *CodeHubClient) ShowDiffCommitInvoker(request *model.ShowDiffCommitRequest) *ShowDiffCommitInvoker {
	requestDef := GenReqDefForShowDiffCommit()
	return &ShowDiffCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSingleCommit 查询某个仓库的特定提交信息
//
// 获取由commit id或分支或标记的名称标识的特定提交。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowSingleCommit(request *model.ShowSingleCommitRequest) (*model.ShowSingleCommitResponse, error) {
	requestDef := GenReqDefForShowSingleCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleCommitResponse), nil
	}
}

// ShowSingleCommitInvoker 查询某个仓库的特定提交信息
func (c *CodeHubClient) ShowSingleCommitInvoker(request *model.ShowSingleCommitRequest) *ShowSingleCommitInvoker {
	requestDef := GenReqDefForShowSingleCommit()
	return &ShowSingleCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussion 创建MR检视意见
//
// 创建MR检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequestDiscussion(request *model.CreateMergeRequestDiscussionRequest) (*model.CreateMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestDiscussionResponse), nil
	}
}

// CreateMergeRequestDiscussionInvoker 创建MR检视意见
func (c *CodeHubClient) CreateMergeRequestDiscussionInvoker(request *model.CreateMergeRequestDiscussionRequest) *CreateMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussion()
	return &CreateMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussionNote 回复MR检视意见
//
// 回复MR检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequestDiscussionNote(request *model.CreateMergeRequestDiscussionNoteRequest) (*model.CreateMergeRequestDiscussionNoteResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestDiscussionNote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestDiscussionNoteResponse), nil
	}
}

// CreateMergeRequestDiscussionNoteInvoker 回复MR检视意见
func (c *CodeHubClient) CreateMergeRequestDiscussionNoteInvoker(request *model.CreateMergeRequestDiscussionNoteRequest) *CreateMergeRequestDiscussionNoteInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussionNote()
	return &CreateMergeRequestDiscussionNoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReviewSetting 获取检视意见设置
//
// 获取检视意见设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowReviewSetting(request *model.ShowReviewSettingRequest) (*model.ShowReviewSettingResponse, error) {
	requestDef := GenReqDefForShowReviewSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReviewSettingResponse), nil
	}
}

// ShowReviewSettingInvoker 获取检视意见设置
func (c *CodeHubClient) ShowReviewSettingInvoker(request *model.ShowReviewSettingRequest) *ShowReviewSettingInvoker {
	requestDef := GenReqDefForShowReviewSetting()
	return &ShowReviewSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFilesByQuery 查询某个仓库的文件信息
//
// 获取仓库中文件的信息，如名称、大小、内容。请注意，文件内容是Base64编码的。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListFilesByQuery(request *model.ListFilesByQueryRequest) (*model.ListFilesByQueryResponse, error) {
	requestDef := GenReqDefForListFilesByQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFilesByQueryResponse), nil
	}
}

// ListFilesByQueryInvoker 查询某个仓库的文件信息
func (c *CodeHubClient) ListFilesByQueryInvoker(request *model.ListFilesByQueryRequest) *ListFilesByQueryInvoker {
	requestDef := GenReqDefForListFilesByQuery()
	return &ListFilesByQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowFile 查询某个仓库的文件信息
//
// 获取仓库中文件的信息，如名称、大小、内容。请注意，文件内容是Base64编码的。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowFile(request *model.ShowFileRequest) (*model.ShowFileResponse, error) {
	requestDef := GenReqDefForShowFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowFileInvoker 查询某个仓库的文件信息
func (c *CodeHubClient) ShowFileInvoker(request *model.ShowFileRequest) *ShowFileInvoker {
	requestDef := GenReqDefForShowFile()
	return &ShowFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetAllRepositoryByProjectId 获取项目下所有仓库信息
//
// 获取仓库列表 模糊查询支持范围,如果未传入project_id，则支持按仓库名或项目名模糊查询，否则，只按仓库名模糊匹配。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) GetAllRepositoryByProjectId(request *model.GetAllRepositoryByProjectIdRequest) (*model.GetAllRepositoryByProjectIdResponse, error) {
	requestDef := GenReqDefForGetAllRepositoryByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAllRepositoryByProjectIdResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetAllRepositoryByProjectIdInvoker 获取项目下所有仓库信息
func (c *CodeHubClient) GetAllRepositoryByProjectIdInvoker(request *model.GetAllRepositoryByProjectIdRequest) *GetAllRepositoryByProjectIdInvoker {
	requestDef := GenReqDefForGetAllRepositoryByProjectId()
	return &GetAllRepositoryByProjectIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetProductTemplates 获取一个项目下可以设置为公开状态的仓库列表
//
// 获取一个项目下可以设置为公开状态的仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) GetProductTemplates(request *model.GetProductTemplatesRequest) (*model.GetProductTemplatesResponse, error) {
	requestDef := GenReqDefForGetProductTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetProductTemplatesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetProductTemplatesInvoker 获取一个项目下可以设置为公开状态的仓库列表
func (c *CodeHubClient) GetProductTemplatesInvoker(request *model.GetProductTemplatesRequest) *GetProductTemplatesInvoker {
	requestDef := GenReqDefForGetProductTemplates()
	return &GetProductTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductTwoTemplates 获取一个项目下可以设置为公开状态的仓库列表
//
// 获取一个项目下可以设置为公开状态的仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProductTwoTemplates(request *model.ListProductTwoTemplatesRequest) (*model.ListProductTwoTemplatesResponse, error) {
	requestDef := GenReqDefForListProductTwoTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTwoTemplatesResponse), nil
	}
}

// ListProductTwoTemplatesInvoker 获取一个项目下可以设置为公开状态的仓库列表
func (c *CodeHubClient) ListProductTwoTemplatesInvoker(request *model.ListProductTwoTemplatesRequest) *ListProductTwoTemplatesInvoker {
	requestDef := GenReqDefForListProductTwoTemplates()
	return &ListProductTwoTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryNameExist 校验指定项目下的仓库名
//
// 一般创建仓库时调用。通过传入项目ID，获取方式请参见[获取项目ID](codehub_api_0014.xml)。,仓库名，来判断仓库是否重名。仓库存在result:false,仓库不存在result:true。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryNameExist(request *model.ShowRepositoryNameExistRequest) (*model.ShowRepositoryNameExistResponse, error) {
	requestDef := GenReqDefForShowRepositoryNameExist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryNameExistResponse), nil
	}
}

// ShowRepositoryNameExistInvoker 校验指定项目下的仓库名
func (c *CodeHubClient) ShowRepositoryNameExistInvoker(request *model.ShowRepositoryNameExistRequest) *ShowRepositoryNameExistInvoker {
	requestDef := GenReqDefForShowRepositoryNameExist()
	return &ShowRepositoryNameExistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRepoMembers 添加仓库成员
//
// 添加仓库成员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddRepoMembers(request *model.AddRepoMembersRequest) (*model.AddRepoMembersResponse, error) {
	requestDef := GenReqDefForAddRepoMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepoMembersResponse), nil
	}
}

// AddRepoMembersInvoker 添加仓库成员
func (c *CodeHubClient) AddRepoMembersInvoker(request *model.AddRepoMembersRequest) *AddRepoMembersInvoker {
	requestDef := GenReqDefForAddRepoMembers()
	return &AddRepoMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepoMember 删除仓库成员
//
// 删除仓库成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteRepoMember(request *model.DeleteRepoMemberRequest) (*model.DeleteRepoMemberResponse, error) {
	requestDef := GenReqDefForDeleteRepoMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoMemberResponse), nil
	}
}

// DeleteRepoMemberInvoker 删除仓库成员
func (c *CodeHubClient) DeleteRepoMemberInvoker(request *model.DeleteRepoMemberRequest) *DeleteRepoMemberInvoker {
	requestDef := GenReqDefForDeleteRepoMember()
	return &DeleteRepoMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepoMembers 获取仓库所有成员记录
//
// 获取仓库成员列表,可通过关键字搜索某成员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepoMembers(request *model.ListRepoMembersRequest) (*model.ListRepoMembersResponse, error) {
	requestDef := GenReqDefForListRepoMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoMembersResponse), nil
	}
}

// ListRepoMembersInvoker 获取仓库所有成员记录
func (c *CodeHubClient) ListRepoMembersInvoker(request *model.ListRepoMembersRequest) *ListRepoMembersInvoker {
	requestDef := GenReqDefForListRepoMembers()
	return &ListRepoMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRepoRole 设置成员在仓库中的角色
//
// 给仓库中成员设置仓库的操作权限，
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) SetRepoRole(request *model.SetRepoRoleRequest) (*model.SetRepoRoleResponse, error) {
	requestDef := GenReqDefForSetRepoRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRepoRoleResponse), nil
	}
}

// SetRepoRoleInvoker 设置成员在仓库中的角色
func (c *CodeHubClient) SetRepoRoleInvoker(request *model.SetRepoRoleRequest) *SetRepoRoleInvoker {
	requestDef := GenReqDefForSetRepoRole()
	return &SetRepoRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddDeployKey 添加部署密钥
//
// 添加部署密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddDeployKey(request *model.AddDeployKeyRequest) (*model.AddDeployKeyResponse, error) {
	requestDef := GenReqDefForAddDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeployKeyResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddDeployKeyInvoker 添加部署密钥
func (c *CodeHubClient) AddDeployKeyInvoker(request *model.AddDeployKeyRequest) *AddDeployKeyInvoker {
	requestDef := GenReqDefForAddDeployKey()
	return &AddDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDeployKeyV2 添加部署密钥
//
// 添加部署密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddDeployKeyV2(request *model.AddDeployKeyV2Request) (*model.AddDeployKeyV2Response, error) {
	requestDef := GenReqDefForAddDeployKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeployKeyV2Response), nil
	}
}

// AddDeployKeyV2Invoker 添加部署密钥
func (c *CodeHubClient) AddDeployKeyV2Invoker(request *model.AddDeployKeyV2Request) *AddDeployKeyV2Invoker {
	requestDef := GenReqDefForAddDeployKeyV2()
	return &AddDeployKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProtectBranchV2 新建保护分支
//
// 新建保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddProtectBranchV2(request *model.AddProtectBranchV2Request) (*model.AddProtectBranchV2Response, error) {
	requestDef := GenReqDefForAddProtectBranchV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProtectBranchV2Response), nil
	}
}

// AddProtectBranchV2Invoker 新建保护分支
func (c *CodeHubClient) AddProtectBranchV2Invoker(request *model.AddProtectBranchV2Request) *AddProtectBranchV2Invoker {
	requestDef := GenReqDefForAddProtectBranchV2()
	return &AddProtectBranchV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTagV2 新建标签
//
// 新建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddTagV2(request *model.AddTagV2Request) (*model.AddTagV2Response, error) {
	requestDef := GenReqDefForAddTagV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTagV2Response), nil
	}
}

// AddTagV2Invoker 新建标签
func (c *CodeHubClient) AddTagV2Invoker(request *model.AddTagV2Request) *AddTagV2Invoker {
	requestDef := GenReqDefForAddTagV2()
	return &AddTagV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepository 创建仓库
//
// 用指定的名称在指定项目上创建仓库。传入参数：仓库名、模板id、是否导入项目成员、归属项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateRepository(request *model.CreateRepositoryRequest) (*model.CreateRepositoryResponse, error) {
	requestDef := GenReqDefForCreateRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositoryResponse), nil
	}
}

// CreateRepositoryInvoker 创建仓库
func (c *CodeHubClient) CreateRepositoryInvoker(request *model.CreateRepositoryRequest) *CreateRepositoryInvoker {
	requestDef := GenReqDefForCreateRepository()
	return &CreateRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteDeployKey 删除仓库部署密钥
//
// 删除仓库部署密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteDeployKey(request *model.DeleteDeployKeyRequest) (*model.DeleteDeployKeyResponse, error) {
	requestDef := GenReqDefForDeleteDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployKeyResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteDeployKeyInvoker 删除仓库部署密钥
func (c *CodeHubClient) DeleteDeployKeyInvoker(request *model.DeleteDeployKeyRequest) *DeleteDeployKeyInvoker {
	requestDef := GenReqDefForDeleteDeployKey()
	return &DeleteDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployKeyV2 删除仓库部署密钥
//
// 删除仓库部署密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteDeployKeyV2(request *model.DeleteDeployKeyV2Request) (*model.DeleteDeployKeyV2Response, error) {
	requestDef := GenReqDefForDeleteDeployKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployKeyV2Response), nil
	}
}

// DeleteDeployKeyV2Invoker 删除仓库部署密钥
func (c *CodeHubClient) DeleteDeployKeyV2Invoker(request *model.DeleteDeployKeyV2Request) *DeleteDeployKeyV2Invoker {
	requestDef := GenReqDefForDeleteDeployKeyV2()
	return &DeleteDeployKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepository 删除仓库
//
// 根据仓库32位uuid删除指定的仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteRepository(request *model.DeleteRepositoryRequest) (*model.DeleteRepositoryResponse, error) {
	requestDef := GenReqDefForDeleteRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepositoryResponse), nil
	}
}

// DeleteRepositoryInvoker 删除仓库
func (c *CodeHubClient) DeleteRepositoryInvoker(request *model.DeleteRepositoryRequest) *DeleteRepositoryInvoker {
	requestDef := GenReqDefForDeleteRepository()
	return &DeleteRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetRepositoryByProjectId 查询项目下的某个仓库
//
// 不建议再使用,建议使用/{repository_uuid}/status
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) GetRepositoryByProjectId(request *model.GetRepositoryByProjectIdRequest) (*model.GetRepositoryByProjectIdResponse, error) {
	requestDef := GenReqDefForGetRepositoryByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetRepositoryByProjectIdResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetRepositoryByProjectIdInvoker 查询项目下的某个仓库
func (c *CodeHubClient) GetRepositoryByProjectIdInvoker(request *model.GetRepositoryByProjectIdRequest) *GetRepositoryByProjectIdInvoker {
	requestDef := GenReqDefForGetRepositoryByProjectId()
	return &GetRepositoryByProjectIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetTemplates 获取公开示例模板列表
//
// 获取公开示例模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) GetTemplates(request *model.GetTemplatesRequest) (*model.GetTemplatesResponse, error) {
	requestDef := GenReqDefForGetTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetTemplatesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// GetTemplatesInvoker 获取公开示例模板列表
func (c *CodeHubClient) GetTemplatesInvoker(request *model.GetTemplatesRequest) *GetTemplatesInvoker {
	requestDef := GenReqDefForGetTemplates()
	return &GetTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranchesByRepositoryId 获取仓库分支列表
//
// 获取仓库分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListBranchesByRepositoryId(request *model.ListBranchesByRepositoryIdRequest) (*model.ListBranchesByRepositoryIdResponse, error) {
	requestDef := GenReqDefForListBranchesByRepositoryId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchesByRepositoryIdResponse), nil
	}
}

// ListBranchesByRepositoryIdInvoker 获取仓库分支列表
func (c *CodeHubClient) ListBranchesByRepositoryIdInvoker(request *model.ListBranchesByRepositoryIdRequest) *ListBranchesByRepositoryIdInvoker {
	requestDef := GenReqDefForListBranchesByRepositoryId()
	return &ListBranchesByRepositoryIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitStatistics 获取仓库上一次的提交统计信息
//
// 获取仓库上一次的提交统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListCommitStatistics(request *model.ListCommitStatisticsRequest) (*model.ListCommitStatisticsResponse, error) {
	requestDef := GenReqDefForListCommitStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitStatisticsResponse), nil
	}
}

// ListCommitStatisticsInvoker 获取仓库上一次的提交统计信息
func (c *CodeHubClient) ListCommitStatisticsInvoker(request *model.ListCommitStatisticsRequest) *ListCommitStatisticsInvoker {
	requestDef := GenReqDefForListCommitStatistics()
	return &ListCommitStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFiles 获取一个仓库下特定分支指定文件内容
//
// 获取一个仓库下特定分支指定文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListFiles(request *model.ListFilesRequest) (*model.ListFilesResponse, error) {
	requestDef := GenReqDefForListFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFilesResponse), nil
	}
}

// ListFilesInvoker 获取一个仓库下特定分支指定文件内容
func (c *CodeHubClient) ListFilesInvoker(request *model.ListFilesRequest) *ListFilesInvoker {
	requestDef := GenReqDefForListFiles()
	return &ListFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeChanges 获取变更文件
//
// 获取变更文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeChanges(request *model.ListMergeChangesRequest) (*model.ListMergeChangesResponse, error) {
	requestDef := GenReqDefForListMergeChanges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeChangesResponse), nil
	}
}

// ListMergeChangesInvoker 获取变更文件
func (c *CodeHubClient) ListMergeChangesInvoker(request *model.ListMergeChangesRequest) *ListMergeChangesInvoker {
	requestDef := GenReqDefForListMergeChanges()
	return &ListMergeChangesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeChangesTrees 获取变更文件列表
//
// 获取变更文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeChangesTrees(request *model.ListMergeChangesTreesRequest) (*model.ListMergeChangesTreesResponse, error) {
	requestDef := GenReqDefForListMergeChangesTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeChangesTreesResponse), nil
	}
}

// ListMergeChangesTreesInvoker 获取变更文件列表
func (c *CodeHubClient) ListMergeChangesTreesInvoker(request *model.ListMergeChangesTreesRequest) *ListMergeChangesTreesInvoker {
	requestDef := GenReqDefForListMergeChangesTrees()
	return &ListMergeChangesTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequest 获取仓库合并请求列表
//
// 获取仓库合并请求列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequest(request *model.ListMergeRequestRequest) (*model.ListMergeRequestResponse, error) {
	requestDef := GenReqDefForListMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestResponse), nil
	}
}

// ListMergeRequestInvoker 获取仓库合并请求列表
func (c *CodeHubClient) ListMergeRequestInvoker(request *model.ListMergeRequestRequest) *ListMergeRequestInvoker {
	requestDef := GenReqDefForListMergeRequest()
	return &ListMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestReviewers 根据仓库短ID和合并请求短ID获取检视人信息
//
// 根据仓库短ID和合并请求短ID获取检视人信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestReviewers(request *model.ListMergeRequestReviewersRequest) (*model.ListMergeRequestReviewersResponse, error) {
	requestDef := GenReqDefForListMergeRequestReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestReviewersResponse), nil
	}
}

// ListMergeRequestReviewersInvoker 根据仓库短ID和合并请求短ID获取检视人信息
func (c *CodeHubClient) ListMergeRequestReviewersInvoker(request *model.ListMergeRequestReviewersRequest) *ListMergeRequestReviewersInvoker {
	requestDef := GenReqDefForListMergeRequestReviewers()
	return &ListMergeRequestReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRelatedCommits 获取关联工作项信息
//
// 获取关联工作项信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRelatedCommits(request *model.ListRelatedCommitsRequest) (*model.ListRelatedCommitsResponse, error) {
	requestDef := GenReqDefForListRelatedCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelatedCommitsResponse), nil
	}
}

// ListRelatedCommitsInvoker 获取关联工作项信息
func (c *CodeHubClient) ListRelatedCommitsInvoker(request *model.ListRelatedCommitsRequest) *ListRelatedCommitsInvoker {
	requestDef := GenReqDefForListRelatedCommits()
	return &ListRelatedCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryStatus 查看仓库的创建状态
//
// 获取仓库状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryStatus(request *model.ListRepositoryStatusRequest) (*model.ListRepositoryStatusResponse, error) {
	requestDef := GenReqDefForListRepositoryStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryStatusResponse), nil
	}
}

// ListRepositoryStatusInvoker 查看仓库的创建状态
func (c *CodeHubClient) ListRepositoryStatusInvoker(request *model.ListRepositoryStatusRequest) *ListRepositoryStatusInvoker {
	requestDef := GenReqDefForListRepositoryStatus()
	return &ListRepositoryStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubfiles 获取分支目录下的文件
//
// 获取分支目录下的文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListSubfiles(request *model.ListSubfilesRequest) (*model.ListSubfilesResponse, error) {
	requestDef := GenReqDefForListSubfiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubfilesResponse), nil
	}
}

// ListSubfilesInvoker 获取分支目录下的文件
func (c *CodeHubClient) ListSubfilesInvoker(request *model.ListSubfilesRequest) *ListSubfilesInvoker {
	requestDef := GenReqDefForListSubfiles()
	return &ListSubfilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplatesTwo 设置仓库是公开状态还是私有状态
//
// 设置仓库是公开状态还是私有状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTemplatesTwo(request *model.ListTemplatesTwoRequest) (*model.ListTemplatesTwoResponse, error) {
	requestDef := GenReqDefForListTemplatesTwo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesTwoResponse), nil
	}
}

// ListTemplatesTwoInvoker 设置仓库是公开状态还是私有状态
func (c *CodeHubClient) ListTemplatesTwoInvoker(request *model.ListTemplatesTwoRequest) *ListTemplatesTwoInvoker {
	requestDef := GenReqDefForListTemplatesTwo()
	return &ListTemplatesTwoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTwoTemplates 获取公开示例模板列表
//
// 获取公开示例模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTwoTemplates(request *model.ListTwoTemplatesRequest) (*model.ListTwoTemplatesResponse, error) {
	requestDef := GenReqDefForListTwoTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTwoTemplatesResponse), nil
	}
}

// ListTwoTemplatesInvoker 获取公开示例模板列表
func (c *CodeHubClient) ListTwoTemplatesInvoker(request *model.ListTwoTemplatesRequest) *ListTwoTemplatesInvoker {
	requestDef := GenReqDefForListTwoTemplates()
	return &ListTwoTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShareTemplates 设置仓库是公开状态还是私有状态
//
// 设置仓库是公开状态还是私有状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShareTemplates(request *model.ShareTemplatesRequest) (*model.ShareTemplatesResponse, error) {
	requestDef := GenReqDefForShareTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShareTemplatesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShareTemplatesInvoker 设置仓库是公开状态还是私有状态
func (c *CodeHubClient) ShareTemplatesInvoker(request *model.ShareTemplatesRequest) *ShareTemplatesInvoker {
	requestDef := GenReqDefForShareTemplates()
	return &ShareTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranchesByRepositoryId 查询某仓库对应的分支
//
// 根据仓库id获取指定仓库的分支列表.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowBranchesByRepositoryId(request *model.ShowBranchesByRepositoryIdRequest) (*model.ShowBranchesByRepositoryIdResponse, error) {
	requestDef := GenReqDefForShowBranchesByRepositoryId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchesByRepositoryIdResponse), nil
	}
}

// ShowBranchesByRepositoryIdInvoker 查询某仓库对应的分支
func (c *CodeHubClient) ShowBranchesByRepositoryIdInvoker(request *model.ShowBranchesByRepositoryIdRequest) *ShowBranchesByRepositoryIdInvoker {
	requestDef := GenReqDefForShowBranchesByRepositoryId()
	return &ShowBranchesByRepositoryIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranchesByTwoRepositoryId 查询某仓库的标签列表
//
// 查询指定仓库对应的分支。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowBranchesByTwoRepositoryId(request *model.ShowBranchesByTwoRepositoryIdRequest) (*model.ShowBranchesByTwoRepositoryIdResponse, error) {
	requestDef := GenReqDefForShowBranchesByTwoRepositoryId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchesByTwoRepositoryIdResponse), nil
	}
}

// ShowBranchesByTwoRepositoryIdInvoker 查询某仓库的标签列表
func (c *CodeHubClient) ShowBranchesByTwoRepositoryIdInvoker(request *model.ShowBranchesByTwoRepositoryIdRequest) *ShowBranchesByTwoRepositoryIdInvoker {
	requestDef := GenReqDefForShowBranchesByTwoRepositoryId()
	return &ShowBranchesByTwoRepositoryIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitsByBranch 根据组名和仓库名查询某仓库某分支对应的提交
//
// 根据仓库组名、仓库名和分支获取提交列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitsByBranch(request *model.ShowCommitsByBranchRequest) (*model.ShowCommitsByBranchResponse, error) {
	requestDef := GenReqDefForShowCommitsByBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitsByBranchResponse), nil
	}
}

// ShowCommitsByBranchInvoker 根据组名和仓库名查询某仓库某分支对应的提交
func (c *CodeHubClient) ShowCommitsByBranchInvoker(request *model.ShowCommitsByBranchRequest) *ShowCommitsByBranchInvoker {
	requestDef := GenReqDefForShowCommitsByBranch()
	return &ShowCommitsByBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitsByRepoId 根据仓库id查询仓库某分支对应的提交，提供更多可选参数
//
// 根据仓库id查询仓库某分支对应的提交.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitsByRepoId(request *model.ShowCommitsByRepoIdRequest) (*model.ShowCommitsByRepoIdResponse, error) {
	requestDef := GenReqDefForShowCommitsByRepoId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitsByRepoIdResponse), nil
	}
}

// ShowCommitsByRepoIdInvoker 根据仓库id查询仓库某分支对应的提交，提供更多可选参数
func (c *CodeHubClient) ShowCommitsByRepoIdInvoker(request *model.ShowCommitsByRepoIdRequest) *ShowCommitsByRepoIdInvoker {
	requestDef := GenReqDefForShowCommitsByRepoId()
	return &ShowCommitsByRepoIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHasPipeline 修改被流水线引用的仓库状态
//
// 修改被流水线引用的仓库状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowHasPipeline(request *model.ShowHasPipelineRequest) (*model.ShowHasPipelineResponse, error) {
	requestDef := GenReqDefForShowHasPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHasPipelineResponse), nil
	}
}

// ShowHasPipelineInvoker 修改被流水线引用的仓库状态
func (c *CodeHubClient) ShowHasPipelineInvoker(request *model.ShowHasPipelineRequest) *ShowHasPipelineInvoker {
	requestDef := GenReqDefForShowHasPipeline()
	return &ShowHasPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageBlob 获取一个仓库下特定分支的图片文件
//
// 获取一个仓库下特定分支的图片文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowImageBlob(request *model.ShowImageBlobRequest) (*model.ShowImageBlobResponse, error) {
	requestDef := GenReqDefForShowImageBlob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageBlobResponse), nil
	}
}

// ShowImageBlobInvoker 获取一个仓库下特定分支的图片文件
func (c *CodeHubClient) ShowImageBlobInvoker(request *model.ShowImageBlobRequest) *ShowImageBlobInvoker {
	requestDef := GenReqDefForShowImageBlob()
	return &ShowImageBlobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMaster 判断用户是否有仓库的管理员权限
//
// 判断用户是否有仓库的管理员权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMaster(request *model.ShowMasterRequest) (*model.ShowMasterResponse, error) {
	requestDef := GenReqDefForShowMaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterResponse), nil
	}
}

// ShowMasterInvoker 判断用户是否有仓库的管理员权限
func (c *CodeHubClient) ShowMasterInvoker(request *model.ShowMasterRequest) *ShowMasterInvoker {
	requestDef := GenReqDefForShowMaster()
	return &ShowMasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequest 获取仓库合并请求详情
//
// 获取仓库合并请求详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequest(request *model.ShowMergeRequestRequest) (*model.ShowMergeRequestResponse, error) {
	requestDef := GenReqDefForShowMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestResponse), nil
	}
}

// ShowMergeRequestInvoker 获取仓库合并请求详情
func (c *CodeHubClient) ShowMergeRequestInvoker(request *model.ShowMergeRequestRequest) *ShowMergeRequestInvoker {
	requestDef := GenReqDefForShowMergeRequest()
	return &ShowMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepoId 根据仓库名组名获取仓库短id，用以拼接与commitid对应提交详情页面url
//
// 获取仓库短id,用于获取仓库详情页面url
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepoId(request *model.ShowRepoIdRequest) (*model.ShowRepoIdResponse, error) {
	requestDef := GenReqDefForShowRepoId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepoIdResponse), nil
	}
}

// ShowRepoIdInvoker 根据仓库名组名获取仓库短id，用以拼接与commitid对应提交详情页面url
func (c *CodeHubClient) ShowRepoIdInvoker(request *model.ShowRepoIdRequest) *ShowRepoIdInvoker {
	requestDef := GenReqDefForShowRepoId()
	return &ShowRepoIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryArchive 下载仓库
//
// 按照指定格式下载仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryArchive(request *model.ShowRepositoryArchiveRequest) (*model.ShowRepositoryArchiveResponse, error) {
	requestDef := GenReqDefForShowRepositoryArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryArchiveResponse), nil
	}
}

// ShowRepositoryArchiveInvoker 下载仓库
func (c *CodeHubClient) ShowRepositoryArchiveInvoker(request *model.ShowRepositoryArchiveRequest) *ShowRepositoryArchiveInvoker {
	requestDef := GenReqDefForShowRepositoryArchive()
	return &ShowRepositoryArchiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryByUuid 查询某个仓库的详细信息
//
// 根据仓库UUID(由CreateRepository接口返回)获取仓库信息仓库信息。返回 包含id，name，组名，仓库访问URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryByUuid(request *model.ShowRepositoryByUuidRequest) (*model.ShowRepositoryByUuidResponse, error) {
	requestDef := GenReqDefForShowRepositoryByUuid()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryByUuidResponse), nil
	}
}

// ShowRepositoryByUuidInvoker 查询某个仓库的详细信息
func (c *CodeHubClient) ShowRepositoryByUuidInvoker(request *model.ShowRepositoryByUuidRequest) *ShowRepositoryByUuidInvoker {
	requestDef := GenReqDefForShowRepositoryByUuid()
	return &ShowRepositoryByUuidInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatistics 仓库统计
//
// 根据仓库短id，查询仓库的代码提交记录统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryStatistics(request *model.ShowRepositoryStatisticsRequest) (*model.ShowRepositoryStatisticsResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsResponse), nil
	}
}

// ShowRepositoryStatisticsInvoker 仓库统计
func (c *CodeHubClient) ShowRepositoryStatisticsInvoker(request *model.ShowRepositoryStatisticsRequest) *ShowRepositoryStatisticsInvoker {
	requestDef := GenReqDefForShowRepositoryStatistics()
	return &ShowRepositoryStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowStatisticCommit 获取代码提交行数
//
// 获取指定日期内代码仓指定分支的代码提交行数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticCommit(request *model.ShowStatisticCommitRequest) (*model.ShowStatisticCommitResponse, error) {
	requestDef := GenReqDefForShowStatisticCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticCommitResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowStatisticCommitInvoker 获取代码提交行数
func (c *CodeHubClient) ShowStatisticCommitInvoker(request *model.ShowStatisticCommitRequest) *ShowStatisticCommitInvoker {
	requestDef := GenReqDefForShowStatisticCommit()
	return &ShowStatisticCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStatisticCommitV3 获取代码提交行数
//
// 获取指定日期内代码仓指定分支的代码提交行数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticCommitV3(request *model.ShowStatisticCommitV3Request) (*model.ShowStatisticCommitV3Response, error) {
	requestDef := GenReqDefForShowStatisticCommitV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticCommitV3Response), nil
	}
}

// ShowStatisticCommitV3Invoker 获取代码提交行数
func (c *CodeHubClient) ShowStatisticCommitV3Invoker(request *model.ShowStatisticCommitV3Request) *ShowStatisticCommitV3Invoker {
	requestDef := GenReqDefForShowStatisticCommitV3()
	return &ShowStatisticCommitV3Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStatisticalData 获取仓库统计数据
//
// 获取仓库统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowStatisticalData(request *model.ShowStatisticalDataRequest) (*model.ShowStatisticalDataResponse, error) {
	requestDef := GenReqDefForShowStatisticalData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticalDataResponse), nil
	}
}

// ShowStatisticalDataInvoker 获取仓库统计数据
func (c *CodeHubClient) ShowStatisticalDataInvoker(request *model.ShowStatisticalDataRequest) *ShowStatisticalDataInvoker {
	requestDef := GenReqDefForShowStatisticalData()
	return &ShowStatisticalDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestApprovalState 合并请求代码审核
//
// 合并请求代码审核
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestApprovalState(request *model.UpdateMergeRequestApprovalStateRequest) (*model.UpdateMergeRequestApprovalStateResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestApprovalState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestApprovalStateResponse), nil
	}
}

// UpdateMergeRequestApprovalStateInvoker 合并请求代码审核
func (c *CodeHubClient) UpdateMergeRequestApprovalStateInvoker(request *model.UpdateMergeRequestApprovalStateRequest) *UpdateMergeRequestApprovalStateInvoker {
	requestDef := GenReqDefForUpdateMergeRequestApprovalState()
	return &UpdateMergeRequestApprovalStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSshKey 添加ssh key
//
// 添加ssh key
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddSshKey(request *model.AddSshKeyRequest) (*model.AddSshKeyResponse, error) {
	requestDef := GenReqDefForAddSshKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSshKeyResponse), nil
	}
}

// AddSshKeyInvoker 添加ssh key
func (c *CodeHubClient) AddSshKeyInvoker(request *model.AddSshKeyRequest) *AddSshKeyInvoker {
	requestDef := GenReqDefForAddSshKey()
	return &AddSshKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteSShkey 删除用户公钥
//
// 删除用户公钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteSShkey(request *model.DeleteSShkeyRequest) (*model.DeleteSShkeyResponse, error) {
	requestDef := GenReqDefForDeleteSShkey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSShkeyResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteSShkeyInvoker 删除用户公钥
func (c *CodeHubClient) DeleteSShkeyInvoker(request *model.DeleteSShkeyRequest) *DeleteSShkeyInvoker {
	requestDef := GenReqDefForDeleteSShkey()
	return &DeleteSShkeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSshKeys 获取ssh key列表
//
// 获取ssh key列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListSshKeys(request *model.ListSshKeysRequest) (*model.ListSshKeysResponse, error) {
	requestDef := GenReqDefForListSshKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSshKeysResponse), nil
	}
}

// ListSshKeysInvoker 获取ssh key列表
func (c *CodeHubClient) ListSshKeysInvoker(request *model.ListSshKeysRequest) *ListSshKeysInvoker {
	requestDef := GenReqDefForListSshKeys()
	return &ListSshKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateKeyVerify 检验私钥是否有拉取代码的权限
//
// 检验私钥是否有拉取代码的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowPrivateKeyVerify(request *model.ShowPrivateKeyVerifyRequest) (*model.ShowPrivateKeyVerifyResponse, error) {
	requestDef := GenReqDefForShowPrivateKeyVerify()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateKeyVerifyResponse), nil
	}
}

// ShowPrivateKeyVerifyInvoker 检验私钥是否有拉取代码的权限
func (c *CodeHubClient) ShowPrivateKeyVerifyInvoker(request *model.ShowPrivateKeyVerifyRequest) *ShowPrivateKeyVerifyInvoker {
	requestDef := GenReqDefForShowPrivateKeyVerify()
	return &ShowPrivateKeyVerifyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ValidateHttpsInfo  https账号密码校验
//
// 判断用户使用 https 上传/下载代码时输入的用户名和密码是否合法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ValidateHttpsInfo(request *model.ValidateHttpsInfoRequest) (*model.ValidateHttpsInfoResponse, error) {
	requestDef := GenReqDefForValidateHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateHttpsInfoResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ValidateHttpsInfoInvoker  https账号密码校验
func (c *CodeHubClient) ValidateHttpsInfoInvoker(request *model.ValidateHttpsInfoRequest) *ValidateHttpsInfoInvoker {
	requestDef := GenReqDefForValidateHttpsInfo()
	return &ValidateHttpsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateHttpsInfoV2 https账号密码校验
//
// 判断用户使用 https 上传/下载代码时输入的用户名和密码是否合法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ValidateHttpsInfoV2(request *model.ValidateHttpsInfoV2Request) (*model.ValidateHttpsInfoV2Response, error) {
	requestDef := GenReqDefForValidateHttpsInfoV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateHttpsInfoV2Response), nil
	}
}

// ValidateHttpsInfoV2Invoker https账号密码校验
func (c *CodeHubClient) ValidateHttpsInfoV2Invoker(request *model.ValidateHttpsInfoV2Request) *ValidateHttpsInfoV2Invoker {
	requestDef := GenReqDefForValidateHttpsInfoV2()
	return &ValidateHttpsInfoV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNewBranch 创建分支
//
// 根据仓库id在指定仓库中创建分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateNewBranch(request *model.CreateNewBranchRequest) (*model.CreateNewBranchResponse, error) {
	requestDef := GenReqDefForCreateNewBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNewBranchResponse), nil
	}
}

// CreateNewBranchInvoker 创建分支
func (c *CodeHubClient) CreateNewBranchInvoker(request *model.CreateNewBranchRequest) *CreateNewBranchInvoker {
	requestDef := GenReqDefForCreateNewBranch()
	return &CreateNewBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateIssues 分支关联工作项
//
// 分支关联工作项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateIssues(request *model.AssociateIssuesRequest) (*model.AssociateIssuesResponse, error) {
	requestDef := GenReqDefForAssociateIssues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateIssuesResponse), nil
	}
}

// AssociateIssuesInvoker 分支关联工作项
func (c *CodeHubClient) AssociateIssuesInvoker(request *model.AssociateIssuesRequest) *AssociateIssuesInvoker {
	requestDef := GenReqDefForAssociateIssues()
	return &AssociateIssuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectAndRepositories 创建项目、仓库
//
// 创建项目后，创建仓库组由后台生成方式 传入参数：仓库名、模板id、是否导入项目成员、归属项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectAndRepositories(request *model.CreateProjectAndRepositoriesRequest) (*model.CreateProjectAndRepositoriesResponse, error) {
	requestDef := GenReqDefForCreateProjectAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectAndRepositoriesResponse), nil
	}
}

// CreateProjectAndRepositoriesInvoker 创建项目、仓库
func (c *CodeHubClient) CreateProjectAndRepositoriesInvoker(request *model.CreateProjectAndRepositoriesRequest) *CreateProjectAndRepositoriesInvoker {
	requestDef := GenReqDefForCreateProjectAndRepositories()
	return &CreateProjectAndRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectAndforkRepositories 创建项目并fork仓库
//
// 创建仓库后fork仓库 传入参数：仓库名、是否导入项目成员、归属项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectAndforkRepositories(request *model.CreateProjectAndforkRepositoriesRequest) (*model.CreateProjectAndforkRepositoriesResponse, error) {
	requestDef := GenReqDefForCreateProjectAndforkRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectAndforkRepositoriesResponse), nil
	}
}

// CreateProjectAndforkRepositoriesInvoker 创建项目并fork仓库
func (c *CodeHubClient) CreateProjectAndforkRepositoriesInvoker(request *model.CreateProjectAndforkRepositoriesRequest) *CreateProjectAndforkRepositoriesInvoker {
	requestDef := GenReqDefForCreateProjectAndforkRepositories()
	return &CreateProjectAndforkRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserAllRepositories 查询用户的所有仓库
//
// 获取用户的所有仓库信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListUserAllRepositories(request *model.ListUserAllRepositoriesRequest) (*model.ListUserAllRepositoriesResponse, error) {
	requestDef := GenReqDefForListUserAllRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserAllRepositoriesResponse), nil
	}
}

// ListUserAllRepositoriesInvoker 查询用户的所有仓库
func (c *CodeHubClient) ListUserAllRepositoriesInvoker(request *model.ListUserAllRepositoriesRequest) *ListUserAllRepositoriesInvoker {
	requestDef := GenReqDefForListUserAllRepositories()
	return &ListUserAllRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllRepositoryByTwoProjectId 查询项目下的所有仓库
//
// 获取仓库列表,模糊查询支持范围,如果未传入project uuid，则支持按仓库名或项目名模糊查询，否则，只按仓库名模糊匹配
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowAllRepositoryByTwoProjectId(request *model.ShowAllRepositoryByTwoProjectIdRequest) (*model.ShowAllRepositoryByTwoProjectIdResponse, error) {
	requestDef := GenReqDefForShowAllRepositoryByTwoProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllRepositoryByTwoProjectIdResponse), nil
	}
}

// ShowAllRepositoryByTwoProjectIdInvoker 查询项目下的所有仓库
func (c *CodeHubClient) ShowAllRepositoryByTwoProjectIdInvoker(request *model.ShowAllRepositoryByTwoProjectIdRequest) *ShowAllRepositoryByTwoProjectIdInvoker {
	requestDef := GenReqDefForShowAllRepositoryByTwoProjectId()
	return &ShowAllRepositoryByTwoProjectIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddHooks 为指定仓库添加hook
//
// 提交代码自动触发编译构建，添加仓库钩子
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddHooks(request *model.AddHooksRequest) (*model.AddHooksResponse, error) {
	requestDef := GenReqDefForAddHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddHooksResponse), nil
	}
}

// AddHooksInvoker 为指定仓库添加hook
func (c *CodeHubClient) AddHooksInvoker(request *model.AddHooksRequest) *AddHooksInvoker {
	requestDef := GenReqDefForAddHooks()
	return &AddHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHooks 删除指定仓库的 hook
//
// 提交代码自动触发编译构建，删除仓库钩子
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteHooks(request *model.DeleteHooksRequest) (*model.DeleteHooksResponse, error) {
	requestDef := GenReqDefForDeleteHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHooksResponse), nil
	}
}

// DeleteHooksInvoker 删除指定仓库的 hook
func (c *CodeHubClient) DeleteHooksInvoker(request *model.DeleteHooksRequest) *DeleteHooksInvoker {
	requestDef := GenReqDefForDeleteHooks()
	return &DeleteHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHooks 查询指定仓库的webhook
//
// 获取仓库webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListHooks(request *model.ListHooksRequest) (*model.ListHooksResponse, error) {
	requestDef := GenReqDefForListHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHooksResponse), nil
	}
}

// ListHooksInvoker 查询指定仓库的webhook
func (c *CodeHubClient) ListHooksInvoker(request *model.ListHooksRequest) *ListHooksInvoker {
	requestDef := GenReqDefForListHooks()
	return &ListHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
