package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListReadonlyNodesResult struct {

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 节点状态。 **取值范围**:  - normal：节点正常。  - abnormal：节点异常。  - creating：节点正在创建。  - createfail：节点创建失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 节点所在的可用区。 **取值范围**: 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListReadonlyNodesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadonlyNodesResult struct{}"
	}

	return strings.Join([]string{"ListReadonlyNodesResult", string(data)}, " ")
}
