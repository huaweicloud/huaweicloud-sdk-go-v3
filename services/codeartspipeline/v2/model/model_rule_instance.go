package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstance struct {

	// 规则实例ID
	Id string `json:"id"`

	// 规则类型ID
	Type string `json:"type"`

	// 规则名称
	Name string `json:"name"`

	// 规则版本
	Version string `json:"version"`

	// 插件ID
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件版本号
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 是否生效
	IsValid bool `json:"is_valid"`

	// 是否可编辑
	Editable *bool `json:"editable,omitempty"`

	// 规则属性列表
	Content []RuleInstanceContent `json:"content"`

	Parent *RuleSet `json:"parent,omitempty"`
}

func (o RuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstance struct{}"
	}

	return strings.Join([]string{"RuleInstance", string(data)}, " ")
}
