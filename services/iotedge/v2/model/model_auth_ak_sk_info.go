package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthAkSkInfo struct {
	// 鉴权秘钥

	Secret *string `json:"secret,omitempty"`
}

func (o AuthAkSkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthAkSkInfo struct{}"
	}

	return strings.Join([]string{"AuthAkSkInfo", string(data)}, " ")
}
