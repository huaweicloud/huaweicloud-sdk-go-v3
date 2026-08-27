package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaurusDbAdvancedBackupPolicyResponse Response Object
type ShowTaurusDbAdvancedBackupPolicyResponse struct {

	// **参数解释**：  备份时间段开始时间。  **取值范围**：  非空，格式必须为hh:mm且有效，当前时间指UTC时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  备份时间段结束时间。  **取值范围**：  非空，格式必须为hh:mm且有效，当前时间指UTC时间。end_time必须大于begin_time。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  一级备份保留数量。当一级备份开关开启时，返回此参数。  **取值范围**：  不涉及。
	RetentionNumBackupLevel1 *int32 `json:"retention_num_backup_level1,omitempty"`

	// **参数解释**：  备份策略集。
	Policies       *[]BackupPolicyInfo `json:"policies,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowTaurusDbAdvancedBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaurusDbAdvancedBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowTaurusDbAdvancedBackupPolicyResponse", string(data)}, " ")
}
