package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateIdentityProviderRequest struct {
	// 待注册的身份提供商ID。

	Id string `json:"id"`

	Body *KeystoneCreateIdentityProviderRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateIdentityProviderRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderRequest", string(data)}, " ")
}
