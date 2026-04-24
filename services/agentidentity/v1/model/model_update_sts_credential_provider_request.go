package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStsCredentialProviderRequest Request Object
type UpdateStsCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`

	Body *UpdateStsCredentialProviderReqBody `json:"body,omitempty"`
}

func (o UpdateStsCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStsCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"UpdateStsCredentialProviderRequest", string(data)}, " ")
}
