package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行计划的更多细节
type ExecutionPlanDiffAttribute struct {

	// 当前资源被修改的参数的名字。
	Name *string `json:"name,omitempty"`

	// 当前资源被修改前参数的值，新创建时为空。
	PreviousValue *string `json:"previous_value,omitempty"`

	// 当前资源被修改的参数的目的值，删除时为空。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ExecutionPlanDiffAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanDiffAttribute struct{}"
	}

	return strings.Join([]string{"ExecutionPlanDiffAttribute", string(data)}, " ")
}
