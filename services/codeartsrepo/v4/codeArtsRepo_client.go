package v4

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsrepo/v4/model"
)

type CodeArtsRepoClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsRepoClient(hcClient *httpclient.HcHttpClient) *CodeArtsRepoClient {
	return &CodeArtsRepoClient{HcClient: hcClient}
}

func CodeArtsRepoClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateCommit 创建提交信息
//
// 创建提交信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateCommit(request *model.CreateCommitRequest) (*model.CreateCommitResponse, error) {
	requestDef := GenReqDefForCreateCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommitResponse), nil
	}
}

// CreateCommitInvoker 创建提交信息
func (c *CodeArtsRepoClient) CreateCommitInvoker(request *model.CreateCommitRequest) *CreateCommitInvoker {
	requestDef := GenReqDefForCreateCommit()
	return &CreateCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCommitRevert 回退提交
//
// 回退提交
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateCommitRevert(request *model.CreateCommitRevertRequest) (*model.CreateCommitRevertResponse, error) {
	requestDef := GenReqDefForCreateCommitRevert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommitRevertResponse), nil
	}
}

// CreateCommitRevertInvoker 回退提交
func (c *CodeArtsRepoClient) CreateCommitRevertInvoker(request *model.CreateCommitRevertRequest) *CreateCommitRevertInvoker {
	requestDef := GenReqDefForCreateCommitRevert()
	return &CreateCommitRevertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitAssociatedRefs 根据提交ID查询分支、标签列表
//
// 根据提交ID查询分支、标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListCommitAssociatedRefs(request *model.ListCommitAssociatedRefsRequest) (*model.ListCommitAssociatedRefsResponse, error) {
	requestDef := GenReqDefForListCommitAssociatedRefs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitAssociatedRefsResponse), nil
	}
}

// ListCommitAssociatedRefsInvoker 根据提交ID查询分支、标签列表
func (c *CodeArtsRepoClient) ListCommitAssociatedRefsInvoker(request *model.ListCommitAssociatedRefsRequest) *ListCommitAssociatedRefsInvoker {
	requestDef := GenReqDefForListCommitAssociatedRefs()
	return &ListCommitAssociatedRefsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommits 查看提交列表
//
// 查看提交列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListCommits(request *model.ListCommitsRequest) (*model.ListCommitsResponse, error) {
	requestDef := GenReqDefForListCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitsResponse), nil
	}
}

// ListCommitsInvoker 查看提交列表
func (c *CodeArtsRepoClient) ListCommitsInvoker(request *model.ListCommitsRequest) *ListCommitsInvoker {
	requestDef := GenReqDefForListCommits()
	return &ListCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommit 获取特定提交信息
//
// 获取特定提交信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowCommit(request *model.ShowCommitRequest) (*model.ShowCommitResponse, error) {
	requestDef := GenReqDefForShowCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitResponse), nil
	}
}

// ShowCommitInvoker 获取特定提交信息
func (c *CodeArtsRepoClient) ShowCommitInvoker(request *model.ShowCommitRequest) *ShowCommitInvoker {
	requestDef := GenReqDefForShowCommit()
	return &ShowCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitDiffMetadata 获取commit引入的文件变更元数据
//
// 获取commit引入的文件变更元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowCommitDiffMetadata(request *model.ShowCommitDiffMetadataRequest) (*model.ShowCommitDiffMetadataResponse, error) {
	requestDef := GenReqDefForShowCommitDiffMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitDiffMetadataResponse), nil
	}
}

// ShowCommitDiffMetadataInvoker 获取commit引入的文件变更元数据
func (c *CodeArtsRepoClient) ShowCommitDiffMetadataInvoker(request *model.ShowCommitDiffMetadataRequest) *ShowCommitDiffMetadataInvoker {
	requestDef := GenReqDefForShowCommitDiffMetadata()
	return &ShowCommitDiffMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitFileDiff 获取commit引入的指定文件的变更内容
//
// 获取commit引入的指定文件的变更内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowCommitFileDiff(request *model.ShowCommitFileDiffRequest) (*model.ShowCommitFileDiffResponse, error) {
	requestDef := GenReqDefForShowCommitFileDiff()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitFileDiffResponse), nil
	}
}

// ShowCommitFileDiffInvoker 获取commit引入的指定文件的变更内容
func (c *CodeArtsRepoClient) ShowCommitFileDiffInvoker(request *model.ShowCommitFileDiffRequest) *ShowCommitFileDiffInvoker {
	requestDef := GenReqDefForShowCommitFileDiff()
	return &ShowCommitFileDiffInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiffCommit 获取提交差异
//
// 获取提交差异
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowDiffCommit(request *model.ShowDiffCommitRequest) (*model.ShowDiffCommitResponse, error) {
	requestDef := GenReqDefForShowDiffCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiffCommitResponse), nil
	}
}

// ShowDiffCommitInvoker 获取提交差异
func (c *CodeArtsRepoClient) ShowDiffCommitInvoker(request *model.ShowDiffCommitRequest) *ShowDiffCommitInvoker {
	requestDef := GenReqDefForShowDiffCommit()
	return &ShowDiffCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussion 创建合并请求检视意见
//
// 创建合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateMergeRequestDiscussion(request *model.CreateMergeRequestDiscussionRequest) (*model.CreateMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestDiscussionResponse), nil
	}
}

// CreateMergeRequestDiscussionInvoker 创建合并请求检视意见
func (c *CodeArtsRepoClient) CreateMergeRequestDiscussionInvoker(request *model.CreateMergeRequestDiscussionRequest) *CreateMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussion()
	return &CreateMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestDiscussionResponse 回复合并请求检视意见
//
// 回复合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateMergeRequestDiscussionResponse(request *model.CreateMergeRequestDiscussionResponseRequest) (*model.CreateMergeRequestDiscussionResponseResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestDiscussionResponse()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestDiscussionResponseResponse), nil
	}
}

// CreateMergeRequestDiscussionResponseInvoker 回复合并请求检视意见
func (c *CodeArtsRepoClient) CreateMergeRequestDiscussionResponseInvoker(request *model.CreateMergeRequestDiscussionResponseRequest) *CreateMergeRequestDiscussionResponseInvoker {
	requestDef := GenReqDefForCreateMergeRequestDiscussionResponse()
	return &CreateMergeRequestDiscussionResponseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReviewSetting 创建/更新检视意见设置
//
// 创建/更新检视意见设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateReviewSetting(request *model.CreateReviewSettingRequest) (*model.CreateReviewSettingResponse, error) {
	requestDef := GenReqDefForCreateReviewSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReviewSettingResponse), nil
	}
}

// CreateReviewSettingInvoker 创建/更新检视意见设置
func (c *CodeArtsRepoClient) CreateReviewSettingInvoker(request *model.CreateReviewSettingRequest) *CreateReviewSettingInvoker {
	requestDef := GenReqDefForCreateReviewSetting()
	return &CreateReviewSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestDiscussion 删除合并请求检视意见
//
// 删除合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteMergeRequestDiscussion(request *model.DeleteMergeRequestDiscussionRequest) (*model.DeleteMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestDiscussionResponse), nil
	}
}

// DeleteMergeRequestDiscussionInvoker 删除合并请求检视意见
func (c *CodeArtsRepoClient) DeleteMergeRequestDiscussionInvoker(request *model.DeleteMergeRequestDiscussionRequest) *DeleteMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForDeleteMergeRequestDiscussion()
	return &DeleteMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitDiscussions 获取代码页单个提交下检视意见列表
//
// 获取代码页单个提交下检视意见列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListCommitDiscussions(request *model.ListCommitDiscussionsRequest) (*model.ListCommitDiscussionsResponse, error) {
	requestDef := GenReqDefForListCommitDiscussions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitDiscussionsResponse), nil
	}
}

// ListCommitDiscussionsInvoker 获取代码页单个提交下检视意见列表
func (c *CodeArtsRepoClient) ListCommitDiscussionsInvoker(request *model.ListCommitDiscussionsRequest) *ListCommitDiscussionsInvoker {
	requestDef := GenReqDefForListCommitDiscussions()
	return &ListCommitDiscussionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDefaultReviewCategories 获取默认的检视意见分类
//
// 获取默认的检视意见分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListDefaultReviewCategories(request *model.ListDefaultReviewCategoriesRequest) (*model.ListDefaultReviewCategoriesResponse, error) {
	requestDef := GenReqDefForListDefaultReviewCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDefaultReviewCategoriesResponse), nil
	}
}

// ListDefaultReviewCategoriesInvoker 获取默认的检视意见分类
func (c *CodeArtsRepoClient) ListDefaultReviewCategoriesInvoker(request *model.ListDefaultReviewCategoriesRequest) *ListDefaultReviewCategoriesInvoker {
	requestDef := GenReqDefForListDefaultReviewCategories()
	return &ListDefaultReviewCategoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestDiscussions 获取合并请求检视意见列表
//
// 获取合并请求检视意见列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestDiscussions(request *model.ListMergeRequestDiscussionsRequest) (*model.ListMergeRequestDiscussionsResponse, error) {
	requestDef := GenReqDefForListMergeRequestDiscussions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestDiscussionsResponse), nil
	}
}

// ListMergeRequestDiscussionsInvoker 获取合并请求检视意见列表
func (c *CodeArtsRepoClient) ListMergeRequestDiscussionsInvoker(request *model.ListMergeRequestDiscussionsRequest) *ListMergeRequestDiscussionsInvoker {
	requestDef := GenReqDefForListMergeRequestDiscussions()
	return &ListMergeRequestDiscussionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestSystemNotes 获取合并请求动态列表
//
// 获取合并请求动态列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestSystemNotes(request *model.ListMergeRequestSystemNotesRequest) (*model.ListMergeRequestSystemNotesResponse, error) {
	requestDef := GenReqDefForListMergeRequestSystemNotes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestSystemNotesResponse), nil
	}
}

// ListMergeRequestSystemNotesInvoker 获取合并请求动态列表
func (c *CodeArtsRepoClient) ListMergeRequestSystemNotesInvoker(request *model.ListMergeRequestSystemNotesRequest) *ListMergeRequestSystemNotesInvoker {
	requestDef := GenReqDefForListMergeRequestSystemNotes()
	return &ListMergeRequestSystemNotesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectNoteRequiredAttributes 获取项目检视意见必填项
//
// 获取项目检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectNoteRequiredAttributes(request *model.ListProjectNoteRequiredAttributesRequest) (*model.ListProjectNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForListProjectNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectNoteRequiredAttributesResponse), nil
	}
}

// ListProjectNoteRequiredAttributesInvoker 获取项目检视意见必填项
func (c *CodeArtsRepoClient) ListProjectNoteRequiredAttributesInvoker(request *model.ListProjectNoteRequiredAttributesRequest) *ListProjectNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForListProjectNoteRequiredAttributes()
	return &ListProjectNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryReviewAuthors 获取仓库下检视意见作者列表
//
// 获取仓库下检视意见作者列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryReviewAuthors(request *model.ListRepositoryReviewAuthorsRequest) (*model.ListRepositoryReviewAuthorsResponse, error) {
	requestDef := GenReqDefForListRepositoryReviewAuthors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryReviewAuthorsResponse), nil
	}
}

// ListRepositoryReviewAuthorsInvoker 获取仓库下检视意见作者列表
func (c *CodeArtsRepoClient) ListRepositoryReviewAuthorsInvoker(request *model.ListRepositoryReviewAuthorsRequest) *ListRepositoryReviewAuthorsInvoker {
	requestDef := GenReqDefForListRepositoryReviewAuthors()
	return &ListRepositoryReviewAuthorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryReviews 获取仓库检视意见列表
//
// 获取仓库检视意见列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryReviews(request *model.ListRepositoryReviewsRequest) (*model.ListRepositoryReviewsResponse, error) {
	requestDef := GenReqDefForListRepositoryReviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryReviewsResponse), nil
	}
}

// ListRepositoryReviewsInvoker 获取仓库检视意见列表
func (c *CodeArtsRepoClient) ListRepositoryReviewsInvoker(request *model.ListRepositoryReviewsRequest) *ListRepositoryReviewsInvoker {
	requestDef := GenReqDefForListRepositoryReviews()
	return &ListRepositoryReviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupNoteRequiredAttributes 获取代码组检视意见必填项
//
// 获取代码组检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupNoteRequiredAttributes(request *model.ShowGroupNoteRequiredAttributesRequest) (*model.ShowGroupNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForShowGroupNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupNoteRequiredAttributesResponse), nil
	}
}

// ShowGroupNoteRequiredAttributesInvoker 获取代码组检视意见必填项
func (c *CodeArtsRepoClient) ShowGroupNoteRequiredAttributesInvoker(request *model.ShowGroupNoteRequiredAttributesRequest) *ShowGroupNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForShowGroupNoteRequiredAttributes()
	return &ShowGroupNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupReviewSettings 获取代码组检视意见设置(不含必填项)
//
// 获取代码组检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupReviewSettings(request *model.ShowGroupReviewSettingsRequest) (*model.ShowGroupReviewSettingsResponse, error) {
	requestDef := GenReqDefForShowGroupReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupReviewSettingsResponse), nil
	}
}

// ShowGroupReviewSettingsInvoker 获取代码组检视意见设置(不含必填项)
func (c *CodeArtsRepoClient) ShowGroupReviewSettingsInvoker(request *model.ShowGroupReviewSettingsRequest) *ShowGroupReviewSettingsInvoker {
	requestDef := GenReqDefForShowGroupReviewSettings()
	return &ShowGroupReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestDiscussion 根据discussion_id获取合并请求检视意见
//
// 根据discussion_id获取合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestDiscussion(request *model.ShowMergeRequestDiscussionRequest) (*model.ShowMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForShowMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestDiscussionResponse), nil
	}
}

// ShowMergeRequestDiscussionInvoker 根据discussion_id获取合并请求检视意见
func (c *CodeArtsRepoClient) ShowMergeRequestDiscussionInvoker(request *model.ShowMergeRequestDiscussionRequest) *ShowMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForShowMergeRequestDiscussion()
	return &ShowMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNoteRequiredAttributes 获取仓库检视意见必填项
//
// 获取仓库检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowNoteRequiredAttributes(request *model.ShowNoteRequiredAttributesRequest) (*model.ShowNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForShowNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNoteRequiredAttributesResponse), nil
	}
}

// ShowNoteRequiredAttributesInvoker 获取仓库检视意见必填项
func (c *CodeArtsRepoClient) ShowNoteRequiredAttributesInvoker(request *model.ShowNoteRequiredAttributesRequest) *ShowNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForShowNoteRequiredAttributes()
	return &ShowNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectReviewSettings 获取项目检视意见设置(不含必填项)
//
// 获取项目检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectReviewSettings(request *model.ShowProjectReviewSettingsRequest) (*model.ShowProjectReviewSettingsResponse, error) {
	requestDef := GenReqDefForShowProjectReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectReviewSettingsResponse), nil
	}
}

// ShowProjectReviewSettingsInvoker 获取项目检视意见设置(不含必填项)
func (c *CodeArtsRepoClient) ShowProjectReviewSettingsInvoker(request *model.ShowProjectReviewSettingsRequest) *ShowProjectReviewSettingsInvoker {
	requestDef := GenReqDefForShowProjectReviewSettings()
	return &ShowProjectReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReviewSetting 获取检视意见设置
//
// 获取检视意见设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowReviewSetting(request *model.ShowReviewSettingRequest) (*model.ShowReviewSettingResponse, error) {
	requestDef := GenReqDefForShowReviewSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReviewSettingResponse), nil
	}
}

// ShowReviewSettingInvoker 获取检视意见设置
func (c *CodeArtsRepoClient) ShowReviewSettingInvoker(request *model.ShowReviewSettingRequest) *ShowReviewSettingInvoker {
	requestDef := GenReqDefForShowReviewSetting()
	return &ShowReviewSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupNoteRequiredAttributes 创建/更新代码组检视意见必填项
//
// 创建/更新代码组检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupNoteRequiredAttributes(request *model.UpdateGroupNoteRequiredAttributesRequest) (*model.UpdateGroupNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateGroupNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupNoteRequiredAttributesResponse), nil
	}
}

// UpdateGroupNoteRequiredAttributesInvoker 创建/更新代码组检视意见必填项
func (c *CodeArtsRepoClient) UpdateGroupNoteRequiredAttributesInvoker(request *model.UpdateGroupNoteRequiredAttributesRequest) *UpdateGroupNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateGroupNoteRequiredAttributes()
	return &UpdateGroupNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupReviewSettings 创建/更新代码组检视意见设置(不含必填项)
//
// 创建/更新代码组检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupReviewSettings(request *model.UpdateGroupReviewSettingsRequest) (*model.UpdateGroupReviewSettingsResponse, error) {
	requestDef := GenReqDefForUpdateGroupReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupReviewSettingsResponse), nil
	}
}

// UpdateGroupReviewSettingsInvoker 创建/更新代码组检视意见设置(不含必填项)
func (c *CodeArtsRepoClient) UpdateGroupReviewSettingsInvoker(request *model.UpdateGroupReviewSettingsRequest) *UpdateGroupReviewSettingsInvoker {
	requestDef := GenReqDefForUpdateGroupReviewSettings()
	return &UpdateGroupReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestDiscussion 更新合并请求检视意见
//
// 更新合并请求检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestDiscussion(request *model.UpdateMergeRequestDiscussionRequest) (*model.UpdateMergeRequestDiscussionResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestDiscussion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestDiscussionResponse), nil
	}
}

// UpdateMergeRequestDiscussionInvoker 更新合并请求检视意见
func (c *CodeArtsRepoClient) UpdateMergeRequestDiscussionInvoker(request *model.UpdateMergeRequestDiscussionRequest) *UpdateMergeRequestDiscussionInvoker {
	requestDef := GenReqDefForUpdateMergeRequestDiscussion()
	return &UpdateMergeRequestDiscussionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestDiscussionInfo 更新合并请求检视意见中除评论内容以外的信息
//
// 更新合并请求检视意见中除评论内容以外的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestDiscussionInfo(request *model.UpdateMergeRequestDiscussionInfoRequest) (*model.UpdateMergeRequestDiscussionInfoResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestDiscussionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestDiscussionInfoResponse), nil
	}
}

// UpdateMergeRequestDiscussionInfoInvoker 更新合并请求检视意见中除评论内容以外的信息
func (c *CodeArtsRepoClient) UpdateMergeRequestDiscussionInfoInvoker(request *model.UpdateMergeRequestDiscussionInfoRequest) *UpdateMergeRequestDiscussionInfoInvoker {
	requestDef := GenReqDefForUpdateMergeRequestDiscussionInfo()
	return &UpdateMergeRequestDiscussionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNoteRequiredAttributes 创建/更新仓库检视意见必填项
//
// 创建/更新仓库检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateNoteRequiredAttributes(request *model.UpdateNoteRequiredAttributesRequest) (*model.UpdateNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNoteRequiredAttributesResponse), nil
	}
}

// UpdateNoteRequiredAttributesInvoker 创建/更新仓库检视意见必填项
func (c *CodeArtsRepoClient) UpdateNoteRequiredAttributesInvoker(request *model.UpdateNoteRequiredAttributesRequest) *UpdateNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateNoteRequiredAttributes()
	return &UpdateNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectNoteRequiredAttributes 创建/更新项目检视意见必填项
//
// 创建/更新项目检视意见必填项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectNoteRequiredAttributes(request *model.UpdateProjectNoteRequiredAttributesRequest) (*model.UpdateProjectNoteRequiredAttributesResponse, error) {
	requestDef := GenReqDefForUpdateProjectNoteRequiredAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectNoteRequiredAttributesResponse), nil
	}
}

// UpdateProjectNoteRequiredAttributesInvoker 创建/更新项目检视意见必填项
func (c *CodeArtsRepoClient) UpdateProjectNoteRequiredAttributesInvoker(request *model.UpdateProjectNoteRequiredAttributesRequest) *UpdateProjectNoteRequiredAttributesInvoker {
	requestDef := GenReqDefForUpdateProjectNoteRequiredAttributes()
	return &UpdateProjectNoteRequiredAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectReviewSettings 创建/更新项目检视意见设置(不含必填项)
//
// 创建/更新项目检视意见设置(不含必填项)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectReviewSettings(request *model.UpdateProjectReviewSettingsRequest) (*model.UpdateProjectReviewSettingsResponse, error) {
	requestDef := GenReqDefForUpdateProjectReviewSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectReviewSettingsResponse), nil
	}
}

// UpdateProjectReviewSettingsInvoker 创建/更新项目检视意见设置(不含必填项)
func (c *CodeArtsRepoClient) UpdateProjectReviewSettingsInvoker(request *model.UpdateProjectReviewSettingsRequest) *UpdateProjectReviewSettingsInvoker {
	requestDef := GenReqDefForUpdateProjectReviewSettings()
	return &UpdateProjectReviewSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建文件
//
// 创建文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建文件
func (c *CodeArtsRepoClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除文件
//
// 删除文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除文件
func (c *CodeArtsRepoClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBlobsRaw 获取仓库单个文件内容(下载单个文件)
//
// 获取仓库单个文件内容(下载单个文件)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DownloadBlobsRaw(request *model.DownloadBlobsRawRequest) (*model.DownloadBlobsRawResponse, error) {
	requestDef := GenReqDefForDownloadBlobsRaw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlobsRawResponse), nil
	}
}

// DownloadBlobsRawInvoker 获取仓库单个文件内容(下载单个文件)
func (c *CodeArtsRepoClient) DownloadBlobsRawInvoker(request *model.DownloadBlobsRawRequest) *DownloadBlobsRawInvoker {
	requestDef := GenReqDefForDownloadBlobsRaw()
	return &DownloadBlobsRawInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileBlameLines 获取文件追溯信息
//
// 获取文件追溯信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListFileBlameLines(request *model.ListFileBlameLinesRequest) (*model.ListFileBlameLinesResponse, error) {
	requestDef := GenReqDefForListFileBlameLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileBlameLinesResponse), nil
	}
}

// ListFileBlameLinesInvoker 获取文件追溯信息
func (c *CodeArtsRepoClient) ListFileBlameLinesInvoker(request *model.ListFileBlameLinesRequest) *ListFileBlameLinesInvoker {
	requestDef := GenReqDefForListFileBlameLines()
	return &ListFileBlameLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileUpperTreeEntries 获取当前文件上级树结构
//
// 获取当前文件上级树结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListFileUpperTreeEntries(request *model.ListFileUpperTreeEntriesRequest) (*model.ListFileUpperTreeEntriesResponse, error) {
	requestDef := GenReqDefForListFileUpperTreeEntries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileUpperTreeEntriesResponse), nil
	}
}

// ListFileUpperTreeEntriesInvoker 获取当前文件上级树结构
func (c *CodeArtsRepoClient) ListFileUpperTreeEntriesInvoker(request *model.ListFileUpperTreeEntriesRequest) *ListFileUpperTreeEntriesInvoker {
	requestDef := GenReqDefForListFileUpperTreeEntries()
	return &ListFileUpperTreeEntriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFiles 获取指定分支下所有的文件列表
//
// 获取指定分支下所有的文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListFiles(request *model.ListFilesRequest) (*model.ListFilesResponse, error) {
	requestDef := GenReqDefForListFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFilesResponse), nil
	}
}

// ListFilesInvoker 获取指定分支下所有的文件列表
func (c *CodeArtsRepoClient) ListFilesInvoker(request *model.ListFilesRequest) *ListFilesInvoker {
	requestDef := GenReqDefForListFiles()
	return &ListFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogsTree 查看文件树
//
// 查看文件树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListLogsTree(request *model.ListLogsTreeRequest) (*model.ListLogsTreeResponse, error) {
	requestDef := GenReqDefForListLogsTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsTreeResponse), nil
	}
}

