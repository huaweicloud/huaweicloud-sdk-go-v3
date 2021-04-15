package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateVpcResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVpcResourceTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagResponse", string(data)}, " ")
}
