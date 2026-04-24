package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/agentidentity/v1/model"
)

type AgentIdentityClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAgentIdentityClient(hcClient *httpclient.HcHttpClient) *AgentIdentityClient {
	return &AgentIdentityClient{HcClient: hcClient}
}

func AgentIdentityClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateApiKeyCredentialProvider 创建API密钥凭证提供者
//
// Creates a new API key credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateApiKeyCredentialProvider(request *model.CreateApiKeyCredentialProviderRequest) (*model.CreateApiKeyCredentialProviderResponse, error) {
	requestDef := GenReqDefForCreateApiKeyCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiKeyCredentialProviderResponse), nil
	}
}

// CreateApiKeyCredentialProviderInvoker 创建API密钥凭证提供者
func (c *AgentIdentityClient) CreateApiKeyCredentialProviderInvoker(request *model.CreateApiKeyCredentialProviderRequest) *CreateApiKeyCredentialProviderInvoker {
	requestDef := GenReqDefForCreateApiKeyCredentialProvider()
	return &CreateApiKeyCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiKeyCredentialProvider 删除API密钥凭证提供者
//
// Deletes an API key credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) DeleteApiKeyCredentialProvider(request *model.DeleteApiKeyCredentialProviderRequest) (*model.DeleteApiKeyCredentialProviderResponse, error) {
	requestDef := GenReqDefForDeleteApiKeyCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiKeyCredentialProviderResponse), nil
	}
}

// DeleteApiKeyCredentialProviderInvoker 删除API密钥凭证提供者
func (c *AgentIdentityClient) DeleteApiKeyCredentialProviderInvoker(request *model.DeleteApiKeyCredentialProviderRequest) *DeleteApiKeyCredentialProviderInvoker {
	requestDef := GenReqDefForDeleteApiKeyCredentialProvider()
	return &DeleteApiKeyCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetApiKeyCredentialProvider 查询API密钥凭证提供者详情
//
// Gets details of an API key credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetApiKeyCredentialProvider(request *model.GetApiKeyCredentialProviderRequest) (*model.GetApiKeyCredentialProviderResponse, error) {
	requestDef := GenReqDefForGetApiKeyCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetApiKeyCredentialProviderResponse), nil
	}
}

// GetApiKeyCredentialProviderInvoker 查询API密钥凭证提供者详情
func (c *AgentIdentityClient) GetApiKeyCredentialProviderInvoker(request *model.GetApiKeyCredentialProviderRequest) *GetApiKeyCredentialProviderInvoker {
	requestDef := GenReqDefForGetApiKeyCredentialProvider()
	return &GetApiKeyCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiKeyCredentialProviders 查询API密钥凭证提供者列表
//
// Lists API key credential providers.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) ListApiKeyCredentialProviders(request *model.ListApiKeyCredentialProvidersRequest) (*model.ListApiKeyCredentialProvidersResponse, error) {
	requestDef := GenReqDefForListApiKeyCredentialProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiKeyCredentialProvidersResponse), nil
	}
}

// ListApiKeyCredentialProvidersInvoker 查询API密钥凭证提供者列表
func (c *AgentIdentityClient) ListApiKeyCredentialProvidersInvoker(request *model.ListApiKeyCredentialProvidersRequest) *ListApiKeyCredentialProvidersInvoker {
	requestDef := GenReqDefForListApiKeyCredentialProviders()
	return &ListApiKeyCredentialProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApiKeyCredentialProvider 更新API密钥凭证提供者
//
// Updates the API key of an existing API key credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) UpdateApiKeyCredentialProvider(request *model.UpdateApiKeyCredentialProviderRequest) (*model.UpdateApiKeyCredentialProviderResponse, error) {
	requestDef := GenReqDefForUpdateApiKeyCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiKeyCredentialProviderResponse), nil
	}
}

