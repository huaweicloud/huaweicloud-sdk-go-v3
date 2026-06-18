package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sts/v1/model"
)

type StsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewStsClient(hcClient *httpclient.HcHttpClient) *StsClient {
	return &StsClient{HcClient: hcClient}
}

func StsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AssumeAgency 通过委托或者信任委托获取临时安全凭证
//
// 通过委托或者信任委托获取临时安全凭证，临时安全凭证可用于对云资源发起访问。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *StsClient) AssumeAgency(request *model.AssumeAgencyRequest) (*model.AssumeAgencyResponse, error) {
	requestDef := GenReqDefForAssumeAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssumeAgencyResponse), nil
	}
}

// AssumeAgencyInvoker 通过委托或者信任委托获取临时安全凭证
func (c *StsClient) AssumeAgencyInvoker(request *model.AssumeAgencyRequest) *AssumeAgencyInvoker {
	requestDef := GenReqDefForAssumeAgency()
	return &AssumeAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssumeAgencyWithOIDC 通过使用OIDC协议SSO的信任委托获取临时安全凭证
//
// 为通过 OIDC 身份提供商令牌验证的用户返回一组临时安全凭证。此操作提供了一种机制，可以让外部的任何兼容 OIDC 身份提供商使用信任委托的临时安全凭证华为云访问，而无需使用 IAM 用户的凭证。
//
// **会话持续时间：**
// 默认情况下，通过 AssumeAgencyWithOIDC 创建的临时安全凭证有效期为一小时。您可以使用可选的 duration_seconds 参数来指定会话的持续时间，duration_seconds 取值范围是从 900 秒（15 分钟）到该信任委托设置的最大会话持续时长，最大会话持续时长的取值范围可以从 1 小时到 12 小时。注意：委托链会将您的会话持续时间限制为最多一小时，当您使用 AssumeAgency API 操作来进行委托链的切换时，如果您提供了大于一小时的 duration_seconds 参数值，该操作将会失败。
//
// **权限：**
// 调用 AssumeAgencyWithOIDC 不需要使用华为云凭证。调用者的身份是通过使用您 JWKS 端点中的公钥进行验证的。
// 您可以使用 policy 和 policy_ids 参数传递自定义策略和已有的身份策略来限制本次会话获得的临时安全凭证的权限范围，最终获得临时安全凭证的权限是 policy 和 policy_ids 与信任委托身上附加的身份策略的交集。
//
// **标签：**
// 在信任委托的信任策略中添加了 sts::tagSession 授权项时，您可以配置您的身份提供商，将属性作为会话标签传递到您的 ID Token 中。每个会话标签由一个键（Key）和一个值（Value）组成。您最多可以传递 20 个会话标签。纯文本形式的会话标签键不得超过 128 个字符，值不得超过 255 个字符。您也可以传递与信任委托身上标签同名的会话标签，此时会话标签会覆盖具有相同键的信任委托标签。您可以将会话标签设置为可传递的 (Transitive)，可传递的会话标签在角色链期间会持续保留。
//
// **身份：**
// 在您的应用程序可以调用 AssumeAgencyWithOIDC 之前，您必须使用您的账号在 IAM 中创建 OIDC 提供商和信任委托，并在信任委托的信任策略中指定该 OIDC 提供商，然后还需要配置您的 OIDC 身份提供商以颁发 IAM 所需的 ID Token。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *StsClient) AssumeAgencyWithOIDC(request *model.AssumeAgencyWithOidcRequest) (*model.AssumeAgencyWithOidcResponse, error) {
	requestDef := GenReqDefForAssumeAgencyWithOIDC()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssumeAgencyWithOidcResponse), nil
	}
}

