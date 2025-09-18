package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StateInfo 任务基本信息
type StateInfo struct {

	// 任务执行状态
	Status *string `json:"status,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 执行记录id
	RecordId *string `json:"record_id,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 步骤状态
	Step *string `json:"step,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 执行人
	Executor *string `json:"executor,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 子步骤信息
	StepState *[]StepInfo `json:"step_state,omitempty"`
}

func (o StateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StateInfo struct{}"
	}

	return strings.Join([]string{"StateInfo", string(data)}, " ")
}