// ListLogsTreeInvoker 查看文件树
func (c *CodeArtsRepoClient) ListLogsTreeInvoker(request *model.ListLogsTreeRequest) *ListLogsTreeInvoker {
	requestDef := GenReqDefForListLogsTree()
	return &ListLogsTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrees 查看分支文件列表
//
// 查看分支文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListTrees(request *model.ListTreesRequest) (*model.ListTreesResponse, error) {
	requestDef := GenReqDefForListTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTreesResponse), nil
	}
}

// ListTreesInvoker 查看分支文件列表
func (c *CodeArtsRepoClient) ListTreesInvoker(request *model.ListTreesRequest) *ListTreesInvoker {
	requestDef := GenReqDefForListTrees()
	return &ListTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenameFile 文件重命名
//
// 文件重命名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RenameFile(request *model.RenameFileRequest) (*model.RenameFileResponse, error) {
	requestDef := GenReqDefForRenameFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenameFileResponse), nil
	}
}

// RenameFileInvoker 文件重命名
func (c *CodeArtsRepoClient) RenameFileInvoker(request *model.RenameFileRequest) *RenameFileInvoker {
	requestDef := GenReqDefForRenameFile()
	return &RenameFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFile 查看文件属性与内容
//
// 查看文件属性与内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowFile(request *model.ShowFileRequest) (*model.ShowFileResponse, error) {
	requestDef := GenReqDefForShowFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileResponse), nil
	}
}

// ShowFileInvoker 查看文件属性与内容
func (c *CodeArtsRepoClient) ShowFileInvoker(request *model.ShowFileRequest) *ShowFileInvoker {
	requestDef := GenReqDefForShowFile()
	return &ShowFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileContent 获取文件内容
//
// 获取文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowFileContent(request *model.ShowFileContentRequest) (*model.ShowFileContentResponse, error) {
	requestDef := GenReqDefForShowFileContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileContentResponse), nil
	}
}

// ShowFileContentInvoker 获取文件内容
func (c *CodeArtsRepoClient) ShowFileContentInvoker(request *model.ShowFileContentRequest) *ShowFileContentInvoker {
	requestDef := GenReqDefForShowFileContent()
	return &ShowFileContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileRaw 获取仓库单个文件内容
//
// 获取仓库单个文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowFileRaw(request *model.ShowFileRawRequest) (*model.ShowFileRawResponse, error) {
	requestDef := GenReqDefForShowFileRaw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileRawResponse), nil
	}
}

// ShowFileRawInvoker 获取仓库单个文件内容
func (c *CodeArtsRepoClient) ShowFileRawInvoker(request *model.ShowFileRawRequest) *ShowFileRawInvoker {
	requestDef := GenReqDefForShowFileRaw()
	return &ShowFileRawInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReadmeFile 获取仓库默认分支的Readme文件内容
//
// 获取仓库默认分支的Readme文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowReadmeFile(request *model.ShowReadmeFileRequest) (*model.ShowReadmeFileResponse, error) {
	requestDef := GenReqDefForShowReadmeFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReadmeFileResponse), nil
	}
}

// ShowReadmeFileInvoker 获取仓库默认分支的Readme文件内容
func (c *CodeArtsRepoClient) ShowReadmeFileInvoker(request *model.ShowReadmeFileRequest) *ShowReadmeFileInvoker {
	requestDef := GenReqDefForShowReadmeFile()
	return &ShowReadmeFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFile 更新文件内容
//
// 修改仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateFile(request *model.UpdateFileRequest) (*model.UpdateFileResponse, error) {
	requestDef := GenReqDefForUpdateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFileResponse), nil
	}
}

// UpdateFileInvoker 更新文件内容
func (c *CodeArtsRepoClient) UpdateFileInvoker(request *model.UpdateFileRequest) *UpdateFileInvoker {
	requestDef := GenReqDefForUpdateFile()
	return &UpdateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateGroupUserGroup 关联代码组与成员组
//
// 关联代码组与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AssociateGroupUserGroup(request *model.AssociateGroupUserGroupRequest) (*model.AssociateGroupUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateGroupUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateGroupUserGroupResponse), nil
	}
}

// AssociateGroupUserGroupInvoker 关联代码组与成员组
func (c *CodeArtsRepoClient) AssociateGroupUserGroupInvoker(request *model.AssociateGroupUserGroupRequest) *AssociateGroupUserGroupInvoker {
	requestDef := GenReqDefForAssociateGroupUserGroup()
	return &AssociateGroupUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroup 创建代码组
//
// 创建代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建代码组
func (c *CodeArtsRepoClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除代码组
//
// 删除代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除代码组
func (c *CodeArtsRepoClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupAddableMembers 获取代码组下可添加的成员列表
//
// 获取代码组下可添加的成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupAddableMembers(request *model.ListGroupAddableMembersRequest) (*model.ListGroupAddableMembersResponse, error) {
	requestDef := GenReqDefForListGroupAddableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupAddableMembersResponse), nil
	}
}

// ListGroupAddableMembersInvoker 获取代码组下可添加的成员列表
func (c *CodeArtsRepoClient) ListGroupAddableMembersInvoker(request *model.ListGroupAddableMembersRequest) *ListGroupAddableMembersInvoker {
	requestDef := GenReqDefForListGroupAddableMembers()
	return &ListGroupAddableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupAddableUserGroups 获取代码组下可添加的成员组
//
// 获取代码组下可添加的成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupAddableUserGroups(request *model.ListGroupAddableUserGroupsRequest) (*model.ListGroupAddableUserGroupsResponse, error) {
	requestDef := GenReqDefForListGroupAddableUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupAddableUserGroupsResponse), nil
	}
}

// ListGroupAddableUserGroupsInvoker 获取代码组下可添加的成员组
func (c *CodeArtsRepoClient) ListGroupAddableUserGroupsInvoker(request *model.ListGroupAddableUserGroupsRequest) *ListGroupAddableUserGroupsInvoker {
	requestDef := GenReqDefForListGroupAddableUserGroups()
	return &ListGroupAddableUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMembers 获取代码组下可添加的成员列表
//
// 获取代码组下可添加的成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupMembers(request *model.ListGroupMembersRequest) (*model.ListGroupMembersResponse, error) {
	requestDef := GenReqDefForListGroupMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMembersResponse), nil
	}
}

// ListGroupMembersInvoker 获取代码组下可添加的成员列表
func (c *CodeArtsRepoClient) ListGroupMembersInvoker(request *model.ListGroupMembersRequest) *ListGroupMembersInvoker {
	requestDef := GenReqDefForListGroupMembers()
	return &ListGroupMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupPermissionResources 获取代码组权限资源点列表
//
// 获取代码组权限资源点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupPermissionResources(request *model.ListGroupPermissionResourcesRequest) (*model.ListGroupPermissionResourcesResponse, error) {
	requestDef := GenReqDefForListGroupPermissionResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupPermissionResourcesResponse), nil
	}
}

// ListGroupPermissionResourcesInvoker 获取代码组权限资源点列表
func (c *CodeArtsRepoClient) ListGroupPermissionResourcesInvoker(request *model.ListGroupPermissionResourcesRequest) *ListGroupPermissionResourcesInvoker {
	requestDef := GenReqDefForListGroupPermissionResources()
	return &ListGroupPermissionResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupSubgroupsAndRepositories 获取代码组下的子代码组和仓库列表
//
// 获取代码组下的子代码组和仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupSubgroupsAndRepositories(request *model.ListGroupSubgroupsAndRepositoriesRequest) (*model.ListGroupSubgroupsAndRepositoriesResponse, error) {
	requestDef := GenReqDefForListGroupSubgroupsAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupSubgroupsAndRepositoriesResponse), nil
	}
}

// ListGroupSubgroupsAndRepositoriesInvoker 获取代码组下的子代码组和仓库列表
func (c *CodeArtsRepoClient) ListGroupSubgroupsAndRepositoriesInvoker(request *model.ListGroupSubgroupsAndRepositoriesRequest) *ListGroupSubgroupsAndRepositoriesInvoker {
	requestDef := GenReqDefForListGroupSubgroupsAndRepositories()
	return &ListGroupSubgroupsAndRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupUserGroups 组织下查询成员组列表
//
// 组织下查询成员组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupUserGroups(request *model.ListGroupUserGroupsRequest) (*model.ListGroupUserGroupsResponse, error) {
	requestDef := GenReqDefForListGroupUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupUserGroupsResponse), nil
	}
}

// ListGroupUserGroupsInvoker 组织下查询成员组列表
func (c *CodeArtsRepoClient) ListGroupUserGroupsInvoker(request *model.ListGroupUserGroupsRequest) *ListGroupUserGroupsInvoker {
	requestDef := GenReqDefForListGroupUserGroups()
	return &ListGroupUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 获取代码组列表
//
// 获取代码组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 获取代码组列表
func (c *CodeArtsRepoClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManageableGroups 获取项目下当前用户有管理权限的代码组列表
//
// 获取项目下当前用户有管理权限的代码组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListManageableGroups(request *model.ListManageableGroupsRequest) (*model.ListManageableGroupsResponse, error) {
	requestDef := GenReqDefForListManageableGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManageableGroupsResponse), nil
	}
}

// ListManageableGroupsInvoker 获取项目下当前用户有管理权限的代码组列表
func (c *CodeArtsRepoClient) ListManageableGroupsInvoker(request *model.ListManageableGroupsRequest) *ListManageableGroupsInvoker {
	requestDef := GenReqDefForListManageableGroups()
	return &ListManageableGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroup 获取代码组信息
//
// 获取代码组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

// ShowGroupInvoker 获取代码组信息
func (c *CodeArtsRepoClient) ShowGroupInvoker(request *model.ShowGroupRequest) *ShowGroupInvoker {
	requestDef := GenReqDefForShowGroup()
	return &ShowGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupGeneralPolicy 获取指定代码组的基本设置信息
//
// 获取指定代码组的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupGeneralPolicy(request *model.ShowGroupGeneralPolicyRequest) (*model.ShowGroupGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowGroupGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupGeneralPolicyResponse), nil
	}
}

// ShowGroupGeneralPolicyInvoker 获取指定代码组的基本设置信息
func (c *CodeArtsRepoClient) ShowGroupGeneralPolicyInvoker(request *model.ShowGroupGeneralPolicyRequest) *ShowGroupGeneralPolicyInvoker {
	requestDef := GenReqDefForShowGroupGeneralPolicy()
	return &ShowGroupGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupInheritSetting 获取代码组继承设置项
//
// 获取代码组继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupInheritSetting(request *model.ShowGroupInheritSettingRequest) (*model.ShowGroupInheritSettingResponse, error) {
	requestDef := GenReqDefForShowGroupInheritSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupInheritSettingResponse), nil
	}
}

// ShowGroupInheritSettingInvoker 获取代码组继承设置项
func (c *CodeArtsRepoClient) ShowGroupInheritSettingInvoker(request *model.ShowGroupInheritSettingRequest) *ShowGroupInheritSettingInvoker {
	requestDef := GenReqDefForShowGroupInheritSetting()
	return &ShowGroupInheritSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupPermissionInheritEnabled 获取代码组继承权限设置开关
//
// 获取代码组继承权限设置开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupPermissionInheritEnabled(request *model.ShowGroupPermissionInheritEnabledRequest) (*model.ShowGroupPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForShowGroupPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupPermissionInheritEnabledResponse), nil
	}
}

// ShowGroupPermissionInheritEnabledInvoker 获取代码组继承权限设置开关
func (c *CodeArtsRepoClient) ShowGroupPermissionInheritEnabledInvoker(request *model.ShowGroupPermissionInheritEnabledRequest) *ShowGroupPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForShowGroupPermissionInheritEnabled()
	return &ShowGroupPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupSettingsInheritCfg 获取代码组继承设置项
//
// 获取代码组继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupSettingsInheritCfg(request *model.ShowGroupSettingsInheritCfgRequest) (*model.ShowGroupSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForShowGroupSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupSettingsInheritCfgResponse), nil
	}
}

// ShowGroupSettingsInheritCfgInvoker 获取代码组继承设置项
func (c *CodeArtsRepoClient) ShowGroupSettingsInheritCfgInvoker(request *model.ShowGroupSettingsInheritCfgRequest) *ShowGroupSettingsInheritCfgInvoker {
	requestDef := GenReqDefForShowGroupSettingsInheritCfg()
	return &ShowGroupSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWatermark 获取代码组水印设置
//
// 获取代码组水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupWatermark(request *model.ShowGroupWatermarkRequest) (*model.ShowGroupWatermarkResponse, error) {
	requestDef := GenReqDefForShowGroupWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWatermarkResponse), nil
	}
}

// ShowGroupWatermarkInvoker 获取代码组水印设置
func (c *CodeArtsRepoClient) ShowGroupWatermarkInvoker(request *model.ShowGroupWatermarkRequest) *ShowGroupWatermarkInvoker {
	requestDef := GenReqDefForShowGroupWatermark()
	return &ShowGroupWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupsGeneralPolicy 获取指定代码组的基本设置信息
//
// 获取指定代码组的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupsGeneralPolicy(request *model.ShowGroupsGeneralPolicyRequest) (*model.ShowGroupsGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowGroupsGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsGeneralPolicyResponse), nil
	}
}

// ShowGroupsGeneralPolicyInvoker 获取指定代码组的基本设置信息
func (c *CodeArtsRepoClient) ShowGroupsGeneralPolicyInvoker(request *model.ShowGroupsGeneralPolicyRequest) *ShowGroupsGeneralPolicyInvoker {
	requestDef := GenReqDefForShowGroupsGeneralPolicy()
	return &ShowGroupsGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupsInherit 获取代码组的继承设置
//
// 获取代码组的继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupsInherit(request *model.ShowGroupsInheritRequest) (*model.ShowGroupsInheritResponse, error) {
	requestDef := GenReqDefForShowGroupsInherit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsInheritResponse), nil
	}
}

// ShowGroupsInheritInvoker 获取代码组的继承设置
func (c *CodeArtsRepoClient) ShowGroupsInheritInvoker(request *model.ShowGroupsInheritRequest) *ShowGroupsInheritInvoker {
	requestDef := GenReqDefForShowGroupsInherit()
	return &ShowGroupsInheritInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferGroup 移交代码组
//
// 移交代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) TransferGroup(request *model.TransferGroupRequest) (*model.TransferGroupResponse, error) {
	requestDef := GenReqDefForTransferGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferGroupResponse), nil
	}
}

// TransferGroupInvoker 移交代码组
func (c *CodeArtsRepoClient) TransferGroupInvoker(request *model.TransferGroupRequest) *TransferGroupInvoker {
	requestDef := GenReqDefForTransferGroup()
	return &TransferGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupGeneralPolicy 更新代码组的基本设置信息
//
// 更新代码组的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupGeneralPolicy(request *model.UpdateGroupGeneralPolicyRequest) (*model.UpdateGroupGeneralPolicyResponse, error) {
	requestDef := GenReqDefForUpdateGroupGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupGeneralPolicyResponse), nil
	}
}

// UpdateGroupGeneralPolicyInvoker 更新代码组的基本设置信息
func (c *CodeArtsRepoClient) UpdateGroupGeneralPolicyInvoker(request *model.UpdateGroupGeneralPolicyRequest) *UpdateGroupGeneralPolicyInvoker {
	requestDef := GenReqDefForUpdateGroupGeneralPolicy()
	return &UpdateGroupGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupWatermark 更新代码组水印设置
//
// 更新代码组水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupWatermark(request *model.UpdateGroupWatermarkRequest) (*model.UpdateGroupWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateGroupWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupWatermarkResponse), nil
	}
}

// UpdateGroupWatermarkInvoker 更新代码组水印设置
func (c *CodeArtsRepoClient) UpdateGroupWatermarkInvoker(request *model.UpdateGroupWatermarkRequest) *UpdateGroupWatermarkInvoker {
	requestDef := GenReqDefForUpdateGroupWatermark()
	return &UpdateGroupWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRepositoryMembers 批量添加仓库成员
//
// 批量添加仓库成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddRepositoryMembers(request *model.AddRepositoryMembersRequest) (*model.AddRepositoryMembersResponse, error) {
	requestDef := GenReqDefForAddRepositoryMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepositoryMembersResponse), nil
	}
}

// AddRepositoryMembersInvoker 批量添加仓库成员
func (c *CodeArtsRepoClient) AddRepositoryMembersInvoker(request *model.AddRepositoryMembersRequest) *AddRepositoryMembersInvoker {
	requestDef := GenReqDefForAddRepositoryMembers()
	return &AddRepositoryMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupProtectedRefsUserGroups 获取代码组下成员组列表
//
// 获取代码组下成员组列表(保护分支保护Tags设置中使用)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupProtectedRefsUserGroups(request *model.ListGroupProtectedRefsUserGroupsRequest) (*model.ListGroupProtectedRefsUserGroupsResponse, error) {
	requestDef := GenReqDefForListGroupProtectedRefsUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupProtectedRefsUserGroupsResponse), nil
	}
}

// ListGroupProtectedRefsUserGroupsInvoker 获取代码组下成员组列表
func (c *CodeArtsRepoClient) ListGroupProtectedRefsUserGroupsInvoker(request *model.ListGroupProtectedRefsUserGroupsRequest) *ListGroupProtectedRefsUserGroupsInvoker {
	requestDef := GenReqDefForListGroupProtectedRefsUserGroups()
	return &ListGroupProtectedRefsUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 获取仓库成员列表
//
// 获取仓库成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 获取仓库成员列表
func (c *CodeArtsRepoClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductPermissionResourcesGrantedUsers 获取项目下成员列表
//
// 获取项目下成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProductPermissionResourcesGrantedUsers(request *model.ListProductPermissionResourcesGrantedUsersRequest) (*model.ListProductPermissionResourcesGrantedUsersResponse, error) {
	requestDef := GenReqDefForListProductPermissionResourcesGrantedUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductPermissionResourcesGrantedUsersResponse), nil
	}
}

// ListProductPermissionResourcesGrantedUsersInvoker 获取项目下成员列表
func (c *CodeArtsRepoClient) ListProductPermissionResourcesGrantedUsersInvoker(request *model.ListProductPermissionResourcesGrantedUsersRequest) *ListProductPermissionResourcesGrantedUsersInvoker {
	requestDef := GenReqDefForListProductPermissionResourcesGrantedUsers()
	return &ListProductPermissionResourcesGrantedUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectProtectedRefsUserGroups 获取项目下成员组列表
//
// 获取项目下成员组列表(保护分支保护Tags设置中使用)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectProtectedRefsUserGroups(request *model.ListProjectProtectedRefsUserGroupsRequest) (*model.ListProjectProtectedRefsUserGroupsResponse, error) {
	requestDef := GenReqDefForListProjectProtectedRefsUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectProtectedRefsUserGroupsResponse), nil
	}
}

// ListProjectProtectedRefsUserGroupsInvoker 获取项目下成员组列表
func (c *CodeArtsRepoClient) ListProjectProtectedRefsUserGroupsInvoker(request *model.ListProjectProtectedRefsUserGroupsRequest) *ListProjectProtectedRefsUserGroupsInvoker {
	requestDef := GenReqDefForListProjectProtectedRefsUserGroups()
	return &ListProjectProtectedRefsUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryProtectedRefsUserGroups 获取仓库下成员组列表
//
// 获取仓库下成员组列表(保护分支保护Tags设置中使用)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryProtectedRefsUserGroups(request *model.ListRepositoryProtectedRefsUserGroupsRequest) (*model.ListRepositoryProtectedRefsUserGroupsResponse, error) {
	requestDef := GenReqDefForListRepositoryProtectedRefsUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryProtectedRefsUserGroupsResponse), nil
	}
}

// ListRepositoryProtectedRefsUserGroupsInvoker 获取仓库下成员组列表
func (c *CodeArtsRepoClient) ListRepositoryProtectedRefsUserGroupsInvoker(request *model.ListRepositoryProtectedRefsUserGroupsRequest) *ListRepositoryProtectedRefsUserGroupsInvoker {
	requestDef := GenReqDefForListRepositoryProtectedRefsUserGroups()
	return &ListRepositoryProtectedRefsUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryUserGroups 获取成员组列表
//
// 获取成员组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryUserGroups(request *model.ListRepositoryUserGroupsRequest) (*model.ListRepositoryUserGroupsResponse, error) {
	requestDef := GenReqDefForListRepositoryUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryUserGroupsResponse), nil
	}
}

// ListRepositoryUserGroupsInvoker 获取成员组列表
func (c *CodeArtsRepoClient) ListRepositoryUserGroupsInvoker(request *model.ListRepositoryUserGroupsRequest) *ListRepositoryUserGroupsInvoker {
	requestDef := GenReqDefForListRepositoryUserGroups()
	return &ListRepositoryUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApprovalMergeRequest 审核合并请求
//
// 审核合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ApprovalMergeRequest(request *model.ApprovalMergeRequestRequest) (*model.ApprovalMergeRequestResponse, error) {
	requestDef := GenReqDefForApprovalMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApprovalMergeRequestResponse), nil
	}
}

// ApprovalMergeRequestInvoker 审核合并请求
func (c *CodeArtsRepoClient) ApprovalMergeRequestInvoker(request *model.ApprovalMergeRequestRequest) *ApprovalMergeRequestInvoker {
	requestDef := GenReqDefForApprovalMergeRequest()
	return &ApprovalMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCherryPickMergeRequest Cherry pick合并请求
//
// Cherry pick合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateCherryPickMergeRequest(request *model.CreateCherryPickMergeRequestRequest) (*model.CreateCherryPickMergeRequestResponse, error) {
	requestDef := GenReqDefForCreateCherryPickMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCherryPickMergeRequestResponse), nil
	}
}

// CreateCherryPickMergeRequestInvoker Cherry pick合并请求
func (c *CodeArtsRepoClient) CreateCherryPickMergeRequestInvoker(request *model.CreateCherryPickMergeRequestRequest) *CreateCherryPickMergeRequestInvoker {
	requestDef := GenReqDefForCreateCherryPickMergeRequest()
	return &CreateCherryPickMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupMergeRequestApproverSetting 创建代码组合并请求审核设置
//
// 创建代码组合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateGroupMergeRequestApproverSetting(request *model.CreateGroupMergeRequestApproverSettingRequest) (*model.CreateGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupMergeRequestApproverSettingResponse), nil
	}
}

// CreateGroupMergeRequestApproverSettingInvoker 创建代码组合并请求审核设置
func (c *CodeArtsRepoClient) CreateGroupMergeRequestApproverSettingInvoker(request *model.CreateGroupMergeRequestApproverSettingRequest) *CreateGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateGroupMergeRequestApproverSetting()
	return &CreateGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupMergeRequestTemplate 创建代码组合并请求模板
//
// 创建代码组合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateGroupMergeRequestTemplate(request *model.CreateGroupMergeRequestTemplateRequest) (*model.CreateGroupMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForCreateGroupMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupMergeRequestTemplateResponse), nil
	}
}

// CreateGroupMergeRequestTemplateInvoker 创建代码组合并请求模板
func (c *CodeArtsRepoClient) CreateGroupMergeRequestTemplateInvoker(request *model.CreateGroupMergeRequestTemplateRequest) *CreateGroupMergeRequestTemplateInvoker {
	requestDef := GenReqDefForCreateGroupMergeRequestTemplate()
	return &CreateGroupMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequest 创建合并请求
//
// 创建合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateMergeRequest(request *model.CreateMergeRequestRequest) (*model.CreateMergeRequestResponse, error) {
	requestDef := GenReqDefForCreateMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestResponse), nil
	}
}

