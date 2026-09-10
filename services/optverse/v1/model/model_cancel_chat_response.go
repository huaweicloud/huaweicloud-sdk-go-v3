package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelChatResponse Response Object
type CancelChatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelChatResponse struct{}"
	}

	return strings.Join([]string{"CancelChatResponse", string(data)}, " ")
}
