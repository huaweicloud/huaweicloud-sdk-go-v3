package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestRuleInstance struct {

	// 规则实例ID
	Id *string `json:"id,omitempty"`

	// 规则类型
	Type string `json:"type"`

	// 规则名称
	Name string `json:"name"`

	// 规则实例状态
	IsValid *bool `json:"is_valid,omitempty"`

	// 规则部署
	LayoutContent *string `json:"layout_content,omitempty"`

	// 插件Id
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件版本
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 规则属性列表
	Content []RuleInstanceContent `json:"content"`
}

func (o RequestRuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestRuleInstance struct{}"
	}

	return strings.Join([]string{"RequestRuleInstance", string(data)}, " ")
}
