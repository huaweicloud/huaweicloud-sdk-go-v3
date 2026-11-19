package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetainedBackupPolicy 保留策略信息。
type RetainedBackupPolicy struct {

	// **参数解释**：  实例释放保留策略。  **约束限制**：  不涉及。  **默认取值**：  不涉及。
	BackupRetentionStrategy *string `json:"backup_retention_strategy,omitempty"`
}

func (o RetainedBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetainedBackupPolicy struct{}"
	}

	return strings.Join([]string{"RetainedBackupPolicy", string(data)}, " ")
}
