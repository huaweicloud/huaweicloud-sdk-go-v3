package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeAgentStatusRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *AgentStatusChangeParam `json:"body,omitempty"`
}

func (o ChangeAgentStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAgentStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeAgentStatusRequest", string(data)}, " ")
}
