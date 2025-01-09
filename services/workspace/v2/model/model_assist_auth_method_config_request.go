package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssistAuthMethodConfigRequest 辅助认证策略请求
type AssistAuthMethodConfigRequest struct {
	AuthType *AuthAssistEnum `json:"auth_type,omitempty"`

	OtpConfigInfo *OtpConfigInfo `json:"otp_config_info,omitempty"`

	RadiusAuthConfig *RadiusAuthConfig `json:"radius_auth_config,omitempty"`

	RadiusGatewayConfig *RadiusGatewayConfig `json:"radius_gateway_config,omitempty"`
}

func (o AssistAuthMethodConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssistAuthMethodConfigRequest struct{}"
	}

	return strings.Join([]string{"AssistAuthMethodConfigRequest", string(data)}, " ")
}
