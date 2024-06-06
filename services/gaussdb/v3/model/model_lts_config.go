package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsConfig struct {

	// 日志类型。
	LogType *string `json:"log_type,omitempty"`

	// 日志组ID。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// 日志流ID。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`

	// 是否开启
	Enabled *bool `json:"enabled,omitempty"`
}

func (o LtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfig struct{}"
	}

	return strings.Join([]string{"LtsConfig", string(data)}, " ")
}
