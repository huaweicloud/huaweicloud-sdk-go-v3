package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EmailDto The email address associated with the user.
type EmailDto struct {

	// 一个布尔值，表示这是否是用户的主电子邮箱
	Primary bool `json:"primary"`

	// 表示电子邮箱类型的字符串
	Type string `json:"type"`

	// 包含电子邮箱地址的字符串
	Value string `json:"value"`
}

func (o EmailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailDto struct{}"
	}

	return strings.Join([]string{"EmailDto", string(data)}, " ")
}
