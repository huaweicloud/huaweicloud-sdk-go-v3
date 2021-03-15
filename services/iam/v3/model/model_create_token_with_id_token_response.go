package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTokenWithIdTokenResponse struct {
	Token          *ScopedTokenInfo `json:"token,omitempty"`
	XSubjectToken  *string          `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateTokenWithIdTokenResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTokenWithIdTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateTokenWithIdTokenResponse", string(data)}, " ")
}
