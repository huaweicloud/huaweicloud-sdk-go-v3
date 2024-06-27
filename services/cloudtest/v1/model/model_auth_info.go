package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthInfo struct {

	// 访问密钥
	AccessKey *string `json:"access_key,omitempty"`

	// 私钥
	SecretKey *string `json:"secret_key,omitempty"`
}

func (o AuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthInfo struct{}"
	}

	return strings.Join([]string{"AuthInfo", string(data)}, " ")
}
