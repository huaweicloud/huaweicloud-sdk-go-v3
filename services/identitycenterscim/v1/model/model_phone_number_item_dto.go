package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PhoneNumberItemDto struct {

	// 包含电话号码的字符串
	Value *string `json:"value,omitempty"`

	// 表示电话号码类型的字符串
	Type *string `json:"type,omitempty"`

	// 一个布尔值，表示这是否是用户的主电话号码
	Primary *bool `json:"primary,omitempty"`
}

func (o PhoneNumberItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneNumberItemDto struct{}"
	}

	return strings.Join([]string{"PhoneNumberItemDto", string(data)}, " ")
}
