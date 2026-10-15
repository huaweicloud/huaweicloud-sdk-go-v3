package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSqlGatewayResourceSpec **参数解释**：SQL Gateway的资源规格，仅FLINK_SQL_GATEWAY类型使用。 **约束限制**：不涉及。
type FlinkSqlGatewayResourceSpec struct {

	// **参数解释**：SQL Gateway的CPU配额，单位为毫核。 **约束限制**：不涉及。 **取值范围**：100~48000。 **默认取值**：500。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：SQL Gateway的内存配额，单位为MB。 **约束限制**：不涉及。 **取值范围**：128~393216。 **默认取值**：1000。
	Memory *int32 `json:"memory,omitempty"`
}

func (o FlinkSqlGatewayResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlGatewayResourceSpec struct{}"
	}

	return strings.Join([]string{"FlinkSqlGatewayResourceSpec", string(data)}, " ")
}
