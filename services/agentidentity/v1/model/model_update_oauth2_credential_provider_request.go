package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOauth2CredentialProviderRequest Request Object
type UpdateOauth2CredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`

	Body *UpdateOauth2CredentialProviderReqBody `json:"body,omitempty"`
}

func (o UpdateOauth2CredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOauth2CredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"UpdateOauth2CredentialProviderRequest", string(data)}, " ")
}
