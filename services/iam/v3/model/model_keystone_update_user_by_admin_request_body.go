package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateUserByAdminRequestBody struct {
	User *KeystoneUpdateUserOption `json:"user"`
}

func (o KeystoneUpdateUserByAdminRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserByAdminRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserByAdminRequestBody", string(data)}, " ")
}
