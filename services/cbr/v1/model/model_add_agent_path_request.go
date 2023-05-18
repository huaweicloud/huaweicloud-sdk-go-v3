package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddAgentPathRequest struct {

	// 客户端ID
	AgentId string `json:"agent_id"`

	Body *AgentAddPathReq `json:"body,omitempty"`
}

func (o AddAgentPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAgentPathRequest struct{}"
	}

	return strings.Join([]string{"AddAgentPathRequest", string(data)}, " ")
}
