package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogStreamResponse Response Object
type CreateLogStreamResponse struct {

	// 创建的日志流的Id。
	LogStreamId    *string `json:"log_stream_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamResponse struct{}"
	}

	return strings.Join([]string{"CreateLogStreamResponse", string(data)}, " ")
}
