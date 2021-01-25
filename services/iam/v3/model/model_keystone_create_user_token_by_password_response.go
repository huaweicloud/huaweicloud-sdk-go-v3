package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateUserTokenByPasswordResponse struct {
	Token          *TokenResult `json:"token,omitempty"`
	XSubjectToken  *string      `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o KeystoneCreateUserTokenByPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordResponse", string(data)}, " ")
}
