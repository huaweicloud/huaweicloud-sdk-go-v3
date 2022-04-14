package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckConfigInfo struct {
	// 检查参数名称

	Name *string `json:"name,omitempty"`
	// 检查参数对应的key值

	CfgKey *string `json:"cfg_key,omitempty"`
	// 检查参数默认值

	DefaultValue *string `json:"default_value,omitempty"`
	// 检查参数可选项

	OptionValue *string `json:"option_value,omitempty"`
	// 0：非必填，1：必填

	IsRequired *int32 `json:"is_required,omitempty"`
	// 检查参数说明

	Description *string `json:"description,omitempty"`
	// 参数类型，0：文本，2：有可选项

	Type *int32 `json:"type,omitempty"`
	// 参数状态，on：启用，off：未启用

	Status *string `json:"status,omitempty"`
}

func (o CheckConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConfigInfo struct{}"
	}

	return strings.Join([]string{"CheckConfigInfo", string(data)}, " ")
}
