package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopSqlsResponse Response Object
type ListTopSqlsResponse struct {

	// **参数解释**: Top SQL总条数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: Top SQL信息列表。
	TopSqlInfos    *[]TopSqlInfoResult `json:"top_sql_infos,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListTopSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListTopSqlsResponse", string(data)}, " ")
}
