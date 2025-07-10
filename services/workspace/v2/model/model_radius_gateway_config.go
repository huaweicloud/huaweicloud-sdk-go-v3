package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RadiusGatewayConfig 短信辅助认证配置。
type RadiusGatewayConfig struct {

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// 用户名。
	AppId *string `json:"app_id,omitempty"`

	// 密码。
	Password *string `json:"password,omitempty"`

	// 获取token地址。
	TokenUrl *string `json:"token_url,omitempty"`

	// 获取验证码地址。
	VerificationCipherUrl *string `json:"verification_cipher_url,omitempty"`

	// 证书内容（PEM）。
	CertContent *string `json:"cert_content,omitempty"`
}

func (o RadiusGatewayConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadiusGatewayConfig struct{}"
	}

	return strings.Join([]string{"RadiusGatewayConfig", string(data)}, " ")
}
