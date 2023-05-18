package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgentRegisterReq struct {
	Agent *AgentRegister `json:"agent"`
}

func (o AgentRegisterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentRegisterReq struct{}"
	}

	return strings.Join([]string{"AgentRegisterReq", string(data)}, " ")
}
