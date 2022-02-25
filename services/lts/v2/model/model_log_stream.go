package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	// 日志流所属标签

	Tag map[string]string `json:"tag,omitempty"`
}

func (o LogStream) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStream struct{}"
	}

	return strings.Join([]string{"LogStream", string(data)}, " ")
}
