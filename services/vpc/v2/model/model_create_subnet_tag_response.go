package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubnetTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagResponse", string(data)}, " ")
}
