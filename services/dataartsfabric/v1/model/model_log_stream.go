package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogStream 日志流信息
type LogStream struct {

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o LogStream) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStream struct{}"
	}

	return strings.Join([]string{"LogStream", string(data)}, " ")
}
