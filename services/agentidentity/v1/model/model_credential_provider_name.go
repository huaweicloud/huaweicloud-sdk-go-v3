package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CredentialProviderName The name of the credential provider.
type CredentialProviderName struct {
}

func (o CredentialProviderName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CredentialProviderName struct{}"
	}

	return strings.Join([]string{"CredentialProviderName", string(data)}, " ")
}
