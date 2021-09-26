package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEditingJobRequest struct {
	Body *CreateEditingJobReq `json:"body,omitempty"`
}

func (o CreateEditingJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEditingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateEditingJobRequest", string(data)}, " ")
}
