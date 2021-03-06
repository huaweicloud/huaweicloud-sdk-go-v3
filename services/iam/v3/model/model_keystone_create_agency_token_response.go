package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateAgencyTokenResponse struct {
	Token *AgencyTokenResult `json:"token,omitempty"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneCreateAgencyTokenResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateAgencyTokenResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateAgencyTokenResponse", string(data)}, " ")
}
