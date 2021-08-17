package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteServersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServersResponse struct{}"
	}

	return strings.Join([]string{"DeleteServersResponse", string(data)}, " ")
}
