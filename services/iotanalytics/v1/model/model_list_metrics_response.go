package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMetricsResponse struct {

	// 时间序列
	Timestamps *[]string `json:"timestamps,omitempty" xml:"timestamps"`

	// 查询的测量指标列表
	Metrics        *[]MetricList `json:"metrics,omitempty" xml:"metrics"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
