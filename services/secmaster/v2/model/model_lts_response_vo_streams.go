package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsResponseVoStreams struct {

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o LtsResponseVoStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsResponseVoStreams struct{}"
	}

	return strings.Join([]string{"LtsResponseVoStreams", string(data)}, " ")
}
