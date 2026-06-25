package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceBackupSummaryResponse Response Object
type ListInstanceBackupSummaryResponse struct {

	// **参数解释**：  实例备份概览列表  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Infos *[]InstanceBackupSummary `json:"infos,omitempty"`

	// **参数解释**：  总记录数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceBackupSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceBackupSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceBackupSummaryResponse", string(data)}, " ")
}
