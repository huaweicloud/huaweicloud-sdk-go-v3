package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMigprojectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMigprojectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateMigprojectResponse", string(data)}, " ")
}
