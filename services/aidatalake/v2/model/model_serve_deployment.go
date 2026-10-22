package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeDeployment **参数解释**：Serve 应用 Deployment 配置。
type ServeDeployment struct {

	// **参数解释**：Deployment 名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：副本数。 **约束限制**：不涉及。 **取值范围**：不小于1。 **默认取值**：不涉及。
	NumReplicas int32 `json:"num_replicas"`

	// **参数解释**：单节点最大副本数。 **约束限制**：不涉及。 **取值范围**：不小于1。 **默认取值**：不涉及。
	MaxReplicasPerNode *int32 `json:"max_replicas_per_node,omitempty"`

	// **参数解释**：最大进行中请求数。 **约束限制**：不涉及。 **取值范围**：不小于1。 **默认取值**：不涉及。
	MaxOngoingRequests *int32 `json:"max_ongoing_requests,omitempty"`

	// **参数解释**：最大排队请求数。 **约束限制**：不涉及。 **取值范围**：不小于1。 **默认取值**：不涉及。
	MaxQueuedRequests *int32 `json:"max_queued_requests,omitempty"`

	RayActorOptions *ServeRayActorOptions `json:"ray_actor_options,omitempty"`
}

func (o ServeDeployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeDeployment struct{}"
	}

	return strings.Join([]string{"ServeDeployment", string(data)}, " ")
}
