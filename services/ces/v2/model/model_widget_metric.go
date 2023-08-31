package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WidgetMetric struct {

	// 服务维度
	Namespace *string `json:"namespace,omitempty"`

	Dimensions *DimensionInfo `json:"dimensions,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 监控视图的指标别名列表
	Alias *[]string `json:"alias,omitempty"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`
}

func (o WidgetMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetMetric struct{}"
	}

	return strings.Join([]string{"WidgetMetric", string(data)}, " ")
}
