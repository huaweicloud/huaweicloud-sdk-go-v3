package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateProjectResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// project_id

	ProjectId      *int32 `json:"project_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectResponse", string(data)}, " ")
}
