package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
)

type IdentityCenterPortalAPICredentials struct {
}

func (s *IdentityCenterPortalAPICredentials) ProcessAuthParams(httpClient *impl.DefaultHttpClient, region string) (auth.ICredential, error) {
	return s, nil
}

func (s *IdentityCenterPortalAPICredentials) ProcessAuthRequest(httpClient *impl.DefaultHttpClient, httpRequest *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	return httpRequest, nil
}

func NewIdentityCenterPortalAPICredentialsBuilder() *IdentityCenterPortalAPICredentialsBuilder {
	return &IdentityCenterPortalAPICredentialsBuilder{IdentityCenterPortalAPICredentials: &IdentityCenterPortalAPICredentials{}}
}

type IdentityCenterPortalAPICredentialsBuilder struct {
	IdentityCenterPortalAPICredentials *IdentityCenterPortalAPICredentials
}

func (builder *IdentityCenterPortalAPICredentialsBuilder) Build() *IdentityCenterPortalAPICredentials {
	return builder.IdentityCenterPortalAPICredentials
}
