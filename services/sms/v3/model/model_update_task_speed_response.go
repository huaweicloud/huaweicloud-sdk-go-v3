package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTaskSpeedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskSpeedResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskSpeedResponse", string(data)}, " ")
}
