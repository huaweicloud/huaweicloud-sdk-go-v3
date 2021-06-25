package model

import (
	"encoding/json"

	"strings"
)

// 创建日志流参数。
type CreateLogStreamParams struct {
	// 需要创建的日志流名称。

	LogStreamName string `json:"log_stream_name"`
	// 日志流Tag集合，不同日志流属性不同。

	Tag *interface{} `json:"tag,omitempty"`
}

func (o CreateLogStreamParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}