// CreateMergeRequestInvoker 创建合并请求
func (c *CodeArtsRepoClient) CreateMergeRequestInvoker(request *model.CreateMergeRequestRequest) *CreateMergeRequestInvoker {
	requestDef := GenReqDefForCreateMergeRequest()
	return &CreateMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestApproverSetting 创建合并请求审核设置
//
// 创建合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateMergeRequestApproverSetting(request *model.CreateMergeRequestApproverSettingRequest) (*model.CreateMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestApproverSettingResponse), nil
	}
}

// CreateMergeRequestApproverSettingInvoker 创建合并请求审核设置
func (c *CodeArtsRepoClient) CreateMergeRequestApproverSettingInvoker(request *model.CreateMergeRequestApproverSettingRequest) *CreateMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateMergeRequestApproverSetting()
	return &CreateMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMergeRequestTemplate 创建合并请求模板
//
// 创建合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateMergeRequestTemplate(request *model.CreateMergeRequestTemplateRequest) (*model.CreateMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForCreateMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeRequestTemplateResponse), nil
	}
}

// CreateMergeRequestTemplateInvoker 创建合并请求模板
func (c *CodeArtsRepoClient) CreateMergeRequestTemplateInvoker(request *model.CreateMergeRequestTemplateRequest) *CreateMergeRequestTemplateInvoker {
	requestDef := GenReqDefForCreateMergeRequestTemplate()
	return &CreateMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectMergeRequestApproverSetting 创建项目合并请求审核设置
//
// 创建项目合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateProjectMergeRequestApproverSetting(request *model.CreateProjectMergeRequestApproverSettingRequest) (*model.CreateProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForCreateProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectMergeRequestApproverSettingResponse), nil
	}
}

// CreateProjectMergeRequestApproverSettingInvoker 创建项目合并请求审核设置
func (c *CodeArtsRepoClient) CreateProjectMergeRequestApproverSettingInvoker(request *model.CreateProjectMergeRequestApproverSettingRequest) *CreateProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForCreateProjectMergeRequestApproverSetting()
	return &CreateProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectMergeRequestTemplate 创建项目合并请求模板
//
// 创建项目合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateProjectMergeRequestTemplate(request *model.CreateProjectMergeRequestTemplateRequest) (*model.CreateProjectMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForCreateProjectMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectMergeRequestTemplateResponse), nil
	}
}

// CreateProjectMergeRequestTemplateInvoker 创建项目合并请求模板
func (c *CodeArtsRepoClient) CreateProjectMergeRequestTemplateInvoker(request *model.CreateProjectMergeRequestTemplateRequest) *CreateProjectMergeRequestTemplateInvoker {
	requestDef := GenReqDefForCreateProjectMergeRequestTemplate()
	return &CreateProjectMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupMergeRequestApproverSetting 删除代码组合并请求审核配置
//
// 删除代码组合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteGroupMergeRequestApproverSetting(request *model.DeleteGroupMergeRequestApproverSettingRequest) (*model.DeleteGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupMergeRequestApproverSettingResponse), nil
	}
}

// DeleteGroupMergeRequestApproverSettingInvoker 删除代码组合并请求审核配置
func (c *CodeArtsRepoClient) DeleteGroupMergeRequestApproverSettingInvoker(request *model.DeleteGroupMergeRequestApproverSettingRequest) *DeleteGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteGroupMergeRequestApproverSetting()
	return &DeleteGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupMergeRequestTemplate 删除代码组合并请求模板
//
// 删除代码组合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteGroupMergeRequestTemplate(request *model.DeleteGroupMergeRequestTemplateRequest) (*model.DeleteGroupMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForDeleteGroupMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupMergeRequestTemplateResponse), nil
	}
}

// DeleteGroupMergeRequestTemplateInvoker 删除代码组合并请求模板
func (c *CodeArtsRepoClient) DeleteGroupMergeRequestTemplateInvoker(request *model.DeleteGroupMergeRequestTemplateRequest) *DeleteGroupMergeRequestTemplateInvoker {
	requestDef := GenReqDefForDeleteGroupMergeRequestTemplate()
	return &DeleteGroupMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestApproverSetting 删除合并请求审核配置
//
// 删除合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteMergeRequestApproverSetting(request *model.DeleteMergeRequestApproverSettingRequest) (*model.DeleteMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestApproverSettingResponse), nil
	}
}

// DeleteMergeRequestApproverSettingInvoker 删除合并请求审核配置
func (c *CodeArtsRepoClient) DeleteMergeRequestApproverSettingInvoker(request *model.DeleteMergeRequestApproverSettingRequest) *DeleteMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteMergeRequestApproverSetting()
	return &DeleteMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestTemplate 删除合并请求模板
//
// 删除合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteMergeRequestTemplate(request *model.DeleteMergeRequestTemplateRequest) (*model.DeleteMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestTemplateResponse), nil
	}
}

// DeleteMergeRequestTemplateInvoker 删除合并请求模板
func (c *CodeArtsRepoClient) DeleteMergeRequestTemplateInvoker(request *model.DeleteMergeRequestTemplateRequest) *DeleteMergeRequestTemplateInvoker {
	requestDef := GenReqDefForDeleteMergeRequestTemplate()
	return &DeleteMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMergeRequestVote 删除合并请求打分
//
// 删除合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteMergeRequestVote(request *model.DeleteMergeRequestVoteRequest) (*model.DeleteMergeRequestVoteResponse, error) {
	requestDef := GenReqDefForDeleteMergeRequestVote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeRequestVoteResponse), nil
	}
}

// DeleteMergeRequestVoteInvoker 删除合并请求打分
func (c *CodeArtsRepoClient) DeleteMergeRequestVoteInvoker(request *model.DeleteMergeRequestVoteRequest) *DeleteMergeRequestVoteInvoker {
	requestDef := GenReqDefForDeleteMergeRequestVote()
	return &DeleteMergeRequestVoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProjectMergeRequestApproverSetting 删除项目合并请求审核配置
//
// 删除项目合并请求审核配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteProjectMergeRequestApproverSetting(request *model.DeleteProjectMergeRequestApproverSettingRequest) (*model.DeleteProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForDeleteProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectMergeRequestApproverSettingResponse), nil
	}
}

// DeleteProjectMergeRequestApproverSettingInvoker 删除项目合并请求审核配置
func (c *CodeArtsRepoClient) DeleteProjectMergeRequestApproverSettingInvoker(request *model.DeleteProjectMergeRequestApproverSettingRequest) *DeleteProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForDeleteProjectMergeRequestApproverSetting()
	return &DeleteProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProjectMergeRequestTemplate 删除项目合并请求模板
//
// 删除项目合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteProjectMergeRequestTemplate(request *model.DeleteProjectMergeRequestTemplateRequest) (*model.DeleteProjectMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForDeleteProjectMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectMergeRequestTemplateResponse), nil
	}
}

// DeleteProjectMergeRequestTemplateInvoker 删除项目合并请求模板
func (c *CodeArtsRepoClient) DeleteProjectMergeRequestTemplateInvoker(request *model.DeleteProjectMergeRequestTemplateRequest) *DeleteProjectMergeRequestTemplateInvoker {
	requestDef := GenReqDefForDeleteProjectMergeRequestTemplate()
	return &DeleteProjectMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportMergeRequest 导入合并请求
//
// 导入合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ImportMergeRequest(request *model.ImportMergeRequestRequest) (*model.ImportMergeRequestResponse, error) {
	requestDef := GenReqDefForImportMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportMergeRequestResponse), nil
	}
}

// ImportMergeRequestInvoker 导入合并请求
func (c *CodeArtsRepoClient) ImportMergeRequestInvoker(request *model.ImportMergeRequestRequest) *ImportMergeRequestInvoker {
	requestDef := GenReqDefForImportMergeRequest()
	return &ImportMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommitAssociatedMergeRequests 获取提交关联的合并请求
//
// 获取提交关联的合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListCommitAssociatedMergeRequests(request *model.ListCommitAssociatedMergeRequestsRequest) (*model.ListCommitAssociatedMergeRequestsResponse, error) {
	requestDef := GenReqDefForListCommitAssociatedMergeRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitAssociatedMergeRequestsResponse), nil
	}
}

// ListCommitAssociatedMergeRequestsInvoker 获取提交关联的合并请求
func (c *CodeArtsRepoClient) ListCommitAssociatedMergeRequestsInvoker(request *model.ListCommitAssociatedMergeRequestsRequest) *ListCommitAssociatedMergeRequestsInvoker {
	requestDef := GenReqDefForListCommitAssociatedMergeRequests()
	return &ListCommitAssociatedMergeRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiscussionTemplates 获取检视意见模板列表
//
// 获取检视意见模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListDiscussionTemplates(request *model.ListDiscussionTemplatesRequest) (*model.ListDiscussionTemplatesResponse, error) {
	requestDef := GenReqDefForListDiscussionTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiscussionTemplatesResponse), nil
	}
}

// ListDiscussionTemplatesInvoker 获取检视意见模板列表
func (c *CodeArtsRepoClient) ListDiscussionTemplatesInvoker(request *model.ListDiscussionTemplatesRequest) *ListDiscussionTemplatesInvoker {
	requestDef := GenReqDefForListDiscussionTemplates()
	return &ListDiscussionTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestApproverSettings 获取代码组合并请求审核设置列表
//
// 获取代码组合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupMergeRequestApproverSettings(request *model.ListGroupMergeRequestApproverSettingsRequest) (*model.ListGroupMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestApproverSettingsResponse), nil
	}
}

// ListGroupMergeRequestApproverSettingsInvoker 获取代码组合并请求审核设置列表
func (c *CodeArtsRepoClient) ListGroupMergeRequestApproverSettingsInvoker(request *model.ListGroupMergeRequestApproverSettingsRequest) *ListGroupMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListGroupMergeRequestApproverSettings()
	return &ListGroupMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestCanBeAssignedReviewers 获取代码组检视人
//
// 获取代码组检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupMergeRequestCanBeAssignedReviewers(request *model.ListGroupMergeRequestCanBeAssignedReviewersRequest) (*model.ListGroupMergeRequestCanBeAssignedReviewersResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestCanBeAssignedReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

// ListGroupMergeRequestCanBeAssignedReviewersInvoker 获取代码组检视人
func (c *CodeArtsRepoClient) ListGroupMergeRequestCanBeAssignedReviewersInvoker(request *model.ListGroupMergeRequestCanBeAssignedReviewersRequest) *ListGroupMergeRequestCanBeAssignedReviewersInvoker {
	requestDef := GenReqDefForListGroupMergeRequestCanBeAssignedReviewers()
	return &ListGroupMergeRequestCanBeAssignedReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestTemplates 获取代码组合并请求模板列表
//
// 获取代码组合并请求模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupMergeRequestTemplates(request *model.ListGroupMergeRequestTemplatesRequest) (*model.ListGroupMergeRequestTemplatesResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestTemplatesResponse), nil
	}
}

// ListGroupMergeRequestTemplatesInvoker 获取代码组合并请求模板列表
func (c *CodeArtsRepoClient) ListGroupMergeRequestTemplatesInvoker(request *model.ListGroupMergeRequestTemplatesRequest) *ListGroupMergeRequestTemplatesInvoker {
	requestDef := GenReqDefForListGroupMergeRequestTemplates()
	return &ListGroupMergeRequestTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMergeRequestValidAssignedCandidates 获取代码组审核人或合并人
//
// 获取代码组审核人或合并人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupMergeRequestValidAssignedCandidates(request *model.ListGroupMergeRequestValidAssignedCandidatesRequest) (*model.ListGroupMergeRequestValidAssignedCandidatesResponse, error) {
	requestDef := GenReqDefForListGroupMergeRequestValidAssignedCandidates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMergeRequestValidAssignedCandidatesResponse), nil
	}
}

// ListGroupMergeRequestValidAssignedCandidatesInvoker 获取代码组审核人或合并人
func (c *CodeArtsRepoClient) ListGroupMergeRequestValidAssignedCandidatesInvoker(request *model.ListGroupMergeRequestValidAssignedCandidatesRequest) *ListGroupMergeRequestValidAssignedCandidatesInvoker {
	requestDef := GenReqDefForListGroupMergeRequestValidAssignedCandidates()
	return &ListGroupMergeRequestValidAssignedCandidatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestApproverSettings 获取合并请求审核设置列表
//
// 获取合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestApproverSettings(request *model.ListMergeRequestApproverSettingsRequest) (*model.ListMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestApproverSettingsResponse), nil
	}
}

// ListMergeRequestApproverSettingsInvoker 获取合并请求审核设置列表
func (c *CodeArtsRepoClient) ListMergeRequestApproverSettingsInvoker(request *model.ListMergeRequestApproverSettingsRequest) *ListMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListMergeRequestApproverSettings()
	return &ListMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestApprovers 获取合并请求审核人列表
//
// 获取合并请求审核人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestApprovers(request *model.ListMergeRequestApproversRequest) (*model.ListMergeRequestApproversResponse, error) {
	requestDef := GenReqDefForListMergeRequestApprovers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestApproversResponse), nil
	}
}

// ListMergeRequestApproversInvoker 获取合并请求审核人列表
func (c *CodeArtsRepoClient) ListMergeRequestApproversInvoker(request *model.ListMergeRequestApproversRequest) *ListMergeRequestApproversInvoker {
	requestDef := GenReqDefForListMergeRequestApprovers()
	return &ListMergeRequestApproversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestChanges 获取合并请求文件变更列表
//
// 获取合并请求文件变更列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestChanges(request *model.ListMergeRequestChangesRequest) (*model.ListMergeRequestChangesResponse, error) {
	requestDef := GenReqDefForListMergeRequestChanges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestChangesResponse), nil
	}
}

// ListMergeRequestChangesInvoker 获取合并请求文件变更列表
func (c *CodeArtsRepoClient) ListMergeRequestChangesInvoker(request *model.ListMergeRequestChangesRequest) *ListMergeRequestChangesInvoker {
	requestDef := GenReqDefForListMergeRequestChanges()
	return &ListMergeRequestChangesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestChangesTrees 获取合并请求文件变更列表树
//
// 获取合并请求文件变更列表树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestChangesTrees(request *model.ListMergeRequestChangesTreesRequest) (*model.ListMergeRequestChangesTreesResponse, error) {
	requestDef := GenReqDefForListMergeRequestChangesTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestChangesTreesResponse), nil
	}
}

// ListMergeRequestChangesTreesInvoker 获取合并请求文件变更列表树
func (c *CodeArtsRepoClient) ListMergeRequestChangesTreesInvoker(request *model.ListMergeRequestChangesTreesRequest) *ListMergeRequestChangesTreesInvoker {
	requestDef := GenReqDefForListMergeRequestChangesTrees()
	return &ListMergeRequestChangesTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestCommits 获取合并请求commit列表
//
// 获取合并请求commit列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestCommits(request *model.ListMergeRequestCommitsRequest) (*model.ListMergeRequestCommitsResponse, error) {
	requestDef := GenReqDefForListMergeRequestCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestCommitsResponse), nil
	}
}

// ListMergeRequestCommitsInvoker 获取合并请求commit列表
func (c *CodeArtsRepoClient) ListMergeRequestCommitsInvoker(request *model.ListMergeRequestCommitsRequest) *ListMergeRequestCommitsInvoker {
	requestDef := GenReqDefForListMergeRequestCommits()
	return &ListMergeRequestCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestConflictFiles 获取所有的冲突文件
//
// 获取所有的冲突文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestConflictFiles(request *model.ListMergeRequestConflictFilesRequest) (*model.ListMergeRequestConflictFilesResponse, error) {
	requestDef := GenReqDefForListMergeRequestConflictFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestConflictFilesResponse), nil
	}
}

// ListMergeRequestConflictFilesInvoker 获取所有的冲突文件
func (c *CodeArtsRepoClient) ListMergeRequestConflictFilesInvoker(request *model.ListMergeRequestConflictFilesRequest) *ListMergeRequestConflictFilesInvoker {
	requestDef := GenReqDefForListMergeRequestConflictFiles()
	return &ListMergeRequestConflictFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestEvaluations 获取合并请求评价列表
//
// 获取合并请求评价列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestEvaluations(request *model.ListMergeRequestEvaluationsRequest) (*model.ListMergeRequestEvaluationsResponse, error) {
	requestDef := GenReqDefForListMergeRequestEvaluations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestEvaluationsResponse), nil
	}
}

// ListMergeRequestEvaluationsInvoker 获取合并请求评价列表
func (c *CodeArtsRepoClient) ListMergeRequestEvaluationsInvoker(request *model.ListMergeRequestEvaluationsRequest) *ListMergeRequestEvaluationsInvoker {
	requestDef := GenReqDefForListMergeRequestEvaluations()
	return &ListMergeRequestEvaluationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestParticipants 获取合并请求参与者
//
// 获取合并请求参与者
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestParticipants(request *model.ListMergeRequestParticipantsRequest) (*model.ListMergeRequestParticipantsResponse, error) {
	requestDef := GenReqDefForListMergeRequestParticipants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestParticipantsResponse), nil
	}
}

// ListMergeRequestParticipantsInvoker 获取合并请求参与者
func (c *CodeArtsRepoClient) ListMergeRequestParticipantsInvoker(request *model.ListMergeRequestParticipantsRequest) *ListMergeRequestParticipantsInvoker {
	requestDef := GenReqDefForListMergeRequestParticipants()
	return &ListMergeRequestParticipantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestReviewers 获取合并请求检视人列表
//
// 获取合并请求检视人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestReviewers(request *model.ListMergeRequestReviewersRequest) (*model.ListMergeRequestReviewersResponse, error) {
	requestDef := GenReqDefForListMergeRequestReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestReviewersResponse), nil
	}
}

// ListMergeRequestReviewersInvoker 获取合并请求检视人列表
func (c *CodeArtsRepoClient) ListMergeRequestReviewersInvoker(request *model.ListMergeRequestReviewersRequest) *ListMergeRequestReviewersInvoker {
	requestDef := GenReqDefForListMergeRequestReviewers()
	return &ListMergeRequestReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestTemplates 获取合并请求模板列表
//
// 获取合并请求模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestTemplates(request *model.ListMergeRequestTemplatesRequest) (*model.ListMergeRequestTemplatesResponse, error) {
	requestDef := GenReqDefForListMergeRequestTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestTemplatesResponse), nil
	}
}

// ListMergeRequestTemplatesInvoker 获取合并请求模板列表
func (c *CodeArtsRepoClient) ListMergeRequestTemplatesInvoker(request *model.ListMergeRequestTemplatesRequest) *ListMergeRequestTemplatesInvoker {
	requestDef := GenReqDefForListMergeRequestTemplates()
	return &ListMergeRequestTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestValidAssignedCandidates 获取可选的合并请求检视人
//
// 获取可选的合并请求检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestValidAssignedCandidates(request *model.ListMergeRequestValidAssignedCandidatesRequest) (*model.ListMergeRequestValidAssignedCandidatesResponse, error) {
	requestDef := GenReqDefForListMergeRequestValidAssignedCandidates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestValidAssignedCandidatesResponse), nil
	}
}

// ListMergeRequestValidAssignedCandidatesInvoker 获取可选的合并请求检视人
func (c *CodeArtsRepoClient) ListMergeRequestValidAssignedCandidatesInvoker(request *model.ListMergeRequestValidAssignedCandidatesRequest) *ListMergeRequestValidAssignedCandidatesInvoker {
	requestDef := GenReqDefForListMergeRequestValidAssignedCandidates()
	return &ListMergeRequestValidAssignedCandidatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMergeRequestVersions 获取文件变更历史版本列表
//
// 获取文件变更历史版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListMergeRequestVersions(request *model.ListMergeRequestVersionsRequest) (*model.ListMergeRequestVersionsResponse, error) {
	requestDef := GenReqDefForListMergeRequestVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeRequestVersionsResponse), nil
	}
}

// ListMergeRequestVersionsInvoker 获取文件变更历史版本列表
func (c *CodeArtsRepoClient) ListMergeRequestVersionsInvoker(request *model.ListMergeRequestVersionsRequest) *ListMergeRequestVersionsInvoker {
	requestDef := GenReqDefForListMergeRequestVersions()
	return &ListMergeRequestVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPersonalMergeRequests 获取个人首页mr列表
//
// 获取个人首页mr列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListPersonalMergeRequests(request *model.ListPersonalMergeRequestsRequest) (*model.ListPersonalMergeRequestsResponse, error) {
	requestDef := GenReqDefForListPersonalMergeRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPersonalMergeRequestsResponse), nil
	}
}

// ListPersonalMergeRequestsInvoker 获取个人首页mr列表
func (c *CodeArtsRepoClient) ListPersonalMergeRequestsInvoker(request *model.ListPersonalMergeRequestsRequest) *ListPersonalMergeRequestsInvoker {
	requestDef := GenReqDefForListPersonalMergeRequests()
	return &ListPersonalMergeRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestApproverSettings 获取项目合并请求审核设置列表
//
// 获取项目合并请求审核设置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectMergeRequestApproverSettings(request *model.ListProjectMergeRequestApproverSettingsRequest) (*model.ListProjectMergeRequestApproverSettingsResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestApproverSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestApproverSettingsResponse), nil
	}
}

// ListProjectMergeRequestApproverSettingsInvoker 获取项目合并请求审核设置列表
func (c *CodeArtsRepoClient) ListProjectMergeRequestApproverSettingsInvoker(request *model.ListProjectMergeRequestApproverSettingsRequest) *ListProjectMergeRequestApproverSettingsInvoker {
	requestDef := GenReqDefForListProjectMergeRequestApproverSettings()
	return &ListProjectMergeRequestApproverSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestCanBeAssignedReviewers 获取项目检视人
//
// 获取代码组检视人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectMergeRequestCanBeAssignedReviewers(request *model.ListProjectMergeRequestCanBeAssignedReviewersRequest) (*model.ListProjectMergeRequestCanBeAssignedReviewersResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

// ListProjectMergeRequestCanBeAssignedReviewersInvoker 获取项目检视人
func (c *CodeArtsRepoClient) ListProjectMergeRequestCanBeAssignedReviewersInvoker(request *model.ListProjectMergeRequestCanBeAssignedReviewersRequest) *ListProjectMergeRequestCanBeAssignedReviewersInvoker {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedReviewers()
	return &ListProjectMergeRequestCanBeAssignedReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestCanBeAssignedUsers 获取项目审核人或合并人
//
// 获取代码组审核人或合并人
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectMergeRequestCanBeAssignedUsers(request *model.ListProjectMergeRequestCanBeAssignedUsersRequest) (*model.ListProjectMergeRequestCanBeAssignedUsersResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestCanBeAssignedUsersResponse), nil
	}
}

// ListProjectMergeRequestCanBeAssignedUsersInvoker 获取项目审核人或合并人
func (c *CodeArtsRepoClient) ListProjectMergeRequestCanBeAssignedUsersInvoker(request *model.ListProjectMergeRequestCanBeAssignedUsersRequest) *ListProjectMergeRequestCanBeAssignedUsersInvoker {
	requestDef := GenReqDefForListProjectMergeRequestCanBeAssignedUsers()
	return &ListProjectMergeRequestCanBeAssignedUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMergeRequestTemplates 获取项目合并请求模板列表
//
// 获取项目合并请求模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectMergeRequestTemplates(request *model.ListProjectMergeRequestTemplatesRequest) (*model.ListProjectMergeRequestTemplatesResponse, error) {
	requestDef := GenReqDefForListProjectMergeRequestTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMergeRequestTemplatesResponse), nil
	}
}

