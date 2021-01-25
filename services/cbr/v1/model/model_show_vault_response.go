package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultResponse", string(data)}, " ")
}
