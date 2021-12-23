package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SendMessagesResponse struct {
	// 消息列表。

	Messages       *[]SendMessagesRespMessages `json:"messages,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o SendMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessagesResponse struct{}"
	}

	return strings.Join([]string{"SendMessagesResponse", string(data)}, " ")
}
