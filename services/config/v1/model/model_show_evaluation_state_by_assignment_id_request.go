package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEvaluationStateByAssignmentIdRequest Request Object
type ShowEvaluationStateByAssignmentIdRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowEvaluationStateByAssignmentIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationStateByAssignmentIdRequest struct{}"
	}

	return strings.Join([]string{"ShowEvaluationStateByAssignmentIdRequest", string(data)}, " ")
}
