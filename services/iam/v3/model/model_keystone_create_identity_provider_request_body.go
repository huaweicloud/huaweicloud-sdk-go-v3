package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateIdentityProviderRequestBody struct {
	IdentityProvider *IdentityproviderOption `json:"identity_provider"`
}

func (o KeystoneCreateIdentityProviderRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderRequestBody", string(data)}, " ")
}
