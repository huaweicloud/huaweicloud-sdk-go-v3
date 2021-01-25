package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComponentResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteComponentResponse struct{}"
	}

	return strings.Join([]string{"DeleteComponentResponse", string(data)}, " ")
}
