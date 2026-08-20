package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionRadioOption 单选选项
type ExtensionRadioOption struct {

	// 选项显示名
	DisplayName *string `json:"displayName,omitempty"`

	// 选项值
	Value *string `json:"value,omitempty"`
}

func (o ExtensionRadioOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionRadioOption struct{}"
	}

	return strings.Join([]string{"ExtensionRadioOption", string(data)}, " ")
}
