package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Oauth2CredentialProvider struct {

	// The name of the credential provider.
	Name string `json:"name"`

	// 凭证提供者的唯一资源名称（URN）。
	Urn string `json:"urn"`

	CredentialProviderVendor *CredentialProviderVendor `json:"credential_provider_vendor"`

	ClientSecret *Secret `json:"client_secret"`

	// OAuth2 提供方认证完成后跳转回调的 URL。
	CallbackUrl string `json:"callback_url"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	Oauth2ProviderConfigOutput *Oauth2ProviderConfigOutput `json:"oauth2_provider_config_output,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o Oauth2CredentialProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2CredentialProvider struct{}"
	}

	return strings.Join([]string{"Oauth2CredentialProvider", string(data)}, " ")
}
