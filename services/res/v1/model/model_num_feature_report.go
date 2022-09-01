package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户数字类型特征统计
type NumFeatureReport struct {

	// 平均值。
	Avg *float64 `json:"avg,omitempty" xml:"avg"`

	// 最小值。
	Min *float64 `json:"min,omitempty" xml:"min"`

	// 最大值。
	Max *float64 `json:"max,omitempty" xml:"max"`

	// 中位数。
	Median *float64 `json:"median,omitempty" xml:"median"`

	// 百分位统计。
	PercentsAgg *string `json:"percents_agg,omitempty" xml:"percents_agg"`

	// 特征名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 特征类型。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`
}

func (o NumFeatureReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NumFeatureReport struct{}"
	}

	return strings.Join([]string{"NumFeatureReport", string(data)}, " ")
}
