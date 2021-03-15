package model

import (
	"encoding/json"

	"strings"
)

// 事件服务对象。
type ServiceEvent struct {
	// 设备事件类型。
	EventType string `json:"event_type"`
	// 设备事件的参数列表。
	Paras *[]ServiceCommandPara `json:"paras,omitempty"`
}

func (o ServiceEvent) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceEvent struct{}"
	}

	return strings.Join([]string{"ServiceEvent", string(data)}, " ")
}
