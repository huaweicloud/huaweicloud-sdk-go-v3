package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改云模式域名的请求体
type UpdateHostRequestBody struct {
	// 是否使用代理

	Proxy *bool `json:"proxy,omitempty"`
	// 证书ID

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名称

	Certificatename *string `json:"certificatename,omitempty"`
	// 独享模式回源服务器配置

	Server *[]PremiumWafServer `json:"server,omitempty"`
	// 支持最低的TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）

	Tls *string `json:"tls,omitempty"`
	// 加密套件代码（cipher_default/cipher_1/cipher_2）

	Cipher *string `json:"cipher,omitempty"`
}

func (o UpdateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostRequestBody", string(data)}, " ")
}
