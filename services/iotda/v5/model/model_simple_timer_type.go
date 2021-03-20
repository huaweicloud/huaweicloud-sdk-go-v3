package model

import (
	"encoding/json"

	"strings"
)

// 条件中简单定时类型的信息，自定义结构。
type SimpleTimerType struct {
	// 规则触发的开始时间，使用UTC时区，格式：yyyyMMdd'T'HHmmss'Z'。

	StartTime string `json:"start_time"`
	// 规则触发的重复时间间隔，单位为秒。

	RepeatInterval int32 `json:"repeat_interval"`
	// 规则触发的重复次数。

	RepeatCount int32 `json:"repeat_count"`
}

func (o SimpleTimerType) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SimpleTimerType struct{}"
	}

	return strings.Join([]string{"SimpleTimerType", string(data)}, " ")
}
