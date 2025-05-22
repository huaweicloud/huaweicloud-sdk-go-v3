package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterRingVo **参数解释**：  集群主机环信息。 **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
type ClusterRingVo struct {

	// **参数解释**：  集群主机环信息。 **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	RingHosts *[]RingHostVo `json:"ring_hosts,omitempty"`
}

func (o ClusterRingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRingVo struct{}"
	}

	return strings.Join([]string{"ClusterRingVo", string(data)}, " ")
}
