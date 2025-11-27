package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthInfo 用户认证信息
type AuthInfo struct {

	// 客户端证书
	ClientCertificateData *string `json:"client-certificate-data,omitempty"`

	// 包含来自TLS客户端密钥文件的PEM编码数据
	ClientKeyData *string `json:"client-key-data,omitempty"`

	// 身份验证令牌
	Token *string `json:"token,omitempty"`
}

func (o AuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthInfo struct{}"
	}

	return strings.Join([]string{"AuthInfo", string(data)}, " ")
}
