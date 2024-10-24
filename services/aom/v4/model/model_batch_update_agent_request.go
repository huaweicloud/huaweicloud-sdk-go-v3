package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAgentRequest Request Object
type BatchUpdateAgentRequest struct {
	Body *AgentUpgradeParam `json:"body,omitempty"`
}

func (o BatchUpdateAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAgentRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAgentRequest", string(data)}, " ")
}
