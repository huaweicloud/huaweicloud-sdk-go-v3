package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsRquestVo struct {

	// 配置名称
	ConfigName *string `json:"config_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否开启
	Enable *string `json:"enable,omitempty"`

	// 日志ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志类型
	LogType *string `json:"log_type,omitempty"`

	// 日志前缀
	LogTypePrefix *string `json:"log_type_prefix,omitempty"`

	// 日志别名
	PipeAlias *string `json:"pipe_alias,omitempty"`
}

func (o LtsRquestVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsRquestVo struct{}"
	}

	return strings.Join([]string{"LtsRquestVo", string(data)}, " ")
}
