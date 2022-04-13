package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchAgentRequest struct {
	// instance_id

	InstanceId string `json:"instance_id"`

	Body *AgentSwitchRequest `json:"body,omitempty"`
}

func (o SwitchAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAgentRequest struct{}"
	}

	return strings.Join([]string{"SwitchAgentRequest", string(data)}, " ")
}
