package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListMetricDataResponse struct {
	// 监控指标。

	Metrics        *[]BatchMetricData `json:"metrics,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchListMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataResponse struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataResponse", string(data)}, " ")
}
