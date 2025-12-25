package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeVulStatusRequestInfoCustomBackupHosts struct {

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 存储库id **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	VaultId *string `json:"vault_id,omitempty"`

	// **参数解释**: 备份名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	BackupName *string `json:"backup_name,omitempty"`
}

func (o ChangeVulStatusRequestInfoCustomBackupHosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusRequestInfoCustomBackupHosts struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusRequestInfoCustomBackupHosts", string(data)}, " ")
}