// ListProjectMergeRequestTemplatesInvoker 获取项目合并请求模板列表
func (c *CodeArtsRepoClient) ListProjectMergeRequestTemplatesInvoker(request *model.ListProjectMergeRequestTemplatesRequest) *ListProjectMergeRequestTemplatesInvoker {
	requestDef := GenReqDefForListProjectMergeRequestTemplates()
	return &ListProjectMergeRequestTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryMergeRequests 获取仓库MR列表
//
// 获取仓库MR列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryMergeRequests(request *model.ListRepositoryMergeRequestsRequest) (*model.ListRepositoryMergeRequestsResponse, error) {
	requestDef := GenReqDefForListRepositoryMergeRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryMergeRequestsResponse), nil
	}
}

// ListRepositoryMergeRequestsInvoker 获取仓库MR列表
func (c *CodeArtsRepoClient) ListRepositoryMergeRequestsInvoker(request *model.ListRepositoryMergeRequestsRequest) *ListRepositoryMergeRequestsInvoker {
	requestDef := GenReqDefForListRepositoryMergeRequests()
	return &ListRepositoryMergeRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MergeMergeRequest 合入合并请求
//
// 合入合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) MergeMergeRequest(request *model.MergeMergeRequestRequest) (*model.MergeMergeRequestResponse, error) {
	requestDef := GenReqDefForMergeMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MergeMergeRequestResponse), nil
	}
}

// MergeMergeRequestInvoker 合入合并请求
func (c *CodeArtsRepoClient) MergeMergeRequestInvoker(request *model.MergeMergeRequestRequest) *MergeMergeRequestInvoker {
	requestDef := GenReqDefForMergeMergeRequest()
	return &MergeMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebaseMergeRequestForOpenApi 变基合并请求
//
// 变基合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RebaseMergeRequestForOpenApi(request *model.RebaseMergeRequestForOpenApiRequest) (*model.RebaseMergeRequestForOpenApiResponse, error) {
	requestDef := GenReqDefForRebaseMergeRequestForOpenApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebaseMergeRequestForOpenApiResponse), nil
	}
}

// RebaseMergeRequestForOpenApiInvoker 变基合并请求
func (c *CodeArtsRepoClient) RebaseMergeRequestForOpenApiInvoker(request *model.RebaseMergeRequestForOpenApiRequest) *RebaseMergeRequestForOpenApiInvoker {
	requestDef := GenReqDefForRebaseMergeRequestForOpenApi()
	return &RebaseMergeRequestForOpenApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResolveMergeRequestConflicts 在线解决合并请求冲突
//
// 在线解决合并请求冲突
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ResolveMergeRequestConflicts(request *model.ResolveMergeRequestConflictsRequest) (*model.ResolveMergeRequestConflictsResponse, error) {
	requestDef := GenReqDefForResolveMergeRequestConflicts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResolveMergeRequestConflictsResponse), nil
	}
}

// ResolveMergeRequestConflictsInvoker 在线解决合并请求冲突
func (c *CodeArtsRepoClient) ResolveMergeRequestConflictsInvoker(request *model.ResolveMergeRequestConflictsRequest) *ResolveMergeRequestConflictsInvoker {
	requestDef := GenReqDefForResolveMergeRequestConflicts()
	return &ResolveMergeRequestConflictsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReviewMergeRequest 检视合并请求
//
// 检视合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ReviewMergeRequest(request *model.ReviewMergeRequestRequest) (*model.ReviewMergeRequestResponse, error) {
	requestDef := GenReqDefForReviewMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReviewMergeRequestResponse), nil
	}
}

// ReviewMergeRequestInvoker 检视合并请求
func (c *CodeArtsRepoClient) ReviewMergeRequestInvoker(request *model.ReviewMergeRequestRequest) *ReviewMergeRequestInvoker {
	requestDef := GenReqDefForReviewMergeRequest()
	return &ReviewMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowActualHeadPipeline 获取合并请求关联的最新流水线
//
// 获取合并请求关联的最新流水线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowActualHeadPipeline(request *model.ShowActualHeadPipelineRequest) (*model.ShowActualHeadPipelineResponse, error) {
	requestDef := GenReqDefForShowActualHeadPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowActualHeadPipelineResponse), nil
	}
}

// ShowActualHeadPipelineInvoker 获取合并请求关联的最新流水线
func (c *CodeArtsRepoClient) ShowActualHeadPipelineInvoker(request *model.ShowActualHeadPipelineRequest) *ShowActualHeadPipelineInvoker {
	requestDef := GenReqDefForShowActualHeadPipeline()
	return &ShowActualHeadPipelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAverageEvaluation 获取合并请求的平均评价
//
// 获取合并请求的平均评价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowAverageEvaluation(request *model.ShowAverageEvaluationRequest) (*model.ShowAverageEvaluationResponse, error) {
	requestDef := GenReqDefForShowAverageEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAverageEvaluationResponse), nil
	}
}

// ShowAverageEvaluationInvoker 获取合并请求的平均评价
func (c *CodeArtsRepoClient) ShowAverageEvaluationInvoker(request *model.ShowAverageEvaluationRequest) *ShowAverageEvaluationInvoker {
	requestDef := GenReqDefForShowAverageEvaluation()
	return &ShowAverageEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranchConflict 获取分支代码冲突
//
// 获取分支代码冲突
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowBranchConflict(request *model.ShowBranchConflictRequest) (*model.ShowBranchConflictResponse, error) {
	requestDef := GenReqDefForShowBranchConflict()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchConflictResponse), nil
	}
}

// ShowBranchConflictInvoker 获取分支代码冲突
func (c *CodeArtsRepoClient) ShowBranchConflictInvoker(request *model.ShowBranchConflictRequest) *ShowBranchConflictInvoker {
	requestDef := GenReqDefForShowBranchConflict()
	return &ShowBranchConflictInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitCommentsByLine 获取代码页单个提交下文件的检视意见
//
// 获取代码页单个提交下文件的检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowCommitCommentsByLine(request *model.ShowCommitCommentsByLineRequest) (*model.ShowCommitCommentsByLineResponse, error) {
	requestDef := GenReqDefForShowCommitCommentsByLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitCommentsByLineResponse), nil
	}
}

// ShowCommitCommentsByLineInvoker 获取代码页单个提交下文件的检视意见
func (c *CodeArtsRepoClient) ShowCommitCommentsByLineInvoker(request *model.ShowCommitCommentsByLineRequest) *ShowCommitCommentsByLineInvoker {
	requestDef := GenReqDefForShowCommitCommentsByLine()
	return &ShowCommitCommentsByLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupMergeRequestSetting 获取代码组合并请求设置
//
// 获取代码组合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupMergeRequestSetting(request *model.ShowGroupMergeRequestSettingRequest) (*model.ShowGroupMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowGroupMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupMergeRequestSettingResponse), nil
	}
}

// ShowGroupMergeRequestSettingInvoker 获取代码组合并请求设置
func (c *CodeArtsRepoClient) ShowGroupMergeRequestSettingInvoker(request *model.ShowGroupMergeRequestSettingRequest) *ShowGroupMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowGroupMergeRequestSetting()
	return &ShowGroupMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestCommentsByLine 获取合并请求文件变更页单个文件下的检视意见
//
// 获取合并请求文件变更页单个文件下的检视意见
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestCommentsByLine(request *model.ShowMergeRequestCommentsByLineRequest) (*model.ShowMergeRequestCommentsByLineResponse, error) {
	requestDef := GenReqDefForShowMergeRequestCommentsByLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestCommentsByLineResponse), nil
	}
}

// ShowMergeRequestCommentsByLineInvoker 获取合并请求文件变更页单个文件下的检视意见
func (c *CodeArtsRepoClient) ShowMergeRequestCommentsByLineInvoker(request *model.ShowMergeRequestCommentsByLineRequest) *ShowMergeRequestCommentsByLineInvoker {
	requestDef := GenReqDefForShowMergeRequestCommentsByLine()
	return &ShowMergeRequestCommentsByLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestDetail 获取MR详情
//
// 获取MR详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestDetail(request *model.ShowMergeRequestDetailRequest) (*model.ShowMergeRequestDetailResponse, error) {
	requestDef := GenReqDefForShowMergeRequestDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestDetailResponse), nil
	}
}

// ShowMergeRequestDetailInvoker 获取MR详情
func (c *CodeArtsRepoClient) ShowMergeRequestDetailInvoker(request *model.ShowMergeRequestDetailRequest) *ShowMergeRequestDetailInvoker {
	requestDef := GenReqDefForShowMergeRequestDetail()
	return &ShowMergeRequestDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestSetting 获取仓库合并请求设置
//
// 获取仓库合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestSetting(request *model.ShowMergeRequestSettingRequest) (*model.ShowMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestSettingResponse), nil
	}
}

// ShowMergeRequestSettingInvoker 获取仓库合并请求设置
func (c *CodeArtsRepoClient) ShowMergeRequestSettingInvoker(request *model.ShowMergeRequestSettingRequest) *ShowMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowMergeRequestSetting()
	return &ShowMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestTemplate 获取单个合并请求模板
//
// 获取单个合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestTemplate(request *model.ShowMergeRequestTemplateRequest) (*model.ShowMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForShowMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestTemplateResponse), nil
	}
}

// ShowMergeRequestTemplateInvoker 获取单个合并请求模板
func (c *CodeArtsRepoClient) ShowMergeRequestTemplateInvoker(request *model.ShowMergeRequestTemplateRequest) *ShowMergeRequestTemplateInvoker {
	requestDef := GenReqDefForShowMergeRequestTemplate()
	return &ShowMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeRequestVotesDetail 获取合并请求打分
//
// 获取合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeRequestVotesDetail(request *model.ShowMergeRequestVotesDetailRequest) (*model.ShowMergeRequestVotesDetailResponse, error) {
	requestDef := GenReqDefForShowMergeRequestVotesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeRequestVotesDetailResponse), nil
	}
}

// ShowMergeRequestVotesDetailInvoker 获取合并请求打分
func (c *CodeArtsRepoClient) ShowMergeRequestVotesDetailInvoker(request *model.ShowMergeRequestVotesDetailRequest) *ShowMergeRequestVotesDetailInvoker {
	requestDef := GenReqDefForShowMergeRequestVotesDetail()
	return &ShowMergeRequestVotesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMergeableStateOuter 获取合并请求的可合入状态。
//
// 获取合并请求的可合入状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowMergeableStateOuter(request *model.ShowMergeableStateOuterRequest) (*model.ShowMergeableStateOuterResponse, error) {
	requestDef := GenReqDefForShowMergeableStateOuter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMergeableStateOuterResponse), nil
	}
}

// ShowMergeableStateOuterInvoker 获取合并请求的可合入状态。
func (c *CodeArtsRepoClient) ShowMergeableStateOuterInvoker(request *model.ShowMergeableStateOuterRequest) *ShowMergeableStateOuterInvoker {
	requestDef := GenReqDefForShowMergeableStateOuter()
	return &ShowMergeableStateOuterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectMergeRequestSetting 获取项目合并请求设置
//
// 获取项目合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectMergeRequestSetting(request *model.ShowProjectMergeRequestSettingRequest) (*model.ShowProjectMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForShowProjectMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectMergeRequestSettingResponse), nil
	}
}

// ShowProjectMergeRequestSettingInvoker 获取项目合并请求设置
func (c *CodeArtsRepoClient) ShowProjectMergeRequestSettingInvoker(request *model.ShowProjectMergeRequestSettingRequest) *ShowProjectMergeRequestSettingInvoker {
	requestDef := GenReqDefForShowProjectMergeRequestSetting()
	return &ShowProjectMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryMergeRequestsStatistic 获取仓库合并请求统计数据
//
// 获取仓库合并请求统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryMergeRequestsStatistic(request *model.ShowRepositoryMergeRequestsStatisticRequest) (*model.ShowRepositoryMergeRequestsStatisticResponse, error) {
	requestDef := GenReqDefForShowRepositoryMergeRequestsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryMergeRequestsStatisticResponse), nil
	}
}

// ShowRepositoryMergeRequestsStatisticInvoker 获取仓库合并请求统计数据
func (c *CodeArtsRepoClient) ShowRepositoryMergeRequestsStatisticInvoker(request *model.ShowRepositoryMergeRequestsStatisticRequest) *ShowRepositoryMergeRequestsStatisticInvoker {
	requestDef := GenReqDefForShowRepositoryMergeRequestsStatistic()
	return &ShowRepositoryMergeRequestsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupMergeRequestApproverSetting 更新代码组合并请求审核设置
//
// 更新代码组合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupMergeRequestApproverSetting(request *model.UpdateGroupMergeRequestApproverSettingRequest) (*model.UpdateGroupMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateGroupMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupMergeRequestApproverSettingResponse), nil
	}
}

// UpdateGroupMergeRequestApproverSettingInvoker 更新代码组合并请求审核设置
func (c *CodeArtsRepoClient) UpdateGroupMergeRequestApproverSettingInvoker(request *model.UpdateGroupMergeRequestApproverSettingRequest) *UpdateGroupMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateGroupMergeRequestApproverSetting()
	return &UpdateGroupMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupMergeRequestTemplate 更新代码组合并请求模板
//
// 更新代码组合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupMergeRequestTemplate(request *model.UpdateGroupMergeRequestTemplateRequest) (*model.UpdateGroupMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForUpdateGroupMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupMergeRequestTemplateResponse), nil
	}
}

// UpdateGroupMergeRequestTemplateInvoker 更新代码组合并请求模板
func (c *CodeArtsRepoClient) UpdateGroupMergeRequestTemplateInvoker(request *model.UpdateGroupMergeRequestTemplateRequest) *UpdateGroupMergeRequestTemplateInvoker {
	requestDef := GenReqDefForUpdateGroupMergeRequestTemplate()
	return &UpdateGroupMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequest 更新合并请求
//
// 更新合并请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequest(request *model.UpdateMergeRequestRequest) (*model.UpdateMergeRequestResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestResponse), nil
	}
}

// UpdateMergeRequestInvoker 更新合并请求
func (c *CodeArtsRepoClient) UpdateMergeRequestInvoker(request *model.UpdateMergeRequestRequest) *UpdateMergeRequestInvoker {
	requestDef := GenReqDefForUpdateMergeRequest()
	return &UpdateMergeRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestApproverSetting 更新合并请求审核设置
//
// 更新合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestApproverSetting(request *model.UpdateMergeRequestApproverSettingRequest) (*model.UpdateMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestApproverSettingResponse), nil
	}
}

// UpdateMergeRequestApproverSettingInvoker 更新合并请求审核设置
func (c *CodeArtsRepoClient) UpdateMergeRequestApproverSettingInvoker(request *model.UpdateMergeRequestApproverSettingRequest) *UpdateMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateMergeRequestApproverSetting()
	return &UpdateMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestApprovers 更新合并请求的审核人列表
//
// 更新合并请求的审核人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestApprovers(request *model.UpdateMergeRequestApproversRequest) (*model.UpdateMergeRequestApproversResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestApprovers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestApproversResponse), nil
	}
}

// UpdateMergeRequestApproversInvoker 更新合并请求的审核人列表
func (c *CodeArtsRepoClient) UpdateMergeRequestApproversInvoker(request *model.UpdateMergeRequestApproversRequest) *UpdateMergeRequestApproversInvoker {
	requestDef := GenReqDefForUpdateMergeRequestApprovers()
	return &UpdateMergeRequestApproversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestReviewers 更新合并请求的检视人列表
//
// 更新合并请求的检视人列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestReviewers(request *model.UpdateMergeRequestReviewersRequest) (*model.UpdateMergeRequestReviewersResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestReviewers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestReviewersResponse), nil
	}
}

// UpdateMergeRequestReviewersInvoker 更新合并请求的检视人列表
func (c *CodeArtsRepoClient) UpdateMergeRequestReviewersInvoker(request *model.UpdateMergeRequestReviewersRequest) *UpdateMergeRequestReviewersInvoker {
	requestDef := GenReqDefForUpdateMergeRequestReviewers()
	return &UpdateMergeRequestReviewersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestSetting 更新仓库合并请求设置
//
// 更新仓库合并请求设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestSetting(request *model.UpdateMergeRequestSettingRequest) (*model.UpdateMergeRequestSettingResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestSettingResponse), nil
	}
}

// UpdateMergeRequestSettingInvoker 更新仓库合并请求设置
func (c *CodeArtsRepoClient) UpdateMergeRequestSettingInvoker(request *model.UpdateMergeRequestSettingRequest) *UpdateMergeRequestSettingInvoker {
	requestDef := GenReqDefForUpdateMergeRequestSetting()
	return &UpdateMergeRequestSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestTemplate 更新合并请求模板
//
// 更新合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestTemplate(request *model.UpdateMergeRequestTemplateRequest) (*model.UpdateMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestTemplateResponse), nil
	}
}

// UpdateMergeRequestTemplateInvoker 更新合并请求模板
func (c *CodeArtsRepoClient) UpdateMergeRequestTemplateInvoker(request *model.UpdateMergeRequestTemplateRequest) *UpdateMergeRequestTemplateInvoker {
	requestDef := GenReqDefForUpdateMergeRequestTemplate()
	return &UpdateMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMergeRequestVote 更新合并请求打分
//
// 更新合并请求打分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateMergeRequestVote(request *model.UpdateMergeRequestVoteRequest) (*model.UpdateMergeRequestVoteResponse, error) {
	requestDef := GenReqDefForUpdateMergeRequestVote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMergeRequestVoteResponse), nil
	}
}

// UpdateMergeRequestVoteInvoker 更新合并请求打分
func (c *CodeArtsRepoClient) UpdateMergeRequestVoteInvoker(request *model.UpdateMergeRequestVoteRequest) *UpdateMergeRequestVoteInvoker {
	requestDef := GenReqDefForUpdateMergeRequestVote()
	return &UpdateMergeRequestVoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectMergeRequestApproverSetting 更新项目合并请求审核设置
//
// 更新项目合并请求审核设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectMergeRequestApproverSetting(request *model.UpdateProjectMergeRequestApproverSettingRequest) (*model.UpdateProjectMergeRequestApproverSettingResponse, error) {
	requestDef := GenReqDefForUpdateProjectMergeRequestApproverSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectMergeRequestApproverSettingResponse), nil
	}
}

// UpdateProjectMergeRequestApproverSettingInvoker 更新项目合并请求审核设置
func (c *CodeArtsRepoClient) UpdateProjectMergeRequestApproverSettingInvoker(request *model.UpdateProjectMergeRequestApproverSettingRequest) *UpdateProjectMergeRequestApproverSettingInvoker {
	requestDef := GenReqDefForUpdateProjectMergeRequestApproverSetting()
	return &UpdateProjectMergeRequestApproverSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectMergeRequestTemplate 更新项目合并请求模板
//
// 更新项目合并请求模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectMergeRequestTemplate(request *model.UpdateProjectMergeRequestTemplateRequest) (*model.UpdateProjectMergeRequestTemplateResponse, error) {
	requestDef := GenReqDefForUpdateProjectMergeRequestTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectMergeRequestTemplateResponse), nil
	}
}

// UpdateProjectMergeRequestTemplateInvoker 更新项目合并请求模板
func (c *CodeArtsRepoClient) UpdateProjectMergeRequestTemplateInvoker(request *model.UpdateProjectMergeRequestTemplateRequest) *UpdateProjectMergeRequestTemplateInvoker {
	requestDef := GenReqDefForUpdateProjectMergeRequestTemplate()
	return &UpdateProjectMergeRequestTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRepositoryFilePushPermissions 批量删除仓库文件推送权限
//
// 批量删除仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchDeleteRepositoryFilePushPermissions(request *model.BatchDeleteRepositoryFilePushPermissionsRequest) (*model.BatchDeleteRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRepositoryFilePushPermissionsResponse), nil
	}
}

// BatchDeleteRepositoryFilePushPermissionsInvoker 批量删除仓库文件推送权限
func (c *CodeArtsRepoClient) BatchDeleteRepositoryFilePushPermissionsInvoker(request *model.BatchDeleteRepositoryFilePushPermissionsRequest) *BatchDeleteRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForBatchDeleteRepositoryFilePushPermissions()
	return &BatchDeleteRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateRepositoryFilePushPermissions 批量更新仓库文件推送权限
//
// 批量更新仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchUpdateRepositoryFilePushPermissions(request *model.BatchUpdateRepositoryFilePushPermissionsRequest) (*model.BatchUpdateRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForBatchUpdateRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateRepositoryFilePushPermissionsResponse), nil
	}
}

// BatchUpdateRepositoryFilePushPermissionsInvoker 批量更新仓库文件推送权限
func (c *CodeArtsRepoClient) BatchUpdateRepositoryFilePushPermissionsInvoker(request *model.BatchUpdateRepositoryFilePushPermissionsRequest) *BatchUpdateRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForBatchUpdateRepositoryFilePushPermissions()
	return &BatchUpdateRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFilePushPermission 创建仓库文件推送权限
//
// 创建仓库文件推送权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateFilePushPermission(request *model.CreateFilePushPermissionRequest) (*model.CreateFilePushPermissionResponse, error) {
	requestDef := GenReqDefForCreateFilePushPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFilePushPermissionResponse), nil
	}
}

// CreateFilePushPermissionInvoker 创建仓库文件推送权限
func (c *CodeArtsRepoClient) CreateFilePushPermissionInvoker(request *model.CreateFilePushPermissionRequest) *CreateFilePushPermissionInvoker {
	requestDef := GenReqDefForCreateFilePushPermission()
	return &CreateFilePushPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryFilePushPermissions 获取仓库文件推送权限列表
//
// 获取仓库文件推送权限列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryFilePushPermissions(request *model.ListRepositoryFilePushPermissionsRequest) (*model.ListRepositoryFilePushPermissionsResponse, error) {
	requestDef := GenReqDefForListRepositoryFilePushPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryFilePushPermissionsResponse), nil
	}
}

// ListRepositoryFilePushPermissionsInvoker 获取仓库文件推送权限列表
func (c *CodeArtsRepoClient) ListRepositoryFilePushPermissionsInvoker(request *model.ListRepositoryFilePushPermissionsRequest) *ListRepositoryFilePushPermissionsInvoker {
	requestDef := GenReqDefForListRepositoryFilePushPermissions()
	return &ListRepositoryFilePushPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryResourcePermissions 查询仓库权限矩阵配置
//
// 查询仓库权限矩阵配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryResourcePermissions(request *model.ListRepositoryResourcePermissionsRequest) (*model.ListRepositoryResourcePermissionsResponse, error) {
	requestDef := GenReqDefForListRepositoryResourcePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryResourcePermissionsResponse), nil
	}
}

// ListRepositoryResourcePermissionsInvoker 查询仓库权限矩阵配置
func (c *CodeArtsRepoClient) ListRepositoryResourcePermissionsInvoker(request *model.ListRepositoryResourcePermissionsRequest) *ListRepositoryResourcePermissionsInvoker {
	requestDef := GenReqDefForListRepositoryResourcePermissions()
	return &ListRepositoryResourcePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryPermissionInheritEnabled 查询仓库权限配置信息
//
// 查询仓库权限配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryPermissionInheritEnabled(request *model.ShowRepositoryPermissionInheritEnabledRequest) (*model.ShowRepositoryPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForShowRepositoryPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryPermissionInheritEnabledResponse), nil
	}
}

// ShowRepositoryPermissionInheritEnabledInvoker 查询仓库权限配置信息
func (c *CodeArtsRepoClient) ShowRepositoryPermissionInheritEnabledInvoker(request *model.ShowRepositoryPermissionInheritEnabledRequest) *ShowRepositoryPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForShowRepositoryPermissionInheritEnabled()
	return &ShowRepositoryPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryPermissionInheritEnabled 更新仓库权限继承配置
