package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListInstanceMetricResponse struct {
	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListInstanceMetricResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceMetricResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricResponse", string(data)}, " ")
}
