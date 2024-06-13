package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRemediationExceptionsRequest Request Object
type BatchDeleteRemediationExceptionsRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	Body *BatchDeleteRemediationExceptionsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRemediationExceptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRemediationExceptionsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRemediationExceptionsRequest", string(data)}, " ")
}
