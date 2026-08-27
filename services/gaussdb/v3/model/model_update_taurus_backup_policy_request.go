package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaurusBackupPolicyRequest 修改同区域备份策略请求体。
type UpdateTaurusBackupPolicyRequest struct {

	// **参数解释**：  备份时间段开始时间。  **约束限制**：  不涉及。  **取值范围**：  格式必须为hh:mm且有效，h为0~23的数字，m为0~59的数字，当前时间指UTC时间。  **默认取值**：  不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释**：  备份时间段结束时间。  **约束限制**：  end_time必须大于begin_time。  **取值范围**：  格式必须为hh:mm且有效，h为0~23的数字，m为0~59的数字，当前时间指UTC时间。  **默认取值**：  不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**：  一级备份保留数量。  **约束限制**：  当一级备份开关开启时，该参数必传。反之，不能传。  **取值范围**：  - 0：不保留一级备份。 - 1：一级备份保留数量，单位为个。  **默认取值**：  0。
	RetentionNumBackupLevel1 *int32 `json:"retention_num_backup_level1,omitempty"`

	// **参数解释**：  备份策略集，包含备份周期、保留天数和策略类型等配置信息，详见Policy数据结构。  **约束限制**：  不涉及。
	Policies []Policy `json:"policies"`
}

func (o UpdateTaurusBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaurusBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaurusBackupPolicyRequest", string(data)}, " ")
}
