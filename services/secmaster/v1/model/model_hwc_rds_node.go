package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRdsNode 主备实例信息
type HwcRdsNode struct {

	// 节点ID。
	Id string `json:"id"`

	// 节点名称。
	Name string `json:"name"`

	// 节点类型，取值为“master”、“slave”或“readreplica”，分别对应于主节点、备节点和只读节点。
	Role string `json:"role"`

	// 节点状态。
	Status string `json:"status"`

	// 可用区。
	AvailabilityZone string `json:"availability_zone"`
}

func (o HwcRdsNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRdsNode struct{}"
	}

	return strings.Join([]string{"HwcRdsNode", string(data)}, " ")
}
