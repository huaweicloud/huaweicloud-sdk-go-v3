package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginBasicVo struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 版本说明
	VersionDescription *string `json:"version_description,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 版本属性
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 操作人
	OpUser *string `json:"op_user,omitempty"`

	// 操作时间
	OpTime *string `json:"op_time,omitempty"`

	// 组合类型
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// 属性
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// 租户ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 业务类型
	BusinessType *string `json:"business_type,omitempty"`

	// 业务类型展示名
	BusinessTypeDisplayName *string `json:"business_type_display_name,omitempty"`

	// 维护者
	Maintainers *string `json:"maintainers,omitempty"`

	// 图标URL
	IconUrl *string `json:"icon_url,omitempty"`

	// 引用次数
	ReferCount *int32 `json:"refer_count,omitempty"`

	// 使用次数
	UsageCount *int32 `json:"usage_count,omitempty"`

	// 运行属性
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// 是否激活
	Active *int32 `json:"active,omitempty"`
}

func (o PluginBasicVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginBasicVo struct{}"
	}

	return strings.Join([]string{"PluginBasicVo", string(data)}, " ")
}
