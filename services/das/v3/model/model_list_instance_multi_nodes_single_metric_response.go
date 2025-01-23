package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceMultiNodesSingleMetricResponse Response Object
type ListInstanceMultiNodesSingleMetricResponse struct {

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 指标值
	Metrics        *[]MultiNodesSingleMetricMetrics `json:"metrics,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListInstanceMultiNodesSingleMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMultiNodesSingleMetricResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceMultiNodesSingleMetricResponse", string(data)}, " ")
}
