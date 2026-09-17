package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndexUsageStatisticsResponse Response Object
type ShowIndexUsageStatisticsResponse struct {

	// 采集时间(ms)
	CollectTime *int64 `json:"collect_time,omitempty"`

	// 索引使用总数
	TotalIndexUsageCount *int64 `json:"total_index_usage_count,omitempty"`

	// 索引总空间(MB)
	IndexSizeMb *float64 `json:"index_size_mb,omitempty"`

	// 碎片率大于30%的数量
	FragmentationGl30Count *int64 `json:"fragmentation_gl30_count,omitempty"`

	// 查找次数小于100的数量
	KeyLookupLt100Count *int64 `json:"key_lookup_lt100_count,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowIndexUsageStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndexUsageStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowIndexUsageStatisticsResponse", string(data)}, " ")
}
