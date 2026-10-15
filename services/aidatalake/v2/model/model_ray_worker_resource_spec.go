package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayWorkerResourceSpec **参数解释**：Ray Worker资源配置。 **约束限制**：不涉及。
type RayWorkerResourceSpec struct {

	// **参数解释**：CPU核数量，单位为毫核。 **约束限制**：不涉及。 **取值范围**：最小值为1000。 **默认取值**：不涉及。
	Cpu int32 `json:"cpu"`

	// **参数解释**：内存大小，单位MB。 **约束限制**：不涉及。 **取值范围**：最小值为1024。 **默认取值**：不涉及。
	Memory int32 `json:"memory"`

	// **参数解释**：节点的GPU数量。 **约束限制**：不支持小于1的值或带小数部分的值（如0.5, 1.5）。 **取值范围**：需配置为大于等于1.0，且必须为整数值（如1.0, 2.0, 3.0, ...）。 **默认取值**：不涉及。
	Gpu *float64 `json:"gpu,omitempty"`

	// **参数解释**：节点的NPU数量。 **约束限制**：不支持小于1的值或带小数部分的值（如0.5, 1.5）。 **取值范围**：需配置为大于等于1.0，且必须为整数值（如1.0, 2.0, 3.0, ...）。 **默认取值**：不涉及。
	Npu *float64 `json:"npu,omitempty"`

	// **参数解释**：Worker副本数量。 **约束限制**：不涉及。 **取值范围**：1~64。 **默认取值**：不涉及。
	WorkerReplicas *int32 `json:"worker_replicas,omitempty"`

	// **参数解释**：是否开启worker节点的弹性伸缩。 **约束限制**：不涉及。 **取值范围**：可选值有：   - true：开启弹性伸缩。取该值时worker_min_replicas和worker_max_replicas必须设置。   - false：不开启弹性伸缩。取该值时worker_replicas必须设置。 **默认取值**：false。
	EnableAutoscaling bool `json:"enable_autoscaling"`

	// **参数解释**：Worker最小副本数量。 **约束限制**：不涉及。 **取值范围**：0~64。 **默认取值**：不涉及。
	WorkerMinReplicas *int32 `json:"worker_min_replicas,omitempty"`

	// **参数解释**：Worker最大副本数量。 **约束限制**：不涉及。 **取值范围**：1~64。 **默认取值**：不涉及。
	WorkerMaxReplicas *int32 `json:"worker_max_replicas,omitempty"`
}

func (o RayWorkerResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayWorkerResourceSpec struct{}"
	}

	return strings.Join([]string{"RayWorkerResourceSpec", string(data)}, " ")
}
