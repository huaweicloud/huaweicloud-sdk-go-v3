package model

import (
	"encoding/json"

	"strings"
)

type Event struct {
	// 事件类型

	EventType *string `json:"eventType,omitempty"`
	// DIS通道名称

	Channel *string `json:"channel,omitempty"`
	// 执行失败处理策略

	FailPolicy *string `json:"failPolicy,omitempty"`
	// 调度并发数

	Concurrent *int32 `json:"concurrent,omitempty"`
	// 读取策略

	ReadPolicy *string `json:"readPolicy,omitempty"`
}

func (o Event) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
