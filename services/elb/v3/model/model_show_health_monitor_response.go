package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHealthMonitorResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHealthMonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthMonitorResponse", string(data)}, " ")
}
