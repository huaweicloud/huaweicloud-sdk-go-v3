package model

import (
	"encoding/json"

	"strings"
)

type ErrorlogResult struct {
	// 节点名称。

	NodeName string `json:"node_name"`
	// 日志级别。

	Level string `json:"level"`
	// 发生时间，UTC时间。

	Time string `json:"time"`
	// 日志内容。

	Content string `json:"content"`
}

func (o ErrorlogResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ErrorlogResult struct{}"
	}

	return strings.Join([]string{"ErrorlogResult", string(data)}, " ")
}
