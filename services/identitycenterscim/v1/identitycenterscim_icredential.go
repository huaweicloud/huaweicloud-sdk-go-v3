package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
)

type IdentityCenterSCIMCredentials struct {
}

func (s *IdentityCenterSCIMCredentials) ProcessAuthParams(httpClient *impl.DefaultHttpClient, region string) auth.ICredential {
	return s
}

func (s *IdentityCenterSCIMCredentials) ProcessAuthRequest(httpClient *impl.DefaultHttpClient, httpRequest *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	return httpRequest, nil
}

func NewIdentityCenterSCIMCredentialsBuilder() *IdentityCenterSCIMCredentialsBuilder {
	return &IdentityCenterSCIMCredentialsBuilder{IdentityCenterSCIMCredentials: &IdentityCenterSCIMCredentials{}}
}

type IdentityCenterSCIMCredentialsBuilder struct {
	IdentityCenterSCIMCredentials *IdentityCenterSCIMCredentials
}

func (builder *IdentityCenterSCIMCredentialsBuilder) Build() *IdentityCenterSCIMCredentials {
	return builder.IdentityCenterSCIMCredentials
}
