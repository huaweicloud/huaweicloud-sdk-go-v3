package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetApiKeyCredentialProviderRequest Request Object
type GetApiKeyCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o GetApiKeyCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetApiKeyCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"GetApiKeyCredentialProviderRequest", string(data)}, " ")
}
