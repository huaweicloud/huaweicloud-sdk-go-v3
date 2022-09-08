package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAgentStatusResponse struct {

	// Agent状态
	Status *string `json:"status,omitempty"`

	// AgentID
	AgentId        *string `json:"agent_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentStatusResponse", string(data)}, " ")
}
