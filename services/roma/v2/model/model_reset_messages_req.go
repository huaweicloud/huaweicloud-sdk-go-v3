package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetMessagesReq struct {
	// topic名称。

	Topic string `json:"topic"`
	// 分区。

	Partition int32 `json:"partition"`
	// 消息偏移量。

	MessageOffset float32 `json:"message_offset"`
	// 应用key。在该消息头中添加一个consumer_key的消息头。

	ConsumerKey *string `json:"consumer_key,omitempty"`
}

func (o ResetMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessagesReq struct{}"
	}

	return strings.Join([]string{"ResetMessagesReq", string(data)}, " ")
}
