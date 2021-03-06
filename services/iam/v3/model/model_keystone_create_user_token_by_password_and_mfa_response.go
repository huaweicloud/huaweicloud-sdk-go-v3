package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateUserTokenByPasswordAndMfaResponse struct {
	Token *TokenResult `json:"token,omitempty"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordAndMfaResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaResponse", string(data)}, " ")
}
