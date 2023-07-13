package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthInfo struct {

	// Base64加密的认证信息
	Auth string `json:"auth"`
}

func (o AuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthInfo struct{}"
	}

	return strings.Join([]string{"AuthInfo", string(data)}, " ")
}
