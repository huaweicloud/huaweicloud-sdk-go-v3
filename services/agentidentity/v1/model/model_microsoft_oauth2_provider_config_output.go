package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicrosoftOauth2ProviderConfigOutput Output configuration for a Microsoft OAuth2 provider.
type MicrosoftOauth2ProviderConfigOutput struct {
	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`

	// Client ID for OAuth2 application.
	ClientId *string `json:"client_id,omitempty"`
}

func (o MicrosoftOauth2ProviderConfigOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicrosoftOauth2ProviderConfigOutput struct{}"
	}

	return strings.Join([]string{"MicrosoftOauth2ProviderConfigOutput", string(data)}, " ")
}
