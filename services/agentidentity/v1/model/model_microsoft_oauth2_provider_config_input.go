package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicrosoftOauth2ProviderConfigInput Configuration settings for Microsoft OAuth2 provider integration.
type MicrosoftOauth2ProviderConfigInput struct {

	// Client ID for OAuth2 application.
	ClientId string `json:"client_id"`

	// Client secret for OAuth2 application.
	ClientSecret string `json:"client_secret"`

	// The tenant ID for the Microsoft OAuth2 provider.
	TenantId string `json:"tenant_id"`
}

func (o MicrosoftOauth2ProviderConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicrosoftOauth2ProviderConfigInput struct{}"
	}

	return strings.Join([]string{"MicrosoftOauth2ProviderConfigInput", string(data)}, " ")
}
