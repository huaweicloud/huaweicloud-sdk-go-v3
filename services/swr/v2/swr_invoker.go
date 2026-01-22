package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

type CheckAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckAgencyInvoker) Invoke() (*model.CheckAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckAgencyResponse), nil
	}
}

type CreateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyInvoker) Invoke() (*model.CreateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyResponse), nil
	}
}

type CreateAuthorizationTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorizationTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorizationTokenInvoker) Invoke() (*model.CreateAuthorizationTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorizationTokenResponse), nil
	}
}

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

type CreateRepoTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepoTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepoTagInvoker) Invoke() (*model.CreateRepoTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepoTagResponse), nil
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

type ListReferencesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReferencesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReferencesInvoker) Invoke() (*model.ListReferencesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReferencesResponse), nil
	}
}

type ListRepoAccessoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepoAccessoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepoAccessoriesInvoker) Invoke() (*model.ListRepoAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepoAccessoriesResponse), nil
	}
}

type ListRepoDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepoDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepoDetailsInvoker) Invoke() (*model.ListRepoDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepoDetailsResponse), nil
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

type ListRepositoryTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryTagInvoker) Invoke() (*model.ListRepositoryTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryTagResponse), nil
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

type ListSharedRepoDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedRepoDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharedRepoDetailsInvoker) Invoke() (*model.ListSharedRepoDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedRepoDetailsResponse), nil
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

type ListSyncRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSyncRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSyncRegionsInvoker) Invoke() (*model.ListSyncRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSyncRegionsResponse), nil
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

type ShowDomainOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainOverviewInvoker) Invoke() (*model.ShowDomainOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainOverviewResponse), nil
	}
}

type ShowDomainResourceReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainResourceReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainResourceReportsInvoker) Invoke() (*model.ShowDomainResourceReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainResourceReportsResponse), nil
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

type ShowRepoTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepoTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepoTagInvoker) Invoke() (*model.ShowRepoTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepoTagResponse), nil
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

type ShowShareFeatureGatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShareFeatureGatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShareFeatureGatesInvoker) Invoke() (*model.ShowShareFeatureGatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShareFeatureGatesResponse), nil
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

type AddDomainNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDomainNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDomainNameInvoker) Invoke() (*model.AddDomainNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDomainNameResponse), nil
	}
}

type CreateImmutableRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImmutableRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImmutableRuleInvoker) Invoke() (*model.CreateImmutableRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImmutableRuleResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateInstanceEndpointPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceEndpointPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceEndpointPolicyInvoker) Invoke() (*model.CreateInstanceEndpointPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceEndpointPolicyResponse), nil
	}
}

type CreateInstanceInternalEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInternalEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInternalEndpointInvoker) Invoke() (*model.CreateInstanceInternalEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceInternalEndpointResponse), nil
	}
}

type CreateInstanceLtCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceLtCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceLtCredentialInvoker) Invoke() (*model.CreateInstanceLtCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceLtCredentialResponse), nil
	}
}

type CreateInstanceNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceNamespaceInvoker) Invoke() (*model.CreateInstanceNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceNamespaceResponse), nil
	}
}

type CreateInstanceRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceRegistryInvoker) Invoke() (*model.CreateInstanceRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceRegistryResponse), nil
	}
}

type CreateInstanceReplicationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceReplicationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceReplicationPolicyInvoker) Invoke() (*model.CreateInstanceReplicationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceReplicationPolicyResponse), nil
	}
}

type CreateInstanceResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceResourceTagsInvoker) Invoke() (*model.CreateInstanceResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResourceTagsResponse), nil
	}
}

type CreateInstanceRetentionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceRetentionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceRetentionPolicyInvoker) Invoke() (*model.CreateInstanceRetentionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceRetentionPolicyResponse), nil
	}
}

type CreateInstanceSignPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceSignPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceSignPolicyInvoker) Invoke() (*model.CreateInstanceSignPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceSignPolicyResponse), nil
	}
}

type CreateInstanceTempCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceTempCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceTempCredentialInvoker) Invoke() (*model.CreateInstanceTempCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceTempCredentialResponse), nil
	}
}

type CreateInstanceWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceWebhookInvoker) Invoke() (*model.CreateInstanceWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceWebhookResponse), nil
	}
}

type CreateSubResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubResourceTagsInvoker) Invoke() (*model.CreateSubResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubResourceTagsResponse), nil
	}
}

type DeleteDomainNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainNameInvoker) Invoke() (*model.DeleteDomainNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainNameResponse), nil
	}
}

type DeleteImmutableRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImmutableRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImmutableRuleInvoker) Invoke() (*model.DeleteImmutableRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImmutableRuleResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteInstanceArtifactInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceArtifactInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceArtifactInvoker) Invoke() (*model.DeleteInstanceArtifactResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceArtifactResponse), nil
	}
}

type DeleteInstanceInternalEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInternalEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInternalEndpointInvoker) Invoke() (*model.DeleteInstanceInternalEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceInternalEndpointResponse), nil
	}
}

type DeleteInstanceJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceJobInvoker) Invoke() (*model.DeleteInstanceJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceJobResponse), nil
	}
}

type DeleteInstanceLtCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceLtCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceLtCredentialInvoker) Invoke() (*model.DeleteInstanceLtCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceLtCredentialResponse), nil
	}
}

type DeleteInstanceNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceNamespaceInvoker) Invoke() (*model.DeleteInstanceNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceNamespaceResponse), nil
	}
}

type DeleteInstanceRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceRegistryInvoker) Invoke() (*model.DeleteInstanceRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceRegistryResponse), nil
	}
}

type DeleteInstanceReplicationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceReplicationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceReplicationPolicyInvoker) Invoke() (*model.DeleteInstanceReplicationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceReplicationPolicyResponse), nil
	}
}

type DeleteInstanceRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceRepositoryInvoker) Invoke() (*model.DeleteInstanceRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceRepositoryResponse), nil
	}
}

type DeleteInstanceResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceResourceTagsInvoker) Invoke() (*model.DeleteInstanceResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResourceTagsResponse), nil
	}
}

type DeleteInstanceRetentionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceRetentionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceRetentionPolicyInvoker) Invoke() (*model.DeleteInstanceRetentionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceRetentionPolicyResponse), nil
	}
}

type DeleteInstanceSignPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceSignPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceSignPolicyInvoker) Invoke() (*model.DeleteInstanceSignPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceSignPolicyResponse), nil
	}
}

type DeleteInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceTagInvoker) Invoke() (*model.DeleteInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceTagResponse), nil
	}
}

type DeleteInstanceWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceWebhookInvoker) Invoke() (*model.DeleteInstanceWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceWebhookResponse), nil
	}
}

type DeleteSubResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubResourceTagsInvoker) Invoke() (*model.DeleteSubResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubResourceTagsResponse), nil
	}
}

type ExecuteInstanceReplicationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteInstanceReplicationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteInstanceReplicationPolicyInvoker) Invoke() (*model.ExecuteInstanceReplicationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteInstanceReplicationPolicyResponse), nil
	}
}

type ExecuteInstanceRetentionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteInstanceRetentionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteInstanceRetentionPolicyInvoker) Invoke() (*model.ExecuteInstanceRetentionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteInstanceRetentionPolicyResponse), nil
	}
}

type ExecuteInstanceSignPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteInstanceSignPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteInstanceSignPolicyInvoker) Invoke() (*model.ExecuteInstanceSignPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteInstanceSignPolicyResponse), nil
	}
}

type ListAuditLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuditLogsInvoker) Invoke() (*model.ListAuditLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditLogsResponse), nil
	}
}

type ListDomainNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainNamesInvoker) Invoke() (*model.ListDomainNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainNamesResponse), nil
	}
}

type ListFeatureGatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeatureGatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeatureGatesInvoker) Invoke() (*model.ListFeatureGatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeatureGatesResponse), nil
	}
}

type ListGlobalFeatureGatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalFeatureGatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalFeatureGatesInvoker) Invoke() (*model.ListGlobalFeatureGatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalFeatureGatesResponse), nil
	}
}

type ListImmutableRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImmutableRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImmutableRulesInvoker) Invoke() (*model.ListImmutableRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImmutableRulesResponse), nil
	}
}

type ListInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceInvoker) Invoke() (*model.ListInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResponse), nil
	}
}

type ListInstanceAccessoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceAccessoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceAccessoriesInvoker) Invoke() (*model.ListInstanceAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceAccessoriesResponse), nil
	}
}

type ListInstanceAllArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceAllArtifactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceAllArtifactsInvoker) Invoke() (*model.ListInstanceAllArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceAllArtifactsResponse), nil
	}
}

type ListInstanceArtifactVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceArtifactVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceArtifactVulnerabilitiesInvoker) Invoke() (*model.ListInstanceArtifactVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceArtifactVulnerabilitiesResponse), nil
	}
}

type ListInstanceArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceArtifactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceArtifactsInvoker) Invoke() (*model.ListInstanceArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceArtifactsResponse), nil
	}
}

type ListInstanceInternalEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInternalEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceInternalEndpointsInvoker) Invoke() (*model.ListInstanceInternalEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceInternalEndpointsResponse), nil
	}
}

type ListInstanceJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceJobsInvoker) Invoke() (*model.ListInstanceJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceJobsResponse), nil
	}
}

type ListInstanceLtCredentialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceLtCredentialsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceLtCredentialsInvoker) Invoke() (*model.ListInstanceLtCredentialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceLtCredentialsResponse), nil
	}
}

type ListInstanceNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceNamespacesInvoker) Invoke() (*model.ListInstanceNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceNamespacesResponse), nil
	}
}

type ListInstanceProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceProjectTagsInvoker) Invoke() (*model.ListInstanceProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceProjectTagsResponse), nil
	}
}

type ListInstanceRegistriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRegistriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRegistriesInvoker) Invoke() (*model.ListInstanceRegistriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRegistriesResponse), nil
	}
}

type ListInstanceReplicationPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceReplicationPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceReplicationPoliciesInvoker) Invoke() (*model.ListInstanceReplicationPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceReplicationPoliciesResponse), nil
	}
}

type ListInstanceReplicationPolicyExecSubTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecSubTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecSubTasksInvoker) Invoke() (*model.ListInstanceReplicationPolicyExecSubTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceReplicationPolicyExecSubTasksResponse), nil
	}
}

type ListInstanceReplicationPolicyExecTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecTasksInvoker) Invoke() (*model.ListInstanceReplicationPolicyExecTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceReplicationPolicyExecTasksResponse), nil
	}
}

type ListInstanceReplicationPolicyExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceReplicationPolicyExecutionsInvoker) Invoke() (*model.ListInstanceReplicationPolicyExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceReplicationPolicyExecutionsResponse), nil
	}
}

type ListInstanceRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRepositoriesInvoker) Invoke() (*model.ListInstanceRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRepositoriesResponse), nil
	}
}

type ListInstanceResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceResourceInstancesInvoker) Invoke() (*model.ListInstanceResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResourceInstancesResponse), nil
	}
}

type ListInstanceResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceResourceTagsInvoker) Invoke() (*model.ListInstanceResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResourceTagsResponse), nil
	}
}

type ListInstanceRetentionPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRetentionPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRetentionPoliciesInvoker) Invoke() (*model.ListInstanceRetentionPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRetentionPoliciesResponse), nil
	}
}

type ListInstanceRetentionPolicyExecSubTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecSubTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecSubTasksInvoker) Invoke() (*model.ListInstanceRetentionPolicyExecSubTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRetentionPolicyExecSubTasksResponse), nil
	}
}

type ListInstanceRetentionPolicyExecTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecTasksInvoker) Invoke() (*model.ListInstanceRetentionPolicyExecTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRetentionPolicyExecTasksResponse), nil
	}
}

type ListInstanceRetentionPolicyExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceRetentionPolicyExecutionsInvoker) Invoke() (*model.ListInstanceRetentionPolicyExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceRetentionPolicyExecutionsResponse), nil
	}
}

type ListInstanceSignPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceSignPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceSignPoliciesInvoker) Invoke() (*model.ListInstanceSignPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceSignPoliciesResponse), nil
	}
}

type ListInstanceSignPolicyExecTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceSignPolicyExecTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceSignPolicyExecTasksInvoker) Invoke() (*model.ListInstanceSignPolicyExecTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceSignPolicyExecTasksResponse), nil
	}
}

type ListInstanceSignPolicyExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceSignPolicyExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceSignPolicyExecutionsInvoker) Invoke() (*model.ListInstanceSignPolicyExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceSignPolicyExecutionsResponse), nil
	}
}

type ListInstanceSignatureExecutionSubtasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceSignatureExecutionSubtasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceSignatureExecutionSubtasksInvoker) Invoke() (*model.ListInstanceSignatureExecutionSubtasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceSignatureExecutionSubtasksResponse), nil
	}
}

type ListInstanceStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceStatisticsInvoker) Invoke() (*model.ListInstanceStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceStatisticsResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListInstanceWebhookJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceWebhookJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceWebhookJobsInvoker) Invoke() (*model.ListInstanceWebhookJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceWebhookJobsResponse), nil
	}
}

type ListInstanceWebhooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceWebhooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceWebhooksInvoker) Invoke() (*model.ListInstanceWebhooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceWebhooksResponse), nil
	}
}

type ListNamespaceRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNamespaceRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNamespaceRepositoriesInvoker) Invoke() (*model.ListNamespaceRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNamespaceRepositoriesResponse), nil
	}
}

type ListNamespaceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNamespaceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNamespaceTagsInvoker) Invoke() (*model.ListNamespaceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNamespaceTagsResponse), nil
	}
}

type ListSubResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubResourceInstancesInvoker) Invoke() (*model.ListSubResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubResourceInstancesResponse), nil
	}
}

type ListSubResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubResourceTagsInvoker) Invoke() (*model.ListSubResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubResourceTagsResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowInstanceArtifactInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceArtifactInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceArtifactInvoker) Invoke() (*model.ShowInstanceArtifactResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceArtifactResponse), nil
	}
}

type ShowInstanceArtifactAdditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceArtifactAdditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceArtifactAdditionInvoker) Invoke() (*model.ShowInstanceArtifactAdditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceArtifactAdditionResponse), nil
	}
}

type ShowInstanceConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceConfigurationInvoker) Invoke() (*model.ShowInstanceConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceConfigurationResponse), nil
	}
}

type ShowInstanceEndpointPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceEndpointPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceEndpointPolicyInvoker) Invoke() (*model.ShowInstanceEndpointPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceEndpointPolicyResponse), nil
	}
}

type ShowInstanceInternalEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInternalEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInternalEndpointInvoker) Invoke() (*model.ShowInstanceInternalEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceInternalEndpointResponse), nil
	}
}

type ShowInstanceJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceJobInvoker) Invoke() (*model.ShowInstanceJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceJobResponse), nil
	}
}

type ShowInstanceNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceNamespaceInvoker) Invoke() (*model.ShowInstanceNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceNamespaceResponse), nil
	}
}

type ShowInstanceRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceRegistryInvoker) Invoke() (*model.ShowInstanceRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceRegistryResponse), nil
	}
}

type ShowInstanceReplicationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceReplicationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceReplicationPolicyInvoker) Invoke() (*model.ShowInstanceReplicationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceReplicationPolicyResponse), nil
	}
}

type ShowInstanceRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceRepositoryInvoker) Invoke() (*model.ShowInstanceRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceRepositoryResponse), nil
	}
}

type ShowInstanceResourceInstancesCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceResourceInstancesCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceResourceInstancesCountInvoker) Invoke() (*model.ShowInstanceResourceInstancesCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResourceInstancesCountResponse), nil
	}
}

type ShowInstanceRetentionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceRetentionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceRetentionPolicyInvoker) Invoke() (*model.ShowInstanceRetentionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceRetentionPolicyResponse), nil
	}
}

type ShowInstanceSignPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceSignPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceSignPolicyInvoker) Invoke() (*model.ShowInstanceSignPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceSignPolicyResponse), nil
	}
}

type ShowInstanceWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceWebhookInvoker) Invoke() (*model.ShowInstanceWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceWebhookResponse), nil
	}
}

type ShowSubResourceInstancesCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubResourceInstancesCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSubResourceInstancesCountInvoker) Invoke() (*model.ShowSubResourceInstancesCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubResourceInstancesCountResponse), nil
	}
}

type StartManualScanningInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartManualScanningInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartManualScanningInvoker) Invoke() (*model.StartManualScanningResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartManualScanningResponse), nil
	}
}

type StopInstanceReplicationPolicyExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInstanceReplicationPolicyExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInstanceReplicationPolicyExecutionInvoker) Invoke() (*model.StopInstanceReplicationPolicyExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInstanceReplicationPolicyExecutionResponse), nil
	}
}

type UpdateDomainNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainNameInvoker) Invoke() (*model.UpdateDomainNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainNameResponse), nil
	}
}

type UpdateImmutableRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImmutableRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateImmutableRuleInvoker) Invoke() (*model.UpdateImmutableRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImmutableRuleResponse), nil
	}
}

type UpdateInstanceConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceConfigurationInvoker) Invoke() (*model.UpdateInstanceConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

type UpdateInstanceEndpointPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceEndpointPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceEndpointPolicyInvoker) Invoke() (*model.UpdateInstanceEndpointPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceEndpointPolicyResponse), nil
	}
}

type UpdateInstanceLtCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceLtCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceLtCredentialInvoker) Invoke() (*model.UpdateInstanceLtCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceLtCredentialResponse), nil
	}
}

type UpdateInstanceNamespaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceNamespaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceNamespaceInvoker) Invoke() (*model.UpdateInstanceNamespaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNamespaceResponse), nil
	}
}

type UpdateInstanceRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceRegistryInvoker) Invoke() (*model.UpdateInstanceRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceRegistryResponse), nil
	}
}

type UpdateInstanceReplicationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceReplicationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceReplicationPolicyInvoker) Invoke() (*model.UpdateInstanceReplicationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceReplicationPolicyResponse), nil
	}
}

type UpdateInstanceRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceRepositoryInvoker) Invoke() (*model.UpdateInstanceRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceRepositoryResponse), nil
	}
}

type UpdateInstanceRetentionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceRetentionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceRetentionPolicyInvoker) Invoke() (*model.UpdateInstanceRetentionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceRetentionPolicyResponse), nil
	}
}

type UpdateInstanceSignPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceSignPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceSignPolicyInvoker) Invoke() (*model.UpdateInstanceSignPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceSignPolicyResponse), nil
	}
}

type UpdateInstanceWebhookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceWebhookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceWebhookInvoker) Invoke() (*model.UpdateInstanceWebhookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceWebhookResponse), nil
	}
}
