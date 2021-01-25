package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateLoginTokenResponse struct {
	Logintoken         *LoginToken `json:"logintoken,omitempty"`
	XSubjectLoginToken *string     `json:"X-Subject-LoginToken,omitempty"`
	HttpStatusCode     int         `json:"-"`
}

func (o CreateLoginTokenResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoginTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateLoginTokenResponse", string(data)}, " ")
}
