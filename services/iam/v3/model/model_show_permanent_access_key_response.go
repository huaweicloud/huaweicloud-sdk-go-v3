package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPermanentAccessKeyResponse struct {
	Credential     *ShowCredential `json:"credential,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowPermanentAccessKeyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowPermanentAccessKeyResponse", string(data)}, " ")
}
