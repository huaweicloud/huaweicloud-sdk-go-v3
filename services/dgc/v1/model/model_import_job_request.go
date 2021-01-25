package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportJobRequest struct {
	Body *ImportFileReq `json:"body,omitempty"`
}

func (o ImportJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportJobRequest struct{}"
	}

	return strings.Join([]string{"ImportJobRequest", string(data)}, " ")
}
