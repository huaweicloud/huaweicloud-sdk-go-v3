package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateProjectStatusRequest struct {
	ProjectId string `json:"project_id"`

	Body *UpdateProjectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateProjectStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProjectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectStatusRequest", string(data)}, " ")
}
