package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAgentPathRequest Request Object
type RemoveAgentPathRequest struct {

	// 客户端ID
	AgentId string `json:"agent_id"`

	Body *AgentRemovePathReq `json:"body,omitempty"`
}

func (o RemoveAgentPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAgentPathRequest struct{}"
	}

	return strings.Join([]string{"RemoveAgentPathRequest", string(data)}, " ")
}
