package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"net/url"
)

type SMSApiCredentials struct {
	AK string
	SK string
}

func (s *SMSApiCredentials) ProcessAuthParams(client *impl.DefaultHttpClient, region string) auth.ICredential {
	return s
}

func (s *SMSApiCredentials) ProcessAuthRequest(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	var additionalHeaders map[string]string

	// Go版SDK签名接口默认未自动新增Host头域，未保持跟其他语言实现的一致性，新增Host参数，提升安全性
	if _, ok := req.GetHeaderParams()["Host"]; !ok {
		if u, err := url.Parse(req.GetEndpoint()); err == nil {
			req.AddHeaderParam("Host", u.Host)
		}
	}

	sn, err := signer.GetSigner(req.GetSigningAlgorithm())
	if err != nil {
		return nil, err
	}

	reqBuilder := req.Builder()
	additionalHeaders, err = sn.Sign(reqBuilder.Build(), s.AK, s.SK)
	if err != nil {
		return nil, err
	}

	for key, value := range additionalHeaders {
		req.AddHeaderParam(key, value)
	}
	return req, nil
}

type SMSApiCredentialsBuilder struct {
	SMSApiCredentials *SMSApiCredentials
}

func NewSMSApiCredentialsBuilder() *SMSApiCredentialsBuilder {

	return &SMSApiCredentialsBuilder{SMSApiCredentials: &SMSApiCredentials{}}
}

func (builder *SMSApiCredentialsBuilder) WithAk(ak string) *SMSApiCredentialsBuilder {
	builder.SMSApiCredentials.AK = ak
	return builder
}

func (builder *SMSApiCredentialsBuilder) WithSk(sk string) *SMSApiCredentialsBuilder {
	builder.SMSApiCredentials.SK = sk
	return builder
}

func (builder *SMSApiCredentialsBuilder) Build() *SMSApiCredentials {
	return builder.SMSApiCredentials
}
