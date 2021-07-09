package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestartInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestartInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}
