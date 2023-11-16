package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoFullStagePluginsItemList struct {

	// 插件列表
	PluginsList *[]FullStagePluginsRelationVoPluginsList `json:"plugins_list,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 业务类型
	BusinessType *string `json:"business_type,omitempty"`

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 条件
	Conditions *[]string `json:"conditions,omitempty"`

	// 额外属性
	Addables *[]FullStagePluginsRelationVoAddables `json:"addables,omitempty"`

	// 是否可编辑
	Editable *bool `json:"editable,omitempty"`

	// 是否可移除
	Removable *bool `json:"removable,omitempty"`

	// 是否可复制
	Cloneable *bool `json:"cloneable,omitempty"`

	// 禁用
	Disabled *bool `json:"disabled,omitempty"`
}

func (o FullStagePluginsRelationVoFullStagePluginsItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoFullStagePluginsItemList struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoFullStagePluginsItemList", string(data)}, " ")
}
