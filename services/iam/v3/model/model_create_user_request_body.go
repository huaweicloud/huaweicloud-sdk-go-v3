package model

import (
	"encoding/json"

	"strings"
)

//
type CreateUserRequestBody struct {
	User *CreateUserOption `json:"user"`
}

func (o CreateUserRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUserRequestBody", string(data)}, " ")
}
