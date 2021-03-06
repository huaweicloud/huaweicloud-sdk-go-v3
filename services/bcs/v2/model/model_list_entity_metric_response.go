package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEntityMetricResponse struct {
	// 指标对象列表。

	Metrics        *[]EntityMetricList `json:"metrics,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEntityMetricResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEntityMetricResponse struct{}"
	}

	return strings.Join([]string{"ListEntityMetricResponse", string(data)}, " ")
}
