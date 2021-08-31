package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectResponse", string(data)}, " ")
}
