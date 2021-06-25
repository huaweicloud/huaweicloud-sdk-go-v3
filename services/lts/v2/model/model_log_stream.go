package model

import (
	"encoding/json"

	"strings"
)

// 返回的日志流信息
type LogStream struct {
	// 创建时间

	CreationTime int64 `json:"creation_time"`
	// 日志流名称

	LogStreamName string `json:"log_stream_name"`
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 过滤器个数

	FilterCount int32 `json:"filter_count"`
}

func (o LogStream) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LogStream struct{}"
	}

	return strings.Join([]string{"LogStream", string(data)}, " ")
}
