package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Queue struct {

	// 队列ID。
	Id *int32 `json:"id,omitempty"`

	// 队列消费堆积总数。
	Lag *int64 `json:"lag,omitempty"`

	// 队列消息总数。
	BrokerOffset *int64 `json:"broker_offset,omitempty"`

	// 已消费消息数。
	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	// 最新消费消息的存储时间，unix毫秒时间戳格式。
	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o Queue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Queue struct{}"
	}

	return strings.Join([]string{"Queue", string(data)}, " ")
}
