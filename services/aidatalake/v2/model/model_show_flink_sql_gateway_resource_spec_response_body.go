package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkSqlGatewayResourceSpecResponseBody **参数解释**：SQL Gateway sidecar的资源规格，包含CPU和内存配额，仅FLINK_SQL_GATEWAY类型返回。
type ShowFlinkSqlGatewayResourceSpecResponseBody struct {

	// **参数解释**：SQL Gateway的CPU配额，单位为毫核。 **取值范围**：100~48000。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：SQL Gateway的内存配额，单位为MB。 **取值范围**：128~393216。
	Memory *int32 `json:"memory,omitempty"`
}

func (o ShowFlinkSqlGatewayResourceSpecResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSqlGatewayResourceSpecResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkSqlGatewayResourceSpecResponseBody", string(data)}, " ")
}
