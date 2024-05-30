package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Components 组件列表
type Components struct {

	// 组件id，当组件类型为DN，组件id为6001，则对应的值为dn_6001。
	Id *string `json:"id,omitempty"`

	// 节点类型，取值为“master”、“slave”，分别对应于主节点、备节点。
	Role *string `json:"role,omitempty"`

	// 组件状态。 Primary：该组件为主。 Normal：该组件状态正常。 Down：该组件处于宕机状态。 Standby：该组件为备。 StateFollower：该ETCD为备。 StateLeader：该ETCD为主。 StateCandidate：该ETCD为仲裁。
	Status *string `json:"status,omitempty"`

	// 分组id，只有dn组件有分组id，用于区分是否是同一个分片下的组件。其他组件为空字符串。
	DistributedId *string `json:"distributed_id,omitempty"`

	// 节点类型，包括：DN, CN, GTM, CM, ETCD。
	Type *string `json:"type,omitempty"`

	// 详情。
	Detail *string `json:"detail,omitempty"`
}

func (o Components) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Components struct{}"
	}

	return strings.Join([]string{"Components", string(data)}, " ")
}
