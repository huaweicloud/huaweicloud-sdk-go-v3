package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentStatus **参数解释**： Agent状态 **取值范围**: - installed：已安装 - not_installed：未安 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中 - not_online：不在线的（除了在线以外的所有状态，仅作为查询条件）
type AgentStatus struct {
}

func (o AgentStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentStatus struct{}"
	}

	return strings.Join([]string{"AgentStatus", string(data)}, " ")
}
