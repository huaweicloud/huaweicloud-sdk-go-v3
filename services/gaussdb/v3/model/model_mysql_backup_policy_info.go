package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlBackupPolicyInfo struct {

	// **参数解释**：  备份时间段开始时间。  **约束限制**：  不涉及。  **取值范围**：  非空，格式必须为hh:mm且有效，当前时间指UTC时间。  **默认取值**：  不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释**：  备份时间段结束时间。  **约束限制**：  end_time必须大于begin_time。  **取值范围**：  非空，格式必须为hh:mm且有效，当前时间指UTC时间。  **默认取值**：  不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**：  一级备份保留数量。  **约束限制**：  当一级备份开关开启时，该参数必传。反之，不能传。  **取值范围**：  - 0：不保留一级备份。 - 1：保留1个一级备份。  **默认取值**：  0。
	RetentionNumBackupLevel1 *int32 `json:"retention_num_backup_level1,omitempty"`

	// **参数解释**：  备份策略集。  **约束限制**：  不涉及。
	Policies []PolicyInfo `json:"policies"`
}

func (o MysqlBackupPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlBackupPolicyInfo struct{}"
	}

	return strings.Join([]string{"MysqlBackupPolicyInfo", string(data)}, " ")
}
