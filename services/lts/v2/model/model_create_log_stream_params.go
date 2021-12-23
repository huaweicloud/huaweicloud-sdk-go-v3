package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建日志流参数。
type CreateLogStreamParams struct {
	// 需要创建的日志流名称。

	LogStreamName string `json:"log_stream_name"`
}

func (o CreateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}
