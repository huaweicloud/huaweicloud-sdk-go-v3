package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentPluginInfoQueryDto struct {
	PluginName *string `json:"plugin_name,omitempty"`

	RegexName *string `json:"regex_name,omitempty"`

	Maintainer *string `json:"maintainer,omitempty"`

	BusinessType *[]string `json:"business_type,omitempty"`

	PluginAttribution *string `json:"plugin_attribution,omitempty"`
}

func (o AgentPluginInfoQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentPluginInfoQueryDto struct{}"
	}

	return strings.Join([]string{"AgentPluginInfoQueryDto", string(data)}, " ")
}
