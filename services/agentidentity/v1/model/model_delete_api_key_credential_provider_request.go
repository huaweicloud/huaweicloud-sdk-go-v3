package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiKeyCredentialProviderRequest Request Object
type DeleteApiKeyCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o DeleteApiKeyCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiKeyCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"DeleteApiKeyCredentialProviderRequest", string(data)}, " ")
}
