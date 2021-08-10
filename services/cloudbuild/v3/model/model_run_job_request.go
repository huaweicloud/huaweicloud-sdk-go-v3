package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunJobRequest struct {
	Body *RunJobRequestBody `json:"body,omitempty"`
}

func (o RunJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunJobRequest struct{}"
	}

	return strings.Join([]string{"RunJobRequest", string(data)}, " ")
}
