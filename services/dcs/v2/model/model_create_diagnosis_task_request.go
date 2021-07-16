package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDiagnosisTaskRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateDiagnosisTaskBody `json:"body,omitempty"`
}

func (o CreateDiagnosisTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskRequest", string(data)}, " ")
}
