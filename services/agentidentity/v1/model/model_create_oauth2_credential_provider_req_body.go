package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOauth2CredentialProviderReqBody struct {

	// The name of the credential provider.
	Name string `json:"name"`

	CredentialProviderVendor *CredentialProviderVendor `json:"credential_provider_vendor"`

	Oauth2ProviderConfigInput *Oauth2ProviderConfigInput `json:"oauth2_provider_config_input"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateOauth2CredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOauth2CredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"CreateOauth2CredentialProviderReqBody", string(data)}, " ")
}
