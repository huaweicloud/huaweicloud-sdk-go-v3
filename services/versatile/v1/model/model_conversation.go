package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Conversation 会话记录
type Conversation struct {
	LastUpdateTime *int64 `json:"lastUpdateTime,omitempty"`

	MessageList *[]Message `json:"messageList,omitempty"`
}

func (o Conversation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Conversation struct{}"
	}

	return strings.Join([]string{"Conversation", string(data)}, " ")
}
