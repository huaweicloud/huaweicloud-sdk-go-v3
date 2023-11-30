package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoBusinessTypeDefinitionVoPluginsList struct {

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 组名
	GroupName *string `json:"group_name,omitempty"`

	// 组类型
	GroupType *string `json:"group_type,omitempty"`

	// 属性
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// 组合插件
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// 运行属性
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// 基础插件列表
	AllSteps *[]FullStagePluginsRelationVoAllSteps `json:"all_steps,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 版本属性
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// 图标URL
	IconUrl *string `json:"icon_url,omitempty"`

	// 可编辑
	MultiStepEditable *int32 `json:"multi_step_editable,omitempty"`

	// 地址
	Location *string `json:"location,omitempty"`

	// 发布商ID
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// 版本
	ManifestVersion *string `json:"manifest_version,omitempty"`
}

func (o PageInfoBusinessTypeDefinitionVoPluginsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoBusinessTypeDefinitionVoPluginsList struct{}"
	}

	return strings.Join([]string{"PageInfoBusinessTypeDefinitionVoPluginsList", string(data)}, " ")
}
