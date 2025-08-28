package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostVaultInfoDefaultVault **参数解释**: 默认选中的存储库 **取值范围**: 不涉及
type HostVaultInfoDefaultVault struct {

	// **参数解释**: 存储库名称 **取值范围**: 字符长度1-128位
	VaultName *string `json:"vault_name,omitempty"`

	// **参数解释**: 存储库id **取值范围**: 字符长度1-128位
	VaultId *string `json:"vault_id,omitempty"`
}

func (o HostVaultInfoDefaultVault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVaultInfoDefaultVault struct{}"
	}

	return strings.Join([]string{"HostVaultInfoDefaultVault", string(data)}, " ")
}
