package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchBindPoolNodesReq struct {

	// **参数解释**：需要进行换绑的节点列表。 **约束限制**：不涉及。
	Nodes []BindNodeItem `json:"nodes"`

	// **参数解释**：是否对换绑的节点进行排水。 **约束限制**：不涉及。 **取值范围**： - true：排水 - false：不排水 **默认取值**：不涉及。
	Drain *bool `json:"drain,omitempty"`
}

func (o BatchBindPoolNodesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindPoolNodesReq struct{}"
	}

	return strings.Join([]string{"BatchBindPoolNodesReq", string(data)}, " ")
}
