package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsResponseVo struct {

	// 配置id
	ConfigId *string `json:"config_id,omitempty"`

	// 配置名称
	ConfigName *string `json:"config_name,omitempty"`

	// 是否开启
	Enable *string `json:"enable,omitempty"`

	// 日志组Id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志类型
	LogType *string `json:"log_type,omitempty"`

	LogTypes *LtsResponseVoLogTypes `json:"log_types,omitempty"`

	// lts日志信息map
	LtsInfos *[]LtsResponseVoLtsInfos `json:"lts_infos,omitempty"`

	// 管道别名
	PipeAlias *string `json:"pipe_alias,omitempty"`

	// 类型前缀
	TypePrefix *string `json:"type_prefix,omitempty"`
}

func (o LtsResponseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsResponseVo struct{}"
	}

	return strings.Join([]string{"LtsResponseVo", string(data)}, " ")
}
