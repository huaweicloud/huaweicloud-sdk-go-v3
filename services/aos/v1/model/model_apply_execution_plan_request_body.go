package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// apply_execution_plan request body
type ApplyExecutionPlanRequestBody struct {

	// 堆栈id
	StackId *string `json:"stack_id,omitempty"`

	// 执行计划Id，在domain_id+region+project_id+stack_id下应唯一
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`

	// 执行操作者的名字，用于审计工作
	Executor *string `json:"executor,omitempty"`
}

func (o ApplyExecutionPlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyExecutionPlanRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyExecutionPlanRequestBody", string(data)}, " ")
}
