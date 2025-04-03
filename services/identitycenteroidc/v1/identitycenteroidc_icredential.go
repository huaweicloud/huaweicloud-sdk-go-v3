package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
)

type IdentityCenterOIDCCredentials struct {
}

func (s *IdentityCenterOIDCCredentials) ProcessAuthParams(httpClient *impl.DefaultHttpClient, region string) auth.ICredential {
	return s
}

func (s *IdentityCenterOIDCCredentials) ProcessAuthRequest(httpClient *impl.DefaultHttpClient, httpRequest *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	return httpRequest, nil
}

func NewIdentityCenterOIDCCredentialsBuilder() *IdentityCenterOIDCCredentialsBuilder {
	return &IdentityCenterOIDCCredentialsBuilder{IdentityCenterOIDCCredentials: &IdentityCenterOIDCCredentials{}}
}

type IdentityCenterOIDCCredentialsBuilder struct {
	IdentityCenterOIDCCredentials *IdentityCenterOIDCCredentials
}

func (builder *IdentityCenterOIDCCredentialsBuilder) Build() *IdentityCenterOIDCCredentials {
	return builder.IdentityCenterOIDCCredentials
}
