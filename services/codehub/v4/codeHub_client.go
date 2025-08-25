package v4

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codehub/v4/model"
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

// CreateBranch 创建分支
//
// 创建分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateBranch(request *model.CreateBranchRequest) (*model.CreateBranchResponse, error) {
	requestDef := GenReqDefForCreateBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBranchResponse), nil
	}
}

// CreateBranchInvoker 创建分支
func (c *CodeHubClient) CreateBranchInvoker(request *model.CreateBranchRequest) *CreateBranchInvoker {
	requestDef := GenReqDefForCreateBranch()
	return &CreateBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBranch 删除分支
//
// 删除分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteBranch(request *model.DeleteBranchRequest) (*model.DeleteBranchResponse, error) {
	requestDef := GenReqDefForDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBranchResponse), nil
	}
}

// DeleteBranchInvoker 删除分支
func (c *CodeHubClient) DeleteBranchInvoker(request *model.DeleteBranchRequest) *DeleteBranchInvoker {
	requestDef := GenReqDefForDeleteBranch()
	return &DeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranches 获取分支列表
//
// 获取分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListBranches(request *model.ListBranchesRequest) (*model.ListBranchesResponse, error) {
	requestDef := GenReqDefForListBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchesResponse), nil
	}
}

// ListBranchesInvoker 获取分支列表
func (c *CodeHubClient) ListBranchesInvoker(request *model.ListBranchesRequest) *ListBranchesInvoker {
	requestDef := GenReqDefForListBranches()
	return &ListBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranch 获取分支详情
//
// 获取分支详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowBranch(request *model.ShowBranchRequest) (*model.ShowBranchResponse, error) {
	requestDef := GenReqDefForShowBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchResponse), nil
	}
}

// ShowBranchInvoker 获取分支详情
func (c *CodeHubClient) ShowBranchInvoker(request *model.ShowBranchRequest) *ShowBranchInvoker {
	requestDef := GenReqDefForShowBranch()
	return &ShowBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBranchName 分支重命名
//
// 分支重命名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateBranchName(request *model.UpdateBranchNameRequest) (*model.UpdateBranchNameResponse, error) {
	requestDef := GenReqDefForUpdateBranchName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBranchNameResponse), nil
	}
}

// UpdateBranchNameInvoker 分支重命名
func (c *CodeHubClient) UpdateBranchNameInvoker(request *model.UpdateBranchNameRequest) *UpdateBranchNameInvoker {
	requestDef := GenReqDefForUpdateBranchName()
	return &UpdateBranchNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCommit 创建提交信息
//
// 创建提交信息
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

// CreateCommitInvoker 创建提交信息
func (c *CodeHubClient) CreateCommitInvoker(request *model.CreateCommitRequest) *CreateCommitInvoker {
	requestDef := GenReqDefForCreateCommit()
	return &CreateCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCommitRevert 回退提交
//
// 回退提交
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateCommitRevert(request *model.CreateCommitRevertRequest) (*model.CreateCommitRevertResponse, error) {
	requestDef := GenReqDefForCreateCommitRevert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommitRevertResponse), nil
	}
}

// CreateCommitRevertInvoker 回退提交
func (c *CodeHubClient) CreateCommitRevertInvoker(request *model.CreateCommitRevertRequest) *CreateCommitRevertInvoker {
	requestDef := GenReqDefForCreateCommitRevert()
	return &CreateCommitRevertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitAssociatedRefs 根据提交ID查询分支、标签列表
//
// 根据提交ID查询分支、标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListCommitAssociatedRefs(request *model.ListCommitAssociatedRefsRequest) (*model.ListCommitAssociatedRefsResponse, error) {
	requestDef := GenReqDefForListCommitAssociatedRefs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitAssociatedRefsResponse), nil
	}
}

// ListCommitAssociatedRefsInvoker 根据提交ID查询分支、标签列表
func (c *CodeHubClient) ListCommitAssociatedRefsInvoker(request *model.ListCommitAssociatedRefsRequest) *ListCommitAssociatedRefsInvoker {
	requestDef := GenReqDefForListCommitAssociatedRefs()
	return &ListCommitAssociatedRefsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommit 获取特定提交信息
//
// 获取特定提交信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommit(request *model.ShowCommitRequest) (*model.ShowCommitResponse, error) {
	requestDef := GenReqDefForShowCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitResponse), nil
	}
}

// ShowCommitInvoker 获取特定提交信息
func (c *CodeHubClient) ShowCommitInvoker(request *model.ShowCommitRequest) *ShowCommitInvoker {
	requestDef := GenReqDefForShowCommit()
	return &ShowCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitDiffMetadata 获取commit引入的文件变更元数据
//
// 获取commit引入的文件变更元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitDiffMetadata(request *model.ShowCommitDiffMetadataRequest) (*model.ShowCommitDiffMetadataResponse, error) {
	requestDef := GenReqDefForShowCommitDiffMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitDiffMetadataResponse), nil
	}
}

// ShowCommitDiffMetadataInvoker 获取commit引入的文件变更元数据
func (c *CodeHubClient) ShowCommitDiffMetadataInvoker(request *model.ShowCommitDiffMetadataRequest) *ShowCommitDiffMetadataInvoker {
	requestDef := GenReqDefForShowCommitDiffMetadata()
	return &ShowCommitDiffMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitFileDiff 获取commit引入的指定文件的变更内容
//
// 获取commit引入的指定文件的变更内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitFileDiff(request *model.ShowCommitFileDiffRequest) (*model.ShowCommitFileDiffResponse, error) {
	requestDef := GenReqDefForShowCommitFileDiff()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitFileDiffResponse), nil
	}
}

// ShowCommitFileDiffInvoker 获取commit引入的指定文件的变更内容
func (c *CodeHubClient) ShowCommitFileDiffInvoker(request *model.ShowCommitFileDiffRequest) *ShowCommitFileDiffInvoker {
	requestDef := GenReqDefForShowCommitFileDiff()
	return &ShowCommitFileDiffInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiffCommit 获取提交差异
//
// 获取提交差异
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

// ShowDiffCommitInvoker 获取提交差异
func (c *CodeHubClient) ShowDiffCommitInvoker(request *model.ShowDiffCommitRequest) *ShowDiffCommitInvoker {
	requestDef := GenReqDefForShowDiffCommit()
	return &ShowDiffCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussion 创建合并请求检视意见
//
// 创建合并请求检视意见
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

// CreateMergeRequestDiscussionInvoker 创建合并请求检视意见
func (c *CodeHubClient) CreateMergeRequestDiscussionInvoker(request *model.CreateMergeRequestDiscussionRequest) *CreateMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussion()
	return &CreateMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussionResponse 回复合并请求检视意见
//
// 回复合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequestDiscussionResponse(request *model.CreateMergeRequestDiscussionResponseRequest) (*model.CreateMergeRequestDiscussionResponseResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestDiscussionResponse()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestDiscussionResponseResponse), nil
	}
}

// CreateMergeRequestDiscussionResponseInvoker 回复合并请求检视意见
func (c *CodeHubClient) CreateMergeRequestDiscussionResponseInvoker(request *model.CreateMergeRequestDiscussionResponseRequest) *CreateMergeRequestDiscussionResponseInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussionResponse()
	return &CreateMergeRequestDiscussionResponseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReviewSetting 创建/更新检视意见设置
//
// 创建/更新检视意见设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateReviewSetting(request *model.CreateReviewSettingRequest) (*model.CreateReviewSettingResponse, error) {
	requestDef := GenReqDefForCreateReviewSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReviewSettingResponse), nil
	}
}

// CreateReviewSettingInvoker 创建/更新检视意见设置
func (c *CodeHubClient) CreateReviewSettingInvoker(request *model.CreateReviewSettingRequest) *CreateReviewSettingInvoker {
	requestDef := GenReqDefForCreateReviewSetting()
	return &CreateReviewSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDefaultReviewCategories 获取默认的检视意见分类
//
// 获取默认的检视意见分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListDefaultReviewCategories(request *model.ListDefaultReviewCategoriesRequest) (*model.ListDefaultReviewCategoriesResponse, error) {
	requestDef := GenReqDefForListDefaultReviewCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDefaultReviewCategoriesResponse), nil
	}
}

// ListDefaultReviewCategoriesInvoker 获取默认的检视意见分类
func (c *CodeHubClient) ListDefaultReviewCategoriesInvoker(request *model.ListDefaultReviewCategoriesRequest) *ListDefaultReviewCategoriesInvoker {
	requestDef := GenReqDefForListDefaultReviewCategories()
	return &ListDefaultReviewCategoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestDiscussions 获取合并请求检视意见列表
//
// 获取合并请求检视意见列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestDiscussions(request *model.ListMergeRequestDiscussionsRequest) (*model.ListMergeRequestDiscussionsResponse, error) {
	requestDef := GenReqDefForListMergeRequestDiscussions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestDiscussionsResponse), nil
	}
}

// ListMergeRequestDiscussionsInvoker 获取合并请求检视意见列表
func (c *CodeHubClient) ListMergeRequestDiscussionsInvoker(request *model.ListMergeRequestDiscussionsRequest) *ListMergeRequestDiscussionsInvoker {
	requestDef := GenReqDefForListMergeRequestDiscussions()
	return &ListMergeRequestDiscussionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectNoteRequiredAttributes 获取项目检视意见必填项
//
// 获取项目检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectNoteRequiredAttributes(request *model.ListProjectNoteRequiredAttributesRequest) (*model.ListProjectNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForListProjectNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectNoteRequiredAttributesResponse), nil
	}
}

// ListProjectNoteRequiredAttributesInvoker 获取项目检视意见必填项
func (c *CodeHubClient) ListProjectNoteRequiredAttributesInvoker(request *model.ListProjectNoteRequiredAttributesRequest) *ListProjectNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForListProjectNoteRequiredAttributes()
	return &ListProjectNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryReviewAuthors 获取仓库下检视意见作者列表
//
// 获取仓库下检视意见作者列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryReviewAuthors(request *model.ListRepositoryReviewAuthorsRequest) (*model.ListRepositoryReviewAuthorsResponse, error) {
	requestDef := GenReqDefForListRepositoryReviewAuthors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryReviewAuthorsResponse), nil
	}
}

// ListRepositoryReviewAuthorsInvoker 获取仓库下检视意见作者列表
func (c *CodeHubClient) ListRepositoryReviewAuthorsInvoker(request *model.ListRepositoryReviewAuthorsRequest) *ListRepositoryReviewAuthorsInvoker {
	requestDef := GenReqDefForListRepositoryReviewAuthors()
	return &ListRepositoryReviewAuthorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryReviews 获取仓库检视意见列表
//
// 获取仓库检视意见列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryReviews(request *model.ListRepositoryReviewsRequest) (*model.ListRepositoryReviewsResponse, error) {
	requestDef := GenReqDefForListRepositoryReviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryReviewsResponse), nil
	}
}

// ListRepositoryReviewsInvoker 获取仓库检视意见列表
func (c *CodeHubClient) ListRepositoryReviewsInvoker(request *model.ListRepositoryReviewsRequest) *ListRepositoryReviewsInvoker {
	requestDef := GenReqDefForListRepositoryReviews()
	return &ListRepositoryReviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupNoteRequiredAttributes 获取代码组检视意见必填项
//
// 获取代码组检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupNoteRequiredAttributes(request *model.ShowGroupNoteRequiredAttributesRequest) (*model.ShowGroupNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForShowGroupNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupNoteRequiredAttributesResponse), nil
	}
}

// ShowGroupNoteRequiredAttributesInvoker 获取代码组检视意见必填项
func (c *CodeHubClient) ShowGroupNoteRequiredAttributesInvoker(request *model.ShowGroupNoteRequiredAttributesRequest) *ShowGroupNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForShowGroupNoteRequiredAttributes()
	return &ShowGroupNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupReviewSettings 获取代码组检视意见设置(不含必填项)
//
// 获取代码组检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupReviewSettings(request *model.ShowGroupReviewSettingsRequest) (*model.ShowGroupReviewSettingsResponse, error) {
	requestDef := GenReqDefForShowGroupReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupReviewSettingsResponse), nil
	}
}

// ShowGroupReviewSettingsInvoker 获取代码组检视意见设置(不含必填项)
func (c *CodeHubClient) ShowGroupReviewSettingsInvoker(request *model.ShowGroupReviewSettingsRequest) *ShowGroupReviewSettingsInvoker {
	requestDef := GenReqDefForShowGroupReviewSettings()
	return &ShowGroupReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestDiscussion 根据discussion_id获取合并请求检视意见
//
// 根据discussion_id获取合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestDiscussion(request *model.ShowMergeRequestDiscussionRequest) (*model.ShowMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForShowMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestDiscussionResponse), nil
	}
}

// ShowMergeRequestDiscussionInvoker 根据discussion_id获取合并请求检视意见
func (c *CodeHubClient) ShowMergeRequestDiscussionInvoker(request *model.ShowMergeRequestDiscussionRequest) *ShowMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForShowMergeRequestDiscussion()
	return &ShowMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNoteRequiredAttributes 获取仓库检视意见必填项
//
// 获取仓库检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowNoteRequiredAttributes(request *model.ShowNoteRequiredAttributesRequest) (*model.ShowNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForShowNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNoteRequiredAttributesResponse), nil
	}
}

// ShowNoteRequiredAttributesInvoker 获取仓库检视意见必填项
func (c *CodeHubClient) ShowNoteRequiredAttributesInvoker(request *model.ShowNoteRequiredAttributesRequest) *ShowNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForShowNoteRequiredAttributes()
	return &ShowNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectReviewSettings 获取项目检视意见设置(不含必填项)
//
// 获取项目检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectReviewSettings(request *model.ShowProjectReviewSettingsRequest) (*model.ShowProjectReviewSettingsResponse, error) {
	requestDef := GenReqDefForShowProjectReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectReviewSettingsResponse), nil
	}
}

// ShowProjectReviewSettingsInvoker 获取项目检视意见设置(不含必填项)
func (c *CodeHubClient) ShowProjectReviewSettingsInvoker(request *model.ShowProjectReviewSettingsRequest) *ShowProjectReviewSettingsInvoker {
	requestDef := GenReqDefForShowProjectReviewSettings()
	return &ShowProjectReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateGroupNoteRequiredAttributes 创建/更新代码组检视意见必填项
//
// 创建/更新代码组检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateGroupNoteRequiredAttributes(request *model.UpdateGroupNoteRequiredAttributesRequest) (*model.UpdateGroupNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateGroupNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupNoteRequiredAttributesResponse), nil
	}
}

// UpdateGroupNoteRequiredAttributesInvoker 创建/更新代码组检视意见必填项
func (c *CodeHubClient) UpdateGroupNoteRequiredAttributesInvoker(request *model.UpdateGroupNoteRequiredAttributesRequest) *UpdateGroupNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateGroupNoteRequiredAttributes()
	return &UpdateGroupNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupReviewSettings 创建/更新代码组检视意见设置(不含必填项)
//
// 创建/更新代码组检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateGroupReviewSettings(request *model.UpdateGroupReviewSettingsRequest) (*model.UpdateGroupReviewSettingsResponse, error) {
	requestDef := GenReqDefForUpdateGroupReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupReviewSettingsResponse), nil
	}
}

// UpdateGroupReviewSettingsInvoker 创建/更新代码组检视意见设置(不含必填项)
func (c *CodeHubClient) UpdateGroupReviewSettingsInvoker(request *model.UpdateGroupReviewSettingsRequest) *UpdateGroupReviewSettingsInvoker {
	requestDef := GenReqDefForUpdateGroupReviewSettings()
	return &UpdateGroupReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestDiscussion 更新合并请求检视意见
//
// 更新合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestDiscussion(request *model.UpdateMergeRequestDiscussionRequest) (*model.UpdateMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestDiscussionResponse), nil
	}
}

// UpdateMergeRequestDiscussionInvoker 更新合并请求检视意见
func (c *CodeHubClient) UpdateMergeRequestDiscussionInvoker(request *model.UpdateMergeRequestDiscussionRequest) *UpdateMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForUpdateMergeRequestDiscussion()
	return &UpdateMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNoteRequiredAttributes 创建/更新仓库检视意见必填项
//
// 创建/更新仓库检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateNoteRequiredAttributes(request *model.UpdateNoteRequiredAttributesRequest) (*model.UpdateNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNoteRequiredAttributesResponse), nil
	}
}

// UpdateNoteRequiredAttributesInvoker 创建/更新仓库检视意见必填项
func (c *CodeHubClient) UpdateNoteRequiredAttributesInvoker(request *model.UpdateNoteRequiredAttributesRequest) *UpdateNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateNoteRequiredAttributes()
	return &UpdateNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectNoteRequiredAttributes 创建/更新项目检视意见必填项
//
// 创建/更新项目检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectNoteRequiredAttributes(request *model.UpdateProjectNoteRequiredAttributesRequest) (*model.UpdateProjectNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateProjectNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectNoteRequiredAttributesResponse), nil
	}
}

// UpdateProjectNoteRequiredAttributesInvoker 创建/更新项目检视意见必填项
func (c *CodeHubClient) UpdateProjectNoteRequiredAttributesInvoker(request *model.UpdateProjectNoteRequiredAttributesRequest) *UpdateProjectNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateProjectNoteRequiredAttributes()
	return &UpdateProjectNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectReviewSettings 创建/更新项目检视意见设置(不含必填项)
//
// 创建/更新项目检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectReviewSettings(request *model.UpdateProjectReviewSettingsRequest) (*model.UpdateProjectReviewSettingsResponse, error) {
	requestDef := GenReqDefForUpdateProjectReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectReviewSettingsResponse), nil
	}
}

// UpdateProjectReviewSettingsInvoker 创建/更新项目检视意见设置(不含必填项)
func (c *CodeHubClient) UpdateProjectReviewSettingsInvoker(request *model.UpdateProjectReviewSettingsRequest) *UpdateProjectReviewSettingsInvoker {
	requestDef := GenReqDefForUpdateProjectReviewSettings()
	return &UpdateProjectReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建文件
//
// 创建文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建文件
func (c *CodeHubClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除文件
//
// 删除文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除文件
func (c *CodeHubClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBlobsRaw 获取仓库单个文件内容(下载单个文件)
//
// 获取仓库单个文件内容(下载单个文件)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DownloadBlobsRaw(request *model.DownloadBlobsRawRequest) (*model.DownloadBlobsRawResponse, error) {
	requestDef := GenReqDefForDownloadBlobsRaw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlobsRawResponse), nil
	}
}

