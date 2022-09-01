package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标描述信息
type MetricDemision struct {

	// 命名空间
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 指标名称
	MetricName *string `json:"metricName,omitempty" xml:"metricName"`

	// 维度列表
	Dimensions *[]Dimension `json:"dimensions,omitempty" xml:"dimensions"`
}

func (o MetricDemision) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDemision struct{}"
	}

	return strings.Join([]string{"MetricDemision", string(data)}, " ")
}
