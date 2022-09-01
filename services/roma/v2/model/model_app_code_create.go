package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppCodeCreate struct {

	// App Code值  支持英文，+_!@#$%+/=，且只能以英文和+、/开头。
	AppCode string `json:"app_code" xml:"app_code"`
}

func (o AppCodeCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCodeCreate struct{}"
	}

	return strings.Join([]string{"AppCodeCreate", string(data)}, " ")
}
