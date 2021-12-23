package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 鉴权配置
type AppAuth struct {
	// 开启或关闭URL鉴权

	Enable *bool `json:"enable,omitempty"`
	// 接入RTC建链认证时的signature的有效期。单位：秒。默认300秒。signature由app_key生成

	Expire *int32 `json:"expire,omitempty"`
	// app鉴权秘钥

	AppKey *string `json:"app_key,omitempty"`
	// app鉴权的更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o AppAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAuth struct{}"
	}

	return strings.Join([]string{"AppAuth", string(data)}, " ")
}
