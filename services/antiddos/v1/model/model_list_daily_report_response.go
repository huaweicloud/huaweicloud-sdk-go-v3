package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDailyReportResponse struct {
	// 24小时内的流量数据

	Data           *[]DailyData `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDailyReportResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDailyReportResponse struct{}"
	}

	return strings.Join([]string{"ListDailyReportResponse", string(data)}, " ")
}
