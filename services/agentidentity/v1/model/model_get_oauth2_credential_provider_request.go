package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetOauth2CredentialProviderRequest Request Object
type GetOauth2CredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o GetOauth2CredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetOauth2CredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"GetOauth2CredentialProviderRequest", string(data)}, " ")
}
