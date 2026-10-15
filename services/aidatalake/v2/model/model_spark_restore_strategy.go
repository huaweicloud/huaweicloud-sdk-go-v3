package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkRestoreStrategy **参数解释**：作业恢复策略配置，用于定义作业失败后的重试机制和超时控制。 **约束限制**：不涉及。
type SparkRestoreStrategy struct {

	// **参数解释**：最大重试次数，用于控制作业失败后的自动重试次数。如果配置的值大于0，则自动开启失败重试；如果不配置或者配置为0，则不开启作业失败重试。 **约束限制**：不涉及。 **取值范围**：0~65535。 **默认取值**：0。
	MaxRetry *int32 `json:"max_retry,omitempty"`

	// **参数解释**：重试间隔时间，用于指定作业失败重试机制中每次重试的时间间隔，单位为秒。 **约束限制**：不涉及。 **取值范围**：1~3600。 **默认取值**：30。
	RetryDelay *int32 `json:"retry_delay,omitempty"`

	// **参数解释**：排队超时时间，用于指定作业提交后等待运行的最长时间，单位为分钟。如果超过此时间作业仍未运行，则作业失败。 **约束限制**：不涉及。 **取值范围**：10~180。 **默认取值**：180。
	QueuedTimeout *int32 `json:"queued_timeout,omitempty"`

	// **参数解释**：运行超时时间，用于指定作业运行的最大时长，单位为分钟。如果超过此时间作业还未运行结束，则作业会取消运行并标记为运行超时。 **约束限制**：不涉及。 **取值范围**：10~525600。 **默认取值**：不涉及。
	RunningTimeout *int32 `json:"running_timeout,omitempty"`
}

func (o SparkRestoreStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkRestoreStrategy struct{}"
	}

	return strings.Join([]string{"SparkRestoreStrategy", string(data)}, " ")
}
