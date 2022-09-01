package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type MigrateVaultResourceRequest struct {
	VaultId string `json:"vault_id" xml:"vault_id"`

	Body *VaultMigrateResourceReq `json:"body,omitempty" xml:"body"`
}

func (o MigrateVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceRequest", string(data)}, " ")
}
