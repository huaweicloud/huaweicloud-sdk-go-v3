package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolNodeAutoscaling 节点池自动伸缩相关配置
type NodePoolNodeAutoscaling struct {

	// 是否开启自动扩缩容
	Enable *bool `json:"enable,omitempty"`

	// 若开启自动扩缩容，最小能缩容的节点个数。不可大于集群规格所允许的节点上限
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// 若开启自动扩缩容，最大能扩容的节点个数，应大于等于 minNodeCount，且不超过集群规格对应的节点数量上限。
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	// 节点保留时间，单位为分钟，扩容出来的节点在这个时间内不会被缩掉
	ScaleDownCooldownTime *int32 `json:"scaleDownCooldownTime,omitempty"`

	// 节点池权重，更高的权重在扩容时拥有更高的优先级
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释**： 缩容非必要时间。单位为分钟，该参数用于指定在确定可以进行缩容操作之前，节点处于不需要状态的持续时间。 当节点在指定的这段时间内一直处于不需要的状态时，autoscaler 插件才会考虑对其进行缩容操作。 这样可以避免因资源的短暂波动而频繁触发缩容，增强系统的稳定性。如果未设置该参数，autoscaler 插件插件会使用默认的时间阈值。 **约束限制**： 不涉及 **取值范围**： 0-2147483647。 > 注意：如果传值为-1，代表取值为空。  **默认取值**： 默认为空
	ScaleDownUnneededTime *int32 `json:"scaleDownUnneededTime,omitempty"`

	// **参数解释**： 缩容利用率阈值。运行在该节点上的所有 Pod 的 CPU 或内存总和除以该节点相应的可分配资源， 当该比值低于此阈值时，该节点可被考虑进行缩容。例如，设置为 0.3 表示当资源利用率低于 30% 时， 会触发缩容操作的评估。如果未设置该参数，autoscaler 插件会使用默认的利用率阈值。 **约束限制**： 不涉及 **取值范围**： 0-1。 > 注意：如果传值为-1，代表取值为空。  **默认取值**： 默认为空
	ScaleDownUtilizationThreshold *float32 `json:"scaleDownUtilizationThreshold,omitempty"`
}

func (o NodePoolNodeAutoscaling) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolNodeAutoscaling struct{}"
	}

	return strings.Join([]string{"NodePoolNodeAutoscaling", string(data)}, " ")
}
