package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDtoInputInfo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 输入类型
	Type *string `json:"type,omitempty"`

	// 验证
	Validation *string `json:"validation,omitempty"`

	// 样式信息
	LayoutContent *string `json:"layout_content,omitempty"`
}

func (o PluginDtoInputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoInputInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoInputInfo", string(data)}, " ")
}
