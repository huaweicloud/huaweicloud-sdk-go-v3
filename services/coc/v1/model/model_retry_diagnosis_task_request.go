package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryDiagnosisTaskRequest Request Object
type RetryDiagnosisTaskRequest struct {

	// 诊断任务ID
	TaskId string `json:"task_id"`

	Body *RetryDiagnosisTaskRequestBody `json:"body,omitempty"`
}

func (o RetryDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"RetryDiagnosisTaskRequest", string(data)}, " ")
}
