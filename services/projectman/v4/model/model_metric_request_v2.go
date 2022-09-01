package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricRequestV2 struct {

	// 统计周期
	DateRange *string `json:"date_range,omitempty" xml:"date_range"`

	// 指标类型
	MetricType *string `json:"metric_type,omitempty" xml:"metric_type"`

	Dividend *MetricRequestV2Dividend `json:"dividend,omitempty" xml:"dividend"`

	// 指标分母过滤条件
	Divisor *interface{} `json:"divisor,omitempty" xml:"divisor"`
}

func (o MetricRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricRequestV2 struct{}"
	}

	return strings.Join([]string{"MetricRequestV2", string(data)}, " ")
}
