package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableInfoOpen struct {

	// **参数解释**： 数据库名称。 **默认取值**： 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**： 模式名。 **默认取值**： 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**： 表名。 **默认取值**： 不涉及。
	TableName *string `json:"table_name,omitempty"`
}

func (o TableInfoOpen) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableInfoOpen struct{}"
	}

	return strings.Join([]string{"TableInfoOpen", string(data)}, " ")
}