// UpdateApiKeyCredentialProviderInvoker 更新API密钥凭证提供者
func (c *AgentIdentityClient) UpdateApiKeyCredentialProviderInvoker(request *model.UpdateApiKeyCredentialProviderRequest) *UpdateApiKeyCredentialProviderInvoker {
	requestDef := GenReqDefForUpdateApiKeyCredentialProvider()
	return &UpdateApiKeyCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompleteResourceTokenAuth Confirm user authentication session for OAuth2.0 tokens
//
// Confirms the user authentication session to obtain OAuth2.0 tokens for a resource
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CompleteResourceTokenAuth(request *model.CompleteResourceTokenAuthRequest) (*model.CompleteResourceTokenAuthResponse, error) {
	requestDef := GenReqDefForCompleteResourceTokenAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompleteResourceTokenAuthResponse), nil
	}
}

// CompleteResourceTokenAuthInvoker Confirm user authentication session for OAuth2.0 tokens
func (c *AgentIdentityClient) CompleteResourceTokenAuthInvoker(request *model.CompleteResourceTokenAuthRequest) *CompleteResourceTokenAuthInvoker {
	requestDef := GenReqDefForCompleteResourceTokenAuth()
	return &CompleteResourceTokenAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetResourceApiKey Retrieve API key from resource credential provider
//
// Retrieves the API key associated with a specified resource credential provider
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetResourceApiKey(request *model.GetResourceApiKeyRequest) (*model.GetResourceApiKeyResponse, error) {
	requestDef := GenReqDefForGetResourceApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetResourceApiKeyResponse), nil
	}
}

// GetResourceApiKeyInvoker Retrieve API key from resource credential provider
func (c *AgentIdentityClient) GetResourceApiKeyInvoker(request *model.GetResourceApiKeyRequest) *GetResourceApiKeyInvoker {
	requestDef := GenReqDefForGetResourceApiKey()
	return &GetResourceApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetResourceOauth2Token Retrieve OAuth2.0 token from resource credential provider
//
// Returns the OAuth2.0 token for the specified resource using the configured flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetResourceOauth2Token(request *model.GetResourceOauth2TokenRequest) (*model.GetResourceOauth2TokenResponse, error) {
	requestDef := GenReqDefForGetResourceOauth2Token()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetResourceOauth2TokenResponse), nil
	}
}

// GetResourceOauth2TokenInvoker Retrieve OAuth2.0 token from resource credential provider
func (c *AgentIdentityClient) GetResourceOauth2TokenInvoker(request *model.GetResourceOauth2TokenRequest) *GetResourceOauth2TokenInvoker {
	requestDef := GenReqDefForGetResourceOauth2Token()
	return &GetResourceOauth2TokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetResourceStsToken Retrieve STS credentials from STS credential provider
//
// Retrieves temporary STS credentials from a specified STS credential provider
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetResourceStsToken(request *model.GetResourceStsTokenRequest) (*model.GetResourceStsTokenResponse, error) {
	requestDef := GenReqDefForGetResourceStsToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetResourceStsTokenResponse), nil
	}
}

// GetResourceStsTokenInvoker Retrieve STS credentials from STS credential provider
func (c *AgentIdentityClient) GetResourceStsTokenInvoker(request *model.GetResourceStsTokenRequest) *GetResourceStsTokenInvoker {
	requestDef := GenReqDefForGetResourceStsToken()
	return &GetResourceStsTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Oauth2Authorize OAuth2.0 Pushed Authorization Request (PAR) standard authorize API
//
// Core OAuth2 authorization endpoint following RFC 9126 PAR spec, only accepts authorization request via request_uri parameter to trigger user authorization flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) Oauth2Authorize(request *model.Oauth2AuthorizeRequest) (*model.Oauth2AuthorizeResponse, error) {
	requestDef := GenReqDefForOauth2Authorize()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Oauth2AuthorizeResponse), nil
	}
}

