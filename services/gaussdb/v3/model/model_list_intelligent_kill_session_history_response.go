package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIntelligentKillSessionHistoryResponse Response Object
type ListIntelligentKillSessionHistoryResponse struct {

	// **参数解释**：  智能Kill会话历史记录列表。
	IntelligentKillSessionHistories *[]IntelligentKillSessionHistory `json:"intelligent_kill_session_histories,omitempty"`

	// **参数解释**：  数据总数。  **取值范围**：  不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIntelligentKillSessionHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIntelligentKillSessionHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListIntelligentKillSessionHistoryResponse", string(data)}, " ")
}
