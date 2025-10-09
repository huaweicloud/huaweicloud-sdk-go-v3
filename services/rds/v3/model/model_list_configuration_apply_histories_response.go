package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationApplyHistoriesResponse Response Object
type ListConfigurationApplyHistoriesResponse struct {

	// **参数解释**：  参数组应用历史记录总数  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  应用历史详情列表  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Histories      *[]ApplyHistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListConfigurationApplyHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationApplyHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationApplyHistoriesResponse", string(data)}, " ")
}
