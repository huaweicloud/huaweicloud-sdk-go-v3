package model

import (
	"encoding/json"

	"strings"
)

// CreateTempRequestBody
type CreateTempRequestBody struct {
	// project_id

	ProjectId int32 `json:"project_id"`
	// temp_type

	TempType int32 `json:"temp_type"`
	// name

	Name string `json:"name"`
	// description

	Description *string `json:"description,omitempty"`
}

func (o CreateTempRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTempRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTempRequestBody", string(data)}, " ")
}
