package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CnComponentInfoResult struct {

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 组件ID。 **取值范围**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`
}

func (o CnComponentInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CnComponentInfoResult struct{}"
	}

	return strings.Join([]string{"CnComponentInfoResult", string(data)}, " ")
}
