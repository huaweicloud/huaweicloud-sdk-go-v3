package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePermanentAccessKeyResponse struct {
	Credential     *CreateCredentialResult `json:"credential,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreatePermanentAccessKeyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"CreatePermanentAccessKeyResponse", string(data)}, " ")
}
