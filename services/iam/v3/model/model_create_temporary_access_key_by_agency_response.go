package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTemporaryAccessKeyByAgencyResponse struct {
	Credential     *Credential `json:"credential,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateTemporaryAccessKeyByAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyResponse", string(data)}, " ")
}
