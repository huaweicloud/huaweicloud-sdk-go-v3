package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionExtendProp struct {

	// API 选项
	ApiOptions *string `json:"api_options,omitempty"`

	// API 类型
	ApiType *string `json:"api_type,omitempty"`

	// 选项
	Options *string `json:"options,omitempty"`

	// 禁用条件
	DisabledConditions *string `json:"disabled_conditions,omitempty"`

	// 可见条件
	VisibleConditions *string `json:"visible_conditions,omitempty"`
}

func (o ExtensionExtendProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionExtendProp struct{}"
	}

	return strings.Join([]string{"ExtensionExtendProp", string(data)}, " ")
}
