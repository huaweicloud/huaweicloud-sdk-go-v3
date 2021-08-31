package model

import (
	"encoding/json"

	"strings"
)

// CreateProjectRequestBody
type CreateProjectRequestBody struct {
	// name

	Name string `json:"name"`
	// description

	Description *string `json:"description,omitempty"`
}

func (o CreateProjectRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProjectRequestBody", string(data)}, " ")
}
