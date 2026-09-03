package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdentityProviderListSummary struct {

	// Identity provider code.
	IdentityProvider string `json:"identity_provider"`

	// Display name of the identity provider.
	DisplayName string `json:"display_name"`

	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`
}

func (o IdentityProviderListSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityProviderListSummary struct{}"
	}

	return strings.Join([]string{"IdentityProviderListSummary", string(data)}, " ")
}
