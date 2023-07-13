package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StringColumnStatisticsData struct {

	// 列中字符串平均长度
	AverageLength float64 `json:"average_length"`

	// 列中字符串最长长度
	MaximumLength int64 `json:"maximum_length"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`

	// 列中去重后字符串个数
	NumberOfDistinctValue int64 `json:"number_of_distinct_value"`

	// 估算唯一值使用的位图
	BitVector string `json:"bit_vector,omitempty"`
}

func (o StringColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StringColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"StringColumnStatisticsData", string(data)}, " ")
}
