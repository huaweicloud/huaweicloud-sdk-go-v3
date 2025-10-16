package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionTopSqlStatisticInfo struct {

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**: 归一化SQL ID。 **取值范围**: 不涉及。
	UniqueSqlId *string `json:"unique_sql_id,omitempty"`

	// **参数解释**: 查询语句。 **取值范围**: 不涉及。
	Query *string `json:"query,omitempty"`
}

func (o SessionTopSqlStatisticInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionTopSqlStatisticInfo struct{}"
	}

	return strings.Join([]string{"SessionTopSqlStatisticInfo", string(data)}, " ")
}
