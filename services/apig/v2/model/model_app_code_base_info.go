package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppCodeBaseInfo struct {
	// App Code值  支持英文，+_!@#$%+/=，且只能以英文和+、/开头，64-180个字符。

	AppCode string `json:"app_code"`
	// 编号

	Id *string `json:"id,omitempty"`
	// 应用编号

	AppId *string `json:"app_id,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
}

func (o AppCodeBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCodeBaseInfo struct{}"
	}

	return strings.Join([]string{"AppCodeBaseInfo", string(data)}, " ")
}
