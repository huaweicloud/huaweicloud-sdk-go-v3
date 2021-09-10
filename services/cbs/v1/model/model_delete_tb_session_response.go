package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTbSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTbSessionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTbSessionResponse struct{}"
	}

	return strings.Join([]string{"DeleteTbSessionResponse", string(data)}, " ")
}
