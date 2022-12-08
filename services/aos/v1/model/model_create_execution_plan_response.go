package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateExecutionPlanResponse struct {

	// 执行计划ID
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
