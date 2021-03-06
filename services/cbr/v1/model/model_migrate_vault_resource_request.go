package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type MigrateVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultMigrateResourceReq `json:"body,omitempty"`
}

func (o MigrateVaultResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceRequest", string(data)}, " ")
}
