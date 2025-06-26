package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableDetail **参数解释**： 恢复表信息。 **取值范围**： 不涉及。
type TableDetail struct {

	// **参数解释**： Schema名称。 **取值范围**： 不涉及。
	SchemaName string `json:"schema_name"`

	// **参数解释**： 表名称。 **取值范围**： 不涉及。
	TableName string `json:"table_name"`
}

func (o TableDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableDetail struct{}"
	}

	return strings.Join([]string{"TableDetail", string(data)}, " ")
}
