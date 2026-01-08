package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthMethodConfigRequest 认证配置请求。
type AuthMethodConfigRequest struct {

	// 认证配置id。
	Id *string `json:"id,omitempty"`

	// 是否支持多域。
	IsMultiDomainAuthenticateEnabled *bool `json:"is_multi_domain_authenticate_enabled,omitempty"`

	AuthType *AuthTypeEnum `json:"auth_type,omitempty"`

	RadiusGatewayConfig *RadiusGatewayConfig `json:"radius_gateway_config,omitempty"`

	ThirdPartyAuthConfig *ThirdPartyAuthConfig `json:"third_party_auth_config,omitempty"`

	// 应急登录模式。
	EmergencyLoginMode *string `json:"emergency_login_mode,omitempty"`

	Saml2AuthConfig *Saml2AuthConfig `json:"saml2_auth_config,omitempty"`
}

func (o AuthMethodConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthMethodConfigRequest struct{}"
	}

	return strings.Join([]string{"AuthMethodConfigRequest", string(data)}, " ")
}
