package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkLogicalClusterRequestBody **参数解释**： 逻辑集群缩容请求体。  **约束限制**：  必须是非空值。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
type ShrinkLogicalClusterRequestBody struct {

	// **参数解释**： 缩容主机环信息。  **约束限制**：  不涉及。 **取值范围**：  不涉及。 **默认取值**：  不涉及。
	ClusterRings []ClusterRingVo `json:"cluster_rings"`

	// **参数解释**： 重分布并发配置数。  **约束限制**：  不涉及。 **取值范围**：  1~200。 **默认取值**：  4。
	ParallelJobs *int32 `json:"parallel_jobs,omitempty"`

	// **参数解释**：  缩容模式。 **约束限制**：  不涉及。 **取值范围**：  read-only：离线模式 insert：在线模式 **默认取值**：  insert
	Mode *string `json:"mode,omitempty"`
}

func (o ShrinkLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"ShrinkLogicalClusterRequestBody", string(data)}, " ")
}
