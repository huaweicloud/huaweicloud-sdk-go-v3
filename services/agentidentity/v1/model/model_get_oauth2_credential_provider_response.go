package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetOauth2CredentialProviderResponse Response Object
type GetOauth2CredentialProviderResponse struct {
	CredentialProvider *Oauth2CredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o GetOauth2CredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetOauth2CredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"GetOauth2CredentialProviderResponse", string(data)}, " ")
}
