package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapNodeInfoResponseBodyNodeList struct {

	// **参数解释**： HTAP标准版实例节点ID。  **取值范围**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： HTAP标准版实例节点名称。  **取值范围**：  不涉及。
	NodeName string `json:"node_name"`

	// **参数解释**： HTAP标准版实例节点角色。  **取值范围**： - fe-leader - fe-follower - fe-observer - be
	Role string `json:"role"`
}

func (o HtapNodeInfoResponseBodyNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapNodeInfoResponseBodyNodeList struct{}"
	}

	return strings.Join([]string{"HtapNodeInfoResponseBodyNodeList", string(data)}, " ")
}
