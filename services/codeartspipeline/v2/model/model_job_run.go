package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRun 任务运行信息
type JobRun struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务类型
	Category *string `json:"category,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 是否异步
	Async *string `json:"async,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务唯一标识
	Identifier *string `json:"identifier,omitempty"`

	// 依赖
	DependsOn *[]string `json:"depends_on,omitempty"`

	// 运行条件
	Condition *string `json:"condition,omitempty"`

	// 执行资源
	Resource *string `json:"resource,omitempty"`

	// 是否选中
	IsSelect *bool `json:"is_select,omitempty"`

	// 任务超时设置
	Timeout *string `json:"timeout,omitempty"`

	// 任务上次下发ID
	LastDispatchId *string `json:"last_dispatch_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 步骤
	Steps *[]StepRun `json:"steps,omitempty"`

	// 任务执行ID
	ExecId *string `json:"exec_id,omitempty"`
}

func (o JobRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRun struct{}"
	}

	return strings.Join([]string{"JobRun", string(data)}, " ")
}
