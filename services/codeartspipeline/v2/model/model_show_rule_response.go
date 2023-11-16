package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleResponse Response Object
type ShowRuleResponse struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则类型
	Type *string `json:"type,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则版本
	Version *string `json:"version,omitempty"`

	// 布局内容
	LayoutContent *string `json:"layout_content,omitempty"`

	// 插件ID
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件版本号
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新人
	Updater *string `json:"updater,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 规则属性列表
	Content        *[]RuleContent `json:"content,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleResponse", string(data)}, " ")
}
