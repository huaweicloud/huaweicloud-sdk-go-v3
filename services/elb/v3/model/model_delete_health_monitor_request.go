package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteHealthMonitorRequest struct {
	// 健康检查ID。

	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthMonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorRequest", string(data)}, " ")
}
