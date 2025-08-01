package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterRing **参数解释**： 集群主机信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterRing struct {

	// **参数解释**： 集群主机信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RingHosts []RingHost `json:"ring_hosts"`

	// **参数解释**： 是否可以缩容。 **约束限制**： 不涉及。 **取值范围**： false|true。 **默认取值**： 不涉及。
	UnShrinkableClusterRing *bool `json:"un_shrinkable_cluster_ring,omitempty"`
}

func (o ClusterRing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRing struct{}"
	}

	return strings.Join([]string{"ClusterRing", string(data)}, " ")
}
