package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteL7policyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7policyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7policyResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7policyResponse", string(data)}, " ")
}
