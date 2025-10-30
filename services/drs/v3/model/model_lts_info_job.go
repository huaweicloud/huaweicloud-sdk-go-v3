package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsInfoJob LTS信息。
type LtsInfoJob struct {

	// 日志组ID。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流ID。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 是否开启LTS。
	LtsEnabled *bool `json:"lts_enabled,omitempty"`
}

func (o LtsInfoJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsInfoJob struct{}"
	}

	return strings.Join([]string{"LtsInfoJob", string(data)}, " ")
}
