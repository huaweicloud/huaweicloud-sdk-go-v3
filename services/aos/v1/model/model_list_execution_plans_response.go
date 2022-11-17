package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListExecutionPlansResponse struct {

	// 执行计划的元数据
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
