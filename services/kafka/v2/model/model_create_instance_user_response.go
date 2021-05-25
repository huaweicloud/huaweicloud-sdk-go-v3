package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateInstanceUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceUserResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserResponse", string(data)}, " ")
}
