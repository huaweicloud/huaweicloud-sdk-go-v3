package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMetricDataRequest struct {
	// 添加一条或多条自定义指标监控数据，请求参数。

	Body *[]MetricDataItem `json:"body,omitempty"`
}

func (o CreateMetricDataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMetricDataRequest struct{}"
	}

	return strings.Join([]string{"CreateMetricDataRequest", string(data)}, " ")
}
