package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleMetricResponse Response Object
type ShowSingleMetricResponse struct {

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 指标值
	Metrics        *[]MetricDataItem `json:"metrics,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSingleMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleMetricResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleMetricResponse", string(data)}, " ")
}
