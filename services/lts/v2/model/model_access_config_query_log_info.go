package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志接入日志详情
type AccessConfigQueryLogInfo struct {

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty" xml:"log_group_id"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty" xml:"log_stream_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty" xml:"log_group_name"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty" xml:"log_stream_name"`
}

func (o AccessConfigQueryLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigQueryLogInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigQueryLogInfo", string(data)}, " ")
}
