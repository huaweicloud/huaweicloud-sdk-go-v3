package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作任务动作参数。
type ActionParams struct {

	// 测试连接数据库信息。
	Endpoints *[]JobEndpointInfo `json:"endpoints,omitempty"`

	// 预检查模式。
	PrecheckMode *string `json:"precheck_mode,omitempty"`

	SkipPrecheckInfo *SkipPreCheckInfo `json:"skip_precheck_info,omitempty"`

	// 任务暂停模式。
	PauseMode *string `json:"pause_mode,omitempty"`

	// 任务定时启动时间。
	StartTime *string `json:"start_time,omitempty"`

	CompareTaskParam *CompareTaskParams `json:"compare_task_param,omitempty"`
}

func (o ActionParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionParams struct{}"
	}

	return strings.Join([]string{"ActionParams", string(data)}, " ")
}
