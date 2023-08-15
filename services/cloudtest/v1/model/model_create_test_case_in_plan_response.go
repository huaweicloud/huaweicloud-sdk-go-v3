package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTestCaseInPlanResponse Response Object
type CreateTestCaseInPlanResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTestCaseInPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseInPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateTestCaseInPlanResponse", string(data)}, " ")
}
