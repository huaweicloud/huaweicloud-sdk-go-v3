package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSqlGatewayResourceSpec **参数解释**：SQL Gateway的资源规格，仅FLINK_SQL_GATEWAY类型使用。 **约束限制**：SQL Gateway与JobManager的CPU之和不能超过48000毫核，内存之和不能超过393216 MB。
type FlinkSqlGatewayResourceSpec struct {

	// **参数解释**：CPU毫核数，单位为毫核，用于指定SQL Gateway的CPU资源。 **约束限制**：SQL Gateway与JobManager的CPU之和不能超过48000毫核。 **取值范围**：500~47500。 **默认取值**：1000。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存大小，单位为MB，用于指定SQL Gateway的内存资源。 **约束限制**：SQL Gateway与JobManager的内存之和不能超过393216 MB。 **取值范围**：2048~391168。 **默认取值**：4096。
	Memory *int32 `json:"memory,omitempty"`
}

func (o FlinkSqlGatewayResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlGatewayResourceSpec struct{}"
	}

	return strings.Join([]string{"FlinkSqlGatewayResourceSpec", string(data)}, " ")
}
