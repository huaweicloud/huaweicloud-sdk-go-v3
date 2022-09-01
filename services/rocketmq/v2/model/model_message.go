package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息。
type Message struct {

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty" xml:"msg_id"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 主题名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 存储消息的时间。
	StoreTimestamp float32 `json:"store_timestamp,omitempty" xml:"store_timestamp"`

	// 产生消息的时间。
	BornTimestamp float32 `json:"born_timestamp,omitempty" xml:"born_timestamp"`

	// 重试次数。
	ReconsumeTimes *string `json:"reconsume_times,omitempty" xml:"reconsume_times"`

	// 消息体。
	Body *string `json:"body,omitempty" xml:"body"`

	// 消息体校验和。
	BodyCrc float32 `json:"body_crc,omitempty" xml:"body_crc"`

	// 存储大小。
	StoreSize float32 `json:"store_size,omitempty" xml:"store_size"`

	// 消息属性列表。
	PropertyList *[]MessagePropertyList `json:"property_list,omitempty" xml:"property_list"`

	// 产生消息的主机IP。
	BornHost *string `json:"born_host,omitempty" xml:"born_host"`

	// 存储消息的主机IP。
	StoreHost *string `json:"store_host,omitempty" xml:"store_host"`

	// 队列ID。
	QueueId *string `json:"queue_id,omitempty" xml:"queue_id"`

	// 在队列中的偏移量。
	QueueOffset *string `json:"queue_offset,omitempty" xml:"queue_offset"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
