package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAutoTerminateTimeServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAutoTerminateTimeServerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAutoTerminateTimeServerResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoTerminateTimeServerResponse", string(data)}, " ")
}
