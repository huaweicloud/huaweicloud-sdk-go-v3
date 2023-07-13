package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAgentRequest Request Object
type CheckAgentRequest struct {
	Body *ProtectableAgentReq `json:"body,omitempty"`
}

func (o CheckAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAgentRequest struct{}"
	}

	return strings.Join([]string{"CheckAgentRequest", string(data)}, " ")
}
