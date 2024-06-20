package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginBasicDto struct {

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 图标URL
	IconUrl *string `json:"icon_url,omitempty"`

	// 运行属性
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 展示名
	DisplayName string `json:"display_name"`

	// 业务类型
	BusinessType string `json:"business_type"`

	// 业务类型展示名
	BusinessTypeDisplayName string `json:"business_type_display_name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否私有
	IsPrivate *int32 `json:"is_private,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 维护者
	Maintainers *string `json:"maintainers,omitempty"`

	// 组合插件类型
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`
}

func (o PluginBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginBasicDto struct{}"
	}

	return strings.Join([]string{"PluginBasicDto", string(data)}, " ")
}
