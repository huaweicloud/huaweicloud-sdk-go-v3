package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`

	Body *PolicyAssignmentRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyAssignmentRequest", string(data)}, " ")
}
