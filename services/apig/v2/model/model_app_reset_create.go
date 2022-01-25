package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppResetCreate struct {
	// 密钥支持英文，数字，“_”,“-”,“!”,“@”,“#”,“$”,“%”,且只能以英文或数字开头，8 ~ 64个字符。用户自定义APP的密钥需要开启配额开关

	AppSecret *string `json:"app_secret,omitempty"`
}

func (o AppResetCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResetCreate struct{}"
	}

	return strings.Join([]string{"AppResetCreate", string(data)}, " ")
}
