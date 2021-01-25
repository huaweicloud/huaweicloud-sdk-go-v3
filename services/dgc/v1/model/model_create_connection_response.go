package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
