package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiKeyCredentialProviderRequest Request Object
type UpdateApiKeyCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`

	Body *UpdateApiKeyCredentialProviderReqBody `json:"body,omitempty"`
}

func (o UpdateApiKeyCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiKeyCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"UpdateApiKeyCredentialProviderRequest", string(data)}, " ")
}
