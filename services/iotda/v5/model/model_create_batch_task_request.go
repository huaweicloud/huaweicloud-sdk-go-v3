package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateBatchTaskRequest struct {
	InstanceId *string          `json:"Instance-Id,omitempty"`
	Body       *CreateBatchTask `json:"body,omitempty"`
}

func (o CreateBatchTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateBatchTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchTaskRequest", string(data)}, " ")
}
