package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProcessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProcessResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProcessResponse struct{}"
	}

	return strings.Join([]string{"DeleteProcessResponse", string(data)}, " ")
}
