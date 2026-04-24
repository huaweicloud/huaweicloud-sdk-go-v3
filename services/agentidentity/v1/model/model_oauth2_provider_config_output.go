package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2ProviderConfigOutput OAuth2 provider config output (UNION). Only one of the following members can be specified.
type Oauth2ProviderConfigOutput struct {
	MicrosoftOauth2ProviderConfig *MicrosoftOauth2ProviderConfigOutput `json:"microsoft_oauth2_provider_config,omitempty"`

	GoogleOauth2ProviderConfig *GoogleOauth2ProviderConfigOutput `json:"google_oauth2_provider_config,omitempty"`

	GithubOauth2ProviderConfig *GithubOauth2ProviderConfigOutput `json:"github_oauth2_provider_config,omitempty"`

	CustomOauth2ProviderConfig *CustomOauth2ProviderConfigOutput `json:"custom_oauth2_provider_config,omitempty"`
}

func (o Oauth2ProviderConfigOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2ProviderConfigOutput struct{}"
	}

	return strings.Join([]string{"Oauth2ProviderConfigOutput", string(data)}, " ")
}
