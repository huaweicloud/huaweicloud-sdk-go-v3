package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSessionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSessionResponse struct{}"
	}

	return strings.Join([]string{"DeleteSessionResponse", string(data)}, " ")
}
