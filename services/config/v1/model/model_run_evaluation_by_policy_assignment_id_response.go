package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunEvaluationByPolicyAssignmentIdResponse Response Object
type RunEvaluationByPolicyAssignmentIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunEvaluationByPolicyAssignmentIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEvaluationByPolicyAssignmentIdResponse struct{}"
	}

	return strings.Join([]string{"RunEvaluationByPolicyAssignmentIdResponse", string(data)}, " ")
}
