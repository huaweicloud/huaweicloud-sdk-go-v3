package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowSqlDetailsResponse Response Object
type ListSlowSqlDetailsResponse struct {

	// **参数解释**: 总条数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 慢SQL详情信息列表。
	SlowSqlDetails *[]SlowSqlDetailResult `json:"slow_sql_details,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSlowSqlDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowSqlDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowSqlDetailsResponse", string(data)}, " ")
}
