package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenteroidc/v1/model"
)

type IdentityCenterOIDCClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterOIDCClient(hcClient *httpclient.HcHttpClient) *IdentityCenterOIDCClient {
	return &IdentityCenterOIDCClient{HcClient: hcClient}
}

func IdentityCenterOIDCClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("v1.IdentityCenterOIDCCredentials")
	return builder
}

// RegisterClient 注册客户端
//
// 向IAM身份中心注册客户端，这允许客户端启动设备授权，输出应该持久化以便于身份验证请求重用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterOIDCClient) RegisterClient(request *model.RegisterClientRequest) (*model.RegisterClientResponse, error) {
	requestDef := GenReqDefForRegisterClient()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterClientResponse), nil
	}
}

// RegisterClientInvoker 注册客户端
func (c *IdentityCenterOIDCClient) RegisterClientInvoker(request *model.RegisterClientRequest) *RegisterClientInvoker {
	requestDef := GenReqDefForRegisterClient()
	return &RegisterClientInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDeviceAuthorization 请求设备授权
//
// 发起设备授权请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterOIDCClient) StartDeviceAuthorization(request *model.StartDeviceAuthorizationRequest) (*model.StartDeviceAuthorizationResponse, error) {
	requestDef := GenReqDefForStartDeviceAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDeviceAuthorizationResponse), nil
	}
}

// StartDeviceAuthorizationInvoker 请求设备授权
func (c *IdentityCenterOIDCClient) StartDeviceAuthorizationInvoker(request *model.StartDeviceAuthorizationRequest) *StartDeviceAuthorizationInvoker {
	requestDef := GenReqDefForStartDeviceAuthorization()
	return &StartDeviceAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateToken 创建令牌
//
// 获取访问令牌
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterOIDCClient) CreateToken(request *model.CreateTokenRequest) (*model.CreateTokenResponse, error) {
	requestDef := GenReqDefForCreateToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTokenResponse), nil
	}
}

// CreateTokenInvoker 创建令牌
func (c *IdentityCenterOIDCClient) CreateTokenInvoker(request *model.CreateTokenRequest) *CreateTokenInvoker {
	requestDef := GenReqDefForCreateToken()
	return &CreateTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
