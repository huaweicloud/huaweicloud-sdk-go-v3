package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepRun 步骤运行信息
type StepRun struct {

	// 步骤名称
	Name *string `json:"name,omitempty"`

	// 步骤插件
	Task *string `json:"task,omitempty"`

	// 插件业务类型
	BusinessType *string `json:"business_type,omitempty"`

	// 输入参数
	Inputs *[]StepRunInputs `json:"inputs,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 官方插件版本号
	OfficialTaskVersion *string `json:"official_task_version,omitempty"`

	// 唯一标识符
	Identifier *string `json:"identifier,omitempty"`

	// 是否可编辑
	MultiStepEditable *int32 `json:"multi_step_editable,omitempty"`

	// 步骤ID
	Id *string `json:"id,omitempty"`

	// 扩展点
	EndpointIds *[]string `json:"endpoint_ids,omitempty"`

	// 上次下发任务ID
	LastDispatchId *string `json:"last_dispatch_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误消息
	Message *string `json:"message,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StepRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepRun struct{}"
	}

	return strings.Join([]string{"StepRun", string(data)}, " ")
}
