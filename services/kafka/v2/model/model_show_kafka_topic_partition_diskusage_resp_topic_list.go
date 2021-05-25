package model

import (
	"encoding/json"

	"strings"
)

type ShowKafkaTopicPartitionDiskusageRespTopicList struct {
	// 磁盘使用量。

	Size *string `json:"size,omitempty"`
	// topic名称。

	TopicName *string `json:"topic_name,omitempty"`
	// 分区。

	TopicPartition *string `json:"topic_partition,omitempty"`
	// 磁盘使用量的占比。

	Percentage *float64 `json:"percentage,omitempty"`
}

func (o ShowKafkaTopicPartitionDiskusageRespTopicList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicPartitionDiskusageRespTopicList struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicPartitionDiskusageRespTopicList", string(data)}, " ")
}
