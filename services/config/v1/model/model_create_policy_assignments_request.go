package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyAssignmentsRequest Request Object
type CreatePolicyAssignmentsRequest struct {
	Body *PolicyAssignmentRequestBody `json:"body,omitempty"`
}

func (o CreatePolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyAssignmentsRequest", string(data)}, " ")
}
