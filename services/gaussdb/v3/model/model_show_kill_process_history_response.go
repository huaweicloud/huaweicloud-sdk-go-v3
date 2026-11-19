package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessHistoryResponse Response Object
type ShowKillProcessHistoryResponse struct {

	// **参数解释**：  历史记录总数。  **取值范围**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  历史记录列表。
	Processes      *[]KillProcessHistoryItem `json:"processes,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowKillProcessHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowKillProcessHistoryResponse", string(data)}, " ")
}
