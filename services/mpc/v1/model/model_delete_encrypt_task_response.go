package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEncryptTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEncryptTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskResponse", string(data)}, " ")
}
