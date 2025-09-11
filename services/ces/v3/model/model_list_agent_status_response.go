package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentStatusResponse Response Object
type ListAgentStatusResponse struct {

	// **参数解释**: agent插件状态列表 **取值范围**: 数组长度为[1,2000]
	AgentStatus    *[]AgentStatusInfo `json:"agent_status,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAgentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAgentStatusResponse", string(data)}, " ")
}
