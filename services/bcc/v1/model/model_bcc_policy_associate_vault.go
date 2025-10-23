package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BccPolicyAssociateVault policy关联存储库vault
type BccPolicyAssociateVault struct {

	// 目标region vaultId
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	// vaultId
	VaultId *string `json:"vault_id,omitempty"`
}

func (o BccPolicyAssociateVault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BccPolicyAssociateVault struct{}"
	}

	return strings.Join([]string{"BccPolicyAssociateVault", string(data)}, " ")
}
