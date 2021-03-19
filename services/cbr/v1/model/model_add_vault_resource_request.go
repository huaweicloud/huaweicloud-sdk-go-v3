package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultAddResourceReq `json:"body,omitempty"`
}

func (o AddVaultResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"AddVaultResourceRequest", string(data)}, " ")
}
