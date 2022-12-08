package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetExecutionPlanResponse struct {

	// 执行计划元素的列表，只有当状态为'AVAILABLE'、'APPLIED'、’APPLY_IN_PROGRESS‘等完成创建后的状态才会被赋值，而'CREATION_IN_PROGRESS'或'CREATION_FAILED'会返回错误。
	ExecutionPlanItems *[]ExecutionPlanItem `json:"execution_plan_items,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o GetExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"GetExecutionPlanResponse", string(data)}, " ")
}
