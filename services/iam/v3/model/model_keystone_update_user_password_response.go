package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneUpdateUserPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneUpdateUserPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserPasswordResponse", string(data)}, " ")
}
