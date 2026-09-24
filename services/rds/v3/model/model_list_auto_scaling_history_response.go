package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoScalingHistoryResponse Response Object
type ListAutoScalingHistoryResponse struct {

	// **参数解释**：  记录条数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  记录数组。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Records        *[]MysqlAutoScalingRecord `json:"records,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAutoScalingHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoScalingHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListAutoScalingHistoryResponse", string(data)}, " ")
}
