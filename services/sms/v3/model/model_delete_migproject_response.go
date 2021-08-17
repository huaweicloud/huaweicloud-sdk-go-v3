package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMigprojectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMigprojectResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectResponse", string(data)}, " ")
}
