package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentPluginInfoQueryDto struct {

	// 可选的查询条件-插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 可选的查询条件-匹配名称
	RegexName *string `json:"regex_name,omitempty"`

	// 维护者
	Maintainer *string `json:"maintainer,omitempty"`

	// 业务类型,[Build,Gate,Deploy,Test,Normal]
	BusinessType *[]string `json:"business_type,omitempty"`

	// 插件属性，official/custom
	PluginAttribution *string `json:"plugin_attribution,omitempty"`
}

func (o AgentPluginInfoQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentPluginInfoQueryDto struct{}"
	}

	return strings.Join([]string{"AgentPluginInfoQueryDto", string(data)}, " ")
}
