package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnosisTaskRequest Request Object
type CreateDiagnosisTaskRequest struct {
	Body *DiagnosisTaskSubmitBody `json:"body,omitempty"`
}

func (o CreateDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskRequest", string(data)}, " ")
}
