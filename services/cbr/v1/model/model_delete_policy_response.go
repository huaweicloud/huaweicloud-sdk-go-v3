package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePolicyResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyResponse", string(data)}, " ")
}
