package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterRequestBody **参数解释**： 编辑逻辑集群请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type UpdateLogicalClusterRequestBody struct {

	// **参数解释**： 逻辑集群编辑环列表信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterRings []ClusterRing `json:"cluster_rings"`

	// **参数解释**： 重分布模式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 查杀作业等待时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WaitingForKilling *int32 `json:"waiting_for_killing,omitempty"`
}

func (o UpdateLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterRequestBody", string(data)}, " ")
}
