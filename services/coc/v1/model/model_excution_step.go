package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExcutionStep struct {

	// 工单步骤id
	ExecutionStepId *string `json:"execution_step_id,omitempty"`

	// 原子能力action
	Action *string `json:"action,omitempty"`

	// 工单步骤开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 工单步骤结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 工单步骤执行信息
	Message *string `json:"message,omitempty"`

	// 工单步骤名称
	Name *string `json:"name,omitempty"`

	// 工单步骤执行状态
	Status *string `json:"status,omitempty"`

	// 步骤输入参数
	Inputs *[]ExcutionStepInputs `json:"inputs,omitempty"`

	// 步骤输出参数
	Outputs *[]ExcutionStepInputs `json:"outputs,omitempty"`

	// 工单步骤附加属性，map形式存储，如展示内网ip，则为{\"fixed_ip\": \"192.168.1.228\"}
	Properties map[string]string `json:"properties,omitempty"`
}

func (o ExcutionStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExcutionStep struct{}"
	}

	return strings.Join([]string{"ExcutionStep", string(data)}, " ")
}
