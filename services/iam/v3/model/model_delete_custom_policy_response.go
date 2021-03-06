package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteCustomPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomPolicyResponse", string(data)}, " ")
}
