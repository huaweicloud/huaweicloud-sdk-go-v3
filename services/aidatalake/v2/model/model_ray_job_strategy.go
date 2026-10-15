package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayJobStrategy **参数解释**：Ray Job提交策略。 **约束限制**：不涉及。
type RayJobStrategy struct {

	// 作业提交超时时间，如果超过此时间作业仍未运行，则作业失败。默认作业提交超时时间为3小时，最大值为1天，最小值为10分钟。单位：分钟。取值范围为10~1440，单位为分钟。
	QueuedTimeout *int32 `json:"queued_timeout,omitempty"`

	// 作业运行超时时间，如果超过此时间作业还未运行结束，则作业会取消运行，并标记失败。默认值为-1，表示没有配置作业运行超时时间。最小时间为10分钟，最大时间为7天。单位：分钟。取值范围为-1~10080（-1表示不配置超时），单位为分钟。
	RunningTimeout *int32 `json:"running_timeout,omitempty"`

	// **参数解释**：最大重试次数，用于控制作业失败后的自动重试次数。如果配置的值大于0，则自动开启失败重试；如果不配置或者配置为0，则不开启作业失败重试。 **约束限制**：不涉及。 **取值范围**：0~65535。 **默认取值**：0。
	MaxRetry *int32 `json:"max_retry,omitempty"`

	// **参数解释**：重试间隔时间，用于指定作业失败重试机制中每次重试的时间间隔，单位为秒。 **约束限制**：不涉及。 **取值范围**：1~3600。 **默认取值**：30。
	RetryDelay *int32 `json:"retry_delay,omitempty"`
}

func (o RayJobStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayJobStrategy struct{}"
	}

	return strings.Join([]string{"RayJobStrategy", string(data)}, " ")
}
