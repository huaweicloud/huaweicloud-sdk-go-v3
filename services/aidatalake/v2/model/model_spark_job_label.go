package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobLabel **参数解释**：作业标签，用于标识和分类Spark作业。 **约束限制**：不涉及。
type SparkJobLabel struct {

	// **参数解释**：标签键，用于标识标签的唯一键名。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签值，用于标识标签的具体值。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o SparkJobLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobLabel struct{}"
	}

	return strings.Join([]string{"SparkJobLabel", string(data)}, " ")
}
