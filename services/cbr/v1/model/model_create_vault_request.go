package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVaultRequest struct {
	Body *VaultCreateReq `json:"body,omitempty"`
}

func (o CreateVaultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVaultRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultRequest", string(data)}, " ")
}
