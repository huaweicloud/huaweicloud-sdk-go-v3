package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateCaseResponse", string(data)}, " ")
}
