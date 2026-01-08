package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolNode
type PoolNode struct {

	// **取值范围**：后端服务器组id。  **取值范围**：不涉及
	Id string `json:"id"`

	// **取值范围**：后端服务器组名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **取值范围**：后端服务器组协议。  **取值范围**：不涉及
	Protocol string `json:"protocol"`

	// **取值范围**：后端服务器组类型。  **取值范围**：不涉及
	Type string `json:"type"`

	// **取值范围**：后端服务器组负载均衡算法。  **取值范围**：不涉及
	LbAlgorithm string `json:"lb_algorithm"`
}

func (o PoolNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolNode struct{}"
	}

	return strings.Join([]string{"PoolNode", string(data)}, " ")
}
