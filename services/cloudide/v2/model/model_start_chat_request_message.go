package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartChatRequestMessage start chat request message
type StartChatRequestMessage struct {

	// user type
	UserType *string `json:"user_type,omitempty"`

	// user id
	UserId *string `json:"user_id,omitempty"`
}

func (o StartChatRequestMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartChatRequestMessage struct{}"
	}

	return strings.Join([]string{"StartChatRequestMessage", string(data)}, " ")
}
