package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAlarmRequest struct {
	// 告警规则的ID。

	AlarmId string `json:"alarm_id"`
}

func (o DeleteAlarmRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAlarmRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRequest", string(data)}, " ")
}
