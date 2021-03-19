package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateAndDeleteVaultTagsRequest struct {
	VaultId string `json:"vault_id"`

	Body *BulkCreateAndDeleteVaultTagsReq `json:"body,omitempty"`
}

func (o BatchCreateAndDeleteVaultTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsRequest", string(data)}, " ")
}
