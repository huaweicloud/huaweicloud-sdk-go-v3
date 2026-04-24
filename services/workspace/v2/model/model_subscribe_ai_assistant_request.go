package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeAiAssistantRequest Request Object
type SubscribeAiAssistantRequest struct {
	Body *SubscribeAiAssistantReq `json:"body,omitempty"`
}

func (o SubscribeAiAssistantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAiAssistantRequest struct{}"
	}

	return strings.Join([]string{"SubscribeAiAssistantRequest", string(data)}, " ")
}
