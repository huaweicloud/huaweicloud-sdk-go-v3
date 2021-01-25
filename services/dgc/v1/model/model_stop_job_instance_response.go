package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopJobInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopJobInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopJobInstanceResponse", string(data)}, " ")
}
