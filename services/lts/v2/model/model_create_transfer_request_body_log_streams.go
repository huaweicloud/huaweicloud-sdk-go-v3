package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTransferRequestBodyLogStreams struct {
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
}

func (o CreateTransferRequestBodyLogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequestBodyLogStreams struct{}"
	}

	return strings.Join([]string{"CreateTransferRequestBodyLogStreams", string(data)}, " ")
}