//
// 更新仓库权限继承配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryPermissionInheritEnabled(request *model.UpdateRepositoryPermissionInheritEnabledRequest) (*model.UpdateRepositoryPermissionInheritEnabledResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryPermissionInheritEnabled()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryPermissionInheritEnabledResponse), nil
	}
}

// UpdateRepositoryPermissionInheritEnabledInvoker 更新仓库权限继承配置
func (c *CodeArtsRepoClient) UpdateRepositoryPermissionInheritEnabledInvoker(request *model.UpdateRepositoryPermissionInheritEnabledRequest) *UpdateRepositoryPermissionInheritEnabledInvoker {
	requestDef := GenReqDefForUpdateRepositoryPermissionInheritEnabled()
	return &UpdateRepositoryPermissionInheritEnabledInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryResourcePermissions 更新仓库权限矩阵配置
//
// 更新仓库权限矩阵配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryResourcePermissions(request *model.UpdateRepositoryResourcePermissionsRequest) (*model.UpdateRepositoryResourcePermissionsResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryResourcePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryResourcePermissionsResponse), nil
	}
}

// UpdateRepositoryResourcePermissionsInvoker 更新仓库权限矩阵配置
func (c *CodeArtsRepoClient) UpdateRepositoryResourcePermissionsInvoker(request *model.UpdateRepositoryResourcePermissionsRequest) *UpdateRepositoryResourcePermissionsInvoker {
	requestDef := GenReqDefForUpdateRepositoryResourcePermissions()
	return &UpdateRepositoryResourcePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatestPipelineJobs 获取流水线的关联的最新任务
//
// 获取流水线的关联的最新任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListLatestPipelineJobs(request *model.ListLatestPipelineJobsRequest) (*model.ListLatestPipelineJobsResponse, error) {
	requestDef := GenReqDefForListLatestPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatestPipelineJobsResponse), nil
	}
}

// ListLatestPipelineJobsInvoker 获取流水线的关联的最新任务
func (c *CodeArtsRepoClient) ListLatestPipelineJobsInvoker(request *model.ListLatestPipelineJobsRequest) *ListLatestPipelineJobsInvoker {
	requestDef := GenReqDefForListLatestPipelineJobs()
	return &ListLatestPipelineJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipelineJobs 获取流水线的关联的任务列表
//
// 获取流水线的关联的任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListPipelineJobs(request *model.ListPipelineJobsRequest) (*model.ListPipelineJobsResponse, error) {
	requestDef := GenReqDefForListPipelineJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelineJobsResponse), nil
	}
}

// ListPipelineJobsInvoker 获取流水线的关联的任务列表
func (c *CodeArtsRepoClient) ListPipelineJobsInvoker(request *model.ListPipelineJobsRequest) *ListPipelineJobsInvoker {
	requestDef := GenReqDefForListPipelineJobs()
	return &ListPipelineJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListItemCommits 工作项关联的提交信息
//
// 工作项关联的提交信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListItemCommits(request *model.ListItemCommitsRequest) (*model.ListItemCommitsResponse, error) {
	requestDef := GenReqDefForListItemCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListItemCommitsResponse), nil
	}
}

// ListItemCommitsInvoker 工作项关联的提交信息
func (c *CodeArtsRepoClient) ListItemCommitsInvoker(request *model.ListItemCommitsRequest) *ListItemCommitsInvoker {
	requestDef := GenReqDefForListItemCommits()
	return &ListItemCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectSubgroupsAndRepositories 获取项目下的代码组和仓库列表
//
// 获取项目下的代码组和仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectSubgroupsAndRepositories(request *model.ListProjectSubgroupsAndRepositoriesRequest) (*model.ListProjectSubgroupsAndRepositoriesResponse, error) {
	requestDef := GenReqDefForListProjectSubgroupsAndRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectSubgroupsAndRepositoriesResponse), nil
	}
}

// ListProjectSubgroupsAndRepositoriesInvoker 获取项目下的代码组和仓库列表
func (c *CodeArtsRepoClient) ListProjectSubgroupsAndRepositoriesInvoker(request *model.ListProjectSubgroupsAndRepositoriesRequest) *ListProjectSubgroupsAndRepositoriesInvoker {
	requestDef := GenReqDefForListProjectSubgroupsAndRepositories()
	return &ListProjectSubgroupsAndRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectGeneralPolicy 获取指定项目的基本设置信息
//
// 获取指定项目的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectGeneralPolicy(request *model.ShowProjectGeneralPolicyRequest) (*model.ShowProjectGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowProjectGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectGeneralPolicyResponse), nil
	}
}

// ShowProjectGeneralPolicyInvoker 获取指定项目的基本设置信息
func (c *CodeArtsRepoClient) ShowProjectGeneralPolicyInvoker(request *model.ShowProjectGeneralPolicyRequest) *ShowProjectGeneralPolicyInvoker {
	requestDef := GenReqDefForShowProjectGeneralPolicy()
	return &ShowProjectGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectMemberSetting 获取项目成员设置
//
// 获取项目成员设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectMemberSetting(request *model.ShowProjectMemberSettingRequest) (*model.ShowProjectMemberSettingResponse, error) {
	requestDef := GenReqDefForShowProjectMemberSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectMemberSettingResponse), nil
	}
}

// ShowProjectMemberSettingInvoker 获取项目成员设置
func (c *CodeArtsRepoClient) ShowProjectMemberSettingInvoker(request *model.ShowProjectMemberSettingRequest) *ShowProjectMemberSettingInvoker {
	requestDef := GenReqDefForShowProjectMemberSetting()
	return &ShowProjectMemberSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectSettingsInheritCfg 获取项目继承设置项
//
// 获取项目继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectSettingsInheritCfg(request *model.ShowProjectSettingsInheritCfgRequest) (*model.ShowProjectSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForShowProjectSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectSettingsInheritCfgResponse), nil
	}
}

// ShowProjectSettingsInheritCfgInvoker 获取项目继承设置项
func (c *CodeArtsRepoClient) ShowProjectSettingsInheritCfgInvoker(request *model.ShowProjectSettingsInheritCfgRequest) *ShowProjectSettingsInheritCfgInvoker {
	requestDef := GenReqDefForShowProjectSettingsInheritCfg()
	return &ShowProjectSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWatermark 获取项目水印设置
//
// 获取项目水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectWatermark(request *model.ShowProjectWatermarkRequest) (*model.ShowProjectWatermarkResponse, error) {
	requestDef := GenReqDefForShowProjectWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWatermarkResponse), nil
	}
}

// ShowProjectWatermarkInvoker 获取项目水印设置
func (c *CodeArtsRepoClient) ShowProjectWatermarkInvoker(request *model.ShowProjectWatermarkRequest) *ShowProjectWatermarkInvoker {
	requestDef := GenReqDefForShowProjectWatermark()
	return &ShowProjectWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectsGeneralPolicy 获取指定项目的基本设置信息
//
// 获取指定项目的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectsGeneralPolicy(request *model.ShowProjectsGeneralPolicyRequest) (*model.ShowProjectsGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowProjectsGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectsGeneralPolicyResponse), nil
	}
}

// ShowProjectsGeneralPolicyInvoker 获取指定项目的基本设置信息
func (c *CodeArtsRepoClient) ShowProjectsGeneralPolicyInvoker(request *model.ShowProjectsGeneralPolicyRequest) *ShowProjectsGeneralPolicyInvoker {
	requestDef := GenReqDefForShowProjectsGeneralPolicy()
	return &ShowProjectsGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourcePermissions 获取资源点对应的角色和权限
//
// 获取资源点对应的角色和权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowResourcePermissions(request *model.ShowResourcePermissionsRequest) (*model.ShowResourcePermissionsResponse, error) {
	requestDef := GenReqDefForShowResourcePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourcePermissionsResponse), nil
	}
}

// ShowResourcePermissionsInvoker 获取资源点对应的角色和权限
func (c *CodeArtsRepoClient) ShowResourcePermissionsInvoker(request *model.ShowResourcePermissionsRequest) *ShowResourcePermissionsInvoker {
	requestDef := GenReqDefForShowResourcePermissions()
	return &ShowResourcePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectGeneralPolicy 更新项目的基本设置信息
//
// 更新项目的基本设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectGeneralPolicy(request *model.UpdateProjectGeneralPolicyRequest) (*model.UpdateProjectGeneralPolicyResponse, error) {
	requestDef := GenReqDefForUpdateProjectGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectGeneralPolicyResponse), nil
	}
}

// UpdateProjectGeneralPolicyInvoker 更新项目的基本设置信息
func (c *CodeArtsRepoClient) UpdateProjectGeneralPolicyInvoker(request *model.UpdateProjectGeneralPolicyRequest) *UpdateProjectGeneralPolicyInvoker {
	requestDef := GenReqDefForUpdateProjectGeneralPolicy()
	return &UpdateProjectGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectSettingsInheritCfg 更新项目继承设置项
//
// 更新项目继承设置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectSettingsInheritCfg(request *model.UpdateProjectSettingsInheritCfgRequest) (*model.UpdateProjectSettingsInheritCfgResponse, error) {
	requestDef := GenReqDefForUpdateProjectSettingsInheritCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectSettingsInheritCfgResponse), nil
	}
}

// UpdateProjectSettingsInheritCfgInvoker 更新项目继承设置项
func (c *CodeArtsRepoClient) UpdateProjectSettingsInheritCfgInvoker(request *model.UpdateProjectSettingsInheritCfgRequest) *UpdateProjectSettingsInheritCfgInvoker {
	requestDef := GenReqDefForUpdateProjectSettingsInheritCfg()
	return &UpdateProjectSettingsInheritCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectWatermark 更新项目水印设置
//
// 更新项目水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectWatermark(request *model.UpdateProjectWatermarkRequest) (*model.UpdateProjectWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateProjectWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectWatermarkResponse), nil
	}
}

// UpdateProjectWatermarkInvoker 更新项目水印设置
func (c *CodeArtsRepoClient) UpdateProjectWatermarkInvoker(request *model.UpdateProjectWatermarkRequest) *UpdateProjectWatermarkInvoker {
	requestDef := GenReqDefForUpdateProjectWatermark()
	return &UpdateProjectWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateProtectedBranch 批量创建仓库保护分支
//
// 批量创建仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchCreateProtectedBranch(request *model.BatchCreateProtectedBranchRequest) (*model.BatchCreateProtectedBranchResponse, error) {
	requestDef := GenReqDefForBatchCreateProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateProtectedBranchResponse), nil
	}
}

// BatchCreateProtectedBranchInvoker 批量创建仓库保护分支
func (c *CodeArtsRepoClient) BatchCreateProtectedBranchInvoker(request *model.BatchCreateProtectedBranchRequest) *BatchCreateProtectedBranchInvoker {
	requestDef := GenReqDefForBatchCreateProtectedBranch()
	return &BatchCreateProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateProtectedTags 批量创建仓库保护Tag
//
// 批量创建仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchCreateProtectedTags(request *model.BatchCreateProtectedTagsRequest) (*model.BatchCreateProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateProtectedTagsResponse), nil
	}
}

// BatchCreateProtectedTagsInvoker 批量创建仓库保护Tag
func (c *CodeArtsRepoClient) BatchCreateProtectedTagsInvoker(request *model.BatchCreateProtectedTagsRequest) *BatchCreateProtectedTagsInvoker {
	requestDef := GenReqDefForBatchCreateProtectedTags()
	return &BatchCreateProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteProtectedBranches 批量删除仓库保护分支
//
// 批量删除仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchDeleteProtectedBranches(request *model.BatchDeleteProtectedBranchesRequest) (*model.BatchDeleteProtectedBranchesResponse, error) {
	requestDef := GenReqDefForBatchDeleteProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProtectedBranchesResponse), nil
	}
}

// BatchDeleteProtectedBranchesInvoker 批量删除仓库保护分支
func (c *CodeArtsRepoClient) BatchDeleteProtectedBranchesInvoker(request *model.BatchDeleteProtectedBranchesRequest) *BatchDeleteProtectedBranchesInvoker {
	requestDef := GenReqDefForBatchDeleteProtectedBranches()
	return &BatchDeleteProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteProtectedTags 批量删除仓库保护Tag
//
// 批量删除仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchDeleteProtectedTags(request *model.BatchDeleteProtectedTagsRequest) (*model.BatchDeleteProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProtectedTagsResponse), nil
	}
}

// BatchDeleteProtectedTagsInvoker 批量删除仓库保护Tag
func (c *CodeArtsRepoClient) BatchDeleteProtectedTagsInvoker(request *model.BatchDeleteProtectedTagsRequest) *BatchDeleteProtectedTagsInvoker {
	requestDef := GenReqDefForBatchDeleteProtectedTags()
	return &BatchDeleteProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateProtectedBranches 批量更新仓库保护分支
//
// 批量更新仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchUpdateProtectedBranches(request *model.BatchUpdateProtectedBranchesRequest) (*model.BatchUpdateProtectedBranchesResponse, error) {
	requestDef := GenReqDefForBatchUpdateProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateProtectedBranchesResponse), nil
	}
}

// BatchUpdateProtectedBranchesInvoker 批量更新仓库保护分支
func (c *CodeArtsRepoClient) BatchUpdateProtectedBranchesInvoker(request *model.BatchUpdateProtectedBranchesRequest) *BatchUpdateProtectedBranchesInvoker {
	requestDef := GenReqDefForBatchUpdateProtectedBranches()
	return &BatchUpdateProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateProtectedTags 批量更新仓库保护Tag
//
// 批量更新仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchUpdateProtectedTags(request *model.BatchUpdateProtectedTagsRequest) (*model.BatchUpdateProtectedTagsResponse, error) {
	requestDef := GenReqDefForBatchUpdateProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateProtectedTagsResponse), nil
	}
}

// BatchUpdateProtectedTagsInvoker 批量更新仓库保护Tag
func (c *CodeArtsRepoClient) BatchUpdateProtectedTagsInvoker(request *model.BatchUpdateProtectedTagsRequest) *BatchUpdateProtectedTagsInvoker {
	requestDef := GenReqDefForBatchUpdateProtectedTags()
	return &BatchUpdateProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectProtectedBranches 创建项目下保护分支
//
// 创建项目下保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateProjectProtectedBranches(request *model.CreateProjectProtectedBranchesRequest) (*model.CreateProjectProtectedBranchesResponse, error) {
	requestDef := GenReqDefForCreateProjectProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectProtectedBranchesResponse), nil
	}
}

// CreateProjectProtectedBranchesInvoker 创建项目下保护分支
func (c *CodeArtsRepoClient) CreateProjectProtectedBranchesInvoker(request *model.CreateProjectProtectedBranchesRequest) *CreateProjectProtectedBranchesInvoker {
	requestDef := GenReqDefForCreateProjectProtectedBranches()
	return &CreateProjectProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectProtectedTags 创建项目下的保护tag
//
// 创建项目下的保护tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateProjectProtectedTags(request *model.CreateProjectProtectedTagsRequest) (*model.CreateProjectProtectedTagsResponse, error) {
	requestDef := GenReqDefForCreateProjectProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectProtectedTagsResponse), nil
	}
}

// CreateProjectProtectedTagsInvoker 创建项目下的保护tag
func (c *CodeArtsRepoClient) CreateProjectProtectedTagsInvoker(request *model.CreateProjectProtectedTagsRequest) *CreateProjectProtectedTagsInvoker {
	requestDef := GenReqDefForCreateProjectProtectedTags()
	return &CreateProjectProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProtectedBranch 删除仓库保护分支
//
// 删除仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteProtectedBranch(request *model.DeleteProtectedBranchRequest) (*model.DeleteProtectedBranchResponse, error) {
	requestDef := GenReqDefForDeleteProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedBranchResponse), nil
	}
}

// DeleteProtectedBranchInvoker 删除仓库保护分支
func (c *CodeArtsRepoClient) DeleteProtectedBranchInvoker(request *model.DeleteProtectedBranchRequest) *DeleteProtectedBranchInvoker {
	requestDef := GenReqDefForDeleteProtectedBranch()
	return &DeleteProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProtectedTag 删除仓库保护Tag
//
// 删除仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteProtectedTag(request *model.DeleteProtectedTagRequest) (*model.DeleteProtectedTagResponse, error) {
	requestDef := GenReqDefForDeleteProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedTagResponse), nil
	}
}

// DeleteProtectedTagInvoker 删除仓库保护Tag
func (c *CodeArtsRepoClient) DeleteProtectedTagInvoker(request *model.DeleteProtectedTagRequest) *DeleteProtectedTagInvoker {
	requestDef := GenReqDefForDeleteProtectedTag()
	return &DeleteProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectProtectedBranches 获取项目下保护分支列表
//
// 获取项目下保护分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectProtectedBranches(request *model.ListProjectProtectedBranchesRequest) (*model.ListProjectProtectedBranchesResponse, error) {
	requestDef := GenReqDefForListProjectProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectProtectedBranchesResponse), nil
	}
}

// ListProjectProtectedBranchesInvoker 获取项目下保护分支列表
func (c *CodeArtsRepoClient) ListProjectProtectedBranchesInvoker(request *model.ListProjectProtectedBranchesRequest) *ListProjectProtectedBranchesInvoker {
	requestDef := GenReqDefForListProjectProtectedBranches()
	return &ListProjectProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectProtectedTags 获取指定项目的保护tag详情
//
// 获取指定项目的保护tag详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectProtectedTags(request *model.ListProjectProtectedTagsRequest) (*model.ListProjectProtectedTagsResponse, error) {
	requestDef := GenReqDefForListProjectProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectProtectedTagsResponse), nil
	}
}

// ListProjectProtectedTagsInvoker 获取指定项目的保护tag详情
func (c *CodeArtsRepoClient) ListProjectProtectedTagsInvoker(request *model.ListProjectProtectedTagsRequest) *ListProjectProtectedTagsInvoker {
	requestDef := GenReqDefForListProjectProtectedTags()
	return &ListProjectProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedBranches 获取仓库保护分支列表
//
// 获取仓库保护分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProtectedBranches(request *model.ListProtectedBranchesRequest) (*model.ListProtectedBranchesResponse, error) {
	requestDef := GenReqDefForListProtectedBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedBranchesResponse), nil
	}
}

// ListProtectedBranchesInvoker 获取仓库保护分支列表
func (c *CodeArtsRepoClient) ListProtectedBranchesInvoker(request *model.ListProtectedBranchesRequest) *ListProtectedBranchesInvoker {
	requestDef := GenReqDefForListProtectedBranches()
	return &ListProtectedBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedTags 获取仓库保护Tag列表
//
// 获取仓库保护Tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProtectedTags(request *model.ListProtectedTagsRequest) (*model.ListProtectedTagsResponse, error) {
	requestDef := GenReqDefForListProtectedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedTagsResponse), nil
	}
}

// ListProtectedTagsInvoker 获取仓库保护Tag列表
func (c *CodeArtsRepoClient) ListProtectedTagsInvoker(request *model.ListProtectedTagsRequest) *ListProtectedTagsInvoker {
	requestDef := GenReqDefForListProtectedTags()
	return &ListProtectedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProtectedBranch 获取仓库保护分支
//
// 获取仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProtectedBranch(request *model.ShowProtectedBranchRequest) (*model.ShowProtectedBranchResponse, error) {
	requestDef := GenReqDefForShowProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectedBranchResponse), nil
	}
}

// ShowProtectedBranchInvoker 获取仓库保护分支
func (c *CodeArtsRepoClient) ShowProtectedBranchInvoker(request *model.ShowProtectedBranchRequest) *ShowProtectedBranchInvoker {
	requestDef := GenReqDefForShowProtectedBranch()
	return &ShowProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProtectedTag 获取仓库保护Tag
//
// 获取仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProtectedTag(request *model.ShowProtectedTagRequest) (*model.ShowProtectedTagResponse, error) {
	requestDef := GenReqDefForShowProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectedTagResponse), nil
	}
}

// ShowProtectedTagInvoker 获取仓库保护Tag
func (c *CodeArtsRepoClient) ShowProtectedTagInvoker(request *model.ShowProtectedTagRequest) *ShowProtectedTagInvoker {
	requestDef := GenReqDefForShowProtectedTag()
	return &ShowProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProtectedBranch 更新仓库保护分支
//
// 更新仓库保护分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProtectedBranch(request *model.UpdateProtectedBranchRequest) (*model.UpdateProtectedBranchResponse, error) {
	requestDef := GenReqDefForUpdateProtectedBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectedBranchResponse), nil
	}
}

// UpdateProtectedBranchInvoker 更新仓库保护分支
func (c *CodeArtsRepoClient) UpdateProtectedBranchInvoker(request *model.UpdateProtectedBranchRequest) *UpdateProtectedBranchInvoker {
	requestDef := GenReqDefForUpdateProtectedBranch()
	return &UpdateProtectedBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProtectedTag 更新仓库保护Tag
//
// 更新仓库保护Tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProtectedTag(request *model.UpdateProtectedTagRequest) (*model.UpdateProtectedTagResponse, error) {
	requestDef := GenReqDefForUpdateProtectedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectedTagResponse), nil
	}
}

// UpdateProtectedTagInvoker 更新仓库保护Tag
func (c *CodeArtsRepoClient) UpdateProtectedTagInvoker(request *model.UpdateProtectedTagRequest) *UpdateProtectedTagInvoker {
	requestDef := GenReqDefForUpdateProtectedTag()
	return &UpdateProtectedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteBranch 批量删除分支
//
// 批量删除分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchDeleteBranch(request *model.BatchDeleteBranchRequest) (*model.BatchDeleteBranchResponse, error) {
	requestDef := GenReqDefForBatchDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteBranchResponse), nil
	}
}

// BatchDeleteBranchInvoker 批量删除分支
func (c *CodeArtsRepoClient) BatchDeleteBranchInvoker(request *model.BatchDeleteBranchRequest) *BatchDeleteBranchInvoker {
	requestDef := GenReqDefForBatchDeleteBranch()
	return &BatchDeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBranch 创建分支
//
// 创建分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateBranch(request *model.CreateBranchRequest) (*model.CreateBranchResponse, error) {
	requestDef := GenReqDefForCreateBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBranchResponse), nil
	}
}

// CreateBranchInvoker 创建分支
func (c *CodeArtsRepoClient) CreateBranchInvoker(request *model.CreateBranchRequest) *CreateBranchInvoker {
	requestDef := GenReqDefForCreateBranch()
	return &CreateBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 创建标签
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 创建标签
func (c *CodeArtsRepoClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBranch 删除分支
//
// 删除分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteBranch(request *model.DeleteBranchRequest) (*model.DeleteBranchResponse, error) {
	requestDef := GenReqDefForDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBranchResponse), nil
	}
}

// DeleteBranchInvoker 删除分支
func (c *CodeArtsRepoClient) DeleteBranchInvoker(request *model.DeleteBranchRequest) *DeleteBranchInvoker {
	requestDef := GenReqDefForDeleteBranch()
	return &DeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除标签
func (c *CodeArtsRepoClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranches 获取分支列表
//
// 获取分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListBranches(request *model.ListBranchesRequest) (*model.ListBranchesResponse, error) {
	requestDef := GenReqDefForListBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchesResponse), nil
	}
}

// ListBranchesInvoker 获取分支列表
func (c *CodeArtsRepoClient) ListBranchesInvoker(request *model.ListBranchesRequest) *ListBranchesInvoker {
	requestDef := GenReqDefForListBranches()
	return &ListBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRefsList 查看分支/tag列表
//
// 查看分支/tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRefsList(request *model.ListRefsListRequest) (*model.ListRefsListResponse, error) {
	requestDef := GenReqDefForListRefsList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRefsListResponse), nil
	}
}

