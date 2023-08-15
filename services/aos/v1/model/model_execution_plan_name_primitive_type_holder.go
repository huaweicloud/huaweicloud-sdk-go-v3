package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlanNamePrimitiveTypeHolder struct {

	// 执行计划的名称。此名字在domain_id+区域+project_id+stack_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ExecutionPlanName string `json:"execution_plan_name"`
}

func (o ExecutionPlanNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutionPlanNamePrimitiveTypeHolder", string(data)}, " ")
}
