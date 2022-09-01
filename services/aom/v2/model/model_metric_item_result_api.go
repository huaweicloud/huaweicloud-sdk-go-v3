package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维度信息。
type MetricItemResultApi struct {

	// 指标维度列表。
	Dimensions *[]Dimension `json:"dimensions,omitempty" xml:"dimensions"`

	// 指标哈希值。
	Dimensionvaluehash *string `json:"dimensionvaluehash,omitempty" xml:"dimensionvaluehash"`

	// 指标名称。
	MetricName *string `json:"metricName,omitempty" xml:"metricName"`

	// 命名空间。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 指标单位。
	Unit *string `json:"unit,omitempty" xml:"unit"`
}

func (o MetricItemResultApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricItemResultApi struct{}"
	}

	return strings.Join([]string{"MetricItemResultApi", string(data)}, " ")
}
