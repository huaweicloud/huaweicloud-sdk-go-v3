package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionPlansResponse Response Object
type ListExecutionPlansResponse struct {

	// 执行计划列表。默认按照生成时间降序排序，最新生成的在最前
	ExecutionPlans *[]ExecutionPlan `json:"execution_plans,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListExecutionPlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionPlansResponse struct{}"
	}

	return strings.Join([]string{"ListExecutionPlansResponse", string(data)}, " ")
}
