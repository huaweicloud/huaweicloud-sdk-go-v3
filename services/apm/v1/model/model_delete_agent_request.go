package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAgentRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *AgentDeleteParam `json:"body,omitempty"`
}

func (o DeleteAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgentRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgentRequest", string(data)}, " ")
}
