package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTempRequest struct {
	Body *CreateTempRequestBody `json:"body,omitempty"`
}

func (o CreateTempRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTempRequest struct{}"
	}

	return strings.Join([]string{"CreateTempRequest", string(data)}, " ")
}
