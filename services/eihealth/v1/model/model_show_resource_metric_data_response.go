package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceMetricDataResponse Response Object
type ShowResourceMetricDataResponse struct {

	// 监控数据列表
	DataPoints *[]DataPointDto `json:"data_points,omitempty"`

	// 监控指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 监控资源id
	ResourceId     *string `json:"resource_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceMetricDataResponse", string(data)}, " ")
}
