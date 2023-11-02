package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncGetChatResultResponse Response Object
type SyncGetChatResultResponse struct {

	// text
	Text *string `json:"text,omitempty"`

	// request id
	RequestId *string `json:"request_id,omitempty"`

	// status
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncGetChatResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncGetChatResultResponse struct{}"
	}

	return strings.Join([]string{"SyncGetChatResultResponse", string(data)}, " ")
}
