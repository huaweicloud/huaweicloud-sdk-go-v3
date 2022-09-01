package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMqsInstanceMessagesRespMessages struct {

	// topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 消息所在的分区。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// 消息key。
	Key *string `json:"key,omitempty" xml:"key"`

	// 消息内容。
	Value *string `json:"value,omitempty" xml:"value"`

	// 消息大小。
	Size *int64 `json:"size,omitempty" xml:"size"`

	// topic名称。
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 大数据标识。
	HugeMessage *bool `json:"huge_message,omitempty" xml:"huge_message"`

	// 消息偏移量。
	MessageOffset *int64 `json:"message_offset,omitempty" xml:"message_offset"`

	// 消息ID。
	MessageId *string `json:"message_id,omitempty" xml:"message_id"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 消息标签。
	Tag *string `json:"tag,omitempty" xml:"tag"`
}

func (o ShowMqsInstanceMessagesRespMessages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceMessagesRespMessages struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceMessagesRespMessages", string(data)}, " ")
}
