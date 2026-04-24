package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateOauth2CredentialProviderReqBody struct {
	Oauth2ProviderConfigInput *Oauth2ProviderConfigInput `json:"oauth2_provider_config_input,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateOauth2CredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOauth2CredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"UpdateOauth2CredentialProviderReqBody", string(data)}, " ")
}
