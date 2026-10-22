package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkDriverResourceSpec **参数解释**：SparkSql端点的Driver资源配置。
type SparkDriverResourceSpec struct {

	// **参数解释**：Spark Driver CPU，单位毫核。 **约束限制**：不涉及。 **取值范围**：1000~48000。 **默认取值**：不涉及。
	Cpu int32 `json:"cpu"`

	// **参数解释**：Spark Driver内存大小，单位MB，1GB=1024MB。 **约束限制**：不涉及。 **取值范围**：1024~393216。 **默认取值**：不涉及。
	Memory int32 `json:"memory"`
}

func (o SparkDriverResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkDriverResourceSpec struct{}"
	}

	return strings.Join([]string{"SparkDriverResourceSpec", string(data)}, " ")
}
