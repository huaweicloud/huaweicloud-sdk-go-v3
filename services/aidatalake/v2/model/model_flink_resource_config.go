package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkResourceConfig **参数解释**：Flink作业的资源配置，配置并行度和各组件资源规格。 **约束限制**：不涉及。
type FlinkResourceConfig struct {

	// **参数解释**：作业并行度，决定作业的并发处理能力。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：1。
	Parallelism *int32 `json:"parallelism,omitempty"`

	JobmanagerResourceSpec *FlinkResourceSpec `json:"jobmanager_resource_spec,omitempty"`

	TaskmanagerResourceSpec *FlinkResourceSpec `json:"taskmanager_resource_spec,omitempty"`

	SqlGatewayResourceSpec *FlinkSqlGatewayResourceSpec `json:"sql_gateway_resource_spec,omitempty"`
}

func (o FlinkResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkResourceConfig struct{}"
	}

	return strings.Join([]string{"FlinkResourceConfig", string(data)}, " ")
}
