package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDefaultMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDefaultMigprojectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectResponse", string(data)}, " ")
}
