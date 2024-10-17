package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAgentRequest Request Object
type SwitchAgentRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
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
