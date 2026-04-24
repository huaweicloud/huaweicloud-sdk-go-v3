package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GoogleOauth2ProviderConfigInput Configuration settings for Google OAuth2 provider integration.
type GoogleOauth2ProviderConfigInput struct {

	// Client ID for OAuth2 application.
	ClientId string `json:"client_id"`

	// Client secret for OAuth2 application.
	ClientSecret string `json:"client_secret"`
}

func (o GoogleOauth2ProviderConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GoogleOauth2ProviderConfigInput struct{}"
	}

	return strings.Join([]string{"GoogleOauth2ProviderConfigInput", string(data)}, " ")
}
