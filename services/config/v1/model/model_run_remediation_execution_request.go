package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRemediationExecutionRequest Request Object
type RunRemediationExecutionRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	Body *RemediationRunRequestBody `json:"body,omitempty"`
}

func (o RunRemediationExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRemediationExecutionRequest struct{}"
	}

	return strings.Join([]string{"RunRemediationExecutionRequest", string(data)}, " ")
}
