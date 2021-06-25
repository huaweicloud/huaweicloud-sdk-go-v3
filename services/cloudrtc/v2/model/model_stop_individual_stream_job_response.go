package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopIndividualStreamJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopIndividualStreamJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopIndividualStreamJobResponse struct{}"
	}

	return strings.Join([]string{"StopIndividualStreamJobResponse", string(data)}, " ")
}
