package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentInfo struct {

	// 插件名称。
	AgentName *string `json:"agent_name,omitempty"`

	// 插件版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 是否安装插件（是否）。
	IsInstalled *bool `json:"is_installed,omitempty"`
}

func (o AgentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentInfo struct{}"
	}

	return strings.Join([]string{"AgentInfo", string(data)}, " ")
}
