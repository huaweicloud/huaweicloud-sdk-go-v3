package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncChatRequest Request Object
type SyncChatRequest struct {
	Body *ChatRequestMessage `json:"body,omitempty"`
}

func (o SyncChatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncChatRequest struct{}"
	}

	return strings.Join([]string{"SyncChatRequest", string(data)}, " ")
}
