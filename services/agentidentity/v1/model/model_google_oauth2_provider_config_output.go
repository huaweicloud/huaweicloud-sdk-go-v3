package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GoogleOauth2ProviderConfigOutput Output configuration for a Google OAuth2 provider.
type GoogleOauth2ProviderConfigOutput struct {
	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`

	// Client ID for OAuth2 application.
	ClientId *string `json:"client_id,omitempty"`
}

func (o GoogleOauth2ProviderConfigOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GoogleOauth2ProviderConfigOutput struct{}"
	}

	return strings.Join([]string{"GoogleOauth2ProviderConfigOutput", string(data)}, " ")
}
