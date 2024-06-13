package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRemediationExceptionsRequest Request Object
type BatchCreateRemediationExceptionsRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	Body *BatchCreateRemediationExceptionsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateRemediationExceptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRemediationExceptionsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateRemediationExceptionsRequest", string(data)}, " ")
}
