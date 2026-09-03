package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceBackupMethodResponse Response Object
type UpdateInstanceBackupMethodResponse struct {

	// **参数解释**：  成功修改后的备份方式。  **约束限制**：  不涉及。
	BackupMethod   *string `json:"backup_method,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceBackupMethodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBackupMethodResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBackupMethodResponse", string(data)}, " ")
}
