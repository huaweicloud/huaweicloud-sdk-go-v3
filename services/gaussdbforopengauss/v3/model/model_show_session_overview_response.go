package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSessionOverviewResponse Response Object
type ShowSessionOverviewResponse struct {

	// **参数解释**: 实时会话的活跃会话数。 **取值范围**: 不涉及。
	ActiveNum *string `json:"active_num,omitempty"`

	// **参数解释**: 实时会话的总会话数。 **取值范围**: 不涉及。
	TotalNum *string `json:"total_num,omitempty"`

	// **参数解释**: 实时会话中的慢SQL数量。 **取值范围**: 不涉及。
	SlowSqlNum *string `json:"slow_sql_num,omitempty"`

	// **参数解释**: 实时会话中的锁等会话数。 **取值范围**: 不涉及。
	LockNum        *string `json:"lock_num,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSessionOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowSessionOverviewResponse", string(data)}, " ")
}
