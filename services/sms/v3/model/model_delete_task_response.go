package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskResponse", string(data)}, " ")
}