// Oauth2AuthorizeInvoker OAuth2.0 Pushed Authorization Request (PAR) standard authorize API
func (c *AgentIdentityClient) Oauth2AuthorizeInvoker(request *model.Oauth2AuthorizeRequest) *Oauth2AuthorizeInvoker {
	requestDef := GenReqDefForOauth2Authorize()
	return &Oauth2AuthorizeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Oauth2Callback OAuth2.0 Standard Authorization Callback API
//
// OAuth2 redirect callback endpoint to receive authorization result after user consent/denial
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) Oauth2Callback(request *model.Oauth2CallbackRequest) (*model.Oauth2CallbackResponse, error) {
	requestDef := GenReqDefForOauth2Callback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Oauth2CallbackResponse), nil
	}
}

// Oauth2CallbackInvoker OAuth2.0 Standard Authorization Callback API
func (c *AgentIdentityClient) Oauth2CallbackInvoker(request *model.Oauth2CallbackRequest) *Oauth2CallbackInvoker {
	requestDef := GenReqDefForOauth2Callback()
	return &Oauth2CallbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOauth2CredentialProvider 创建OAuth2凭证提供者
//
// Creates a new OAuth2 credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateOauth2CredentialProvider(request *model.CreateOauth2CredentialProviderRequest) (*model.CreateOauth2CredentialProviderResponse, error) {
	requestDef := GenReqDefForCreateOauth2CredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOauth2CredentialProviderResponse), nil
	}
}

// CreateOauth2CredentialProviderInvoker 创建OAuth2凭证提供者
func (c *AgentIdentityClient) CreateOauth2CredentialProviderInvoker(request *model.CreateOauth2CredentialProviderRequest) *CreateOauth2CredentialProviderInvoker {
	requestDef := GenReqDefForCreateOauth2CredentialProvider()
	return &CreateOauth2CredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOauth2CredentialProvider 删除OAuth2凭证提供者
//
// Deletes an OAuth2 credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) DeleteOauth2CredentialProvider(request *model.DeleteOauth2CredentialProviderRequest) (*model.DeleteOauth2CredentialProviderResponse, error) {
	requestDef := GenReqDefForDeleteOauth2CredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOauth2CredentialProviderResponse), nil
	}
}

// DeleteOauth2CredentialProviderInvoker 删除OAuth2凭证提供者
func (c *AgentIdentityClient) DeleteOauth2CredentialProviderInvoker(request *model.DeleteOauth2CredentialProviderRequest) *DeleteOauth2CredentialProviderInvoker {
	requestDef := GenReqDefForDeleteOauth2CredentialProvider()
	return &DeleteOauth2CredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetOauth2CredentialProvider 查询OAuth2凭证提供者详情
//
// Gets details of an OAuth2 credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetOauth2CredentialProvider(request *model.GetOauth2CredentialProviderRequest) (*model.GetOauth2CredentialProviderResponse, error) {
	requestDef := GenReqDefForGetOauth2CredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetOauth2CredentialProviderResponse), nil
	}
}

// GetOauth2CredentialProviderInvoker 查询OAuth2凭证提供者详情
func (c *AgentIdentityClient) GetOauth2CredentialProviderInvoker(request *model.GetOauth2CredentialProviderRequest) *GetOauth2CredentialProviderInvoker {
	requestDef := GenReqDefForGetOauth2CredentialProvider()
	return &GetOauth2CredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOauth2CredentialProviders 查询OAuth2凭证提供者列表
//
// Lists OAuth2 credential providers.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) ListOauth2CredentialProviders(request *model.ListOauth2CredentialProvidersRequest) (*model.ListOauth2CredentialProvidersResponse, error) {
	requestDef := GenReqDefForListOauth2CredentialProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOauth2CredentialProvidersResponse), nil
	}
}

// ListOauth2CredentialProvidersInvoker 查询OAuth2凭证提供者列表
func (c *AgentIdentityClient) ListOauth2CredentialProvidersInvoker(request *model.ListOauth2CredentialProvidersRequest) *ListOauth2CredentialProvidersInvoker {
	requestDef := GenReqDefForListOauth2CredentialProviders()
	return &ListOauth2CredentialProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOauth2CredentialProvider 更新OAuth2凭证提供者
//
// Updates an existing OAuth2 credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) UpdateOauth2CredentialProvider(request *model.UpdateOauth2CredentialProviderRequest) (*model.UpdateOauth2CredentialProviderResponse, error) {
	requestDef := GenReqDefForUpdateOauth2CredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOauth2CredentialProviderResponse), nil
	}
}