// DownloadBlobsRawInvoker 获取仓库单个文件内容(下载单个文件)
func (c *CodeHubClient) DownloadBlobsRawInvoker(request *model.DownloadBlobsRawRequest) *DownloadBlobsRawInvoker {
	requestDef := GenReqDefForDownloadBlobsRaw()
	return &DownloadBlobsRawInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileBlameLines 获取文件追溯信息
//
// 获取文件追溯信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListFileBlameLines(request *model.ListFileBlameLinesRequest) (*model.ListFileBlameLinesResponse, error) {
	requestDef := GenReqDefForListFileBlameLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileBlameLinesResponse), nil
	}
}

// ListFileBlameLinesInvoker 获取文件追溯信息
func (c *CodeHubClient) ListFileBlameLinesInvoker(request *model.ListFileBlameLinesRequest) *ListFileBlameLinesInvoker {
	requestDef := GenReqDefForListFileBlameLines()
	return &ListFileBlameLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileUpperTreeEntries 获取当前文件上级树结构
//
// 获取当前文件上级树结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListFileUpperTreeEntries(request *model.ListFileUpperTreeEntriesRequest) (*model.ListFileUpperTreeEntriesResponse, error) {
	requestDef := GenReqDefForListFileUpperTreeEntries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileUpperTreeEntriesResponse), nil
	}
}

// ListFileUpperTreeEntriesInvoker 获取当前文件上级树结构
func (c *CodeHubClient) ListFileUpperTreeEntriesInvoker(request *model.ListFileUpperTreeEntriesRequest) *ListFileUpperTreeEntriesInvoker {
	requestDef := GenReqDefForListFileUpperTreeEntries()
	return &ListFileUpperTreeEntriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFiles 获取指定分支下所有的文件列表
//
// 获取指定分支下所有的文件列表
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

// ListFilesInvoker 获取指定分支下所有的文件列表
func (c *CodeHubClient) ListFilesInvoker(request *model.ListFilesRequest) *ListFilesInvoker {
	requestDef := GenReqDefForListFiles()
	return &ListFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFile 查看文件属性与内容
//
// 查看文件属性与内容
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

// ShowFileInvoker 查看文件属性与内容
func (c *CodeHubClient) ShowFileInvoker(request *model.ShowFileRequest) *ShowFileInvoker {
	requestDef := GenReqDefForShowFile()
	return &ShowFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileContent 获取文件内容
//
// 获取文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowFileContent(request *model.ShowFileContentRequest) (*model.ShowFileContentResponse, error) {
	requestDef := GenReqDefForShowFileContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileContentResponse), nil
	}
}

// ShowFileContentInvoker 获取文件内容
func (c *CodeHubClient) ShowFileContentInvoker(request *model.ShowFileContentRequest) *ShowFileContentInvoker {
	requestDef := GenReqDefForShowFileContent()
	return &ShowFileContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReadmeFile 获取仓库默认分支的Readme文件内容
//
// 获取仓库默认分支的Readme文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowReadmeFile(request *model.ShowReadmeFileRequest) (*model.ShowReadmeFileResponse, error) {
	requestDef := GenReqDefForShowReadmeFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReadmeFileResponse), nil
	}
}

// ShowReadmeFileInvoker 获取仓库默认分支的Readme文件内容
func (c *CodeHubClient) ShowReadmeFileInvoker(request *model.ShowReadmeFileRequest) *ShowReadmeFileInvoker {
	requestDef := GenReqDefForShowReadmeFile()
	return &ShowReadmeFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFile 更新文件内容
//
// 修改仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateFile(request *model.UpdateFileRequest) (*model.UpdateFileResponse, error) {
	requestDef := GenReqDefForUpdateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFileResponse), nil
	}
}

// UpdateFileInvoker 更新文件内容
func (c *CodeHubClient) UpdateFileInvoker(request *model.UpdateFileRequest) *UpdateFileInvoker {
	requestDef := GenReqDefForUpdateFile()
	return &UpdateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRepositoryFilePushPermissions 批量删除仓库文件推送权限
//
// 批量删除仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchDeleteRepositoryFilePushPermissions(request *model.BatchDeleteRepositoryFilePushPermissionsRequest) (*model.BatchDeleteRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRepositoryFilePushPermissionsResponse), nil
	}
}

// BatchDeleteRepositoryFilePushPermissionsInvoker 批量删除仓库文件推送权限
func (c *CodeHubClient) BatchDeleteRepositoryFilePushPermissionsInvoker(request *model.BatchDeleteRepositoryFilePushPermissionsRequest) *BatchDeleteRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForBatchDeleteRepositoryFilePushPermissions()
	return &BatchDeleteRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateRepositoryFilePushPermissions 批量更新仓库文件推送权限
//
// 批量更新仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchUpdateRepositoryFilePushPermissions(request *model.BatchUpdateRepositoryFilePushPermissionsRequest) (*model.BatchUpdateRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForBatchUpdateRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateRepositoryFilePushPermissionsResponse), nil
	}
}

// BatchUpdateRepositoryFilePushPermissionsInvoker 批量更新仓库文件推送权限
func (c *CodeHubClient) BatchUpdateRepositoryFilePushPermissionsInvoker(request *model.BatchUpdateRepositoryFilePushPermissionsRequest) *BatchUpdateRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForBatchUpdateRepositoryFilePushPermissions()
	return &BatchUpdateRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFilePushPermission 创建仓库文件推送权限
//
// 创建仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateFilePushPermission(request *model.CreateFilePushPermissionRequest) (*model.CreateFilePushPermissionResponse, error) {
	requestDef := GenReqDefForCreateFilePushPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFilePushPermissionResponse), nil
	}
}

// CreateFilePushPermissionInvoker 创建仓库文件推送权限
func (c *CodeHubClient) CreateFilePushPermissionInvoker(request *model.CreateFilePushPermissionRequest) *CreateFilePushPermissionInvoker {
	requestDef := GenReqDefForCreateFilePushPermission()
	return &CreateFilePushPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryFilePushPermissions 获取仓库文件推送权限列表
//
// 获取仓库文件推送权限列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryFilePushPermissions(request *model.ListRepositoryFilePushPermissionsRequest) (*model.ListRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForListRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryFilePushPermissionsResponse), nil
	}
}

// ListRepositoryFilePushPermissionsInvoker 获取仓库文件推送权限列表
func (c *CodeHubClient) ListRepositoryFilePushPermissionsInvoker(request *model.ListRepositoryFilePushPermissionsRequest) *ListRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForListRepositoryFilePushPermissions()
	return &ListRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateGroupUserGroup 关联代码组与成员组
//
// 关联代码组与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateGroupUserGroup(request *model.AssociateGroupUserGroupRequest) (*model.AssociateGroupUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateGroupUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateGroupUserGroupResponse), nil
	}
}

// AssociateGroupUserGroupInvoker 关联代码组与成员组
func (c *CodeHubClient) AssociateGroupUserGroupInvoker(request *model.AssociateGroupUserGroupRequest) *AssociateGroupUserGroupInvoker {
	requestDef := GenReqDefForAssociateGroupUserGroup()
	return &AssociateGroupUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroup 创建代码组
//
// 创建代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建代码组
func (c *CodeHubClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除代码组
//
// 删除代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除代码组
func (c *CodeHubClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupAddableMembers 获取代码组下可添加的成员列表
//
// 获取代码组下可添加的成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupAddableMembers(request *model.ListGroupAddableMembersRequest) (*model.ListGroupAddableMembersResponse, error) {
	requestDef := GenReqDefForListGroupAddableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupAddableMembersResponse), nil
	}
}

// ListGroupAddableMembersInvoker 获取代码组下可添加的成员列表
func (c *CodeHubClient) ListGroupAddableMembersInvoker(request *model.ListGroupAddableMembersRequest) *ListGroupAddableMembersInvoker {
	requestDef := GenReqDefForListGroupAddableMembers()
	return &ListGroupAddableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupAddableUserGroups 获取代码组下可添加的成员组
//
// 获取代码组下可添加的成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupAddableUserGroups(request *model.ListGroupAddableUserGroupsRequest) (*model.ListGroupAddableUserGroupsResponse, error) {
	requestDef := GenReqDefForListGroupAddableUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupAddableUserGroupsResponse), nil
	}
}

// ListGroupAddableUserGroupsInvoker 获取代码组下可添加的成员组
func (c *CodeHubClient) ListGroupAddableUserGroupsInvoker(request *model.ListGroupAddableUserGroupsRequest) *ListGroupAddableUserGroupsInvoker {
	requestDef := GenReqDefForListGroupAddableUserGroups()
	return &ListGroupAddableUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMembers 获取代码组下可添加的成员列表
//
// 获取代码组下可添加的成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupMembers(request *model.ListGroupMembersRequest) (*model.ListGroupMembersResponse, error) {
	requestDef := GenReqDefForListGroupMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMembersResponse), nil
	}
}

// ListGroupMembersInvoker 获取代码组下可添加的成员列表
func (c *CodeHubClient) ListGroupMembersInvoker(request *model.ListGroupMembersRequest) *ListGroupMembersInvoker {
	requestDef := GenReqDefForListGroupMembers()
	return &ListGroupMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupPermissionResources 获取代码组权限资源点列表
//
// 获取代码组权限资源点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupPermissionResources(request *model.ListGroupPermissionResourcesRequest) (*model.ListGroupPermissionResourcesResponse, error) {
	requestDef := GenReqDefForListGroupPermissionResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupPermissionResourcesResponse), nil
	}
}

// ListGroupPermissionResourcesInvoker 获取代码组权限资源点列表
func (c *CodeHubClient) ListGroupPermissionResourcesInvoker(request *model.ListGroupPermissionResourcesRequest) *ListGroupPermissionResourcesInvoker {
	requestDef := GenReqDefForListGroupPermissionResources()
	return &ListGroupPermissionResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupSubgroupsAndRepositories 获取代码组下的子代码组和仓库列表
//
// 获取代码组下的子代码组和仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupSubgroupsAndRepositories(request *model.ListGroupSubgroupsAndRepositoriesRequest) (*model.ListGroupSubgroupsAndRepositoriesResponse, error) {
	requestDef := GenReqDefForListGroupSubgroupsAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupSubgroupsAndRepositoriesResponse), nil
	}
}

// ListGroupSubgroupsAndRepositoriesInvoker 获取代码组下的子代码组和仓库列表
func (c *CodeHubClient) ListGroupSubgroupsAndRepositoriesInvoker(request *model.ListGroupSubgroupsAndRepositoriesRequest) *ListGroupSubgroupsAndRepositoriesInvoker {
	requestDef := GenReqDefForListGroupSubgroupsAndRepositories()
	return &ListGroupSubgroupsAndRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 获取代码组列表
//
// 获取代码组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 获取代码组列表
func (c *CodeHubClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImpersonationTokens 获取用户的个人访问令牌
//
// 获取用户的个人访问令牌
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListImpersonationTokens(request *model.ListImpersonationTokensRequest) (*model.ListImpersonationTokensResponse, error) {
	requestDef := GenReqDefForListImpersonationTokens()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImpersonationTokensResponse), nil
	}
}

// ListImpersonationTokensInvoker 获取用户的个人访问令牌
func (c *CodeHubClient) ListImpersonationTokensInvoker(request *model.ListImpersonationTokensRequest) *ListImpersonationTokensInvoker {
	requestDef := GenReqDefForListImpersonationTokens()
	return &ListImpersonationTokensInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductPermissionResourcesGrantedUsers 获取项目下成员列表
//
// 获取项目下成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProductPermissionResourcesGrantedUsers(request *model.ListProductPermissionResourcesGrantedUsersRequest) (*model.ListProductPermissionResourcesGrantedUsersResponse, error) {
	requestDef := GenReqDefForListProductPermissionResourcesGrantedUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductPermissionResourcesGrantedUsersResponse), nil
	}
}

// ListProductPermissionResourcesGrantedUsersInvoker 获取项目下成员列表
func (c *CodeHubClient) ListProductPermissionResourcesGrantedUsersInvoker(request *model.ListProductPermissionResourcesGrantedUsersRequest) *ListProductPermissionResourcesGrantedUsersInvoker {
	requestDef := GenReqDefForListProductPermissionResourcesGrantedUsers()
	return &ListProductPermissionResourcesGrantedUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectSubgroupsAndRepositories 获取项目下的代码组和仓库列表
//
// 获取项目下的代码组和仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectSubgroupsAndRepositories(request *model.ListProjectSubgroupsAndRepositoriesRequest) (*model.ListProjectSubgroupsAndRepositoriesResponse, error) {
	requestDef := GenReqDefForListProjectSubgroupsAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectSubgroupsAndRepositoriesResponse), nil
	}
}

// ListProjectSubgroupsAndRepositoriesInvoker 获取项目下的代码组和仓库列表
func (c *CodeHubClient) ListProjectSubgroupsAndRepositoriesInvoker(request *model.ListProjectSubgroupsAndRepositoriesRequest) *ListProjectSubgroupsAndRepositoriesInvoker {
	requestDef := GenReqDefForListProjectSubgroupsAndRepositories()
	return &ListProjectSubgroupsAndRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroup 获取代码组信息
//
// 获取代码组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

// ShowGroupInvoker 获取代码组信息
func (c *CodeHubClient) ShowGroupInvoker(request *model.ShowGroupRequest) *ShowGroupInvoker {
	requestDef := GenReqDefForShowGroup()
	return &ShowGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupGeneralPolicy 获取指定代码组的基本设置信息
//
// 获取指定代码组的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupGeneralPolicy(request *model.ShowGroupGeneralPolicyRequest) (*model.ShowGroupGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowGroupGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupGeneralPolicyResponse), nil
	}
}

// ShowGroupGeneralPolicyInvoker 获取指定代码组的基本设置信息
func (c *CodeHubClient) ShowGroupGeneralPolicyInvoker(request *model.ShowGroupGeneralPolicyRequest) *ShowGroupGeneralPolicyInvoker {
	requestDef := GenReqDefForShowGroupGeneralPolicy()
	return &ShowGroupGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupPermissionInheritEnabled 获取代码组继承权限设置开关
//
// 获取代码组继承权限设置开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupPermissionInheritEnabled(request *model.ShowGroupPermissionInheritEnabledRequest) (*model.ShowGroupPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForShowGroupPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupPermissionInheritEnabledResponse), nil
	}
}

// ShowGroupPermissionInheritEnabledInvoker 获取代码组继承权限设置开关
func (c *CodeHubClient) ShowGroupPermissionInheritEnabledInvoker(request *model.ShowGroupPermissionInheritEnabledRequest) *ShowGroupPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForShowGroupPermissionInheritEnabled()
	return &ShowGroupPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupSettingsInheritCfg 获取代码组继承设置项
//
// 获取代码组继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupSettingsInheritCfg(request *model.ShowGroupSettingsInheritCfgRequest) (*model.ShowGroupSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForShowGroupSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupSettingsInheritCfgResponse), nil
	}
}

// ShowGroupSettingsInheritCfgInvoker 获取代码组继承设置项
func (c *CodeHubClient) ShowGroupSettingsInheritCfgInvoker(request *model.ShowGroupSettingsInheritCfgRequest) *ShowGroupSettingsInheritCfgInvoker {
	requestDef := GenReqDefForShowGroupSettingsInheritCfg()
	return &ShowGroupSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWatermark 获取代码组水印设置
//
// 获取代码组水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupWatermark(request *model.ShowGroupWatermarkRequest) (*model.ShowGroupWatermarkResponse, error) {
	requestDef := GenReqDefForShowGroupWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWatermarkResponse), nil
	}
}

// ShowGroupWatermarkInvoker 获取代码组水印设置
func (c *CodeHubClient) ShowGroupWatermarkInvoker(request *model.ShowGroupWatermarkRequest) *ShowGroupWatermarkInvoker {
	requestDef := GenReqDefForShowGroupWatermark()
	return &ShowGroupWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupsInherit 获取代码组的继承设置
//
// 获取代码组的继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupsInherit(request *model.ShowGroupsInheritRequest) (*model.ShowGroupsInheritResponse, error) {
	requestDef := GenReqDefForShowGroupsInherit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsInheritResponse), nil
	}
}

// ShowGroupsInheritInvoker 获取代码组的继承设置
func (c *CodeHubClient) ShowGroupsInheritInvoker(request *model.ShowGroupsInheritRequest) *ShowGroupsInheritInvoker {
	requestDef := GenReqDefForShowGroupsInherit()
	return &ShowGroupsInheritInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectGeneralPolicy 获取指定项目的基本设置信息
//
// 获取指定项目的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectGeneralPolicy(request *model.ShowProjectGeneralPolicyRequest) (*model.ShowProjectGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowProjectGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectGeneralPolicyResponse), nil
	}
}

// ShowProjectGeneralPolicyInvoker 获取指定项目的基本设置信息
func (c *CodeHubClient) ShowProjectGeneralPolicyInvoker(request *model.ShowProjectGeneralPolicyRequest) *ShowProjectGeneralPolicyInvoker {
	requestDef := GenReqDefForShowProjectGeneralPolicy()
	return &ShowProjectGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectMemberSetting 获取项目成员设置
//
// 获取项目成员设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectMemberSetting(request *model.ShowProjectMemberSettingRequest) (*model.ShowProjectMemberSettingResponse, error) {
	requestDef := GenReqDefForShowProjectMemberSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectMemberSettingResponse), nil
	}
}

