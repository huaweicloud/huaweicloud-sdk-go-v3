package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaslPlainAuthInfo sasl plain 认证信息
type SaslPlainAuthInfo struct {

	// 用户名。支持大小写字母、数字、“.”、“-” 、“_”
	Username *string `json:"username,omitempty"`

	// 密码
	Password *string `json:"password,omitempty"`
}

func (o SaslPlainAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaslPlainAuthInfo struct{}"
	}

	return strings.Join([]string{"SaslPlainAuthInfo", string(data)}, " ")
}
