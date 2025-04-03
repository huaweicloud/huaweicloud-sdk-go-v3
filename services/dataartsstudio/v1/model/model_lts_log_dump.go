package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsLogDump struct {

	// 是否开启LTS日志转储功能，true表示开启，false表示关闭。
	LogDump *bool `json:"log_dump,omitempty"`

	// 日志组ID。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志组名称。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流ID。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称。
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o LtsLogDump) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsLogDump struct{}"
	}

	return strings.Join([]string{"LtsLogDump", string(data)}, " ")
}