// ListRefsListInvoker 查看分支/tag列表
func (c *CodeArtsRepoClient) ListRefsListInvoker(request *model.ListRefsListRequest) *ListRefsListInvoker {
	requestDef := GenReqDefForListRefsList()
	return &ListRefsListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 获取标签列表
//
// 获取标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 获取标签列表
func (c *CodeArtsRepoClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBranch 获取分支详情
//
// 获取分支详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowBranch(request *model.ShowBranchRequest) (*model.ShowBranchResponse, error) {
	requestDef := GenReqDefForShowBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchResponse), nil
	}
}

// ShowBranchInvoker 获取分支详情
func (c *CodeArtsRepoClient) ShowBranchInvoker(request *model.ShowBranchRequest) *ShowBranchInvoker {
	requestDef := GenReqDefForShowBranch()
	return &ShowBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTag 查看标签详情
//
// 查看标签详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowTag(request *model.ShowTagRequest) (*model.ShowTagResponse, error) {
	requestDef := GenReqDefForShowTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagResponse), nil
	}
}

// ShowTagInvoker 查看标签详情
func (c *CodeArtsRepoClient) ShowTagInvoker(request *model.ShowTagRequest) *ShowTagInvoker {
	requestDef := GenReqDefForShowTag()
	return &ShowTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBranchName 分支重命名
//
// 分支重命名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateBranchName(request *model.UpdateBranchNameRequest) (*model.UpdateBranchNameResponse, error) {
	requestDef := GenReqDefForUpdateBranchName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBranchNameResponse), nil
	}
}

// UpdateBranchNameInvoker 分支重命名
func (c *CodeArtsRepoClient) UpdateBranchNameInvoker(request *model.UpdateBranchNameRequest) *UpdateBranchNameInvoker {
	requestDef := GenReqDefForUpdateBranchName()
	return &UpdateBranchNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSubmodule 创建子模块
//
// 创建子模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddSubmodule(request *model.AddSubmoduleRequest) (*model.AddSubmoduleResponse, error) {
	requestDef := GenReqDefForAddSubmodule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubmoduleResponse), nil
	}
}

// AddSubmoduleInvoker 创建子模块
func (c *CodeArtsRepoClient) AddSubmoduleInvoker(request *model.AddSubmoduleRequest) *AddSubmoduleInvoker {
	requestDef := GenReqDefForAddSubmodule()
	return &AddSubmoduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTrustedIpAddress 添加仓库ip白名单
//
// 添加仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddTrustedIpAddress(request *model.AddTrustedIpAddressRequest) (*model.AddTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTrustedIpAddressResponse), nil
	}
}

// AddTrustedIpAddressInvoker 添加仓库ip白名单
func (c *CodeArtsRepoClient) AddTrustedIpAddressInvoker(request *model.AddTrustedIpAddressRequest) *AddTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTrustedIpAddress()
	return &AddTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRemoteMirror 将普通仓库与远程镜像关联
//
// 将普通仓库与远程镜像关联
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AssociateRemoteMirror(request *model.AssociateRemoteMirrorRequest) (*model.AssociateRemoteMirrorResponse, error) {
	requestDef := GenReqDefForAssociateRemoteMirror()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRemoteMirrorResponse), nil
	}
}

// AssociateRemoteMirrorInvoker 将普通仓库与远程镜像关联
func (c *CodeArtsRepoClient) AssociateRemoteMirrorInvoker(request *model.AssociateRemoteMirrorRequest) *AssociateRemoteMirrorInvoker {
	requestDef := GenReqDefForAssociateRemoteMirror()
	return &AssociateRemoteMirrorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRepositoryUserGroup 关联仓库与成员组
//
// 关联仓库与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AssociateRepositoryUserGroup(request *model.AssociateRepositoryUserGroupRequest) (*model.AssociateRepositoryUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRepositoryUserGroupResponse), nil
	}
}

// AssociateRepositoryUserGroupInvoker 关联仓库与成员组
func (c *CodeArtsRepoClient) AssociateRepositoryUserGroupInvoker(request *model.AssociateRepositoryUserGroupRequest) *AssociateRepositoryUserGroupInvoker {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()
	return &AssociateRepositoryUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchValidateRepoNames 批量检查仓库名
//
// 批量检查仓库名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchValidateRepoNames(request *model.BatchValidateRepoNamesRequest) (*model.BatchValidateRepoNamesResponse, error) {
	requestDef := GenReqDefForBatchValidateRepoNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateRepoNamesResponse), nil
	}
}

// BatchValidateRepoNamesInvoker 批量检查仓库名
func (c *CodeArtsRepoClient) BatchValidateRepoNamesInvoker(request *model.BatchValidateRepoNamesRequest) *BatchValidateRepoNamesInvoker {
	requestDef := GenReqDefForBatchValidateRepoNames()
	return &BatchValidateRepoNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDir 创建指定分支下的目录
//
// 创建指定分支下的目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateDir(request *model.CreateDirRequest) (*model.CreateDirResponse, error) {
	requestDef := GenReqDefForCreateDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDirResponse), nil
	}
}

// CreateDirInvoker 创建指定分支下的目录
func (c *CodeArtsRepoClient) CreateDirInvoker(request *model.CreateDirRequest) *CreateDirInvoker {
	requestDef := GenReqDefForCreateDir()
	return &CreateDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepositoryCommitRule 创建仓库提交规则
//
// 创建仓库提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateRepositoryCommitRule(request *model.CreateRepositoryCommitRuleRequest) (*model.CreateRepositoryCommitRuleResponse, error) {
	requestDef := GenReqDefForCreateRepositoryCommitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositoryCommitRuleResponse), nil
	}
}

// CreateRepositoryCommitRuleInvoker 创建仓库提交规则
func (c *CodeArtsRepoClient) CreateRepositoryCommitRuleInvoker(request *model.CreateRepositoryCommitRuleRequest) *CreateRepositoryCommitRuleInvoker {
	requestDef := GenReqDefForCreateRepositoryCommitRule()
	return &CreateRepositoryCommitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepositoryLabel 创建仓库标签
//
// 创建仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateRepositoryLabel(request *model.CreateRepositoryLabelRequest) (*model.CreateRepositoryLabelResponse, error) {
	requestDef := GenReqDefForCreateRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositoryLabelResponse), nil
	}
}

// CreateRepositoryLabelInvoker 创建仓库标签
func (c *CodeArtsRepoClient) CreateRepositoryLabelInvoker(request *model.CreateRepositoryLabelRequest) *CreateRepositoryLabelInvoker {
	requestDef := GenReqDefForCreateRepositoryLabel()
	return &CreateRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepositorySystemLabels 创建仓库系统标签
//
// 创建仓库系统标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CreateRepositorySystemLabels(request *model.CreateRepositorySystemLabelsRequest) (*model.CreateRepositorySystemLabelsResponse, error) {
	requestDef := GenReqDefForCreateRepositorySystemLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepositorySystemLabelsResponse), nil
	}
}

// CreateRepositorySystemLabelsInvoker 创建仓库系统标签
func (c *CodeArtsRepoClient) CreateRepositorySystemLabelsInvoker(request *model.CreateRepositorySystemLabelsRequest) *CreateRepositorySystemLabelsInvoker {
	requestDef := GenReqDefForCreateRepositorySystemLabels()
	return &CreateRepositorySystemLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepositoryLabel 删除仓库标签
//
// 删除仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteRepositoryLabel(request *model.DeleteRepositoryLabelRequest) (*model.DeleteRepositoryLabelResponse, error) {
	requestDef := GenReqDefForDeleteRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepositoryLabelResponse), nil
	}
}

// DeleteRepositoryLabelInvoker 删除仓库标签
func (c *CodeArtsRepoClient) DeleteRepositoryLabelInvoker(request *model.DeleteRepositoryLabelRequest) *DeleteRepositoryLabelInvoker {
	requestDef := GenReqDefForDeleteRepositoryLabel()
	return &DeleteRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrustedIpAddress 删除仓库ip白名单
//
// 删除仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteTrustedIpAddress(request *model.DeleteTrustedIpAddressRequest) (*model.DeleteTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrustedIpAddressResponse), nil
	}
}

// DeleteTrustedIpAddressInvoker 删除仓库ip白名单
func (c *CodeArtsRepoClient) DeleteTrustedIpAddressInvoker(request *model.DeleteTrustedIpAddressRequest) *DeleteTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTrustedIpAddress()
	return &DeleteTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadArchive 仓库下载
//
// 仓库下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DownloadArchive(request *model.DownloadArchiveRequest) (*model.DownloadArchiveResponse, error) {
	requestDef := GenReqDefForDownloadArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadArchiveResponse), nil
	}
}

// DownloadArchiveInvoker 仓库下载
func (c *CodeArtsRepoClient) DownloadArchiveInvoker(request *model.DownloadArchiveRequest) *DownloadArchiveInvoker {
	requestDef := GenReqDefForDownloadArchive()
	return &DownloadArchiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteRepositoryStatistics 触发仓库统计任务
//
// 触发仓库统计任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ExecuteRepositoryStatistics(request *model.ExecuteRepositoryStatisticsRequest) (*model.ExecuteRepositoryStatisticsResponse, error) {
	requestDef := GenReqDefForExecuteRepositoryStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteRepositoryStatisticsResponse), nil
	}
}

// ExecuteRepositoryStatisticsInvoker 触发仓库统计任务
func (c *CodeArtsRepoClient) ExecuteRepositoryStatisticsInvoker(request *model.ExecuteRepositoryStatisticsRequest) *ExecuteRepositoryStatisticsInvoker {
	requestDef := GenReqDefForExecuteRepositoryStatistics()
	return &ExecuteRepositoryStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCurrentUserRepositories 获取当前登录用户仓库
//
// 获取当前登录用户仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListCurrentUserRepositories(request *model.ListCurrentUserRepositoriesRequest) (*model.ListCurrentUserRepositoriesResponse, error) {
	requestDef := GenReqDefForListCurrentUserRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCurrentUserRepositoriesResponse), nil
	}
}

// ListCurrentUserRepositoriesInvoker 获取当前登录用户仓库
func (c *CodeArtsRepoClient) ListCurrentUserRepositoriesInvoker(request *model.ListCurrentUserRepositoriesRequest) *ListCurrentUserRepositoriesInvoker {
	requestDef := GenReqDefForListCurrentUserRepositories()
	return &ListCurrentUserRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupRepositories 获取代码组下仓库列表
//
// 获取代码组下仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupRepositories(request *model.ListGroupRepositoriesRequest) (*model.ListGroupRepositoriesResponse, error) {
	requestDef := GenReqDefForListGroupRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupRepositoriesResponse), nil
	}
}

// ListGroupRepositoriesInvoker 获取代码组下仓库列表
func (c *CodeArtsRepoClient) ListGroupRepositoriesInvoker(request *model.ListGroupRepositoriesRequest) *ListGroupRepositoriesInvoker {
	requestDef := GenReqDefForListGroupRepositories()
	return &ListGroupRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPersonalRecentPushEvents 获取当前用户最近提交动态列表
//
// 查询当前最近前N条提交动态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListPersonalRecentPushEvents(request *model.ListPersonalRecentPushEventsRequest) (*model.ListPersonalRecentPushEventsResponse, error) {
	requestDef := GenReqDefForListPersonalRecentPushEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPersonalRecentPushEventsResponse), nil
	}
}

// ListPersonalRecentPushEventsInvoker 获取当前用户最近提交动态列表
func (c *CodeArtsRepoClient) ListPersonalRecentPushEventsInvoker(request *model.ListPersonalRecentPushEventsRequest) *ListPersonalRecentPushEventsInvoker {
	requestDef := GenReqDefForListPersonalRecentPushEvents()
	return &ListPersonalRecentPushEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPersonalRepositoryImportRecords 查看当前用户仓库导入任务列表
//
// 查看当前用户仓库导入任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListPersonalRepositoryImportRecords(request *model.ListPersonalRepositoryImportRecordsRequest) (*model.ListPersonalRepositoryImportRecordsResponse, error) {
	requestDef := GenReqDefForListPersonalRepositoryImportRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPersonalRepositoryImportRecordsResponse), nil
	}
}

// ListPersonalRepositoryImportRecordsInvoker 查看当前用户仓库导入任务列表
func (c *CodeArtsRepoClient) ListPersonalRepositoryImportRecordsInvoker(request *model.ListPersonalRepositoryImportRecordsRequest) *ListPersonalRepositoryImportRecordsInvoker {
	requestDef := GenReqDefForListPersonalRepositoryImportRecords()
	return &ListPersonalRepositoryImportRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectRepositories 获取项目下仓库列表
//
// 获取项目下仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectRepositories(request *model.ListProjectRepositoriesRequest) (*model.ListProjectRepositoriesResponse, error) {
	requestDef := GenReqDefForListProjectRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectRepositoriesResponse), nil
	}
}

// ListProjectRepositoriesInvoker 获取项目下仓库列表
func (c *CodeArtsRepoClient) ListProjectRepositoriesInvoker(request *model.ListProjectRepositoriesRequest) *ListProjectRepositoriesInvoker {
	requestDef := GenReqDefForListProjectRepositories()
	return &ListProjectRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryCommitRules 查看仓库提交规则
//
// 查看仓库提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryCommitRules(request *model.ListRepositoryCommitRulesRequest) (*model.ListRepositoryCommitRulesResponse, error) {
	requestDef := GenReqDefForListRepositoryCommitRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryCommitRulesResponse), nil
	}
}

// ListRepositoryCommitRulesInvoker 查看仓库提交规则
func (c *CodeArtsRepoClient) ListRepositoryCommitRulesInvoker(request *model.ListRepositoryCommitRulesRequest) *ListRepositoryCommitRulesInvoker {
	requestDef := GenReqDefForListRepositoryCommitRules()
	return &ListRepositoryCommitRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryContributors 获取仓库贡献者列表
//
// 获取仓库贡献者列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryContributors(request *model.ListRepositoryContributorsRequest) (*model.ListRepositoryContributorsResponse, error) {
	requestDef := GenReqDefForListRepositoryContributors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryContributorsResponse), nil
	}
}

// ListRepositoryContributorsInvoker 获取仓库贡献者列表
func (c *CodeArtsRepoClient) ListRepositoryContributorsInvoker(request *model.ListRepositoryContributorsRequest) *ListRepositoryContributorsInvoker {
	requestDef := GenReqDefForListRepositoryContributors()
	return &ListRepositoryContributorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryEvents 获取仓库动态
//
// 获取仓库动态（当前仅开放推送动态）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryEvents(request *model.ListRepositoryEventsRequest) (*model.ListRepositoryEventsResponse, error) {
	requestDef := GenReqDefForListRepositoryEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryEventsResponse), nil
	}
}

// ListRepositoryEventsInvoker 获取仓库动态
func (c *CodeArtsRepoClient) ListRepositoryEventsInvoker(request *model.ListRepositoryEventsRequest) *ListRepositoryEventsInvoker {
	requestDef := GenReqDefForListRepositoryEvents()
	return &ListRepositoryEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryForks 获取仓库Fork列表
//
// 获取仓库Fork列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryForks(request *model.ListRepositoryForksRequest) (*model.ListRepositoryForksResponse, error) {
	requestDef := GenReqDefForListRepositoryForks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryForksResponse), nil
	}
}

// ListRepositoryForksInvoker 获取仓库Fork列表
func (c *CodeArtsRepoClient) ListRepositoryForksInvoker(request *model.ListRepositoryForksRequest) *ListRepositoryForksInvoker {
	requestDef := GenReqDefForListRepositoryForks()
	return &ListRepositoryForksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryLabels 获取仓库标签列表
//
// 获取仓库标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryLabels(request *model.ListRepositoryLabelsRequest) (*model.ListRepositoryLabelsResponse, error) {
	requestDef := GenReqDefForListRepositoryLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryLabelsResponse), nil
	}
}

// ListRepositoryLabelsInvoker 获取仓库标签列表
func (c *CodeArtsRepoClient) ListRepositoryLabelsInvoker(request *model.ListRepositoryLabelsRequest) *ListRepositoryLabelsInvoker {
	requestDef := GenReqDefForListRepositoryLabels()
	return &ListRepositoryLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryLanguages 获取仓库默认分支语言统计
//
// 获取仓库默认分支语言统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryLanguages(request *model.ListRepositoryLanguagesRequest) (*model.ListRepositoryLanguagesResponse, error) {
	requestDef := GenReqDefForListRepositoryLanguages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryLanguagesResponse), nil
	}
}

// ListRepositoryLanguagesInvoker 获取仓库默认分支语言统计
func (c *CodeArtsRepoClient) ListRepositoryLanguagesInvoker(request *model.ListRepositoryLanguagesRequest) *ListRepositoryLanguagesInvoker {
	requestDef := GenReqDefForListRepositoryLanguages()
	return &ListRepositoryLanguagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryTemplates 模板仓列表
//
// 模板仓列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryTemplates(request *model.ListRepositoryTemplatesRequest) (*model.ListRepositoryTemplatesResponse, error) {
	requestDef := GenReqDefForListRepositoryTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryTemplatesResponse), nil
	}
}

// ListRepositoryTemplatesInvoker 模板仓列表
func (c *CodeArtsRepoClient) ListRepositoryTemplatesInvoker(request *model.ListRepositoryTemplatesRequest) *ListRepositoryTemplatesInvoker {
	requestDef := GenReqDefForListRepositoryTemplates()
	return &ListRepositoryTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubmodules 获取仓库指定分支或者标签子模块列表
//
// 获取仓库指定分支或者标签子模块列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListSubmodules(request *model.ListSubmodulesRequest) (*model.ListSubmodulesResponse, error) {
	requestDef := GenReqDefForListSubmodules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubmodulesResponse), nil
	}
}

// ListSubmodulesInvoker 获取仓库指定分支或者标签子模块列表
func (c *CodeArtsRepoClient) ListSubmodulesInvoker(request *model.ListSubmodulesRequest) *ListSubmodulesInvoker {
	requestDef := GenReqDefForListSubmodules()
	return &ListSubmodulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrustedIpAddresses 获取仓库ip白名单
//
// 获取仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListTrustedIpAddresses(request *model.ListTrustedIpAddressesRequest) (*model.ListTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrustedIpAddressesResponse), nil
	}
}

// ListTrustedIpAddressesInvoker 获取仓库ip白名单
func (c *CodeArtsRepoClient) ListTrustedIpAddressesInvoker(request *model.ListTrustedIpAddressesRequest) *ListTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTrustedIpAddresses()
	return &ListTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockRepository 锁定仓库
//
// 锁定仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) LockRepository(request *model.LockRepositoryRequest) (*model.LockRepositoryResponse, error) {
	requestDef := GenReqDefForLockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockRepositoryResponse), nil
	}
}

// LockRepositoryInvoker 锁定仓库
func (c *CodeArtsRepoClient) LockRepositoryInvoker(request *model.LockRepositoryRequest) *LockRepositoryInvoker {
	requestDef := GenReqDefForLockRepository()
	return &LockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveDeployKey 删除仓库部署密钥
//
// 删除仓库部署密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RemoveDeployKey(request *model.RemoveDeployKeyRequest) (*model.RemoveDeployKeyResponse, error) {
	requestDef := GenReqDefForRemoveDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveDeployKeyResponse), nil
	}
}

// RemoveDeployKeyInvoker 删除仓库部署密钥
func (c *CodeArtsRepoClient) RemoveDeployKeyInvoker(request *model.RemoveDeployKeyRequest) *RemoveDeployKeyInvoker {
	requestDef := GenReqDefForRemoveDeployKey()
	return &RemoveDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveDeployKeyFromSubmodules 删除子仓库部署密钥
//
// 将该该仓库的部署密钥从子模组中删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RemoveDeployKeyFromSubmodules(request *model.RemoveDeployKeyFromSubmodulesRequest) (*model.RemoveDeployKeyFromSubmodulesResponse, error) {
	requestDef := GenReqDefForRemoveDeployKeyFromSubmodules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveDeployKeyFromSubmodulesResponse), nil
	}
}

// RemoveDeployKeyFromSubmodulesInvoker 删除子仓库部署密钥
func (c *CodeArtsRepoClient) RemoveDeployKeyFromSubmodulesInvoker(request *model.RemoveDeployKeyFromSubmodulesRequest) *RemoveDeployKeyFromSubmodulesInvoker {
	requestDef := GenReqDefForRemoveDeployKeyFromSubmodules()
	return &RemoveDeployKeyFromSubmodulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlobs 获取文件内容
//
// 获取文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowBlobs(request *model.ShowBlobsRequest) (*model.ShowBlobsResponse, error) {
	requestDef := GenReqDefForShowBlobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlobsResponse), nil
	}
}

// ShowBlobsInvoker 获取文件内容
func (c *CodeArtsRepoClient) ShowBlobsInvoker(request *model.ShowBlobsRequest) *ShowBlobsInvoker {
	requestDef := GenReqDefForShowBlobs()
	return &ShowBlobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommitStatistics 获取仓库指定分支的提交统计信息
//
// 获取仓库指定分支的提交统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowCommitStatistics(request *model.ShowCommitStatisticsRequest) (*model.ShowCommitStatisticsResponse, error) {
	requestDef := GenReqDefForShowCommitStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommitStatisticsResponse), nil
	}
}

// ShowCommitStatisticsInvoker 获取仓库指定分支的提交统计信息
func (c *CodeArtsRepoClient) ShowCommitStatisticsInvoker(request *model.ShowCommitStatisticsRequest) *ShowCommitStatisticsInvoker {
	requestDef := GenReqDefForShowCommitStatistics()
	return &ShowCommitStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiffLines 按行数查询提交文件内容
//
// 按行数查询提交文件内容，最大显示行数为1000行
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowDiffLines(request *model.ShowDiffLinesRequest) (*model.ShowDiffLinesResponse, error) {
	requestDef := GenReqDefForShowDiffLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiffLinesResponse), nil
	}
}

// ShowDiffLinesInvoker 按行数查询提交文件内容
func (c *CodeArtsRepoClient) ShowDiffLinesInvoker(request *model.ShowDiffLinesRequest) *ShowDiffLinesInvoker {
	requestDef := GenReqDefForShowDiffLines()
	return &ShowDiffLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLastPushEventInRepository 获取仓库最近推送事件
//
// 获取仓库最近推送事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowLastPushEventInRepository(request *model.ShowLastPushEventInRepositoryRequest) (*model.ShowLastPushEventInRepositoryResponse, error) {
	requestDef := GenReqDefForShowLastPushEventInRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastPushEventInRepositoryResponse), nil
	}
}

// ShowLastPushEventInRepositoryInvoker 获取仓库最近推送事件
func (c *CodeArtsRepoClient) ShowLastPushEventInRepositoryInvoker(request *model.ShowLastPushEventInRepositoryRequest) *ShowLastPushEventInRepositoryInvoker {
	requestDef := GenReqDefForShowLastPushEventInRepository()
	return &ShowLastPushEventInRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotificationSubscription 获取仓库通知设置
//
// 获取仓库通知设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowNotificationSubscription(request *model.ShowNotificationSubscriptionRequest) (*model.ShowNotificationSubscriptionResponse, error) {
	requestDef := GenReqDefForShowNotificationSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationSubscriptionResponse), nil
	}
}

// ShowNotificationSubscriptionInvoker 获取仓库通知设置
func (c *CodeArtsRepoClient) ShowNotificationSubscriptionInvoker(request *model.ShowNotificationSubscriptionRequest) *ShowNotificationSubscriptionInvoker {
	requestDef := GenReqDefForShowNotificationSubscription()
	return &ShowNotificationSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotificationSubscriptionsStatus 获取仓库通知设置启用状态
//
// 获取仓库通知设置启用状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowNotificationSubscriptionsStatus(request *model.ShowNotificationSubscriptionsStatusRequest) (*model.ShowNotificationSubscriptionsStatusResponse, error) {
	requestDef := GenReqDefForShowNotificationSubscriptionsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationSubscriptionsStatusResponse), nil
	}
}

