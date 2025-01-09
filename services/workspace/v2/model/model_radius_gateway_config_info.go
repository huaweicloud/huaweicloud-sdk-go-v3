package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RadiusGatewayConfigInfo 短信辅助认证配置
type RadiusGatewayConfigInfo struct {

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 用户名
	AppId *string `json:"app_id,omitempty"`

	// 证书域名
	CertDomainName *string `json:"cert_domain_name,omitempty"`

	// 获取token地址
	TokenUrl *string `json:"token_url,omitempty"`

	// 获取验证码地址
	VerificationCipherUrl *string `json:"verification_cipher_url,omitempty"`

	// 认证类型
	AuthType *string `json:"auth_type,omitempty"`

	// 辅助认证类型
	AssistAuthType *string `json:"assist_auth_type,omitempty"`

	// 过期时间
	Expiration *string `json:"expiration,omitempty"`
}

func (o RadiusGatewayConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadiusGatewayConfigInfo struct{}"
	}

	return strings.Join([]string{"RadiusGatewayConfigInfo", string(data)}, " ")
}
