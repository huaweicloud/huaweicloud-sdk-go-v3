package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopologyNodes
type TopologyNodes struct {

	// **取值范围**：拓扑LB节点信息。
	Loadbalancers []LoadBalancerNode `json:"loadbalancers"`

	// **取值范围**：拓扑EIP节点信息。
	Eips []EipNode `json:"eips"`

	// **取值范围**：拓扑监听器节点信息。
	Listeners []ListenerNode `json:"listeners"`

	// **取值范围**：拓扑后端服务器组节点信息。
	Pools []PoolNode `json:"pools"`
}

func (o TopologyNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyNodes struct{}"
	}

	return strings.Join([]string{"TopologyNodes", string(data)}, " ")
}
