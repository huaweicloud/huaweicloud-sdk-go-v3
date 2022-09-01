package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息体。
type ShowPartitionMessageEntity struct {

	// 消息的key。
	Key *string `json:"key,omitempty" xml:"key"`

	// 消息内容。
	Value *string `json:"value,omitempty" xml:"value"`

	// Topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 分区编号。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// 消息位置。
	MessageOffset *int64 `json:"message_offset,omitempty" xml:"message_offset"`

	// 消息大小，单位字节。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 生产消息的时间。 格式为Unix时间戳。单位为毫秒。
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`
}

func (o ShowPartitionMessageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMessageEntity struct{}"
	}

	return strings.Join([]string{"ShowPartitionMessageEntity", string(data)}, " ")
}
