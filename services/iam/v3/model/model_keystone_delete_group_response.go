package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneDeleteGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteGroupResponse", string(data)}, " ")
}
