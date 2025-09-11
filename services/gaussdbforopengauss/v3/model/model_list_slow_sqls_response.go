package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowSqlsResponse Response Object
type ListSlowSqlsResponse struct {

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 慢SQL列表。
	SlowSqlInfos   *[]SlowSqlInfoResult `json:"slow_sql_infos,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSlowSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowSqlsResponse", string(data)}, " ")
}
