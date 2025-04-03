package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterportalapi/v1/model"
)

type IdentityCenterPortalAPIClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterPortalAPIClient(hcClient *httpclient.HcHttpClient) *IdentityCenterPortalAPIClient {
	return &IdentityCenterPortalAPIClient{HcClient: hcClient}
}

func IdentityCenterPortalAPIClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("v1.IdentityCenterPortalAPICredentials")
	return builder
}

// ListAccounts 列出账号
//
// 列出分配给用户的所有账号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterPortalAPIClient) ListAccounts(request *model.ListAccountsRequest) (*model.ListAccountsResponse, error) {
	requestDef := GenReqDefForListAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountsResponse), nil
	}
}

// ListAccountsInvoker 列出账号
func (c *IdentityCenterPortalAPIClient) ListAccountsInvoker(request *model.ListAccountsRequest) *ListAccountsInvoker {
	requestDef := GenReqDefForListAccounts()
	return &ListAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAgencies 列出账号委托
//
// 列出账号分配给用户的所有委托或信任委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterPortalAPIClient) ListAccountAgencies(request *model.ListAccountAgenciesRequest) (*model.ListAccountAgenciesResponse, error) {
	requestDef := GenReqDefForListAccountAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountAgenciesResponse), nil
	}
}

// ListAccountAgenciesInvoker 列出账号委托
func (c *IdentityCenterPortalAPIClient) ListAccountAgenciesInvoker(request *model.ListAccountAgenciesRequest) *ListAccountAgenciesInvoker {
	requestDef := GenReqDefForListAccountAgencies()
	return &ListAccountAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetAgencyCredentials 获取委托凭证
//
// 获取分配给用户的指定委托或信任委托的STS短期凭证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterPortalAPIClient) GetAgencyCredentials(request *model.GetAgencyCredentialsRequest) (*model.GetAgencyCredentialsResponse, error) {
	requestDef := GenReqDefForGetAgencyCredentials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAgencyCredentialsResponse), nil
	}
}

// GetAgencyCredentialsInvoker 获取委托凭证
func (c *IdentityCenterPortalAPIClient) GetAgencyCredentialsInvoker(request *model.GetAgencyCredentialsRequest) *GetAgencyCredentialsInvoker {
	requestDef := GenReqDefForGetAgencyCredentials()
	return &GetAgencyCredentialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Logout 登出用户
//
// 向IAM身份中心服务发送API调用，使相应的服务端IAM身份中心登录会话失效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterPortalAPIClient) Logout(request *model.LogoutRequest) (*model.LogoutResponse, error) {
	requestDef := GenReqDefForLogout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LogoutResponse), nil
	}
}

// LogoutInvoker 登出用户
func (c *IdentityCenterPortalAPIClient) LogoutInvoker(request *model.LogoutRequest) *LogoutInvoker {
	requestDef := GenReqDefForLogout()
	return &LogoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
