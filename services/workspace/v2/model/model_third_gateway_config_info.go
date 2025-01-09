package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdGatewayConfigInfo 第三方网关配置信息。
type ThirdGatewayConfigInfo struct {

	// 第三方登录url
	LoginUrl string `json:"login_url"`

	// 第三方登出url
	LogoutUrl string `json:"logout_url"`

	// 第三方网关token校验地址
	TokenUrl string `json:"token_url"`

	// 校验证书url
	VerificationCipherUrl string `json:"verification_cipher_url"`

	// 第三方证书
	CertContent string `json:"cert_content"`
}

func (o ThirdGatewayConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdGatewayConfigInfo struct{}"
	}

	return strings.Join([]string{"ThirdGatewayConfigInfo", string(data)}, " ")
}
