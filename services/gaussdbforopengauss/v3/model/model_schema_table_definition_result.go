package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchemaTableDefinitionResult struct {

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 数据库表定义信息。 **取值范围**: 不涉及。
	TableDefinition *string `json:"table_definition,omitempty"`
}

func (o SchemaTableDefinitionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaTableDefinitionResult struct{}"
	}

	return strings.Join([]string{"SchemaTableDefinitionResult", string(data)}, " ")
}
