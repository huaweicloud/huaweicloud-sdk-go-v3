package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GithubOauth2ProviderConfigOutput Output configuration for a GitHub OAuth2 provider.
type GithubOauth2ProviderConfigOutput struct {
	Oauth2Discovery *Oauth2Discovery `json:"oauth2_discovery"`

	// Client ID for OAuth2 application.
	ClientId *string `json:"client_id,omitempty"`
}

func (o GithubOauth2ProviderConfigOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GithubOauth2ProviderConfigOutput struct{}"
	}

	return strings.Join([]string{"GithubOauth2ProviderConfigOutput", string(data)}, " ")
}