// ShowProjectMemberSettingInvoker 获取项目成员设置
func (c *CodeHubClient) ShowProjectMemberSettingInvoker(request *model.ShowProjectMemberSettingRequest) *ShowProjectMemberSettingInvoker {
	requestDef := GenReqDefForShowProjectMemberSetting()
	return &ShowProjectMemberSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourcePermissions 获取资源点对应的角色和权限
//
// 获取资源点对应的角色和权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowResourcePermissions(request *model.ShowResourcePermissionsRequest) (*model.ShowResourcePermissionsResponse, error) {
	requestDef := GenReqDefForShowResourcePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourcePermissionsResponse), nil
	}
}

// ShowResourcePermissionsInvoker 获取资源点对应的角色和权限
func (c *CodeHubClient) ShowResourcePermissionsInvoker(request *model.ShowResourcePermissionsRequest) *ShowResourcePermissionsInvoker {
	requestDef := GenReqDefForShowResourcePermissions()
	return &ShowResourcePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferGroup 移交代码组
//
// 移交代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) TransferGroup(request *model.TransferGroupRequest) (*model.TransferGroupResponse, error) {
	requestDef := GenReqDefForTransferGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferGroupResponse), nil
	}
}

// TransferGroupInvoker 移交代码组
func (c *CodeHubClient) TransferGroupInvoker(request *model.TransferGroupRequest) *TransferGroupInvoker {
	requestDef := GenReqDefForTransferGroup()
	return &TransferGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupWatermark 更新代码组水印设置
//
// 更新代码组水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateGroupWatermark(request *model.UpdateGroupWatermarkRequest) (*model.UpdateGroupWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateGroupWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupWatermarkResponse), nil
	}
}

// UpdateGroupWatermarkInvoker 更新代码组水印设置
func (c *CodeHubClient) UpdateGroupWatermarkInvoker(request *model.UpdateGroupWatermarkRequest) *UpdateGroupWatermarkInvoker {
	requestDef := GenReqDefForUpdateGroupWatermark()
	return &UpdateGroupWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepositoryLabel 创建仓库标签
//
// 创建仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateRepositoryLabel(request *model.CreateRepositoryLabelRequest) (*model.CreateRepositoryLabelResponse, error) {
	requestDef := GenReqDefForCreateRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositoryLabelResponse), nil
	}
}

// CreateRepositoryLabelInvoker 创建仓库标签
func (c *CodeHubClient) CreateRepositoryLabelInvoker(request *model.CreateRepositoryLabelRequest) *CreateRepositoryLabelInvoker {
	requestDef := GenReqDefForCreateRepositoryLabel()
	return &CreateRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepositorySystemLabels 创建仓库系统标签
//
// 创建仓库系统标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateRepositorySystemLabels(request *model.CreateRepositorySystemLabelsRequest) (*model.CreateRepositorySystemLabelsResponse, error) {
	requestDef := GenReqDefForCreateRepositorySystemLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositorySystemLabelsResponse), nil
	}
}

// CreateRepositorySystemLabelsInvoker 创建仓库系统标签
func (c *CodeHubClient) CreateRepositorySystemLabelsInvoker(request *model.CreateRepositorySystemLabelsRequest) *CreateRepositorySystemLabelsInvoker {
	requestDef := GenReqDefForCreateRepositorySystemLabels()
	return &CreateRepositorySystemLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepositoryLabel 删除仓库标签
//
// 删除仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteRepositoryLabel(request *model.DeleteRepositoryLabelRequest) (*model.DeleteRepositoryLabelResponse, error) {
	requestDef := GenReqDefForDeleteRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepositoryLabelResponse), nil
	}
}

// DeleteRepositoryLabelInvoker 删除仓库标签
func (c *CodeHubClient) DeleteRepositoryLabelInvoker(request *model.DeleteRepositoryLabelRequest) *DeleteRepositoryLabelInvoker {
	requestDef := GenReqDefForDeleteRepositoryLabel()
	return &DeleteRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryLabels 获取仓库标签列表
//
// 获取仓库标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryLabels(request *model.ListRepositoryLabelsRequest) (*model.ListRepositoryLabelsResponse, error) {
	requestDef := GenReqDefForListRepositoryLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryLabelsResponse), nil
	}
}

// ListRepositoryLabelsInvoker 获取仓库标签列表
func (c *CodeHubClient) ListRepositoryLabelsInvoker(request *model.ListRepositoryLabelsRequest) *ListRepositoryLabelsInvoker {
	requestDef := GenReqDefForListRepositoryLabels()
	return &ListRepositoryLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryLabel 修改仓库标签
//
// 修改仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryLabel(request *model.UpdateRepositoryLabelRequest) (*model.UpdateRepositoryLabelResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryLabelResponse), nil
	}
}

// UpdateRepositoryLabelInvoker 修改仓库标签
func (c *CodeHubClient) UpdateRepositoryLabelInvoker(request *model.UpdateRepositoryLabelRequest) *UpdateRepositoryLabelInvoker {
	requestDef := GenReqDefForUpdateRepositoryLabel()
	return &UpdateRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRepositoryMembers 批量添加仓库成员
//
// 批量添加仓库成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddRepositoryMembers(request *model.AddRepositoryMembersRequest) (*model.AddRepositoryMembersResponse, error) {
	requestDef := GenReqDefForAddRepositoryMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepositoryMembersResponse), nil
	}
}

// AddRepositoryMembersInvoker 批量添加仓库成员
func (c *CodeHubClient) AddRepositoryMembersInvoker(request *model.AddRepositoryMembersRequest) *AddRepositoryMembersInvoker {
	requestDef := GenReqDefForAddRepositoryMembers()
	return &AddRepositoryMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 获取仓库成员列表
//
// 获取仓库成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 获取仓库成员列表
func (c *CodeHubClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryUserGroups 获取成员组列表
//
// 获取成员组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryUserGroups(request *model.ListRepositoryUserGroupsRequest) (*model.ListRepositoryUserGroupsResponse, error) {
	requestDef := GenReqDefForListRepositoryUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryUserGroupsResponse), nil
	}
}

// ListRepositoryUserGroupsInvoker 获取成员组列表
func (c *CodeHubClient) ListRepositoryUserGroupsInvoker(request *model.ListRepositoryUserGroupsRequest) *ListRepositoryUserGroupsInvoker {
	requestDef := GenReqDefForListRepositoryUserGroups()
	return &ListRepositoryUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApprovalMergeRequest 审核合并请求
//
// 审核合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ApprovalMergeRequest(request *model.ApprovalMergeRequestRequest) (*model.ApprovalMergeRequestResponse, error) {
	requestDef := GenReqDefForApprovalMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApprovalMergeRequestResponse), nil
	}
}

// ApprovalMergeRequestInvoker 审核合并请求
func (c *CodeHubClient) ApprovalMergeRequestInvoker(request *model.ApprovalMergeRequestRequest) *ApprovalMergeRequestInvoker {
	requestDef := GenReqDefForApprovalMergeRequest()
	return &ApprovalMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupMergeRequestApproverSetting 创建代码组合并请求审核设置
//
// 创建代码组合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateGroupMergeRequestApproverSetting(request *model.CreateGroupMergeRequestApproverSettingRequest) (*model.CreateGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupMergeRequestApproverSettingResponse), nil
	}
}

// CreateGroupMergeRequestApproverSettingInvoker 创建代码组合并请求审核设置
func (c *CodeHubClient) CreateGroupMergeRequestApproverSettingInvoker(request *model.CreateGroupMergeRequestApproverSettingRequest) *CreateGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateGroupMergeRequestApproverSetting()
	return &CreateGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequest 创建合并请求
//
// 创建合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequest(request *model.CreateMergeRequestRequest) (*model.CreateMergeRequestResponse, error) {
	requestDef := GenReqDefForCreateMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestResponse), nil
	}
}

// CreateMergeRequestInvoker 创建合并请求
func (c *CodeHubClient) CreateMergeRequestInvoker(request *model.CreateMergeRequestRequest) *CreateMergeRequestInvoker {
	requestDef := GenReqDefForCreateMergeRequest()
	return &CreateMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestApproverSetting 创建合并请求审核设置
//
// 创建合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequestApproverSetting(request *model.CreateMergeRequestApproverSettingRequest) (*model.CreateMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestApproverSettingResponse), nil
	}
}

// CreateMergeRequestApproverSettingInvoker 创建合并请求审核设置
func (c *CodeHubClient) CreateMergeRequestApproverSettingInvoker(request *model.CreateMergeRequestApproverSettingRequest) *CreateMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateMergeRequestApproverSetting()
	return &CreateMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestTemplate 创建合并请求模板
//
// 创建合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateMergeRequestTemplate(request *model.CreateMergeRequestTemplateRequest) (*model.CreateMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestTemplateResponse), nil
	}
}

// CreateMergeRequestTemplateInvoker 创建合并请求模板
func (c *CodeHubClient) CreateMergeRequestTemplateInvoker(request *model.CreateMergeRequestTemplateRequest) *CreateMergeRequestTemplateInvoker {
	requestDef := GenReqDefForCreateMergeRequestTemplate()
	return &CreateMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectMergeRequestApproverSetting 创建项目合并请求审核设置
//
// 创建项目合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectMergeRequestApproverSetting(request *model.CreateProjectMergeRequestApproverSettingRequest) (*model.CreateProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectMergeRequestApproverSettingResponse), nil
	}
}

// CreateProjectMergeRequestApproverSettingInvoker 创建项目合并请求审核设置
func (c *CodeHubClient) CreateProjectMergeRequestApproverSettingInvoker(request *model.CreateProjectMergeRequestApproverSettingRequest) *CreateProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateProjectMergeRequestApproverSetting()
	return &CreateProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupMergeRequestApproverSetting 删除代码组合并请求审核配置
//
// 删除代码组合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteGroupMergeRequestApproverSetting(request *model.DeleteGroupMergeRequestApproverSettingRequest) (*model.DeleteGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupMergeRequestApproverSettingResponse), nil
	}
}

// DeleteGroupMergeRequestApproverSettingInvoker 删除代码组合并请求审核配置
func (c *CodeHubClient) DeleteGroupMergeRequestApproverSettingInvoker(request *model.DeleteGroupMergeRequestApproverSettingRequest) *DeleteGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteGroupMergeRequestApproverSetting()
	return &DeleteGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestApproverSetting 删除合并请求审核配置
//
// 删除合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteMergeRequestApproverSetting(request *model.DeleteMergeRequestApproverSettingRequest) (*model.DeleteMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestApproverSettingResponse), nil
	}
}

// DeleteMergeRequestApproverSettingInvoker 删除合并请求审核配置
func (c *CodeHubClient) DeleteMergeRequestApproverSettingInvoker(request *model.DeleteMergeRequestApproverSettingRequest) *DeleteMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteMergeRequestApproverSetting()
	return &DeleteMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestTemplate 删除合并请求模板
//
// 删除合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteMergeRequestTemplate(request *model.DeleteMergeRequestTemplateRequest) (*model.DeleteMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestTemplateResponse), nil
	}
}

// DeleteMergeRequestTemplateInvoker 删除合并请求模板
func (c *CodeHubClient) DeleteMergeRequestTemplateInvoker(request *model.DeleteMergeRequestTemplateRequest) *DeleteMergeRequestTemplateInvoker {
	requestDef := GenReqDefForDeleteMergeRequestTemplate()
	return &DeleteMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestVote 删除合并请求打分
//
// 删除合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteMergeRequestVote(request *model.DeleteMergeRequestVoteRequest) (*model.DeleteMergeRequestVoteResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestVote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestVoteResponse), nil
	}
}

// DeleteMergeRequestVoteInvoker 删除合并请求打分
func (c *CodeHubClient) DeleteMergeRequestVoteInvoker(request *model.DeleteMergeRequestVoteRequest) *DeleteMergeRequestVoteInvoker {
	requestDef := GenReqDefForDeleteMergeRequestVote()
	return &DeleteMergeRequestVoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProjectMergeRequestApproverSetting 删除项目合并请求审核配置
//
// 删除项目合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteProjectMergeRequestApproverSetting(request *model.DeleteProjectMergeRequestApproverSettingRequest) (*model.DeleteProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectMergeRequestApproverSettingResponse), nil
	}
}

// DeleteProjectMergeRequestApproverSettingInvoker 删除项目合并请求审核配置
func (c *CodeHubClient) DeleteProjectMergeRequestApproverSettingInvoker(request *model.DeleteProjectMergeRequestApproverSettingRequest) *DeleteProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteProjectMergeRequestApproverSetting()
	return &DeleteProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportMergeRequest 导入合并请求
//
// 导入合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ImportMergeRequest(request *model.ImportMergeRequestRequest) (*model.ImportMergeRequestResponse, error) {
	requestDef := GenReqDefForImportMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportMergeRequestResponse), nil
	}
}

// ImportMergeRequestInvoker 导入合并请求
func (c *CodeHubClient) ImportMergeRequestInvoker(request *model.ImportMergeRequestRequest) *ImportMergeRequestInvoker {
	requestDef := GenReqDefForImportMergeRequest()
	return &ImportMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiscussionTemplates 获取检视意见模板列表
//
// 获取检视意见模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListDiscussionTemplates(request *model.ListDiscussionTemplatesRequest) (*model.ListDiscussionTemplatesResponse, error) {
	requestDef := GenReqDefForListDiscussionTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiscussionTemplatesResponse), nil
	}
}

// ListDiscussionTemplatesInvoker 获取检视意见模板列表
func (c *CodeHubClient) ListDiscussionTemplatesInvoker(request *model.ListDiscussionTemplatesRequest) *ListDiscussionTemplatesInvoker {
	requestDef := GenReqDefForListDiscussionTemplates()
	return &ListDiscussionTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestApproverSettings 获取代码组合并请求审核设置列表
//
// 获取代码组合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupMergeRequestApproverSettings(request *model.ListGroupMergeRequestApproverSettingsRequest) (*model.ListGroupMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestApproverSettingsResponse), nil
	}
}

// ListGroupMergeRequestApproverSettingsInvoker 获取代码组合并请求审核设置列表
func (c *CodeHubClient) ListGroupMergeRequestApproverSettingsInvoker(request *model.ListGroupMergeRequestApproverSettingsRequest) *ListGroupMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListGroupMergeRequestApproverSettings()
	return &ListGroupMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestCanBeAssignedReviewers 获取代码组检视人
//
// 获取代码组检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupMergeRequestCanBeAssignedReviewers(request *model.ListGroupMergeRequestCanBeAssignedReviewersRequest) (*model.ListGroupMergeRequestCanBeAssignedReviewersResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestCanBeAssignedReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

// ListGroupMergeRequestCanBeAssignedReviewersInvoker 获取代码组检视人
func (c *CodeHubClient) ListGroupMergeRequestCanBeAssignedReviewersInvoker(request *model.ListGroupMergeRequestCanBeAssignedReviewersRequest) *ListGroupMergeRequestCanBeAssignedReviewersInvoker {
	requestDef := GenReqDefForListGroupMergeRequestCanBeAssignedReviewers()
	return &ListGroupMergeRequestCanBeAssignedReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestValidAssignedCandidates 获取代码组审核人或合并人
//
// 获取代码组审核人或合并人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupMergeRequestValidAssignedCandidates(request *model.ListGroupMergeRequestValidAssignedCandidatesRequest) (*model.ListGroupMergeRequestValidAssignedCandidatesResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestValidAssignedCandidates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestValidAssignedCandidatesResponse), nil
	}
}

// ListGroupMergeRequestValidAssignedCandidatesInvoker 获取代码组审核人或合并人
func (c *CodeHubClient) ListGroupMergeRequestValidAssignedCandidatesInvoker(request *model.ListGroupMergeRequestValidAssignedCandidatesRequest) *ListGroupMergeRequestValidAssignedCandidatesInvoker {
	requestDef := GenReqDefForListGroupMergeRequestValidAssignedCandidates()
	return &ListGroupMergeRequestValidAssignedCandidatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestApproverSettings 获取合并请求审核设置列表
//
// 获取合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestApproverSettings(request *model.ListMergeRequestApproverSettingsRequest) (*model.ListMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestApproverSettingsResponse), nil
	}
}

// ListMergeRequestApproverSettingsInvoker 获取合并请求审核设置列表
func (c *CodeHubClient) ListMergeRequestApproverSettingsInvoker(request *model.ListMergeRequestApproverSettingsRequest) *ListMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListMergeRequestApproverSettings()
	return &ListMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestApprovers 获取合并请求审核人列表
//
// 获取合并请求审核人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestApprovers(request *model.ListMergeRequestApproversRequest) (*model.ListMergeRequestApproversResponse, error) {
	requestDef := GenReqDefForListMergeRequestApprovers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestApproversResponse), nil
	}
}

// ListMergeRequestApproversInvoker 获取合并请求审核人列表
func (c *CodeHubClient) ListMergeRequestApproversInvoker(request *model.ListMergeRequestApproversRequest) *ListMergeRequestApproversInvoker {
	requestDef := GenReqDefForListMergeRequestApprovers()
	return &ListMergeRequestApproversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestChanges 获取合并请求文件变更列表
//
// 获取合并请求文件变更列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestChanges(request *model.ListMergeRequestChangesRequest) (*model.ListMergeRequestChangesResponse, error) {
	requestDef := GenReqDefForListMergeRequestChanges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestChangesResponse), nil
	}
}

// ListMergeRequestChangesInvoker 获取合并请求文件变更列表
func (c *CodeHubClient) ListMergeRequestChangesInvoker(request *model.ListMergeRequestChangesRequest) *ListMergeRequestChangesInvoker {
	requestDef := GenReqDefForListMergeRequestChanges()
	return &ListMergeRequestChangesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestChangesTrees 获取合并请求文件变更列表树
//
// 获取合并请求文件变更列表树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestChangesTrees(request *model.ListMergeRequestChangesTreesRequest) (*model.ListMergeRequestChangesTreesResponse, error) {
	requestDef := GenReqDefForListMergeRequestChangesTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestChangesTreesResponse), nil
	}
}

// ListMergeRequestChangesTreesInvoker 获取合并请求文件变更列表树
func (c *CodeHubClient) ListMergeRequestChangesTreesInvoker(request *model.ListMergeRequestChangesTreesRequest) *ListMergeRequestChangesTreesInvoker {
	requestDef := GenReqDefForListMergeRequestChangesTrees()
	return &ListMergeRequestChangesTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestCommits 获取合并请求commit列表
//
// 获取合并请求commit列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestCommits(request *model.ListMergeRequestCommitsRequest) (*model.ListMergeRequestCommitsResponse, error) {
	requestDef := GenReqDefForListMergeRequestCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestCommitsResponse), nil
	}
}

