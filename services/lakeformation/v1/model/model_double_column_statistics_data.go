package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DoubleColumnStatisticsData 浮点数统计信息
type DoubleColumnStatisticsData struct {

	// 列中浮点数最小值
	MinimumValue float64 `json:"minimum_value"`

	// 列中浮点数最大值
	MaximumValue float64 `json:"maximum_value"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`

	// 列中去重后浮点数个数
	NumberOfDistinctValue int64 `json:"number_of_distinct_value"`

	// 估算唯一值使用的位图
	BitVector string `json:"bit_vector,omitempty"`
}

func (o DoubleColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DoubleColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"DoubleColumnStatisticsData", string(data)}, " ")
}
