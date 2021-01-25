package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CancelRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRemuxTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskResponse", string(data)}, " ")
}
