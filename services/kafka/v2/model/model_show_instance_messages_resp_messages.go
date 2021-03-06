package model

import (
	"encoding/json"

	"strings"
)

type ShowInstanceMessagesRespMessages struct {
	// topic名称。

	Topic *string `json:"topic,omitempty"`
	// 消息所在的分区。

	Partition *int32 `json:"partition,omitempty"`
	// 消息key。

	Key *string `json:"key,omitempty"`
	// 消息内容。

	Value *string `json:"value,omitempty"`
	// 消息大小。

	Size *int32 `json:"size,omitempty"`
	// topic名称。

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 大数据标识。

	HugeMessage *bool `json:"huge_message,omitempty"`
	// 消息偏移量。

	MessageOffset *int32 `json:"message_offset,omitempty"`
	// 消息ID。

	MessageId *string `json:"message_id,omitempty"`
	// 应用ID。

	AppId *string `json:"app_id,omitempty"`
	// 消息标签。

	Tag *string `json:"tag,omitempty"`
}

func (o ShowInstanceMessagesRespMessages) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceMessagesRespMessages struct{}"
	}

	return strings.Join([]string{"ShowInstanceMessagesRespMessages", string(data)}, " ")
}
