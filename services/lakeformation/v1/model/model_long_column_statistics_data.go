package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LongColumnStatisticsData struct {

	// 列中长整数最小值
	MinimumValue int64 `json:"minimum_value"`

	// 列中长整数最大值
	MaximumValue int64 `json:"maximum_value"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`

	// 列中去重后的长整数个数
	NumberOfDistinctValue int64 `json:"number_of_distinct_value"`

	// 估算唯一值使用的位图
	BitVector string `json:"bit_vector,omitempty"`
}

func (o LongColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LongColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"LongColumnStatisticsData", string(data)}, " ")
}
