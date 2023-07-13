package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlExplainResponse Response Object
type ShowSqlExplainResponse struct {

	// SQL执行计划列表
	ExecutionPlans *[]ExecutionPlan `json:"execution_plans,omitempty"`

	// SQL执行失败时，显示执行错误信息
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlExplainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlExplainResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlExplainResponse", string(data)}, " ")
}
