package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdGatewayInfo 第三方网关信息
type ThirdGatewayInfo struct {

	// 第三方登录url
	LoginUrl *string `json:"login_url,omitempty"`

	// 第三方登出url
	LogoutUrl *string `json:"logout_url,omitempty"`

	// 第三方网关token校验url
	TokenUrl *string `json:"token_url,omitempty"`

	// 校验证书url
	VerificationCipherUrl *string `json:"verification_cipher_url,omitempty"`

	// 第三方网关认证证书
	CertContent *string `json:"cert_content,omitempty"`
}

func (o ThirdGatewayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdGatewayInfo struct{}"
	}

	return strings.Join([]string{"ThirdGatewayInfo", string(data)}, " ")
}
