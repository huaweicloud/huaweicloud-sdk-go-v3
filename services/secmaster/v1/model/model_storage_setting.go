package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageSetting 存储设置
type StorageSetting struct {

	// 数据转换 CU
	DataTransformationCu float32 `json:"data_transformation_cu,omitempty"`

	// 索引副本数
	IndexReplicas *int32 `json:"index_replicas,omitempty"`

	// 索引分片数
	IndexShards *int32 `json:"index_shards,omitempty"`

	// 索引存储周期
	IndexStoragePeriod *int32 `json:"index_storage_period,omitempty"`

	// 索引存储大小
	IndexStorageSize *int32 `json:"index_storage_size,omitempty"`

	// 湖存储周期
	LakeStoragePeriod *int32 `json:"lake_storage_period,omitempty"`

	// 流式带宽
	StreamingBandwidth float32 `json:"streaming_bandwidth,omitempty"`

	// 流式数据空间ID
	StreamingDataspaceId *string `json:"streaming_dataspace_id,omitempty"`

	// 流式分区数
	StreamingPartition *int32 `json:"streaming_partition,omitempty"`

	// 流式保留大小
	StreamingRetentionSize *int32 `json:"streaming_retention_size,omitempty"`

	// 流式保留时间
	StreamingRetentionTime *int32 `json:"streaming_retention_time,omitempty"`
}

func (o StorageSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageSetting struct{}"
	}

	return strings.Join([]string{"StorageSetting", string(data)}, " ")
}