// ListMergeRequestCommitsInvoker 获取合并请求commit列表
func (c *CodeHubClient) ListMergeRequestCommitsInvoker(request *model.ListMergeRequestCommitsRequest) *ListMergeRequestCommitsInvoker {
	requestDef := GenReqDefForListMergeRequestCommits()
	return &ListMergeRequestCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestConflictFiles 获取所有的冲突文件
//
// 获取所有的冲突文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestConflictFiles(request *model.ListMergeRequestConflictFilesRequest) (*model.ListMergeRequestConflictFilesResponse, error) {
	requestDef := GenReqDefForListMergeRequestConflictFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestConflictFilesResponse), nil
	}
}

// ListMergeRequestConflictFilesInvoker 获取所有的冲突文件
func (c *CodeHubClient) ListMergeRequestConflictFilesInvoker(request *model.ListMergeRequestConflictFilesRequest) *ListMergeRequestConflictFilesInvoker {
	requestDef := GenReqDefForListMergeRequestConflictFiles()
	return &ListMergeRequestConflictFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestReviewers 获取合并请求检视人列表
//
// 获取合并请求检视人列表
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

// ListMergeRequestReviewersInvoker 获取合并请求检视人列表
func (c *CodeHubClient) ListMergeRequestReviewersInvoker(request *model.ListMergeRequestReviewersRequest) *ListMergeRequestReviewersInvoker {
	requestDef := GenReqDefForListMergeRequestReviewers()
	return &ListMergeRequestReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestTemplates 获取合并请求模板列表
//
// 获取合并请求模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestTemplates(request *model.ListMergeRequestTemplatesRequest) (*model.ListMergeRequestTemplatesResponse, error) {
	requestDef := GenReqDefForListMergeRequestTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestTemplatesResponse), nil
	}
}

// ListMergeRequestTemplatesInvoker 获取合并请求模板列表
func (c *CodeHubClient) ListMergeRequestTemplatesInvoker(request *model.ListMergeRequestTemplatesRequest) *ListMergeRequestTemplatesInvoker {
	requestDef := GenReqDefForListMergeRequestTemplates()
	return &ListMergeRequestTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestValidAssignedCandidates 获取可选的合并请求检视人
//
// 获取可选的合并请求检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestValidAssignedCandidates(request *model.ListMergeRequestValidAssignedCandidatesRequest) (*model.ListMergeRequestValidAssignedCandidatesResponse, error) {
	requestDef := GenReqDefForListMergeRequestValidAssignedCandidates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestValidAssignedCandidatesResponse), nil
	}
}

// ListMergeRequestValidAssignedCandidatesInvoker 获取可选的合并请求检视人
func (c *CodeHubClient) ListMergeRequestValidAssignedCandidatesInvoker(request *model.ListMergeRequestValidAssignedCandidatesRequest) *ListMergeRequestValidAssignedCandidatesInvoker {
	requestDef := GenReqDefForListMergeRequestValidAssignedCandidates()
	return &ListMergeRequestValidAssignedCandidatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestApproverSettings 获取项目合并请求审核设置列表
//
// 获取项目合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectMergeRequestApproverSettings(request *model.ListProjectMergeRequestApproverSettingsRequest) (*model.ListProjectMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestApproverSettingsResponse), nil
	}
}

// ListProjectMergeRequestApproverSettingsInvoker 获取项目合并请求审核设置列表
func (c *CodeHubClient) ListProjectMergeRequestApproverSettingsInvoker(request *model.ListProjectMergeRequestApproverSettingsRequest) *ListProjectMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListProjectMergeRequestApproverSettings()
	return &ListProjectMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestCanBeAssignedReviewers 获取项目检视人
//
// 获取代码组检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectMergeRequestCanBeAssignedReviewers(request *model.ListProjectMergeRequestCanBeAssignedReviewersRequest) (*model.ListProjectMergeRequestCanBeAssignedReviewersResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

// ListProjectMergeRequestCanBeAssignedReviewersInvoker 获取项目检视人
func (c *CodeHubClient) ListProjectMergeRequestCanBeAssignedReviewersInvoker(request *model.ListProjectMergeRequestCanBeAssignedReviewersRequest) *ListProjectMergeRequestCanBeAssignedReviewersInvoker {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedReviewers()
	return &ListProjectMergeRequestCanBeAssignedReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestCanBeAssignedUsers 获取项目审核人或合并人
//
// 获取代码组审核人或合并人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectMergeRequestCanBeAssignedUsers(request *model.ListProjectMergeRequestCanBeAssignedUsersRequest) (*model.ListProjectMergeRequestCanBeAssignedUsersResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestCanBeAssignedUsersResponse), nil
	}
}

// ListProjectMergeRequestCanBeAssignedUsersInvoker 获取项目审核人或合并人
func (c *CodeHubClient) ListProjectMergeRequestCanBeAssignedUsersInvoker(request *model.ListProjectMergeRequestCanBeAssignedUsersRequest) *ListProjectMergeRequestCanBeAssignedUsersInvoker {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedUsers()
	return &ListProjectMergeRequestCanBeAssignedUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryMergeRequests 获取仓库MR列表
//
// 获取仓库MR列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryMergeRequests(request *model.ListRepositoryMergeRequestsRequest) (*model.ListRepositoryMergeRequestsResponse, error) {
	requestDef := GenReqDefForListRepositoryMergeRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryMergeRequestsResponse), nil
	}
}

// ListRepositoryMergeRequestsInvoker 获取仓库MR列表
func (c *CodeHubClient) ListRepositoryMergeRequestsInvoker(request *model.ListRepositoryMergeRequestsRequest) *ListRepositoryMergeRequestsInvoker {
	requestDef := GenReqDefForListRepositoryMergeRequests()
	return &ListRepositoryMergeRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MergeMergeRequest 合入合并请求
//
// 合入合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) MergeMergeRequest(request *model.MergeMergeRequestRequest) (*model.MergeMergeRequestResponse, error) {
	requestDef := GenReqDefForMergeMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MergeMergeRequestResponse), nil
	}
}

// MergeMergeRequestInvoker 合入合并请求
func (c *CodeHubClient) MergeMergeRequestInvoker(request *model.MergeMergeRequestRequest) *MergeMergeRequestInvoker {
	requestDef := GenReqDefForMergeMergeRequest()
	return &MergeMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebaseMergeRequestForOpenApi 变基合并请求
//
// 变基合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) RebaseMergeRequestForOpenApi(request *model.RebaseMergeRequestForOpenApiRequest) (*model.RebaseMergeRequestForOpenApiResponse, error) {
	requestDef := GenReqDefForRebaseMergeRequestForOpenApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebaseMergeRequestForOpenApiResponse), nil
	}
}

// RebaseMergeRequestForOpenApiInvoker 变基合并请求
func (c *CodeHubClient) RebaseMergeRequestForOpenApiInvoker(request *model.RebaseMergeRequestForOpenApiRequest) *RebaseMergeRequestForOpenApiInvoker {
	requestDef := GenReqDefForRebaseMergeRequestForOpenApi()
	return &RebaseMergeRequestForOpenApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResolveMergeRequestConflicts 在线解决合并请求冲突
//
// 在线解决合并请求冲突
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ResolveMergeRequestConflicts(request *model.ResolveMergeRequestConflictsRequest) (*model.ResolveMergeRequestConflictsResponse, error) {
	requestDef := GenReqDefForResolveMergeRequestConflicts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResolveMergeRequestConflictsResponse), nil
	}
}

// ResolveMergeRequestConflictsInvoker 在线解决合并请求冲突
func (c *CodeHubClient) ResolveMergeRequestConflictsInvoker(request *model.ResolveMergeRequestConflictsRequest) *ResolveMergeRequestConflictsInvoker {
	requestDef := GenReqDefForResolveMergeRequestConflicts()
	return &ResolveMergeRequestConflictsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReviewMergeRequest 检视合并请求
//
// 检视合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ReviewMergeRequest(request *model.ReviewMergeRequestRequest) (*model.ReviewMergeRequestResponse, error) {
	requestDef := GenReqDefForReviewMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReviewMergeRequestResponse), nil
	}
}

// ReviewMergeRequestInvoker 检视合并请求
func (c *CodeHubClient) ReviewMergeRequestInvoker(request *model.ReviewMergeRequestRequest) *ReviewMergeRequestInvoker {
	requestDef := GenReqDefForReviewMergeRequest()
	return &ReviewMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowActualHeadPipeline 获取合并请求关联的最新流水线
//
// 获取合并请求关联的最新流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowActualHeadPipeline(request *model.ShowActualHeadPipelineRequest) (*model.ShowActualHeadPipelineResponse, error) {
	requestDef := GenReqDefForShowActualHeadPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowActualHeadPipelineResponse), nil
	}
}

// ShowActualHeadPipelineInvoker 获取合并请求关联的最新流水线
func (c *CodeHubClient) ShowActualHeadPipelineInvoker(request *model.ShowActualHeadPipelineRequest) *ShowActualHeadPipelineInvoker {
	requestDef := GenReqDefForShowActualHeadPipeline()
	return &ShowActualHeadPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAverageEvaluation 获取合并请求的平均评价
//
// 获取合并请求的平均评价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowAverageEvaluation(request *model.ShowAverageEvaluationRequest) (*model.ShowAverageEvaluationResponse, error) {
	requestDef := GenReqDefForShowAverageEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAverageEvaluationResponse), nil
	}
}

// ShowAverageEvaluationInvoker 获取合并请求的平均评价
func (c *CodeHubClient) ShowAverageEvaluationInvoker(request *model.ShowAverageEvaluationRequest) *ShowAverageEvaluationInvoker {
	requestDef := GenReqDefForShowAverageEvaluation()
	return &ShowAverageEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranchConflict 获取分支代码冲突
//
// 获取分支代码冲突
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowBranchConflict(request *model.ShowBranchConflictRequest) (*model.ShowBranchConflictResponse, error) {
	requestDef := GenReqDefForShowBranchConflict()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchConflictResponse), nil
	}
}

// ShowBranchConflictInvoker 获取分支代码冲突
func (c *CodeHubClient) ShowBranchConflictInvoker(request *model.ShowBranchConflictRequest) *ShowBranchConflictInvoker {
	requestDef := GenReqDefForShowBranchConflict()
	return &ShowBranchConflictInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupMergeRequestSetting 获取代码组合并请求设置
//
// 获取代码组合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupMergeRequestSetting(request *model.ShowGroupMergeRequestSettingRequest) (*model.ShowGroupMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowGroupMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupMergeRequestSettingResponse), nil
	}
}

// ShowGroupMergeRequestSettingInvoker 获取代码组合并请求设置
func (c *CodeHubClient) ShowGroupMergeRequestSettingInvoker(request *model.ShowGroupMergeRequestSettingRequest) *ShowGroupMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowGroupMergeRequestSetting()
	return &ShowGroupMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestCommentsByLine 获取合并请求文件变更页单个文件下的检视意见
//
// 获取合并请求文件变更页单个文件下的检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestCommentsByLine(request *model.ShowMergeRequestCommentsByLineRequest) (*model.ShowMergeRequestCommentsByLineResponse, error) {
	requestDef := GenReqDefForShowMergeRequestCommentsByLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestCommentsByLineResponse), nil
	}
}

// ShowMergeRequestCommentsByLineInvoker 获取合并请求文件变更页单个文件下的检视意见
func (c *CodeHubClient) ShowMergeRequestCommentsByLineInvoker(request *model.ShowMergeRequestCommentsByLineRequest) *ShowMergeRequestCommentsByLineInvoker {
	requestDef := GenReqDefForShowMergeRequestCommentsByLine()
	return &ShowMergeRequestCommentsByLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestDetail 获取MR详情
//
// 获取MR详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestDetail(request *model.ShowMergeRequestDetailRequest) (*model.ShowMergeRequestDetailResponse, error) {
	requestDef := GenReqDefForShowMergeRequestDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestDetailResponse), nil
	}
}

// ShowMergeRequestDetailInvoker 获取MR详情
func (c *CodeHubClient) ShowMergeRequestDetailInvoker(request *model.ShowMergeRequestDetailRequest) *ShowMergeRequestDetailInvoker {
	requestDef := GenReqDefForShowMergeRequestDetail()
	return &ShowMergeRequestDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestSetting 获取仓库合并请求设置
//
// 获取仓库合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestSetting(request *model.ShowMergeRequestSettingRequest) (*model.ShowMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestSettingResponse), nil
	}
}

// ShowMergeRequestSettingInvoker 获取仓库合并请求设置
func (c *CodeHubClient) ShowMergeRequestSettingInvoker(request *model.ShowMergeRequestSettingRequest) *ShowMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowMergeRequestSetting()
	return &ShowMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestTemplate 获取单个合并请求模板
//
// 获取单个合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestTemplate(request *model.ShowMergeRequestTemplateRequest) (*model.ShowMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForShowMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestTemplateResponse), nil
	}
}

// ShowMergeRequestTemplateInvoker 获取单个合并请求模板
func (c *CodeHubClient) ShowMergeRequestTemplateInvoker(request *model.ShowMergeRequestTemplateRequest) *ShowMergeRequestTemplateInvoker {
	requestDef := GenReqDefForShowMergeRequestTemplate()
	return &ShowMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestVotesDetail 获取合并请求打分
//
// 获取合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeRequestVotesDetail(request *model.ShowMergeRequestVotesDetailRequest) (*model.ShowMergeRequestVotesDetailResponse, error) {
	requestDef := GenReqDefForShowMergeRequestVotesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestVotesDetailResponse), nil
	}
}

// ShowMergeRequestVotesDetailInvoker 获取合并请求打分
func (c *CodeHubClient) ShowMergeRequestVotesDetailInvoker(request *model.ShowMergeRequestVotesDetailRequest) *ShowMergeRequestVotesDetailInvoker {
	requestDef := GenReqDefForShowMergeRequestVotesDetail()
	return &ShowMergeRequestVotesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeableStateOuter 获取合并请求的可合入状态。
//
// 获取合并请求的可合入状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowMergeableStateOuter(request *model.ShowMergeableStateOuterRequest) (*model.ShowMergeableStateOuterResponse, error) {
	requestDef := GenReqDefForShowMergeableStateOuter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeableStateOuterResponse), nil
	}
}

// ShowMergeableStateOuterInvoker 获取合并请求的可合入状态。
func (c *CodeHubClient) ShowMergeableStateOuterInvoker(request *model.ShowMergeableStateOuterRequest) *ShowMergeableStateOuterInvoker {
	requestDef := GenReqDefForShowMergeableStateOuter()
	return &ShowMergeableStateOuterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectMergeRequestSetting 获取项目合并请求设置
//
// 获取项目合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectMergeRequestSetting(request *model.ShowProjectMergeRequestSettingRequest) (*model.ShowProjectMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowProjectMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectMergeRequestSettingResponse), nil
	}
}

// ShowProjectMergeRequestSettingInvoker 获取项目合并请求设置
func (c *CodeHubClient) ShowProjectMergeRequestSettingInvoker(request *model.ShowProjectMergeRequestSettingRequest) *ShowProjectMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowProjectMergeRequestSetting()
	return &ShowProjectMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupMergeRequestApproverSetting 更新代码组合并请求审核设置
//
// 更新代码组合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateGroupMergeRequestApproverSetting(request *model.UpdateGroupMergeRequestApproverSettingRequest) (*model.UpdateGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupMergeRequestApproverSettingResponse), nil
	}
}

// UpdateGroupMergeRequestApproverSettingInvoker 更新代码组合并请求审核设置
func (c *CodeHubClient) UpdateGroupMergeRequestApproverSettingInvoker(request *model.UpdateGroupMergeRequestApproverSettingRequest) *UpdateGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateGroupMergeRequestApproverSetting()
	return &UpdateGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequest 更新合并请求
//
// 更新合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequest(request *model.UpdateMergeRequestRequest) (*model.UpdateMergeRequestResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestResponse), nil
	}
}

// UpdateMergeRequestInvoker 更新合并请求
func (c *CodeHubClient) UpdateMergeRequestInvoker(request *model.UpdateMergeRequestRequest) *UpdateMergeRequestInvoker {
	requestDef := GenReqDefForUpdateMergeRequest()
	return &UpdateMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestApproverSetting 更新合并请求审核设置
//
// 更新合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestApproverSetting(request *model.UpdateMergeRequestApproverSettingRequest) (*model.UpdateMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestApproverSettingResponse), nil
	}
}

// UpdateMergeRequestApproverSettingInvoker 更新合并请求审核设置
func (c *CodeHubClient) UpdateMergeRequestApproverSettingInvoker(request *model.UpdateMergeRequestApproverSettingRequest) *UpdateMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateMergeRequestApproverSetting()
	return &UpdateMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestApprovers 更新合并请求的审核人列表
//
// 更新合并请求的审核人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestApprovers(request *model.UpdateMergeRequestApproversRequest) (*model.UpdateMergeRequestApproversResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestApprovers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestApproversResponse), nil
	}
}

// UpdateMergeRequestApproversInvoker 更新合并请求的审核人列表
func (c *CodeHubClient) UpdateMergeRequestApproversInvoker(request *model.UpdateMergeRequestApproversRequest) *UpdateMergeRequestApproversInvoker {
	requestDef := GenReqDefForUpdateMergeRequestApprovers()
	return &UpdateMergeRequestApproversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestReviewers 更新合并请求的检视人列表
//
// 更新合并请求的检视人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestReviewers(request *model.UpdateMergeRequestReviewersRequest) (*model.UpdateMergeRequestReviewersResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestReviewersResponse), nil
	}
}

// UpdateMergeRequestReviewersInvoker 更新合并请求的检视人列表
func (c *CodeHubClient) UpdateMergeRequestReviewersInvoker(request *model.UpdateMergeRequestReviewersRequest) *UpdateMergeRequestReviewersInvoker {
	requestDef := GenReqDefForUpdateMergeRequestReviewers()
	return &UpdateMergeRequestReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestSetting 更新仓库合并请求设置
//
// 更新仓库合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestSetting(request *model.UpdateMergeRequestSettingRequest) (*model.UpdateMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestSettingResponse), nil
	}
}

