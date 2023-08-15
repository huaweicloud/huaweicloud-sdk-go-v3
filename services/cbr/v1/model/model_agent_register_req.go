package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentRegisterReq
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
