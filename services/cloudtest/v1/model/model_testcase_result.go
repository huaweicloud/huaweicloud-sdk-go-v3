package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestcaseResult 测试用例结果
type TestcaseResult struct {

	// 结果
	ExecuteResultId *string `json:"execute_result_id,omitempty"`

	// 测试用例状态
	ExecuteStatus *string `json:"execute_status,omitempty"`

	// 失败原因
	FailureCause *string `json:"failure_cause,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 执行测试用例用户id
	ExecutorId *string `json:"executor_id,omitempty"`

	// 执行测试用例用户name
	ExecutorName *string `json:"executor_name,omitempty"`
}

func (o TestcaseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestcaseResult struct{}"
	}

	return strings.Join([]string{"TestcaseResult", string(data)}, " ")
}
