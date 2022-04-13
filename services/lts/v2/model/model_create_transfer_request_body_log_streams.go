package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTransferRequestBodyLogStreams struct {
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 日志流名称

	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o CreateTransferRequestBodyLogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequestBodyLogStreams struct{}"
	}

	return strings.Join([]string{"CreateTransferRequestBodyLogStreams", string(data)}, " ")
}
