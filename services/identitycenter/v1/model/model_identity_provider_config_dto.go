package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentityProviderConfigDto 身份提供商
type IdentityProviderConfigDto struct {

	// 身份提供商issuer
	IssuerUrl string `json:"issuer_url"`

	// 身份提供商元数据
	MetadataUrl string `json:"metadata_url"`

	// 身份提供商远程登录链接
	RemoteLoginUrl string `json:"remote_login_url"`

	// 身份提供商远程登出链接
	RemoteLogoutUrl string `json:"remote_logout_url"`
}

func (o IdentityProviderConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityProviderConfigDto struct{}"
	}

	return strings.Join([]string{"IdentityProviderConfigDto", string(data)}, " ")
}
