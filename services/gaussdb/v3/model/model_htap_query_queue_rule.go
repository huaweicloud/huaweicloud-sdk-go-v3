package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapQueryQueueRule **参数解释**：  HTAP标准版查询队列规则。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认值**：  不涉及。
type HtapQueryQueueRule struct {

	// **参数解释**：  查询队列中的查询数量上限。  **约束限制**：  不涉及。  **取值范围**：  非负整数。  **默认值**：  1024。
	QueryQueueMaxQueuedQueries int32 `json:"query_queue_max_queued_queries"`

	// **参数解释**：  查询阻塞时间上限。  **约束限制**：  不涉及。  **取值范围**：  非负整数。  **默认值**：  300。
	QueryQueuePendingTimeoutSecond int32 `json:"query_queue_pending_timeout_second"`

	// **参数解释**：  查询队列并发值。  **约束限制**：  不涉及。  **取值范围**：  0~100的整数，0表示没有限制。  **默认值**：  0。
	QueryQueueConcurrencyLimit int32 `json:"query_queue_concurrency_limit"`

	// **参数解释**：  内存使用百分比。  **约束限制**：  不涉及。  **取值范围**：  0~100的整数，0表示没有限制。  **默认值**：  0。
	QueryQueueMemUsedPctLimit int32 `json:"query_queue_mem_used_pct_limit"`

	// **参数解释**：  CPU使用百分比。  **约束限制**：  不涉及。  **取值范围**：  0~100的整数，0表示没有限制。  **默认值**：  0。
	QueryQueueCpuUsedPctLimit int32 `json:"query_queue_cpu_used_pct_limit"`

	// **参数解释**：  查询队列开关状态。  **约束限制**：  不涉及。  **取值范围**：  - true：表示开启。 - false：表示关闭。  **默认值**：  false。
	EnableQueryQueueSelect *string `json:"enable_query_queue_select,omitempty"`
}

func (o HtapQueryQueueRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapQueryQueueRule struct{}"
	}

	return strings.Join([]string{"HtapQueryQueueRule", string(data)}, " ")
}
