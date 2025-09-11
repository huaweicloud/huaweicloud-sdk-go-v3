package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalSlowSqlDetailResponse Response Object
type ShowGlobalSlowSqlDetailResponse struct {

	// **参数解释**: 慢SQL详情列表。
	SlowSqlDetails *[]SlowSqlDetailResult `json:"slow_sql_details,omitempty"`

	// **参数解释**: 返回的慢SQL数量。 **取值范围**: 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGlobalSlowSqlDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalSlowSqlDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalSlowSqlDetailResponse", string(data)}, " ")
}
