package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVaultTagsRequest struct {
	// 资源id

	VaultId string `json:"vault_id"`

	Body *VaultTagsCreateReq `json:"body,omitempty"`
}

func (o CreateVaultTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsRequest", string(data)}, " ")
}
