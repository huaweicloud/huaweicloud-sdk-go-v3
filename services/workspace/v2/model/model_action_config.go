package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionConfig 动作配置。
type ActionConfig struct {

	// 动作名称。STOP：关机，HIBERNATE：休眠，REBOOT：重启，EXECUTE_SCRIPT：执行脚本。
	Action *string `json:"action,omitempty"`

	// 最小等待时长，单位分钟。如果不填，则使用父级的 min_wait_time。
	MinWaitTime *int32 `json:"min_wait_time,omitempty"`

	// 默认等待时长，单位分钟。如果不填，则使用父级的 default_wait_time。
	DefaultWaitTime *int32 `json:"default_wait_time,omitempty"`

	// 最小执行周期，单位分钟。如果不填，则使用父级的 min_exec_time。
	MinExecTime *int32 `json:"min_exec_time,omitempty"`

	// 默认执行周期，单位分钟。如果不填，则使用父级的 default_exec_time。
	DefaultExecTime *int32 `json:"default_exec_time,omitempty"`
}

func (o ActionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionConfig struct{}"
	}

	return strings.Join([]string{"ActionConfig", string(data)}, " ")
}
