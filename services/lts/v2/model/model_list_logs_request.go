package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsRequest Request Object
type ListLogsRequest struct {

	// 日志组id。
	LogGroupId string `json:"log_group_id"`

	// 日志流id。
	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsLogParams `json:"body,omitempty"`
}

func (o ListLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLogsRequest", string(data)}, " ")
}
