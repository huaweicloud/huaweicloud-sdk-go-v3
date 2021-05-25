package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateVaultRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`

	Body *VaultUpdateReq `json:"body,omitempty"`
}

func (o UpdateVaultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVaultRequest struct{}"
	}

	return strings.Join([]string{"UpdateVaultRequest", string(data)}, " ")
}
