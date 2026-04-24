package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeAiAssistantResponse Response Object
type SubscribeAiAssistantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SubscribeAiAssistantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAiAssistantResponse struct{}"
	}

	return strings.Join([]string{"SubscribeAiAssistantResponse", string(data)}, " ")
}
