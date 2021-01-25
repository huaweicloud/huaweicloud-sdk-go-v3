package model

import (
	"encoding/json"

	"strings"
)

//
type CreateTemporaryAccessKeyByTokenRequestBody struct {
	Auth *TokenAuth `json:"auth"`
}

func (o CreateTemporaryAccessKeyByTokenRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByTokenRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenRequestBody", string(data)}, " ")
}
