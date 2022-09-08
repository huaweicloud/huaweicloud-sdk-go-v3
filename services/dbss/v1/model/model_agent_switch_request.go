package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentSwitchRequest struct {

	// 审计agent的ID
	AgentId string `json:"agent_id"`

	// Agent开关状态 1：开启 0：关闭
	Status int32 `json:"status"`
}

func (o AgentSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentSwitchRequest struct{}"
	}

	return strings.Join([]string{"AgentSwitchRequest", string(data)}, " ")
}
