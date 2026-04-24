package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2ProviderConfigInput OAuth2 provider config input (UNION). Only one of the following members can be specified.
type Oauth2ProviderConfigInput struct {
	MicrosoftOauth2ProviderConfig *MicrosoftOauth2ProviderConfigInput `json:"microsoft_oauth2_provider_config,omitempty"`

	GoogleOauth2ProviderConfig *GoogleOauth2ProviderConfigInput `json:"google_oauth2_provider_config,omitempty"`

	GithubOauth2ProviderConfig *GithubOauth2ProviderConfigInput `json:"github_oauth2_provider_config,omitempty"`

	CustomOauth2ProviderConfig *CustomOauth2ProviderConfigInput `json:"custom_oauth2_provider_config,omitempty"`
}

func (o Oauth2ProviderConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2ProviderConfigInput struct{}"
	}

	return strings.Join([]string{"Oauth2ProviderConfigInput", string(data)}, " ")
}
