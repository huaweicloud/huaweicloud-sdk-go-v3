package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkResourceConfig **参数解释**：Spark作业资源配置，用于配置Driver和Executor的资源规格和数量。 **约束限制**：不涉及。
type SparkResourceConfig struct {
	DriverResourceSpec *ResourceSpec `json:"driver_resource_spec,omitempty"`

	// **参数解释**：Executor数量，用于指定Spark作业的最大Executor个数。如果配置此参数，则表示启用动态资源分配，动态资源分配最小Executor数为1，初始Executor数为1，最大Executor数为配置值。 **约束限制**：不涉及。 **取值范围**：0~65535。 **默认取值**：不涉及。
	ExecutorNumber *int32 `json:"executor_number,omitempty"`

	ExecutorResourceSpec *ResourceSpec `json:"executor_resource_spec,omitempty"`
}

func (o SparkResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkResourceConfig struct{}"
	}

	return strings.Join([]string{"SparkResourceConfig", string(data)}, " ")
}
