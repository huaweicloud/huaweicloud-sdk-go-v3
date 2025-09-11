package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseSchemaTableResult struct {

	// **参数解释**: 表名称。 **取值范围**: 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`
}

func (o DatabaseSchemaTableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseSchemaTableResult struct{}"
	}

	return strings.Join([]string{"DatabaseSchemaTableResult", string(data)}, " ")
}
