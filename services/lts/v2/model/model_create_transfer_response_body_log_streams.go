package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTransferResponseBodyLogStreams struct {
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 日志流名称

	LogStreamName string `json:"log_stream_name"`
}

func (o CreateTransferResponseBodyLogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferResponseBodyLogStreams struct{}"
	}

	return strings.Join([]string{"CreateTransferResponseBodyLogStreams", string(data)}, " ")
}
