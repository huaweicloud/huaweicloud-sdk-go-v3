package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerDetails struct {

	// 消费者标识
	ConsumerTag *string `json:"consumer_tag,omitempty"`

	ChannelDetails *ChannelDetails `json:"channel_details,omitempty"`

	// 消费者客户端是否设置手动ack
	AckRequired *bool `json:"ack_required,omitempty"`

	// 消费者客户端预取值
	PrefetchCount *int32 `json:"prefetch_count,omitempty"`
}

func (o ConsumerDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerDetails struct{}"
	}

	return strings.Join([]string{"ConsumerDetails", string(data)}, " ")
}
