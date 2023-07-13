package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// 时间序列
	Timestamps *[]string `json:"timestamps,omitempty"`

	// 查询的测量指标列表
	Metrics        *[]MetricList `json:"metrics,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
