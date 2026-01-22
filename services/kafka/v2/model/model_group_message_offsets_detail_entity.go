package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupMessageOffsetsDetailEntity 消费组消息位点详情
type GroupMessageOffsetsDetailEntity struct {

	// 分区
	Partition *string `json:"partition,omitempty"`

	// 消息当前位点
	MessageCurrentOffset *string `json:"message_current_offset,omitempty"`

	// topic名称
	Topic *string `json:"topic,omitempty"`

	// 消息开始位点
	MessageLogStartOffset *int32 `json:"message_log_start_offset,omitempty"`

	// 剩余可消费消息数，即消息堆积数
	Lag *int32 `json:"lag,omitempty"`

	// 消息结束位点
	MessageLogEndOffset *int32 `json:"message_log_end_offset,omitempty"`

	// 消费者Id
	ConsumerId *string `json:"consumer_id,omitempty"`

	// host名称
	Host *string `json:"host,omitempty"`

	// 客户端Id
	ClientId *string `json:"client_id,omitempty"`
}

func (o GroupMessageOffsetsDetailEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMessageOffsetsDetailEntity struct{}"
	}

	return strings.Join([]string{"GroupMessageOffsetsDetailEntity", string(data)}, " ")
}
