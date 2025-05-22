package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterRingInfo **参数解释**： 集群实例环信息。 **取值范围**： 不涉及。
type LogicalClusterRingInfo struct {

	// **参数解释**： 集群实例环信息。 **取值范围**： 不涉及。
	RingHosts *[]RingHost `json:"ring_hosts,omitempty"`
}

func (o LogicalClusterRingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterRingInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterRingInfo", string(data)}, " ")
}
