package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgentRequest Request Object
type UpdateAgentRequest struct {
	AgentId string `json:"agent_id"`

	Body *AgentUpdateReq `json:"body,omitempty"`
}

func (o UpdateAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgentRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgentRequest", string(data)}, " ")
}
