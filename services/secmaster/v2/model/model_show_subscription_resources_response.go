package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionResourcesResponse Response Object
type ShowSubscriptionResourcesResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	Sku *SkuEnum `json:"sku,omitempty"`

	// 资源上限值
	UpperLimit *float64 `json:"upper_limit,omitempty"`

	// 配额单位（如 GB、条、分片等）
	Unit *string `json:"unit,omitempty"`

	// 配额步长
	Step *float64 `json:"step,omitempty"`

	// 已使用的资源数量
	UsedAmount *float64 `json:"used_amount,omitempty"`

	// 剩余可用的资源数量
	UnusedAmount *float64 `json:"unused_amount,omitempty"`

	// 版本号
	Version *int64 `json:"version,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 索引存储上限
	IndexStorageUpperLimit *int64 `json:"index_storage_upper_limit,omitempty"`

	// 索引分片上限
	IndexShardsUpperLimit *int64 `json:"index_shards_upper_limit,omitempty"`

	// 剩余可用索引分片数量
	IndexShardsUnused *int64 `json:"index_shards_unused,omitempty"`

	// 剩余可用分区数量
	PartitionsUnused *int64 `json:"partitions_unused,omitempty"`

	// 分区上限
	PartitionUpperLimit *int64 `json:"partition_upper_limit,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowSubscriptionResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionResourcesResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionResourcesResponse", string(data)}, " ")
}
