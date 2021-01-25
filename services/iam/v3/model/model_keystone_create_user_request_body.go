package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateUserRequestBody struct {
	User *KeystoneCreateUserOption `json:"user"`
}

func (o KeystoneCreateUserRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserRequestBody", string(data)}, " ")
}
