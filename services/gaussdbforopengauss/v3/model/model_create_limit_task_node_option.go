package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLimitTaskNodeOption struct {

	// **参数解释**: 节点ID。 **约束限制**: 必须是当前实例的某一个节点ID。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 该节点执行的SQL语句ID。 **约束限制**: 如果“limit_type”为SQL_ID，必须与“limit_type_value”值一致。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId string `json:"sql_id"`
}

func (o CreateLimitTaskNodeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskNodeOption struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskNodeOption", string(data)}, " ")
}
