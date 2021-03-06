package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateScopedTokenRequestBody struct {
	Auth *ScopedTokenAuth `json:"auth"`
}

func (o KeystoneCreateScopedTokenRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateScopedTokenRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateScopedTokenRequestBody", string(data)}, " ")
}
