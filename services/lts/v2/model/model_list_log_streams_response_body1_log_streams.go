package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLogStreamsResponseBody1LogStreams struct {
	// 日志流创建时间

	CreationTime int64 `json:"creation_time"`
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 日志流名称

	LogStreamName string `json:"log_stream_name"`
	// 日志流所属标签

	Tag map[string]string `json:"tag"`
	// 过滤器个数

	FilterCount int32 `json:"filter_count"`
}

func (o ListLogStreamsResponseBody1LogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamsResponseBody1LogStreams struct{}"
	}

	return strings.Join([]string{"ListLogStreamsResponseBody1LogStreams", string(data)}, " ")
}
