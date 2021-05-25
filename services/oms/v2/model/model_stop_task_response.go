package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopTaskResponse struct{}"
	}

	return strings.Join([]string{"StopTaskResponse", string(data)}, " ")
}