// AssumeAgencyWithOIDCInvoker 通过使用OIDC协议SSO的信任委托获取临时安全凭证
func (c *StsClient) AssumeAgencyWithOIDCInvoker(request *model.AssumeAgencyWithOidcRequest) *AssumeAgencyWithOIDCInvoker {
	requestDef := GenReqDefForAssumeAgencyWithOIDC()
	return &AssumeAgencyWithOIDCInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssumeAgencyWithSAML 通过使用SAML协议SSO的信任委托获取临时安全凭证
//
// 为通过 SAML 身份验证响应（SAML Authentication Response）验证的用户返回一组临时安全凭证。此操作提供了一种机制，可以让外部的 SAML 身份提供商使用信任委托的临时安全凭证华为云访问，而无需使用 IAM 用户的凭证。
//
// **会话持续时间：**
// 默认情况下，通过 AssumeAgencyWithSAML 创建的临时安全凭证有效期为一小时。您可以使用可选的 duration_seconds 参数或者 SAML 身份验证响应中 SessionNotOnOrAfter 值和 SessionDuration 值来指定会话的持续时间，最终的会话持续时间以三者中较短的一个为准，且会话的持续时间不能超过委托设置的最大会话时长限制。duration_seconds 取值范围是从 900 秒（15 分钟）到该信任委托设置的最大会话持续时长，最大会话持续时长的取值范围可以从 1 小时到 12 小时。注意：委托链会将您的会话持续时间限制为最多一小时，当您使用 AssumeAgency API 操作来进行委托链的切换时，如果您提供了大于一小时的 duration_seconds 参数值，该操作将会失败。
//
// **权限：**
// 调用 AssumeAgencyWithSAML 不需要使用华为云凭证。调用者的身份是通过使用您上传的 SAML 提供商元数据文档中的密钥进行验证的。
// 您可以使用 policy 和 policy_ids 参数传递自定义策略和已有的身份策略来限制本次会话获得的临时安全凭证的权限范围，最终获得临时安全凭证的权限是 policy 和 policy_ids 与信任委托身上附加的身份策略的交集。
//
// **标签：**
// 在信任委托的信任策略中添加了 sts::tagSession 授权项时，您可以配置您的身份提供商，将属性作为会话标签 (Session Tags) 传递到 SAML 断言中。每个会话标签由一个键（Key）和一个值（Value）组成。您最多可以传递 20 个会话标签。纯文本形式的会话标签键不得超过 128 个字符，值不得超过 255 个字符。您也可以传递与信任委托身上标签同名的会话标签，此时会话标签会覆盖具有相同键的信任委托标签。您可以将会话标签设置为可传递的 (Transitive)，可传递的会话标签在角色链期间会持续保留。
//
// **SAML 配置：**
// 在您的应用程序调用 AssumeAgencyWithSAML 之前，您必须使用您的账号在 IAM 中创建 SAML 提供商和信任委托，并在信任委托的信任策略中指定该 SAML 提供商，然后还需要配置您的 SAML 身份提供商以发布 IAM 所需的声明 (Claims)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *StsClient) AssumeAgencyWithSAML(request *model.AssumeAgencyWithSamlRequest) (*model.AssumeAgencyWithSamlResponse, error) {
	requestDef := GenReqDefForAssumeAgencyWithSAML()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssumeAgencyWithSamlResponse), nil
	}
}

// AssumeAgencyWithSAMLInvoker 通过使用SAML协议SSO的信任委托获取临时安全凭证
func (c *StsClient) AssumeAgencyWithSAMLInvoker(request *model.AssumeAgencyWithSamlRequest) *AssumeAgencyWithSAMLInvoker {
	requestDef := GenReqDefForAssumeAgencyWithSAML()
	return &AssumeAgencyWithSAMLInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DecodeAuthorizationMessage 解密鉴权失败的原因
//
// 解密鉴权失败的原因。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *StsClient) DecodeAuthorizationMessage(request *model.DecodeAuthorizationMessageRequest) (*model.DecodeAuthorizationMessageResponse, error) {
	requestDef := GenReqDefForDecodeAuthorizationMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecodeAuthorizationMessageResponse), nil
	}
}

// DecodeAuthorizationMessageInvoker 解密鉴权失败的原因
func (c *StsClient) DecodeAuthorizationMessageInvoker(request *model.DecodeAuthorizationMessageRequest) *DecodeAuthorizationMessageInvoker {
	requestDef := GenReqDefForDecodeAuthorizationMessage()
	return &DecodeAuthorizationMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetCallerIdentity 获取调用者身份信息
//
// 获取调用者（用户，委托等）身份信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *StsClient) GetCallerIdentity(request *model.GetCallerIdentityRequest) (*model.GetCallerIdentityResponse, error) {
	requestDef := GenReqDefForGetCallerIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetCallerIdentityResponse), nil
	}
}

// GetCallerIdentityInvoker 获取调用者身份信息
func (c *StsClient) GetCallerIdentityInvoker(request *model.GetCallerIdentityRequest) *GetCallerIdentityInvoker {
	requestDef := GenReqDefForGetCallerIdentity()
	return &GetCallerIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
