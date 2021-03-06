package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateProjectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProjectStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProjectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectStatusResponse", string(data)}, " ")
}
