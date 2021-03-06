package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateHealthMonitorResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateHealthMonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorResponse", string(data)}, " ")
}
