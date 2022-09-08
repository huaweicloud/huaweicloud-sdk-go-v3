package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetMessagesResponse struct {

	// topic名称。
	Topic *string `json:"topic,omitempty"`

	// 分区。
	Partition *int32 `json:"partition,omitempty"`

	// 消息偏移量。
	MessageOffset *int64 `json:"message_offset,omitempty"`

	// 应用key。在该消息头中添加一个consumer_key的消息头。
	ConsumerKey    *string `json:"consumer_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessagesResponse struct{}"
	}

	return strings.Join([]string{"ResetMessagesResponse", string(data)}, " ")
}
