package model

import (
	"encoding/json"

	"strings"
)

type ProbeDto struct {
	// 执行探测的命令行命令

	ExecCommand *string `json:"exec_command,omitempty"`

	HttpGet *HttpGetDto `json:"http_get,omitempty"`
	// 表示从工作负载启动后从多久开始探测

	InitialDelaySeconds int32 `json:"initial_delay_seconds"`
	// 表示探测超时时间

	TimeoutSeconds int32 `json:"timeout_seconds"`
}

func (o ProbeDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProbeDto struct{}"
	}

	return strings.Join([]string{"ProbeDto", string(data)}, " ")
}
