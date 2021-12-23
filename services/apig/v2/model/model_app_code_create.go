package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppCodeCreate struct {
	// App Code值  支持英文，+_!@#$%+/=，且只能以英文和+、/开头，64-180个字符。

	AppCode string `json:"app_code"`
}

func (o AppCodeCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCodeCreate struct{}"
	}

	return strings.Join([]string{"AppCodeCreate", string(data)}, " ")
}
