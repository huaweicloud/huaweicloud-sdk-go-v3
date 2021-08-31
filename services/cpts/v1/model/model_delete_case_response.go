package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteCaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteCaseResponse", string(data)}, " ")
}
