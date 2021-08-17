package model

import (
	"encoding/json"

	"strings"
)

// EIP异常事件响应体
type DailyLog struct {
	// 开始时间

	StartTime int64 `json:"start_time"`
	// 结束时间

	EndTime int64 `json:"end_time"`
	// 防护状态，可选范围： - 1：表示清洗 - 2：表示黑洞

	Status int32 `json:"status"`
	// 触发时流量

	TriggerBps int32 `json:"trigger_bps"`
	// 触发时报文速率

	TriggerPps int32 `json:"trigger_pps"`
	// 触发时HTTP请求速率

	TriggerHttpPps int32 `json:"trigger_http_pps"`
}

func (o DailyLog) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DailyLog struct{}"
	}

	return strings.Join([]string{"DailyLog", string(data)}, " ")
}
