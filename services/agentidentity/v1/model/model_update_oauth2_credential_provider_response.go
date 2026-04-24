package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOauth2CredentialProviderResponse Response Object
type UpdateOauth2CredentialProviderResponse struct {
	CredentialProvider *Oauth2CredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o UpdateOauth2CredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOauth2CredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"UpdateOauth2CredentialProviderResponse", string(data)}, " ")
}
