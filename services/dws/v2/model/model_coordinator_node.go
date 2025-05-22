package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CoordinatorNode **参数解释**： CN节点详情。 **取值范围**： 不涉及。
type CoordinatorNode struct {

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 内网IP。 **取值范围**： 不涉及。
	PrivateIp string `json:"private_ip"`
}

func (o CoordinatorNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoordinatorNode struct{}"
	}

	return strings.Join([]string{"CoordinatorNode", string(data)}, " ")
}
