package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSqlExecutionPlanResponse struct {
	// SQL执行计划列表

	ExecutionPlans *[]ExecutionPlan `json:"execution_plans,omitempty"`
	// SQL执行失败时，显示执行错误信息

	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlExecutionPlanResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSqlExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlExecutionPlanResponse", string(data)}, " ")
}