// UpdateMergeRequestSettingInvoker 更新仓库合并请求设置
func (c *CodeHubClient) UpdateMergeRequestSettingInvoker(request *model.UpdateMergeRequestSettingRequest) *UpdateMergeRequestSettingInvoker {
	requestDef := GenReqDefForUpdateMergeRequestSetting()
	return &UpdateMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestTemplate 更新合并请求模板
//
// 更新合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestTemplate(request *model.UpdateMergeRequestTemplateRequest) (*model.UpdateMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestTemplateResponse), nil
	}
}

// UpdateMergeRequestTemplateInvoker 更新合并请求模板
func (c *CodeHubClient) UpdateMergeRequestTemplateInvoker(request *model.UpdateMergeRequestTemplateRequest) *UpdateMergeRequestTemplateInvoker {
	requestDef := GenReqDefForUpdateMergeRequestTemplate()
	return &UpdateMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestVote 更新合并请求打分
//
// 更新合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateMergeRequestVote(request *model.UpdateMergeRequestVoteRequest) (*model.UpdateMergeRequestVoteResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestVote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestVoteResponse), nil
	}
}

// UpdateMergeRequestVoteInvoker 更新合并请求打分
func (c *CodeHubClient) UpdateMergeRequestVoteInvoker(request *model.UpdateMergeRequestVoteRequest) *UpdateMergeRequestVoteInvoker {
	requestDef := GenReqDefForUpdateMergeRequestVote()
	return &UpdateMergeRequestVoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectMergeRequestApproverSetting 更新项目合并请求审核设置
//
// 更新项目合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectMergeRequestApproverSetting(request *model.UpdateProjectMergeRequestApproverSettingRequest) (*model.UpdateProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectMergeRequestApproverSettingResponse), nil
	}
}

// UpdateProjectMergeRequestApproverSettingInvoker 更新项目合并请求审核设置
func (c *CodeHubClient) UpdateProjectMergeRequestApproverSettingInvoker(request *model.UpdateProjectMergeRequestApproverSettingRequest) *UpdateProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateProjectMergeRequestApproverSetting()
	return &UpdateProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitAssociatedMergeRequests 获取提交关联的合并请求
//
// 获取提交关联的合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListCommitAssociatedMergeRequests(request *model.ListCommitAssociatedMergeRequestsRequest) (*model.ListCommitAssociatedMergeRequestsResponse, error) {
	requestDef := GenReqDefForListCommitAssociatedMergeRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitAssociatedMergeRequestsResponse), nil
	}
}

// ListCommitAssociatedMergeRequestsInvoker 获取提交关联的合并请求
func (c *CodeHubClient) ListCommitAssociatedMergeRequestsInvoker(request *model.ListCommitAssociatedMergeRequestsRequest) *ListCommitAssociatedMergeRequestsInvoker {
	requestDef := GenReqDefForListCommitAssociatedMergeRequests()
	return &ListCommitAssociatedMergeRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryMergeRequestsStatistic 获取仓库合并请求统计数据
//
// 获取仓库合并请求统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryMergeRequestsStatistic(request *model.ShowRepositoryMergeRequestsStatisticRequest) (*model.ShowRepositoryMergeRequestsStatisticResponse, error) {
	requestDef := GenReqDefForShowRepositoryMergeRequestsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryMergeRequestsStatisticResponse), nil
	}
}

// ShowRepositoryMergeRequestsStatisticInvoker 获取仓库合并请求统计数据
func (c *CodeHubClient) ShowRepositoryMergeRequestsStatisticInvoker(request *model.ShowRepositoryMergeRequestsStatisticRequest) *ShowRepositoryMergeRequestsStatisticInvoker {
	requestDef := GenReqDefForShowRepositoryMergeRequestsStatistic()
	return &ShowRepositoryMergeRequestsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestParticipants 获取合并请求参与者
//
// 获取合并请求参与者
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListMergeRequestParticipants(request *model.ListMergeRequestParticipantsRequest) (*model.ListMergeRequestParticipantsResponse, error) {
	requestDef := GenReqDefForListMergeRequestParticipants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestParticipantsResponse), nil
	}
}

// ListMergeRequestParticipantsInvoker 获取合并请求参与者
func (c *CodeHubClient) ListMergeRequestParticipantsInvoker(request *model.ListMergeRequestParticipantsRequest) *ListMergeRequestParticipantsInvoker {
	requestDef := GenReqDefForListMergeRequestParticipants()
	return &ListMergeRequestParticipantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryPermissionInheritEnabled 查询仓库权限配置信息
//
// 查询仓库权限配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryPermissionInheritEnabled(request *model.ShowRepositoryPermissionInheritEnabledRequest) (*model.ShowRepositoryPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForShowRepositoryPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryPermissionInheritEnabledResponse), nil
	}
}

// ShowRepositoryPermissionInheritEnabledInvoker 查询仓库权限配置信息
func (c *CodeHubClient) ShowRepositoryPermissionInheritEnabledInvoker(request *model.ShowRepositoryPermissionInheritEnabledRequest) *ShowRepositoryPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForShowRepositoryPermissionInheritEnabled()
	return &ShowRepositoryPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryPermissionInheritEnabled 更新仓库权限继承配置
//
// 更新仓库权限继承配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryPermissionInheritEnabled(request *model.UpdateRepositoryPermissionInheritEnabledRequest) (*model.UpdateRepositoryPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryPermissionInheritEnabledResponse), nil
	}
}

// UpdateRepositoryPermissionInheritEnabledInvoker 更新仓库权限继承配置
func (c *CodeHubClient) UpdateRepositoryPermissionInheritEnabledInvoker(request *model.UpdateRepositoryPermissionInheritEnabledRequest) *UpdateRepositoryPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForUpdateRepositoryPermissionInheritEnabled()
	return &UpdateRepositoryPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatestPipelineJobs 获取流水线的关联的最新任务
//
// 获取流水线的关联的最新任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListLatestPipelineJobs(request *model.ListLatestPipelineJobsRequest) (*model.ListLatestPipelineJobsResponse, error) {
	requestDef := GenReqDefForListLatestPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatestPipelineJobsResponse), nil
	}
}

// ListLatestPipelineJobsInvoker 获取流水线的关联的最新任务
func (c *CodeHubClient) ListLatestPipelineJobsInvoker(request *model.ListLatestPipelineJobsRequest) *ListLatestPipelineJobsInvoker {
	requestDef := GenReqDefForListLatestPipelineJobs()
	return &ListLatestPipelineJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineJobs 获取流水线的关联的任务列表
//
// 获取流水线的关联的任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListPipelineJobs(request *model.ListPipelineJobsRequest) (*model.ListPipelineJobsResponse, error) {
	requestDef := GenReqDefForListPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineJobsResponse), nil
	}
}

// ListPipelineJobsInvoker 获取流水线的关联的任务列表
func (c *CodeHubClient) ListPipelineJobsInvoker(request *model.ListPipelineJobsRequest) *ListPipelineJobsInvoker {
	requestDef := GenReqDefForListPipelineJobs()
	return &ListPipelineJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectSettingsInheritCfg 获取项目继承设置项
//
// 获取项目继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectSettingsInheritCfg(request *model.ShowProjectSettingsInheritCfgRequest) (*model.ShowProjectSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForShowProjectSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectSettingsInheritCfgResponse), nil
	}
}

// ShowProjectSettingsInheritCfgInvoker 获取项目继承设置项
func (c *CodeHubClient) ShowProjectSettingsInheritCfgInvoker(request *model.ShowProjectSettingsInheritCfgRequest) *ShowProjectSettingsInheritCfgInvoker {
	requestDef := GenReqDefForShowProjectSettingsInheritCfg()
	return &ShowProjectSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWatermark 获取项目水印设置
//
// 获取项目水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectWatermark(request *model.ShowProjectWatermarkRequest) (*model.ShowProjectWatermarkResponse, error) {
	requestDef := GenReqDefForShowProjectWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWatermarkResponse), nil
	}
}

// ShowProjectWatermarkInvoker 获取项目水印设置
func (c *CodeHubClient) ShowProjectWatermarkInvoker(request *model.ShowProjectWatermarkRequest) *ShowProjectWatermarkInvoker {
	requestDef := GenReqDefForShowProjectWatermark()
	return &ShowProjectWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectSettingsInheritCfg 更新项目继承设置项
//
// 更新项目继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectSettingsInheritCfg(request *model.UpdateProjectSettingsInheritCfgRequest) (*model.UpdateProjectSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForUpdateProjectSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectSettingsInheritCfgResponse), nil
	}
}

// UpdateProjectSettingsInheritCfgInvoker 更新项目继承设置项
func (c *CodeHubClient) UpdateProjectSettingsInheritCfgInvoker(request *model.UpdateProjectSettingsInheritCfgRequest) *UpdateProjectSettingsInheritCfgInvoker {
	requestDef := GenReqDefForUpdateProjectSettingsInheritCfg()
	return &UpdateProjectSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectWatermark 更新项目水印设置
//
// 更新项目水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectWatermark(request *model.UpdateProjectWatermarkRequest) (*model.UpdateProjectWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateProjectWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectWatermarkResponse), nil
	}
}

// UpdateProjectWatermarkInvoker 更新项目水印设置
func (c *CodeHubClient) UpdateProjectWatermarkInvoker(request *model.UpdateProjectWatermarkRequest) *UpdateProjectWatermarkInvoker {
	requestDef := GenReqDefForUpdateProjectWatermark()
	return &UpdateProjectWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateProtectedBranch 批量创建仓库保护分支
//
// 批量创建仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchCreateProtectedBranch(request *model.BatchCreateProtectedBranchRequest) (*model.BatchCreateProtectedBranchResponse, error) {
	requestDef := GenReqDefForBatchCreateProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateProtectedBranchResponse), nil
	}
}

// BatchCreateProtectedBranchInvoker 批量创建仓库保护分支
func (c *CodeHubClient) BatchCreateProtectedBranchInvoker(request *model.BatchCreateProtectedBranchRequest) *BatchCreateProtectedBranchInvoker {
	requestDef := GenReqDefForBatchCreateProtectedBranch()
	return &BatchCreateProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteProtectedBranches 批量删除仓库保护分支
//
// 批量删除仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchDeleteProtectedBranches(request *model.BatchDeleteProtectedBranchesRequest) (*model.BatchDeleteProtectedBranchesResponse, error) {
	requestDef := GenReqDefForBatchDeleteProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProtectedBranchesResponse), nil
	}
}

// BatchDeleteProtectedBranchesInvoker 批量删除仓库保护分支
func (c *CodeHubClient) BatchDeleteProtectedBranchesInvoker(request *model.BatchDeleteProtectedBranchesRequest) *BatchDeleteProtectedBranchesInvoker {
	requestDef := GenReqDefForBatchDeleteProtectedBranches()
	return &BatchDeleteProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateProtectedBranches 批量更新仓库保护分支
//
// 批量更新仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchUpdateProtectedBranches(request *model.BatchUpdateProtectedBranchesRequest) (*model.BatchUpdateProtectedBranchesResponse, error) {
	requestDef := GenReqDefForBatchUpdateProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateProtectedBranchesResponse), nil
	}
}

// BatchUpdateProtectedBranchesInvoker 批量更新仓库保护分支
func (c *CodeHubClient) BatchUpdateProtectedBranchesInvoker(request *model.BatchUpdateProtectedBranchesRequest) *BatchUpdateProtectedBranchesInvoker {
	requestDef := GenReqDefForBatchUpdateProtectedBranches()
	return &BatchUpdateProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectProtectedBranches 创建项目下保护分支
//
// 创建项目下保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateProjectProtectedBranches(request *model.CreateProjectProtectedBranchesRequest) (*model.CreateProjectProtectedBranchesResponse, error) {
	requestDef := GenReqDefForCreateProjectProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectProtectedBranchesResponse), nil
	}
}

// CreateProjectProtectedBranchesInvoker 创建项目下保护分支
func (c *CodeHubClient) CreateProjectProtectedBranchesInvoker(request *model.CreateProjectProtectedBranchesRequest) *CreateProjectProtectedBranchesInvoker {
	requestDef := GenReqDefForCreateProjectProtectedBranches()
	return &CreateProjectProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProtectedBranch 删除仓库保护分支
//
// 删除仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteProtectedBranch(request *model.DeleteProtectedBranchRequest) (*model.DeleteProtectedBranchResponse, error) {
	requestDef := GenReqDefForDeleteProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedBranchResponse), nil
	}
}

// DeleteProtectedBranchInvoker 删除仓库保护分支
func (c *CodeHubClient) DeleteProtectedBranchInvoker(request *model.DeleteProtectedBranchRequest) *DeleteProtectedBranchInvoker {
	requestDef := GenReqDefForDeleteProtectedBranch()
	return &DeleteProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectProtectedBranches 获取项目下保护分支列表
//
// 获取项目下保护分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectProtectedBranches(request *model.ListProjectProtectedBranchesRequest) (*model.ListProjectProtectedBranchesResponse, error) {
	requestDef := GenReqDefForListProjectProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectProtectedBranchesResponse), nil
	}
}

// ListProjectProtectedBranchesInvoker 获取项目下保护分支列表
func (c *CodeHubClient) ListProjectProtectedBranchesInvoker(request *model.ListProjectProtectedBranchesRequest) *ListProjectProtectedBranchesInvoker {
	requestDef := GenReqDefForListProjectProtectedBranches()
	return &ListProjectProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedBranches 获取仓库保护分支列表
//
// 获取仓库保护分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProtectedBranches(request *model.ListProtectedBranchesRequest) (*model.ListProtectedBranchesResponse, error) {
	requestDef := GenReqDefForListProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedBranchesResponse), nil
	}
}

// ListProtectedBranchesInvoker 获取仓库保护分支列表
func (c *CodeHubClient) ListProtectedBranchesInvoker(request *model.ListProtectedBranchesRequest) *ListProtectedBranchesInvoker {
	requestDef := GenReqDefForListProtectedBranches()
	return &ListProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProtectedBranch 获取仓库保护分支
//
// 获取仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProtectedBranch(request *model.ShowProtectedBranchRequest) (*model.ShowProtectedBranchResponse, error) {
	requestDef := GenReqDefForShowProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectedBranchResponse), nil
	}
}

// ShowProtectedBranchInvoker 获取仓库保护分支
func (c *CodeHubClient) ShowProtectedBranchInvoker(request *model.ShowProtectedBranchRequest) *ShowProtectedBranchInvoker {
	requestDef := GenReqDefForShowProtectedBranch()
	return &ShowProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProtectedBranch 更新仓库保护分支
//
// 更新仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProtectedBranch(request *model.UpdateProtectedBranchRequest) (*model.UpdateProtectedBranchResponse, error) {
	requestDef := GenReqDefForUpdateProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectedBranchResponse), nil
	}
}

// UpdateProtectedBranchInvoker 更新仓库保护分支
func (c *CodeHubClient) UpdateProtectedBranchInvoker(request *model.UpdateProtectedBranchRequest) *UpdateProtectedBranchInvoker {
	requestDef := GenReqDefForUpdateProtectedBranch()
	return &UpdateProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateProtectedTags 批量创建仓库保护Tag
//
// 批量创建仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchCreateProtectedTags(request *model.BatchCreateProtectedTagsRequest) (*model.BatchCreateProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateProtectedTagsResponse), nil
	}
}

// BatchCreateProtectedTagsInvoker 批量创建仓库保护Tag
func (c *CodeHubClient) BatchCreateProtectedTagsInvoker(request *model.BatchCreateProtectedTagsRequest) *BatchCreateProtectedTagsInvoker {
	requestDef := GenReqDefForBatchCreateProtectedTags()
	return &BatchCreateProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteProtectedTags 批量删除仓库保护Tag
//
// 批量删除仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchDeleteProtectedTags(request *model.BatchDeleteProtectedTagsRequest) (*model.BatchDeleteProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProtectedTagsResponse), nil
	}
}

// BatchDeleteProtectedTagsInvoker 批量删除仓库保护Tag
func (c *CodeHubClient) BatchDeleteProtectedTagsInvoker(request *model.BatchDeleteProtectedTagsRequest) *BatchDeleteProtectedTagsInvoker {
	requestDef := GenReqDefForBatchDeleteProtectedTags()
	return &BatchDeleteProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateProtectedTags 批量更新仓库保护Tag
//
// 批量更新仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) BatchUpdateProtectedTags(request *model.BatchUpdateProtectedTagsRequest) (*model.BatchUpdateProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchUpdateProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateProtectedTagsResponse), nil
	}
}

// BatchUpdateProtectedTagsInvoker 批量更新仓库保护Tag
func (c *CodeHubClient) BatchUpdateProtectedTagsInvoker(request *model.BatchUpdateProtectedTagsRequest) *BatchUpdateProtectedTagsInvoker {
	requestDef := GenReqDefForBatchUpdateProtectedTags()
	return &BatchUpdateProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProtectedTag 删除仓库保护Tag
//
// 删除仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteProtectedTag(request *model.DeleteProtectedTagRequest) (*model.DeleteProtectedTagResponse, error) {
	requestDef := GenReqDefForDeleteProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedTagResponse), nil
	}
}

// DeleteProtectedTagInvoker 删除仓库保护Tag
func (c *CodeHubClient) DeleteProtectedTagInvoker(request *model.DeleteProtectedTagRequest) *DeleteProtectedTagInvoker {
	requestDef := GenReqDefForDeleteProtectedTag()
	return &DeleteProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedTags 获取仓库保护Tag列表
//
// 获取仓库保护Tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProtectedTags(request *model.ListProtectedTagsRequest) (*model.ListProtectedTagsResponse, error) {
	requestDef := GenReqDefForListProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedTagsResponse), nil
	}
}

