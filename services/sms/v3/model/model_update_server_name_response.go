package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateServerNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateServerNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerNameResponse", string(data)}, " ")
}
