package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteBackupResult struct {

	// **参数解释**：  备份ID。  **取值范围**：  不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释**：  提交任务异常时返回的编码。  **取值范围**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  提交任务异常时返回的描述信息。  **取值范围**：  不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DeleteBackupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupResult struct{}"
	}

	return strings.Join([]string{"DeleteBackupResult", string(data)}, " ")
}
