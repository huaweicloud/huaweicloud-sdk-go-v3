package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesNodesResult struct {

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 组件ID列表。
	ComponentIds *[]string `json:"component_ids,omitempty"`
}

func (o InstancesNodesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesNodesResult struct{}"
	}

	return strings.Join([]string{"InstancesNodesResult", string(data)}, " ")
}
