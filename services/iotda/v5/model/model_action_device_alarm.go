package model

import (
	"encoding/json"

	"strings"
)

// 上报设备告警消息结构
type ActionDeviceAlarm struct {
	// 告警名称。

	Name string `json:"name"`
	// 告警状态。 - fault：上报告警。 - recovery：恢复告警。

	AlarmStatus string `json:"alarm_status"`
	// 告警级别,取值范围：warning（警告）、minor（一般）、major（严重）和critical（致命）。

	Severity string `json:"severity"`
	// 告警的描述信息。

	Description *string `json:"description,omitempty"`
}

func (o ActionDeviceAlarm) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionDeviceAlarm struct{}"
	}

	return strings.Join([]string{"ActionDeviceAlarm", string(data)}, " ")
}
