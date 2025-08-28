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
