package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleReq struct {

	// 规则类型
	Type string `json:"type"`

	// 规则名称
	Name string `json:"name"`

	// 布局内容
	LayoutContent string `json:"layout_content"`

	// 插件ID
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件版本号
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 规则属性集
	Content []RuleContent `json:"content"`
}

func (o UpdateRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateRuleReq", string(data)}, " ")
}
