package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkExecutorResourceSpec **参数解释**：SparkSql端点的Executor资源配置。
type SparkExecutorResourceSpec struct {

	// **参数解释**：Spark Executor CPU，单位毫核。 **约束限制**：不涉及。 **取值范围**：1000~48000。 **默认取值**：不涉及。
	Cpu int32 `json:"cpu"`

	// **参数解释**：Spark Executor内存大小，单位MB，1GB=1024MB。 **约束限制**：不涉及。 **取值范围**：1024~393216。 **默认取值**：不涉及。
	Memory int32 `json:"memory"`

	// **参数解释**：Spark Executor最小数量。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	MinExecutorNumbers int32 `json:"min_executor_numbers"`

	// **参数解释**：Spark Executor最大数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：不涉及。
	MaxExecutorNumbers int32 `json:"max_executor_numbers"`
}

func (o SparkExecutorResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkExecutorResourceSpec struct{}"
	}

	return strings.Join([]string{"SparkExecutorResourceSpec", string(data)}, " ")
}
