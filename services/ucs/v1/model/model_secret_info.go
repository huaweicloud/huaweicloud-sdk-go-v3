package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecretInfo struct {

	// 使用的认证模式
	AuthMode *string `json:"authMode,omitempty"`

	// 存储了实际认证凭据的Secret
	Secret *interface{} `json:"secret,omitempty"`
}

func (o SecretInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretInfo struct{}"
	}

	return strings.Join([]string{"SecretInfo", string(data)}, " ")
}
