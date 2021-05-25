package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVaultTagRequest struct {
	// 资源id

	VaultId string `json:"vault_id"`
}

func (o ShowVaultTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultTagRequest", string(data)}, " ")
}
