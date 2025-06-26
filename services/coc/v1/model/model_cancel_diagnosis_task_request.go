package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelDiagnosisTaskRequest Request Object
type CancelDiagnosisTaskRequest struct {

	// 诊断任务ID
	TaskId string `json:"task_id"`
}

func (o CancelDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelDiagnosisTaskRequest", string(data)}, " ")
}