// UpdateOauth2CredentialProviderInvoker 更新OAuth2凭证提供者
func (c *AgentIdentityClient) UpdateOauth2CredentialProviderInvoker(request *model.UpdateOauth2CredentialProviderRequest) *UpdateOauth2CredentialProviderInvoker {
	requestDef := GenReqDefForUpdateOauth2CredentialProvider()
	return &UpdateOauth2CredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStsCredentialProvider 创建STS凭证提供者
//
// Creates a new STS credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateStsCredentialProvider(request *model.CreateStsCredentialProviderRequest) (*model.CreateStsCredentialProviderResponse, error) {
	requestDef := GenReqDefForCreateStsCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStsCredentialProviderResponse), nil
	}
}

// CreateStsCredentialProviderInvoker 创建STS凭证提供者
func (c *AgentIdentityClient) CreateStsCredentialProviderInvoker(request *model.CreateStsCredentialProviderRequest) *CreateStsCredentialProviderInvoker {
	requestDef := GenReqDefForCreateStsCredentialProvider()
	return &CreateStsCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStsCredentialProvider 删除STS凭证提供者
//
// Deletes an STS credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) DeleteStsCredentialProvider(request *model.DeleteStsCredentialProviderRequest) (*model.DeleteStsCredentialProviderResponse, error) {
	requestDef := GenReqDefForDeleteStsCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStsCredentialProviderResponse), nil
	}
}

// DeleteStsCredentialProviderInvoker 删除STS凭证提供者
func (c *AgentIdentityClient) DeleteStsCredentialProviderInvoker(request *model.DeleteStsCredentialProviderRequest) *DeleteStsCredentialProviderInvoker {
	requestDef := GenReqDefForDeleteStsCredentialProvider()
	return &DeleteStsCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetStsCredentialProvider 查询STS凭证提供者详情
//
// Gets details of an STS credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetStsCredentialProvider(request *model.GetStsCredentialProviderRequest) (*model.GetStsCredentialProviderResponse, error) {
	requestDef := GenReqDefForGetStsCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetStsCredentialProviderResponse), nil
	}
}

// GetStsCredentialProviderInvoker 查询STS凭证提供者详情
func (c *AgentIdentityClient) GetStsCredentialProviderInvoker(request *model.GetStsCredentialProviderRequest) *GetStsCredentialProviderInvoker {
	requestDef := GenReqDefForGetStsCredentialProvider()
	return &GetStsCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStsCredentialProviders 查询STS凭证提供者列表
//
// Lists STS credential providers.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) ListStsCredentialProviders(request *model.ListStsCredentialProvidersRequest) (*model.ListStsCredentialProvidersResponse, error) {
	requestDef := GenReqDefForListStsCredentialProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStsCredentialProvidersResponse), nil
	}
}

// ListStsCredentialProvidersInvoker 查询STS凭证提供者列表
func (c *AgentIdentityClient) ListStsCredentialProvidersInvoker(request *model.ListStsCredentialProvidersRequest) *ListStsCredentialProvidersInvoker {
	requestDef := GenReqDefForListStsCredentialProviders()
	return &ListStsCredentialProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStsCredentialProvider 更新STS凭证提供者
//
// Updates an existing STS credential provider.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) UpdateStsCredentialProvider(request *model.UpdateStsCredentialProviderRequest) (*model.UpdateStsCredentialProviderResponse, error) {
	requestDef := GenReqDefForUpdateStsCredentialProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStsCredentialProviderResponse), nil
	}
}

