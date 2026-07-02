package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupVaultLockInfo struct {

	// **参数解释**：  合规锁保留期，单位是天。  **取值范围**：  1-36500。
	LockRetentionDays int32 `json:"lock_retention_days"`

	// **参数解释**：  合规锁配置策略。  **取值范围**：  - true：已开启合规锁配置。 - false: 未开启合规锁配置。
	LockPolicy bool `json:"lock_policy"`

	// **参数解释**：  合规锁开始时间，时间戳格式。  **取值范围**：  不涉及。
	LockStartTime int64 `json:"lock_start_time"`

	// **参数解释**：  合规锁结束时间，时间戳格式。  **取值范围**：  不涉及。
	LockEndTime int64 `json:"lock_end_time"`
}

func (o BackupVaultLockInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupVaultLockInfo struct{}"
	}

	return strings.Join([]string{"BackupVaultLockInfo", string(data)}, " ")
}
