package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthMethodConfigRequest 认证配置请求
type AuthMethodConfigRequest struct {
	AuthType *AuthTypeEnum `json:"auth_type,omitempty"`

	RadiusGatewayConfig *RadiusGatewayConfig `json:"radius_gateway_config,omitempty"`

	ThirdPartyAuthConfig *ThirdPartyAuthConfig `json:"third_party_auth_config,omitempty"`

	// 应急登录模式。
	EmergencyLoginMode *string `json:"emergency_login_mode,omitempty"`
}

func (o AuthMethodConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthMethodConfigRequest struct{}"
	}

	return strings.Join([]string{"AuthMethodConfigRequest", string(data)}, " ")
}
