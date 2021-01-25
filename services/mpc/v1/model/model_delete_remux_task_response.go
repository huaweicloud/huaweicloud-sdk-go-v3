package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRemuxTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteRemuxTaskResponse", string(data)}, " ")
}
