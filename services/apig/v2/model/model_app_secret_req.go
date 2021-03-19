package model

import (
	"encoding/json"

	"strings"
)

type AppSecretReq struct {
	// 密钥支持英文，数字，“_”,“-”,“_”,“!”,“@”,“#”,“$”,“%”,且只能以英文或数字开头，8 ~ 64个字符。用户自定义APP的密钥需要开启配额开关

	AppSecret *string `json:"app_secret,omitempty"`
}

func (o AppSecretReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AppSecretReq struct{}"
	}

	return strings.Join([]string{"AppSecretReq", string(data)}, " ")
}
