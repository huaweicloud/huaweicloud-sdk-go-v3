package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志接入日志详情
type AccessConfigQueryLogInfo struct {
	// 日志组ID

	LogGroupId *string `json:"log_group_id,omitempty"`
	// 日志流ID

	LogStreamId *string `json:"log_stream_id,omitempty"`
	// 日志组名称

	LogGroupName *string `json:"log_group_name,omitempty"`
	// 日志流名称

	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o AccessConfigQueryLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigQueryLogInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigQueryLogInfo", string(data)}, " ")
}
