package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateUserRequest struct {
	Body *KeystoneCreateUserRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserRequest", string(data)}, " ")
}