// ListProtectedTagsInvoker 获取仓库保护Tag列表
func (c *CodeHubClient) ListProtectedTagsInvoker(request *model.ListProtectedTagsRequest) *ListProtectedTagsInvoker {
	requestDef := GenReqDefForListProtectedTags()
	return &ListProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProtectedTag 获取仓库保护Tag
//
// 获取仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProtectedTag(request *model.ShowProtectedTagRequest) (*model.ShowProtectedTagResponse, error) {
	requestDef := GenReqDefForShowProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectedTagResponse), nil
	}
}

// ShowProtectedTagInvoker 获取仓库保护Tag
func (c *CodeHubClient) ShowProtectedTagInvoker(request *model.ShowProtectedTagRequest) *ShowProtectedTagInvoker {
	requestDef := GenReqDefForShowProtectedTag()
	return &ShowProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProtectedTag 更新仓库保护Tag
//
// 更新仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProtectedTag(request *model.UpdateProtectedTagRequest) (*model.UpdateProtectedTagResponse, error) {
	requestDef := GenReqDefForUpdateProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectedTagResponse), nil
	}
}

// UpdateProtectedTagInvoker 更新仓库保护Tag
func (c *CodeHubClient) UpdateProtectedTagInvoker(request *model.UpdateProtectedTagRequest) *UpdateProtectedTagInvoker {
	requestDef := GenReqDefForUpdateProtectedTag()
	return &UpdateProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSubmodule 创建子模块
//
// 创建子模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddSubmodule(request *model.AddSubmoduleRequest) (*model.AddSubmoduleResponse, error) {
	requestDef := GenReqDefForAddSubmodule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubmoduleResponse), nil
	}
}

// AddSubmoduleInvoker 创建子模块
func (c *CodeHubClient) AddSubmoduleInvoker(request *model.AddSubmoduleRequest) *AddSubmoduleInvoker {
	requestDef := GenReqDefForAddSubmodule()
	return &AddSubmoduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTrustedIpAddress 添加仓库ip白名单
//
// 添加仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddTrustedIpAddress(request *model.AddTrustedIpAddressRequest) (*model.AddTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTrustedIpAddressResponse), nil
	}
}

// AddTrustedIpAddressInvoker 添加仓库ip白名单
func (c *CodeHubClient) AddTrustedIpAddressInvoker(request *model.AddTrustedIpAddressRequest) *AddTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTrustedIpAddress()
	return &AddTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRemoteMirror 将普通仓库与远程镜像关联
//
// 将普通仓库与远程镜像关联
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateRemoteMirror(request *model.AssociateRemoteMirrorRequest) (*model.AssociateRemoteMirrorResponse, error) {
	requestDef := GenReqDefForAssociateRemoteMirror()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRemoteMirrorResponse), nil
	}
}

// AssociateRemoteMirrorInvoker 将普通仓库与远程镜像关联
func (c *CodeHubClient) AssociateRemoteMirrorInvoker(request *model.AssociateRemoteMirrorRequest) *AssociateRemoteMirrorInvoker {
	requestDef := GenReqDefForAssociateRemoteMirror()
	return &AssociateRemoteMirrorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRepositoryUserGroup 关联仓库与成员组
//
// 关联仓库与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateRepositoryUserGroup(request *model.AssociateRepositoryUserGroupRequest) (*model.AssociateRepositoryUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRepositoryUserGroupResponse), nil
	}
}

// AssociateRepositoryUserGroupInvoker 关联仓库与成员组
func (c *CodeHubClient) AssociateRepositoryUserGroupInvoker(request *model.AssociateRepositoryUserGroupRequest) *AssociateRepositoryUserGroupInvoker {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()
	return &AssociateRepositoryUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDir 创建指定分支下的目录
//
// 创建指定分支下的目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateDir(request *model.CreateDirRequest) (*model.CreateDirResponse, error) {
	requestDef := GenReqDefForCreateDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDirResponse), nil
	}
}

// CreateDirInvoker 创建指定分支下的目录
func (c *CodeHubClient) CreateDirInvoker(request *model.CreateDirRequest) *CreateDirInvoker {
	requestDef := GenReqDefForCreateDir()
	return &CreateDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrustedIpAddress 删除仓库ip白名单
//
// 删除仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteTrustedIpAddress(request *model.DeleteTrustedIpAddressRequest) (*model.DeleteTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrustedIpAddressResponse), nil
	}
}

// DeleteTrustedIpAddressInvoker 删除仓库ip白名单
func (c *CodeHubClient) DeleteTrustedIpAddressInvoker(request *model.DeleteTrustedIpAddressRequest) *DeleteTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTrustedIpAddress()
	return &DeleteTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadArchive 仓库下载
//
// 仓库下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DownloadArchive(request *model.DownloadArchiveRequest) (*model.DownloadArchiveResponse, error) {
	requestDef := GenReqDefForDownloadArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadArchiveResponse), nil
	}
}

// DownloadArchiveInvoker 仓库下载
func (c *CodeHubClient) DownloadArchiveInvoker(request *model.DownloadArchiveRequest) *DownloadArchiveInvoker {
	requestDef := GenReqDefForDownloadArchive()
	return &DownloadArchiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCurrentUserRepositories 获取当前登录用户仓库
//
// 获取当前登录用户仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListCurrentUserRepositories(request *model.ListCurrentUserRepositoriesRequest) (*model.ListCurrentUserRepositoriesResponse, error) {
	requestDef := GenReqDefForListCurrentUserRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCurrentUserRepositoriesResponse), nil
	}
}

// ListCurrentUserRepositoriesInvoker 获取当前登录用户仓库
func (c *CodeHubClient) ListCurrentUserRepositoriesInvoker(request *model.ListCurrentUserRepositoriesRequest) *ListCurrentUserRepositoriesInvoker {
	requestDef := GenReqDefForListCurrentUserRepositories()
	return &ListCurrentUserRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManageableGroups 获取项目下当前用户有管理权限的代码组列表
//
// 获取项目下当前用户有管理权限的代码组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListManageableGroups(request *model.ListManageableGroupsRequest) (*model.ListManageableGroupsResponse, error) {
	requestDef := GenReqDefForListManageableGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManageableGroupsResponse), nil
	}
}

// ListManageableGroupsInvoker 获取项目下当前用户有管理权限的代码组列表
func (c *CodeHubClient) ListManageableGroupsInvoker(request *model.ListManageableGroupsRequest) *ListManageableGroupsInvoker {
	requestDef := GenReqDefForListManageableGroups()
	return &ListManageableGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPersonalRepositoryImportRecords 查看当前用户仓库导入任务列表
//
// 查看当前用户仓库导入任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListPersonalRepositoryImportRecords(request *model.ListPersonalRepositoryImportRecordsRequest) (*model.ListPersonalRepositoryImportRecordsResponse, error) {
	requestDef := GenReqDefForListPersonalRepositoryImportRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPersonalRepositoryImportRecordsResponse), nil
	}
}

// ListPersonalRepositoryImportRecordsInvoker 查看当前用户仓库导入任务列表
func (c *CodeHubClient) ListPersonalRepositoryImportRecordsInvoker(request *model.ListPersonalRepositoryImportRecordsRequest) *ListPersonalRepositoryImportRecordsInvoker {
	requestDef := GenReqDefForListPersonalRepositoryImportRecords()
	return &ListPersonalRepositoryImportRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryCommitRules 查看仓库提交规则
//
// 查看仓库提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryCommitRules(request *model.ListRepositoryCommitRulesRequest) (*model.ListRepositoryCommitRulesResponse, error) {
	requestDef := GenReqDefForListRepositoryCommitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryCommitRulesResponse), nil
	}
}

// ListRepositoryCommitRulesInvoker 查看仓库提交规则
func (c *CodeHubClient) ListRepositoryCommitRulesInvoker(request *model.ListRepositoryCommitRulesRequest) *ListRepositoryCommitRulesInvoker {
	requestDef := GenReqDefForListRepositoryCommitRules()
	return &ListRepositoryCommitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryContributors 获取仓库贡献者列表
//
// 获取仓库贡献者列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryContributors(request *model.ListRepositoryContributorsRequest) (*model.ListRepositoryContributorsResponse, error) {
	requestDef := GenReqDefForListRepositoryContributors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryContributorsResponse), nil
	}
}

// ListRepositoryContributorsInvoker 获取仓库贡献者列表
func (c *CodeHubClient) ListRepositoryContributorsInvoker(request *model.ListRepositoryContributorsRequest) *ListRepositoryContributorsInvoker {
	requestDef := GenReqDefForListRepositoryContributors()
	return &ListRepositoryContributorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryEvents 获取仓库动态
//
// 获取仓库动态（当前仅开放推送动态）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryEvents(request *model.ListRepositoryEventsRequest) (*model.ListRepositoryEventsResponse, error) {
	requestDef := GenReqDefForListRepositoryEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryEventsResponse), nil
	}
}

// ListRepositoryEventsInvoker 获取仓库动态
func (c *CodeHubClient) ListRepositoryEventsInvoker(request *model.ListRepositoryEventsRequest) *ListRepositoryEventsInvoker {
	requestDef := GenReqDefForListRepositoryEvents()
	return &ListRepositoryEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryForks 获取仓库Fork列表
//
// 获取仓库Fork列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryForks(request *model.ListRepositoryForksRequest) (*model.ListRepositoryForksResponse, error) {
	requestDef := GenReqDefForListRepositoryForks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryForksResponse), nil
	}
}

// ListRepositoryForksInvoker 获取仓库Fork列表
func (c *CodeHubClient) ListRepositoryForksInvoker(request *model.ListRepositoryForksRequest) *ListRepositoryForksInvoker {
	requestDef := GenReqDefForListRepositoryForks()
	return &ListRepositoryForksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryLanguages 获取仓库默认分支语言统计
//
// 获取仓库默认分支语言统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryLanguages(request *model.ListRepositoryLanguagesRequest) (*model.ListRepositoryLanguagesResponse, error) {
	requestDef := GenReqDefForListRepositoryLanguages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryLanguagesResponse), nil
	}
}

// ListRepositoryLanguagesInvoker 获取仓库默认分支语言统计
func (c *CodeHubClient) ListRepositoryLanguagesInvoker(request *model.ListRepositoryLanguagesRequest) *ListRepositoryLanguagesInvoker {
	requestDef := GenReqDefForListRepositoryLanguages()
	return &ListRepositoryLanguagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryTemplates 模板仓列表
//
// 模板仓列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryTemplates(request *model.ListRepositoryTemplatesRequest) (*model.ListRepositoryTemplatesResponse, error) {
	requestDef := GenReqDefForListRepositoryTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryTemplatesResponse), nil
	}
}

// ListRepositoryTemplatesInvoker 模板仓列表
func (c *CodeHubClient) ListRepositoryTemplatesInvoker(request *model.ListRepositoryTemplatesRequest) *ListRepositoryTemplatesInvoker {
	requestDef := GenReqDefForListRepositoryTemplates()
	return &ListRepositoryTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubmodules 获取仓库指定分支或者标签子模块列表
//
// 获取仓库指定分支或者标签子模块列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListSubmodules(request *model.ListSubmodulesRequest) (*model.ListSubmodulesResponse, error) {
	requestDef := GenReqDefForListSubmodules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubmodulesResponse), nil
	}
}

// ListSubmodulesInvoker 获取仓库指定分支或者标签子模块列表
func (c *CodeHubClient) ListSubmodulesInvoker(request *model.ListSubmodulesRequest) *ListSubmodulesInvoker {
	requestDef := GenReqDefForListSubmodules()
	return &ListSubmodulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrees 查看分支文件列表
//
// 查看分支文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTrees(request *model.ListTreesRequest) (*model.ListTreesResponse, error) {
	requestDef := GenReqDefForListTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTreesResponse), nil
	}
}

// ListTreesInvoker 查看分支文件列表
func (c *CodeHubClient) ListTreesInvoker(request *model.ListTreesRequest) *ListTreesInvoker {
	requestDef := GenReqDefForListTrees()
	return &ListTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrustedIpAddresses 获取仓库ip白名单
//
// 获取仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTrustedIpAddresses(request *model.ListTrustedIpAddressesRequest) (*model.ListTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrustedIpAddressesResponse), nil
	}
}

// ListTrustedIpAddressesInvoker 获取仓库ip白名单
func (c *CodeHubClient) ListTrustedIpAddressesInvoker(request *model.ListTrustedIpAddressesRequest) *ListTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTrustedIpAddresses()
	return &ListTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockRepository 锁定仓库
//
// 锁定仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) LockRepository(request *model.LockRepositoryRequest) (*model.LockRepositoryResponse, error) {
	requestDef := GenReqDefForLockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockRepositoryResponse), nil
	}
}

// LockRepositoryInvoker 锁定仓库
func (c *CodeHubClient) LockRepositoryInvoker(request *model.LockRepositoryRequest) *LockRepositoryInvoker {
	requestDef := GenReqDefForLockRepository()
	return &LockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveDeployKey 删除仓库部署秘钥
//
// 删除仓库部署秘钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) RemoveDeployKey(request *model.RemoveDeployKeyRequest) (*model.RemoveDeployKeyResponse, error) {
	requestDef := GenReqDefForRemoveDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveDeployKeyResponse), nil
	}
}

// RemoveDeployKeyInvoker 删除仓库部署秘钥
func (c *CodeHubClient) RemoveDeployKeyInvoker(request *model.RemoveDeployKeyRequest) *RemoveDeployKeyInvoker {
	requestDef := GenReqDefForRemoveDeployKey()
	return &RemoveDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlobs 获取文件内容
//
// 获取文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowBlobs(request *model.ShowBlobsRequest) (*model.ShowBlobsResponse, error) {
	requestDef := GenReqDefForShowBlobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlobsResponse), nil
	}
}

// ShowBlobsInvoker 获取文件内容
func (c *CodeHubClient) ShowBlobsInvoker(request *model.ShowBlobsRequest) *ShowBlobsInvoker {
	requestDef := GenReqDefForShowBlobs()
	return &ShowBlobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitStatistics 获取仓库指定分支的提交统计信息
//
// 获取仓库指定分支的提交统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowCommitStatistics(request *model.ShowCommitStatisticsRequest) (*model.ShowCommitStatisticsResponse, error) {
	requestDef := GenReqDefForShowCommitStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitStatisticsResponse), nil
	}
}

// ShowCommitStatisticsInvoker 获取仓库指定分支的提交统计信息
func (c *CodeHubClient) ShowCommitStatisticsInvoker(request *model.ShowCommitStatisticsRequest) *ShowCommitStatisticsInvoker {
	requestDef := GenReqDefForShowCommitStatistics()
	return &ShowCommitStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiffLines 按行数查询提交文件内容
//
// 按行数查询提交文件内容，最大显示行数为1000行
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowDiffLines(request *model.ShowDiffLinesRequest) (*model.ShowDiffLinesResponse, error) {
	requestDef := GenReqDefForShowDiffLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiffLinesResponse), nil
	}
}

// ShowDiffLinesInvoker 按行数查询提交文件内容
func (c *CodeHubClient) ShowDiffLinesInvoker(request *model.ShowDiffLinesRequest) *ShowDiffLinesInvoker {
	requestDef := GenReqDefForShowDiffLines()
	return &ShowDiffLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLastPushEventInRepository 获取仓库最近推送事件
//
// 获取仓库最近推送事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowLastPushEventInRepository(request *model.ShowLastPushEventInRepositoryRequest) (*model.ShowLastPushEventInRepositoryResponse, error) {
	requestDef := GenReqDefForShowLastPushEventInRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastPushEventInRepositoryResponse), nil
	}
}

// ShowLastPushEventInRepositoryInvoker 获取仓库最近推送事件
func (c *CodeHubClient) ShowLastPushEventInRepositoryInvoker(request *model.ShowLastPushEventInRepositoryRequest) *ShowLastPushEventInRepositoryInvoker {
	requestDef := GenReqDefForShowLastPushEventInRepository()
	return &ShowLastPushEventInRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotificationSubscription 获取仓库通知设置
//
// 获取仓库通知设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowNotificationSubscription(request *model.ShowNotificationSubscriptionRequest) (*model.ShowNotificationSubscriptionResponse, error) {
	requestDef := GenReqDefForShowNotificationSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationSubscriptionResponse), nil
	}
}

// ShowNotificationSubscriptionInvoker 获取仓库通知设置
func (c *CodeHubClient) ShowNotificationSubscriptionInvoker(request *model.ShowNotificationSubscriptionRequest) *ShowNotificationSubscriptionInvoker {
	requestDef := GenReqDefForShowNotificationSubscription()
	return &ShowNotificationSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotificationSubscriptionsStatus 获取仓库通知设置启用状态
//
// 获取仓库通知设置启用状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowNotificationSubscriptionsStatus(request *model.ShowNotificationSubscriptionsStatusRequest) (*model.ShowNotificationSubscriptionsStatusResponse, error) {
	requestDef := GenReqDefForShowNotificationSubscriptionsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationSubscriptionsStatusResponse), nil
	}
}

// ShowNotificationSubscriptionsStatusInvoker 获取仓库通知设置启用状态
func (c *CodeHubClient) ShowNotificationSubscriptionsStatusInvoker(request *model.ShowNotificationSubscriptionsStatusRequest) *ShowNotificationSubscriptionsStatusInvoker {
	requestDef := GenReqDefForShowNotificationSubscriptionsStatus()
	return &ShowNotificationSubscriptionsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRefCompare 分支、tags、提交对比
//
// 分支、tags、提交对比
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRefCompare(request *model.ShowRefCompareRequest) (*model.ShowRefCompareResponse, error) {
	requestDef := GenReqDefForShowRefCompare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefCompareResponse), nil
	}
}

// ShowRefCompareInvoker 分支、tags、提交对比
func (c *CodeHubClient) ShowRefCompareInvoker(request *model.ShowRefCompareRequest) *ShowRefCompareInvoker {
	requestDef := GenReqDefForShowRefCompare()
	return &ShowRefCompareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRemoteMirror 获取仓库镜像详情
//
// 获取仓库镜像详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRemoteMirror(request *model.ShowRemoteMirrorRequest) (*model.ShowRemoteMirrorResponse, error) {
	requestDef := GenReqDefForShowRemoteMirror()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRemoteMirrorResponse), nil
	}
}

