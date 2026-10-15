package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SparkSqlParameter struct {

	// **参数解释**：占位符的键，用于在SQL语句中标识参数位置。例如：SQL语句中的${key}。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：占位符的值，用于替换SQL语句中的占位符。根据value_type的不同，值的格式也不同。 - STRING：字符串，例如：“xxx”。 - DECIMAL：定点数，例如：“12.1”。 - INTEGER：整数，例如：“13”。 - DATE：日期时间戳，例如：“1779188276372”。 - TIMESTAMP：时间戳，例如：“1779188276372”。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	Value string `json:"value"`

	ValueType *SparkSqlParameterValueType `json:"value_type,omitempty"`
}

func (o SparkSqlParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlParameter struct{}"
	}

	return strings.Join([]string{"SparkSqlParameter", string(data)}, " ")
}
