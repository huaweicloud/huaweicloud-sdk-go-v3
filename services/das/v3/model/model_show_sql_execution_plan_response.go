package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSqlExecutionPlanResponse struct {

	// SQL执行计划列表
	ExecutionPlans *[]ExecutionPlan `json:"execution_plans,omitempty" xml:"execution_plans"`

	// SQL执行失败时，显示执行错误信息
	ErrorMessage   *string `json:"error_message,omitempty" xml:"error_message"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlExecutionPlanResponse", string(data)}, " ")
}
