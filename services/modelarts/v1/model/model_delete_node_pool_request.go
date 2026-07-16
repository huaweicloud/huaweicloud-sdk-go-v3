package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodePoolRequest Request Object
type DeleteNodePoolRequest struct {

	// **参数解释**：资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：节点池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodepoolName string `json:"nodepool_name"`
}

func (o DeleteNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodePoolRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodePoolRequest", string(data)}, " ")
}
