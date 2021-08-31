package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateProjectRequest struct {
	Body *CreateProjectRequestBody `json:"body,omitempty"`
}

func (o CreateProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRequest", string(data)}, " ")
}
