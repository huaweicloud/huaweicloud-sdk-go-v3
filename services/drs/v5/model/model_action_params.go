package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionParams 操作任务动作参数。
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

	// 再编辑任务启动时取值true。
	IsSyncReEdit *bool `json:"is_sync_re_edit,omitempty"`

	// 是否支持只初始化任务。仅支持白名单用户使用，需要提交工单申请才能使用。
	IsOnlyInitTask *bool `json:"is_only_init_task,omitempty"`

	// 强制结束时取值为true。
	ForceDelete *bool `json:"force_delete,omitempty"`

	PublicIpConfig *PublicIpConfig `json:"public_ip_config,omitempty"`

	ReplayConfig *ReplayConfigInfo `json:"replay_config,omitempty"`
}

func (o ActionParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionParams struct{}"
	}

	return strings.Join([]string{"ActionParams", string(data)}, " ")
}
