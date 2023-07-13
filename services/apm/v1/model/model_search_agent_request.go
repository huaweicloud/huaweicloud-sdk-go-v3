package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAgentRequest Request Object
type SearchAgentRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *AgentSearchParam `json:"body,omitempty"`
}

func (o SearchAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAgentRequest struct{}"
	}

	return strings.Join([]string{"SearchAgentRequest", string(data)}, " ")
}
