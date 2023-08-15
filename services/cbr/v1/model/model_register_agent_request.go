package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAgentRequest Request Object
type RegisterAgentRequest struct {
	Body *AgentRegisterReq `json:"body,omitempty"`
}

func (o RegisterAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAgentRequest struct{}"
	}

	return strings.Join([]string{"RegisterAgentRequest", string(data)}, " ")
}
