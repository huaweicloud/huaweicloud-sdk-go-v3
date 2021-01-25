package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreJobInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreJobInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreJobInstanceResponse", string(data)}, " ")
}
