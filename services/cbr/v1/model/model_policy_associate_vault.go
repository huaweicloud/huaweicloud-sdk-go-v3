package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyAssociateVault
type PolicyAssociateVault struct {

	// 关联的远端存储库ID
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	// 存储库ID
	VaultId string `json:"vault_id"`
}

func (o PolicyAssociateVault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssociateVault struct{}"
	}

	return strings.Join([]string{"PolicyAssociateVault", string(data)}, " ")
}
