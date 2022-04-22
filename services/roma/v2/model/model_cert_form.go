package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertForm struct {

	// 证书内容
	CertContent string `json:"cert_content"`

	// 证书名称。长度为4 ~ 50位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 私钥内容
	PrivateKey string `json:"private_key"`
}

func (o CertForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertForm struct{}"
	}

	return strings.Join([]string{"CertForm", string(data)}, " ")
}
