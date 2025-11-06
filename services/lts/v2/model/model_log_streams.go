package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogStreams struct {

	// 日志流ID。
	LogStreamId string `json:"log_stream_id"`

	// 日志流名称。
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o LogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStreams struct{}"
	}

	return strings.Join([]string{"LogStreams", string(data)}, " ")
}
