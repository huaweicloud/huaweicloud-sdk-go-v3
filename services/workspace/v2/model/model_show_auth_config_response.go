package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthConfigResponse Response Object
type ShowAuthConfigResponse struct {

	// 认证配置ID。
	Id *string `json:"id,omitempty"`

	// 认证类型 LOCAL_PASSWORD：本地密码认证模式 KERBEROS：Windows AD认证模式 LDAP：第三方LDAP模式 CLIENT_TOKEN：金审UKEY客户端Token认证模式 OAUTH2：第三方单点登录模式
	AuthType *string `json:"auth_type,omitempty"`

	// 当前状态。
	Enable *bool `json:"enable,omitempty"`

	// 当前状态。
	IsMultiDomainAuthenticateEnabled *bool `json:"is_multi_domain_authenticate_enabled,omitempty"`

	RadiusGatewayConfig *RadiusGatewayConfigInfo `json:"radius_gateway_config,omitempty"`

	// 第三方认证接口配置信息。
	ThirdPartyAuthConfig *[]ThirdPartyAuthConfig `json:"third_party_auth_config,omitempty"`

	// 应急登录模式。
	EmergencyLoginMode *string `json:"emergency_login_mode,omitempty"`

	Saml2AuthConfig *Saml2AuthConfig `json:"saml2_auth_config,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o ShowAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAuthConfigResponse", string(data)}, " ")
}
