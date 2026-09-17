package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndexUsageTrendPoint 索引使用趋势点
type IndexUsageTrendPoint struct {

	// 采集时间
	CollectTime *int64 `json:"collect_time,omitempty"`

	// TOP1碎片率
	MaxFragmentationPercentage *float64 `json:"max_fragmentation_percentage,omitempty"`

	// 总空间大小(MB)
	IndexSizeMb *float64 `json:"index_size_mb,omitempty"`

	// 页数量
	PageCount *float64 `json:"page_count,omitempty"`
}

func (o IndexUsageTrendPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndexUsageTrendPoint struct{}"
	}

	return strings.Join([]string{"IndexUsageTrendPoint", string(data)}, " ")
}
