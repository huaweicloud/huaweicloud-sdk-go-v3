package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UnregisterAgentRequest struct {

	// 客户端ID
	AgentId string `json:"agent_id"`
}

func (o UnregisterAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnregisterAgentRequest struct{}"
	}

	return strings.Join([]string{"UnregisterAgentRequest", string(data)}, " ")
}