// UpdateStsCredentialProviderInvoker 更新STS凭证提供者
func (c *AgentIdentityClient) UpdateStsCredentialProviderInvoker(request *model.UpdateStsCredentialProviderRequest) *UpdateStsCredentialProviderInvoker {
	requestDef := GenReqDefForUpdateStsCredentialProvider()
	return &UpdateStsCredentialProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetTokenVault 查询令牌保管库详情
//
// Gets details of a token vault.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetTokenVault(request *model.GetTokenVaultRequest) (*model.GetTokenVaultResponse, error) {
	requestDef := GenReqDefForGetTokenVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetTokenVaultResponse), nil
	}
}

// GetTokenVaultInvoker 查询令牌保管库详情
func (c *AgentIdentityClient) GetTokenVaultInvoker(request *model.GetTokenVaultRequest) *GetTokenVaultInvoker {
	requestDef := GenReqDefForGetTokenVault()
	return &GetTokenVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkloadAccessToken Create workload access token (not acting on behalf of a user)
//
// Retrieves a workload access token for agentic workloads that do not act on behalf of a user
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateWorkloadAccessToken(request *model.CreateWorkloadAccessTokenRequest) (*model.CreateWorkloadAccessTokenResponse, error) {
	requestDef := GenReqDefForCreateWorkloadAccessToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkloadAccessTokenResponse), nil
	}
}

// CreateWorkloadAccessTokenInvoker Create workload access token (not acting on behalf of a user)
func (c *AgentIdentityClient) CreateWorkloadAccessTokenInvoker(request *model.CreateWorkloadAccessTokenRequest) *CreateWorkloadAccessTokenInvoker {
	requestDef := GenReqDefForCreateWorkloadAccessToken()
	return &CreateWorkloadAccessTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkloadAccessTokenForJwt Create workload access token using JWT (acting on behalf of a user)
//
// Retrieves a workload access token for agentic workloads acting on behalf of a user, using a JWT token
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateWorkloadAccessTokenForJwt(request *model.CreateWorkloadAccessTokenForJwtRequest) (*model.CreateWorkloadAccessTokenForJwtResponse, error) {
	requestDef := GenReqDefForCreateWorkloadAccessTokenForJwt()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkloadAccessTokenForJwtResponse), nil
	}
}

// CreateWorkloadAccessTokenForJwtInvoker Create workload access token using JWT (acting on behalf of a user)
func (c *AgentIdentityClient) CreateWorkloadAccessTokenForJwtInvoker(request *model.CreateWorkloadAccessTokenForJwtRequest) *CreateWorkloadAccessTokenForJwtInvoker {
	requestDef := GenReqDefForCreateWorkloadAccessTokenForJwt()
	return &CreateWorkloadAccessTokenForJwtInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkloadAccessTokenForUserId Create workload access token using user ID (acting on behalf of a user)
//
// Retrieves a workload access token for agentic workloads acting on behalf of a user, using the user&#39;s ID
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateWorkloadAccessTokenForUserId(request *model.CreateWorkloadAccessTokenForUserIdRequest) (*model.CreateWorkloadAccessTokenForUserIdResponse, error) {
	requestDef := GenReqDefForCreateWorkloadAccessTokenForUserId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkloadAccessTokenForUserIdResponse), nil
	}
}

// CreateWorkloadAccessTokenForUserIdInvoker Create workload access token using user ID (acting on behalf of a user)
func (c *AgentIdentityClient) CreateWorkloadAccessTokenForUserIdInvoker(request *model.CreateWorkloadAccessTokenForUserIdRequest) *CreateWorkloadAccessTokenForUserIdInvoker {
	requestDef := GenReqDefForCreateWorkloadAccessTokenForUserId()
	return &CreateWorkloadAccessTokenForUserIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkloadIdentity 创建工作负载身份
//
// Creates a new workload identity.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) CreateWorkloadIdentity(request *model.CreateWorkloadIdentityRequest) (*model.CreateWorkloadIdentityResponse, error) {
	requestDef := GenReqDefForCreateWorkloadIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkloadIdentityResponse), nil
	}
}

