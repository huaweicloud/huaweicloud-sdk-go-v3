package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionsLogConfig struct {

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o ConnectionsLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionsLogConfig struct{}"
	}

	return strings.Join([]string{"ConnectionsLogConfig", string(data)}, " ")
}
