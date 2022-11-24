package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateExecutionPlanResponse struct {

	// 栈的名字
	StackName *string `json:"stack_name,omitempty"`

	// 栈的唯一Id
	StackId *string `json:"stack_id,omitempty"`

	// 执行计划的名字。如果未指定，则使用execution_plan_id作为execution_plan_name。
	ExecutionPlanName *string `json:"execution_plan_name,omitempty"`

	// 执行计划的唯一Id，由IaC随机生成
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateExecutionPlanResponse", string(data)}, " ")
}
