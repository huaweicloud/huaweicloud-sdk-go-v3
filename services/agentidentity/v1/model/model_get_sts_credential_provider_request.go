package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetStsCredentialProviderRequest Request Object
type GetStsCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o GetStsCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetStsCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"GetStsCredentialProviderRequest", string(data)}, " ")
}
