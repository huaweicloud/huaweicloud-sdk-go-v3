package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmailItemDto struct {

	// 包含电子邮箱地址的字符串
	Value string `json:"value"`

	// 表示电子邮箱类型的字符串
	Type string `json:"type"`

	// 一个布尔值，表示这是否是用户的主电子邮箱
	Primary bool `json:"primary"`
}

func (o EmailItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailItemDto struct{}"
	}

	return strings.Join([]string{"EmailItemDto", string(data)}, " ")
}
