package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupVaultInfo struct {

	// 存储库ID
	VaultId *string `json:"vault_id,omitempty"`

	// 存储库名称
	VaultName *string `json:"vault_name,omitempty"`

	// 存储库总容量，单位GB
	VaultSize *int32 `json:"vault_size,omitempty"`

	// 已使用容量，单位MB，指的是已有备份占用的容量，例如绑定了1台主机，已经有两个备份数，两个备份60G,则已使用容量为60G。
	VaultUsed *int32 `json:"vault_used,omitempty"`

	// 已分配容量，单位GB，指绑定的服务器大小，例如绑定了1台主机，主机大小40G，则已分配容量为40G。
	VaultAllocated *int32 `json:"vault_allocated,omitempty"`

	// 存储库创建模式：   - 按需：post_paid   - 包周期：pre_paid
	VaultChargingMode *string `json:"vault_charging_mode,omitempty"`

	// 存储库状态。   - available：可用。   - lock：被锁定。   - frozen：冻结。   - deleting：删除中。   - error：错误。
	VaultStatus *string `json:"vault_status,omitempty"`

	// 备份策略ID，若为空，则为未绑定状态，若不为空，通过backup_policy_enabled字段判断策略是否启用
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	// 备份策略名称
	BackupPolicyName *string `json:"backup_policy_name,omitempty"`

	// 备份策略是否启用
	BackupPolicyEnabled *bool `json:"backup_policy_enabled,omitempty"`

	// 已绑定服务器（个）
	ResourcesNum *int32 `json:"resources_num,omitempty"`
}

func (o BackupVaultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupVaultInfo struct{}"
	}

	return strings.Join([]string{"BackupVaultInfo", string(data)}, " ")
}
