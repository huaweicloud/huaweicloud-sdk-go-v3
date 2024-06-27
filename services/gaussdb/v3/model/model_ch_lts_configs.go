package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChLtsConfigs struct {

	// 日志类型。
	LogType *string `json:"log_type,omitempty"`

	// LTS日志组id。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// LTS日志流id。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`

	// LTS配置开关状态。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ChLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChLtsConfigs struct{}"
	}

	return strings.Join([]string{"ChLtsConfigs", string(data)}, " ")
}