// ShowRemoteMirrorInvoker 获取仓库镜像详情
func (c *CodeHubClient) ShowRemoteMirrorInvoker(request *model.ShowRemoteMirrorRequest) *ShowRemoteMirrorInvoker {
	requestDef := GenReqDefForShowRemoteMirror()
	return &ShowRemoteMirrorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepository 获取仓库详情
//
// 获取仓库详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepository(request *model.ShowRepositoryRequest) (*model.ShowRepositoryResponse, error) {
	requestDef := GenReqDefForShowRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryResponse), nil
	}
}

// ShowRepositoryInvoker 获取仓库详情
func (c *CodeHubClient) ShowRepositoryInvoker(request *model.ShowRepositoryRequest) *ShowRepositoryInvoker {
	requestDef := GenReqDefForShowRepository()
	return &ShowRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryGeneralCommitRule 查看仓库通用提交规则
//
// 查看仓库通用提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryGeneralCommitRule(request *model.ShowRepositoryGeneralCommitRuleRequest) (*model.ShowRepositoryGeneralCommitRuleResponse, error) {
	requestDef := GenReqDefForShowRepositoryGeneralCommitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryGeneralCommitRuleResponse), nil
	}
}

// ShowRepositoryGeneralCommitRuleInvoker 查看仓库通用提交规则
func (c *CodeHubClient) ShowRepositoryGeneralCommitRuleInvoker(request *model.ShowRepositoryGeneralCommitRuleRequest) *ShowRepositoryGeneralCommitRuleInvoker {
	requestDef := GenReqDefForShowRepositoryGeneralCommitRule()
	return &ShowRepositoryGeneralCommitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryGeneralPolicy 查看仓库通用策略
//
// 查看仓库通用策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryGeneralPolicy(request *model.ShowRepositoryGeneralPolicyRequest) (*model.ShowRepositoryGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowRepositoryGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryGeneralPolicyResponse), nil
	}
}

// ShowRepositoryGeneralPolicyInvoker 查看仓库通用策略
func (c *CodeHubClient) ShowRepositoryGeneralPolicyInvoker(request *model.ShowRepositoryGeneralPolicyRequest) *ShowRepositoryGeneralPolicyInvoker {
	requestDef := GenReqDefForShowRepositoryGeneralPolicy()
	return &ShowRepositoryGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryInheritSetting 查看仓库继承设置
//
// 查看仓库继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryInheritSetting(request *model.ShowRepositoryInheritSettingRequest) (*model.ShowRepositoryInheritSettingResponse, error) {
	requestDef := GenReqDefForShowRepositoryInheritSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryInheritSettingResponse), nil
	}
}

// ShowRepositoryInheritSettingInvoker 查看仓库继承设置
func (c *CodeHubClient) ShowRepositoryInheritSettingInvoker(request *model.ShowRepositoryInheritSettingRequest) *ShowRepositoryInheritSettingInvoker {
	requestDef := GenReqDefForShowRepositoryInheritSetting()
	return &ShowRepositoryInheritSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryInheritSettingSource 查看仓库继承设置源
//
// 查看仓库继承设置源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryInheritSettingSource(request *model.ShowRepositoryInheritSettingSourceRequest) (*model.ShowRepositoryInheritSettingSourceResponse, error) {
	requestDef := GenReqDefForShowRepositoryInheritSettingSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryInheritSettingSourceResponse), nil
	}
}

// ShowRepositoryInheritSettingSourceInvoker 查看仓库继承设置源
func (c *CodeHubClient) ShowRepositoryInheritSettingSourceInvoker(request *model.ShowRepositoryInheritSettingSourceRequest) *ShowRepositoryInheritSettingSourceInvoker {
	requestDef := GenReqDefForShowRepositoryInheritSettingSource()
	return &ShowRepositoryInheritSettingSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatisticsStatus 获取仓库统计任务状态
//
// 获取仓库统计任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryStatisticsStatus(request *model.ShowRepositoryStatisticsStatusRequest) (*model.ShowRepositoryStatisticsStatusResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatisticsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsStatusResponse), nil
	}
}

// ShowRepositoryStatisticsStatusInvoker 获取仓库统计任务状态
func (c *CodeHubClient) ShowRepositoryStatisticsStatusInvoker(request *model.ShowRepositoryStatisticsStatusRequest) *ShowRepositoryStatisticsStatusInvoker {
	requestDef := GenReqDefForShowRepositoryStatisticsStatus()
	return &ShowRepositoryStatisticsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatisticsSummary 获取仓库统计摘要
//
// 获取仓库统计摘要
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryStatisticsSummary(request *model.ShowRepositoryStatisticsSummaryRequest) (*model.ShowRepositoryStatisticsSummaryResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatisticsSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsSummaryResponse), nil
	}
}

// ShowRepositoryStatisticsSummaryInvoker 获取仓库统计摘要
func (c *CodeHubClient) ShowRepositoryStatisticsSummaryInvoker(request *model.ShowRepositoryStatisticsSummaryRequest) *ShowRepositoryStatisticsSummaryInvoker {
	requestDef := GenReqDefForShowRepositoryStatisticsSummary()
	return &ShowRepositoryStatisticsSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWatermark 获取仓库水印设置
//
// 获取仓库水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryWatermark(request *model.ShowRepositoryWatermarkRequest) (*model.ShowRepositoryWatermarkResponse, error) {
	requestDef := GenReqDefForShowRepositoryWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWatermarkResponse), nil
	}
}

// ShowRepositoryWatermarkInvoker 获取仓库水印设置
func (c *CodeHubClient) ShowRepositoryWatermarkInvoker(request *model.ShowRepositoryWatermarkRequest) *ShowRepositoryWatermarkInvoker {
	requestDef := GenReqDefForShowRepositoryWatermark()
	return &ShowRepositoryWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserRefPermission 获取CR仓库用户分支或标签级权限
//
// 获取CR仓库用户分支或标签级权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowUserRefPermission(request *model.ShowUserRefPermissionRequest) (*model.ShowUserRefPermissionResponse, error) {
	requestDef := GenReqDefForShowUserRefPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserRefPermissionResponse), nil
	}
}

// ShowUserRefPermissionInvoker 获取CR仓库用户分支或标签级权限
func (c *CodeHubClient) ShowUserRefPermissionInvoker(request *model.ShowUserRefPermissionRequest) *ShowUserRefPermissionInvoker {
	requestDef := GenReqDefForShowUserRefPermission()
	return &ShowUserRefPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartRemoteMirrorSynchronization 启动仓库镜像同步
//
// 启动仓库镜像同步
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) StartRemoteMirrorSynchronization(request *model.StartRemoteMirrorSynchronizationRequest) (*model.StartRemoteMirrorSynchronizationResponse, error) {
	requestDef := GenReqDefForStartRemoteMirrorSynchronization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartRemoteMirrorSynchronizationResponse), nil
	}
}

// StartRemoteMirrorSynchronizationInvoker 启动仓库镜像同步
func (c *CodeHubClient) StartRemoteMirrorSynchronizationInvoker(request *model.StartRemoteMirrorSynchronizationRequest) *StartRemoteMirrorSynchronizationInvoker {
	requestDef := GenReqDefForStartRemoteMirrorSynchronization()
	return &StartRemoteMirrorSynchronizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockRepository 解锁仓库
//
// 解锁仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UnlockRepository(request *model.UnlockRepositoryRequest) (*model.UnlockRepositoryResponse, error) {
	requestDef := GenReqDefForUnlockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockRepositoryResponse), nil
	}
}

// UnlockRepositoryInvoker 解锁仓库
func (c *CodeHubClient) UnlockRepositoryInvoker(request *model.UnlockRepositoryRequest) *UnlockRepositoryInvoker {
	requestDef := GenReqDefForUnlockRepository()
	return &UnlockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotificationSubscription 修改仓库通知设置
//
// 修改仓库通知设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateNotificationSubscription(request *model.UpdateNotificationSubscriptionRequest) (*model.UpdateNotificationSubscriptionResponse, error) {
	requestDef := GenReqDefForUpdateNotificationSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationSubscriptionResponse), nil
	}
}

// UpdateNotificationSubscriptionInvoker 修改仓库通知设置
func (c *CodeHubClient) UpdateNotificationSubscriptionInvoker(request *model.UpdateNotificationSubscriptionRequest) *UpdateNotificationSubscriptionInvoker {
	requestDef := GenReqDefForUpdateNotificationSubscription()
	return &UpdateNotificationSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryGeneralPolicy 修改仓库通用策略
//
// 修改仓库通用策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryGeneralPolicy(request *model.UpdateRepositoryGeneralPolicyRequest) (*model.UpdateRepositoryGeneralPolicyResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryGeneralPolicyResponse), nil
	}
}

// UpdateRepositoryGeneralPolicyInvoker 修改仓库通用策略
func (c *CodeHubClient) UpdateRepositoryGeneralPolicyInvoker(request *model.UpdateRepositoryGeneralPolicyRequest) *UpdateRepositoryGeneralPolicyInvoker {
	requestDef := GenReqDefForUpdateRepositoryGeneralPolicy()
	return &UpdateRepositoryGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryInheritSetting 修改仓库继承设置
//
// 修改仓库继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryInheritSetting(request *model.UpdateRepositoryInheritSettingRequest) (*model.UpdateRepositoryInheritSettingResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryInheritSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryInheritSettingResponse), nil
	}
}

// UpdateRepositoryInheritSettingInvoker 修改仓库继承设置
func (c *CodeHubClient) UpdateRepositoryInheritSettingInvoker(request *model.UpdateRepositoryInheritSettingRequest) *UpdateRepositoryInheritSettingInvoker {
	requestDef := GenReqDefForUpdateRepositoryInheritSetting()
	return &UpdateRepositoryInheritSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryWatermark 更新仓库水印设置
//
// 更新仓库水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryWatermark(request *model.UpdateRepositoryWatermarkRequest) (*model.UpdateRepositoryWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryWatermarkResponse), nil
	}
}

// UpdateRepositoryWatermarkInvoker 更新仓库水印设置
func (c *CodeHubClient) UpdateRepositoryWatermarkInvoker(request *model.UpdateRepositoryWatermarkRequest) *UpdateRepositoryWatermarkInvoker {
	requestDef := GenReqDefForUpdateRepositoryWatermark()
	return &UpdateRepositoryWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrustedIpAddress 修改仓库ip白名单
//
// 修改仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateTrustedIpAddress(request *model.UpdateTrustedIpAddressRequest) (*model.UpdateTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrustedIpAddressResponse), nil
	}
}

// UpdateTrustedIpAddressInvoker 修改仓库ip白名单
func (c *CodeHubClient) UpdateTrustedIpAddressInvoker(request *model.UpdateTrustedIpAddressRequest) *UpdateTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTrustedIpAddress()
	return &UpdateTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 创建标签
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 创建标签
func (c *CodeHubClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除标签
func (c *CodeHubClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 获取标签列表
//
// 获取标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 获取标签列表
func (c *CodeHubClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTag 查看标签详情
//
// 查看标签详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowTag(request *model.ShowTagRequest) (*model.ShowTagResponse, error) {
	requestDef := GenReqDefForShowTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagResponse), nil
	}
}

// ShowTagInvoker 查看标签详情
func (c *CodeHubClient) ShowTagInvoker(request *model.ShowTagRequest) *ShowTagInvoker {
	requestDef := GenReqDefForShowTag()
	return &ShowTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTenantTrustedIpAddress 添加租户ip白名单
//
// 添加租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddTenantTrustedIpAddress(request *model.AddTenantTrustedIpAddressRequest) (*model.AddTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTenantTrustedIpAddressResponse), nil
	}
}

// AddTenantTrustedIpAddressInvoker 添加租户ip白名单
func (c *CodeHubClient) AddTenantTrustedIpAddressInvoker(request *model.AddTenantTrustedIpAddressRequest) *AddTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()
	return &AddTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTenantTrustedIpAddress 删除租户ip白名单
//
// 删除租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteTenantTrustedIpAddress(request *model.DeleteTenantTrustedIpAddressRequest) (*model.DeleteTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTenantTrustedIpAddressResponse), nil
	}
}

// DeleteTenantTrustedIpAddressInvoker 删除租户ip白名单
func (c *CodeHubClient) DeleteTenantTrustedIpAddressInvoker(request *model.DeleteTenantTrustedIpAddressRequest) *DeleteTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()
	return &DeleteTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantTrustedIpAddresses 获取租户ip白名单
//
// 获取租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTenantTrustedIpAddresses(request *model.ListTenantTrustedIpAddressesRequest) (*model.ListTenantTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantTrustedIpAddressesResponse), nil
	}
}

// ListTenantTrustedIpAddressesInvoker 获取租户ip白名单
func (c *CodeHubClient) ListTenantTrustedIpAddressesInvoker(request *model.ListTenantTrustedIpAddressesRequest) *ListTenantTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()
	return &ListTenantTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantTrustedIpAddress 修改租户ip白名单
//
// 修改租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateTenantTrustedIpAddress(request *model.UpdateTenantTrustedIpAddressRequest) (*model.UpdateTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantTrustedIpAddressResponse), nil
	}
}

// UpdateTenantTrustedIpAddressInvoker 修改租户ip白名单
func (c *CodeHubClient) UpdateTenantTrustedIpAddressInvoker(request *model.UpdateTenantTrustedIpAddressRequest) *UpdateTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()
	return &UpdateTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDeployKey 校验部署秘钥在上层代码组或项目是否配置
//
// 校验部署秘钥在上层代码组或项目是否配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CheckDeployKey(request *model.CheckDeployKeyRequest) (*model.CheckDeployKeyResponse, error) {
	requestDef := GenReqDefForCheckDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDeployKeyResponse), nil
	}
}

// CheckDeployKeyInvoker 校验部署秘钥在上层代码组或项目是否配置
func (c *CodeHubClient) CheckDeployKeyInvoker(request *model.CheckDeployKeyRequest) *CheckDeployKeyInvoker {
	requestDef := GenReqDefForCheckDeployKey()
	return &CheckDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckGroupDeployKey 校验代码组部署秘钥在上层代码组或项目是否配置
//
// 校验代码组部署秘钥在上层代码组或项目是否配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CheckGroupDeployKey(request *model.CheckGroupDeployKeyRequest) (*model.CheckGroupDeployKeyResponse, error) {
	requestDef := GenReqDefForCheckGroupDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckGroupDeployKeyResponse), nil
	}
}

// CheckGroupDeployKeyInvoker 校验代码组部署秘钥在上层代码组或项目是否配置
func (c *CodeHubClient) CheckGroupDeployKeyInvoker(request *model.CheckGroupDeployKeyRequest) *CheckGroupDeployKeyInvoker {
	requestDef := GenReqDefForCheckGroupDeployKey()
	return &CheckGroupDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranchRelatedWorkItems 获取仓库下指定分支的关联工作项列表
//
// 获取仓库下指定分支的关联工作项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListBranchRelatedWorkItems(request *model.ListBranchRelatedWorkItemsRequest) (*model.ListBranchRelatedWorkItemsResponse, error) {
	requestDef := GenReqDefForListBranchRelatedWorkItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchRelatedWorkItemsResponse), nil
	}
}

// ListBranchRelatedWorkItemsInvoker 获取仓库下指定分支的关联工作项列表
func (c *CodeHubClient) ListBranchRelatedWorkItemsInvoker(request *model.ListBranchRelatedWorkItemsRequest) *ListBranchRelatedWorkItemsInvoker {
	requestDef := GenReqDefForListBranchRelatedWorkItems()
	return &ListBranchRelatedWorkItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupDeployKeys 获取代码组下部署密钥列表
//
// 获取代码组下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupDeployKeys(request *model.ListGroupDeployKeysRequest) (*model.ListGroupDeployKeysResponse, error) {
	requestDef := GenReqDefForListGroupDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupDeployKeysResponse), nil
	}
}

// ListGroupDeployKeysInvoker 获取代码组下部署密钥列表
func (c *CodeHubClient) ListGroupDeployKeysInvoker(request *model.ListGroupDeployKeysRequest) *ListGroupDeployKeysInvoker {
	requestDef := GenReqDefForListGroupDeployKeys()
	return &ListGroupDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectDeployKeys 获取项目下部署密钥列表
//
// 获取项目下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectDeployKeys(request *model.ListProjectDeployKeysRequest) (*model.ListProjectDeployKeysResponse, error) {
	requestDef := GenReqDefForListProjectDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectDeployKeysResponse), nil
	}
}

// ListProjectDeployKeysInvoker 获取项目下部署密钥列表
func (c *CodeHubClient) ListProjectDeployKeysInvoker(request *model.ListProjectDeployKeysRequest) *ListProjectDeployKeysInvoker {
	requestDef := GenReqDefForListProjectDeployKeys()
	return &ListProjectDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryDeployKeys 获取仓库下部署密钥列表
//
// 获取仓库下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryDeployKeys(request *model.ListRepositoryDeployKeysRequest) (*model.ListRepositoryDeployKeysResponse, error) {
	requestDef := GenReqDefForListRepositoryDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryDeployKeysResponse), nil
	}
}

// ListRepositoryDeployKeysInvoker 获取仓库下部署密钥列表
func (c *CodeHubClient) ListRepositoryDeployKeysInvoker(request *model.ListRepositoryDeployKeysRequest) *ListRepositoryDeployKeysInvoker {
	requestDef := GenReqDefForListRepositoryDeployKeys()
	return &ListRepositoryDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWorkItems 获取仓库下工作项列表
//
// 获取仓库下工作项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryWorkItems(request *model.ListRepositoryWorkItemsRequest) (*model.ListRepositoryWorkItemsResponse, error) {
	requestDef := GenReqDefForListRepositoryWorkItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWorkItemsResponse), nil
	}
}

