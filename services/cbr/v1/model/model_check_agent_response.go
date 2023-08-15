package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAgentResponse Response Object
type CheckAgentResponse struct {

	// 状态列表
	AgentStatus    *[]ProtectableAgentStatus `json:"agent_status,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CheckAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAgentResponse struct{}"
	}

	return strings.Join([]string{"CheckAgentResponse", string(data)}, " ")
}
