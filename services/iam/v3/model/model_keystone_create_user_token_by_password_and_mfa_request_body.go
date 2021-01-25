package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateUserTokenByPasswordAndMfaRequestBody struct {
	Auth *MfaAuth `json:"auth"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordAndMfaRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaRequestBody", string(data)}, " ")
}
