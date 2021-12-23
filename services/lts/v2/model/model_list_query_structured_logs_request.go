package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQueryStructuredLogsRequest struct {
	// 日志组id。

	LogGroupId string `json:"log_group_id"`
	// 日志流id。

	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsStructLogParams `json:"body,omitempty"`
}

func (o ListQueryStructuredLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryStructuredLogsRequest struct{}"
	}

	return strings.Join([]string{"ListQueryStructuredLogsRequest", string(data)}, " ")
}
