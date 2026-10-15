package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SparkSqlLabel struct {

	// **参数解释**：标签的键，用于标识作业的分类维度。例如：env、project等。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签的值，用于标识作业的具体分类。例如：production、data-analytics等。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o SparkSqlLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlLabel struct{}"
	}

	return strings.Join([]string{"SparkSqlLabel", string(data)}, " ")
}
