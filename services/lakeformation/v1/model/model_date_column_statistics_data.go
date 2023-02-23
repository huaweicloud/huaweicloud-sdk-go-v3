package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DateColumnStatisticsData struct {

	// 列中的最小时间戳
	MinimumValue *sdktime.SdkTime `json:"minimum_value,omitempty"`

	// 列中的最大时间戳
	MaximumValue *sdktime.SdkTime `json:"maximum_value,omitempty"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`

	// 列中去重后的时间戳个数
	NumberOfDistinctValue int64 `json:"number_of_distinct_value"`

	// 估算唯一值使用的位图
	BitVector string `json:"bit_vector,omitempty"`
}

func (o DateColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DateColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"DateColumnStatisticsData", string(data)}, " ")
}
