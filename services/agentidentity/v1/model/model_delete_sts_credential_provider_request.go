package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStsCredentialProviderRequest Request Object
type DeleteStsCredentialProviderRequest struct {

	// The name of the credential provider.
	CredentialProviderName string `json:"credential_provider_name"`
}

func (o DeleteStsCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStsCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"DeleteStsCredentialProviderRequest", string(data)}, " ")
}
