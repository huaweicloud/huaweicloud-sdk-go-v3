package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkRestoreStrategyResponse **参数解释**：运行及恢复策略配置，包含重试次数、超时时间等信息。
type ShowSparkRestoreStrategyResponse struct {

	// **参数解释**：最大重试次数，用于控制作业失败后的自动重试。 **取值范围**：大于0表示开启失败重试，0表示不开启重试。
	MaxRetry *int32 `json:"max_retry,omitempty"`

	// **参数解释**：重试间隔时间，用于指定每次重试之间的等待时间。 **取值范围**：单位为秒，默认值为30秒。
	RetryDelay *int32 `json:"retry_delay,omitempty"`

	// **参数解释**：排队超时时间，用于指定作业提交后的最大等待时间。 **取值范围**：单位为分钟，默认值为180分钟，最小值为10分钟，最大值180分钟。超过此时间作业仍未运行则作业失败。
	QueuedTimeout *int32 `json:"queued_timeout,omitempty"`

	// **参数解释**：运行超时时间，用于指定作业运行的最大时长。 **取值范围**：单位为分钟，默认值为-1（表示不限制），最小值为10分钟，最大值为525600分钟（10年）。超过此时间作业未结束则作业会取消运行并标记失败。
	RunningTimeout *int32 `json:"running_timeout,omitempty"`
}

func (o ShowSparkRestoreStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkRestoreStrategyResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkRestoreStrategyResponse", string(data)}, " ")
}
