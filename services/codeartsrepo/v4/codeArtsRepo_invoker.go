package v4

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsrepo/v4/model"
)

type CreateCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCommitInvoker) Invoke() (*model.CreateCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCommitResponse), nil
	}
}

type CreateCommitRevertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCommitRevertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCommitRevertInvoker) Invoke() (*model.CreateCommitRevertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCommitRevertResponse), nil
	}
}

type ListCommitAssociatedRefsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitAssociatedRefsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitAssociatedRefsInvoker) Invoke() (*model.ListCommitAssociatedRefsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitAssociatedRefsResponse), nil
	}
}

type ListCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitsInvoker) Invoke() (*model.ListCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitsResponse), nil
	}
}

type ShowCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitInvoker) Invoke() (*model.ShowCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitResponse), nil
	}
}

type ShowCommitDiffMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitDiffMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitDiffMetadataInvoker) Invoke() (*model.ShowCommitDiffMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitDiffMetadataResponse), nil
	}
}

type ShowCommitFileDiffInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitFileDiffInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitFileDiffInvoker) Invoke() (*model.ShowCommitFileDiffResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitFileDiffResponse), nil
	}
}

type ShowDiffCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiffCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiffCommitInvoker) Invoke() (*model.ShowDiffCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiffCommitResponse), nil
	}
}

type CreateMergeRequestDiscussionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestDiscussionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestDiscussionInvoker) Invoke() (*model.CreateMergeRequestDiscussionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestDiscussionResponse), nil
	}
}

type CreateMergeRequestDiscussionResponseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestDiscussionResponseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestDiscussionResponseInvoker) Invoke() (*model.CreateMergeRequestDiscussionResponseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestDiscussionResponseResponse), nil
	}
}

type CreateReviewSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReviewSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReviewSettingInvoker) Invoke() (*model.CreateReviewSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReviewSettingResponse), nil
	}
}

type DeleteMergeRequestDiscussionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMergeRequestDiscussionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMergeRequestDiscussionInvoker) Invoke() (*model.DeleteMergeRequestDiscussionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMergeRequestDiscussionResponse), nil
	}
}

type ListCommitDiscussionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitDiscussionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitDiscussionsInvoker) Invoke() (*model.ListCommitDiscussionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitDiscussionsResponse), nil
	}
}

type ListDefaultReviewCategoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDefaultReviewCategoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDefaultReviewCategoriesInvoker) Invoke() (*model.ListDefaultReviewCategoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDefaultReviewCategoriesResponse), nil
	}
}

type ListMergeRequestDiscussionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestDiscussionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestDiscussionsInvoker) Invoke() (*model.ListMergeRequestDiscussionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestDiscussionsResponse), nil
	}
}

type ListMergeRequestSystemNotesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestSystemNotesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestSystemNotesInvoker) Invoke() (*model.ListMergeRequestSystemNotesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestSystemNotesResponse), nil
	}
}

type ListProjectNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectNoteRequiredAttributesInvoker) Invoke() (*model.ListProjectNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectNoteRequiredAttributesResponse), nil
	}
}

type ListRepositoryReviewAuthorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryReviewAuthorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryReviewAuthorsInvoker) Invoke() (*model.ListRepositoryReviewAuthorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryReviewAuthorsResponse), nil
	}
}

type ListRepositoryReviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryReviewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryReviewsInvoker) Invoke() (*model.ListRepositoryReviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryReviewsResponse), nil
	}
}

type ShowGroupNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupNoteRequiredAttributesInvoker) Invoke() (*model.ShowGroupNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupNoteRequiredAttributesResponse), nil
	}
}

type ShowGroupReviewSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupReviewSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupReviewSettingsInvoker) Invoke() (*model.ShowGroupReviewSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupReviewSettingsResponse), nil
	}
}

type ShowMergeRequestDiscussionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestDiscussionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestDiscussionInvoker) Invoke() (*model.ShowMergeRequestDiscussionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestDiscussionResponse), nil
	}
}

type ShowNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNoteRequiredAttributesInvoker) Invoke() (*model.ShowNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNoteRequiredAttributesResponse), nil
	}
}

type ShowProjectReviewSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectReviewSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectReviewSettingsInvoker) Invoke() (*model.ShowProjectReviewSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectReviewSettingsResponse), nil
	}
}

type ShowReviewSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReviewSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReviewSettingInvoker) Invoke() (*model.ShowReviewSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReviewSettingResponse), nil
	}
}

type UpdateGroupNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupNoteRequiredAttributesInvoker) Invoke() (*model.UpdateGroupNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupNoteRequiredAttributesResponse), nil
	}
}

type UpdateGroupReviewSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupReviewSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupReviewSettingsInvoker) Invoke() (*model.UpdateGroupReviewSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupReviewSettingsResponse), nil
	}
}

type UpdateMergeRequestDiscussionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestDiscussionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestDiscussionInvoker) Invoke() (*model.UpdateMergeRequestDiscussionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestDiscussionResponse), nil
	}
}

type UpdateMergeRequestDiscussionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestDiscussionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestDiscussionInfoInvoker) Invoke() (*model.UpdateMergeRequestDiscussionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestDiscussionInfoResponse), nil
	}
}

type UpdateNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNoteRequiredAttributesInvoker) Invoke() (*model.UpdateNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNoteRequiredAttributesResponse), nil
	}
}

type UpdateProjectNoteRequiredAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectNoteRequiredAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectNoteRequiredAttributesInvoker) Invoke() (*model.UpdateProjectNoteRequiredAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectNoteRequiredAttributesResponse), nil
	}
}

type UpdateProjectReviewSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectReviewSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectReviewSettingsInvoker) Invoke() (*model.UpdateProjectReviewSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectReviewSettingsResponse), nil
	}
}

type CreateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFileInvoker) Invoke() (*model.CreateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFileResponse), nil
	}
}

type DeleteFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFileInvoker) Invoke() (*model.DeleteFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFileResponse), nil
	}
}

type DownloadBlobsRawInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBlobsRawInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadBlobsRawInvoker) Invoke() (*model.DownloadBlobsRawResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBlobsRawResponse), nil
	}
}

type ListFileBlameLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileBlameLinesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileBlameLinesInvoker) Invoke() (*model.ListFileBlameLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileBlameLinesResponse), nil
	}
}

type ListFileUpperTreeEntriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileUpperTreeEntriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileUpperTreeEntriesInvoker) Invoke() (*model.ListFileUpperTreeEntriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileUpperTreeEntriesResponse), nil
	}
}

type ListFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFilesInvoker) Invoke() (*model.ListFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFilesResponse), nil
	}
}

type ListLogsTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogsTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLogsTreeInvoker) Invoke() (*model.ListLogsTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogsTreeResponse), nil
	}
}

type ListTreesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTreesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTreesInvoker) Invoke() (*model.ListTreesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTreesResponse), nil
	}
}

type RenameFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenameFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenameFileInvoker) Invoke() (*model.RenameFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenameFileResponse), nil
	}
}

type ShowFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileInvoker) Invoke() (*model.ShowFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileResponse), nil
	}
}

type ShowFileContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileContentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileContentInvoker) Invoke() (*model.ShowFileContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileContentResponse), nil
	}
}

type ShowFileRawInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileRawInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileRawInvoker) Invoke() (*model.ShowFileRawResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileRawResponse), nil
	}
}

type ShowReadmeFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReadmeFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReadmeFileInvoker) Invoke() (*model.ShowReadmeFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReadmeFileResponse), nil
	}
}

type UpdateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFileInvoker) Invoke() (*model.UpdateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFileResponse), nil
	}
}

type AssociateGroupUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateGroupUserGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateGroupUserGroupInvoker) Invoke() (*model.AssociateGroupUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateGroupUserGroupResponse), nil
	}
}

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupInvoker) Invoke() (*model.CreateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupResponse), nil
	}
}

type DeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type ListGroupAddableMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupAddableMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupAddableMembersInvoker) Invoke() (*model.ListGroupAddableMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupAddableMembersResponse), nil
	}
}

type ListGroupAddableUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupAddableUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupAddableUserGroupsInvoker) Invoke() (*model.ListGroupAddableUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupAddableUserGroupsResponse), nil
	}
}

type ListGroupMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMembersInvoker) Invoke() (*model.ListGroupMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMembersResponse), nil
	}
}

type ListGroupPermissionResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupPermissionResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupPermissionResourcesInvoker) Invoke() (*model.ListGroupPermissionResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupPermissionResourcesResponse), nil
	}
}

type ListGroupSubgroupsAndRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupSubgroupsAndRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupSubgroupsAndRepositoriesInvoker) Invoke() (*model.ListGroupSubgroupsAndRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupSubgroupsAndRepositoriesResponse), nil
	}
}

type ListGroupUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupUserGroupsInvoker) Invoke() (*model.ListGroupUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupUserGroupsResponse), nil
	}
}

type ListGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsInvoker) Invoke() (*model.ListGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsResponse), nil
	}
}

type ListManageableGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManageableGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManageableGroupsInvoker) Invoke() (*model.ListManageableGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManageableGroupsResponse), nil
	}
}

type ShowGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupInvoker) Invoke() (*model.ShowGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupResponse), nil
	}
}

type ShowGroupGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupGeneralPolicyInvoker) Invoke() (*model.ShowGroupGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupGeneralPolicyResponse), nil
	}
}

type ShowGroupInheritSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupInheritSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupInheritSettingInvoker) Invoke() (*model.ShowGroupInheritSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupInheritSettingResponse), nil
	}
}

type ShowGroupPermissionInheritEnabledInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupPermissionInheritEnabledInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupPermissionInheritEnabledInvoker) Invoke() (*model.ShowGroupPermissionInheritEnabledResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupPermissionInheritEnabledResponse), nil
	}
}

type ShowGroupSettingsInheritCfgInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupSettingsInheritCfgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupSettingsInheritCfgInvoker) Invoke() (*model.ShowGroupSettingsInheritCfgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupSettingsInheritCfgResponse), nil
	}
}

type ShowGroupWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupWatermarkInvoker) Invoke() (*model.ShowGroupWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupWatermarkResponse), nil
	}
}

type ShowGroupsGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupsGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupsGeneralPolicyInvoker) Invoke() (*model.ShowGroupsGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupsGeneralPolicyResponse), nil
	}
}

type ShowGroupsInheritInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupsInheritInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupsInheritInvoker) Invoke() (*model.ShowGroupsInheritResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupsInheritResponse), nil
	}
}

type TransferGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *TransferGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TransferGroupInvoker) Invoke() (*model.TransferGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TransferGroupResponse), nil
	}
}

type UpdateGroupGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupGeneralPolicyInvoker) Invoke() (*model.UpdateGroupGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupGeneralPolicyResponse), nil
	}
}

type UpdateGroupWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupWatermarkInvoker) Invoke() (*model.UpdateGroupWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupWatermarkResponse), nil
	}
}

type AddRepositoryMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRepositoryMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRepositoryMembersInvoker) Invoke() (*model.AddRepositoryMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRepositoryMembersResponse), nil
	}
}

type ListGroupProtectedRefsUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupProtectedRefsUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupProtectedRefsUserGroupsInvoker) Invoke() (*model.ListGroupProtectedRefsUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupProtectedRefsUserGroupsResponse), nil
	}
}

type ListMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMembersInvoker) Invoke() (*model.ListMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMembersResponse), nil
	}
}

type ListProductPermissionResourcesGrantedUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductPermissionResourcesGrantedUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductPermissionResourcesGrantedUsersInvoker) Invoke() (*model.ListProductPermissionResourcesGrantedUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductPermissionResourcesGrantedUsersResponse), nil
	}
}

type ListProjectProtectedRefsUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectProtectedRefsUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectProtectedRefsUserGroupsInvoker) Invoke() (*model.ListProjectProtectedRefsUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectProtectedRefsUserGroupsResponse), nil
	}
}

type ListRepositoryProtectedRefsUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryProtectedRefsUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryProtectedRefsUserGroupsInvoker) Invoke() (*model.ListRepositoryProtectedRefsUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryProtectedRefsUserGroupsResponse), nil
	}
}

type ListRepositoryUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryUserGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryUserGroupsInvoker) Invoke() (*model.ListRepositoryUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryUserGroupsResponse), nil
	}
}

type ApprovalMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApprovalMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApprovalMergeRequestInvoker) Invoke() (*model.ApprovalMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApprovalMergeRequestResponse), nil
	}
}

type CreateCherryPickMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCherryPickMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCherryPickMergeRequestInvoker) Invoke() (*model.CreateCherryPickMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCherryPickMergeRequestResponse), nil
	}
}

type CreateGroupMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupMergeRequestApproverSettingInvoker) Invoke() (*model.CreateGroupMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupMergeRequestApproverSettingResponse), nil
	}
}

type CreateGroupMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupMergeRequestTemplateInvoker) Invoke() (*model.CreateGroupMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupMergeRequestTemplateResponse), nil
	}
}

type CreateMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestInvoker) Invoke() (*model.CreateMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestResponse), nil
	}
}

type CreateMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestApproverSettingInvoker) Invoke() (*model.CreateMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestApproverSettingResponse), nil
	}
}

type CreateMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestTemplateInvoker) Invoke() (*model.CreateMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestTemplateResponse), nil
	}
}

type CreateProjectMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectMergeRequestApproverSettingInvoker) Invoke() (*model.CreateProjectMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectMergeRequestApproverSettingResponse), nil
	}
}

type CreateProjectMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectMergeRequestTemplateInvoker) Invoke() (*model.CreateProjectMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectMergeRequestTemplateResponse), nil
	}
}

type DeleteGroupMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupMergeRequestApproverSettingInvoker) Invoke() (*model.DeleteGroupMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupMergeRequestApproverSettingResponse), nil
	}
}

type DeleteGroupMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupMergeRequestTemplateInvoker) Invoke() (*model.DeleteGroupMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupMergeRequestTemplateResponse), nil
	}
}

type DeleteMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMergeRequestApproverSettingInvoker) Invoke() (*model.DeleteMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMergeRequestApproverSettingResponse), nil
	}
}

type DeleteMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMergeRequestTemplateInvoker) Invoke() (*model.DeleteMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMergeRequestTemplateResponse), nil
	}
}

type DeleteMergeRequestVoteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMergeRequestVoteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMergeRequestVoteInvoker) Invoke() (*model.DeleteMergeRequestVoteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMergeRequestVoteResponse), nil
	}
}

type DeleteProjectMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProjectMergeRequestApproverSettingInvoker) Invoke() (*model.DeleteProjectMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectMergeRequestApproverSettingResponse), nil
	}
}

type DeleteProjectMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProjectMergeRequestTemplateInvoker) Invoke() (*model.DeleteProjectMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectMergeRequestTemplateResponse), nil
	}
}

type ImportMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportMergeRequestInvoker) Invoke() (*model.ImportMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportMergeRequestResponse), nil
	}
}

type ListCommitAssociatedMergeRequestsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitAssociatedMergeRequestsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitAssociatedMergeRequestsInvoker) Invoke() (*model.ListCommitAssociatedMergeRequestsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitAssociatedMergeRequestsResponse), nil
	}
}

type ListDiscussionTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiscussionTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiscussionTemplatesInvoker) Invoke() (*model.ListDiscussionTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiscussionTemplatesResponse), nil
	}
}

type ListGroupMergeRequestApproverSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMergeRequestApproverSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMergeRequestApproverSettingsInvoker) Invoke() (*model.ListGroupMergeRequestApproverSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMergeRequestApproverSettingsResponse), nil
	}
}

type ListGroupMergeRequestCanBeAssignedReviewersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMergeRequestCanBeAssignedReviewersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMergeRequestCanBeAssignedReviewersInvoker) Invoke() (*model.ListGroupMergeRequestCanBeAssignedReviewersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

type ListGroupMergeRequestTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMergeRequestTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMergeRequestTemplatesInvoker) Invoke() (*model.ListGroupMergeRequestTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMergeRequestTemplatesResponse), nil
	}
}

type ListGroupMergeRequestValidAssignedCandidatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMergeRequestValidAssignedCandidatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMergeRequestValidAssignedCandidatesInvoker) Invoke() (*model.ListGroupMergeRequestValidAssignedCandidatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMergeRequestValidAssignedCandidatesResponse), nil
	}
}

type ListMergeRequestApproverSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestApproverSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestApproverSettingsInvoker) Invoke() (*model.ListMergeRequestApproverSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestApproverSettingsResponse), nil
	}
}

type ListMergeRequestApproversInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestApproversInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestApproversInvoker) Invoke() (*model.ListMergeRequestApproversResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestApproversResponse), nil
	}
}

type ListMergeRequestChangesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestChangesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestChangesInvoker) Invoke() (*model.ListMergeRequestChangesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestChangesResponse), nil
	}
}

type ListMergeRequestChangesTreesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestChangesTreesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestChangesTreesInvoker) Invoke() (*model.ListMergeRequestChangesTreesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestChangesTreesResponse), nil
	}
}

type ListMergeRequestCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestCommitsInvoker) Invoke() (*model.ListMergeRequestCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestCommitsResponse), nil
	}
}

type ListMergeRequestConflictFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestConflictFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestConflictFilesInvoker) Invoke() (*model.ListMergeRequestConflictFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestConflictFilesResponse), nil
	}
}

type ListMergeRequestEvaluationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestEvaluationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestEvaluationsInvoker) Invoke() (*model.ListMergeRequestEvaluationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestEvaluationsResponse), nil
	}
}

type ListMergeRequestParticipantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestParticipantsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestParticipantsInvoker) Invoke() (*model.ListMergeRequestParticipantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestParticipantsResponse), nil
	}
}

type ListMergeRequestReviewersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestReviewersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestReviewersInvoker) Invoke() (*model.ListMergeRequestReviewersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestReviewersResponse), nil
	}
}

type ListMergeRequestTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestTemplatesInvoker) Invoke() (*model.ListMergeRequestTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestTemplatesResponse), nil
	}
}

type ListMergeRequestValidAssignedCandidatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestValidAssignedCandidatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestValidAssignedCandidatesInvoker) Invoke() (*model.ListMergeRequestValidAssignedCandidatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestValidAssignedCandidatesResponse), nil
	}
}

type ListMergeRequestVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestVersionsInvoker) Invoke() (*model.ListMergeRequestVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestVersionsResponse), nil
	}
}

type ListPersonalMergeRequestsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPersonalMergeRequestsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPersonalMergeRequestsInvoker) Invoke() (*model.ListPersonalMergeRequestsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPersonalMergeRequestsResponse), nil
	}
}

type ListProjectMergeRequestApproverSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectMergeRequestApproverSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectMergeRequestApproverSettingsInvoker) Invoke() (*model.ListProjectMergeRequestApproverSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectMergeRequestApproverSettingsResponse), nil
	}
}

type ListProjectMergeRequestCanBeAssignedReviewersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectMergeRequestCanBeAssignedReviewersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectMergeRequestCanBeAssignedReviewersInvoker) Invoke() (*model.ListProjectMergeRequestCanBeAssignedReviewersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectMergeRequestCanBeAssignedReviewersResponse), nil
	}
}

type ListProjectMergeRequestCanBeAssignedUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectMergeRequestCanBeAssignedUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectMergeRequestCanBeAssignedUsersInvoker) Invoke() (*model.ListProjectMergeRequestCanBeAssignedUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectMergeRequestCanBeAssignedUsersResponse), nil
	}
}

type ListProjectMergeRequestTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectMergeRequestTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectMergeRequestTemplatesInvoker) Invoke() (*model.ListProjectMergeRequestTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectMergeRequestTemplatesResponse), nil
	}
}

type ListRepositoryMergeRequestsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryMergeRequestsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryMergeRequestsInvoker) Invoke() (*model.ListRepositoryMergeRequestsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryMergeRequestsResponse), nil
	}
}

type MergeMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *MergeMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MergeMergeRequestInvoker) Invoke() (*model.MergeMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MergeMergeRequestResponse), nil
	}
}

type RebaseMergeRequestForOpenApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebaseMergeRequestForOpenApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebaseMergeRequestForOpenApiInvoker) Invoke() (*model.RebaseMergeRequestForOpenApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebaseMergeRequestForOpenApiResponse), nil
	}
}

type ResolveMergeRequestConflictsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResolveMergeRequestConflictsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResolveMergeRequestConflictsInvoker) Invoke() (*model.ResolveMergeRequestConflictsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResolveMergeRequestConflictsResponse), nil
	}
}

type ReviewMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReviewMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReviewMergeRequestInvoker) Invoke() (*model.ReviewMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewMergeRequestResponse), nil
	}
}

type ShowActualHeadPipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowActualHeadPipelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowActualHeadPipelineInvoker) Invoke() (*model.ShowActualHeadPipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowActualHeadPipelineResponse), nil
	}
}

type ShowAverageEvaluationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAverageEvaluationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAverageEvaluationInvoker) Invoke() (*model.ShowAverageEvaluationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAverageEvaluationResponse), nil
	}
}

type ShowBranchConflictInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBranchConflictInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBranchConflictInvoker) Invoke() (*model.ShowBranchConflictResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBranchConflictResponse), nil
	}
}

type ShowCommitCommentsByLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitCommentsByLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitCommentsByLineInvoker) Invoke() (*model.ShowCommitCommentsByLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitCommentsByLineResponse), nil
	}
}

type ShowGroupMergeRequestSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupMergeRequestSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupMergeRequestSettingInvoker) Invoke() (*model.ShowGroupMergeRequestSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupMergeRequestSettingResponse), nil
	}
}

type ShowMergeRequestCommentsByLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestCommentsByLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestCommentsByLineInvoker) Invoke() (*model.ShowMergeRequestCommentsByLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestCommentsByLineResponse), nil
	}
}

type ShowMergeRequestDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestDetailInvoker) Invoke() (*model.ShowMergeRequestDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestDetailResponse), nil
	}
}

type ShowMergeRequestSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestSettingInvoker) Invoke() (*model.ShowMergeRequestSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestSettingResponse), nil
	}
}

type ShowMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestTemplateInvoker) Invoke() (*model.ShowMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestTemplateResponse), nil
	}
}

type ShowMergeRequestVotesDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestVotesDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestVotesDetailInvoker) Invoke() (*model.ShowMergeRequestVotesDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestVotesDetailResponse), nil
	}
}

type ShowMergeableStateOuterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeableStateOuterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeableStateOuterInvoker) Invoke() (*model.ShowMergeableStateOuterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeableStateOuterResponse), nil
	}
}

type ShowProjectMergeRequestSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectMergeRequestSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectMergeRequestSettingInvoker) Invoke() (*model.ShowProjectMergeRequestSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectMergeRequestSettingResponse), nil
	}
}

type ShowRepositoryMergeRequestsStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryMergeRequestsStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryMergeRequestsStatisticInvoker) Invoke() (*model.ShowRepositoryMergeRequestsStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryMergeRequestsStatisticResponse), nil
	}
}

type UpdateGroupMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupMergeRequestApproverSettingInvoker) Invoke() (*model.UpdateGroupMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupMergeRequestApproverSettingResponse), nil
	}
}

type UpdateGroupMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupMergeRequestTemplateInvoker) Invoke() (*model.UpdateGroupMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupMergeRequestTemplateResponse), nil
	}
}

type UpdateMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestInvoker) Invoke() (*model.UpdateMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestResponse), nil
	}
}

type UpdateMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestApproverSettingInvoker) Invoke() (*model.UpdateMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestApproverSettingResponse), nil
	}
}

type UpdateMergeRequestApproversInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestApproversInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestApproversInvoker) Invoke() (*model.UpdateMergeRequestApproversResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestApproversResponse), nil
	}
}

type UpdateMergeRequestReviewersInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestReviewersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestReviewersInvoker) Invoke() (*model.UpdateMergeRequestReviewersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestReviewersResponse), nil
	}
}

type UpdateMergeRequestSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestSettingInvoker) Invoke() (*model.UpdateMergeRequestSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestSettingResponse), nil
	}
}

type UpdateMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestTemplateInvoker) Invoke() (*model.UpdateMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestTemplateResponse), nil
	}
}

type UpdateMergeRequestVoteInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestVoteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestVoteInvoker) Invoke() (*model.UpdateMergeRequestVoteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestVoteResponse), nil
	}
}

type UpdateProjectMergeRequestApproverSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectMergeRequestApproverSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectMergeRequestApproverSettingInvoker) Invoke() (*model.UpdateProjectMergeRequestApproverSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectMergeRequestApproverSettingResponse), nil
	}
}

type UpdateProjectMergeRequestTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectMergeRequestTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectMergeRequestTemplateInvoker) Invoke() (*model.UpdateProjectMergeRequestTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectMergeRequestTemplateResponse), nil
	}
}

type BatchDeleteRepositoryFilePushPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRepositoryFilePushPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRepositoryFilePushPermissionsInvoker) Invoke() (*model.BatchDeleteRepositoryFilePushPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRepositoryFilePushPermissionsResponse), nil
	}
}

type BatchUpdateRepositoryFilePushPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateRepositoryFilePushPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateRepositoryFilePushPermissionsInvoker) Invoke() (*model.BatchUpdateRepositoryFilePushPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateRepositoryFilePushPermissionsResponse), nil
	}
}

type CreateFilePushPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFilePushPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFilePushPermissionInvoker) Invoke() (*model.CreateFilePushPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFilePushPermissionResponse), nil
	}
}

type ListRepositoryFilePushPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryFilePushPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryFilePushPermissionsInvoker) Invoke() (*model.ListRepositoryFilePushPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryFilePushPermissionsResponse), nil
	}
}

type ListRepositoryResourcePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryResourcePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryResourcePermissionsInvoker) Invoke() (*model.ListRepositoryResourcePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryResourcePermissionsResponse), nil
	}
}

type ShowRepositoryPermissionInheritEnabledInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryPermissionInheritEnabledInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryPermissionInheritEnabledInvoker) Invoke() (*model.ShowRepositoryPermissionInheritEnabledResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryPermissionInheritEnabledResponse), nil
	}
}

type UpdateRepositoryPermissionInheritEnabledInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryPermissionInheritEnabledInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryPermissionInheritEnabledInvoker) Invoke() (*model.UpdateRepositoryPermissionInheritEnabledResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryPermissionInheritEnabledResponse), nil
	}
}

type UpdateRepositoryResourcePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryResourcePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryResourcePermissionsInvoker) Invoke() (*model.UpdateRepositoryResourcePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryResourcePermissionsResponse), nil
	}
}

type ListLatestPipelineJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLatestPipelineJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLatestPipelineJobsInvoker) Invoke() (*model.ListLatestPipelineJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLatestPipelineJobsResponse), nil
	}
}

type ListPipelineJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPipelineJobsInvoker) Invoke() (*model.ListPipelineJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineJobsResponse), nil
	}
}

type ListItemCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListItemCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListItemCommitsInvoker) Invoke() (*model.ListItemCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListItemCommitsResponse), nil
	}
}

type ListProjectSubgroupsAndRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectSubgroupsAndRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectSubgroupsAndRepositoriesInvoker) Invoke() (*model.ListProjectSubgroupsAndRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectSubgroupsAndRepositoriesResponse), nil
	}
}

type ShowProjectGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectGeneralPolicyInvoker) Invoke() (*model.ShowProjectGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectGeneralPolicyResponse), nil
	}
}

type ShowProjectMemberSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectMemberSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectMemberSettingInvoker) Invoke() (*model.ShowProjectMemberSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectMemberSettingResponse), nil
	}
}

type ShowProjectSettingsInheritCfgInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectSettingsInheritCfgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectSettingsInheritCfgInvoker) Invoke() (*model.ShowProjectSettingsInheritCfgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectSettingsInheritCfgResponse), nil
	}
}

type ShowProjectWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectWatermarkInvoker) Invoke() (*model.ShowProjectWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectWatermarkResponse), nil
	}
}

type ShowProjectsGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectsGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectsGeneralPolicyInvoker) Invoke() (*model.ShowProjectsGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectsGeneralPolicyResponse), nil
	}
}

type ShowResourcePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourcePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourcePermissionsInvoker) Invoke() (*model.ShowResourcePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourcePermissionsResponse), nil
	}
}

type UpdateProjectGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectGeneralPolicyInvoker) Invoke() (*model.UpdateProjectGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectGeneralPolicyResponse), nil
	}
}

type UpdateProjectSettingsInheritCfgInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectSettingsInheritCfgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectSettingsInheritCfgInvoker) Invoke() (*model.UpdateProjectSettingsInheritCfgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectSettingsInheritCfgResponse), nil
	}
}

type UpdateProjectWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectWatermarkInvoker) Invoke() (*model.UpdateProjectWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectWatermarkResponse), nil
	}
}

type BatchCreateProtectedBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateProtectedBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateProtectedBranchInvoker) Invoke() (*model.BatchCreateProtectedBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateProtectedBranchResponse), nil
	}
}

type BatchCreateProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateProtectedTagsInvoker) Invoke() (*model.BatchCreateProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateProtectedTagsResponse), nil
	}
}

type BatchDeleteProtectedBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteProtectedBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteProtectedBranchesInvoker) Invoke() (*model.BatchDeleteProtectedBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteProtectedBranchesResponse), nil
	}
}

type BatchDeleteProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteProtectedTagsInvoker) Invoke() (*model.BatchDeleteProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteProtectedTagsResponse), nil
	}
}

type BatchUpdateProtectedBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateProtectedBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateProtectedBranchesInvoker) Invoke() (*model.BatchUpdateProtectedBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateProtectedBranchesResponse), nil
	}
}

type BatchUpdateProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateProtectedTagsInvoker) Invoke() (*model.BatchUpdateProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateProtectedTagsResponse), nil
	}
}

type CreateProjectProtectedBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectProtectedBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectProtectedBranchesInvoker) Invoke() (*model.CreateProjectProtectedBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectProtectedBranchesResponse), nil
	}
}

type CreateProjectProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectProtectedTagsInvoker) Invoke() (*model.CreateProjectProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectProtectedTagsResponse), nil
	}
}

type DeleteProtectedBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectedBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectedBranchInvoker) Invoke() (*model.DeleteProtectedBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectedBranchResponse), nil
	}
}

type DeleteProtectedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectedTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectedTagInvoker) Invoke() (*model.DeleteProtectedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectedTagResponse), nil
	}
}

type ListProjectProtectedBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectProtectedBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectProtectedBranchesInvoker) Invoke() (*model.ListProjectProtectedBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectProtectedBranchesResponse), nil
	}
}

type ListProjectProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectProtectedTagsInvoker) Invoke() (*model.ListProjectProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectProtectedTagsResponse), nil
	}
}

type ListProtectedBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedBranchesInvoker) Invoke() (*model.ListProtectedBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedBranchesResponse), nil
	}
}

type ListProtectedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedTagsInvoker) Invoke() (*model.ListProtectedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedTagsResponse), nil
	}
}

type ShowProtectedBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectedBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProtectedBranchInvoker) Invoke() (*model.ShowProtectedBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectedBranchResponse), nil
	}
}

type ShowProtectedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectedTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProtectedTagInvoker) Invoke() (*model.ShowProtectedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectedTagResponse), nil
	}
}

type UpdateProtectedBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProtectedBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProtectedBranchInvoker) Invoke() (*model.UpdateProtectedBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectedBranchResponse), nil
	}
}

type UpdateProtectedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProtectedTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProtectedTagInvoker) Invoke() (*model.UpdateProtectedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectedTagResponse), nil
	}
}

type BatchDeleteBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteBranchInvoker) Invoke() (*model.BatchDeleteBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteBranchResponse), nil
	}
}

type CreateBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBranchInvoker) Invoke() (*model.CreateBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBranchResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBranchInvoker) Invoke() (*model.DeleteBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBranchResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBranchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBranchesInvoker) Invoke() (*model.ListBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBranchesResponse), nil
	}
}

type ListRefsListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRefsListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRefsListInvoker) Invoke() (*model.ListRefsListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRefsListResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ShowBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBranchInvoker) Invoke() (*model.ShowBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBranchResponse), nil
	}
}

type ShowTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagInvoker) Invoke() (*model.ShowTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagResponse), nil
	}
}

type UpdateBranchNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBranchNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBranchNameInvoker) Invoke() (*model.UpdateBranchNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBranchNameResponse), nil
	}
}

type AddSubmoduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSubmoduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSubmoduleInvoker) Invoke() (*model.AddSubmoduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSubmoduleResponse), nil
	}
}

type AddTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddTrustedIpAddressInvoker) Invoke() (*model.AddTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTrustedIpAddressResponse), nil
	}
}

type AssociateRemoteMirrorInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRemoteMirrorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRemoteMirrorInvoker) Invoke() (*model.AssociateRemoteMirrorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRemoteMirrorResponse), nil
	}
}

type AssociateRepositoryUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRepositoryUserGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRepositoryUserGroupInvoker) Invoke() (*model.AssociateRepositoryUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRepositoryUserGroupResponse), nil
	}
}

type BatchValidateRepoNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchValidateRepoNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchValidateRepoNamesInvoker) Invoke() (*model.BatchValidateRepoNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchValidateRepoNamesResponse), nil
	}
}

type CreateDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDirInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDirInvoker) Invoke() (*model.CreateDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDirResponse), nil
	}
}

type CreateRepositoryCommitRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepositoryCommitRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepositoryCommitRuleInvoker) Invoke() (*model.CreateRepositoryCommitRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepositoryCommitRuleResponse), nil
	}
}

type CreateRepositoryLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepositoryLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepositoryLabelInvoker) Invoke() (*model.CreateRepositoryLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepositoryLabelResponse), nil
	}
}

type CreateRepositorySystemLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepositorySystemLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepositorySystemLabelsInvoker) Invoke() (*model.CreateRepositorySystemLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepositorySystemLabelsResponse), nil
	}
}

type DeleteRepositoryLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepositoryLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepositoryLabelInvoker) Invoke() (*model.DeleteRepositoryLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepositoryLabelResponse), nil
	}
}

type DeleteTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrustedIpAddressInvoker) Invoke() (*model.DeleteTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrustedIpAddressResponse), nil
	}
}

type DownloadArchiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadArchiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadArchiveInvoker) Invoke() (*model.DownloadArchiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadArchiveResponse), nil
	}
}

type ExecuteRepositoryStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteRepositoryStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteRepositoryStatisticsInvoker) Invoke() (*model.ExecuteRepositoryStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteRepositoryStatisticsResponse), nil
	}
}

type ListCurrentUserRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCurrentUserRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCurrentUserRepositoriesInvoker) Invoke() (*model.ListCurrentUserRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCurrentUserRepositoriesResponse), nil
	}
}

type ListGroupRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupRepositoriesInvoker) Invoke() (*model.ListGroupRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupRepositoriesResponse), nil
	}
}

type ListPersonalRecentPushEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPersonalRecentPushEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPersonalRecentPushEventsInvoker) Invoke() (*model.ListPersonalRecentPushEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPersonalRecentPushEventsResponse), nil
	}
}

type ListPersonalRepositoryImportRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPersonalRepositoryImportRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPersonalRepositoryImportRecordsInvoker) Invoke() (*model.ListPersonalRepositoryImportRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPersonalRepositoryImportRecordsResponse), nil
	}
}

type ListProjectRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectRepositoriesInvoker) Invoke() (*model.ListProjectRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectRepositoriesResponse), nil
	}
}

type ListRepositoryCommitRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryCommitRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryCommitRulesInvoker) Invoke() (*model.ListRepositoryCommitRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryCommitRulesResponse), nil
	}
}

type ListRepositoryContributorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryContributorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryContributorsInvoker) Invoke() (*model.ListRepositoryContributorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryContributorsResponse), nil
	}
}

type ListRepositoryEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryEventsInvoker) Invoke() (*model.ListRepositoryEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryEventsResponse), nil
	}
}

type ListRepositoryForksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryForksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryForksInvoker) Invoke() (*model.ListRepositoryForksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryForksResponse), nil
	}
}

type ListRepositoryLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryLabelsInvoker) Invoke() (*model.ListRepositoryLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryLabelsResponse), nil
	}
}

type ListRepositoryLanguagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryLanguagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryLanguagesInvoker) Invoke() (*model.ListRepositoryLanguagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryLanguagesResponse), nil
	}
}

type ListRepositoryTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryTemplatesInvoker) Invoke() (*model.ListRepositoryTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryTemplatesResponse), nil
	}
}

type ListSubmodulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubmodulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubmodulesInvoker) Invoke() (*model.ListSubmodulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubmodulesResponse), nil
	}
}

type ListTrustedIpAddressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrustedIpAddressesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrustedIpAddressesInvoker) Invoke() (*model.ListTrustedIpAddressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrustedIpAddressesResponse), nil
	}
}

type LockRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *LockRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LockRepositoryInvoker) Invoke() (*model.LockRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LockRepositoryResponse), nil
	}
}

type RemoveDeployKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveDeployKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveDeployKeyInvoker) Invoke() (*model.RemoveDeployKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveDeployKeyResponse), nil
	}
}

type RemoveDeployKeyFromSubmodulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveDeployKeyFromSubmodulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveDeployKeyFromSubmodulesInvoker) Invoke() (*model.RemoveDeployKeyFromSubmodulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveDeployKeyFromSubmodulesResponse), nil
	}
}

type ShowBlobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlobsInvoker) Invoke() (*model.ShowBlobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlobsResponse), nil
	}
}

type ShowCommitStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitStatisticsInvoker) Invoke() (*model.ShowCommitStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitStatisticsResponse), nil
	}
}

type ShowDiffLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiffLinesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiffLinesInvoker) Invoke() (*model.ShowDiffLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiffLinesResponse), nil
	}
}

type ShowLastPushEventInRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLastPushEventInRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLastPushEventInRepositoryInvoker) Invoke() (*model.ShowLastPushEventInRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLastPushEventInRepositoryResponse), nil
	}
}

type ShowNotificationSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotificationSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotificationSubscriptionInvoker) Invoke() (*model.ShowNotificationSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotificationSubscriptionResponse), nil
	}
}

type ShowNotificationSubscriptionsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotificationSubscriptionsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotificationSubscriptionsStatusInvoker) Invoke() (*model.ShowNotificationSubscriptionsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotificationSubscriptionsStatusResponse), nil
	}
}

type ShowRefCompareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRefCompareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRefCompareInvoker) Invoke() (*model.ShowRefCompareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRefCompareResponse), nil
	}
}

type ShowRemoteMirrorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRemoteMirrorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRemoteMirrorInvoker) Invoke() (*model.ShowRemoteMirrorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRemoteMirrorResponse), nil
	}
}

type ShowRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryInvoker) Invoke() (*model.ShowRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryResponse), nil
	}
}

type ShowRepositoryGeneralCommitRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryGeneralCommitRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryGeneralCommitRuleInvoker) Invoke() (*model.ShowRepositoryGeneralCommitRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryGeneralCommitRuleResponse), nil
	}
}

type ShowRepositoryGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryGeneralPolicyInvoker) Invoke() (*model.ShowRepositoryGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryGeneralPolicyResponse), nil
	}
}

type ShowRepositoryInheritSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryInheritSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryInheritSettingInvoker) Invoke() (*model.ShowRepositoryInheritSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryInheritSettingResponse), nil
	}
}

type ShowRepositoryInheritSettingSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryInheritSettingSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryInheritSettingSourceInvoker) Invoke() (*model.ShowRepositoryInheritSettingSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryInheritSettingSourceResponse), nil
	}
}

type ShowRepositoryStatisticsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryStatisticsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryStatisticsStatusInvoker) Invoke() (*model.ShowRepositoryStatisticsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryStatisticsStatusResponse), nil
	}
}

type ShowRepositoryStatisticsSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryStatisticsSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryStatisticsSummaryInvoker) Invoke() (*model.ShowRepositoryStatisticsSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryStatisticsSummaryResponse), nil
	}
}

type ShowRepositoryWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryWatermarkInvoker) Invoke() (*model.ShowRepositoryWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryWatermarkResponse), nil
	}
}

type ShowUserRefPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserRefPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserRefPermissionInvoker) Invoke() (*model.ShowUserRefPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserRefPermissionResponse), nil
	}
}

type StartHouseKeepingInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartHouseKeepingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartHouseKeepingInvoker) Invoke() (*model.StartHouseKeepingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartHouseKeepingResponse), nil
	}
}

type StartRemoteMirrorSynchronizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartRemoteMirrorSynchronizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartRemoteMirrorSynchronizationInvoker) Invoke() (*model.StartRemoteMirrorSynchronizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartRemoteMirrorSynchronizationResponse), nil
	}
}

type SyncDeployKeyToSubmodulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncDeployKeyToSubmodulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncDeployKeyToSubmodulesInvoker) Invoke() (*model.SyncDeployKeyToSubmodulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncDeployKeyToSubmodulesResponse), nil
	}
}

type UnlockRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnlockRepositoryInvoker) Invoke() (*model.UnlockRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockRepositoryResponse), nil
	}
}

type UpdateNotificationSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotificationSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotificationSubscriptionInvoker) Invoke() (*model.UpdateNotificationSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotificationSubscriptionResponse), nil
	}
}

type UpdateRepositoryCommitRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryCommitRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryCommitRuleInvoker) Invoke() (*model.UpdateRepositoryCommitRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryCommitRuleResponse), nil
	}
}

type UpdateRepositoryGeneralCommitRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryGeneralCommitRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryGeneralCommitRuleInvoker) Invoke() (*model.UpdateRepositoryGeneralCommitRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryGeneralCommitRuleResponse), nil
	}
}

type UpdateRepositoryGeneralPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryGeneralPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryGeneralPolicyInvoker) Invoke() (*model.UpdateRepositoryGeneralPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryGeneralPolicyResponse), nil
	}
}

type UpdateRepositoryInheritSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryInheritSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryInheritSettingInvoker) Invoke() (*model.UpdateRepositoryInheritSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryInheritSettingResponse), nil
	}
}

type UpdateRepositoryLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryLabelInvoker) Invoke() (*model.UpdateRepositoryLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryLabelResponse), nil
	}
}

type UpdateRepositoryRemoteMirrorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryRemoteMirrorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryRemoteMirrorInvoker) Invoke() (*model.UpdateRepositoryRemoteMirrorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryRemoteMirrorResponse), nil
	}
}

type UpdateRepositoryWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryWatermarkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryWatermarkInvoker) Invoke() (*model.UpdateRepositoryWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryWatermarkResponse), nil
	}
}

type UpdateTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTrustedIpAddressInvoker) Invoke() (*model.UpdateTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrustedIpAddressResponse), nil
	}
}

type AddTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTenantTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddTenantTrustedIpAddressInvoker) Invoke() (*model.AddTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTenantTrustedIpAddressResponse), nil
	}
}

type DeleteTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTenantTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTenantTrustedIpAddressInvoker) Invoke() (*model.DeleteTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTenantTrustedIpAddressResponse), nil
	}
}

type ExportTenantRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTenantRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTenantRepositoriesInvoker) Invoke() (*model.ExportTenantRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTenantRepositoriesResponse), nil
	}
}

type ListTenantRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantRepositoriesInvoker) Invoke() (*model.ListTenantRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantRepositoriesResponse), nil
	}
}

type ListTenantTrustedIpAddressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantTrustedIpAddressesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantTrustedIpAddressesInvoker) Invoke() (*model.ListTenantTrustedIpAddressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantTrustedIpAddressesResponse), nil
	}
}

type UpdateTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantTrustedIpAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTenantTrustedIpAddressInvoker) Invoke() (*model.UpdateTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantTrustedIpAddressResponse), nil
	}
}

type CheckDeployKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDeployKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDeployKeyInvoker) Invoke() (*model.CheckDeployKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDeployKeyResponse), nil
	}
}

type CheckGroupDeployKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckGroupDeployKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckGroupDeployKeyInvoker) Invoke() (*model.CheckGroupDeployKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckGroupDeployKeyResponse), nil
	}
}

type ListBranchRelatedWorkItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBranchRelatedWorkItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBranchRelatedWorkItemsInvoker) Invoke() (*model.ListBranchRelatedWorkItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBranchRelatedWorkItemsResponse), nil
	}
}

type ListGroupDeployKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupDeployKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupDeployKeysInvoker) Invoke() (*model.ListGroupDeployKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupDeployKeysResponse), nil
	}
}

type ListProjectDeployKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectDeployKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectDeployKeysInvoker) Invoke() (*model.ListProjectDeployKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectDeployKeysResponse), nil
	}
}

type ListRepositoryDeployKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryDeployKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryDeployKeysInvoker) Invoke() (*model.ListRepositoryDeployKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryDeployKeysResponse), nil
	}
}

type ListRepositoryWorkItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryWorkItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryWorkItemsInvoker) Invoke() (*model.ListRepositoryWorkItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryWorkItemsResponse), nil
	}
}

type ShowGroupE2eSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupE2eSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupE2eSettingInvoker) Invoke() (*model.ShowGroupE2eSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupE2eSettingResponse), nil
	}
}

type ShowProjectE2eSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectE2eSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectE2eSettingInvoker) Invoke() (*model.ShowProjectE2eSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectE2eSettingResponse), nil
	}
}

type ShowRepositoryE2eSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryE2eSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryE2eSettingInvoker) Invoke() (*model.ShowRepositoryE2eSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryE2eSettingResponse), nil
	}
}

type AddSshKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSshKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSshKeyInvoker) Invoke() (*model.AddSshKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSshKeyResponse), nil
	}
}

type BatchValidateUserGroupPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchValidateUserGroupPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchValidateUserGroupPermissionsInvoker) Invoke() (*model.BatchValidateUserGroupPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchValidateUserGroupPermissionsResponse), nil
	}
}

type DeleteSshKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSshKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSshKeyInvoker) Invoke() (*model.DeleteSshKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSshKeyResponse), nil
	}
}

type ListImpersonationTokensInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImpersonationTokensInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImpersonationTokensInvoker) Invoke() (*model.ListImpersonationTokensResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImpersonationTokensResponse), nil
	}
}

type ListUserGpgKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserGpgKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserGpgKeysInvoker) Invoke() (*model.ListUserGpgKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserGpgKeysResponse), nil
	}
}

type ListUserKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserKeysInvoker) Invoke() (*model.ListUserKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserKeysResponse), nil
	}
}

type SendUserEmailVerifyCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendUserEmailVerifyCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendUserEmailVerifyCodeInvoker) Invoke() (*model.SendUserEmailVerifyCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendUserEmailVerifyCodeResponse), nil
	}
}

type ShowHttpsPasswordSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpsPasswordSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpsPasswordSettingInvoker) Invoke() (*model.ShowHttpsPasswordSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpsPasswordSettingResponse), nil
	}
}

type ShowUserEmailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserEmailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserEmailsInvoker) Invoke() (*model.ShowUserEmailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserEmailsResponse), nil
	}
}

type UpdateHttpsPasswordSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpsPasswordSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpsPasswordSettingInvoker) Invoke() (*model.UpdateHttpsPasswordSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpsPasswordSettingResponse), nil
	}
}

type UpdateUserEmailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserEmailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserEmailsInvoker) Invoke() (*model.UpdateUserEmailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserEmailsResponse), nil
	}
}

type AddGroupWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGroupWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddGroupWebhookInvoker) Invoke() (*model.AddGroupWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGroupWebhookResponse), nil
	}
}

type AddProjectWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddProjectWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddProjectWebhookInvoker) Invoke() (*model.AddProjectWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProjectWebhookResponse), nil
	}
}

type AddRepositoryWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRepositoryWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRepositoryWebhookInvoker) Invoke() (*model.AddRepositoryWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRepositoryWebhookResponse), nil
	}
}

type ListGroupWebhookLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupWebhookLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupWebhookLogsInvoker) Invoke() (*model.ListGroupWebhookLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupWebhookLogsResponse), nil
	}
}

type ListGroupWebhooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupWebhooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupWebhooksInvoker) Invoke() (*model.ListGroupWebhooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupWebhooksResponse), nil
	}
}

type ListProjectWebhookLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectWebhookLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectWebhookLogsInvoker) Invoke() (*model.ListProjectWebhookLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectWebhookLogsResponse), nil
	}
}

type ListProjectWebhooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectWebhooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectWebhooksInvoker) Invoke() (*model.ListProjectWebhooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectWebhooksResponse), nil
	}
}

type ListRepositoryWebhookLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryWebhookLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryWebhookLogsInvoker) Invoke() (*model.ListRepositoryWebhookLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryWebhookLogsResponse), nil
	}
}

type ListRepositoryWebhooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryWebhooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryWebhooksInvoker) Invoke() (*model.ListRepositoryWebhooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryWebhooksResponse), nil
	}
}

type RemoveGroupWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveGroupWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveGroupWebhookInvoker) Invoke() (*model.RemoveGroupWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveGroupWebhookResponse), nil
	}
}

type RemoveProjectWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveProjectWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveProjectWebhookInvoker) Invoke() (*model.RemoveProjectWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveProjectWebhookResponse), nil
	}
}

type RemoveRepositoryWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveRepositoryWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveRepositoryWebhookInvoker) Invoke() (*model.RemoveRepositoryWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveRepositoryWebhookResponse), nil
	}
}

type ShowGroupWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupWebhookInvoker) Invoke() (*model.ShowGroupWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupWebhookResponse), nil
	}
}

type ShowGroupWebhookLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupWebhookLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupWebhookLogInvoker) Invoke() (*model.ShowGroupWebhookLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupWebhookLogResponse), nil
	}
}

type ShowProjectWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectWebhookInvoker) Invoke() (*model.ShowProjectWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectWebhookResponse), nil
	}
}

type ShowProjectWebhookLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectWebhookLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectWebhookLogInvoker) Invoke() (*model.ShowProjectWebhookLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectWebhookLogResponse), nil
	}
}

type ShowRepositoryWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryWebhookInvoker) Invoke() (*model.ShowRepositoryWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryWebhookResponse), nil
	}
}

type ShowRepositoryWebhookLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryWebhookLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryWebhookLogInvoker) Invoke() (*model.ShowRepositoryWebhookLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryWebhookLogResponse), nil
	}
}

type UpdateGroupWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupWebhookInvoker) Invoke() (*model.UpdateGroupWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupWebhookResponse), nil
	}
}

type UpdateProjectWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectWebhookInvoker) Invoke() (*model.UpdateProjectWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectWebhookResponse), nil
	}
}

type UpdateRepositoryWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepositoryWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepositoryWebhookInvoker) Invoke() (*model.UpdateRepositoryWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepositoryWebhookResponse), nil
	}
}
