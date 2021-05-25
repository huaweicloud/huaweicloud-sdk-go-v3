package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVaultRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`
}

func (o ShowVaultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultRequest", string(data)}, " ")
}
