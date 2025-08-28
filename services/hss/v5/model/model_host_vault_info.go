package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostVaultInfo 主机的存储库信息
type HostVaultInfo struct {

	// **参数解释**: 主机id **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机公网ip **取值范围**: 字符长度1-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 主机私网ip **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机的资产重要性 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 默认备份名称的后缀 **取值范围**: 字符长度0-32位
	DefaultBackupNameSuffix *string `json:"default_backup_name_suffix,omitempty"`

	DefaultVault *HostVaultInfoDefaultVault `json:"default_vault,omitempty"`

	// **参数解释**: 主机的存储库列表 **取值范围**: 不涉及
	Vaults *[]HostVaultInfoVaults `json:"vaults,omitempty"`
}

func (o HostVaultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVaultInfo struct{}"
	}

	return strings.Join([]string{"HostVaultInfo", string(data)}, " ")
}
