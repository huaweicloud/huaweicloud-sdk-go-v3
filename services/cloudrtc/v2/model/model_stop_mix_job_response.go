package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopMixJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMixJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopMixJobResponse struct{}"
	}

	return strings.Join([]string{"StopMixJobResponse", string(data)}, " ")
}
