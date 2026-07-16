package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeNetwork 节点网络配置。
type NodeNetwork struct {

	// **参数解释**：vpc id。 **取值范围**：不涉及。
	Vpc *string `json:"vpc,omitempty"`

	// **参数解释**：子网id。 **取值范围**：不涉及。
	Subnet *string `json:"subnet,omitempty"`

	// **参数解释**：安全组id集合。
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
}

func (o NodeNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeNetwork struct{}"
	}

	return strings.Join([]string{"NodeNetwork", string(data)}, " ")
}
