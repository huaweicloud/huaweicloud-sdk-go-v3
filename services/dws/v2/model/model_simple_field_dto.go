package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleFieldDto **参数解释**： 指标表相关字段信息。 **取值范围**： 不涉及。
type SimpleFieldDto struct {

	// **参数解释**： 指标表对应字段名称。 **取值范围**： 不涉及。
	ColumnName *string `json:"column_name,omitempty"`

	// **参数解释**： 指标表对应字段类型。 **取值范围**： 不涉及。
	ColumnType *string `json:"column_type,omitempty"`
}

func (o SimpleFieldDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleFieldDto struct{}"
	}

	return strings.Join([]string{"SimpleFieldDto", string(data)}, " ")
}
