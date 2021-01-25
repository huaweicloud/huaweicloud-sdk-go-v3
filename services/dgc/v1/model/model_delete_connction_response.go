package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteConnctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConnctionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteConnctionResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnctionResponse", string(data)}, " ")
}
