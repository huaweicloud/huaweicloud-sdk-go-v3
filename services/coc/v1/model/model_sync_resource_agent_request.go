package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourceAgentRequest Request Object
type SyncResourceAgentRequest struct {
	Body *SyncResourceAgentReq `json:"body,omitempty"`
}

func (o SyncResourceAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceAgentRequest struct{}"
	}

	return strings.Join([]string{"SyncResourceAgentRequest", string(data)}, " ")
}
