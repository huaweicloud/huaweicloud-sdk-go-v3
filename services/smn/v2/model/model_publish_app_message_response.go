package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppMessageResponse Response Object
type PublishAppMessageResponse struct {

	// 唯一的消息ID。
	MessageId *string `json:"message_id,omitempty"`

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishAppMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppMessageResponse struct{}"
	}

	return strings.Join([]string{"PublishAppMessageResponse", string(data)}, " ")
}
