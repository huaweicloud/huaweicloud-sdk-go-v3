package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateIdentityProviderRequestBody
type KeystoneCreateIdentityProviderRequestBody struct {
	IdentityProvider *IdentityproviderOption `json:"identity_provider"`
}

func (o KeystoneCreateIdentityProviderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderRequestBody", string(data)}, " ")
}
