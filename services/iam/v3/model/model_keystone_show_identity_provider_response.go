package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneShowIdentityProviderResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowIdentityProviderResponse", string(data)}, " ")
}
