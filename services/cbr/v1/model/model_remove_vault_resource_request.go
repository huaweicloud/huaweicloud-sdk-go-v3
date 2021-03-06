package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveVaultResourceRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`

	Body *VaultRemoveResourceReq `json:"body,omitempty"`
}

func (o RemoveVaultResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"RemoveVaultResourceRequest", string(data)}, " ")
}
