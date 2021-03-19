package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateAgencyTokenRequest struct {
	Nocatalog *string `json:"nocatalog,omitempty"`

	Body *KeystoneCreateAgencyTokenRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateAgencyTokenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateAgencyTokenRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateAgencyTokenRequest", string(data)}, " ")
}
