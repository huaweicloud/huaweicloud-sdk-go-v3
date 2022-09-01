package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维度信息。
type SeriesQueryItemResult struct {

	// 命名空间。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 维度列表。
	Dimensions *[]DimensionSeries `json:"dimensions,omitempty" xml:"dimensions"`

	// 时间序列名称。
	MetricName *string `json:"metric_name,omitempty" xml:"metric_name"`

	// 时间序列单位。
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 时间序列哈希值。
	DimensionValueHash *string `json:"dimension_value_hash,omitempty" xml:"dimension_value_hash"`
}

func (o SeriesQueryItemResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SeriesQueryItemResult struct{}"
	}

	return strings.Join([]string{"SeriesQueryItemResult", string(data)}, " ")
}
