package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DeletePolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyAssignmentRequest", string(data)}, " ")
}
