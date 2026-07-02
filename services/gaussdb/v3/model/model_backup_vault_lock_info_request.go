package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupVaultLockInfoRequest struct {

	// **参数解释**：  合规锁保留期，单位是天。  **约束限制**：  仅支持按天配置。  **取值范围**：  1-36500。  **默认取值**：  1。
	LockRetentionDays int32 `json:"lock_retention_days"`

	// **参数解释**：  合规锁配置策略。  **约束限制**：  不涉及。  **取值范围**：  当前仅支持设置为true，表示开启或延期合规锁。  **默认取值**：  true。
	LockPolicy bool `json:"lock_policy"`
}

func (o BackupVaultLockInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupVaultLockInfoRequest struct{}"
	}

	return strings.Join([]string{"BackupVaultLockInfoRequest", string(data)}, " ")
}
