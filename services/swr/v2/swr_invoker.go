package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

type CreateImageSyncRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageSyncRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageSyncRepoInvoker) Invoke() (*model.CreateImageSyncRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageSyncRepoResponse), nil
	}
}

type CreateManualImageSyncRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManualImageSyncRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateManualImageSyncRepoInvoker) Invoke() (*model.CreateManualImageSyncRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManualImageSyncRepoResponse), nil
	}
}

type CreateNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNamespaceInvoker) Invoke() (*model.CreateNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNamespaceResponse), nil
	}
}

type CreateNamespaceAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNamespaceAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNamespaceAuthInvoker) Invoke() (*model.CreateNamespaceAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNamespaceAuthResponse), nil
	}
}

type CreateRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepoInvoker) Invoke() (*model.CreateRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepoResponse), nil
	}
}

type CreateRepoDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepoDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepoDomainsInvoker) Invoke() (*model.CreateRepoDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepoDomainsResponse), nil
	}
}

type CreateRetentionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRetentionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRetentionInvoker) Invoke() (*model.CreateRetentionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRetentionResponse), nil
	}
}

type CreateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretInvoker) Invoke() (*model.CreateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretResponse), nil
	}
}

type CreateTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTriggerInvoker) Invoke() (*model.CreateTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTriggerResponse), nil
	}
}

type CreateUserRepositoryAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserRepositoryAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserRepositoryAuthInvoker) Invoke() (*model.CreateUserRepositoryAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserRepositoryAuthResponse), nil
	}
}

type DeleteImageSyncRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageSyncRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageSyncRepoInvoker) Invoke() (*model.DeleteImageSyncRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageSyncRepoResponse), nil
	}
}

type DeleteNamespaceAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNamespaceAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNamespaceAuthInvoker) Invoke() (*model.DeleteNamespaceAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNamespaceAuthResponse), nil
	}
}

type DeleteNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNamespacesInvoker) Invoke() (*model.DeleteNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNamespacesResponse), nil
	}
}

type DeleteRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepoInvoker) Invoke() (*model.DeleteRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepoResponse), nil
	}
}

type DeleteRepoDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepoDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepoDomainsInvoker) Invoke() (*model.DeleteRepoDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepoDomainsResponse), nil
	}
}

type DeleteRepoTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepoTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepoTagInvoker) Invoke() (*model.DeleteRepoTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepoTagResponse), nil
	}
}

type DeleteRetentionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRetentionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRetentionInvoker) Invoke() (*model.DeleteRetentionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRetentionResponse), nil
	}
}

type DeleteTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTriggerInvoker) Invoke() (*model.DeleteTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTriggerResponse), nil
	}
}

type DeleteUserRepositoryAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserRepositoryAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserRepositoryAuthInvoker) Invoke() (*model.DeleteUserRepositoryAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserRepositoryAuthResponse), nil
	}
}

type ListImageAutoSyncReposDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageAutoSyncReposDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageAutoSyncReposDetailsInvoker) Invoke() (*model.ListImageAutoSyncReposDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageAutoSyncReposDetailsResponse), nil
	}
}

type ListNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNamespacesInvoker) Invoke() (*model.ListNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNamespacesResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListRepoDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepoDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepoDomainsInvoker) Invoke() (*model.ListRepoDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepoDomainsResponse), nil
	}
}

type ListReposDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReposDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReposDetailsInvoker) Invoke() (*model.ListReposDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReposDetailsResponse), nil
	}
}

type ListRepositoryTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryTagsInvoker) Invoke() (*model.ListRepositoryTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryTagsResponse), nil
	}
}

type ListRetentionHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRetentionHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRetentionHistoriesInvoker) Invoke() (*model.ListRetentionHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRetentionHistoriesResponse), nil
	}
}

type ListRetentionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRetentionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRetentionsInvoker) Invoke() (*model.ListRetentionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRetentionsResponse), nil
	}
}

type ListSharedReposDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedReposDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharedReposDetailsInvoker) Invoke() (*model.ListSharedReposDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedReposDetailsResponse), nil
	}
}

type ListTriggersDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTriggersDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTriggersDetailsInvoker) Invoke() (*model.ListTriggersDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTriggersDetailsResponse), nil
	}
}

type ShowAccessDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessDomainInvoker) Invoke() (*model.ShowAccessDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessDomainResponse), nil
	}
}

type ShowNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNamespaceInvoker) Invoke() (*model.ShowNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNamespaceResponse), nil
	}
}

type ShowNamespaceAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNamespaceAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNamespaceAuthInvoker) Invoke() (*model.ShowNamespaceAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNamespaceAuthResponse), nil
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

type ShowRetentionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRetentionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRetentionInvoker) Invoke() (*model.ShowRetentionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRetentionResponse), nil
	}
}

type ShowSyncJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSyncJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSyncJobInvoker) Invoke() (*model.ShowSyncJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSyncJobResponse), nil
	}
}

type ShowTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTriggerInvoker) Invoke() (*model.ShowTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTriggerResponse), nil
	}
}

type ShowUserRepositoryAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserRepositoryAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserRepositoryAuthInvoker) Invoke() (*model.ShowUserRepositoryAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserRepositoryAuthResponse), nil
	}
}

type UpdateNamespaceAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNamespaceAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNamespaceAuthInvoker) Invoke() (*model.UpdateNamespaceAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNamespaceAuthResponse), nil
	}
}

type UpdateRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepoInvoker) Invoke() (*model.UpdateRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepoResponse), nil
	}
}

type UpdateRepoDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepoDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepoDomainsInvoker) Invoke() (*model.UpdateRepoDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepoDomainsResponse), nil
	}
}

type UpdateRetentionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRetentionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRetentionInvoker) Invoke() (*model.UpdateRetentionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRetentionResponse), nil
	}
}

type UpdateTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTriggerInvoker) Invoke() (*model.UpdateTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTriggerResponse), nil
	}
}

type UpdateUserRepositoryAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserRepositoryAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserRepositoryAuthInvoker) Invoke() (*model.UpdateUserRepositoryAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserRepositoryAuthResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}
