package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTaskStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResponse", string(data)}, " ")
}
