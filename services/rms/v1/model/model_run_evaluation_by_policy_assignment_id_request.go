package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunEvaluationByPolicyAssignmentIdRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o RunEvaluationByPolicyAssignmentIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEvaluationByPolicyAssignmentIdRequest struct{}"
	}

	return strings.Join([]string{"RunEvaluationByPolicyAssignmentIdRequest", string(data)}, " ")
}
