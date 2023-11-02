package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncGetChatResultRequest Request Object
type SyncGetChatResultRequest struct {
	Body *ChatResultRequestMessage `json:"body,omitempty"`
}

func (o SyncGetChatResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncGetChatResultRequest struct{}"
	}

	return strings.Join([]string{"SyncGetChatResultRequest", string(data)}, " ")
}
