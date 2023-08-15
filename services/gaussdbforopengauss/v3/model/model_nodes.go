package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nodes 节点信息
type Nodes struct {

	// 节点ID。
	Id string `json:"id"`

	// 节点名字。
	Name *string `json:"name,omitempty"`

	// 节点所在可用区编码。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 可用区描述信息。
	Description *string `json:"description,omitempty"`

	// 节点状态。
	Status *string `json:"status,omitempty"`

	// 组件列表。
	Components []Components `json:"components"`
}

func (o Nodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nodes struct{}"
	}

	return strings.Join([]string{"Nodes", string(data)}, " ")
}
