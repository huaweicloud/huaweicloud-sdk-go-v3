package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeResult 实例节点信息。
type NodeResult struct {

	// 节点ID。
	Id string `json:"id"`

	// 节点名称。
	Name string `json:"name"`

	// 节点类型，取值为“master”、“slave”，分别对应于主节点、备节点。
	Role string `json:"role"`

	// 节点状态。
	Status string `json:"status"`

	// 可用区。
	AvailabilityZone string `json:"availability_zone"`

	// 节点内网IP。分布式实例类型下该参数仅针对CN节点有效，主备版实例类型下该参数对所有节点有效，且在弹性云服务器创建成功后参数值存在。
	PrivateIp string `json:"private_ip"`

	// 绑定的外网IP。分布式实例类型下该参数仅针对CN节点有效，主备版实例类型下该参数对所有节点有效，且在弹性云服务器创建成功并绑定弹性公网IP后参数值存在。
	PublicIp string `json:"public_ip"`

	// 节点上组件信息（例组件ID:分布式ID），多个组件信息用;隔开。
	ComponentNames *string `json:"component_names,omitempty"`
}

func (o NodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResult struct{}"
	}

	return strings.Join([]string{"NodeResult", string(data)}, " ")
}
