package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipeStorageSetting 管道存储设置
type PipeStorageSetting struct {

	// Bandwidth in MB/s
	StreamingBandwidth *float32 `json:"streaming_bandwidth,omitempty"`

	// UUID
	StreamingDataspaceId string `json:"streaming_dataspace_id"`

	// 索引存储周期
	IndexStoragePeriod *int64 `json:"index_storage_period,omitempty"`

	// 索引存储容量，单位：GB
	IndexStorageSize *int64 `json:"index_storage_size,omitempty"`

	// 索引分区数
	IndexShards *int64 `json:"index_shards,omitempty"`

	// 数据转换CU
	DataTransformationCu *float32 `json:"data_transformation_cu,omitempty"`

	// Index shards
	LakeStoragePeriod *int64 `json:"lake_storage_period,omitempty"`
}

func (o PipeStorageSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeStorageSetting struct{}"
	}

	return strings.Join([]string{"PipeStorageSetting", string(data)}, " ")
}
