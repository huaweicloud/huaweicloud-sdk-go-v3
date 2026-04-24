package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomOauth2ProviderConfigInput Input configuration for a custom OAuth2 provider.
type CustomOauth2ProviderConfigInput struct {

	// Client ID for OAuth2 application.
	ClientId string `json:"client_id"`

	// Client secret for OAuth2 application.
	ClientSecret string `json:"client_secret"`

	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`
}

func (o CustomOauth2ProviderConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomOauth2ProviderConfigInput struct{}"
	}

	return strings.Join([]string{"CustomOauth2ProviderConfigInput", string(data)}, " ")
}
