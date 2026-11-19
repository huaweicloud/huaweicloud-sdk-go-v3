package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSessionSummaryUseOpsTunnelResponse Response Object
type ShowSessionSummaryUseOpsTunnelResponse struct {

	// **参数解释**: 活跃会话数量。 **取值范围**: 不涉及。
	ActiveCount *int32 `json:"active_count,omitempty"`

	// **参数解释**: 会话总数量。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 最大SQL执行时间。单位：秒。 **取值范围**: 不涉及。
	MaxQueryTimeCost *int32 `json:"max_query_time_cost,omitempty"`

	// **参数解释**: 最大事务执行时间。单位：秒。 **取值范围**: 不涉及。
	MaxTransactionTimeCost *int32 `json:"max_transaction_time_cost,omitempty"`
	HttpStatusCode         int    `json:"-"`
}

func (o ShowSessionSummaryUseOpsTunnelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionSummaryUseOpsTunnelResponse struct{}"
	}

	return strings.Join([]string{"ShowSessionSummaryUseOpsTunnelResponse", string(data)}, " ")
}
