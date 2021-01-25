package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateUserTokenByPasswordRequestBody struct {
	Auth *PwdAuth `json:"auth"`
}

func (o KeystoneCreateUserTokenByPasswordRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordRequestBody", string(data)}, " ")
}
