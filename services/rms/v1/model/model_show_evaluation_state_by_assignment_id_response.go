package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEvaluationStateByAssignmentIdResponse struct {

	// 规则ID
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty" xml:"policy_assignment_id"`

	// 评估任务执行状态
	State *string `json:"state,omitempty" xml:"state"`

	// 评估任务开始时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 评估任务结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 评估任务失败信息
	ErrorMessage   *string `json:"error_message,omitempty" xml:"error_message"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEvaluationStateByAssignmentIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationStateByAssignmentIdResponse struct{}"
	}

	return strings.Join([]string{"ShowEvaluationStateByAssignmentIdResponse", string(data)}, " ")
}
