package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetExecutionPlanResponse Response Object
type GetExecutionPlanResponse struct {

	// 执行计划项目的列表
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