// ListRepositoryWorkItemsInvoker 获取仓库下工作项列表
func (c *CodeHubClient) ListRepositoryWorkItemsInvoker(request *model.ListRepositoryWorkItemsRequest) *ListRepositoryWorkItemsInvoker {
	requestDef := GenReqDefForListRepositoryWorkItems()
	return &ListRepositoryWorkItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSshKey 添加ssh公钥
//
// 添加ssh公钥
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

// AddSshKeyInvoker 添加ssh公钥
func (c *CodeHubClient) AddSshKeyInvoker(request *model.AddSshKeyRequest) *AddSshKeyInvoker {
	requestDef := GenReqDefForAddSshKey()
	return &AddSshKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSshKey 删除ssh公钥
//
// 删除ssh公钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteSshKey(request *model.DeleteSshKeyRequest) (*model.DeleteSshKeyResponse, error) {
	requestDef := GenReqDefForDeleteSshKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSshKeyResponse), nil
	}
}

// DeleteSshKeyInvoker 删除ssh公钥
func (c *CodeHubClient) DeleteSshKeyInvoker(request *model.DeleteSshKeyRequest) *DeleteSshKeyInvoker {
	requestDef := GenReqDefForDeleteSshKey()
	return &DeleteSshKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserGpgKeys 获取当前用户的gpg_key列表
//
// 获取当前用户的gpg_key列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListUserGpgKeys(request *model.ListUserGpgKeysRequest) (*model.ListUserGpgKeysResponse, error) {
	requestDef := GenReqDefForListUserGpgKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserGpgKeysResponse), nil
	}
}

// ListUserGpgKeysInvoker 获取当前用户的gpg_key列表
func (c *CodeHubClient) ListUserGpgKeysInvoker(request *model.ListUserGpgKeysRequest) *ListUserGpgKeysInvoker {
	requestDef := GenReqDefForListUserGpgKeys()
	return &ListUserGpgKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserKeys 获取当前用户的秘钥列表
//
// 获取当前用户的秘钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListUserKeys(request *model.ListUserKeysRequest) (*model.ListUserKeysResponse, error) {
	requestDef := GenReqDefForListUserKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserKeysResponse), nil
	}
}

// ListUserKeysInvoker 获取当前用户的秘钥列表
func (c *CodeHubClient) ListUserKeysInvoker(request *model.ListUserKeysRequest) *ListUserKeysInvoker {
	requestDef := GenReqDefForListUserKeys()
	return &ListUserKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendUserEmailVerifyCode 发送邮箱验证码
//
// 发送邮箱验证码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) SendUserEmailVerifyCode(request *model.SendUserEmailVerifyCodeRequest) (*model.SendUserEmailVerifyCodeResponse, error) {
	requestDef := GenReqDefForSendUserEmailVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendUserEmailVerifyCodeResponse), nil
	}
}

// SendUserEmailVerifyCodeInvoker 发送邮箱验证码
func (c *CodeHubClient) SendUserEmailVerifyCodeInvoker(request *model.SendUserEmailVerifyCodeRequest) *SendUserEmailVerifyCodeInvoker {
	requestDef := GenReqDefForSendUserEmailVerifyCode()
	return &SendUserEmailVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpsPasswordSetting 获取https的验证方式
//
// 获取https的验证方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowHttpsPasswordSetting(request *model.ShowHttpsPasswordSettingRequest) (*model.ShowHttpsPasswordSettingResponse, error) {
	requestDef := GenReqDefForShowHttpsPasswordSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpsPasswordSettingResponse), nil
	}
}

// ShowHttpsPasswordSettingInvoker 获取https的验证方式
func (c *CodeHubClient) ShowHttpsPasswordSettingInvoker(request *model.ShowHttpsPasswordSettingRequest) *ShowHttpsPasswordSettingInvoker {
	requestDef := GenReqDefForShowHttpsPasswordSetting()
	return &ShowHttpsPasswordSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserEmails 获取用户相关邮箱信息
//
// 获取用户相关邮箱信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowUserEmails(request *model.ShowUserEmailsRequest) (*model.ShowUserEmailsResponse, error) {
	requestDef := GenReqDefForShowUserEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserEmailsResponse), nil
	}
}

// ShowUserEmailsInvoker 获取用户相关邮箱信息
func (c *CodeHubClient) ShowUserEmailsInvoker(request *model.ShowUserEmailsRequest) *ShowUserEmailsInvoker {
	requestDef := GenReqDefForShowUserEmails()
	return &ShowUserEmailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpsPasswordSetting 修改https的验证方式
//
// 修改https的验证方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateHttpsPasswordSetting(request *model.UpdateHttpsPasswordSettingRequest) (*model.UpdateHttpsPasswordSettingResponse, error) {
	requestDef := GenReqDefForUpdateHttpsPasswordSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpsPasswordSettingResponse), nil
	}
}

// UpdateHttpsPasswordSettingInvoker 修改https的验证方式
func (c *CodeHubClient) UpdateHttpsPasswordSettingInvoker(request *model.UpdateHttpsPasswordSettingRequest) *UpdateHttpsPasswordSettingInvoker {
	requestDef := GenReqDefForUpdateHttpsPasswordSetting()
	return &UpdateHttpsPasswordSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserEmails 更新邮箱
//
// 更新邮箱
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateUserEmails(request *model.UpdateUserEmailsRequest) (*model.UpdateUserEmailsResponse, error) {
	requestDef := GenReqDefForUpdateUserEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserEmailsResponse), nil
	}
}

// UpdateUserEmailsInvoker 更新邮箱
func (c *CodeHubClient) UpdateUserEmailsInvoker(request *model.UpdateUserEmailsRequest) *UpdateUserEmailsInvoker {
	requestDef := GenReqDefForUpdateUserEmails()
	return &UpdateUserEmailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGroupWebhook 添加代码组下Webhook
//
// 添加代码组下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddGroupWebhook(request *model.AddGroupWebhookRequest) (*model.AddGroupWebhookResponse, error) {
	requestDef := GenReqDefForAddGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGroupWebhookResponse), nil
	}
}

// AddGroupWebhookInvoker 添加代码组下Webhook
func (c *CodeHubClient) AddGroupWebhookInvoker(request *model.AddGroupWebhookRequest) *AddGroupWebhookInvoker {
	requestDef := GenReqDefForAddGroupWebhook()
	return &AddGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProjectWebhook 添加项目下Webhook
//
// 添加项目下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddProjectWebhook(request *model.AddProjectWebhookRequest) (*model.AddProjectWebhookResponse, error) {
	requestDef := GenReqDefForAddProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProjectWebhookResponse), nil
	}
}

// AddProjectWebhookInvoker 添加项目下Webhook
func (c *CodeHubClient) AddProjectWebhookInvoker(request *model.AddProjectWebhookRequest) *AddProjectWebhookInvoker {
	requestDef := GenReqDefForAddProjectWebhook()
	return &AddProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRepositoryWebhook 添加仓库下Webhook
//
// 添加仓库下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddRepositoryWebhook(request *model.AddRepositoryWebhookRequest) (*model.AddRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForAddRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepositoryWebhookResponse), nil
	}
}

// AddRepositoryWebhookInvoker 添加仓库下Webhook
func (c *CodeHubClient) AddRepositoryWebhookInvoker(request *model.AddRepositoryWebhookRequest) *AddRepositoryWebhookInvoker {
	requestDef := GenReqDefForAddRepositoryWebhook()
	return &AddRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupWebhookLogs 获取代码组下指定Webhook的日志列表
//
// 获取代码组下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupWebhookLogs(request *model.ListGroupWebhookLogsRequest) (*model.ListGroupWebhookLogsResponse, error) {
	requestDef := GenReqDefForListGroupWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupWebhookLogsResponse), nil
	}
}

// ListGroupWebhookLogsInvoker 获取代码组下指定Webhook的日志列表
func (c *CodeHubClient) ListGroupWebhookLogsInvoker(request *model.ListGroupWebhookLogsRequest) *ListGroupWebhookLogsInvoker {
	requestDef := GenReqDefForListGroupWebhookLogs()
	return &ListGroupWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupWebhooks 获取代码组下Webhook列表
//
// 获取代码组下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListGroupWebhooks(request *model.ListGroupWebhooksRequest) (*model.ListGroupWebhooksResponse, error) {
	requestDef := GenReqDefForListGroupWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupWebhooksResponse), nil
	}
}

// ListGroupWebhooksInvoker 获取代码组下Webhook列表
func (c *CodeHubClient) ListGroupWebhooksInvoker(request *model.ListGroupWebhooksRequest) *ListGroupWebhooksInvoker {
	requestDef := GenReqDefForListGroupWebhooks()
	return &ListGroupWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectWebhookLogs 获取项目下指定Webhook的日志列表
//
// 获取项目下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectWebhookLogs(request *model.ListProjectWebhookLogsRequest) (*model.ListProjectWebhookLogsResponse, error) {
	requestDef := GenReqDefForListProjectWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectWebhookLogsResponse), nil
	}
}

// ListProjectWebhookLogsInvoker 获取项目下指定Webhook的日志列表
func (c *CodeHubClient) ListProjectWebhookLogsInvoker(request *model.ListProjectWebhookLogsRequest) *ListProjectWebhookLogsInvoker {
	requestDef := GenReqDefForListProjectWebhookLogs()
	return &ListProjectWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectWebhooks 获取项目下Webhook列表
//
// 获取项目下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListProjectWebhooks(request *model.ListProjectWebhooksRequest) (*model.ListProjectWebhooksResponse, error) {
	requestDef := GenReqDefForListProjectWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectWebhooksResponse), nil
	}
}

// ListProjectWebhooksInvoker 获取项目下Webhook列表
func (c *CodeHubClient) ListProjectWebhooksInvoker(request *model.ListProjectWebhooksRequest) *ListProjectWebhooksInvoker {
	requestDef := GenReqDefForListProjectWebhooks()
	return &ListProjectWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWebhookLogs 获取仓库下指定Webhook的日志列表
//
// 获取仓库下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryWebhookLogs(request *model.ListRepositoryWebhookLogsRequest) (*model.ListRepositoryWebhookLogsResponse, error) {
	requestDef := GenReqDefForListRepositoryWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWebhookLogsResponse), nil
	}
}

// ListRepositoryWebhookLogsInvoker 获取仓库下指定Webhook的日志列表
func (c *CodeHubClient) ListRepositoryWebhookLogsInvoker(request *model.ListRepositoryWebhookLogsRequest) *ListRepositoryWebhookLogsInvoker {
	requestDef := GenReqDefForListRepositoryWebhookLogs()
	return &ListRepositoryWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWebhooks 获取仓库下Webhook列表
//
// 获取仓库下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListRepositoryWebhooks(request *model.ListRepositoryWebhooksRequest) (*model.ListRepositoryWebhooksResponse, error) {
	requestDef := GenReqDefForListRepositoryWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWebhooksResponse), nil
	}
}

// ListRepositoryWebhooksInvoker 获取仓库下Webhook列表
func (c *CodeHubClient) ListRepositoryWebhooksInvoker(request *model.ListRepositoryWebhooksRequest) *ListRepositoryWebhooksInvoker {
	requestDef := GenReqDefForListRepositoryWebhooks()
	return &ListRepositoryWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveGroupWebhook 删除代码组下单个Webhook
//
// 删除代码组下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) RemoveGroupWebhook(request *model.RemoveGroupWebhookRequest) (*model.RemoveGroupWebhookResponse, error) {
	requestDef := GenReqDefForRemoveGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveGroupWebhookResponse), nil
	}
}

// RemoveGroupWebhookInvoker 删除代码组下单个Webhook
func (c *CodeHubClient) RemoveGroupWebhookInvoker(request *model.RemoveGroupWebhookRequest) *RemoveGroupWebhookInvoker {
	requestDef := GenReqDefForRemoveGroupWebhook()
	return &RemoveGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveProjectWebhook 删除项目下单个Webhook
//
// 删除项目下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) RemoveProjectWebhook(request *model.RemoveProjectWebhookRequest) (*model.RemoveProjectWebhookResponse, error) {
	requestDef := GenReqDefForRemoveProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveProjectWebhookResponse), nil
	}
}

// RemoveProjectWebhookInvoker 删除项目下单个Webhook
func (c *CodeHubClient) RemoveProjectWebhookInvoker(request *model.RemoveProjectWebhookRequest) *RemoveProjectWebhookInvoker {
	requestDef := GenReqDefForRemoveProjectWebhook()
	return &RemoveProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveRepositoryWebhook 删除仓库下单个Webhook
//
// 删除仓库下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) RemoveRepositoryWebhook(request *model.RemoveRepositoryWebhookRequest) (*model.RemoveRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForRemoveRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveRepositoryWebhookResponse), nil
	}
}

// RemoveRepositoryWebhookInvoker 删除仓库下单个Webhook
func (c *CodeHubClient) RemoveRepositoryWebhookInvoker(request *model.RemoveRepositoryWebhookRequest) *RemoveRepositoryWebhookInvoker {
	requestDef := GenReqDefForRemoveRepositoryWebhook()
	return &RemoveRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWebhook 获取代码组下单个Webhook数据
//
// 获取代码组下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupWebhook(request *model.ShowGroupWebhookRequest) (*model.ShowGroupWebhookResponse, error) {
	requestDef := GenReqDefForShowGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWebhookResponse), nil
	}
}

// ShowGroupWebhookInvoker 获取代码组下单个Webhook数据
func (c *CodeHubClient) ShowGroupWebhookInvoker(request *model.ShowGroupWebhookRequest) *ShowGroupWebhookInvoker {
	requestDef := GenReqDefForShowGroupWebhook()
	return &ShowGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWebhookLog 获取代码组下指定Webhook的指定日志详情
//
// 获取代码组下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroupWebhookLog(request *model.ShowGroupWebhookLogRequest) (*model.ShowGroupWebhookLogResponse, error) {
	requestDef := GenReqDefForShowGroupWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWebhookLogResponse), nil
	}
}

// ShowGroupWebhookLogInvoker 获取代码组下指定Webhook的指定日志详情
func (c *CodeHubClient) ShowGroupWebhookLogInvoker(request *model.ShowGroupWebhookLogRequest) *ShowGroupWebhookLogInvoker {
	requestDef := GenReqDefForShowGroupWebhookLog()
	return &ShowGroupWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWebhook 获取项目下单个Webhook数据
//
// 获取项目下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectWebhook(request *model.ShowProjectWebhookRequest) (*model.ShowProjectWebhookResponse, error) {
	requestDef := GenReqDefForShowProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWebhookResponse), nil
	}
}

// ShowProjectWebhookInvoker 获取项目下单个Webhook数据
func (c *CodeHubClient) ShowProjectWebhookInvoker(request *model.ShowProjectWebhookRequest) *ShowProjectWebhookInvoker {
	requestDef := GenReqDefForShowProjectWebhook()
	return &ShowProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWebhookLog 获取项目下指定Webhook的指定日志详情
//
// 获取项目下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowProjectWebhookLog(request *model.ShowProjectWebhookLogRequest) (*model.ShowProjectWebhookLogResponse, error) {
	requestDef := GenReqDefForShowProjectWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWebhookLogResponse), nil
	}
}

// ShowProjectWebhookLogInvoker 获取项目下指定Webhook的指定日志详情
func (c *CodeHubClient) ShowProjectWebhookLogInvoker(request *model.ShowProjectWebhookLogRequest) *ShowProjectWebhookLogInvoker {
	requestDef := GenReqDefForShowProjectWebhookLog()
	return &ShowProjectWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWebhook 获取仓库下单个Webhook数据
//
// 获取仓库下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryWebhook(request *model.ShowRepositoryWebhookRequest) (*model.ShowRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForShowRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWebhookResponse), nil
	}
}

// ShowRepositoryWebhookInvoker 获取仓库下单个Webhook数据
func (c *CodeHubClient) ShowRepositoryWebhookInvoker(request *model.ShowRepositoryWebhookRequest) *ShowRepositoryWebhookInvoker {
	requestDef := GenReqDefForShowRepositoryWebhook()
	return &ShowRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWebhookLog 获取仓库下指定Webhook的指定日志详情
//
// 获取仓库下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowRepositoryWebhookLog(request *model.ShowRepositoryWebhookLogRequest) (*model.ShowRepositoryWebhookLogResponse, error) {
	requestDef := GenReqDefForShowRepositoryWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWebhookLogResponse), nil
	}
}

// ShowRepositoryWebhookLogInvoker 获取仓库下指定Webhook的指定日志详情
func (c *CodeHubClient) ShowRepositoryWebhookLogInvoker(request *model.ShowRepositoryWebhookLogRequest) *ShowRepositoryWebhookLogInvoker {
	requestDef := GenReqDefForShowRepositoryWebhookLog()
	return &ShowRepositoryWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupWebhook 更新代码组下单个Webhook数据
//
// 更新代码组下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateGroupWebhook(request *model.UpdateGroupWebhookRequest) (*model.UpdateGroupWebhookResponse, error) {
	requestDef := GenReqDefForUpdateGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupWebhookResponse), nil
	}
}

// UpdateGroupWebhookInvoker 更新代码组下单个Webhook数据
func (c *CodeHubClient) UpdateGroupWebhookInvoker(request *model.UpdateGroupWebhookRequest) *UpdateGroupWebhookInvoker {
	requestDef := GenReqDefForUpdateGroupWebhook()
	return &UpdateGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectWebhook 更新项目下单个Webhook数据
//
// 更新项目下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateProjectWebhook(request *model.UpdateProjectWebhookRequest) (*model.UpdateProjectWebhookResponse, error) {
	requestDef := GenReqDefForUpdateProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectWebhookResponse), nil
	}
}

// UpdateProjectWebhookInvoker 更新项目下单个Webhook数据
func (c *CodeHubClient) UpdateProjectWebhookInvoker(request *model.UpdateProjectWebhookRequest) *UpdateProjectWebhookInvoker {
	requestDef := GenReqDefForUpdateProjectWebhook()
	return &UpdateProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryWebhook 更新仓库下单个Webhook数据
//
// 更新仓库下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateRepositoryWebhook(request *model.UpdateRepositoryWebhookRequest) (*model.UpdateRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryWebhookResponse), nil
	}
}

// UpdateRepositoryWebhookInvoker 更新仓库下单个Webhook数据
func (c *CodeHubClient) UpdateRepositoryWebhookInvoker(request *model.UpdateRepositoryWebhookRequest) *UpdateRepositoryWebhookInvoker {
	requestDef := GenReqDefForUpdateRepositoryWebhook()
	return &UpdateRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
