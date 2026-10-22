package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSparkClusterConfig struct {

	// **参数解释**：集群老化时间（小时）。 **约束限制**：不涉及。 **取值范围**：1~12。 **默认取值**：不涉及。
	AgingTime *int32 `json:"aging_time,omitempty"`

	AutoStopConfig *AutoStopConfig `json:"auto_stop_config,omitempty"`

	DriverResourceSpec *UpdateSparkDriverResourceSpec `json:"driver_resource_spec,omitempty"`

	ExecutorResourceSpec *UpdateSparkExecutorResourceSpec `json:"executor_resource_spec,omitempty"`

	// **参数解释**：最小集群数量。 **约束限制**：不涉及。 **取值范围**：0~20。 **默认取值**：不涉及。
	MinClusters *int32 `json:"min_clusters,omitempty"`

	// **参数解释**：最大集群数量。 **约束限制**：不涉及。 **取值范围**：2~20。 **默认取值**：不涉及。
	MaxClusters *int32 `json:"max_clusters,omitempty"`

	// **参数解释**：集群最大并发SQL数量。 **约束限制**：不涉及。 **取值范围**：1~30。 **默认取值**：不涉及。
	MaxClusterConcurrentQueries *int32 `json:"max_cluster_concurrent_queries,omitempty"`
}

func (o UpdateSparkClusterConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparkClusterConfig struct{}"
	}

	return strings.Join([]string{"UpdateSparkClusterConfig", string(data)}, " ")
}
