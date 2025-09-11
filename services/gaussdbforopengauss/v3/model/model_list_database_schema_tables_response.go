package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseSchemaTablesResponse Response Object
type ListDatabaseSchemaTablesResponse struct {

	// **参数解释**： 列表中每个元素表示一个数据库表。
	DatabaseTables *[]DatabaseForListTableResult `json:"database_tables,omitempty"`

	// **参数解释**： 数据库表总数。 **取值范围**： 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseSchemaTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseSchemaTablesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseSchemaTablesResponse", string(data)}, " ")
}
