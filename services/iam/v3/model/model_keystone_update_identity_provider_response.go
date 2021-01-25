package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneUpdateIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneUpdateIdentityProviderResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateIdentityProviderResponse", string(data)}, " ")
}
