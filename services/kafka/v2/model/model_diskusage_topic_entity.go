package model

import (
	"encoding/json"

	"strings"
)

type DiskusageTopicEntity struct {
	// 磁盘使用量。

	Size *string `json:"size,omitempty"`
	// topic名称。

	TopicName *string `json:"topic_name,omitempty"`
	// 分区。

	TopicPartition *string `json:"topic_partition,omitempty"`
	// 磁盘使用量的占比。

	Percentage *float64 `json:"percentage,omitempty"`
}

func (o DiskusageTopicEntity) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DiskusageTopicEntity struct{}"
	}

	return strings.Join([]string{"DiskusageTopicEntity", string(data)}, " ")
}
