package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoResponseListPluginBasicVoData struct {

	// 插件名
	PluginName *string `json:"pluginName,omitempty"`

	// 展示名
	DisplayName *string `json:"displayName,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 版本说明
	VersionDescription *string `json:"versionDescription,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 版本属性
	VersionAttribution *string `json:"versionAttribution,omitempty"`

	// 唯一ID
	UniqueId *string `json:"uniqueId,omitempty"`

	// 操作人
	OpUser *string `json:"opUser,omitempty"`

	// 操作时间
	OpTime *string `json:"opTime,omitempty"`

	// 组合类型
	PluginCompositionType *string `json:"pluginCompositionType,omitempty"`

	// 属性
	PluginAttribution *string `json:"pluginAttribution,omitempty"`

	// 租户ID
	WorkspaceId *string `json:"workspaceId,omitempty"`

	// 业务类型
	BusinessType *string `json:"businessType,omitempty"`

	// 业务类型展示名
	BusinessTypeDisplayName *string `json:"businessTypeDisplayName,omitempty"`

	// 维护者
	Maintainers *string `json:"maintainers,omitempty"`

	// 图标URL
	IconUrl *string `json:"iconUrl,omitempty"`

	// 引用次数
	ReferCount *int32 `json:"referCount,omitempty"`

	// 使用次数
	UsageCount *int32 `json:"usageCount,omitempty"`

	// 运行属性
	RuntimeAttribution *string `json:"runtimeAttribution,omitempty"`

	// 是否激活
	Active *int32 `json:"active,omitempty"`
}

func (o PageInfoResponseListPluginBasicVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoResponseListPluginBasicVoData struct{}"
	}

	return strings.Join([]string{"PageInfoResponseListPluginBasicVoData", string(data)}, " ")
}
