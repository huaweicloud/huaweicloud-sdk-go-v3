package model

import (
	"encoding/json"

	"strings"
)

//
type LoginTokenAssumedBy struct {
	User *LoginTokenUser `json:"user,omitempty"`
}

func (o LoginTokenAssumedBy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoginTokenAssumedBy struct{}"
	}

	return strings.Join([]string{"LoginTokenAssumedBy", string(data)}, " ")
}
