package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubnetTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubnetTagResponse", string(data)}, " ")
}
