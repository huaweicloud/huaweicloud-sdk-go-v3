package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MissingIndexTrendPoint 缺失索引趋势点
type MissingIndexTrendPoint struct {

	// 采集时间
	CollectTime *int64 `json:"collect_time,omitempty"`

	// 索引缺失总数
	TotalMissingIndexCount *int64 `json:"total_missing_index_count,omitempty"`
}

func (o MissingIndexTrendPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MissingIndexTrendPoint struct{}"
	}

	return strings.Join([]string{"MissingIndexTrendPoint", string(data)}, " ")
}
