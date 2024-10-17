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
