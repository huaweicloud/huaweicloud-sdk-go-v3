package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTestCaseInPlanRequest struct {
	// 计划唯一标识，固定长度32位字符

	PlanId string `json:"plan_id"`

	Body *CreateTestCaseInPlanRequestBody `json:"body,omitempty"`
}

func (o CreateTestCaseInPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseInPlanRequest struct{}"
	}

	return strings.Join([]string{"CreateTestCaseInPlanRequest", string(data)}, " ")
}
