package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAlarmResponse struct {
	// 告警对象列表。

	MetricAlarms   *[]MetricAlarms `json:"metric_alarms,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAlarmResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmResponse", string(data)}, " ")
}
