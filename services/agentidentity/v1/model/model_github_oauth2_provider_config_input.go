package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GithubOauth2ProviderConfigInput Configuration settings for GitHub OAuth2 provider integration.
type GithubOauth2ProviderConfigInput struct {

	// Client ID for OAuth2 application.
	ClientId string `json:"client_id"`

	// Client secret for OAuth2 application.
	ClientSecret string `json:"client_secret"`
}

func (o GithubOauth2ProviderConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GithubOauth2ProviderConfigInput struct{}"
	}

	return strings.Join([]string{"GithubOauth2ProviderConfigInput", string(data)}, " ")
}
