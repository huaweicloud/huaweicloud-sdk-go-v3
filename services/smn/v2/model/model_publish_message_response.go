package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PublishMessageResponse struct {
	// 请求的唯一标识ID。

	RequestId *string `json:"request_id,omitempty"`
	// 唯一的消息ID。

	MessageId      *string `json:"message_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishMessageResponse struct{}"
	}

	return strings.Join([]string{"PublishMessageResponse", string(data)}, " ")
}
