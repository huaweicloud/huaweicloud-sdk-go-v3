package model

import (
	"encoding/json"

	"strings"
)

// 修改独享模式域名的请求
type UpdatePremiumHostRequestBody struct {
	// 是否使用代理

	Proxy *bool `json:"proxy,omitempty"`
	// 证书ID

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名称

	Certificatename *string `json:"certificatename,omitempty"`
	// 支持最低的TLS版本

	Tls *string `json:"tls,omitempty"`
	// 加密套件代码

	Cipher *string `json:"cipher,omitempty"`
}

func (o UpdatePremiumHostRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostRequestBody", string(data)}, " ")
}
