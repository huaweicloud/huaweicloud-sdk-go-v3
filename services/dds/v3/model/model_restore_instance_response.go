package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreInstanceResponse", string(data)}, " ")
}