// ShowNotificationSubscriptionsStatusInvoker 获取仓库通知设置启用状态
func (c *CodeArtsRepoClient) ShowNotificationSubscriptionsStatusInvoker(request *model.ShowNotificationSubscriptionsStatusRequest) *ShowNotificationSubscriptionsStatusInvoker {
	requestDef := GenReqDefForShowNotificationSubscriptionsStatus()
	return &ShowNotificationSubscriptionsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRefCompare 分支、tags、提交对比
//
// 分支、tags、提交对比
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRefCompare(request *model.ShowRefCompareRequest) (*model.ShowRefCompareResponse, error) {
	requestDef := GenReqDefForShowRefCompare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefCompareResponse), nil
	}
}

// ShowRefCompareInvoker 分支、tags、提交对比
func (c *CodeArtsRepoClient) ShowRefCompareInvoker(request *model.ShowRefCompareRequest) *ShowRefCompareInvoker {
	requestDef := GenReqDefForShowRefCompare()
	return &ShowRefCompareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRemoteMirror 获取仓库镜像详情
//
// 获取仓库镜像详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRemoteMirror(request *model.ShowRemoteMirrorRequest) (*model.ShowRemoteMirrorResponse, error) {
	requestDef := GenReqDefForShowRemoteMirror()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRemoteMirrorResponse), nil
	}
}

// ShowRemoteMirrorInvoker 获取仓库镜像详情
func (c *CodeArtsRepoClient) ShowRemoteMirrorInvoker(request *model.ShowRemoteMirrorRequest) *ShowRemoteMirrorInvoker {
	requestDef := GenReqDefForShowRemoteMirror()
	return &ShowRemoteMirrorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepository 获取仓库详情
//
// 获取仓库详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepository(request *model.ShowRepositoryRequest) (*model.ShowRepositoryResponse, error) {
	requestDef := GenReqDefForShowRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryResponse), nil
	}
}

// ShowRepositoryInvoker 获取仓库详情
func (c *CodeArtsRepoClient) ShowRepositoryInvoker(request *model.ShowRepositoryRequest) *ShowRepositoryInvoker {
	requestDef := GenReqDefForShowRepository()
	return &ShowRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryGeneralCommitRule 查看仓库通用提交规则
//
// 查看仓库通用提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryGeneralCommitRule(request *model.ShowRepositoryGeneralCommitRuleRequest) (*model.ShowRepositoryGeneralCommitRuleResponse, error) {
	requestDef := GenReqDefForShowRepositoryGeneralCommitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryGeneralCommitRuleResponse), nil
	}
}

// ShowRepositoryGeneralCommitRuleInvoker 查看仓库通用提交规则
func (c *CodeArtsRepoClient) ShowRepositoryGeneralCommitRuleInvoker(request *model.ShowRepositoryGeneralCommitRuleRequest) *ShowRepositoryGeneralCommitRuleInvoker {
	requestDef := GenReqDefForShowRepositoryGeneralCommitRule()
	return &ShowRepositoryGeneralCommitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryGeneralPolicy 查看仓库通用策略
//
// 查看仓库通用策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryGeneralPolicy(request *model.ShowRepositoryGeneralPolicyRequest) (*model.ShowRepositoryGeneralPolicyResponse, error) {
	requestDef := GenReqDefForShowRepositoryGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryGeneralPolicyResponse), nil
	}
}

// ShowRepositoryGeneralPolicyInvoker 查看仓库通用策略
func (c *CodeArtsRepoClient) ShowRepositoryGeneralPolicyInvoker(request *model.ShowRepositoryGeneralPolicyRequest) *ShowRepositoryGeneralPolicyInvoker {
	requestDef := GenReqDefForShowRepositoryGeneralPolicy()
	return &ShowRepositoryGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryInheritSetting 查看仓库继承设置
//
// 查看仓库继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryInheritSetting(request *model.ShowRepositoryInheritSettingRequest) (*model.ShowRepositoryInheritSettingResponse, error) {
	requestDef := GenReqDefForShowRepositoryInheritSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryInheritSettingResponse), nil
	}
}

// ShowRepositoryInheritSettingInvoker 查看仓库继承设置
func (c *CodeArtsRepoClient) ShowRepositoryInheritSettingInvoker(request *model.ShowRepositoryInheritSettingRequest) *ShowRepositoryInheritSettingInvoker {
	requestDef := GenReqDefForShowRepositoryInheritSetting()
	return &ShowRepositoryInheritSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryInheritSettingSource 查看仓库继承设置源
//
// 查看仓库继承设置源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryInheritSettingSource(request *model.ShowRepositoryInheritSettingSourceRequest) (*model.ShowRepositoryInheritSettingSourceResponse, error) {
	requestDef := GenReqDefForShowRepositoryInheritSettingSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryInheritSettingSourceResponse), nil
	}
}

// ShowRepositoryInheritSettingSourceInvoker 查看仓库继承设置源
func (c *CodeArtsRepoClient) ShowRepositoryInheritSettingSourceInvoker(request *model.ShowRepositoryInheritSettingSourceRequest) *ShowRepositoryInheritSettingSourceInvoker {
	requestDef := GenReqDefForShowRepositoryInheritSettingSource()
	return &ShowRepositoryInheritSettingSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatisticsStatus 获取仓库统计任务状态
//
// 获取仓库统计任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryStatisticsStatus(request *model.ShowRepositoryStatisticsStatusRequest) (*model.ShowRepositoryStatisticsStatusResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatisticsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsStatusResponse), nil
	}
}

// ShowRepositoryStatisticsStatusInvoker 获取仓库统计任务状态
func (c *CodeArtsRepoClient) ShowRepositoryStatisticsStatusInvoker(request *model.ShowRepositoryStatisticsStatusRequest) *ShowRepositoryStatisticsStatusInvoker {
	requestDef := GenReqDefForShowRepositoryStatisticsStatus()
	return &ShowRepositoryStatisticsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryStatisticsSummary 获取仓库统计摘要
//
// 获取仓库统计摘要
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryStatisticsSummary(request *model.ShowRepositoryStatisticsSummaryRequest) (*model.ShowRepositoryStatisticsSummaryResponse, error) {
	requestDef := GenReqDefForShowRepositoryStatisticsSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryStatisticsSummaryResponse), nil
	}
}

// ShowRepositoryStatisticsSummaryInvoker 获取仓库统计摘要
func (c *CodeArtsRepoClient) ShowRepositoryStatisticsSummaryInvoker(request *model.ShowRepositoryStatisticsSummaryRequest) *ShowRepositoryStatisticsSummaryInvoker {
	requestDef := GenReqDefForShowRepositoryStatisticsSummary()
	return &ShowRepositoryStatisticsSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWatermark 获取仓库水印设置
//
// 获取仓库水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryWatermark(request *model.ShowRepositoryWatermarkRequest) (*model.ShowRepositoryWatermarkResponse, error) {
	requestDef := GenReqDefForShowRepositoryWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWatermarkResponse), nil
	}
}

// ShowRepositoryWatermarkInvoker 获取仓库水印设置
func (c *CodeArtsRepoClient) ShowRepositoryWatermarkInvoker(request *model.ShowRepositoryWatermarkRequest) *ShowRepositoryWatermarkInvoker {
	requestDef := GenReqDefForShowRepositoryWatermark()
	return &ShowRepositoryWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserRefPermission 获取CR仓库用户分支或标签级权限
//
// 获取CR仓库用户分支或标签级权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowUserRefPermission(request *model.ShowUserRefPermissionRequest) (*model.ShowUserRefPermissionResponse, error) {
	requestDef := GenReqDefForShowUserRefPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserRefPermissionResponse), nil
	}
}

// ShowUserRefPermissionInvoker 获取CR仓库用户分支或标签级权限
func (c *CodeArtsRepoClient) ShowUserRefPermissionInvoker(request *model.ShowUserRefPermissionRequest) *ShowUserRefPermissionInvoker {
	requestDef := GenReqDefForShowUserRefPermission()
	return &ShowUserRefPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartHouseKeeping 启动仓库加速
//
// 启动仓库加速
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) StartHouseKeeping(request *model.StartHouseKeepingRequest) (*model.StartHouseKeepingResponse, error) {
	requestDef := GenReqDefForStartHouseKeeping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartHouseKeepingResponse), nil
	}
}

// StartHouseKeepingInvoker 启动仓库加速
func (c *CodeArtsRepoClient) StartHouseKeepingInvoker(request *model.StartHouseKeepingRequest) *StartHouseKeepingInvoker {
	requestDef := GenReqDefForStartHouseKeeping()
	return &StartHouseKeepingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartRemoteMirrorSynchronization 启动仓库镜像同步
//
// 启动仓库镜像同步
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) StartRemoteMirrorSynchronization(request *model.StartRemoteMirrorSynchronizationRequest) (*model.StartRemoteMirrorSynchronizationResponse, error) {
	requestDef := GenReqDefForStartRemoteMirrorSynchronization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartRemoteMirrorSynchronizationResponse), nil
	}
}

// StartRemoteMirrorSynchronizationInvoker 启动仓库镜像同步
func (c *CodeArtsRepoClient) StartRemoteMirrorSynchronizationInvoker(request *model.StartRemoteMirrorSynchronizationRequest) *StartRemoteMirrorSynchronizationInvoker {
	requestDef := GenReqDefForStartRemoteMirrorSynchronization()
	return &StartRemoteMirrorSynchronizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncDeployKeyToSubmodules 仓库部署密钥同步到子仓
//
// 将该仓库的部署密钥同步到所有的子模组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) SyncDeployKeyToSubmodules(request *model.SyncDeployKeyToSubmodulesRequest) (*model.SyncDeployKeyToSubmodulesResponse, error) {
	requestDef := GenReqDefForSyncDeployKeyToSubmodules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncDeployKeyToSubmodulesResponse), nil
	}
}

// SyncDeployKeyToSubmodulesInvoker 仓库部署密钥同步到子仓
func (c *CodeArtsRepoClient) SyncDeployKeyToSubmodulesInvoker(request *model.SyncDeployKeyToSubmodulesRequest) *SyncDeployKeyToSubmodulesInvoker {
	requestDef := GenReqDefForSyncDeployKeyToSubmodules()
	return &SyncDeployKeyToSubmodulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockRepository 解锁仓库
//
// 解锁仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UnlockRepository(request *model.UnlockRepositoryRequest) (*model.UnlockRepositoryResponse, error) {
	requestDef := GenReqDefForUnlockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockRepositoryResponse), nil
	}
}

// UnlockRepositoryInvoker 解锁仓库
func (c *CodeArtsRepoClient) UnlockRepositoryInvoker(request *model.UnlockRepositoryRequest) *UnlockRepositoryInvoker {
	requestDef := GenReqDefForUnlockRepository()
	return &UnlockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotificationSubscription 修改仓库通知设置
//
// 修改仓库通知设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateNotificationSubscription(request *model.UpdateNotificationSubscriptionRequest) (*model.UpdateNotificationSubscriptionResponse, error) {
	requestDef := GenReqDefForUpdateNotificationSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationSubscriptionResponse), nil
	}
}

// UpdateNotificationSubscriptionInvoker 修改仓库通知设置
func (c *CodeArtsRepoClient) UpdateNotificationSubscriptionInvoker(request *model.UpdateNotificationSubscriptionRequest) *UpdateNotificationSubscriptionInvoker {
	requestDef := GenReqDefForUpdateNotificationSubscription()
	return &UpdateNotificationSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryCommitRule 修改仓库提交规则
//
// 修改仓库提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryCommitRule(request *model.UpdateRepositoryCommitRuleRequest) (*model.UpdateRepositoryCommitRuleResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryCommitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryCommitRuleResponse), nil
	}
}

// UpdateRepositoryCommitRuleInvoker 修改仓库提交规则
func (c *CodeArtsRepoClient) UpdateRepositoryCommitRuleInvoker(request *model.UpdateRepositoryCommitRuleRequest) *UpdateRepositoryCommitRuleInvoker {
	requestDef := GenReqDefForUpdateRepositoryCommitRule()
	return &UpdateRepositoryCommitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryGeneralCommitRule 修改仓库通用提交规则
//
// 修改仓库通用提交规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryGeneralCommitRule(request *model.UpdateRepositoryGeneralCommitRuleRequest) (*model.UpdateRepositoryGeneralCommitRuleResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryGeneralCommitRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryGeneralCommitRuleResponse), nil
	}
}

// UpdateRepositoryGeneralCommitRuleInvoker 修改仓库通用提交规则
func (c *CodeArtsRepoClient) UpdateRepositoryGeneralCommitRuleInvoker(request *model.UpdateRepositoryGeneralCommitRuleRequest) *UpdateRepositoryGeneralCommitRuleInvoker {
	requestDef := GenReqDefForUpdateRepositoryGeneralCommitRule()
	return &UpdateRepositoryGeneralCommitRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryGeneralPolicy 修改仓库通用策略
//
// 修改仓库通用策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryGeneralPolicy(request *model.UpdateRepositoryGeneralPolicyRequest) (*model.UpdateRepositoryGeneralPolicyResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryGeneralPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryGeneralPolicyResponse), nil
	}
}

// UpdateRepositoryGeneralPolicyInvoker 修改仓库通用策略
func (c *CodeArtsRepoClient) UpdateRepositoryGeneralPolicyInvoker(request *model.UpdateRepositoryGeneralPolicyRequest) *UpdateRepositoryGeneralPolicyInvoker {
	requestDef := GenReqDefForUpdateRepositoryGeneralPolicy()
	return &UpdateRepositoryGeneralPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryInheritSetting 修改仓库继承设置
//
// 修改仓库继承设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryInheritSetting(request *model.UpdateRepositoryInheritSettingRequest) (*model.UpdateRepositoryInheritSettingResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryInheritSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryInheritSettingResponse), nil
	}
}

// UpdateRepositoryInheritSettingInvoker 修改仓库继承设置
func (c *CodeArtsRepoClient) UpdateRepositoryInheritSettingInvoker(request *model.UpdateRepositoryInheritSettingRequest) *UpdateRepositoryInheritSettingInvoker {
	requestDef := GenReqDefForUpdateRepositoryInheritSetting()
	return &UpdateRepositoryInheritSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryLabel 修改仓库标签
//
// 修改仓库标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryLabel(request *model.UpdateRepositoryLabelRequest) (*model.UpdateRepositoryLabelResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryLabelResponse), nil
	}
}

// UpdateRepositoryLabelInvoker 修改仓库标签
func (c *CodeArtsRepoClient) UpdateRepositoryLabelInvoker(request *model.UpdateRepositoryLabelRequest) *UpdateRepositoryLabelInvoker {
	requestDef := GenReqDefForUpdateRepositoryLabel()
	return &UpdateRepositoryLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryRemoteMirror 更新仓库镜像信息
//
// 更新仓库镜像信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryRemoteMirror(request *model.UpdateRepositoryRemoteMirrorRequest) (*model.UpdateRepositoryRemoteMirrorResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryRemoteMirror()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryRemoteMirrorResponse), nil
	}
}

// UpdateRepositoryRemoteMirrorInvoker 更新仓库镜像信息
func (c *CodeArtsRepoClient) UpdateRepositoryRemoteMirrorInvoker(request *model.UpdateRepositoryRemoteMirrorRequest) *UpdateRepositoryRemoteMirrorInvoker {
	requestDef := GenReqDefForUpdateRepositoryRemoteMirror()
	return &UpdateRepositoryRemoteMirrorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryWatermark 更新仓库水印设置
//
// 更新仓库水印设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryWatermark(request *model.UpdateRepositoryWatermarkRequest) (*model.UpdateRepositoryWatermarkResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryWatermarkResponse), nil
	}
}

// UpdateRepositoryWatermarkInvoker 更新仓库水印设置
func (c *CodeArtsRepoClient) UpdateRepositoryWatermarkInvoker(request *model.UpdateRepositoryWatermarkRequest) *UpdateRepositoryWatermarkInvoker {
	requestDef := GenReqDefForUpdateRepositoryWatermark()
	return &UpdateRepositoryWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrustedIpAddress 修改仓库ip白名单
//
// 修改仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateTrustedIpAddress(request *model.UpdateTrustedIpAddressRequest) (*model.UpdateTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrustedIpAddressResponse), nil
	}
}

// UpdateTrustedIpAddressInvoker 修改仓库ip白名单
func (c *CodeArtsRepoClient) UpdateTrustedIpAddressInvoker(request *model.UpdateTrustedIpAddressRequest) *UpdateTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTrustedIpAddress()
	return &UpdateTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTenantTrustedIpAddress 添加租户ip白名单
//
// 添加租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddTenantTrustedIpAddress(request *model.AddTenantTrustedIpAddressRequest) (*model.AddTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTenantTrustedIpAddressResponse), nil
	}
}

// AddTenantTrustedIpAddressInvoker 添加租户ip白名单
func (c *CodeArtsRepoClient) AddTenantTrustedIpAddressInvoker(request *model.AddTenantTrustedIpAddressRequest) *AddTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()
	return &AddTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTenantTrustedIpAddress 删除租户ip白名单
//
// 删除租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteTenantTrustedIpAddress(request *model.DeleteTenantTrustedIpAddressRequest) (*model.DeleteTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTenantTrustedIpAddressResponse), nil
	}
}

// DeleteTenantTrustedIpAddressInvoker 删除租户ip白名单
func (c *CodeArtsRepoClient) DeleteTenantTrustedIpAddressInvoker(request *model.DeleteTenantTrustedIpAddressRequest) *DeleteTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()
	return &DeleteTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportTenantRepositories 租户仓库列表
//
// 租户下所有占用资源的仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ExportTenantRepositories(request *model.ExportTenantRepositoriesRequest) (*model.ExportTenantRepositoriesResponse, error) {
	requestDef := GenReqDefForExportTenantRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportTenantRepositoriesResponse), nil
	}
}

// ExportTenantRepositoriesInvoker 租户仓库列表
func (c *CodeArtsRepoClient) ExportTenantRepositoriesInvoker(request *model.ExportTenantRepositoriesRequest) *ExportTenantRepositoriesInvoker {
	requestDef := GenReqDefForExportTenantRepositories()
	return &ExportTenantRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantRepositories 租户仓库列表
//
// 租户下所有占用资源的仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListTenantRepositories(request *model.ListTenantRepositoriesRequest) (*model.ListTenantRepositoriesResponse, error) {
	requestDef := GenReqDefForListTenantRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantRepositoriesResponse), nil
	}
}

// ListTenantRepositoriesInvoker 租户仓库列表
func (c *CodeArtsRepoClient) ListTenantRepositoriesInvoker(request *model.ListTenantRepositoriesRequest) *ListTenantRepositoriesInvoker {
	requestDef := GenReqDefForListTenantRepositories()
	return &ListTenantRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantTrustedIpAddresses 获取租户ip白名单
//
// 获取租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListTenantTrustedIpAddresses(request *model.ListTenantTrustedIpAddressesRequest) (*model.ListTenantTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantTrustedIpAddressesResponse), nil
	}
}

// ListTenantTrustedIpAddressesInvoker 获取租户ip白名单
func (c *CodeArtsRepoClient) ListTenantTrustedIpAddressesInvoker(request *model.ListTenantTrustedIpAddressesRequest) *ListTenantTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()
	return &ListTenantTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantTrustedIpAddress 修改租户ip白名单
//
// 修改租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateTenantTrustedIpAddress(request *model.UpdateTenantTrustedIpAddressRequest) (*model.UpdateTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantTrustedIpAddressResponse), nil
	}
}

// UpdateTenantTrustedIpAddressInvoker 修改租户ip白名单
func (c *CodeArtsRepoClient) UpdateTenantTrustedIpAddressInvoker(request *model.UpdateTenantTrustedIpAddressRequest) *UpdateTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()
	return &UpdateTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDeployKey 校验部署密钥在上层代码组或项目是否配置
//
// 校验部署密钥在上层代码组或项目是否配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CheckDeployKey(request *model.CheckDeployKeyRequest) (*model.CheckDeployKeyResponse, error) {
	requestDef := GenReqDefForCheckDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDeployKeyResponse), nil
	}
}

// CheckDeployKeyInvoker 校验部署密钥在上层代码组或项目是否配置
func (c *CodeArtsRepoClient) CheckDeployKeyInvoker(request *model.CheckDeployKeyRequest) *CheckDeployKeyInvoker {
	requestDef := GenReqDefForCheckDeployKey()
	return &CheckDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckGroupDeployKey 校验代码组部署密钥在上层代码组或项目是否配置
//
// 校验代码组部署密钥在上层代码组或项目是否配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) CheckGroupDeployKey(request *model.CheckGroupDeployKeyRequest) (*model.CheckGroupDeployKeyResponse, error) {
	requestDef := GenReqDefForCheckGroupDeployKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckGroupDeployKeyResponse), nil
	}
}

// CheckGroupDeployKeyInvoker 校验代码组部署密钥在上层代码组或项目是否配置
func (c *CodeArtsRepoClient) CheckGroupDeployKeyInvoker(request *model.CheckGroupDeployKeyRequest) *CheckGroupDeployKeyInvoker {
	requestDef := GenReqDefForCheckGroupDeployKey()
	return &CheckGroupDeployKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranchRelatedWorkItems 获取仓库下指定分支的关联工作项列表
//
// 获取仓库下指定分支的关联工作项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListBranchRelatedWorkItems(request *model.ListBranchRelatedWorkItemsRequest) (*model.ListBranchRelatedWorkItemsResponse, error) {
	requestDef := GenReqDefForListBranchRelatedWorkItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchRelatedWorkItemsResponse), nil
	}
}

// ListBranchRelatedWorkItemsInvoker 获取仓库下指定分支的关联工作项列表
func (c *CodeArtsRepoClient) ListBranchRelatedWorkItemsInvoker(request *model.ListBranchRelatedWorkItemsRequest) *ListBranchRelatedWorkItemsInvoker {
	requestDef := GenReqDefForListBranchRelatedWorkItems()
	return &ListBranchRelatedWorkItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupDeployKeys 获取代码组下部署密钥列表
//
// 获取代码组下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupDeployKeys(request *model.ListGroupDeployKeysRequest) (*model.ListGroupDeployKeysResponse, error) {
	requestDef := GenReqDefForListGroupDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupDeployKeysResponse), nil
	}
}

// ListGroupDeployKeysInvoker 获取代码组下部署密钥列表
func (c *CodeArtsRepoClient) ListGroupDeployKeysInvoker(request *model.ListGroupDeployKeysRequest) *ListGroupDeployKeysInvoker {
	requestDef := GenReqDefForListGroupDeployKeys()
	return &ListGroupDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectDeployKeys 获取项目下部署密钥列表
//
// 获取项目下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectDeployKeys(request *model.ListProjectDeployKeysRequest) (*model.ListProjectDeployKeysResponse, error) {
	requestDef := GenReqDefForListProjectDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectDeployKeysResponse), nil
	}
}

// ListProjectDeployKeysInvoker 获取项目下部署密钥列表
func (c *CodeArtsRepoClient) ListProjectDeployKeysInvoker(request *model.ListProjectDeployKeysRequest) *ListProjectDeployKeysInvoker {
	requestDef := GenReqDefForListProjectDeployKeys()
	return &ListProjectDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryDeployKeys 获取仓库下部署密钥列表
//
// 获取仓库下部署密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryDeployKeys(request *model.ListRepositoryDeployKeysRequest) (*model.ListRepositoryDeployKeysResponse, error) {
	requestDef := GenReqDefForListRepositoryDeployKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryDeployKeysResponse), nil
	}
}

