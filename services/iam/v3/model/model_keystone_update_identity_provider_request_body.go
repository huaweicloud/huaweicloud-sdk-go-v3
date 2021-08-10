package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateIdentityProviderRequestBody struct {
	IdentityProvider *UpdateIdentityproviderOption `json:"identity_provider"`
}

func (o KeystoneUpdateIdentityProviderRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateIdentityProviderRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateIdentityProviderRequestBody", string(data)}, " ")
}
