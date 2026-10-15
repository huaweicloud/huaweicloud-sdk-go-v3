package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobSqlScriptParameter **参数解释**：SQL脚本占位符参数，用于配置SQL脚本中的参数化变量。
type SparkJobSqlScriptParameter struct {

	// **参数解释**：占位符的键，用于标识SQL脚本中的参数名称。 **取值范围**：长度为1~128个字符。
	Key string `json:"key"`

	// **参数解释**：占位符的值，用于传递给SQL脚本的实际参数值。不同类型的数据格式如下： - STRING：字符串，例如：“xxx”。 - DECIMAL：定点数，例如：“12.1”。 - INTEGER：整数，例如：“13”。 - DATE：日期时间戳，例如：“1779188276372”。 - TIMESTAMP：时间戳，例如：“1779188276372”。 **取值范围**：长度为1~512个字符。
	Value string `json:"value"`

	ValueType *SparkSqlParameterValueType `json:"value_type,omitempty"`
}

func (o SparkJobSqlScriptParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobSqlScriptParameter struct{}"
	}

	return strings.Join([]string{"SparkJobSqlScriptParameter", string(data)}, " ")
}
