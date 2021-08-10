package model

import (
	"encoding/json"

	"strings"
)

// 修改的topic。
type UpdateInstanceTopicReqTopics struct {
	// topic名称

	Id string `json:"id"`
	// 老化时间，单位小时。

	RetentionTime *int32 `json:"retention_time,omitempty"`
	// 是否同步复制。

	SyncReplication *bool `json:"sync_replication,omitempty"`
	// 是否同步落盘。

	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`
	// 分区数。

	NewPartitionNumbers *int32 `json:"new_partition_numbers,omitempty"`
}

func (o UpdateInstanceTopicReqTopics) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopics struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopics", string(data)}, " ")
}
