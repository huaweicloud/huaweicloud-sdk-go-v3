package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEdgeSiteMetricsResponse struct {

	// 监控数据
	MetricData     *[]MetricDataDetail `json:"metric_data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEdgeSiteMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSiteMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeSiteMetricsResponse", string(data)}, " ")
}
