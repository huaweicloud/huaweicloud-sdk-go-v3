package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneCreateIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty" xml:"identity_provider"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneCreateIdentityProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderResponse", string(data)}, " ")
}
