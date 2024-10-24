package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SingleAgentParam struct {

	// agent ID唯一值。
	AgentId string `json:"agent_id"`

	// 主机ip。
	InnerIp string `json:"inner_ip"`
}

func (o SingleAgentParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleAgentParam struct{}"
	}

	return strings.Join([]string{"SingleAgentParam", string(data)}, " ")
}