// CreateWorkloadIdentityInvoker 创建工作负载身份
func (c *AgentIdentityClient) CreateWorkloadIdentityInvoker(request *model.CreateWorkloadIdentityRequest) *CreateWorkloadIdentityInvoker {
	requestDef := GenReqDefForCreateWorkloadIdentity()
	return &CreateWorkloadIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkloadIdentity 删除工作负载身份
//
// Deletes a workload identity.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) DeleteWorkloadIdentity(request *model.DeleteWorkloadIdentityRequest) (*model.DeleteWorkloadIdentityResponse, error) {
	requestDef := GenReqDefForDeleteWorkloadIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkloadIdentityResponse), nil
	}
}

// DeleteWorkloadIdentityInvoker 删除工作负载身份
func (c *AgentIdentityClient) DeleteWorkloadIdentityInvoker(request *model.DeleteWorkloadIdentityRequest) *DeleteWorkloadIdentityInvoker {
	requestDef := GenReqDefForDeleteWorkloadIdentity()
	return &DeleteWorkloadIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetWorkloadIdentity 查询工作负载身份详情
//
// Gets details of a workload identity.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetWorkloadIdentity(request *model.GetWorkloadIdentityRequest) (*model.GetWorkloadIdentityResponse, error) {
	requestDef := GenReqDefForGetWorkloadIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetWorkloadIdentityResponse), nil
	}
}

// GetWorkloadIdentityInvoker 查询工作负载身份详情
func (c *AgentIdentityClient) GetWorkloadIdentityInvoker(request *model.GetWorkloadIdentityRequest) *GetWorkloadIdentityInvoker {
	requestDef := GenReqDefForGetWorkloadIdentity()
	return &GetWorkloadIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetWorkloadIdentityAuthorizerConfiguration 查询工作负载身份的授权配置
//
// Gets the authorizer configuration of a workload identity.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) GetWorkloadIdentityAuthorizerConfiguration(request *model.GetWorkloadIdentityAuthorizerConfigurationRequest) (*model.GetWorkloadIdentityAuthorizerConfigurationResponse, error) {
	requestDef := GenReqDefForGetWorkloadIdentityAuthorizerConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetWorkloadIdentityAuthorizerConfigurationResponse), nil
	}
}

// GetWorkloadIdentityAuthorizerConfigurationInvoker 查询工作负载身份的授权配置
func (c *AgentIdentityClient) GetWorkloadIdentityAuthorizerConfigurationInvoker(request *model.GetWorkloadIdentityAuthorizerConfigurationRequest) *GetWorkloadIdentityAuthorizerConfigurationInvoker {
	requestDef := GenReqDefForGetWorkloadIdentityAuthorizerConfiguration()
	return &GetWorkloadIdentityAuthorizerConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkloadIdentities 查询工作负载身份列表
//
// Lists workload identities.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) ListWorkloadIdentities(request *model.ListWorkloadIdentitiesRequest) (*model.ListWorkloadIdentitiesResponse, error) {
	requestDef := GenReqDefForListWorkloadIdentities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkloadIdentitiesResponse), nil
	}
}

// ListWorkloadIdentitiesInvoker 查询工作负载身份列表
func (c *AgentIdentityClient) ListWorkloadIdentitiesInvoker(request *model.ListWorkloadIdentitiesRequest) *ListWorkloadIdentitiesInvoker {
	requestDef := GenReqDefForListWorkloadIdentities()
	return &ListWorkloadIdentitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkloadIdentity 更新工作负载身份
//
// Updates an existing workload identity.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AgentIdentityClient) UpdateWorkloadIdentity(request *model.UpdateWorkloadIdentityRequest) (*model.UpdateWorkloadIdentityResponse, error) {
	requestDef := GenReqDefForUpdateWorkloadIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkloadIdentityResponse), nil
	}
}

// UpdateWorkloadIdentityInvoker 更新工作负载身份
func (c *AgentIdentityClient) UpdateWorkloadIdentityInvoker(request *model.UpdateWorkloadIdentityRequest) *UpdateWorkloadIdentityInvoker {
	requestDef := GenReqDefForUpdateWorkloadIdentity()
	return &UpdateWorkloadIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
