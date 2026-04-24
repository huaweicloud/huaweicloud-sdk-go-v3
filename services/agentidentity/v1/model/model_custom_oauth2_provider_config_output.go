package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomOauth2ProviderConfigOutput Output configuration for a custom OAuth2 provider.
type CustomOauth2ProviderConfigOutput struct {
	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`

	// Client ID for OAuth2 application.
	ClientId *string `json:"client_id,omitempty"`
}

func (o CustomOauth2ProviderConfigOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomOauth2ProviderConfigOutput struct{}"
	}

	return strings.Join([]string{"CustomOauth2ProviderConfigOutput", string(data)}, " ")
}
