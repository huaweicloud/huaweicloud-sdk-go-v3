package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterInfo **参数解释**： 创建逻辑集群信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CreateLogicalClusterInfo struct {

	// **参数解释**： 逻辑集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName string `json:"logical_cluster_name"`

	// **参数解释**： 逻辑集群环信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterRings []ClusterRing `json:"cluster_rings"`
}

func (o CreateLogicalClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterInfo", string(data)}, " ")
}
