package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricRequest3 struct {

	// 统计周期
	DateRange *string `json:"date_range,omitempty"`

	// 指标类型
	MetricType *string `json:"metric_type,omitempty"`

	// 迭代ID
	SprintId *string `json:"sprint_id,omitempty"`

	Dividend *MetricRequest3Dividend `json:"dividend,omitempty"`

	// 指标分母过滤条件
	Divisor *interface{} `json:"divisor,omitempty"`
}

func (o MetricRequest3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricRequest3 struct{}"
	}

	return strings.Join([]string{"MetricRequest3", string(data)}, " ")
}
