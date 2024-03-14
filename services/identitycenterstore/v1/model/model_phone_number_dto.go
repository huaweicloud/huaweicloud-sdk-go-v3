package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneNumberDto The phone number associated with the user.
type PhoneNumberDto struct {

	// 一个布尔值，表示这是否为用户的主电话号码
	Primary *bool `json:"primary,omitempty"`

	// 表示电话号码类型的字符串
	Type *string `json:"type,omitempty"`

	// 包含电话号码的字符串
	Value *string `json:"value,omitempty"`
}

func (o PhoneNumberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneNumberDto struct{}"
	}

	return strings.Join([]string{"PhoneNumberDto", string(data)}, " ")
}