// ListRepositoryDeployKeysInvoker 获取仓库下部署密钥列表
func (c *CodeArtsRepoClient) ListRepositoryDeployKeysInvoker(request *model.ListRepositoryDeployKeysRequest) *ListRepositoryDeployKeysInvoker {
	requestDef := GenReqDefForListRepositoryDeployKeys()
	return &ListRepositoryDeployKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWorkItems 获取仓库下工作项列表
//
// 获取仓库下工作项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryWorkItems(request *model.ListRepositoryWorkItemsRequest) (*model.ListRepositoryWorkItemsResponse, error) {
	requestDef := GenReqDefForListRepositoryWorkItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWorkItemsResponse), nil
	}
}

// ListRepositoryWorkItemsInvoker 获取仓库下工作项列表
func (c *CodeArtsRepoClient) ListRepositoryWorkItemsInvoker(request *model.ListRepositoryWorkItemsRequest) *ListRepositoryWorkItemsInvoker {
	requestDef := GenReqDefForListRepositoryWorkItems()
	return &ListRepositoryWorkItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupE2eSetting 获取代码组下E2E设置信息
//
// 获取代码组下E2E设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupE2eSetting(request *model.ShowGroupE2eSettingRequest) (*model.ShowGroupE2eSettingResponse, error) {
	requestDef := GenReqDefForShowGroupE2eSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupE2eSettingResponse), nil
	}
}

// ShowGroupE2eSettingInvoker 获取代码组下E2E设置信息
func (c *CodeArtsRepoClient) ShowGroupE2eSettingInvoker(request *model.ShowGroupE2eSettingRequest) *ShowGroupE2eSettingInvoker {
	requestDef := GenReqDefForShowGroupE2eSetting()
	return &ShowGroupE2eSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectE2eSetting 获取项目下E2E设置信息
//
// 获取项目下E2E设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectE2eSetting(request *model.ShowProjectE2eSettingRequest) (*model.ShowProjectE2eSettingResponse, error) {
	requestDef := GenReqDefForShowProjectE2eSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectE2eSettingResponse), nil
	}
}

// ShowProjectE2eSettingInvoker 获取项目下E2E设置信息
func (c *CodeArtsRepoClient) ShowProjectE2eSettingInvoker(request *model.ShowProjectE2eSettingRequest) *ShowProjectE2eSettingInvoker {
	requestDef := GenReqDefForShowProjectE2eSetting()
	return &ShowProjectE2eSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryE2eSetting 获取仓库下E2E设置信息
//
// 获取仓库下E2E设置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryE2eSetting(request *model.ShowRepositoryE2eSettingRequest) (*model.ShowRepositoryE2eSettingResponse, error) {
	requestDef := GenReqDefForShowRepositoryE2eSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryE2eSettingResponse), nil
	}
}

// ShowRepositoryE2eSettingInvoker 获取仓库下E2E设置信息
func (c *CodeArtsRepoClient) ShowRepositoryE2eSettingInvoker(request *model.ShowRepositoryE2eSettingRequest) *ShowRepositoryE2eSettingInvoker {
	requestDef := GenReqDefForShowRepositoryE2eSetting()
	return &ShowRepositoryE2eSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSshKey 添加ssh公钥
//
// 添加ssh公钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddSshKey(request *model.AddSshKeyRequest) (*model.AddSshKeyResponse, error) {
	requestDef := GenReqDefForAddSshKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSshKeyResponse), nil
	}
}

// AddSshKeyInvoker 添加ssh公钥
func (c *CodeArtsRepoClient) AddSshKeyInvoker(request *model.AddSshKeyRequest) *AddSshKeyInvoker {
	requestDef := GenReqDefForAddSshKey()
	return &AddSshKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchValidateUserGroupPermissions 获取当前用户指定的代码组列表中的权限
//
// 获取当前用户指定的代码组列表中的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) BatchValidateUserGroupPermissions(request *model.BatchValidateUserGroupPermissionsRequest) (*model.BatchValidateUserGroupPermissionsResponse, error) {
	requestDef := GenReqDefForBatchValidateUserGroupPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateUserGroupPermissionsResponse), nil
	}
}

// BatchValidateUserGroupPermissionsInvoker 获取当前用户指定的代码组列表中的权限
func (c *CodeArtsRepoClient) BatchValidateUserGroupPermissionsInvoker(request *model.BatchValidateUserGroupPermissionsRequest) *BatchValidateUserGroupPermissionsInvoker {
	requestDef := GenReqDefForBatchValidateUserGroupPermissions()
	return &BatchValidateUserGroupPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSshKey 删除ssh公钥
//
// 删除ssh公钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) DeleteSshKey(request *model.DeleteSshKeyRequest) (*model.DeleteSshKeyResponse, error) {
	requestDef := GenReqDefForDeleteSshKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSshKeyResponse), nil
	}
}

// DeleteSshKeyInvoker 删除ssh公钥
func (c *CodeArtsRepoClient) DeleteSshKeyInvoker(request *model.DeleteSshKeyRequest) *DeleteSshKeyInvoker {
	requestDef := GenReqDefForDeleteSshKey()
	return &DeleteSshKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImpersonationTokens 获取用户的个人访问令牌
//
// 获取用户的个人访问令牌
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListImpersonationTokens(request *model.ListImpersonationTokensRequest) (*model.ListImpersonationTokensResponse, error) {
	requestDef := GenReqDefForListImpersonationTokens()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImpersonationTokensResponse), nil
	}
}

// ListImpersonationTokensInvoker 获取用户的个人访问令牌
func (c *CodeArtsRepoClient) ListImpersonationTokensInvoker(request *model.ListImpersonationTokensRequest) *ListImpersonationTokensInvoker {
	requestDef := GenReqDefForListImpersonationTokens()
	return &ListImpersonationTokensInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserGpgKeys 获取当前用户的gpg_key列表
//
// 获取当前用户的gpg_key列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListUserGpgKeys(request *model.ListUserGpgKeysRequest) (*model.ListUserGpgKeysResponse, error) {
	requestDef := GenReqDefForListUserGpgKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserGpgKeysResponse), nil
	}
}

// ListUserGpgKeysInvoker 获取当前用户的gpg_key列表
func (c *CodeArtsRepoClient) ListUserGpgKeysInvoker(request *model.ListUserGpgKeysRequest) *ListUserGpgKeysInvoker {
	requestDef := GenReqDefForListUserGpgKeys()
	return &ListUserGpgKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserKeys 获取当前用户的密钥列表
//
// 获取当前用户的密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListUserKeys(request *model.ListUserKeysRequest) (*model.ListUserKeysResponse, error) {
	requestDef := GenReqDefForListUserKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserKeysResponse), nil
	}
}

// ListUserKeysInvoker 获取当前用户的密钥列表
func (c *CodeArtsRepoClient) ListUserKeysInvoker(request *model.ListUserKeysRequest) *ListUserKeysInvoker {
	requestDef := GenReqDefForListUserKeys()
	return &ListUserKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendUserEmailVerifyCode 发送邮箱验证码
//
// 发送邮箱验证码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) SendUserEmailVerifyCode(request *model.SendUserEmailVerifyCodeRequest) (*model.SendUserEmailVerifyCodeResponse, error) {
	requestDef := GenReqDefForSendUserEmailVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendUserEmailVerifyCodeResponse), nil
	}
}

// SendUserEmailVerifyCodeInvoker 发送邮箱验证码
func (c *CodeArtsRepoClient) SendUserEmailVerifyCodeInvoker(request *model.SendUserEmailVerifyCodeRequest) *SendUserEmailVerifyCodeInvoker {
	requestDef := GenReqDefForSendUserEmailVerifyCode()
	return &SendUserEmailVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpsPasswordSetting 获取https的验证方式
//
// 获取https的验证方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowHttpsPasswordSetting(request *model.ShowHttpsPasswordSettingRequest) (*model.ShowHttpsPasswordSettingResponse, error) {
	requestDef := GenReqDefForShowHttpsPasswordSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpsPasswordSettingResponse), nil
	}
}

// ShowHttpsPasswordSettingInvoker 获取https的验证方式
func (c *CodeArtsRepoClient) ShowHttpsPasswordSettingInvoker(request *model.ShowHttpsPasswordSettingRequest) *ShowHttpsPasswordSettingInvoker {
	requestDef := GenReqDefForShowHttpsPasswordSetting()
	return &ShowHttpsPasswordSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserEmails 获取用户相关邮箱信息
//
// 获取用户相关邮箱信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowUserEmails(request *model.ShowUserEmailsRequest) (*model.ShowUserEmailsResponse, error) {
	requestDef := GenReqDefForShowUserEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserEmailsResponse), nil
	}
}

// ShowUserEmailsInvoker 获取用户相关邮箱信息
func (c *CodeArtsRepoClient) ShowUserEmailsInvoker(request *model.ShowUserEmailsRequest) *ShowUserEmailsInvoker {
	requestDef := GenReqDefForShowUserEmails()
	return &ShowUserEmailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpsPasswordSetting 修改https的验证方式
//
// 修改https的验证方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateHttpsPasswordSetting(request *model.UpdateHttpsPasswordSettingRequest) (*model.UpdateHttpsPasswordSettingResponse, error) {
	requestDef := GenReqDefForUpdateHttpsPasswordSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpsPasswordSettingResponse), nil
	}
}

// UpdateHttpsPasswordSettingInvoker 修改https的验证方式
func (c *CodeArtsRepoClient) UpdateHttpsPasswordSettingInvoker(request *model.UpdateHttpsPasswordSettingRequest) *UpdateHttpsPasswordSettingInvoker {
	requestDef := GenReqDefForUpdateHttpsPasswordSetting()
	return &UpdateHttpsPasswordSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserEmails 更新邮箱
//
// 更新邮箱
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateUserEmails(request *model.UpdateUserEmailsRequest) (*model.UpdateUserEmailsResponse, error) {
	requestDef := GenReqDefForUpdateUserEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserEmailsResponse), nil
	}
}

// UpdateUserEmailsInvoker 更新邮箱
func (c *CodeArtsRepoClient) UpdateUserEmailsInvoker(request *model.UpdateUserEmailsRequest) *UpdateUserEmailsInvoker {
	requestDef := GenReqDefForUpdateUserEmails()
	return &UpdateUserEmailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGroupWebhook 添加代码组下Webhook
//
// 添加代码组下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddGroupWebhook(request *model.AddGroupWebhookRequest) (*model.AddGroupWebhookResponse, error) {
	requestDef := GenReqDefForAddGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGroupWebhookResponse), nil
	}
}

// AddGroupWebhookInvoker 添加代码组下Webhook
func (c *CodeArtsRepoClient) AddGroupWebhookInvoker(request *model.AddGroupWebhookRequest) *AddGroupWebhookInvoker {
	requestDef := GenReqDefForAddGroupWebhook()
	return &AddGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProjectWebhook 添加项目下Webhook
//
// 添加项目下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddProjectWebhook(request *model.AddProjectWebhookRequest) (*model.AddProjectWebhookResponse, error) {
	requestDef := GenReqDefForAddProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProjectWebhookResponse), nil
	}
}

// AddProjectWebhookInvoker 添加项目下Webhook
func (c *CodeArtsRepoClient) AddProjectWebhookInvoker(request *model.AddProjectWebhookRequest) *AddProjectWebhookInvoker {
	requestDef := GenReqDefForAddProjectWebhook()
	return &AddProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRepositoryWebhook 添加仓库下Webhook
//
// 添加仓库下Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) AddRepositoryWebhook(request *model.AddRepositoryWebhookRequest) (*model.AddRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForAddRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRepositoryWebhookResponse), nil
	}
}

// AddRepositoryWebhookInvoker 添加仓库下Webhook
func (c *CodeArtsRepoClient) AddRepositoryWebhookInvoker(request *model.AddRepositoryWebhookRequest) *AddRepositoryWebhookInvoker {
	requestDef := GenReqDefForAddRepositoryWebhook()
	return &AddRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupWebhookLogs 获取代码组下指定Webhook的日志列表
//
// 获取代码组下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupWebhookLogs(request *model.ListGroupWebhookLogsRequest) (*model.ListGroupWebhookLogsResponse, error) {
	requestDef := GenReqDefForListGroupWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupWebhookLogsResponse), nil
	}
}

// ListGroupWebhookLogsInvoker 获取代码组下指定Webhook的日志列表
func (c *CodeArtsRepoClient) ListGroupWebhookLogsInvoker(request *model.ListGroupWebhookLogsRequest) *ListGroupWebhookLogsInvoker {
	requestDef := GenReqDefForListGroupWebhookLogs()
	return &ListGroupWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupWebhooks 获取代码组下Webhook列表
//
// 获取代码组下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListGroupWebhooks(request *model.ListGroupWebhooksRequest) (*model.ListGroupWebhooksResponse, error) {
	requestDef := GenReqDefForListGroupWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupWebhooksResponse), nil
	}
}

// ListGroupWebhooksInvoker 获取代码组下Webhook列表
func (c *CodeArtsRepoClient) ListGroupWebhooksInvoker(request *model.ListGroupWebhooksRequest) *ListGroupWebhooksInvoker {
	requestDef := GenReqDefForListGroupWebhooks()
	return &ListGroupWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectWebhookLogs 获取项目下指定Webhook的日志列表
//
// 获取项目下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectWebhookLogs(request *model.ListProjectWebhookLogsRequest) (*model.ListProjectWebhookLogsResponse, error) {
	requestDef := GenReqDefForListProjectWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectWebhookLogsResponse), nil
	}
}

// ListProjectWebhookLogsInvoker 获取项目下指定Webhook的日志列表
func (c *CodeArtsRepoClient) ListProjectWebhookLogsInvoker(request *model.ListProjectWebhookLogsRequest) *ListProjectWebhookLogsInvoker {
	requestDef := GenReqDefForListProjectWebhookLogs()
	return &ListProjectWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectWebhooks 获取项目下Webhook列表
//
// 获取项目下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListProjectWebhooks(request *model.ListProjectWebhooksRequest) (*model.ListProjectWebhooksResponse, error) {
	requestDef := GenReqDefForListProjectWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectWebhooksResponse), nil
	}
}

// ListProjectWebhooksInvoker 获取项目下Webhook列表
func (c *CodeArtsRepoClient) ListProjectWebhooksInvoker(request *model.ListProjectWebhooksRequest) *ListProjectWebhooksInvoker {
	requestDef := GenReqDefForListProjectWebhooks()
	return &ListProjectWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWebhookLogs 获取仓库下指定Webhook的日志列表
//
// 获取仓库下指定Webhook的日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryWebhookLogs(request *model.ListRepositoryWebhookLogsRequest) (*model.ListRepositoryWebhookLogsResponse, error) {
	requestDef := GenReqDefForListRepositoryWebhookLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWebhookLogsResponse), nil
	}
}

// ListRepositoryWebhookLogsInvoker 获取仓库下指定Webhook的日志列表
func (c *CodeArtsRepoClient) ListRepositoryWebhookLogsInvoker(request *model.ListRepositoryWebhookLogsRequest) *ListRepositoryWebhookLogsInvoker {
	requestDef := GenReqDefForListRepositoryWebhookLogs()
	return &ListRepositoryWebhookLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryWebhooks 获取仓库下Webhook列表
//
// 获取仓库下Webhook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ListRepositoryWebhooks(request *model.ListRepositoryWebhooksRequest) (*model.ListRepositoryWebhooksResponse, error) {
	requestDef := GenReqDefForListRepositoryWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryWebhooksResponse), nil
	}
}

// ListRepositoryWebhooksInvoker 获取仓库下Webhook列表
func (c *CodeArtsRepoClient) ListRepositoryWebhooksInvoker(request *model.ListRepositoryWebhooksRequest) *ListRepositoryWebhooksInvoker {
	requestDef := GenReqDefForListRepositoryWebhooks()
	return &ListRepositoryWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveGroupWebhook 删除代码组下单个Webhook
//
// 删除代码组下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RemoveGroupWebhook(request *model.RemoveGroupWebhookRequest) (*model.RemoveGroupWebhookResponse, error) {
	requestDef := GenReqDefForRemoveGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveGroupWebhookResponse), nil
	}
}

// RemoveGroupWebhookInvoker 删除代码组下单个Webhook
func (c *CodeArtsRepoClient) RemoveGroupWebhookInvoker(request *model.RemoveGroupWebhookRequest) *RemoveGroupWebhookInvoker {
	requestDef := GenReqDefForRemoveGroupWebhook()
	return &RemoveGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveProjectWebhook 删除项目下单个Webhook
//
// 删除项目下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RemoveProjectWebhook(request *model.RemoveProjectWebhookRequest) (*model.RemoveProjectWebhookResponse, error) {
	requestDef := GenReqDefForRemoveProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveProjectWebhookResponse), nil
	}
}

// RemoveProjectWebhookInvoker 删除项目下单个Webhook
func (c *CodeArtsRepoClient) RemoveProjectWebhookInvoker(request *model.RemoveProjectWebhookRequest) *RemoveProjectWebhookInvoker {
	requestDef := GenReqDefForRemoveProjectWebhook()
	return &RemoveProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveRepositoryWebhook 删除仓库下单个Webhook
//
// 删除仓库下单个Webhook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) RemoveRepositoryWebhook(request *model.RemoveRepositoryWebhookRequest) (*model.RemoveRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForRemoveRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveRepositoryWebhookResponse), nil
	}
}

// RemoveRepositoryWebhookInvoker 删除仓库下单个Webhook
func (c *CodeArtsRepoClient) RemoveRepositoryWebhookInvoker(request *model.RemoveRepositoryWebhookRequest) *RemoveRepositoryWebhookInvoker {
	requestDef := GenReqDefForRemoveRepositoryWebhook()
	return &RemoveRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWebhook 获取代码组下单个Webhook数据
//
// 获取代码组下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupWebhook(request *model.ShowGroupWebhookRequest) (*model.ShowGroupWebhookResponse, error) {
	requestDef := GenReqDefForShowGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWebhookResponse), nil
	}
}

// ShowGroupWebhookInvoker 获取代码组下单个Webhook数据
func (c *CodeArtsRepoClient) ShowGroupWebhookInvoker(request *model.ShowGroupWebhookRequest) *ShowGroupWebhookInvoker {
	requestDef := GenReqDefForShowGroupWebhook()
	return &ShowGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupWebhookLog 获取代码组下指定Webhook的指定日志详情
//
// 获取代码组下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowGroupWebhookLog(request *model.ShowGroupWebhookLogRequest) (*model.ShowGroupWebhookLogResponse, error) {
	requestDef := GenReqDefForShowGroupWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupWebhookLogResponse), nil
	}
}

// ShowGroupWebhookLogInvoker 获取代码组下指定Webhook的指定日志详情
func (c *CodeArtsRepoClient) ShowGroupWebhookLogInvoker(request *model.ShowGroupWebhookLogRequest) *ShowGroupWebhookLogInvoker {
	requestDef := GenReqDefForShowGroupWebhookLog()
	return &ShowGroupWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWebhook 获取项目下单个Webhook数据
//
// 获取项目下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectWebhook(request *model.ShowProjectWebhookRequest) (*model.ShowProjectWebhookResponse, error) {
	requestDef := GenReqDefForShowProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWebhookResponse), nil
	}
}

// ShowProjectWebhookInvoker 获取项目下单个Webhook数据
func (c *CodeArtsRepoClient) ShowProjectWebhookInvoker(request *model.ShowProjectWebhookRequest) *ShowProjectWebhookInvoker {
	requestDef := GenReqDefForShowProjectWebhook()
	return &ShowProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWebhookLog 获取项目下指定Webhook的指定日志详情
//
// 获取项目下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowProjectWebhookLog(request *model.ShowProjectWebhookLogRequest) (*model.ShowProjectWebhookLogResponse, error) {
	requestDef := GenReqDefForShowProjectWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWebhookLogResponse), nil
	}
}

// ShowProjectWebhookLogInvoker 获取项目下指定Webhook的指定日志详情
func (c *CodeArtsRepoClient) ShowProjectWebhookLogInvoker(request *model.ShowProjectWebhookLogRequest) *ShowProjectWebhookLogInvoker {
	requestDef := GenReqDefForShowProjectWebhookLog()
	return &ShowProjectWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWebhook 获取仓库下单个Webhook数据
//
// 获取仓库下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryWebhook(request *model.ShowRepositoryWebhookRequest) (*model.ShowRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForShowRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWebhookResponse), nil
	}
}

// ShowRepositoryWebhookInvoker 获取仓库下单个Webhook数据
func (c *CodeArtsRepoClient) ShowRepositoryWebhookInvoker(request *model.ShowRepositoryWebhookRequest) *ShowRepositoryWebhookInvoker {
	requestDef := GenReqDefForShowRepositoryWebhook()
	return &ShowRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepositoryWebhookLog 获取仓库下指定Webhook的指定日志详情
//
// 获取仓库下指定Webhook的指定日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) ShowRepositoryWebhookLog(request *model.ShowRepositoryWebhookLogRequest) (*model.ShowRepositoryWebhookLogResponse, error) {
	requestDef := GenReqDefForShowRepositoryWebhookLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryWebhookLogResponse), nil
	}
}

// ShowRepositoryWebhookLogInvoker 获取仓库下指定Webhook的指定日志详情
func (c *CodeArtsRepoClient) ShowRepositoryWebhookLogInvoker(request *model.ShowRepositoryWebhookLogRequest) *ShowRepositoryWebhookLogInvoker {
	requestDef := GenReqDefForShowRepositoryWebhookLog()
	return &ShowRepositoryWebhookLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupWebhook 更新代码组下单个Webhook数据
//
// 更新代码组下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateGroupWebhook(request *model.UpdateGroupWebhookRequest) (*model.UpdateGroupWebhookResponse, error) {
	requestDef := GenReqDefForUpdateGroupWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupWebhookResponse), nil
	}
}

// UpdateGroupWebhookInvoker 更新代码组下单个Webhook数据
func (c *CodeArtsRepoClient) UpdateGroupWebhookInvoker(request *model.UpdateGroupWebhookRequest) *UpdateGroupWebhookInvoker {
	requestDef := GenReqDefForUpdateGroupWebhook()
	return &UpdateGroupWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectWebhook 更新项目下单个Webhook数据
//
// 更新项目下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateProjectWebhook(request *model.UpdateProjectWebhookRequest) (*model.UpdateProjectWebhookResponse, error) {
	requestDef := GenReqDefForUpdateProjectWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectWebhookResponse), nil
	}
}

// UpdateProjectWebhookInvoker 更新项目下单个Webhook数据
func (c *CodeArtsRepoClient) UpdateProjectWebhookInvoker(request *model.UpdateProjectWebhookRequest) *UpdateProjectWebhookInvoker {
	requestDef := GenReqDefForUpdateProjectWebhook()
	return &UpdateProjectWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepositoryWebhook 更新仓库下单个Webhook数据
//
// 更新仓库下单个Webhook数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsRepoClient) UpdateRepositoryWebhook(request *model.UpdateRepositoryWebhookRequest) (*model.UpdateRepositoryWebhookResponse, error) {
	requestDef := GenReqDefForUpdateRepositoryWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepositoryWebhookResponse), nil
	}
}

// UpdateRepositoryWebhookInvoker 更新仓库下单个Webhook数据
func (c *CodeArtsRepoClient) UpdateRepositoryWebhookInvoker(request *model.UpdateRepositoryWebhookRequest) *UpdateRepositoryWebhookInvoker {
	requestDef := GenReqDefForUpdateRepositoryWebhook()
	return &UpdateRepositoryWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
