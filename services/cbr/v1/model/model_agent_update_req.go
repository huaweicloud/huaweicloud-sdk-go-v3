package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgentUpdateReq struct {
	Agent *AgentUpdate `json:"agent"`
}

func (o AgentUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentUpdateReq struct{}"
	}

	return strings.Join([]string{"AgentUpdateReq", string(data)}, " ")
}
