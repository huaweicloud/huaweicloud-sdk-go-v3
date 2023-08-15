package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateVaultResourceRequest Request Object
type MigrateVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultMigrateResourceReq `json:"body,omitempty"`
}

func (o MigrateVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceRequest", string(data)}, " ")
}
