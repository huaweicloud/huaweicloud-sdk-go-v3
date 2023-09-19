package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogMappingConfig struct {

	// 源日志组ID
	SourceLogGroupId *string `json:"source_log_group_id,omitempty"`

	// 目标日志组ID
	TargetLogGroupId *string `json:"target_log_group_id,omitempty"`

	// 目标日志组名称
	TargetLogGroupName *string `json:"target_log_group_name,omitempty"`

	// 日志流配置
	LogStreamConfig *[]LogMappingStreamInfo `json:"log_stream_config,omitempty"`
}

func (o LogMappingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogMappingConfig struct{}"
	}

	return strings.Join([]string{"LogMappingConfig", string(data)}, " ")
}
