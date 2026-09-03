package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceBackupMethodRequestBody 切换实例备份方式请求体
type UpdateInstanceBackupMethodRequestBody struct {

	// **参数解释**：  修改的备份方式。Db为物理备份（OBS），EBackup为CBR快照备份。  **约束限制**：  不涉及。
	BackupMethod *string `json:"backup_method,omitempty"`
}

func (o UpdateInstanceBackupMethodRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBackupMethodRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBackupMethodRequestBody", string(data)}, " ")
}
