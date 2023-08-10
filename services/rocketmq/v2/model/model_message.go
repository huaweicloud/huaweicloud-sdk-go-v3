package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Message 消息。
type Message struct {

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 主题名称。
	Topic *string `json:"topic,omitempty"`

	// 存储消息的时间。
	StoreTimestamp float32 `json:"store_timestamp,omitempty"`

	// 产生消息的时间。
	BornTimestamp float32 `json:"born_timestamp,omitempty"`

	// 重试次数。
	ReconsumeTimes *int32 `json:"reconsume_times,omitempty"`

	// 消息体。
	Body *string `json:"body,omitempty"`

	// 消息体校验和。
	BodyCrc float32 `json:"body_crc,omitempty"`

	// 存储大小。
	StoreSize float32 `json:"store_size,omitempty"`

	// 消息属性列表。
	PropertyList *[]MessagePropertyList `json:"property_list,omitempty"`

	// 产生消息的主机IP。
	BornHost *string `json:"born_host,omitempty"`

	// 存储消息的主机IP。
	StoreHost *string `json:"store_host,omitempty"`

	// 队列ID。
	QueueId *int32 `json:"queue_id,omitempty"`

	// 在队列中的偏移量。
	QueueOffset *int32 `json:"queue_offset,omitempty"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
