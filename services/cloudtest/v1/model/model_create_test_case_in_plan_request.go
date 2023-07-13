package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTestCaseInPlanRequest Request Object
type CreateTestCaseInPlanRequest struct {

	// 计划唯一标识，长度11-34位字符
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
