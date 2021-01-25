package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
