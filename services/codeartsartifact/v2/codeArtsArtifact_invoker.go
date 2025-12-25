package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsartifact/v2/model"
)

type BatchDeleteTrashesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTrashesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTrashesInvoker) Invoke() (*model.BatchDeleteTrashesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTrashesResponse), nil
	}
}

type BatchRestoreRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRestoreRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRestoreRepoInvoker) Invoke() (*model.BatchRestoreRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRestoreRepoResponse), nil
	}
}

type CreateArtifactoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateArtifactoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateArtifactoryInvoker) Invoke() (*model.CreateArtifactoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateArtifactoryResponse), nil
	}
}

type CreateAttentionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAttentionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAttentionInvoker) Invoke() (*model.CreateAttentionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAttentionResponse), nil
	}
}

type CreateDockerRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDockerRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDockerRepositoriesInvoker) Invoke() (*model.CreateDockerRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDockerRepositoriesResponse), nil
	}
}

type CreateMavenRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMavenRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMavenRepoInvoker) Invoke() (*model.CreateMavenRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMavenRepoResponse), nil
	}
}

type CreateProjectRelatedRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectRelatedRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectRelatedRepositoryInvoker) Invoke() (*model.CreateProjectRelatedRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectRelatedRepositoryResponse), nil
	}
}

type DeleteArtifactFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteArtifactFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteArtifactFileInvoker) Invoke() (*model.DeleteArtifactFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteArtifactFileResponse), nil
	}
}

type DeleteCompletelyUpdateFileStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCompletelyUpdateFileStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCompletelyUpdateFileStateInvoker) Invoke() (*model.DeleteCompletelyUpdateFileStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCompletelyUpdateFileStateResponse), nil
	}
}

type DeleteRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepositoryInvoker) Invoke() (*model.DeleteRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepositoryResponse), nil
	}
}

type ListAllRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllRepositoriesInvoker) Invoke() (*model.ListAllRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllRepositoriesResponse), nil
	}
}

type ListArtifactoryComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArtifactoryComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListArtifactoryComponentInvoker) Invoke() (*model.ListArtifactoryComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArtifactoryComponentResponse), nil
	}
}

type ListArtifactoryStorageStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArtifactoryStorageStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListArtifactoryStorageStatisticInvoker) Invoke() (*model.ListArtifactoryStorageStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArtifactoryStorageStatisticResponse), nil
	}
}

type ListAttentionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttentionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttentionsInvoker) Invoke() (*model.ListAttentionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttentionsResponse), nil
	}
}

type ListCapacityMessageSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCapacityMessageSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCapacityMessageSettingInvoker) Invoke() (*model.ListCapacityMessageSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCapacityMessageSettingResponse), nil
	}
}

type ListChildProxyRepositoriesListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChildProxyRepositoriesListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListChildProxyRepositoriesListInvoker) Invoke() (*model.ListChildProxyRepositoriesListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChildProxyRepositoriesListResponse), nil
	}
}

type ListDomainIpConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainIpConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainIpConfigInvoker) Invoke() (*model.ListDomainIpConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainIpConfigResponse), nil
	}
}

type ListFileBuildArchivesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileBuildArchivesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileBuildArchivesInvoker) Invoke() (*model.ListFileBuildArchivesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileBuildArchivesResponse), nil
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

type ListLatestVersionFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLatestVersionFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLatestVersionFilesInvoker) Invoke() (*model.ListLatestVersionFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLatestVersionFilesResponse), nil
	}
}

type ListMavenListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMavenListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMavenListInvoker) Invoke() (*model.ListMavenListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMavenListResponse), nil
	}
}

type ListMavenUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMavenUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMavenUserInvoker) Invoke() (*model.ListMavenUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMavenUserResponse), nil
	}
}

type ListNetProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetProxyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetProxyInvoker) Invoke() (*model.ListNetProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetProxyResponse), nil
	}
}

type ListProjectRolePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectRolePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectRolePermissionsInvoker) Invoke() (*model.ListProjectRolePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectRolePermissionsResponse), nil
	}
}

type ListProjectUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectUsersInvoker) Invoke() (*model.ListProjectUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectUsersResponse), nil
	}
}

type ListSecGuardListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecGuardListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecGuardListInvoker) Invoke() (*model.ListSecGuardListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecGuardListResponse), nil
	}
}

type ModifyRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyRepositoryInvoker) Invoke() (*model.ModifyRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyRepositoryResponse), nil
	}
}

type ResetUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetUserPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetUserPasswordInvoker) Invoke() (*model.ResetUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetUserPasswordResponse), nil
	}
}

type SearchArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchArtifactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchArtifactsInvoker) Invoke() (*model.SearchArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchArtifactsResponse), nil
	}
}

type SearchByChecksumInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchByChecksumInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchByChecksumInvoker) Invoke() (*model.SearchByChecksumResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchByChecksumResponse), nil
	}
}

type ShowAuditInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditInvoker) Invoke() (*model.ShowAuditResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditResponse), nil
	}
}

type ShowAutoDeleteJobSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoDeleteJobSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoDeleteJobSettingsInvoker) Invoke() (*model.ShowAutoDeleteJobSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoDeleteJobSettingsResponse), nil
	}
}

type ShowDomainReleaseRepoStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainReleaseRepoStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainReleaseRepoStorageInvoker) Invoke() (*model.ShowDomainReleaseRepoStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainReleaseRepoStorageResponse), nil
	}
}

type ShowFileDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileDetailInvoker) Invoke() (*model.ShowFileDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileDetailResponse), nil
	}
}

type ShowFileDetailByFullNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileDetailByFullNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileDetailByFullNameInvoker) Invoke() (*model.ShowFileDetailByFullNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileDetailByFullNameResponse), nil
	}
}

type ShowFileTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileTreeInvoker) Invoke() (*model.ShowFileTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileTreeResponse), nil
	}
}

type ShowLatestVersionFilesCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestVersionFilesCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestVersionFilesCountInvoker) Invoke() (*model.ShowLatestVersionFilesCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestVersionFilesCountResponse), nil
	}
}

type ShowMavenInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMavenInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMavenInfoInvoker) Invoke() (*model.ShowMavenInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMavenInfoResponse), nil
	}
}

type ShowMultiRolesUserPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiRolesUserPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiRolesUserPermissionsInvoker) Invoke() (*model.ShowMultiRolesUserPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiRolesUserPermissionsResponse), nil
	}
}

type ShowOpenSourceEnabledInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOpenSourceEnabledInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOpenSourceEnabledInvoker) Invoke() (*model.ShowOpenSourceEnabledResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOpenSourceEnabledResponse), nil
	}
}

type ShowPackageDataDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPackageDataDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPackageDataDetailInvoker) Invoke() (*model.ShowPackageDataDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPackageDataDetailResponse), nil
	}
}

type ShowPackageInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPackageInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPackageInfoInvoker) Invoke() (*model.ShowPackageInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPackageInfoResponse), nil
	}
}

type ShowProjectListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectListInvoker) Invoke() (*model.ShowProjectListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectListResponse), nil
	}
}

type ShowProjectReleaseFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectReleaseFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectReleaseFilesInvoker) Invoke() (*model.ShowProjectReleaseFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectReleaseFilesResponse), nil
	}
}

type ShowProjectStorageInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectStorageInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectStorageInfoInvoker) Invoke() (*model.ShowProjectStorageInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectStorageInfoResponse), nil
	}
}

type ShowReleaseProjectFilesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowReleaseProjectFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowReleaseProjectFilesInvoker) Invoke() (*model.ShowReleaseProjectFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReleaseProjectFilesResponse), nil
	}
}

type ShowRepoUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepoUserInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepoUserInfoInvoker) Invoke() (*model.ShowRepoUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepoUserInfoResponse), nil
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

type ShowRepositoryInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryInfoInvoker) Invoke() (*model.ShowRepositoryInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryInfoResponse), nil
	}
}

type ShowStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStorageInvoker) Invoke() (*model.ShowStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStorageResponse), nil
	}
}

type ShowUserPrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserPrivilegesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserPrivilegesInvoker) Invoke() (*model.ShowUserPrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserPrivilegesResponse), nil
	}
}

type ShowUserTicketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserTicketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserTicketInvoker) Invoke() (*model.ShowUserTicketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserTicketResponse), nil
	}
}

type UpdateArtifactoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateArtifactoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateArtifactoryInvoker) Invoke() (*model.UpdateArtifactoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateArtifactoryResponse), nil
	}
}
