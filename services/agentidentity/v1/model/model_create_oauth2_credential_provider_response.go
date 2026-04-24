package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOauth2CredentialProviderResponse Response Object
type CreateOauth2CredentialProviderResponse struct {
	CredentialProvider *Oauth2CredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o CreateOauth2CredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOauth2CredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"CreateOauth2CredentialProviderResponse", string(data)}, " ")
}
