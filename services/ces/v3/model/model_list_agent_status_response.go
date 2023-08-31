package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentStatusResponse Response Object
type ListAgentStatusResponse struct {

	// agent插件状态列表
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
