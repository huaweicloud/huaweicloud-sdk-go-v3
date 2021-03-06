package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateUserResponse struct {
	User           *KeystoneCreateUserResult `json:"user,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o KeystoneCreateUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserResponse", string(data)}, " ")
}
