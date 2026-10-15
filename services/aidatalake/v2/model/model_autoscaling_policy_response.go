package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingPolicyResponse **参数解释**：自动调优配置，用于控制Flink作业自动调优行为。 **约束限制**：不涉及。
type AutoscalingPolicyResponse struct {

	// **参数解释**：自动调优开关。 **约束限制**：不涉及。 **取值范围**：   - true：开启自动调优。   - false：关闭自动调优。 **默认取值**：false。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**：算子最大并行度。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：1000。
	VertexMaxParallelism *int32 `json:"vertex_max_parallelism,omitempty"`

	// **参数解释**：CPU配额（单位：毫核）。 **约束限制**：不涉及。 **取值范围**：500~40000000。 **默认取值**：16000。
	QuotaCpu *int32 `json:"quota_cpu,omitempty"`

	// **参数解释**：内存配额（单位：MB）。 **约束限制**：不涉及。 **取值范围**：2048~81920000。 **默认取值**：65536。
	QuotaMemory *int64 `json:"quota_memory,omitempty"`

	// **参数解释**：用户自定义的自动调优配置参数，为键值对形式，传入Flink Kubernetes Operator支持的job.autoscaler.*参数。最多支持50个参数，每个值长度为1~1024个字符。 **约束限制**：采用全量替换模式。以传入内容覆盖旧内容；不传该字段则保留现有配置不变。参数名必须以job.autoscaler.开头。结构化字段同名键（job.autoscaler.enabled、job.autoscaler.vertex.max-parallelism、job.autoscaler.quota.cpu、job.autoscaler.quota.memory）配置不生效。
	CustomConfig map[string]string `json:"custom_config,omitempty"`

	// **参数解释**：自动调优配置的同步状态。配置下发失败时置为FAILED，可通过重新调用更新接口重新下发。 **约束限制**：不涉及。 **取值范围**：   - PENDING：配置已保存，尚未确认下发。   - SYNCED：配置已确认下发成功。   - FAILED：配置下发失败，需重新下发。 **默认取值**：不涉及。
	SyncStatus *string `json:"sync_status,omitempty"`
}

func (o AutoscalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"AutoscalingPolicyResponse", string(data)}, " ")
}
