package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/agentidentity/v1/model"
)

type CreateApiKeyCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiKeyCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApiKeyCredentialProviderInvoker) Invoke() (*model.CreateApiKeyCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiKeyCredentialProviderResponse), nil
	}
}

type DeleteApiKeyCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiKeyCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiKeyCredentialProviderInvoker) Invoke() (*model.DeleteApiKeyCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiKeyCredentialProviderResponse), nil
	}
}

type GetApiKeyCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetApiKeyCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetApiKeyCredentialProviderInvoker) Invoke() (*model.GetApiKeyCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetApiKeyCredentialProviderResponse), nil
	}
}

type ListApiKeyCredentialProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiKeyCredentialProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiKeyCredentialProvidersInvoker) Invoke() (*model.ListApiKeyCredentialProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiKeyCredentialProvidersResponse), nil
	}
}

type UpdateApiKeyCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApiKeyCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApiKeyCredentialProviderInvoker) Invoke() (*model.UpdateApiKeyCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApiKeyCredentialProviderResponse), nil
	}
}

type CompleteResourceTokenAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompleteResourceTokenAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CompleteResourceTokenAuthInvoker) Invoke() (*model.CompleteResourceTokenAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompleteResourceTokenAuthResponse), nil
	}
}

type GetResourceApiKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetResourceApiKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetResourceApiKeyInvoker) Invoke() (*model.GetResourceApiKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetResourceApiKeyResponse), nil
	}
}

type GetResourceOauth2TokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetResourceOauth2TokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetResourceOauth2TokenInvoker) Invoke() (*model.GetResourceOauth2TokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetResourceOauth2TokenResponse), nil
	}
}

type GetResourceStsTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetResourceStsTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetResourceStsTokenInvoker) Invoke() (*model.GetResourceStsTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetResourceStsTokenResponse), nil
	}
}

type ListIdentityProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIdentityProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIdentityProvidersInvoker) Invoke() (*model.ListIdentityProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIdentityProvidersResponse), nil
	}
}

type Oauth2AuthorizeInvoker struct {
	*invoker.BaseInvoker
}

func (i *Oauth2AuthorizeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *Oauth2AuthorizeInvoker) Invoke() (*model.Oauth2AuthorizeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oauth2AuthorizeResponse), nil
	}
}

type Oauth2CallbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *Oauth2CallbackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *Oauth2CallbackInvoker) Invoke() (*model.Oauth2CallbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oauth2CallbackResponse), nil
	}
}

type CreateOauth2CredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOauth2CredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOauth2CredentialProviderInvoker) Invoke() (*model.CreateOauth2CredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOauth2CredentialProviderResponse), nil
	}
}

type DeleteOauth2CredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOauth2CredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOauth2CredentialProviderInvoker) Invoke() (*model.DeleteOauth2CredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOauth2CredentialProviderResponse), nil
	}
}

type GetOauth2CredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetOauth2CredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetOauth2CredentialProviderInvoker) Invoke() (*model.GetOauth2CredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetOauth2CredentialProviderResponse), nil
	}
}

type ListOauth2CredentialProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOauth2CredentialProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOauth2CredentialProvidersInvoker) Invoke() (*model.ListOauth2CredentialProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOauth2CredentialProvidersResponse), nil
	}
}

type UpdateOauth2CredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOauth2CredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOauth2CredentialProviderInvoker) Invoke() (*model.UpdateOauth2CredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOauth2CredentialProviderResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type GetPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetPolicyInvoker) Invoke() (*model.GetPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPolicyResponse), nil
	}
}

type ListPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesInvoker) Invoke() (*model.ListPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesResponse), nil
	}
}

type UpdatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyInvoker) Invoke() (*model.UpdatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyResponse), nil
	}
}

type CreatePolicyEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyEngineInvoker) Invoke() (*model.CreatePolicyEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyEngineResponse), nil
	}
}

type DeletePolicyEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyEngineInvoker) Invoke() (*model.DeletePolicyEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyEngineResponse), nil
	}
}

type GetPolicyEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetPolicyEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetPolicyEngineInvoker) Invoke() (*model.GetPolicyEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPolicyEngineResponse), nil
	}
}

type ListPolicyEngineAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyEngineAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyEngineAttachmentsInvoker) Invoke() (*model.ListPolicyEngineAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyEngineAttachmentsResponse), nil
	}
}

type ListPolicyEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyEnginesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyEnginesInvoker) Invoke() (*model.ListPolicyEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyEnginesResponse), nil
	}
}

type UpdatePolicyEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyEngineInvoker) Invoke() (*model.UpdatePolicyEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyEngineResponse), nil
	}
}

type CreateStsCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStsCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStsCredentialProviderInvoker) Invoke() (*model.CreateStsCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStsCredentialProviderResponse), nil
	}
}

type DeleteStsCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStsCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStsCredentialProviderInvoker) Invoke() (*model.DeleteStsCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStsCredentialProviderResponse), nil
	}
}

type GetStsCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetStsCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetStsCredentialProviderInvoker) Invoke() (*model.GetStsCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetStsCredentialProviderResponse), nil
	}
}

type ListStsCredentialProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStsCredentialProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStsCredentialProvidersInvoker) Invoke() (*model.ListStsCredentialProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStsCredentialProvidersResponse), nil
	}
}

type UpdateStsCredentialProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStsCredentialProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStsCredentialProviderInvoker) Invoke() (*model.UpdateStsCredentialProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStsCredentialProviderResponse), nil
	}
}

type GetTokenVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetTokenVaultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetTokenVaultInvoker) Invoke() (*model.GetTokenVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetTokenVaultResponse), nil
	}
}

type UpdateTokenVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTokenVaultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTokenVaultInvoker) Invoke() (*model.UpdateTokenVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTokenVaultResponse), nil
	}
}

type CreateWorkloadAccessTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkloadAccessTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkloadAccessTokenInvoker) Invoke() (*model.CreateWorkloadAccessTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkloadAccessTokenResponse), nil
	}
}

type CreateWorkloadAccessTokenForJwtInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkloadAccessTokenForJwtInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkloadAccessTokenForJwtInvoker) Invoke() (*model.CreateWorkloadAccessTokenForJwtResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkloadAccessTokenForJwtResponse), nil
	}
}

type CreateWorkloadAccessTokenForUserIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkloadAccessTokenForUserIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkloadAccessTokenForUserIdInvoker) Invoke() (*model.CreateWorkloadAccessTokenForUserIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkloadAccessTokenForUserIdResponse), nil
	}
}

type CreateWorkloadIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkloadIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkloadIdentityInvoker) Invoke() (*model.CreateWorkloadIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkloadIdentityResponse), nil
	}
}

type DeleteWorkloadIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkloadIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkloadIdentityInvoker) Invoke() (*model.DeleteWorkloadIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkloadIdentityResponse), nil
	}
}

type GetWorkloadIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetWorkloadIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetWorkloadIdentityInvoker) Invoke() (*model.GetWorkloadIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetWorkloadIdentityResponse), nil
	}
}

type GetWorkloadIdentityAuthorizerConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetWorkloadIdentityAuthorizerConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetWorkloadIdentityAuthorizerConfigurationInvoker) Invoke() (*model.GetWorkloadIdentityAuthorizerConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetWorkloadIdentityAuthorizerConfigurationResponse), nil
	}
}

type ListWorkloadIdentitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkloadIdentitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkloadIdentitiesInvoker) Invoke() (*model.ListWorkloadIdentitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkloadIdentitiesResponse), nil
	}
}

type UpdateWorkloadIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkloadIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkloadIdentityInvoker) Invoke() (*model.UpdateWorkloadIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkloadIdentityResponse), nil
	}
}
