package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerResponse", string(data)}, " ")
}
