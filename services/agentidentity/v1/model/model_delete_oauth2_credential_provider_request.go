package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOauth2CredentialProviderRequest Request Object
type DeleteOauth2CredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o DeleteOauth2CredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOauth2CredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"DeleteOauth2CredentialProviderRequest", string(data)}, " ")
}
