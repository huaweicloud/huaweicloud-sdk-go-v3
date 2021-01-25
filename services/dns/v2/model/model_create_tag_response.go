package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
