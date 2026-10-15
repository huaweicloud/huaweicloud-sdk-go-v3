package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkResourceConfigResponse **参数解释**：资源配置信息，包含Driver和Executor的资源配置。
type ShowSparkResourceConfigResponse struct {
	DriverResourceSpec *ShowResourceSpecResponse `json:"driver_resource_spec,omitempty"`

	ExecutorResourceSpec *ShowResourceSpecResponse `json:"executor_resource_spec,omitempty"`

	// **参数解释**：Spark作业的最大executor个数，用于动态资源分配。 **取值范围**：大于0的整数。 当配置“executor_number”时，表示启用动态资源分配：spark.dynamicAllocation.enabled=true,spark.dynamicAllocation.minExecutor=1,spark.dynamicAllocation.maxExecutors=executor_number,spark.dynamicAllocation.initialExecutors=1。如需修改上述参数，可通过properties字段设置spark相关参数。
	ExecutorNumber *int32 `json:"executor_number,omitempty"`
}

func (o ShowSparkResourceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkResourceConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkResourceConfigResponse", string(data)}, " ")
}
