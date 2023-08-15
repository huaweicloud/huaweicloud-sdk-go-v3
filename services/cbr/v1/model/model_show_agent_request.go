package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentRequest Request Object
type ShowAgentRequest struct {

	// 客户端ID
	AgentId string `json:"agent_id"`
}

func (o ShowAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentRequest struct{}"
	}

	return strings.Join([]string{"ShowAgentRequest", string(data)}, " ")
}
