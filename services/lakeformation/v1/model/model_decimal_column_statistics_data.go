package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecimalColumnStatisticsData 小数统计信息
type DecimalColumnStatisticsData struct {
	MinimumValue *Decimal `json:"minimum_value"`

	MaximumValue *Decimal `json:"maximum_value"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`

	// 列中去重后的小数个数
	NumberOfDistinctValue int64 `json:"number_of_distinct_value"`

	// 估算唯一值使用的位图
	BitVector string `json:"bit_vector,omitempty"`
}

func (o DecimalColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecimalColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"DecimalColumnStatisticsData", string(data)}, " ")
}
