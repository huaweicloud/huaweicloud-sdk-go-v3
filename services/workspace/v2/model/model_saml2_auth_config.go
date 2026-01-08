package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Saml2AuthConfig SAML2认证配置信息。
type Saml2AuthConfig struct {

	// 身份提供者名称。
	IdentityProvider *string `json:"identity_provider,omitempty"`

	// 接入服务器地址。
	AccessServerAddress *string `json:"access_server_address,omitempty"`

	// 唯一用户标识符。
	UniqueUserIdentifier *string `json:"unique_user_identifier,omitempty"`

	IdpMetadataInfo *IdpMetadataInfo `json:"idp_metadata_info,omitempty"`
}

func (o Saml2AuthConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Saml2AuthConfig struct{}"
	}

	return strings.Join([]string{"Saml2AuthConfig", string(data)}, " ")
}
