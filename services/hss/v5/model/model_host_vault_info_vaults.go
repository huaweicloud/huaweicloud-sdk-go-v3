package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostVaultInfoVaults **参数解释**: 主机的存储库信息 **取值范围**: 不涉及
type HostVaultInfoVaults struct {

	// **参数解释**: 存储库名称 **取值范围**: 字符长度1-128位
	VaultName *string `json:"vault_name,omitempty"`

	// **参数解释**: 存储库id **取值范围**: 字符长度1-128位
	VaultId *string `json:"vault_id,omitempty"`
}

func (o HostVaultInfoVaults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVaultInfoVaults struct{}"
	}

	return strings.Join([]string{"HostVaultInfoVaults", string(data)}, " ")
}
