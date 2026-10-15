package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkResourceConfigResponseBody **参数解释**：Flink作业资源配置，包含作业的并行度和各组件资源规格。
type ShowFlinkResourceConfigResponseBody struct {

	// **参数解释**：作业并行度。 **取值范围**：1~1000。
	Parallelism int32 `json:"parallelism"`

	JobmanagerResourceSpec *ShowFlinkResourceSpecResponseBody `json:"jobmanager_resource_spec"`

	TaskmanagerResourceSpec *ShowFlinkResourceSpecResponseBody `json:"taskmanager_resource_spec"`

	SqlGatewayResourceSpec *ShowFlinkSqlGatewayResourceSpecResponseBody `json:"sql_gateway_resource_spec,omitempty"`
}

func (o ShowFlinkResourceConfigResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkResourceConfigResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkResourceConfigResponseBody", string(data)}, " ")
}
