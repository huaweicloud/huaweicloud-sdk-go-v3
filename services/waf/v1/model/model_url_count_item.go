package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlCountItem struct {
	// 攻击类型

	Key *string `json:"key,omitempty"`
	// 数量

	Num *int32 `json:"num,omitempty"`
	// 域名

	Host *string `json:"host,omitempty"`
}

func (o UrlCountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlCountItem struct{}"
	}

	return strings.Join([]string{"UrlCountItem", string(data)}, " ")
}
