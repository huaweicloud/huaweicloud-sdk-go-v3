package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadataUpdateAnnotations **参数解释**：资源池的注释信息。
type PoolMetadataUpdateAnnotations struct {

	// **参数解释**：资源池描述信息，用于说明资源池用于某种指定场景。不能包含特殊字符!<>=&\"'。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsDescription *string `json:"os.modelarts/description,omitempty"`

	// **参数解释**：订单id，包周期创建和变更的时候需要传递该参数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsOrderId *string `json:"os.modelarts/order.id,omitempty"`

	// **参数解释**：资源池的调度算法。 **约束限制**：不涉及。 **取值范围**：可选值如下： - FIFO：先进先出，默认值 - Range：遍历调度 - Step: 阶梯轮询 - Rank: 分级调度 **默认取值**：FIFO。
	OsModelartsPoolSchedulerQueueStrategy *string `json:"os.modelarts.pool/scheduler.queue.strategy,omitempty"`

	// **参数解释**：资源池的调度策略配置。 **约束限制**：仅当调度策略为Step或Rank时生效。 **取值范围**：Json字符串，自行配置，示例如下： - Step: {\"step_config\":{\"timeout\":300}} - Rank: {\"rank_config\":[{\"max_priority\":3,\"min_priority\":3,\"rank_strategy\":\"Step\",\"step_config\":{\"timeout\":300}},{\"max_priority\":2,\"min_priority\":1,\"rank_strategy\":\"Range\"}]} **默认取值**：不涉及。
	OsModelartsPoolSchedulerQueueStrategyConfig *string `json:"os.modelarts.pool/scheduler.queue.strategy.config,omitempty"`
}

func (o PoolMetadataUpdateAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadataUpdateAnnotations struct{}"
	}

	return strings.Join([]string{"PoolMetadataUpdateAnnotations", string(data)}, " ")
}
