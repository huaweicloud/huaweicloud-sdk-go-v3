package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改的topic。
type UpdateInstanceTopicReqTopics struct {

	// topic名称，不支持修改。
	Id string `json:"id" xml:"id"`

	// 老化时间，单位小时。
	RetentionTime *int32 `json:"retention_time,omitempty" xml:"retention_time"`

	// 是否同步复制。
	SyncReplication *bool `json:"sync_replication,omitempty" xml:"sync_replication"`

	// 是否同步落盘。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty" xml:"sync_message_flush"`

	// 分区数。
	NewPartitionNumbers *int32 `json:"new_partition_numbers,omitempty" xml:"new_partition_numbers"`
}

func (o UpdateInstanceTopicReqTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopics struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopics", string(data)}, " ")
}
