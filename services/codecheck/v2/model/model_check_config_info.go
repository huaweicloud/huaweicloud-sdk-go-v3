package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckConfigInfo struct {

	// 检查参数名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 检查参数对应的key值
	CfgKey *string `json:"cfg_key,omitempty" xml:"cfg_key"`

	// 检查参数默认值
	DefaultValue *string `json:"default_value,omitempty" xml:"default_value"`

	// 检查参数可选项
	OptionValue *string `json:"option_value,omitempty" xml:"option_value"`

	// 0：非必填，1：必填
	IsRequired *int32 `json:"is_required,omitempty" xml:"is_required"`

	// 检查参数说明
	Description *string `json:"description,omitempty" xml:"description"`

	// 参数类型，0：文本，2：有可选项
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 参数状态，on：启用，off：未启用
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o CheckConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConfigInfo struct{}"
	}

	return strings.Join([]string{"CheckConfigInfo", string(data)}, " ")
}
