package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志接入基础日志信息。
type AccessConfigBaseLogInfo struct {
	// 日志组ID

	LogGroupId *string `json:"log_group_id,omitempty"`
	// 日志流ID

	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o AccessConfigBaseLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigBaseLogInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigBaseLogInfo", string(data)}, " ")
}
