package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePlanRequest struct {
	Body *CreatePlanRequestBody `json:"body,omitempty"`
}

func (o CreatePlanRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePlanRequest struct{}"
	}

	return strings.Join([]string{"CreatePlanRequest", string(data)}, " ")
}
