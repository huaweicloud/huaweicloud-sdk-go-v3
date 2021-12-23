package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLogStreamsRequest struct {
	// 日志组名称

	LogGroupName *string `json:"log_group_name,omitempty"`
	// 日志流名称

	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o ListLogStreamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListLogStreamsRequest", string(data)}, " ")
}
