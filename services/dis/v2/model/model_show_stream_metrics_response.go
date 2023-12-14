package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStreamMetricsResponse Response Object
type ShowStreamMetricsResponse struct {
	Metrics *Metrics `json:"metrics,omitempty"`

	// 监控数据对象列表。
	MetricsList    *[]Metrics `json:"metrics_list,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowStreamMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowStreamMetricsResponse", string(data)}, " ")
}
