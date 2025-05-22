package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopoRingInfo **参数解释**： 集群拓扑实例环信息。 **取值范围**： 不涉及。
type TopoRingInfo struct {

	// **参数解释**： 集群实例列表信息。 **取值范围**： 不涉及。
	InstanceInfoLists *[]TopoInstanceInfo `json:"instance_info_lists,omitempty"`
}

func (o TopoRingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoRingInfo struct{}"
	}

	return strings.Join([]string{"TopoRingInfo", string(data)}, " ")
}
