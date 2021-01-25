package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateUnscopeTokenByIdpInitiatedResponse struct {
	Token          *IdpToken `json:"token,omitempty"`
	XSubjectToken  *string   `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateUnscopeTokenByIdpInitiatedResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUnscopeTokenByIdpInitiatedResponse struct{}"
	}

	return strings.Join([]string{"CreateUnscopeTokenByIdpInitiatedResponse", string(data)}, " ")
}
