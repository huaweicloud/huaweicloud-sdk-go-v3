package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartChatResponse Response Object
type StartChatResponse struct {

	// status
	Status *string `json:"status,omitempty"`

	// chat id
	ChatId *string `json:"chat_id,omitempty"`

	// error message
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartChatResponse struct{}"
	}

	return strings.Join([]string{"StartChatResponse", string(data)}, " ")
}
