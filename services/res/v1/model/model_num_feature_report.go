package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户数字类型特征统计
type NumFeatureReport struct {

	// 平均值。
	Avg *float64 `json:"avg,omitempty"`

	// 最小值。
	Min *float64 `json:"min,omitempty"`

	// 最大值。
	Max *float64 `json:"max,omitempty"`

	// 中位数。
	Median *float64 `json:"median,omitempty"`

	// 百分位统计。
	PercentsAgg *string `json:"percents_agg,omitempty"`

	// 特征名。
	Name *string `json:"name,omitempty"`

	// 特征类型。
	DataType *string `json:"data_type,omitempty"`
}

func (o NumFeatureReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NumFeatureReport struct{}"
	}

	return strings.Join([]string{"NumFeatureReport", string(data)}, " ")
}
